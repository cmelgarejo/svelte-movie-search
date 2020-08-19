package auth

import (
	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src"
	"net/http"
)

// MigrateHandler entry point of the slfn /v{X}/auth/[main]
func MigrateHandler(w http.ResponseWriter, r *http.Request) {
	db := gen.NewDBFromEnvVars()
	defer db.Close()
	gen.GetHTTPVercel(src.New(db, nil), db, src.GetMigrations(db), w, r)
}
