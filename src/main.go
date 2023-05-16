package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	cert := os.Getenv("CERT_CHAIN")
	if cert == "" {
		cert = "none"
		fmt.Fprintf(w, "none\n")
		return
	} else {
		decoded, err := base64.StdEncoding.DecodeString(cert)
		if err != nil {
			fmt.Println("decode error:", err)
			fmt.Fprintf(w, "none\n")
			return
		}
		fmt.Println("decoded string")
		fmt.Fprintf(w, "%s", string(decoded))
	}
}

func main() {

	http.HandleFunc("/", hello)
	fmt.Println("serving on port 8080")
	http.ListenAndServe(":8080", nil)
}
