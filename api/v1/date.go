package v1

import (
	"fmt"
	"net/http"
	"time"
)

// Date entry point of the slfn /v{X}/date
func Date(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}
