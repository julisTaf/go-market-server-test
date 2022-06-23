package middlewares

import (
	"Go-market-test/pkg/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthzNoHeader(t *testing.T) {
	router := gin.Default()
	router.Use(Authz())

	router.GET("/api/protected/profile", controllers.Profile)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/protected/profile", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 403, w.Code)
}

func TestAuthzInvalidTokenFormat(t *testing.T) {
	router := gin.Default()
	router.Use(Authz())

	router.GET("/api/protected/profile", controllers.Profile)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/protected/profile", nil)
	req.Header.Add("Authorization", "test")

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestAuthzInvalidToken(t *testing.T) {
	invalidToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	router := gin.Default()
	router.Use(Authz())

	router.GET("/api/protected/profile", controllers.Profile)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/protected/profile", nil)
	req.Header.Add("Authorization", invalidToken)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}
