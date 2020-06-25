package auth

import (
	"github.com/cmelgarejo/go-gql-server/pkg/utils"
	"net/http"
)

// Handler entry point of the slsfn /v{X}/date
func Handler(w http.ResponseWriter, r *http.Request) {
	Callback(&w, r, &utils.ServerConfig{
		JWT: utils.JWTConfig{
			Algorithm: "HS512",
			Secret:    "328c69c995a14a7f944623af20396c2c6f997ae806df4cf08eaf9f569cf8f8ad",
		},
	})
}
