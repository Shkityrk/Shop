package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// copyHeaders copies relevant headers from the incoming request to the outgoing request
func copyHeaders(from *http.Request, to *http.Request) {
	// Copy cookies
	for _, cookie := range from.Cookies() {
		to.AddCookie(cookie)
	}

	// Copy authorization header
	if auth := from.Header.Get("Authorization"); auth != "" {
		to.Header.Set("Authorization", auth)
	}

	// Copy other relevant headers
	if contentType := from.Header.Get("Content-Type"); contentType != "" {
		to.Header.Set("Content-Type", contentType)
	}

	if accept := from.Header.Get("Accept"); accept != "" {
		to.Header.Set("Accept", accept)
	}
}

// copyCookies copies cookies from HTTP response to Gin context
func copyCookies(from *http.Response, to *gin.Context) {
	for _, cookie := range from.Cookies() {
		http.SetCookie(to.Writer, cookie)
	}
}

