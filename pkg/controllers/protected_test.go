package controllers

import (
	"Go-market-test/pkg/database"
	models "Go-market-test/pkg/user"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProfileNotFound(t *testing.T) {
	var profile models.User

	err := database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&models.User{})

	request, err := http.NewRequest("GET", "/api/protected/profile", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	c.Set("email", "notfound@email.com")

	Profile(c)

	err = json.Unmarshal(w.Body.Bytes(), &profile)
	assert.NoError(t, err)

	assert.Equal(t, 404, w.Code)

	database.GlobalDB.Unscoped().Where("email = ?", "jwt@email.com").Delete(&models.User{})
}
