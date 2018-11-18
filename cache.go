// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

// Package gcscache provides storage, backed by Google Cloud Storage,
// for certificates managed by the golang.org/x/crypto/acme/autocert package.
//
// This package is a work in progress and makes no API stability promises.
package gcscache

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"golang.org/x/crypto/acme/autocert"
)

// Cache implements the autocert.Cache interface using Google Cloud Storage.
type Cache struct {
	client *storage.Client
	bucket string
}

// New creates and initializes a new Cache backed by the given Google Cloud
// Storage bucket.
func New(bucket string) (*Cache, error) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	c := &Cache{client, bucket}

	return c, nil
}

// Get reads a certificate data from the specified object name.
func (c *Cache) Get(ctx context.Context, name string) ([]byte, error) {
	r, err := c.client.Bucket(c.bucket).Object(name).NewReader(context.Background())

	if err == storage.ErrObjectNotExist {
		return nil, autocert.ErrCacheMiss
	}

	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

// Put writes the certificate data to the specified object name.
func (c *Cache) Put(ctx context.Context, name string, data []byte) error {
	w := c.client.Bucket(c.bucket).Object(name).NewWriter(context.Background())
	w.Write(data)
	return w.Close()
}

// Delete removes the specified object name.
func (c *Cache) Delete(ctx context.Context, name string) error {
	o := c.client.Bucket(c.bucket).Object(name)
	err := o.Delete(context.Background())
	if err == storage.ErrObjectNotExist {
		return nil
	}
	return err
}
