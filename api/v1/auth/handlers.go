package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cmelgarejo/go-gql-server/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/markbates/goth/gothic"
)

// Claims JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Begin login with the auth provider
func Begin(w *http.ResponseWriter, r *http.Request) {
	// You have to add value context with provider name to get provider name in GetProviderName method
	r = addProviderToContext(r, r.URL.Query().Get(string(utils.ProjectContextKeys.ProviderCtxKey)))
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(*w, r); err != nil {
		gothic.BeginAuthHandler(*w, r)
	} else {
		fmt.Printf("user: %#v", gothUser)
	}
}

// Callback callback to complete auth provider flow
func Callback(w http.ResponseWriter, r *http.Request, cfg *utils.ServerConfig) {
	// You have to add value context with provider name to get provider name in GetProviderName method
	r = addProviderToContext(r, r.URL.Query().Get(string(utils.ProjectContextKeys.ProviderCtxKey)))
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		abortWithError(&w, http.StatusInternalServerError, err)
		return
	}
	// u, err := orm.FindUserByJWT(user.Email, user.Provider, user.UserID)
	// logger.Debugf("gothUser: %#v", user)
	if err != nil {
		// if u, err = orm.UpsertUserProfile(&user); err != nil {
		// 	logger.Errorf("[Auth.CallBack.UserLoggedIn.UpsertUserProfile.Error]: %v", err)
		// 	abortWithError(http.StatusInternalServerError, err)
		// }
	}
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
