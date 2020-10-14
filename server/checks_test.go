package server

import (
	"testing"

	managementClient "github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/security_checks"
	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestStartChecks(t *testing.T) {
	client := serverClient.Default.Server

	t.Run("with enabled STT", func(t *testing.T) {
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

		resp, err := managementClient.Default.SecurityChecks.StartSecurityChecks(nil)
		require.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("with disabled STT", func(t *testing.T) {
		defer restoreSettingsDefaults(t)
		// Disabled STT
		res, err := client.ChangeSettings(&server.ChangeSettingsParams{
			Body: server.ChangeSettingsBody{
				DisableStt:      true,
				EnableTelemetry: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.False(t, res.Payload.Settings.SttEnabled)
		assert.True(t, res.Payload.Settings.TelemetryEnabled)

		resp, err := managementClient.Default.SecurityChecks.StartSecurityChecks(nil)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, `STT is disabled.`)
		assert.Nil(t, resp)
	})
}

func TestGetSecurityCheckResults(t *testing.T) {
	if !pmmapitests.RunSTTTests {
		t.Skip("Skipping STT tests until we have environment: https://jira.percona.com/browse/PMM-5106")
	}

	client := serverClient.Default.Server

	t.Run("with disabled STT", func(t *testing.T) {
		defer restoreSettingsDefaults(t)
		// Disabled STT
		res, err := client.ChangeSettings(&server.ChangeSettingsParams{
			Body: server.ChangeSettingsBody{
				DisableStt: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.False(t, res.Payload.Settings.SttEnabled)

		results, err := managementClient.Default.SecurityChecks.GetSecurityCheckResults(nil)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, `STT is disabled.`)
		assert.Nil(t, results)
	})

	t.Run("with enabled STT", func(t *testing.T) {
		defer restoreSettingsDefaults(t)
		// Enabled STT
		res, err := client.ChangeSettings(&server.ChangeSettingsParams{
			Body: server.ChangeSettingsBody{
				EnableStt: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.True(t, res.Payload.Settings.SttEnabled)

		resp, err := managementClient.Default.SecurityChecks.StartSecurityChecks(nil)
		require.NoError(t, err)
		assert.NotNil(t, resp)

		results, err := managementClient.Default.SecurityChecks.GetSecurityCheckResults(nil)
		require.NoError(t, err)
		assert.NotNil(t, results)
	})
}

func TestListSecurityChecks(t *testing.T) {
	client := serverClient.Default.Server

	defer restoreSettingsDefaults(t)
	// Enable STT
	res, err := client.ChangeSettings(&server.ChangeSettingsParams{
		Body: server.ChangeSettingsBody{
			EnableStt: true,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)
	assert.True(t, res.Payload.Settings.SttEnabled)

	resp, err := managementClient.Default.SecurityChecks.ListSecurityChecks(nil)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Payload.ChecksStates)
}

func TestToggleSecurityChecks(t *testing.T) {
	client := serverClient.Default.Server

	defer restoreSettingsDefaults(t)
	// Enable STT
	res, err := client.ChangeSettings(&server.ChangeSettingsParams{
		Body: server.ChangeSettingsBody{
			EnableStt: true,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)
	assert.True(t, res.Payload.Settings.SttEnabled)

	resp, err := managementClient.Default.SecurityChecks.ListSecurityChecks(nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp.Payload.ChecksStates)

	var check *security_checks.ChecksStatesItems0
	var disable bool
	var params *security_checks.ToggleSecurityChecksParams

	// enable ⥁ disable loop, it checks current state of first returned check and changes its state,
	// then in second iteration it returns state to its origin.
	for i := 0; i < 2; i++ {
		check = resp.Payload.ChecksStates[0]
		disable = !check.Disabled

		params = &security_checks.ToggleSecurityChecksParams{
			Body: security_checks.ToggleSecurityChecksBody{
				ChecksParams: []*security_checks.ChecksParamsItems0{
					{
						Name:    check.Name,
						Disable: disable,
						Enable:  !disable,
					},
				},
			},
			Context: pmmapitests.Context,
		}

		_, err = managementClient.Default.SecurityChecks.ToggleSecurityChecks(params)
		require.NoError(t, err)

		resp, err = managementClient.Default.SecurityChecks.ListSecurityChecks(nil)
		require.NoError(t, err)
		require.NotEmpty(t, resp.Payload.ChecksStates)

		assert.Equal(t, check.Name, resp.Payload.ChecksStates[0].Name)
		assert.Equal(t, !check.Disabled, resp.Payload.ChecksStates[0].Disabled)
	}
}
