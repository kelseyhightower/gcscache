# GCS Cache

GCS Cache implements the [autocet.Cache](https://godoc.org/golang.org/x/crypto/acme/autocert#Cache) interface using [Google Cloud Storage](https://cloud.google.com/storage/).

## Example Usage

```
package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/kelseyhightower/gcscache"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	cache, err := gcscache.New(bucket, project)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Autocert with Google Cloud Storage Cache"))
	})

	m := autocert.Manager{
		Cache:      cache,
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Handler:    mux,
	}
	s := &http.Server{
		Addr:      "0.0.0.0:443",
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}

	s.ListenAndServeTLS("", "")
}
```
