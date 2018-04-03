package auth

import (
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/keiwi/utils/models"
)

// signinKey set up a global string for our secret
var signinKey = []byte("kdsadsndadsafs")

// JwtMiddleware handler for jwt tokens
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return signinKey, nil
	},
	UserProperty:  "user",
	Debug:         false,
	SigningMethod: jwt.SigningMethodHS256,
})

// GetToken create a jwt token with user claims
func GetToken(user *models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	signedToken, _ := token.SignedString(signinKey)
	return signedToken
}

// GetJSONToken create a JSON token string
func GetJSONToken(user *models.User) string {
	token := GetToken(user)
	jsontoken := "{\"id_token\": \"" + token + "\"}"
	return jsontoken
}

// GetUserClaimsFromContext return "user" claims as a map from request
func GetUserClaimsFromContext(req *http.Request) map[string]interface{} {
	token := req.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims
}
