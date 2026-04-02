// Copyright 2025 Huawei Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package internal

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const credString = `
{"credential":{"access":"ak","expires_at":"2020-01-08T03:50:07.574000Z","secret":"sk","securitytoken":"sec-token"}}
`

type ServerResponse struct {
	Status int
	Body   string
}

func TestMetadataAccessor_GetCredential(t *testing.T) {
	// imds v1
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == GetTokenPath {
			w.WriteHeader(405)
		} else {
			assert.Empty(t, r.Header.Get(XMetadataToken))
			_, err := w.Write([]byte(credString))
			assert.NoError(t, err)
		}
	}))
	defer mockServer.Close()

	MetadataEndpoint = mockServer.URL
	credentials, err := NewMetadataAccessor().GetCredential()
	assert.NoError(t, err)
	assert.Equal(t, "ak", credentials.Access)
	assert.Equal(t, "sk", credentials.Secret)
	assert.Equal(t, "sec-token", credentials.SecurityToken)
}

func TestMetadataAccessor_GetCredential2(t *testing.T) {
	// imds v2
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == GetTokenPath {
			_, err := w.Write([]byte("metadata-token"))
			assert.NoError(t, err)
		} else {
			assert.Equal(t, "metadata-token", r.Header.Get(XMetadataToken))
			_, err := w.Write([]byte(credString))
			assert.NoError(t, err)
		}
	}))
	defer mockServer.Close()

	MetadataEndpoint = mockServer.URL
	credentials, err := NewMetadataAccessor().GetCredential()
	assert.NoError(t, err)
	assert.Equal(t, "ak", credentials.Access)
	assert.Equal(t, "sk", credentials.Secret)
	assert.Equal(t, "sec-token", credentials.SecurityToken)
}

func TestMetadataAccessor_GetCredential3(t *testing.T) {
	status := 404
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, err := w.Write([]byte("some error"))
		assert.NoError(t, err)
		status += 100
	}))
	defer mockServer.Close()

	MetadataEndpoint = mockServer.URL
	for i := 0; i < 2; i++ {
		_, err := NewMetadataAccessor().GetCredential()
		assert.ErrorContains(t, err, "some error")
	}
}

func TestMetadataAccessor_GetCredential4(t *testing.T) {
	// use v1 first, and disable v1
	count := 1
	respMap := map[int]ServerResponse{}
	respMap[1] = ServerResponse{Status: 405, Body: "{}"}
	respMap[2] = ServerResponse{Status: 200, Body: credString}
	respMap[3] = ServerResponse{Status: 401, Body: "{}"}
	respMap[4] = ServerResponse{Status: 200, Body: "metadata-token"}
	respMap[5] = ServerResponse{Status: 200, Body: credString}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sr := respMap[count]
		w.WriteHeader(sr.Status)
		_, err := w.Write([]byte(sr.Body))
		assert.NoError(t, err)

		count += 1
	}))
	defer mockServer.Close()

	MetadataEndpoint = mockServer.URL
	accessor := NewMetadataAccessor()
	for i := 0; i < 2; i++ {
		credentials, err := accessor.GetCredential()
		assert.NoError(t, err)
		assert.Equal(t, "ak", credentials.Access)
		assert.Equal(t, "sk", credentials.Secret)
		assert.Equal(t, "sec-token", credentials.SecurityToken)
	}
}
