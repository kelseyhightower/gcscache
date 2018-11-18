// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"github.com/kelseyhightower/gcscache"
	"golang.org/x/crypto/acme/autocert"
)

var (
	bucket string
	domain string
)

func main() {
	flag.StringVar(&bucket, "bucket", "", "The GCS bucket.")
	flag.StringVar(&domain, "domain", "", "The domain name to secure.")
	flag.Parse()

	cache, err := gcscache.New(bucket)
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
	}
	s := &http.Server{
		Addr:      "0.0.0.0:443",
		Handler:   mux,
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}

	s.ListenAndServeTLS("", "")
}
