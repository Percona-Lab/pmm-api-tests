package management

import (
	"os"
	"testing"

	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/discovery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestRDSDiscovery(t *testing.T) {
	t.Skip("Need to configure Jenkins & Travis to handle credentials")
	t.Run("Basic", func(t *testing.T) {
		awsAccesKey := os.Getenv("AWS_ACCESS_KEY_ID")
		awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		params := &discovery.DiscoverRDSParams{
			Context: pmmapitests.Context,
			Body: discovery.DiscoverRDSBody{
				AWSAccessKey: awsAccesKey,
				AWSSecretKey: awsSecretKey,
			},
		}
		discoverOK, err := client.Default.Discovery.DiscoverRDS(params)
		require.NoError(t, err)
		require.NotNil(t, discoverOK)
		require.NotNil(t, discoverOK.Payload.RDSInstances)
		instances := discoverOK.Payload.RDSInstances
		assert.NotEmpty(t, instances)
	})
}
