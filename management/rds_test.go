package management

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/rds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRDSDiscovery(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		accessKey, secretKey := os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY")
		if accessKey == "" || secretKey == "" {
			// TODO remove skip once secrets are added
			t.Skip("Environment variables AWS_ACCESS_KEY / AWS_SECRET_KEY are not defined, skipping test")
		}

		params := &rds.DiscoverRDSParams{
			Body: rds.DiscoverRDSBody{
				AWSAccessKey: accessKey,
				AWSSecretKey: secretKey,
			},
			Context: pmmapitests.Context,
		}
		discoverOK, err := client.Default.RDS.DiscoverRDS(params)
		require.NoError(t, err)
		require.NotNil(t, discoverOK.Payload)
		assert.NotEmpty(t, discoverOK.Payload.RDSInstances)

		// TODO Better tests: https://jira.percona.com/browse/PMM-4896
	})
}

func TestAddRds(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Run("BasicAddRDS", func(t *testing.T) {
		params := &rds.AddRDSParams{
			Body: rds.AddRDSBody{
				Region:                    "region",
				Az:                        "az",
				InstanceID:                "d752f1a9-31c9-4b8c-bb2d-d26bc000001",
				NodeModel:                 "some-model",
				Address:                   "some.example.rds",
				Port:                      3306,
				Engine:                    pointer.ToString("DISCOVER_RDS_MYSQL"),
				NodeName:                  "some-node-name-000001",
				ServiceName:               "test-add-rds-service000001",
				Environment:               "some-env",
				Cluster:                   "cluster-01",
				ReplicationSet:            "rs-01",
				Username:                  "some-username",
				Password:                  "some-password",
				AWSAccessKey:              "my-aws-access-key",
				AWSSecretKey:              "my-aws-secret-key",
				RDSExporter:               true,
				QANMysqlPerfschema:        true,
				CustomLabels:              map[string]string{},
				SkipConnectionCheck:       true,
				TLS:                       false,
				TLSSkipVerify:             false,
				DisableQueryExamples:      false,
				TablestatsGroupTableLimit: 2000,
			},
			Context: pmmapitests.Context,
		}
		addRDSOK, err := client.Default.RDS.AddRDS(params)
		require.NoError(t, err)
		require.NotNil(t, addRDSOK.Payload)

		body := addRDSOK.Payload
		pmmapitests.RemoveAgents(t, body.MysqldExporter.AgentID)
		pmmapitests.RemoveAgents(t, body.QANMysqlPerfschema.AgentID)
		pmmapitests.RemoveAgents(t, body.RDSExporter.AgentID)
		pmmapitests.RemoveServices(t, body.Mysql.ServiceID)
		pmmapitests.RemoveNodes(t, body.Mysql.NodeID)
	})
}
