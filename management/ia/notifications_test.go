package ia

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	channelsClient "github.com/percona/pmm/api/managementpb/ia/json/client"
	"github.com/percona/pmm/api/managementpb/ia/json/client/channels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAddChannel(t *testing.T) {
	// if !pmmapitests.RunIATests {
	// 	t.Skip("Skipping IA tests until IA will out of beta: https://jira.percona.com/browse/PMM-7001")
	// }

	client := channelsClient.Default.Channels

	t.Run("normal", func(t *testing.T) {
		_, err := client.AddChannel(&channels.AddChannelParams{
			Body: channels.AddChannelBody{
				Summary: gofakeit.Quote(),
				Disabled: gofakeit.Bool(),
				EmailConfig: &channels.AddChannelParamsBodyEmailConfig{
					SendResolved: false,
					To:           []string{gofakeit.Email()},
				},
			},
			Context: pmmapitests.Context,
		})

		require.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {
		_, err := client.AddChannel(&channels.AddChannelParams{
			Body: channels.AddChannelBody{
				Summary: gofakeit.Quote(),
				Disabled: gofakeit.Bool(),
				EmailConfig: &channels.AddChannelParamsBodyEmailConfig{
					SendResolved: false,
				},
			},
			Context: pmmapitests.Context,
		})

		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field EmailConfig.To: value '[]' must contain at least 1 elements")
	})

	t.Run("missing config", func(t *testing.T) {
		_, err := client.AddChannel(&channels.AddChannelParams{
			Body: channels.AddChannelBody{
				Summary: gofakeit.Quote(),
				Disabled: gofakeit.Bool(),
			},
			Context: pmmapitests.Context,
		})

		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "Missing channel configuration.")
	})
}

func TestChangeChannel(t *testing.T) {
	// if !pmmapitests.RunIATests {
	// 	t.Skip("Skipping IA tests until IA will out of beta: https://jira.percona.com/browse/PMM-7001")
	// }

	client := channelsClient.Default.Channels

	t.Run("normal", func(t *testing.T) {
		summary := gofakeit.UUID()
		_, err := client.AddChannel(&channels.AddChannelParams{
			Body: channels.AddChannelBody{
				Summary:  summary,
				Disabled: gofakeit.Bool(),
				EmailConfig: &channels.AddChannelParamsBodyEmailConfig{
					SendResolved: false,
					To:           []string{gofakeit.Email()},
				},
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		resp, err := client.ListChannels(&channels.ListChannelsParams{Context: pmmapitests.Context})
		require.NoError(t, err)

		var id string
		for _, channel := range resp.Payload.Channels {
			if channel.Summary == summary {
				id = channel.ChannelID
			}
		}
		require.NotEmpty(t, id)

		newEmail := []string{gofakeit.Email()}
		_, err = client.ChangeChannel(&channels.ChangeChannelParams{
			Body: channels.ChangeChannelBody{
				ChannelID: id,
				Disabled: gofakeit.Bool(),
				EmailConfig: &channels.ChangeChannelParamsBodyEmailConfig{
					SendResolved: true,
					To:           newEmail,
				},
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		resp, err = client.ListChannels(&channels.ListChannelsParams{Context: pmmapitests.Context})
		require.NoError(t, err)

		assert.NotEmpty(t, resp.Payload.Channels)
		var found bool
		for _, channel := range resp.Payload.Channels {
			if channel.ChannelID == id {
				assert.Equal(t, newEmail, channel.EmailConfig.To)
				assert.True(t, channel.EmailConfig.SendResolved)
				found = true
			}
		}

		assert.True(t, found, "Expected channel not found")
	})
}

func TestRemoveChannel(t *testing.T) {
	// if !pmmapitests.RunIATests {
	// 	t.Skip("Skipping IA tests until IA will out of beta: https://jira.percona.com/browse/PMM-7001")
	// }

	client := channelsClient.Default.Channels

	t.Run("normal", func(t *testing.T) {
		summary := gofakeit.UUID()
		_, err := client.AddChannel(&channels.AddChannelParams{
			Body: channels.AddChannelBody{
				Summary:  summary,
				Disabled: gofakeit.Bool(),
				EmailConfig: &channels.AddChannelParamsBodyEmailConfig{
					SendResolved: false,
					To:           []string{gofakeit.Email()},
				},
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		resp, err := client.ListChannels(&channels.ListChannelsParams{Context: pmmapitests.Context})
		require.NoError(t, err)

		var id string
		for _, channel := range resp.Payload.Channels {
			if channel.Summary == summary {
				id = channel.ChannelID
			}
		}
		require.NotEmpty(t, id)

		_, err = client.RemoveChannel(&channels.RemoveChannelParams{
			Body: channels.RemoveChannelBody{
				ChannelID: id,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		resp, err = client.ListChannels(&channels.ListChannelsParams{Context: pmmapitests.Context})
		require.NoError(t, err)

		assert.NotEmpty(t, resp.Payload.Channels)
		for _, channel := range resp.Payload.Channels {
			if channel.ChannelID == id {
				assert.NotEqual(t, id, channel.ChannelID)
			}
		}
	})
	t.Run("unknown id", func(t *testing.T) {
		_, err := client.AddChannel(&channels.AddChannelParams{
			Body: channels.AddChannelBody{
				Summary:  gofakeit.Quote(),
				Disabled: gofakeit.Bool(),
				EmailConfig: &channels.AddChannelParamsBodyEmailConfig{
					SendResolved: false,
					To:           []string{gofakeit.Email()},
				},
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		_, err = client.RemoveChannel(&channels.RemoveChannelParams{
			Body: channels.RemoveChannelBody{
				ChannelID: gofakeit.UUID(),
			},
			Context: pmmapitests.Context,
		})
		require.Error(t, err)
	})
}

func TestListChannels(t *testing.T) {
	// if !pmmapitests.RunIATests {
	// 	t.Skip("Skipping IA tests until IA will out of beta: https://jira.percona.com/browse/PMM-7001")
	// }

	client := channelsClient.Default.Channels

	summary := gofakeit.UUID()
	email := gofakeit.Email()
	_, err := client.AddChannel(&channels.AddChannelParams{
		Body: channels.AddChannelBody{
			Summary:  summary,
			Disabled: gofakeit.Bool(),
			EmailConfig: &channels.AddChannelParamsBodyEmailConfig{
				SendResolved: true,
				To:           []string{email},
			},
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)

	resp, err := client.ListChannels(&channels.ListChannelsParams{Context: pmmapitests.Context})
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Payload.Channels)
	var found bool
	for _, channel := range resp.Payload.Channels {
		if channel.Summary == summary {
			assert.Equal(t, []string{email}, channel.EmailConfig.To)
			assert.True(t, channel.EmailConfig.SendResolved)
			found = true
		}
	}

	assert.True(t, found, "Expected channel not found")
}
