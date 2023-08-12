package security

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"etentnode-api/app/entity"
	"etentnode-api/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(user entity.User, expire int) (string, error) {
	var mySigningKey = []byte("98hbun98h")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["phone_number"] = user.PhoneNumber
	claims["role_id"] = user.RoleID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 8766).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			resp := helper.Response("Unauthorized", http.StatusUnauthorized, err)

			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}
		user_id, err := ExtractTokenID(c)

		c.Set("user_id", user_id)
		c.Next()
	}
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("98hbun98h"), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (float64, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("98hbun98h"), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["id"].(float64), nil
	}
	return 0, nil
}
