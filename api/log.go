package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogRequestBody struct {
	ID        int    `json:"id"`
	Log       string `json:"log"`
	CreatedAt string `json:"created_at"`
}
type LogResponseBody struct {
}

type CreatePostResponse struct {
	ID int64 `json:"id"`
	SuccessPostResponse
}

type SuccessPostResponse struct {
	Message string `json:"message"`
}

type ErrorPostResponse struct {
	Message string `json:"error"`
}

func (api *API) createPost(c *gin.Context) {
	request := LogRequestBody{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Invalid Message Body"})
	}
	mhsID, err := api.getUserIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Your ID cann't read"})
	}
	logID, err := api.logRepo.InsertLog(mhsID, request.Log, request.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorPostResponse{Message: "Internal Server Error"})
	}
	c.JSON(http.StatusCreated,
		CreatePostResponse{
			ID: logID,
			SuccessPostResponse: SuccessPostResponse{
				Message: "Post Created",
			},
		})
}
