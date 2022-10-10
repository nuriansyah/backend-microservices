package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfile(t *testing.T) {
	r := gin.Default()
	r.Run(":8081")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/profile", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, "Get All")
}
