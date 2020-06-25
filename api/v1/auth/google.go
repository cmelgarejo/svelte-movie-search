package auth

import (
	"fmt"
	"net/http"

	// "github.com/cmelgarejo/go-gql-server/pkg/utils"
	"github.com/markbates/goth/gothic"
)

// ProviderHandler entry point of the slsfn /v{X}/[provider]
func ProviderHandler(w http.ResponseWriter, r *http.Request) {
	// You have to add value context with provider name to get provider name in GetProviderName method
	r = addProviderToContext(r, "google")
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err != nil {
		gothic.BeginAuthHandler(w, r)
	} else {
		fmt.Printf("user: %#v", gothUser)
	}
}
