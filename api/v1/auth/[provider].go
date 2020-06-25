package auth

import (
	"fmt"
	"net/http"
	"time"
)

// Handler entry point of the slsfn /v{X}/[provider]
func Handler(w http.ResponseWriter, r *http.Request) {
	Begin(&w, r)
}
