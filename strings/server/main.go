package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/idwall/desafios/strings/wrapper"
)

const DefaultLimit = 40

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/wrap", PostWrap)

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}

func PostWrap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is accepted", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q := r.URL.Query()
	justify := q.Get("justify") == "true"
	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		limit = DefaultLimit
	}

	result := wrapper.Wrap(string(body), limit, justify)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
