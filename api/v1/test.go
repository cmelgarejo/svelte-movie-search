package v1

import (
	"fmt"
	"net/http"
	"os"
)

// Test entry point of the slsfn /v{X}/Test
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s - %s - %s", os.Getenv("PROVIDER_GOOGLE_KEY"), os.Getenv("PROVIDER_GOOGLE_SECRET"),
		os.Getenv("VERCEL_URL"))
}
