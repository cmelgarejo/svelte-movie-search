package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cmelgarejo/go-gql-server/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

// Claims JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// CallbackHandler entry point of the slsfn /v{X}/date
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	cfg := &utils.ServerConfig{
		JWT: utils.JWTConfig{
			Algorithm: "HS512",
			Secret:    "328c69c995a14a7f944623af20396c2c6f997ae806df4cf08eaf9f569cf8f8ad",
		},
	}
	callback(w, r, cfg)
}

// callback to complete auth provider flow
func callback(w http.ResponseWriter, r *http.Request, cfg *utils.ServerConfig) {
	// You have to add value context with provider name to get provider name in GetProviderName method
	r = addProviderToContext(r, "google")
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		abortWithError(&w, http.StatusInternalServerError, err)
		return
	}
	// u, err := orm.FindUserByJWT(user.Email, user.Provider, user.UserID)
	// logger.Debugf("gothUser: %#v", user)
	// if err != nil {
	// if u, err = orm.UpsertUserProfile(&user); err != nil {
	// 	logger.Errorf("[Auth.CallBack.UserLoggedIn.UpsertUserProfile.Error]: %v", err)
	// 	abortWithError(http.StatusInternalServerError, err)
	// }
	// }
	// logger.Debug("[Auth.CallBack.UserLoggedIn.USER]: ", u)
	// logger.Debug("[Auth.CallBack.UserLoggedIn]: ", u.ID)
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.JWT.Algorithm), Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			Id:        user.UserID,
			Issuer:    user.Provider,
			IssuedAt:  time.Now().UTC().Unix(),
			NotBefore: time.Now().UTC().Unix(),
			ExpiresAt: user.ExpiresAt.UTC().Unix(),
		},
	})
	token, err := jwtToken.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		// logger.Error("[Auth.Callback.JWT] error: ", err)
		abortWithError(&w, http.StatusInternalServerError, err)
		return
	}
	// logger.Debug("token: ", token)
	json := map[string]interface{}{
		"type":          "Bearer",
		"token":         token,
		"refresh_token": user.RefreshToken,
	}
	JSON(&w, http.StatusOK, json)
}

// // Logout logs out of the auth provider
// func Logout() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		r = addProviderToContext(c, c.Param(string(utils.ProjectContextKeys.ProviderCtxKey)))
// 		gothic.Logout(w, r)
// 		w.Header().Set("Location", "/")
// 		w.WriteHeader(http.StatusTemporaryRedirect)
// 	}
// }

func abortWithError(w *http.ResponseWriter, sc int, err error) {
	(*w).WriteHeader(sc)
}

// JSON friday 13th
func JSON(w *http.ResponseWriter, sc int, j interface{}) (err error) {
	b, err := json.Marshal(j)
	fmt.Fprint(*w, b)
	return err
}

func addProviderToContext(r *http.Request, value interface{}) *http.Request {
	// gothic.Store = sessions.NewCookieStore([]byte("<your secret here>"))
	goth.UseProviders(google.New(utils.MustGet("PROVIDER_GOOGLE_KEY"), utils.MustGet("PROVIDER_GOOGLE_SECRET"),
		"https://svelte-movie-search-git-serverless.cmelgarejo.now.sh/api/v1/auth/callback", "email", "profile", "openid"))
	return r.WithContext(context.WithValue(r.Context(),
		string(utils.ProjectContextKeys.GothicProviderCtxKey), value))
}
