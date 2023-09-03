package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/client"
	"github.com/magiclabs/magic-admin-go/token"
)

const authBearer = "Bearer"

var secret_key string

var magicSDK *client.API

func hello(w http.ResponseWriter, req *http.Request) {
	// Check whether or not DIDT exists in HTTP Header Request
	if !strings.HasPrefix(req.Header.Get("Authorization"), authBearer) {
		fmt.Fprintf(w, "Bearer token is required")
		return
	}

	// Retrieve DIDT token from HTTP Header Request
	didToken := req.Header.Get("Authorization")[len(authBearer)+1:]
	fmt.Printf("didToken: %s\n\n", didToken)

	// Create a Token instance to interact with the DID token
	tk, err := token.NewToken(didToken)
	if err != nil {
		fmt.Fprintf(w, "Malformed DID token error: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	// Validate the Token instance before using it
	if err := tk.Validate(); err != nil {
		fmt.Fprintf(w, "DID token failed validation: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "Hello %s\n", tk)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func endpoint(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "{\"d\": 123}")
}

func main() {

	secret_key = os.Getenv("MAGIC_SECRET")

	magicSDK = client.New(secret_key, magic.NewDefaultClient())

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/endpoint", endpoint)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
