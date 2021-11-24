package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rosenchat/src/configs"
	"rosenchat/src/exception"
	"rosenchat/src/jwt"
	"rosenchat/src/logger"
)

var conf = configs.Get()
var log = logger.Get()

// implGoogle implements IOAuthProvider interface for Google.
type implGoogle struct {
	jwtManager jwt.IJWTManager
}

func (i *implGoogle) Name() string {
	return "google"
}

func (i *implGoogle) GetRedirectURL() string {
	return fmt.Sprintf(
		"%s?scope=%s&include_granted_scopes=true&response_type=code&redirect_uri=%s&client_id=%s",
		conf.GoogleOAuth.RedirectURL,
		conf.GoogleOAuth.Scopes,
		fmt.Sprintf("%s/api/auth/%s/callback", conf.GeneralOAuth.ServerCallbackBaseURL, i.Name()),
		conf.GoogleOAuth.ClientID,
	)
}

func (i *implGoogle) Code2Token(code string) (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"code":          code,
		"client_id":     conf.GoogleOAuth.ClientID,
		"client_secret": conf.GoogleOAuth.ClientSecret,
		"redirect_uri":  fmt.Sprintf("%s/api/auth/%s/callback", conf.GeneralOAuth.ServerCallbackBaseURL, i.Name()),
		"grant_type":    "authorization_code",
	})
	if err != nil {
		log.Errorf("Failed to marshal request body: %s", err.Error())
		return "", err
	}

	resp, err := http.Post(conf.GoogleOAuth.TokenEndpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Errorf("Failed to execute HTTP request: %s", err.Error())
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %s", err.Error())
		return "", err
	}

	if resp.StatusCode > 299 {
		log.Errorf("Endpoint returned unsuccessful status code. Body:\n%s", string(bodyBytes))
		return "", exception.InternalServerError()
	}

	bodyMap := map[string]interface{}{}
	if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
		log.Errorf("Failed to unmarshal result into map: %s", err.Error())
		return "", err
	}

	idToken, exists := bodyMap["id_token"]
	if !exists || idToken == nil {
		log.Errorf("Result did not contain an ID Token.")
		return "", exception.InternalServerError()
	}

	return idToken.(string), nil
}

func (i *implGoogle) Token2UserInfo(token string) (*UserInfoDTO, error) {
	claims, err := i.jwtManager.DecodeUnsafe(token)
	if err != nil {
		return nil, err
	}

	email, existsE := claims["email"]
	firstName, existsF := claims["given_name"]
	lastName, existsL := claims["family_name"]
	pictureLink := claims["picture"]

	if !existsE || !existsF || !existsL {
		log.Errorf("Insufficient information from provider.")
		return nil, exception.InternalServerError().AddMessages("insufficient info")
	}

	return &UserInfoDTO{
		Email:       email.(string),
		FirstName:   firstName.(string),
		LastName:    lastName.(string),
		PictureLink: pictureLink.(string),
	}, nil
}

func (i *implGoogle) init() {
	i.jwtManager = jwt.Get()
}
