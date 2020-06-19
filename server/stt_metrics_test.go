package server

import (
	"context"
	"testing"
	"time"

	managementClient "github.com/percona/pmm/api/managementpb/json/client"
	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/prometheus/client_golang/api"
	promapi "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

const promAddress = "http://admin:admin@127.0.0.1:80/prometheus"

func TestSTTMetrics(t *testing.T) {
	client := serverClient.Default.Server

	t.Run("StartSTTChecksAndRecordMetrics", func(t *testing.T) {
		defer restoreSettingsDefaults(t)
		// Enabled STT
		res, err := client.ChangeSettings(&server.ChangeSettingsParams{
			Body: server.ChangeSettingsBody{
				EnableStt:       true,
				EnableTelemetry: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.True(t, res.Payload.Settings.SttEnabled)
		assert.True(t, res.Payload.Settings.TelemetryEnabled)
		assert.Empty(t, err)

		resp, err := managementClient.Default.SecurityChecks.StartSecurityChecks(nil)
		require.NoError(t, err)
		assert.NotNil(t, resp)

		client, err := api.NewClient(api.Config{
			Address: promAddress,
		})
		require.NoError(t, err)
		promClient := promapi.NewAPI(client)

		testCases := []struct {
			query string
			len   int
		}{
			{
				query: "pmm_managed_checks_alerts_generated_total",
				len:   7,
			},
			{
				query: "pmm_managed_checks_scripts_executed_total",
				len:   3,
			},
		}

		for _, tc := range testCases {
			result, _, err := promClient.Query(context.Background(), tc.query, time.Now())
			require.NoError(t, err)
			assert.NotEmpty(t, result)
			assert.Len(t, result, tc.len)
		}
	})
}
