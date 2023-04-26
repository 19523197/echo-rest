package auth

import (
	"belajar-echo/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

const AccessTokenCookieName = "X-Member"

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET_KEY")
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateTokenAndCookies(user *model.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(AccessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)

	return nil
}

func generateAccessToken(user *model.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Hour * 1)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(user *model.User, expire time.Time, secret []byte) (string, time.Time, error) {
	claims := Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(token)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expire, nil
}

func setTokenCookie(name, token string, expire time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expire
	cookie.Path = "/"

	cookie.HttpOnly = true

	c.SetCookie(cookie)

}

func setUserCookie(user *model.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Username
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func JWTErrorChecker(err error, c echo.Context) error {
	return c.Redirect(401, c.Echo().Reverse("SignIn"))
}
