package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go-web/errors"
	"go-web/models"
	"strconv"
)

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account": user.Account,
		"id":      strconv.FormatInt(user.Id, 10),
	})
	return token.SignedString([]byte("secret"))
}

func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("authorization")
		if tokenStr != "" {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.ErrUnauthorized()
				}
				return []byte("secret"), nil
			})
			if token.Valid {
				claims, _ := token.Claims.(jwt.MapClaims)
				userId := claims["id"].(string)
				if err := models.UserHandler.QueryLastByWhere(&models.User{}, "id = ?", []interface{}{userId}); err != nil {
					return errors.ErrUnauthorized()
				}
				c.Request().Header.Set("userId", userId)
				return next(c)
			}
			return errors.ErrUnauthorized()
		}
		return errors.ErrUnauthorized()
	}
}
