package ui

import (
	"fmt"
	"net/http"
	"os"
)

func assets(w http.ResponseWriter, req *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(w, "Render assets error: %v!", err)
	}
	http.FileServer(http.Dir(cwd+"/static")).ServeHTTP(w, req)
}
