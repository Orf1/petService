package main

import (
	"fmt"
	"log"
	"net/http"
)

type petDatabase map[string]string

var (
	db petDatabase
)

func main() {
	db = petDatabase{
		"abcd": "Call John Doe at (805) 861-8822",
		"dcba": "Call Orfeas Magoulas at (805) 861-0195",
		"aaaa": "Call Shimon Navon at (111) 111-1111",
	}

	log.Println("Starting http server")

	getPetHandler := http.HandlerFunc(getPetHandler)
	http.Handle("/pet", getPetHandler)

	log.Println("Listening on http://localhost:80")

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Println("Unable to start http server")
		log.Fatalln(err)
	}
}

func getPetHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, present := query["id"]

	if !present || len(id) == 0 {
		_, _ = fmt.Fprint(w, "id missing from query")
		return
	}

	if db[id[0]] != "" {
		_, _ = fmt.Fprint(w, db[id[0]])
	} else {
		_, _ = fmt.Fprint(w, "pet not found in database")
	}
}
