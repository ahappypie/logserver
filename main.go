package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", RootHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	//var prettyJSON bytes.Buffer
	//error := json.Indent(&prettyJSON, dump, "", "\t")
	//if error != nil {
	//	log.Println("JSON parse error: ", error)
	//}

	log.Println(string(dump) + "\n")
}
