package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CreatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type DetailPostResponse struct {
	PostResponse
}

type PostResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type ErrorPostResponse struct {
	Message string `json:"error"`
}

type CreatePostResponse struct {
	ID int64 `json:"id"`
	SuccessPostResponse
}

type SuccessPostResponse struct {
	Message string `json:"message"`
}

func (api *API) createPost(ctx *gin.Context) {
	var req = CreatePostRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Invalid Request Body"})
		return
	}

	authorID, err := api.getUserIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Your ID can't read"})
	}

	postID, err := api.postRepo.InsertPost(authorID, req.Title, req.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, CreatePostResponse{
		ID: postID,
		SuccessPostResponse: SuccessPostResponse{
			Message: "Post Created",
		},
	})
}

func (api *API) readPosts(ctx *gin.Context) {
	var (
		postID int
		err    error
	)
	authorID := api.getUserIDAvoidPanic(ctx)

	if postID, err = strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Invalid Post ID"})
		return
	}
	posts, err := api.postRepo.FetchPostByID(postID, authorID)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, ErrorPostResponse{Message: "Internal Server Error"})
		return
	}
	if len(posts) == 0 {
		ctx.JSON(http.StatusNotFound, ErrorPostResponse{Message: "Post Not Found"})
		return
	}

	ctx.JSON(http.StatusCreated, DetailPostResponse{
		PostResponse: PostResponse{
			ID:          postID,
			Title:       posts[0].Title,
			Description: posts[0].Description,
		},
	})
}
func (api *API) getUserIDAvoidPanic(ctx *gin.Context) (authorID int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover from panic")
		}
	}()

	authorID, _ = api.getUserIdFromToken(ctx)
	return
}
