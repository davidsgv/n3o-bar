package middlewares

import (
	"log"
	"net/http"

	service "github.com/davidsgv/n3o-bar/Service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := ""
		if len(authHeader) >= len(BEARER_SCHEMA) {
			tokenString = authHeader[len(BEARER_SCHEMA):]
		}

		token, err := service.NewAuthorizationService().ValidateToken(tokenString)

		if err == nil && token.Valid {
			//si no se especifican los roles
			//todos los roles tienen permiso
			if len(roles) == 0 {
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			for i := 0; i < len(roles); i++ {
				if roles[i] == claims["rol"] {
					// 	log.Println("Claims[Usuario]: ", claims["usuario"])
					// 	log.Println("Claims[IdUsuario]: ", claims["usuarioId"])
					// 	log.Println("Claims[Rol]: ", claims["rol"])
					// 	log.Println("Claims[Issuer]: ", claims["iss"])
					// 	log.Println("Claims[IssuedAt]: ", claims["iat"])
					// 	log.Println("Claims[ExpiresAt]: ", claims["exp"])
					return
				}
			}
		}

		log.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
