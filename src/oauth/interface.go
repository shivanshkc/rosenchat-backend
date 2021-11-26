package oauth

import (
	"context"
	"rosenchat/src/database"
	"sync"
)

// IOAuthProvider represents a generic OAuth provider.
type IOAuthProvider interface {
	// Name gives the name of the provider.
	Name(ctx context.Context) string

	// GetRedirectURL returns the URL to the auth page of the provider.
	GetRedirectURL(ctx context.Context) string

	// Code2Token converts the auth code to identity token.
	Code2Token(ctx context.Context, code string) (string, error)

	// Token2UserInfo converts the identity token into the user's info.
	Token2UserInfo(ctx context.Context, token string) (*database.UserInfoDTO, error)

	// init can be used to initialize the implementation.
	init()
}

var googleProviderOnce = &sync.Once{}
var googleProviderSingleton IOAuthProvider

// GetGoogleOAuthProvider returns the Google OAuthProvider singleton.
func GetGoogleOAuthProvider() IOAuthProvider {
	googleProviderOnce.Do(func() {
		googleProviderSingleton = &implGoogle{}
		googleProviderSingleton.init()
	})

	return googleProviderSingleton
}
