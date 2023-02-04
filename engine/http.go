package engine

import (
	"fmt"
	"net/http"
)

func HandleRequestMux(w http.ResponseWriter, r *http.Request) {
	// Handle request
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "form.html")
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("name")
		occupation := r.FormValue("occupation")

		fmt.Fprintf(w, "%s is a %s\n", name, occupation)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func HandleGetRequest() {
	// Handle get request
}

func HandlePutRequest() {
	// Handle put request
}

func HandleDeleteRequest() {
	// Handle delete request
}

func HandlePostRequest() {
	// Handle post request
}
