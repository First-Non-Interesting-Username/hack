package cmd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeRequest(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		assert.Equal(t, "/chat/completions", r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"mocked response"`))
	}))
	defer mockServer.Close()

	viper.Set("base_url", mockServer.URL)
	viper.Set("api_key", "test-api-key")
	viper.Set("model", "test-model")

	prompt = "Hello"
	commandMode = false
	codeMode = false

	restore := mockStdin("")
	defer restore()

	response, err := makeRequest()

	require.NoError(t, err)
	assert.Equal(t, `"mocked response"`, response)
}
