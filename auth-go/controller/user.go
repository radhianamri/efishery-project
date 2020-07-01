package controller

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/radhianamri/efishery-project/auth-go/config"
	"github.com/radhianamri/efishery-project/auth-go/model"
)

type claimsResp struct {
	Name      interface{} `json:"name"`
	Phone     interface{} `json:"phone"`
	Role      interface{} `json:"role"`
	Timestamp interface{} `json:"timestamp"`
}

type login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// CreateUser attempts to create new user
func CreateUser(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return Bad(c)
	}

	if err := model.CreateUser(&u); err != nil {
		//TODO : switch response when record not found
		return Unprocessable(c, err.Error())
	}

	return OK(c, "A new user has been successfully registered.")
}

// LoginUser attempts to login user
func LoginUser(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		return Bad(c)
	}

	//check and retrieve user data
	if err := model.CheckLoginUser(&u); err != nil {
		//TODO : switch response when record not found
		return Unprocessable(c, err.Error())
	}

	//generate new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = u.Name
	claims["phone"] = u.Phone
	claims["role"] = u.Role
	claims["timestamp"] = time.Now().Format(time.RFC3339)

	newToken, err := token.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return Unprocessable(c, err.Error())
	}
	return Data(c, newToken)
}

func GetUserClaims(c echo.Context) error {
	authScheme := "Bearer"
	auth := c.Request().Header.Get("Authorization")
	l := len(authScheme)
	var userToken string
	if len(auth) > l+1 && auth[:l] == authScheme {
		userToken = auth[l+1:]
	}
	if userToken == "" {
		return Unauthorized(c)
	}
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return Unauthorized(c)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["phone"] != nil {
			return Data(c, claimsResp{
				claims["name"],
				claims["phone"],
				claims["role"],
				claims["timestamp"],
			})
		}
	}
	return Unauthorized(c)
}
