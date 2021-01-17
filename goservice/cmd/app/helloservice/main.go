package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
// a simple http web service which is exposing /hello end point on port 7070.
func main() {

	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Println("hello GO http web service")
		data, err := ioutil.ReadAll(r.Body) // r.body - reads the request body
		if err != nil {
			http.Error(rw, "error occured", http.StatusInternalServerError) // returns error to the client
			return

			/* another way of returning error response is through response writer.
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("error occured"))
			return
			*/

		}

		fmt.Fprintf(rw, "HELLO %s\n", data) // writes data to the http response.
		
		/*  another way of writing data to the response using response writer
		rw.WriteHeader(http.StatusAccepted)
		rw.Write(data)
        */ 
	})

	// http webservice will be listenig on any ip address and on port 7070.
	http.ListenAndServe(":7070", nil) // when handler is nill, DefaultServMux handlder is used.

}