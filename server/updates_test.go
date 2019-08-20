package server

import (
	"strings"
	"testing"
	"time"

	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestUpdates(t *testing.T) {
	t.Run("CheckUpdates", func(t *testing.T) {
		// do not run this test in parallel with other tests as it also tests timings

		const fast, slow = 3 * time.Second, 60 * time.Second

		// that call should always be fast
		version, err := serverClient.Default.Server.Version(server.NewVersionParamsWithTimeout(fast))
		require.NoError(t, err)
		if version.Payload.Server == nil || version.Payload.Server.Version == "" {
			t.Skip("skipping test in developer's environment")
		}

		params := &server.CheckUpdatesParams{
			Context: pmmapitests.Context,
		}
		params.SetTimeout(slow) // that call can be slow with a cold cache
		res, err := serverClient.Default.Server.CheckUpdates(params)
		require.NoError(t, err)

		require.NotEmpty(t, res.Payload.Installed)
		assert.True(t, strings.HasPrefix(res.Payload.Installed.Version, "2.0.0-"),
			"installed.version = %q should have '2.0.0-' prefix", res.Payload.Installed.Version)
		assert.NotEmpty(t, res.Payload.Installed.FullVersion)
		require.NotEmpty(t, res.Payload.Installed.Timestamp)
		ts := time.Time(res.Payload.Installed.Timestamp)
		hour, min, _ := ts.Clock()
		assert.Zero(t, hour, "installed.timestamp should contain only date")
		assert.Zero(t, min, "installed.timestamp should contain only date")

		require.NotEmpty(t, res.Payload.Latest)
		assert.True(t, strings.HasPrefix(res.Payload.Latest.Version, "2.0.0-"),
			"latest.version = %q should have '2.0.0-' prefix", res.Payload.Latest.Version)
		assert.NotEmpty(t, res.Payload.Latest.FullVersion)
		require.NotEmpty(t, res.Payload.Latest.Timestamp)
		ts = time.Time(res.Payload.Latest.Timestamp)
		hour, min, _ = ts.Clock()
		assert.Zero(t, hour, "latest.timestamp should contain only date")
		assert.Zero(t, min, "latest.timestamp should contain only date")

		assert.Equal(t, res.Payload.Installed.FullVersion != res.Payload.Latest.FullVersion, res.Payload.UpdateAvailable)
		assert.NotEmpty(t, res.Payload.LastCheck)

		t.Run("HotCache", func(t *testing.T) {
			params = &server.CheckUpdatesParams{
				Context: pmmapitests.Context,
			}
			params.SetTimeout(fast) // that call should be fast with hot cache
			resAgain, err := serverClient.Default.Server.CheckUpdates(params)
			require.NoError(t, err)

			assert.Equal(t, res.Payload, resAgain.Payload)
		})

		t.Run("Force", func(t *testing.T) {
			params = &server.CheckUpdatesParams{
				Body: server.CheckUpdatesBody{
					Force: true,
				},
				Context: pmmapitests.Context,
			}
			params.SetTimeout(slow) // that call with force can be slow
			resForce, err := serverClient.Default.Server.CheckUpdates(params)
			require.NoError(t, err)

			assert.Equal(t, res.Payload.Installed, resForce.Payload.Installed)
			assert.Equal(t, resForce.Payload.Installed.FullVersion != resForce.Payload.Latest.FullVersion, resForce.Payload.UpdateAvailable)
			assert.NotEqual(t, res.Payload.LastCheck, resForce.Payload.LastCheck)
		})
	})

	t.Run("Update", func(t *testing.T) {
		if !pmmapitests.RunUpdateTest {
			t.Skip("skipping PMM Server update test")
		}

		startRes, err := serverClient.Default.Server.StartUpdate(nil)
		require.NoError(t, err)
		assert.Zero(t, startRes.Payload.LogOffset)
		authToken := startRes.Payload.AuthToken
		require.NotEmpty(t, authToken)

		_, err = serverClient.Default.Server.StartUpdate(nil)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, "Update is already running.")
	})
}
