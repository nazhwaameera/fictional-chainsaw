package helpers

import "github.com/gin-gonic/gin"

// memungkinkan client mengirimkan request body melalui JSON ataupun melalui form
func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
