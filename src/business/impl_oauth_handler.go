package business

import (
	"fmt"
	"net/http"
	"rosenchat/src/configs"
	"rosenchat/src/database"
	"rosenchat/src/exception"
	"rosenchat/src/oauth"
	"rosenchat/src/utils/hashutils"
	"rosenchat/src/utils/httputils"
)

var conf = configs.Get()

// implOAuthHandler implements IOAuthHandler.
type implOAuthHandler struct {
	providerMap map[string]oauth.IOAuthProvider
	userInfoDB  database.IUserInfoDB
}

func (i *implOAuthHandler) Redirect(provider string, writer http.ResponseWriter) (*ResponseDTO, error) {
	oAuthProvider, exists := i.providerMap[provider]
	if !exists {
		return nil, exception.ProviderNotFound()
	}

	respHeaders := map[string]string{"location": oAuthProvider.GetRedirectURL()}
	httputils.WriteJSON(writer, nil, respHeaders, http.StatusFound)

	return nil, nil
}

func (i *implOAuthHandler) HandleCallback(provider string, code string, writer http.ResponseWriter) (*ResponseDTO, error) {
	clientCallbackURL := conf.GeneralOAuth.ClientCallbackURL
	excProvider := exception.ProviderNotFound()

	oAuthProvider, exists := i.providerMap[provider]
	if !exists {
		redirectURL := fmt.Sprintf("%s?error=%s", clientCallbackURL, excProvider.Error())
		respHeaders := map[string]string{"location": redirectURL}
		httputils.WriteJSON(writer, nil, respHeaders, excProvider.StatusCode)
	}

	token, err := oAuthProvider.Code2Token(code)
	if err != nil {
		redirectURL := fmt.Sprintf("%s?error=%s", clientCallbackURL, excProvider.Error())
		respHeaders := map[string]string{"location": redirectURL}
		httputils.WriteJSON(writer, nil, respHeaders, excProvider.StatusCode)
	}

	userInfo, err := oAuthProvider.Token2UserInfo(token)
	if err != nil {
		redirectURL := fmt.Sprintf("%s?error=%s", clientCallbackURL, excProvider.Error())
		respHeaders := map[string]string{"location": redirectURL}
		httputils.WriteJSON(writer, nil, respHeaders, excProvider.StatusCode)
	}

	userInfo.ID = hashutils.SHA256Hex(userInfo.Email)

	go func() { _ = i.userInfoDB.PutUserInfo(userInfo) }()

	redirectURL := fmt.Sprintf("%s?id_token=%s&provider=%s", clientCallbackURL, token, provider)
	respHeaders := map[string]string{"location": redirectURL}
	httputils.WriteJSON(writer, nil, respHeaders, http.StatusFound)

	return nil, nil
}

func (i *implOAuthHandler) init() {
	googleProvider := oauth.GetGoogleOAuthProvider()
	i.providerMap = map[string]oauth.IOAuthProvider{
		googleProvider.Name(): googleProvider,
	}

	i.userInfoDB = database.Get()
}
