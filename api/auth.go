package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginReqBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginSuccessResponse struct {
	Token string `json:"token"`
}

var jwtKey = []byte("key")

type Claims struct {
	id    int    `json:"id"`
	email string `json:"email"`
	jwt.StandardClaims
}

func (api API) genereteJWT(useId *int) (string, error) {
	expTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		id: *useId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//tokenString, err := token.SigningString()
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, err
}
func (api *API) getUserIdFromToken(c *gin.Context) (int, error) {
	tokenString := c.GetHeader("Authorization")[(len("Bearer ")):]
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return -1, err
	}
	if token.Valid {
		claim := token.Claims.(*Claims)
		return claim.id, nil
	} else {
		return -1, errors.New("Invalid Tokens")
	}
}

func (api *API) login(c *gin.Context) {
	var loginReq LoginReqBody
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mhsId, err := api.mhsRepo.Login(loginReq.Email, loginReq.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	tokenString, err := api.genereteJWT(mhsId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dosenId, err := api.dosenRepo.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	tokenString, err = api.genereteJWT(dosenId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginSuccessResponse{Token: tokenString})
}
