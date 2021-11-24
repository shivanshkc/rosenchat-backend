package business

import (
	"net/http"
	"rosenchat/src/exception"
	"rosenchat/src/oauth"
	"rosenchat/src/utils/httputils"
)

// implOAuthHandler implements IOAuthHandler.
type implOAuthHandler struct {
	providerMap map[string]oauth.IOAuthProvider
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
	panic("implement me")
}

func (i *implOAuthHandler) init() {
	googleProvider := oauth.GetGoogleOAuthProvider()
	i.providerMap = map[string]oauth.IOAuthProvider{
		googleProvider.Name(): googleProvider,
	}
}
