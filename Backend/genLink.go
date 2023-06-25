package main

import (
	"fmt"
	"net/http"

	"lado.one/linkGen/linkLookup"
	"lado.one/linkGen/writeLink"
)

func reqHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/shrt" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		query := r.URL.Query().Get("q")
		//fmt.Fprintf(w, "GET request successful, query: %s\n", query)
		dest_data, err := linkLookup.GetDest(query)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(dest_data)

		if len(dest_data.Rows) == 0 {
			fmt.Fprint(w, "No such link found.")
			return
		} else if len(dest_data.Rows) > 1 {
			fmt.Fprint(w, "An error occurred: Hash crash")
			return
		} else {
			http.Redirect(w, r, dest_data.Rows[0].Value, http.StatusSeeOther)
		}

	case "POST":
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		dest := r.PostFormValue("dest")

		// Check if dest is aleady in the database
		ref := writeLink.GetSHA1Hash(dest)[:8]
		dest_data, err := linkLookup.GetDest(ref)
		if err != nil {
			fmt.Println(err)
		}

		if len(dest_data.Rows) == 0 {
			writeLink.AddNew(dest)
			fmt.Fprint(w, "Link added:\n Use http://localhost:3000/shrt?q="+ref+" to access it.")
		} else {
			fmt.Fprint(w, "Link already exists:\n Use http://localhost:3000/shrt?q="+ref+" to access it.")
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/shrt", reqHandler)
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}
