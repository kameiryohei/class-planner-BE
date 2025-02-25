package middleware

import (
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	})
}

func CorsMiddleware() echo.MiddlewareFunc {
	var allowedOrigins []string
	if os.Getenv("GO_ENV") == "dev" {
		allowedOrigins = []string{"http://localhost:3000", os.Getenv("FE_URL")}
	} else {
		allowedOrigins = []string{os.Getenv("FE_URL")}
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	})
}

func CsrfMiddleware() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteDefaultMode,
	})
}

// JWTトークンが存在する場合のみ検証を行い、存在しない場合はリクエストを通過させるミドルウェア
func OptionalJwtMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("token")
			// トークンが存在する場合は検証
			if err == nil && cookie.Value != "" {
				// JWT検証を実行
				jwtMiddleware := echojwt.WithConfig(config)
				handler := jwtMiddleware(func(c echo.Context) error {
					// 検証成功時は何もせず次に進む
					return nil
				})
				// エラーが発生しても無視して次に進む
				_ = handler(c)
			}
			return next(c)
		}
	}
}
