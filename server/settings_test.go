package server

import (
	"testing"

	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestSettings(t *testing.T) {
	t.Run("GetSettings", func(t *testing.T) {
		res, err := serverClient.Default.Server.GetSettings(nil)
		require.NoError(t, err)
		assert.True(t, res.Payload.Settings.Telemetry)
		expected := &server.GetSettingsOKBodySettingsMetricsResolutions{
			Hr: "1s",
			Mr: "5s",
			Lr: "60s",
		}
		require.Equal(t, expected, res.Payload.Settings.MetricsResolutions)
		expectedDataRetention := &server.GetSettingsOKBodySettingsQAN{
			DataRetention: "2592000s",
		}
		require.Equal(t, expectedDataRetention, res.Payload.Settings.QAN)

		t.Run("ChangeSettings", func(t *testing.T) {
			// always restore settings on exit
			defer func() {
				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Context: pmmapitests.Context,
					Body: server.ChangeSettingsBody{
						EnableTelemetry: true,
						MetricsResolutions: &server.ChangeSettingsParamsBodyMetricsResolutions{
							Hr: "1s",
							Mr: "5s",
							Lr: "60s",
						},
						QAN: &server.ChangeSettingsParamsBodyQAN{
							DataRetention: "720h",
						},
					},
				})
				require.NoError(t, err)
				assert.True(t, res.Payload.Settings.Telemetry)
				expected := &server.ChangeSettingsOKBodySettingsMetricsResolutions{
					Hr: "1s",
					Mr: "5s",
					Lr: "60s",
				}
				assert.Equal(t, expected, res.Payload.Settings.MetricsResolutions)
				expectedDataRetention := &server.ChangeSettingsOKBodySettingsQAN{
					DataRetention: "2592000s",
				}
				require.Equal(t, expectedDataRetention, res.Payload.Settings.QAN)
			}()

			t.Run("BothEnableAndDisable", func(t *testing.T) {
				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Context: pmmapitests.Context,
					Body: server.ChangeSettingsBody{
						EnableTelemetry:  true,
						DisableTelemetry: true,
					},
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `Both enable_telemetry and disable_telemetry are present.`)
				assert.Empty(t, res)
			})

			t.Run("InvalidDataRetentionDuration", func(t *testing.T) {
				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Context: pmmapitests.Context,
					Body: server.ChangeSettingsBody{
						QAN: &server.ChangeSettingsParamsBodyQAN{
							DataRetention: "INVALID_DURATION",
						},
					},
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `bad Duration: time: invalid duration INVALID_DURATION`)
				assert.Empty(t, res)
			})

			t.Run("InvalidDataRetentionDuration2", func(t *testing.T) {
				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Context: pmmapitests.Context,
					Body: server.ChangeSettingsBody{
						QAN: &server.ChangeSettingsParamsBodyQAN{
							DataRetention: "10s",
						},
					},
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `The data retention duration must be a multiple of 24 hours, but is 10s`)
				assert.Empty(t, res)
			})

			t.Run("TooSmall", func(t *testing.T) {
				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Context: pmmapitests.Context,
					Body: server.ChangeSettingsBody{
						MetricsResolutions: &server.ChangeSettingsParamsBodyMetricsResolutions{
							Hr: "0.1s",
						},
					},
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, `Minimal resolution is 1s.`)
				assert.Empty(t, res)
			})

			t.Run("OK", func(t *testing.T) {
				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Context: pmmapitests.Context,
					Body: server.ChangeSettingsBody{
						DisableTelemetry: true,
						MetricsResolutions: &server.ChangeSettingsParamsBodyMetricsResolutions{
							Hr: "2s",
							Mr: "15s",
							Lr: "2m",
						},
						QAN: &server.ChangeSettingsParamsBodyQAN{
							DataRetention: "240h",
						},
					},
				})
				require.NoError(t, err)
				assert.False(t, res.Payload.Settings.Telemetry)
				expected := &server.ChangeSettingsOKBodySettingsMetricsResolutions{
					Hr: "2s",
					Mr: "15s",
					Lr: "120s",
				}
				assert.Equal(t, expected, res.Payload.Settings.MetricsResolutions)

				getRes, err := serverClient.Default.Server.GetSettings(nil)
				require.NoError(t, err)
				assert.False(t, getRes.Payload.Settings.Telemetry)
				getExpected := &server.GetSettingsOKBodySettingsMetricsResolutions{
					Hr: "2s",
					Mr: "15s",
					Lr: "120s",
				}
				require.Equal(t, getExpected, getRes.Payload.Settings.MetricsResolutions)
				expectedDataRetention := &server.GetSettingsOKBodySettingsQAN{
					DataRetention: "864000s",
				}
				require.Equal(t, expectedDataRetention, getRes.Payload.Settings.QAN)
			})
		})
	})
}
