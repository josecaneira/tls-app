// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//
// Original work by Google Inc. at https://github.com/googlecloudplatform/kubernetes-engine-samples/tree/394477ca2a99/quickstarts/hello-app-tls
// Derivate work by José Caneira at https://github.com/josecaneira for TLS end-to-end Demo.
//

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8443"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	tlsCert, tlsKey := os.Getenv("TLS_CERT"), os.Getenv("TLS_KEY")
	if tlsCert == "" {
		log.Fatal("TLS_CERT environment variable must be set")
	}
	if tlsKey == "" {
		log.Fatal("TLS_KEY environment variable must be set")
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("tls cert: %s", tlsCert)
	log.Printf("tls key: %s", tlsKey)
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServeTLS(":"+port, tlsCert, tlsKey, server)
	log.Fatal(err)
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	fmt.Fprintf(w, "\nTLS end-to-end Hello, world!\n\n")
	fmt.Fprintf(w, "Protocol: %s!\n", r.Proto)
	fmt.Fprintf(w, "Domain: %s\n", r.Host)
	fmt.Fprintf(w, "Backend Pod Hostname: %s\n", host)
	if headerIP := r.Header.Get("X-Forwarded-For"); headerIP != "" {
		fmt.Fprintf(w, "Client IP (X-Forwarded-For): %s\n", headerIP)
		log.Printf("Serving request: %s for: %s", r.URL.Path, headerIP)
	} else {
		fmt.Fprintf(w, "Client IP: %s\n", r.RemoteAddr)
		log.Printf("Serving request: %s for: %s", r.URL.Path, r.RemoteAddr)
	}
}