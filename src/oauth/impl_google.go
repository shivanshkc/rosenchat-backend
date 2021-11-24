package oauth

import (
	"fmt"
	"rosenchat/src/configs"
)

var conf = configs.Get()

// implGoogle implements IOAuthProvider interface for Google.
type implGoogle struct{}

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
	panic("implement me")
}

func (i *implGoogle) Token2UserInfo(token string) (*UserInfoDTO, error) {
	panic("implement me")
}

func (i *implGoogle) init() {}
