package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func getName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params["PARAM"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + param + "!"))
}

func getBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func postData(w http.ResponseWriter, r *http.Request) {
	d, _ := io.ReadAll(r.Body)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I got message\n" + string(d)))
}

func postHeaders(w http.ResponseWriter, r *http.Request) {
	sum := 0
	for i, v := range r.Header {
		if strings.ToLower(i) == "a" || strings.ToLower(i) == "b" {
			tmp, _ := strconv.Atoi(v[0])
			sum += tmp
		}
	}
	w.Header().Add("a+b", strconv.Itoa(sum))
	w.WriteHeader(http.StatusOK)

}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", getName).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBad).Methods(http.MethodGet)

	router.HandleFunc("/data", postData).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
