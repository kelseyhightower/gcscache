# GCS Cache

[![Go Report Card](https://goreportcard.com/badge/github.com/kelseyhightower/gcscache)](https://goreportcard.com/report/github.com/kelseyhightower/gcscache) [![GoDoc](https://godoc.org/github.com/kelseyhightower/gcscache?status.svg)](https://godoc.org/github.com/kelseyhightower/gcscache)

GCS Cache implements the [autocert.Cache](https://godoc.org/golang.org/x/crypto/acme/autocert#Cache) interface using [Google Cloud Storage](https://cloud.google.com/storage/).

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
    cache, err := gcscache.New("bucket")
    if err != nil {
        log.Fatal(err)
    }

    m := autocert.Manager{
        Cache:      cache,
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("example.org"),
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Autocert with Google Cloud Storage Cache"))
    })

    s := &http.Server{
        Addr:      "0.0.0.0:443",
        Handler:   mux,
        TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
    }

    s.ListenAndServeTLS("", "")
}
```
