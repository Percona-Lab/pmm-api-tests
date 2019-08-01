package server

import (
	"bytes"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAuth(t *testing.T) {
	t.Parallel()

	// make a BaseURL without authentication
	baseURL, err := url.Parse(pmmapitests.BaseURL.String())
	require.NoError(t, err)
	baseURL.User = nil

	uri := baseURL.ResolveReference(&url.URL{
		Path: "v1/version",
	})

	t.Logf("URI: %s", uri)
	resp, err := http.Get(uri.String())
	require.NoError(t, err)
	defer resp.Body.Close() //nolint:errcheck
	b, err := httputil.DumpResponse(resp, true)
	require.NoError(t, err)

	assert.Equal(t, 401, resp.StatusCode, "response:\n%s", b)
	assert.False(t, bytes.Contains(b, []byte(`<html>`)), "response:\n%s", b)
}
