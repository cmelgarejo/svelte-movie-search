package date

import (
	"fmt"
	"net/http"
	"time"
)

// Handler entry point of the slsfn /v{X}/date
func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}
