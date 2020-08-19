package auth

import (
	"net/http"

	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src/auth"
)

// GoogleHandler entry point of the slfn /auth/google
func GoogleHandler(w http.ResponseWriter, r *http.Request) {
	db := gen.NewDBFromEnvVars()
	defer db.Close()
	auth.BeginVercel("google", w, r)
}
