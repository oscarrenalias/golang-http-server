package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func writeResponse(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling request!")

	// is there a custom message?
	message := os.Getenv("MESSAGE")
	if message != "" {
		fmt.Println("Using custom message")
	} else {
		message = ""
	}

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}

	message = hostName + "\n" + message + "\n" + string(requestDump)

	w.Write([]byte(message))
}

func main() {
	port := ":8080"

	http.HandleFunc("/", writeResponse)
	fmt.Println("Listening in port " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
