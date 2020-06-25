package auth

import (
	"net/http"
)

// Handler entry point of the slsfn /v{X}/[provider]
func Handler(w http.ResponseWriter, r *http.Request) {
	Begin(&w, r)
}
