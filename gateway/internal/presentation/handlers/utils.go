package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// copyCookies copies cookies from HTTP response to Gin context
func copyCookies(from *http.Response, to *gin.Context) {
	for _, cookie := range from.Cookies() {
		http.SetCookie(to.Writer, cookie)
	}
}

