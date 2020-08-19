package auth

import (
	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src"
	"net/http"
)

// GQLPlaygroundHandler entry point of the slfn /v{X}/auth/[main]
func GQLPlaygroundHandler(w http.ResponseWriter, r *http.Request) {
	db := gen.NewDBFromEnvVars()
	defer db.Close()
	gen.GetHTTPVercel(src.New(db, nil), db, nil, w, r)
}
