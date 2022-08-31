package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CreatePostRequest struct {
	AuthorID    int    `json:"author_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreatePostResponse struct {
	ID int64 `json:"id"`
	SuccessPostResponse
}

type DetailPostResponse struct {
	PostResponse
}

type PostResponse struct {
	ID          int                `json:"id"`
	IsAuthor    bool               `json:"is_author"`
	Author      AuthorPostResponse `json:"author"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	CreatedAt   string             `json:"created_at"`
}
type AuthorPostResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	Program      string `json:"program"`
	Company      string `json:"company"`
	Batch        int    `json:"batch"`
	ProfileImage string `json:"profile_image"`
}

type SuccessPostResponse struct {
	Message string `json:"message"`
}
type ErrorPostResponse struct {
	Message string `json:"error"`
}

func (api *API) createPost(c *gin.Context) {
	req := CreatePostRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Invalid Request Body"})
		return
	}
	authorID, err := api.getUserIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Your ID cann't read"})
	}
	postID, err := api.postRepo.InserPost(authorID, req.Title, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorPostResponse{Message: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, CreatePostResponse{
		ID: postID,
		SuccessPostResponse: SuccessPostResponse{
			Message: "Post Created",
		},
	})
}

func (api *API) readPost(c *gin.Context) {
	var (
		postID int
		err    error
	)

	authorID := api.getUserIDAvoidPanic(c)

	if postID, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Invalid Post ID"})
		return
	}
	posts, err := api.postRepo.FetchPostByID(postID, authorID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, ErrorPostResponse{Message: "Internal server Error"})
		return
	}
	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, ErrorPostResponse{Message: "Post Not Found"})
		return
	}
	var (
		authorProgram, authorCompany string
		authorBatch                  int
	)
	if posts[0].AuthorProgram.Valid {
		authorProgram = posts[0].AuthorProgram.String
	}
	if posts[0].AuthorCompany.Valid {
		authorCompany = posts[0].AuthorCompany.String
	}
	if posts[0].AuthorBatch.Valid {
		authorBatch = int(posts[0].AuthorBatch.Int32)
	}

	c.JSON(http.StatusOK, DetailPostResponse{
		PostResponse{
			ID:       posts[0].ID,
			IsAuthor: posts[0].AuthorID == authorID,
			Author: AuthorPostResponse{
				ID:      posts[0].AuthorID,
				Name:    posts[0].AuthorName,
				Role:    posts[0].AuthorRole,
				Program: authorProgram,
				Company: authorCompany,
				Batch:   authorBatch,
			},
			Title:       posts[0].Title,
			Description: posts[0].Description,
			CreatedAt:   posts[0].CreatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func (api *API) getUserIDAvoidPanic(c *gin.Context) (authorID int) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("Recover from Panic")
		}
	}()
	authorID, _ = api.getUserIdFromToken(c)
	return
}
