package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", PublicIPHandler)

	log.Println("Listening... on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func PublicIPHandler(w http.ResponseWriter, _ *http.Request) {
	resp, err := http.Get("https://ifconfig.me/all.json")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(body))
}
