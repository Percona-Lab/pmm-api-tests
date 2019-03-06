package inventory

import (
	"testing"

	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Percona-Lab/pmm-api-tests"
)

func TestServices(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		t.Parallel()
		node := addRemoteNode(t, withUUID(t, "Remote node for services test"))
		nodeID := node.Remote.NodeID
		defer removeNodes(t, nodeID)

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "Some MySQL Service"),
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		remoteService := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      nodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "Some MySQL Service on remote Node"),
		})
		remoteServiceID := remoteService.Mysql.ServiceID
		defer removeServices(t, remoteServiceID)

		res, err := client.Default.Services.ListServices(&services.ListServicesParams{Context: pmmapitests.Context})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotZerof(t, len(res.Payload.Mysql), "There should be at least one node")
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.Mysql {
				if v.ServiceID == serviceID {
					return true
				}
			}
			return false
		}, "There should be MySQL service with id `%s`", serviceID)
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.Mysql {
				if v.ServiceID == remoteServiceID {
					return true
				}
			}
			return false
		}, "There should be MySQL service with id `%s`", remoteServiceID)
	})

	t.Run("FilterList", func(t *testing.T) {
		t.Skip("Have not implemented yet.")
		t.Parallel()
		node := addRemoteNode(t, withUUID(t, "Remote node to check services filter"))
		nodeID := node.Remote.NodeID
		defer removeNodes(t, nodeID)

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "Some MySQL Service for filters test"),
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		remoteService := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      nodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "Some MySQL Service on remote Node for filters test"),
		})
		remoteServiceID := remoteService.Mysql.ServiceID
		defer removeServices(t, remoteServiceID)

		res, err := client.Default.Services.ListServices(&services.ListServicesParams{
			Body:    services.ListServicesBody{NodeID: nodeID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotZerof(t, len(res.Payload.Mysql), "There should be at least one node")
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.Mysql {
				if v.ServiceID == serviceID {
					return false
				}
			}
			return true
		}, "There should not be MySQL service with id `%s`", serviceID)
		require.Conditionf(t, func() (success bool) {
			for _, v := range res.Payload.Mysql {
				if v.ServiceID == remoteServiceID {
					return true
				}
			}
			return false
		}, "There should be MySQL service with id `%s`", remoteServiceID)
	})
}

func TestGetService(t *testing.T) {
	t.Run("NotFound", func(t *testing.T) {
		params := &services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: "pmm-not-found"},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.GetService(params)
		assertEqualAPIError(t, err, 404)
		assert.Nil(t, res)
	})

	t.Run("EmptyServiceID", func(t *testing.T) {
		params := &services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: ""},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.GetService(params)
		assertEqualAPIError(t, err, 400)
		assert.Nil(t, res)
	})
}

