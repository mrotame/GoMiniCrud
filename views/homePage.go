package views

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Golang API V1.0.0")
}
