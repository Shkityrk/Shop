package http

import "net/http"

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

	// Copy X-User-Id header (set by auth middleware)
	if userID := from.Header.Get("X-User-Id"); userID != "" {
		to.Header.Set("X-User-Id", userID)
	}

	// Copy other relevant headers
	if contentType := from.Header.Get("Content-Type"); contentType != "" {
		to.Header.Set("Content-Type", contentType)
	}

	if accept := from.Header.Get("Accept"); accept != "" {
		to.Header.Set("Accept", accept)
	}
}
