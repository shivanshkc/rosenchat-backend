package jwt

import (
	"github.com/golang-jwt/jwt"
)

// implGolangJWT implements IJWTManager using golang-jwt/jwt package.
type implGolangJWT struct {
	parser *jwt.Parser
}

func (i *implGolangJWT) DecodeUnsafe(token string) (map[string]interface{}, error) {
	claims := &jwt.MapClaims{}
	if _, _, err := i.parser.ParseUnverified(token, claims); err != nil {
		return nil, err
	}
	return *claims, nil
}

func (i *implGolangJWT) init() {
	i.parser = &jwt.Parser{SkipClaimsValidation: true}

}
