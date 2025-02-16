package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Anton-Kraev/avito-test-assignment/internal/lib/tokens"
)

func JWTMiddlewareExcludeAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/api/auth" {
			return next(c)
		}

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		tokenString := authHeader[7:]

		userID, err := tokens.ValidateJWT(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("userID", userID)

		return next(c)
	}
}
