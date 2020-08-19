package auth

import (
	"net/http"

	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src/auth"
)

// CallbackHandler entry point of the slfn auth/callback/google
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	db := gen.NewDBFromEnvVars()
	defer db.Close()
	auth.CallbackHandlerVercel(db, "google")(w, r)
	// auth.CallbackHandler(db)(w, r)
}
