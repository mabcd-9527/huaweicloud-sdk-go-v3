package internal

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewPodIdentityAccessor(t *testing.T) {

}

func TestPodIdentityAccessor_GetCredential(t *testing.T) {
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
	filePath := filepath.Join(tmpDir, "TestTokenFile")
	err := os.WriteFile(filePath, []byte("token_value"), 0600)
	filePath, err = filepath.Abs(filePath)
	assert.NoError(t, err)

	accessor := NewPodIdentityAccessor(mockServer.URL, filePath)
	credential, err := accessor.GetCredential()
	assert.NoError(t, err)
	assert.Equal(t, "ak", credential.Access)
	assert.Equal(t, "sk", credential.Secret)
	assert.Equal(t, "st", credential.SecurityToken)
	assert.Equal(t, int64(1578455407), credential.ExpireAt)
}
