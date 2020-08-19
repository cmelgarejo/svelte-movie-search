package auth

import (
	"net/http"

	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src"
)

// GraphQLHandler entry point of the slfn /v{X}/auth/[main]
func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	db := gen.NewDBFromEnvVars()
	defer db.Close()
	gen.GetHTTPVercel(src.New(db, nil), db, nil, w, r)
}
