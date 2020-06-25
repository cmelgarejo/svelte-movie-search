package date

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().String()
	fmt.Fprintf(w, currentTime)
}
