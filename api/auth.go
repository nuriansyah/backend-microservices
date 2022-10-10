package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nuriansyah/log-mbkm-unpas/helper"
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

type RegisterReqBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,lowercase,oneof=dosen mahasiswa"`
}
type RegisterSuccessResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

var jwtKey = []byte("key")

type Claims struct {
	Id    int
	Email string
	Role  string
	jwt.StandardClaims
}

func (api *API) getUserIdFromToken(c *gin.Context) (int, error) {
	tokenString := c.GetHeader("Authorization")[(len("Bearer ")):]
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return -1, err
	}

	if token.Valid {
		claim := token.Claims.(*Claims)
		return claim.Id, nil

	} else {
		return -1, errors.New("invalid token")
	}
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, err
}

func (api API) generateJWT(userId *int, role *string) (string, error) {
	expTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		Id:   *userId,
		Role: *role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func (api *API) register(c *gin.Context) {
	var input RegisterReqBody
	err := c.BindJSON(&input)
	var ve validator.ValidationErrors

	if err != nil {
		if errors.As(err, &ve) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"errors": helper.GetErrorMessage(ve)},
			)
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	userId, responseCode, err := api.userRepo.InsertUser(input.Name, input.Email, input.Password, input.Role)
	if err != nil {
		c.AbortWithStatusJSON(responseCode, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := api.generateJWT(&userId, &input.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RegisterSuccessResponse{Message: "success", Token: tokenString})
}

func (api *API) login(c *gin.Context) {
	var loginReq LoginReqBody
	err := c.BindJSON(&loginReq)
	var ve validator.ValidationErrors

	if err != nil {
		if errors.As(err, &ve) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"errors": helper.GetErrorMessage(ve)},
			)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	userId, err := api.userRepo.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	role, err := api.userRepo.GetUserRole(*userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := api.generateJWT(userId, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginSuccessResponse{Token: tokenString})
}