func TestMySQLService(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()
		serviceName := withUUID(t, "Basic MySQL Service")
		params := &services.AddMySQLServiceParams{
			Body: services.AddMySQLServiceBody{
				NodeID:      "pmm-server",
				Address:     "localhost",
				Port:        3306,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddMySQLService(params)
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotNil(t, res.Payload.Mysql)
		require.NotEmpty(t, res.Payload.Mysql.ServiceID)
		serviceID := res.Payload.Mysql.ServiceID
		defer removeServices(t, serviceID)
		require.Equal(t, "pmm-server", res.Payload.Mysql.NodeID)
		require.Equal(t, "localhost", res.Payload.Mysql.Address)
		require.Equal(t, int64(3306), res.Payload.Mysql.Port)
		require.Equal(t, serviceName, res.Payload.Mysql.ServiceName)

		// Check if the service saved in PMM-Managed.
		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, serviceRes)
		require.NotNil(t, serviceRes.Payload.Mysql)
		require.Nil(t, serviceRes.Payload.AmazonRDSMysql)
		require.NotEmpty(t, serviceRes.Payload.Mysql.ServiceID)
		require.Equal(t, "pmm-server", serviceRes.Payload.Mysql.NodeID)
		require.Equal(t, "localhost", serviceRes.Payload.Mysql.Address)
		require.Equal(t, int64(3306), serviceRes.Payload.Mysql.Port)
		require.Equal(t, serviceName, serviceRes.Payload.Mysql.ServiceName)

		// Check duplicates.
		params = &services.AddMySQLServiceParams{
			Body: services.AddMySQLServiceBody{
				NodeID:      "pmm-server",
				Address:     "127.0.0.1",
				Port:        3336,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err = client.Default.Services.AddMySQLService(params)
		assertEqualAPIError(t, err, 409)
		assert.Nil(t, res)
	})

	t.Run("ChangeMySQLServiceName", func(t *testing.T) {
		t.Parallel()
		body := services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "MySQL Service to change name"),
		}
		service := addMySQLService(t, body)
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		// Change MySQL service name.
		changedServiceName := withUUID(t, "Changed MySQL Service")
		changeRes, err := client.Default.Services.ChangeMySQLService(&services.ChangeMySQLServiceParams{
			Body: services.ChangeMySQLServiceBody{
				ServiceID:   serviceID,
				ServiceName: changedServiceName,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, changeRes.Payload.Mysql)
		require.Equal(t, serviceID, changeRes.Payload.Mysql.ServiceID)
		require.Equal(t, changedServiceName, changeRes.Payload.Mysql.ServiceName)

		// Check changes in backend.
		changedService, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.Equal(t, changedServiceName, changedService.Payload.Mysql.ServiceName)
		// Check that other fields isn't changed.
		require.Equal(t, serviceRes.Payload.Mysql.Port, changedService.Payload.Mysql.Port)
		require.Equal(t, serviceRes.Payload.Mysql.Address, changedService.Payload.Mysql.Address)
		require.Equal(t, serviceRes.Payload.Mysql.NodeID, changedService.Payload.Mysql.NodeID)
	})

	t.Run("ChangeMySQLServicePort", func(t *testing.T) { //TODO: will we implement this?
		t.Skip("Not implemented yet.")
		body := services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "MySQL Service to change port"),
		}
		service := addMySQLService(t, body)
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		// Change MySQL service name.
		newPort := int64(3337)
		changeRes, err := client.Default.Services.ChangeMySQLService(&services.ChangeMySQLServiceParams{
			Body: services.ChangeMySQLServiceBody{
				ServiceID: serviceID,
				Port:      newPort,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, changeRes.Payload.Mysql)
		require.Equal(t, serviceID, changeRes.Payload.Mysql.ServiceID)
		require.Equal(t, newPort, changeRes.Payload.Mysql.Port)

		// Check changes in backend.
		changedService, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.Equal(t, newPort, changedService.Payload.Mysql.Port)
		// Check that other fields isn't changed.
		require.Equal(t, serviceRes.Payload.Mysql.ServiceName, changedService.Payload.Mysql.ServiceName)
		require.Equal(t, serviceRes.Payload.Mysql.Address, changedService.Payload.Mysql.Address)
		require.Equal(t, serviceRes.Payload.Mysql.NodeID, changedService.Payload.Mysql.NodeID)
	})

	t.Run("AddNodeIDEmpty", func(t *testing.T) {
		t.Parallel()
		params := &services.AddMySQLServiceParams{
			Body: services.AddMySQLServiceBody{
				NodeID:      "",
				Address:     "localhost",
				Port:        3306,
				ServiceName: withUUID(t, "MySQL Service with empty node id"),
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddMySQLService(params)
		assertEqualAPIError(t, err, 400)
		assert.Nil(t, res)
	})
}

func TestAmazonRDSMySQLService(t *testing.T) {
	t.Skip("Not implemented yet.")
	t.Run("Basic", func(t *testing.T) {
		serviceName := withUUID(t, "Basic AmazonRDSMySQL Service")
		params := &services.AddAmazonRDSMySQLServiceParams{
			Body: services.AddAmazonRDSMySQLServiceBody{
				NodeID:      "pmm-server",
				Address:     "localhost",
				Port:        3306,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddAmazonRDSMySQLService(params)
		require.NoError(t, err)
		require.NotNil(t, res)
		require.NotNil(t, res.Payload.AmazonRDSMysql)
		require.NotEmpty(t, res.Payload.AmazonRDSMysql.ServiceID)
		serviceID := res.Payload.AmazonRDSMysql.ServiceID
		defer removeServices(t, serviceID)
		require.Equal(t, "pmm-server", res.Payload.AmazonRDSMysql.NodeID)
		require.Equal(t, "localhost", res.Payload.AmazonRDSMysql.Address)
		require.Equal(t, int64(3306), res.Payload.AmazonRDSMysql.Port)
		require.Equal(t, serviceName, res.Payload.AmazonRDSMysql.ServiceName)

		// Check if the service saved in PMM-Managed.
		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, serviceRes)
		require.NotNil(t, serviceRes.Payload.AmazonRDSMysql)
		require.Nil(t, serviceRes.Payload.Mysql)
		require.NotEmpty(t, serviceRes.Payload.AmazonRDSMysql.ServiceID)
		require.Equal(t, "pmm-server", serviceRes.Payload.AmazonRDSMysql.NodeID)
		require.Equal(t, "localhost", serviceRes.Payload.AmazonRDSMysql.Address)
		require.Equal(t, int64(3306), serviceRes.Payload.AmazonRDSMysql.Port)
		require.Equal(t, serviceName, serviceRes.Payload.AmazonRDSMysql.ServiceName)

		// Check duplicates.
		params = &services.AddAmazonRDSMySQLServiceParams{
			Body: services.AddAmazonRDSMySQLServiceBody{
				NodeID:      "pmm-server",
				Address:     "127.0.0.1",
				Port:        3336,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err = client.Default.Services.AddAmazonRDSMySQLService(params)
		assertEqualAPIError(t, err, 409)
		assert.Nil(t, res)
	})
}
