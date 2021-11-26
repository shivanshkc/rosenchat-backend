package business

import (
	"context"
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

func (i *implOAuthHandler) Redirect(ctx context.Context, provider string, writer http.ResponseWriter) (*ResponseDTO, error) {
	oAuthProvider, exists := i.providerMap[provider]
	if !exists {
		return nil, exception.ProviderNotFound()
	}

	respHeaders := map[string]string{"location": oAuthProvider.GetRedirectURL(ctx)}
	httputils.WriteJSON(ctx, writer, nil, respHeaders, http.StatusFound)

	return nil, nil
}

func (i *implOAuthHandler) HandleCallback(ctx context.Context, provider string, code string, writer http.ResponseWriter) (*ResponseDTO, error) {
	clientCallbackURL := conf.GeneralOAuth.ClientCallbackURL
	excProvider := exception.ProviderNotFound()

	oAuthProvider, exists := i.providerMap[provider]
	if !exists {
		redirectURL := fmt.Sprintf("%s?error=%s", clientCallbackURL, excProvider.Error())
		respHeaders := map[string]string{"location": redirectURL}
		httputils.WriteJSON(ctx, writer, nil, respHeaders, excProvider.StatusCode)
	}

	token, err := oAuthProvider.Code2Token(ctx, code)
	if err != nil {
		redirectURL := fmt.Sprintf("%s?error=%s", clientCallbackURL, excProvider.Error())
		respHeaders := map[string]string{"location": redirectURL}
		httputils.WriteJSON(ctx, writer, nil, respHeaders, excProvider.StatusCode)
	}

	userInfo, err := oAuthProvider.Token2UserInfo(ctx, token)
	if err != nil {
		redirectURL := fmt.Sprintf("%s?error=%s", clientCallbackURL, excProvider.Error())
		respHeaders := map[string]string{"location": redirectURL}
		httputils.WriteJSON(ctx, writer, nil, respHeaders, excProvider.StatusCode)
	}

	userInfo.ID = hashutils.SHA256Hex(userInfo.Email)

	go func() { _ = i.userInfoDB.PutUserInfo(ctx, userInfo) }()

	redirectURL := fmt.Sprintf("%s?id_token=%s&provider=%s", clientCallbackURL, token, provider)
	respHeaders := map[string]string{"location": redirectURL}
	httputils.WriteJSON(ctx, writer, nil, respHeaders, http.StatusFound)

	return nil, nil
}

func (i *implOAuthHandler) init() {
	googleProvider := oauth.GetGoogleOAuthProvider()
	i.providerMap = map[string]oauth.IOAuthProvider{
		googleProvider.Name(context.Background()): googleProvider,
	}

	i.userInfoDB = database.Get()
}
