package auth

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	SecretKey string
}

func (m *AuthMiddleware) ValidateToken(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
	if token == "" {
		ctx.Output.JSON(map[string]interface{}{
			"error": "Token no proporcionado",
		}, 401, false, false)
		return
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.SecretKey), nil
	})

	if err != nil {
		ctx.Output.JSON(map[string]interface{}{
			"error": "Token inválido",
		}, 401, false, false)
		return
	}

	// Agregar claims al contexto
	ctx.Input.SetData("user_id", claims["user_id"])
	ctx.Input.SetData("roles", claims["roles"])
}
