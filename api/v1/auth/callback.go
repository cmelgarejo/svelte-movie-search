package auth

import (
	"net/http"
)

// Handler entry point of the slsfn /v{X}/date
func Handler(w http.ResponseWriter, r *http.Request) {
	Callback(&w, r)
}
