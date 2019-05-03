package management

import (
	"testing"

	"github.com/AlekSi/pointer"
	inventoryClient "github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/agents"
	"github.com/percona/pmm/api/inventorypb/json/client/services"
	"github.com/percona/pmm/api/managementpb/json/client"
	mongodb "github.com/percona/pmm/api/managementpb/json/client/mongo_db"
	"github.com/percona/pmm/api/managementpb/json/client/node"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAddMongoDB(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {

		nodeName := pmmapitests.TestString(t, "node-for-basic-name")
		nodeID, pmmAgentID := registerGenericNode(t, node.RegisterBody{
			NodeName: nodeName,
			NodeType: pointer.ToString(node.RegisterBodyNodeTypeGENERICNODE),
		})
		defer pmmapitests.RemoveNodes(t, nodeID)
		defer removePMMAgentWithSubAgents(t, pmmAgentID)

		serviceName := pmmapitests.TestString(t, "service-name-for-basic-name")

		params := &mongodb.AddMongoDBParams{
			Context: pmmapitests.Context,
			Body: mongodb.AddMongoDBBody{
				NodeID:      nodeID,
				PMMAgentID:  pmmAgentID,
				ServiceName: serviceName,
				Address:     "10.10.10.10",
				Port:        27017,
			},
		}
		addMongoDBOK, err := client.Default.MongoDB.AddMongoDB(params)
		require.NoError(t, err)
		require.NotNil(t, addMongoDBOK)
		require.NotNil(t, addMongoDBOK.Payload.Service)
		serviceID := addMongoDBOK.Payload.Service.ServiceID
		defer pmmapitests.RemoveServices(t, serviceID)

		// Check that service is created and its fields.
		serviceOK, err := inventoryClient.Default.Services.GetService(&services.GetServiceParams{
			Body: services.GetServiceBody{
				ServiceID: serviceID,
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		require.NotNil(t, serviceOK)
		assert.Equal(t, services.GetServiceOKBody{
			Mongodb: &services.GetServiceOKBodyMongodb{
				ServiceID:   serviceID,
				NodeID:      nodeID,
				ServiceName: serviceName,
				Address:     "10.10.10.10",
				Port:        27017,
			},
		}, *serviceOK.Payload)

		// Check that no one exporter is added.
		listAgents, err := inventoryClient.Default.Agents.ListAgents(&agents.ListAgentsParams{
			Context: pmmapitests.Context,
			Body: agents.ListAgentsBody{
				ServiceID: serviceID,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, agents.ListAgentsOKBody{}, *listAgents.Payload)
		defer removeAllAgentsInList(t, listAgents)
	})

	t.Run("All fields", func(t *testing.T) {
		nodeName := pmmapitests.TestString(t, "node-name-for-all-fields")
		nodeID, pmmAgentID := registerGenericNode(t, node.RegisterBody{
			NodeName: nodeName,
			NodeType: pointer.ToString(node.RegisterBodyNodeTypeGENERICNODE),
		})
		defer pmmapitests.RemoveNodes(t, nodeID)
		defer removePMMAgentWithSubAgents(t, pmmAgentID)

		serviceName := pmmapitests.TestString(t, "service-name-for-all-fields")

		params := &mongodb.AddMongoDBParams{
			Context: pmmapitests.Context,
			Body: mongodb.AddMongoDBBody{
				NodeID:             nodeID,
				PMMAgentID:         pmmAgentID,
				ServiceName:        serviceName,
				Address:            "10.10.10.10",
				Port:               27017,
				Username:           "username",
				Password:           "password",
				Environment:        "some-environment",
				Cluster:            "cluster-name",
				ReplicationSet:     "replication-set",
				MongodbExporter:    true,
				QANMongodbProfiler: true,
				CustomLabels:       map[string]string{"bar": "foo"},
			},
		}
		addMongoDBOK, err := client.Default.MongoDB.AddMongoDB(params)
		require.NoError(t, err)
		require.NotNil(t, addMongoDBOK)
		require.NotNil(t, addMongoDBOK.Payload.Service)
		serviceID := addMongoDBOK.Payload.Service.ServiceID
		defer pmmapitests.RemoveServices(t, serviceID)

		// Check that service is created and its fields.
		serviceOK, err := inventoryClient.Default.Services.GetService(&services.GetServiceParams{
			Body: services.GetServiceBody{
				ServiceID: serviceID,
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.NotNil(t, serviceOK)
		assert.Equal(t, services.GetServiceOKBody{
			Mongodb: &services.GetServiceOKBodyMongodb{
				ServiceID:      serviceID,
				NodeID:         nodeID,
				ServiceName:    serviceName,
				Address:        "10.10.10.10",
				Port:           27017,
				Environment:    "some-environment",
				Cluster:        "cluster-name",
				ReplicationSet: "replication-set",
				CustomLabels:   map[string]string{"bar": "foo"},
			},
		}, *serviceOK.Payload)

		// Check that no one exporter is added.
		listAgents, err := inventoryClient.Default.Agents.ListAgents(&agents.ListAgentsParams{
			Context: pmmapitests.Context,
			Body: agents.ListAgentsBody{
				ServiceID: serviceID,
			},
		})
		assert.NoError(t, err)
		require.NotNil(t, listAgents)
		defer removeAllAgentsInList(t, listAgents)

		require.Len(t, listAgents.Payload.MongodbExporter, 1)
		require.Len(t, listAgents.Payload.QANMongodbProfilerAgent, 1)
		assert.Equal(t, agents.ListAgentsOKBody{
			MongodbExporter: []*agents.MongodbExporterItems0{
				{
					AgentID:    listAgents.Payload.MongodbExporter[0].AgentID,
					ServiceID:  serviceID,
					PMMAgentID: pmmAgentID,
					Username:   "username",
					Password:   "password",
				},
			},
			QANMongodbProfilerAgent: []*agents.QANMongodbProfilerAgentItems0{
				{
					AgentID:    listAgents.Payload.QANMongodbProfilerAgent[0].AgentID,
					ServiceID:  serviceID,
					PMMAgentID: pmmAgentID,
					Username:   "username",
					Password:   "password",
				},
			},
		}, *listAgents.Payload)
	})

	t.Run("Empty Node ID", func(t *testing.T) {
		params := &mongodb.AddMongoDBParams{
			Context: pmmapitests.Context,
			Body:    mongodb.AddMongoDBBody{},
		}
		addMongoDBOK, err := client.Default.MongoDB.AddMongoDB(params)
		pmmapitests.AssertEqualAPIError(t, err, pmmapitests.ServerResponse{Code: 400, Error: "invalid field NodeId: value '' must not be an empty string"})
		assert.Nil(t, addMongoDBOK)
	})

	t.Run("Empty Service Name", func(t *testing.T) {
		nodeName := pmmapitests.TestString(t, "node-name")
		nodeID, pmmAgentID := registerGenericNode(t, node.RegisterBody{
			NodeName: nodeName,
			NodeType: pointer.ToString(node.RegisterBodyNodeTypeGENERICNODE),
		})
		defer pmmapitests.RemoveNodes(t, nodeID)
		defer removePMMAgentWithSubAgents(t, pmmAgentID)

		params := &mongodb.AddMongoDBParams{
			Context: pmmapitests.Context,
			Body:    mongodb.AddMongoDBBody{NodeID: nodeID},
		}
		addMongoDBOK, err := client.Default.MongoDB.AddMongoDB(params)
		pmmapitests.AssertEqualAPIError(t, err, pmmapitests.ServerResponse{Code: 400, Error: "invalid field ServiceName: value '' must not be an empty string"})
		assert.Nil(t, addMongoDBOK)
	})

	t.Run("Empty Address", func(t *testing.T) {
		nodeName := pmmapitests.TestString(t, "node-name")
		nodeID, pmmAgentID := registerGenericNode(t, node.RegisterBody{
			NodeName: nodeName,
			NodeType: pointer.ToString(node.RegisterBodyNodeTypeGENERICNODE),
		})
		defer pmmapitests.RemoveNodes(t, nodeID)
		defer removePMMAgentWithSubAgents(t, pmmAgentID)

		serviceName := pmmapitests.TestString(t, "service-name")
		params := &mongodb.AddMongoDBParams{
			Context: pmmapitests.Context,
			Body: mongodb.AddMongoDBBody{
				NodeID:      nodeID,
				ServiceName: serviceName,
			},
		}
		addMongoDBOK, err := client.Default.MongoDB.AddMongoDB(params)
		pmmapitests.AssertEqualAPIError(t, err, pmmapitests.ServerResponse{Code: 400, Error: "invalid field Address: value '' must not be an empty string"})
		assert.Nil(t, addMongoDBOK)
	})

	t.Run("Empty Port", func(t *testing.T) {
		nodeName := pmmapitests.TestString(t, "node-name")
		nodeID, pmmAgentID := registerGenericNode(t, node.RegisterBody{
			NodeName: nodeName,
			NodeType: pointer.ToString(node.RegisterBodyNodeTypeGENERICNODE),
		})
		defer pmmapitests.RemoveNodes(t, nodeID)
		defer removePMMAgentWithSubAgents(t, pmmAgentID)

		serviceName := pmmapitests.TestString(t, "service-name")
		params := &mongodb.AddMongoDBParams{
			Context: pmmapitests.Context,
			Body: mongodb.AddMongoDBBody{
				NodeID:      nodeID,
				ServiceName: serviceName,
				Address:     "10.10.10.10",
			},
		}
		addMongoDBOK, err := client.Default.MongoDB.AddMongoDB(params)
		pmmapitests.AssertEqualAPIError(t, err, pmmapitests.ServerResponse{Code: 400, Error: "invalid field Port: value '0' must be greater than '0'"})
		assert.Nil(t, addMongoDBOK)
	})

	t.Run("Empty Pmm Agent ID", func(t *testing.T) {
		nodeName := pmmapitests.TestString(t, "node-name")
		nodeID, pmmAgentID := registerGenericNode(t, node.RegisterBody{
			NodeName: nodeName,
			NodeType: pointer.ToString(node.RegisterBodyNodeTypeGENERICNODE),
		})
		defer pmmapitests.RemoveNodes(t, nodeID)
		defer removePMMAgentWithSubAgents(t, pmmAgentID)

		serviceName := pmmapitests.TestString(t, "service-name")
		params := &mongodb.AddMongoDBParams{
			Context: pmmapitests.Context,
			Body: mongodb.AddMongoDBBody{
				NodeID:      nodeID,
				ServiceName: serviceName,
				Address:     "10.10.10.10",
				Port:        3306,
			},
		}
		addMongoDBOK, err := client.Default.MongoDB.AddMongoDB(params)
		pmmapitests.AssertEqualAPIError(t, err, pmmapitests.ServerResponse{Code: 400, Error: "invalid field PmmAgentId: value '' must not be an empty string"})
		assert.Nil(t, addMongoDBOK)
	})
}
