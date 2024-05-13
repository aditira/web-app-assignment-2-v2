package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil {
			ctx.AbortWithStatus(http.StatusSeeOther)
			return
		}

		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("id", claims.UserID)
		ctx.Next()
	})
}
