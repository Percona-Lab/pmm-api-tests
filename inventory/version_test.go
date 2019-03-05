package inventory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Percona-Lab/pmm-api-tests"
)

func TestVersion(t *testing.T) {
	type VersionResponse struct {
		Version string
	}

	t.Run("Get", func(t *testing.T) {
		url := fmt.Sprintf("%sv1/version", *pmmapitests.ServerURLF)
		var versionResponse VersionResponse

		resp, err := http.Get(url)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Body)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		err = json.Unmarshal(body, &versionResponse)
		require.NoError(t, err)
		require.NotNil(t, versionResponse)
		require.NotNil(t, versionResponse.Version)
		require.Equal(t, "2.0.0-dev", versionResponse.Version)
	})
}
