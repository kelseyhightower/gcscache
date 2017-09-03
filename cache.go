// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcscache

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

type Cache struct {
	client    *storage.Client
	bucket    string
	projectID string
}

func New(bucket, projectId string) (*Cache, error) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	c := &Cache{bucket, client, projectId}

	return c, nil
}

func (c *Cache) Get(ctx context.Context, name string) ([]byte, error) {
	r, err := c.client.Bucket(c.bucket).Object(name).NewReader(context.Background())
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

func (c *Cache) Put(ctx context.Context, name string, data []byte) error {
	w := c.client.Bucket(c.bucket).Object(name).NewWriter(context.Background())
	w.Write(data)
	return w.Close()
}

func (c *Cache) Delete(ctx context.Context, name string) error {
	o := c.client.Bucket(c.bucket).Object(name)
	return o.Delete(context.Background())
}
