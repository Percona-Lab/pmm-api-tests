package management

import (
	"testing"

	"github.com/AlekSi/pointer"
	inventoryClient "github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/agents"
	"github.com/percona/pmm/api/inventorypb/json/client/services"
	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/annotation"
	mysql "github.com/percona/pmm/api/managementpb/json/client/my_sql"
	"github.com/percona/pmm/api/managementpb/json/client/node"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAddAnnotation(t *testing.T) {
	t.Run("Add Basic Annotation", func(t *testing.T) {
		params := &annotation.AddAnnotationParams{
			Body: annotation.AddAnnotationBody{
				Text: "Annotation Text",
				Tags: []string{"tag1", "tag2"},
			},
			Context: pmmapitests.Context,
		}
		_, err := client.Default.Annotation.AddAnnotation(params)
		require.NoError(t, err)
	})

	t.Run("Add Empty Annotation", func(t *testing.T) {
		params := &annotation.AddAnnotationParams{
			Body: annotation.AddAnnotationBody{
				Text: "",
				Tags: []string{},
			},
			Context: pmmapitests.Context,
		}
		_, err := client.Default.Annotation.AddAnnotation(params)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field Text: value '' must not be an empty string")
	})

	t.Run("Non-existing service", func(t *testing.T) {
		params := &annotation.AddAnnotationParams{
			Body: annotation.AddAnnotationBody{
				Text:         "Some text",
				ServiceNames: []string{"no-service"},
			},
			Context: pmmapitests.Context,
		}
		_, err := client.Default.Annotation.AddAnnotation(params)
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, `Service with name "no-service" not found.`)
	})

	t.Run("Non-existing node", func(t *testing.T) {
		params := &annotation.AddAnnotationParams{
			Body: annotation.AddAnnotationBody{
				Text:     "Some text",
				NodeName: "no-node",
			},
			Context: pmmapitests.Context,
		}
		_, err := client.Default.Annotation.AddAnnotation(params)
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, `Node with name "no-node" not found.`)
	})

	t.Run("Existing service", func(t *testing.T) {
		body := node.RegisterNodeBody{
			NodeName:   "test-node",
			NodeType:   pointer.ToString(node.RegisterNodeBodyNodeTypeGENERICNODE),
			Address:    "node-address",
			Region:     "region",
			Reregister: true,
		}
		paramsRegister := node.RegisterNodeParams{
			Context: pmmapitests.Context,
			Body:    body,
		}
		node, err := client.Default.Node.RegisterNode(&paramsRegister)
		require.NoError(t, err)

		nodeID := node.Payload.GenericNode.NodeID
		pmmAgentID := node.Payload.PMMAgent.AgentID

		defer pmmapitests.RemoveNodes(t, nodeID)
		defer pmmapitests.RemoveAgents(t, pmmAgentID)
		nodeExporterAgentID, ok := assertNodeExporterCreated(t, node.Payload.PMMAgent.AgentID)
		if ok {
			defer pmmapitests.RemoveAgents(t, nodeExporterAgentID)
		}

		params := &mysql.AddMySQLParams{
			Context: pmmapitests.Context,
			Body: mysql.AddMySQLBody{
				NodeID:      nodeID,
				PMMAgentID:  pmmAgentID,
				ServiceName: "test-service-mysql",
				Address:     "10.10.10.10",
				Port:        3306,
				Username:    "username",

				SkipConnectionCheck: true,
			},
		}
		addMySQLOK, err := client.Default.MySQL.AddMySQL(params)
		require.NoError(t, err)
		require.NotNil(t, addMySQLOK)
		require.NotNil(t, addMySQLOK.Payload.Service)
		serviceID := addMySQLOK.Payload.Service.ServiceID
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
			Mysql: &services.GetServiceOKBodyMysql{
				ServiceID:   serviceID,
				NodeID:      nodeID,
				ServiceName: "test-service-mysql",
				Address:     "10.10.10.10",
				Port:        3306,
			},
		}, *serviceOK.Payload)

		// Check that mysqld exporter is added by default.
		listAgents, err := inventoryClient.Default.Agents.ListAgents(&agents.ListAgentsParams{
			Context: pmmapitests.Context,
			Body: agents.ListAgentsBody{
				ServiceID: serviceID,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, agents.ListAgentsOKBody{
			MysqldExporter: []*agents.MysqldExporterItems0{
				{
					AgentID:                   listAgents.Payload.MysqldExporter[0].AgentID,
					ServiceID:                 serviceID,
					PMMAgentID:                pmmAgentID,
					Username:                  "username",
					TablestatsGroupTableLimit: 1000,
				},
			},
		}, *listAgents.Payload)
		defer removeAllAgentsInList(t, listAgents)

		paramsAdd := &annotation.AddAnnotationParams{
			Body: annotation.AddAnnotationBody{
				Text:         "Some text",
				ServiceNames: []string{"test-service-mysql"},
			},
			Context: pmmapitests.Context,
		}
		_, err = client.Default.Annotation.AddAnnotation(paramsAdd)
		require.NoError(t, err)
	})

	t.Run("Existing node", func(t *testing.T) {
		body := node.RegisterNodeBody{
			NodeName:   "test-node",
			NodeType:   pointer.ToString(node.RegisterNodeBodyNodeTypeGENERICNODE),
			Address:    "node-address",
			Region:     "region",
			Reregister: true,
		}
		paramsRegister := node.RegisterNodeParams{
			Context: pmmapitests.Context,
			Body:    body,
		}
		node, err := client.Default.Node.RegisterNode(&paramsRegister)
		require.NoError(t, err)

		defer pmmapitests.RemoveNodes(t, node.Payload.GenericNode.NodeID)
		defer pmmapitests.RemoveAgents(t, node.Payload.PMMAgent.AgentID)
		nodeExporterAgentID, ok := assertNodeExporterCreated(t, node.Payload.PMMAgent.AgentID)
		if ok {
			defer pmmapitests.RemoveAgents(t, nodeExporterAgentID)
		}

		paramsAdd := &annotation.AddAnnotationParams{
			Body: annotation.AddAnnotationBody{
				Text:     "Some text",
				NodeName: "test-node",
			},
			Context: pmmapitests.Context,
		}
		_, err = client.Default.Annotation.AddAnnotation(paramsAdd)
		require.NoError(t, err)
	})
}
