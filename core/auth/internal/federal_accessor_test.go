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
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
)

func TestFederalAccessor_GetCredential(t *testing.T) {
	var count int32
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-IAM-Trace-Id", "trace-id")
		w.Header().Set("X-Request-Id", "request-id")
		if atomic.LoadInt32(&count) == 0 {
			// v3.0/OS-AUTH/id-token/tokens
			assert.Equal(t, "idp_id", r.Header.Get("X-Idp-Id"))
			bytes, err := ioutil.ReadAll(r.Body)
			assert.NoError(t, err)
			body := strings.TrimSpace(string(bytes))
			assert.Equal(t, `{"auth":{"id_token":{"id":"id_token"}}}`, body)

			w.Header().Set("X-Subject-Token", "auth-token")
			_, err = w.Write([]byte(`{"token":{"expires_at":"2018-03-13T03:00:01.168000Z","methods":["mapped"]}}`))
			assert.NoError(t, err)
		} else {
			// v3.0/OS-CREDENTIAL/securitytokens
			bytes, err := ioutil.ReadAll(r.Body)
			assert.NoError(t, err)
			body := strings.TrimSpace(string(bytes))
			assert.Equal(t, `{"auth":{"identity":{"methods":["token"],"token":{"id":"auth-token","duration_seconds":21600}}}}`, body)

			_, err = w.Write([]byte(`{"credential":{"access":"ak","expires_at":"2020-01-08T03:50:07.574000Z","secret":"sk","securitytoken":"sec-token"}}`))
			assert.NoError(t, err)
		}

		atomic.AddInt32(&count, 1)
	}))
	defer mockServer.Close()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "TestFederalAccessorIdTokenFile")
	err := os.WriteFile(filePath, []byte("id_token"), 0600)
	filePath, err = filepath.Abs(filePath)
	assert.NoError(t, err)

	client := impl.NewDefaultHttpClient(config.DefaultHttpConfig().WithIgnoreSSLVerification(true))
	accessor := NewFederalAccessor()
	credentials, err := accessor.GetCredential(
		WithIamEndpoint(mockServer.URL),
		WithClient(client),
		WithIdpId("idp_id"),
		WithIdTokenFile(filePath),
	)
	assert.NoError(t, err)
	assert.Equal(t, "ak", credentials.Access)
	assert.Equal(t, "sk", credentials.Secret)
	assert.Equal(t, "sec-token", credentials.SecurityToken)
}
