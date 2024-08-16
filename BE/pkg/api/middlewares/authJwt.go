package middlewares

import (
	"errors"
	"github.com/Hemtrakan/todolist_connection_mongodb/pkg/shared/models"
	"github.com/Hemtrakan/todolist_connection_mongodb/pkg/shared/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func AuthJwt(permission []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientTime := c.Request.Header.Get("X-Timestamp")
		timeisvalid := validateTimestamp(clientTime)
		if !timeisvalid {
			c.JSON(http.StatusRequestTimeout, gin.H{"status": "INVALID_TIMESTAMP", "message": "Invalid timestamp detected"})
			c.Abort()
			return
		}

		token := c.Request.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authentication not found"})
			c.Abort()
			return
		}
		claims, err := validateToken(token, permission)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": err.Error()})
			c.Abort()
			return
		}
		c.Set("user", claims)
		c.Next()

	}
}

func validateToken(signedToken string, permission []string) (*models.Claims, error) {
	jwtKey := []byte(viper.GetString("authen.secretKey"))
	jwtToken, err := jwt.ParseWithClaims(
		signedToken,
		&models.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*models.Claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	if !validatePermission(claims.Permissions, permission) {
		err = errors.New("not has permission")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, nil
}

func validatePermission(scopes, permissions []string) bool {
	hasPermission := []string{}
	for _, permission := range permissions {
		for _, scope := range scopes {
			if permission == scope {
				hasPermission = append(hasPermission, permission)
				break
			}
		}
	}
	return len(hasPermission) == len(permissions)
}

func validateTimestamp(clientTime string) bool {
	requestTime, err := time.Parse(string(time.RFC3339), clientTime)
	if err != nil {
		return true
	}
	currentTime := time.Now().In(time.Local)
	startServerTime := currentTime.Add(-30 * time.Minute)
	endServerTime := currentTime.Add(30 * time.Minute)
	return utils.InTimeSpan(startServerTime, endServerTime, requestTime)
}
