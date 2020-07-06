package server

import (
	"testing"
	"time"

	"github.com/percona/pmm/api/alertmanager/amclient"
	"github.com/percona/pmm/api/alertmanager/amclient/alert"
	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAlertManager(t *testing.T) {
	res, err := serverClient.Default.Server.GetSettings(nil)
	require.NoError(t, err)
	require.True(t, res.Payload.Settings.TelemetryEnabled)
	require.NoError(t, err)

	t.Run("TestEndsAtForFailedChecksAlerts", func(t *testing.T) {
		if !pmmapitests.RunSTTTests {
			t.Skip("Skipping STT tests until we have environment: https://jira.percona.com/browse/PMM-5106")
		}

		defer restoreSettingsDefaults(t)

		// Enabling STT
		res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
			Body: server.ChangeSettingsBody{
				EnableStt: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.True(t, res.Payload.Settings.TelemetryEnabled)

		const defaultResendInterval = 30

		// 120 sec ping for failed checks alerts to appear in alertmanager
		for i := 0; i < 120; i++ {
			res, err := amclient.Default.Alert.GetAlerts(&alert.GetAlertsParams{
				Filter:  []string{"stt_check=1"},
				Context: pmmapitests.Context,
			})
			require.NoError(t, err)
			if len(res.Payload) == 0 {
				time.Sleep(1 * time.Second)
				continue
			}

			// TODO: Expand this test once we are silencing/removing alerts.
			for _, v := range res.Payload {
				delta := 3 * defaultResendInterval * time.Second
				// Since the `EndsAt` timestamp is always 3 times the
				// `resendInterval` in the future from `UpdatedAt`
				// we check whether they lie in that time delta.
				assert.WithinDuration(t, time.Time(*v.EndsAt), time.Time(*v.UpdatedAt), delta)
			}
			assert.Greater(t, len(res.Payload), 0, "No alerts met")
			break
		}
	})
}
