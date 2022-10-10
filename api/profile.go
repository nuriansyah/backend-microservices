package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DetailProfileRequest struct {
	Nrp     string `json:"nrp"`
	Prodi   string `json:"prodi"`
	Program string `json:"program"`
	Company string `json:"company"`
	Batch   int    `json:"bacth"`
}

type ResponseDetailProfile struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Nrp     string `json:"nrp"`
	Program string `json:"program"`
	Company string `json:"company"`
	Batch   int    `json:"batch"`
}
type Response struct {
	Message string `json:"message"`
}

func (api *API) getProfile(ctx *gin.Context) {
	userID, err := api.getUserIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, Response{"Unauthorized"})
		return
	}
	user, err := api.userRepo.GetUserData(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Message: err.Error()})
		return
	}
	var (
		userCompany string
		userBatch   int
	)

	//if user.Avatar != nil {
	//	userAvatar = *user.Avatar
	//}

	if user.Company != nil {
		userCompany = *user.Company
	}

	if user.Batch != nil {
		userBatch = *user.Batch
	}
	ctx.JSON(http.StatusOK, ResponseDetailProfile{
		ID:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Nrp:     user.Nrp,
		Program: user.Program,
		Company: userCompany,
		Batch:   userBatch,
	})
}

func (api *API) updateProfile(ctx *gin.Context) {
	var request DetailProfileRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Message: "Invalid Request"})
		return
	}

	userID, err := api.getUserIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, Response{"Unauthorized"})
		return
	}

	if err := api.userRepo.UpdateDetailDataUser(userID, request.Batch, request.Nrp, request.Prodi, request.Program, request.Company); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{Message: "Successfully Updated"})
}
