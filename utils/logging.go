package utils

import (
	"net/http"
	// "io/ioutil"
	// "encoding/json"
	"fmt"
	// "bytes"
)

type Body struct {
	OperationName 	string
	Variables		string
	Query			string
}

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Logged connection from %s", r.RemoteAddr)
		// Read the content
		// var bodyBytes []byte
		// if r.Body != nil {
		// 	bodyBytes, _ = ioutil.ReadAll(r.Body)
		// 	bodyString := string(bodyBytes)
		// 	fmt.Println("*******************")
		// 	fmt.Println(bodyString);
		// 	var body Body
		// 	json.Unmarshal([]byte(bodyString), &body)
		// 	fmt.Printf("Operation name: %s", body.OperationName)
		// 	fmt.Println()
		// 	fmt.Printf("Variables: %s", body.Variables)
		// 	fmt.Println()
		// 	fmt.Printf("Query: %s", body.Query)
		// 	fmt.Println()
		// 	fmt.Println("*******************")
		// }

		// // Restore the io.ReadCloser to its original state
		// r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		next.ServeHTTP(w, r)
	}
}