package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
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
		assert.True(t, res.Payload.Settings.TelemetryEnabled)
		expected := &server.GetSettingsOKBodySettingsMetricsResolutions{
			Hr: "5s",
			Mr: "5s",
			Lr: "60s",
		}
		assert.Equal(t, expected, res.Payload.Settings.MetricsResolutions)
		assert.Equal(t, "2592000s", res.Payload.Settings.DataRetention)
		assert.Equal(t, []string{"aws"}, res.Payload.Settings.AWSPartitions)

		t.Run("ChangeSettings", func(t *testing.T) {
			teardown := func(t *testing.T) {
				t.Helper()

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						EnableTelemetry: true,
						MetricsResolutions: &server.ChangeSettingsParamsBodyMetricsResolutions{
							Hr: "5s",
							Mr: "5s",
							Lr: "60s",
						},
						DataRetention: "720h",
						AWSPartitions: []string{"aws"},
					},
					Context: pmmapitests.Context,
				})
				require.NoError(t, err)
				assert.True(t, res.Payload.Settings.TelemetryEnabled)
				expected := &server.ChangeSettingsOKBodySettingsMetricsResolutions{
					Hr: "5s",
					Mr: "5s",
					Lr: "60s",
				}
				assert.Equal(t, expected, res.Payload.Settings.MetricsResolutions)
				assert.Equal(t, "2592000s", res.Payload.Settings.DataRetention)
				assert.Equal(t, []string{"aws"}, res.Payload.Settings.AWSPartitions)
			}

			defer teardown(t)

			t.Run("InvalidBothEnableAndDisable", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						EnableTelemetry:  true,
						DisableTelemetry: true,
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `Both enable_telemetry and disable_telemetry are present.`)
				assert.Empty(t, res)
			})

			t.Run("InvalidPartition", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						AWSPartitions: []string{"aws-123"},
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `aws_partitions: partition "aws-123" is invalid`)
				assert.Empty(t, res)
			})

			t.Run("TooManyPartitions", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						AWSPartitions: []string{"aws", "aws", "aws", "aws", "aws", "aws"},
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `aws_partitions: list is too long`)
				assert.Empty(t, res)
			})

			t.Run("HRInvalid", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						MetricsResolutions: &server.ChangeSettingsParamsBodyMetricsResolutions{
							Hr: "1",
						},
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `bad Duration: time: missing unit in duration 1`)
				assert.Empty(t, res)
			})

			t.Run("HRTooSmall", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						MetricsResolutions: &server.ChangeSettingsParamsBodyMetricsResolutions{
							Hr: "0.5s",
						},
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `hr: minimal resolution is 1s`)
				assert.Empty(t, res)
			})

			t.Run("HRFractional", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						MetricsResolutions: &server.ChangeSettingsParamsBodyMetricsResolutions{
							Hr: "1.5s",
						},
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `hr: should be a natural number of seconds`)
				assert.Empty(t, res)
			})

			t.Run("DataRetentionInvalid", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						DataRetention: "1",
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `bad Duration: time: missing unit in duration 1`)
				assert.Empty(t, res)
			})

			t.Run("DataRetentionInvalidToSmall", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						DataRetention: "10s",
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `data_retention: minimal resolution is 24h`)
				assert.Empty(t, res)
			})

			t.Run("DataRetentionFractional", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						DataRetention: "36h",
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `data_retention: should be a natural number of days`)
				assert.Empty(t, res)
			})

			t.Run("InvalidSSHKey", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						SSHKey: "some-invalid-ssh-key",
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `Invalid ssh key`)
				assert.Empty(t, res)
			})

			t.Run("NoAdminUserForSSH", func(t *testing.T) {
				defer teardown(t)

				sshKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQClY/8sz3w03vA2bY6mBFgUzrvb2FIoHw8ZjUXGGClJzJg5HC3jW1m5df7TOIkx0bt6Da2UOhuCvS4o27IT1aiHXVFydppp6ghQRB6saiiW2TKlQ7B+mXatwVaOIkO381kEjgijAs0LJnNRGpqQW0ZEAxVMz4a8puaZmVNicYSVYs4kV3QZsHuqn7jHbxs5NGAO+uRRSjcuPXregsyd87RAUHkGmNrwNFln/XddMzdGMwqZOuZWuxIXBqSrSX927XGHAJlUaOmLz5etZXHzfAY1Zxfu39r66Sx95bpm3JBmc/Ewfr8T2WL0cqynkpH+3QQBCjweTHzBE+lpXHdR2se1 qsandbox"
				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						SSHKey: sshKey,
					},
					Context: pmmapitests.Context,
				})
				pmmapitests.AssertAPIErrorf(t, err, 500, codes.Internal, `Internal server error.`)
				assert.Empty(t, res)
			})

			t.Run("OK", func(t *testing.T) {
				defer teardown(t)

				res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
					Body: server.ChangeSettingsBody{
						DisableTelemetry: true,
						MetricsResolutions: &server.ChangeSettingsParamsBodyMetricsResolutions{
							Hr: "2s",
							Mr: "15s",
							Lr: "2m",
						},
						DataRetention: "240h",
						AWSPartitions: []string{"aws-cn", "aws", "aws-cn"}, // duplicates are ok
					},
					Context: pmmapitests.Context,
				})
				require.NoError(t, err)
				assert.False(t, res.Payload.Settings.TelemetryEnabled)
				expected := &server.ChangeSettingsOKBodySettingsMetricsResolutions{
					Hr: "2s",
					Mr: "15s",
					Lr: "120s",
				}
				assert.Equal(t, expected, res.Payload.Settings.MetricsResolutions)
				assert.Equal(t, []string{"aws", "aws-cn"}, res.Payload.Settings.AWSPartitions)

				getRes, err := serverClient.Default.Server.GetSettings(nil)
				require.NoError(t, err)
				assert.False(t, getRes.Payload.Settings.TelemetryEnabled)
				getExpected := &server.GetSettingsOKBodySettingsMetricsResolutions{
					Hr: "2s",
					Mr: "15s",
					Lr: "120s",
				}
				assert.Equal(t, getExpected, getRes.Payload.Settings.MetricsResolutions)
				assert.Equal(t, "864000s", res.Payload.Settings.DataRetention)
				assert.Equal(t, []string{"aws", "aws-cn"}, res.Payload.Settings.AWSPartitions)

				t.Run("DefaultsAreNotRestored", func(t *testing.T) {
					defer teardown(t)

					res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
						Body:    server.ChangeSettingsBody{},
						Context: pmmapitests.Context,
					})
					require.NoError(t, err)
					assert.False(t, res.Payload.Settings.TelemetryEnabled)
					expected := &server.ChangeSettingsOKBodySettingsMetricsResolutions{
						Hr: "2s",
						Mr: "15s",
						Lr: "120s",
					}
					assert.Equal(t, expected, res.Payload.Settings.MetricsResolutions)
					assert.Equal(t, []string{"aws", "aws-cn"}, res.Payload.Settings.AWSPartitions)

					getRes, err := serverClient.Default.Server.GetSettings(nil)
					require.NoError(t, err)
					assert.False(t, getRes.Payload.Settings.TelemetryEnabled)
					getExpected := &server.GetSettingsOKBodySettingsMetricsResolutions{
						Hr: "2s",
						Mr: "15s",
						Lr: "120s",
					}
					assert.Equal(t, getExpected, getRes.Payload.Settings.MetricsResolutions)
					assert.Equal(t, "864000s", res.Payload.Settings.DataRetention)
					assert.Equal(t, []string{"aws", "aws-cn"}, res.Payload.Settings.AWSPartitions)
				})

				t.Run("Set Alert Manager Rules", func(t *testing.T) {
					defer teardown(t)
					body := server.ChangeSettingsBody{
						AlertManagerAddress: "localhost:1234",
						AlertManagerRules: `groups:
- name: example
  rules:
  - alert: HighRequestLatency
    expr: job:request_latency_seconds:mean5m{job="myjob"} > 0.5
    for: 10m
    labels:
      severity: page
    annotations:
      summary: High request latency`,
					}

					res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
						Body:    body,
						Context: pmmapitests.Context,
					})
					require.NoError(t, err)
					assert.Equal(t, res.Payload.Settings.AlertManagerAddress, body.AlertManagerAddress)
					assert.Equal(t, res.Payload.Settings.AlertManagerRules, body.AlertManagerRules)
				})

				t.Run("Clear Alert Manager Rules", func(t *testing.T) {
					defer teardown(t)
					body := server.ChangeSettingsBody{
						AlertManagerAddress: "localhost:1234",
						AlertManagerRules: `groups:
- name: example
  rules:
  - alert: HighRequestLatency
    expr: job:request_latency_seconds:mean5m{job="myjob"} > 0.5
    for: 10m
    labels:
      severity: page
    annotations:
      summary: High request latency`,
						RemoveAlertManagerAddress: true,
						RemoveAlertManagerRules:   true,
					}

					res, err := serverClient.Default.Server.ChangeSettings(&server.ChangeSettingsParams{
						Body:    body,
						Context: pmmapitests.Context,
					})
					require.NoError(t, err)
					assert.Equal(t, res.Payload.Settings.AlertManagerAddress, "")
					assert.Equal(t, res.Payload.Settings.AlertManagerRules, "")
				})
			})

			t.Run("grpc-gateway", func(t *testing.T) {
				// Test with pure JSON without swagger for tracking grpc-gateway behavior:
				// https://github.com/grpc-ecosystem/grpc-gateway/issues/400

				// do not use generated types as they can do extra work in generated methods
				type params struct {
					Settings struct {
						MetricsResolutions struct {
							LR string `json:"lr"`
						} `json:"metrics_resolutions"`
					} `json:"settings"`
				}
				changeURI := pmmapitests.BaseURL.ResolveReference(&url.URL{
					Path: "v1/Settings/Change",
				})
				getURI := pmmapitests.BaseURL.ResolveReference(&url.URL{
					Path: "v1/Settings/Get",
				})

				for change, get := range map[string]string{
					"59s": "59s",
					"60s": "60s",
					"61s": "61s",
					"61":  "", // no suffix => error
					"2m":  "120s",
					"1h":  "3600s",
					"1d":  "", // d suffix => error
					"1w":  "", // w suffix => error
				} {
					change, get := change, get
					t.Run(change, func(t *testing.T) {
						defer teardown(t)

						var p params
						p.Settings.MetricsResolutions.LR = change
						b, err := json.Marshal(p.Settings)
						require.NoError(t, err)
						req, err := http.NewRequest("POST", changeURI.String(), bytes.NewReader(b))
						require.NoError(t, err)
						if pmmapitests.Debug {
							b, err = httputil.DumpRequestOut(req, true)
							require.NoError(t, err)
							t.Logf("Request:\n%s", b)
						}

						resp, err := http.DefaultClient.Do(req)
						require.NoError(t, err)
						if pmmapitests.Debug {
							b, err = httputil.DumpResponse(resp, true)
							require.NoError(t, err)
							t.Logf("Response:\n%s", b)
						}
						b, err = ioutil.ReadAll(resp.Body)
						assert.NoError(t, err)
						resp.Body.Close() //nolint:errcheck

						if get == "" {
							assert.Equal(t, 400, resp.StatusCode, "response:\n%s", b)
							return
						}
						assert.Equal(t, 200, resp.StatusCode, "response:\n%s", b)

						p.Settings.MetricsResolutions.LR = ""
						err = json.Unmarshal(b, &p)
						require.NoError(t, err)
						assert.Equal(t, get, p.Settings.MetricsResolutions.LR, "Change")

						req, err = http.NewRequest("POST", getURI.String(), nil)
						require.NoError(t, err)
						if pmmapitests.Debug {
							b, err = httputil.DumpRequestOut(req, true)
							require.NoError(t, err)
							t.Logf("Request:\n%s", b)
						}

						resp, err = http.DefaultClient.Do(req)
						require.NoError(t, err)
						if pmmapitests.Debug {
							b, err = httputil.DumpResponse(resp, true)
							require.NoError(t, err)
							t.Logf("Response:\n%s", b)
						}
						b, err = ioutil.ReadAll(resp.Body)
						assert.NoError(t, err)
						resp.Body.Close() //nolint:errcheck
						assert.Equal(t, 200, resp.StatusCode, "response:\n%s", b)

						p.Settings.MetricsResolutions.LR = ""
						err = json.Unmarshal(b, &p)
						require.NoError(t, err)
						assert.Equal(t, get, p.Settings.MetricsResolutions.LR, "Get")
					})
				}
			})
		})
	})
}
