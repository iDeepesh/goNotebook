package state

import (
	"fmt"
	"net/http"
)

//ServerToParseInputs - A server that parses url/form inputs and returns them
func ServerToParseInputs() {
	parser := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for k, v := range r.Form {
			fmt.Fprintln(w, "Key: "+k+" Value: "+v[0])
		}

		fmt.Fprintln(w, "Done printing form values.")
	}

	http.HandleFunc("/parse", parser)
	http.ListenAndServe(":5080", nil)
}
