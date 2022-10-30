package tests

import (
	"gin-example/app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controllers.Login(c)
	assert.Equal(t, 200, w.Code)

}
