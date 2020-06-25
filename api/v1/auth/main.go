package auth

import (
	"context"
	"net/http"

	"github.com/cmelgarejo/go-gql-server/pkg/utils"
)

func addProviderToContext(r *http.Request, value interface{}) *http.Request {
	return r.WithContext(context.WithValue(r.Context(),
		string(utils.ProjectContextKeys.GothicProviderCtxKey), value))
}
