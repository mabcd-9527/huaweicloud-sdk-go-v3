package provider

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPodIdentityCredentialProvider_GetCredentials(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-Id", "request-id")
		assert.Equal(t, "token_value", r.Header.Get("Authorization"))
		bytes, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		body := strings.TrimSpace(string(bytes))
		assert.Equal(t, `{}`, body)

		_, err = w.Write([]byte("{\"assumedAgency\":{\"urn\":\"\",\"id\":\"\"}," +
			"\"audience\":\"\",\"credentials\":" +
			"{\"accessKeyId\":\"ak\",\"secretAccessKey\":\"sk\"," +
			"\"securityToken\":\"st\",\"expiration\":\"2020-01-08T03:50:07.574000Z\"}," +
			"\"podIdentityAssociationId\":\"\",\"subject\":" +
			"{\"namespace\":\"\",\"serviceAccount\":\"\"}}"))
		assert.NoError(t, err)

	}))
	defer mockServer.Close()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "BaseTestTokenFile")
	err := os.WriteFile(filePath, []byte("token_value"), 0600)
	filePath, err = filepath.Abs(filePath)
	assert.NoError(t, err)

	err = os.Setenv("HC_CONTAINER_CREDENTIALS_FULL_URI", mockServer.URL)
	assert.NoError(t, err)
	err = os.Setenv("HC_CONTAINER_AUTHORIZATION_TOKEN_FILE", filePath)
	assert.NoError(t, err)

	prov := BasicCredentialPodIdentityProvider()
	credentials, err := prov.GetCredentials()
	if basic, ok := credentials.(*auth.BasicCredentials); ok {
		assert.Equal(t, "ak", basic.AK)
		assert.Equal(t, "sk", basic.SK)
		assert.Equal(t, "st", basic.SecurityToken)
	} else {
		assert.Fail(t, "expect type: *BasicCredentials")
	}

	prov = GlobalCredentialPodIdentityProvider()
	credentials, err = prov.GetCredentials()
	if global, ok := credentials.(*auth.GlobalCredentials); ok {
		assert.Equal(t, "ak", global.AK)
		assert.Equal(t, "sk", global.SK)
		assert.Equal(t, "st", global.SecurityToken)
	} else {
		assert.Fail(t, "expect type: *BasicCredentials")
	}

	assert.NoError(t, err)
	err = os.Unsetenv("HC_CONTAINER_CREDENTIALS_FULL_URI")
	assert.NoError(t, err)
	err = os.Unsetenv("HC_CONTAINER_AUTHORIZATION_TOKEN_FILE")
	assert.NoError(t, err)
}
