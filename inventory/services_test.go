package inventory

import (
	"fmt"
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
		genericNodeOKBody := addGenericNode(t, withUUID(t, "Generic node for services test"))
		genericNodeID := genericNodeOKBody.Generic.NodeID
		defer removeNodes(t, genericNodeID)
		remoteNodeOKBody := addRemoteNode(t, withUUID(t, "Remote node for services test"))
		remoteNodeID := remoteNodeOKBody.Remote.NodeID
		defer removeNodes(t, remoteNodeID)

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      genericNodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "Some MySQL Service"),
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		remoteService := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      remoteNodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "Some MySQL Service on remote Node"),
		})
		remoteServiceID := remoteService.Mysql.ServiceID
		defer removeServices(t, remoteServiceID)

		res, err := client.Default.Services.ListServices(&services.ListServicesParams{Context: pmmapitests.Context})
		assert.NoError(t, err)
		require.NotNil(t, res)
		assert.NotZerof(t, len(res.Payload.Mysql), "There should be at least one node")
		assertMySQLServiceExists(t, res, serviceID)
		assertMySQLServiceExists(t, res, remoteServiceID)
	})

	t.Run("FilterList", func(t *testing.T) {
		t.Parallel()
		genericNodeOKBody := addGenericNode(t, withUUID(t, "Generic node for services test"))
		genericNodeID := genericNodeOKBody.Generic.NodeID
		defer removeNodes(t, genericNodeID)
		remoteNodeOKBody := addRemoteNode(t, withUUID(t, "Remote node to check services filter"))
		remoteNodeID := remoteNodeOKBody.Remote.NodeID
		defer removeNodes(t, remoteNodeID)

		service := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      genericNodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "Some MySQL Service for filters test"),
		})
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		remoteService := addMySQLService(t, services.AddMySQLServiceBody{
			NodeID:      remoteNodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: withUUID(t, "Some MySQL Service on remote Node for filters test"),
		})
		remoteServiceID := remoteService.Mysql.ServiceID
		defer removeServices(t, remoteServiceID)

		res, err := client.Default.Services.ListServices(&services.ListServicesParams{
			Body:    services.ListServicesBody{NodeID: remoteNodeID},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		require.NotNil(t, res)
		assert.NotZerof(t, len(res.Payload.Mysql), "There should be at least one node")
		assertMySQLServiceNotExist(t, res, serviceID)
		assertMySQLServiceExists(t, res, remoteServiceID)
	})
}

func TestGetService(t *testing.T) {
	t.Run("NotFound", func(t *testing.T) {
		params := &services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: "pmm-not-found"},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.GetService(params)
		assertEqualAPIError(t, err, ServerResponse{404, "Service with ID \"pmm-not-found\" not found."})
		assert.Nil(t, res)
	})

	t.Run("EmptyServiceID", func(t *testing.T) {
		params := &services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: ""},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.GetService(params)
		assertEqualAPIError(t, err, ServerResponse{400, "Empty Service ID."})
		assert.Nil(t, res)
	})
}

func TestMySQLService(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()
		genericNodeOKBody := addGenericNode(t, withUUID(t, "Generic node for services test"))
		genericNodeID := genericNodeOKBody.Generic.NodeID
		defer removeNodes(t, genericNodeID)

		serviceName := withUUID(t, "Basic MySQL Service")
		params := &services.AddMySQLServiceParams{
			Body: services.AddMySQLServiceBody{
				NodeID:      genericNodeID,
				Address:     "localhost",
				Port:        3306,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddMySQLService(params)
		assert.NoError(t, err)
		require.NotNil(t, res)
		serviceID := res.Payload.Mysql.ServiceID
		assert.Equal(t, &services.AddMySQLServiceOK{
			Payload: &services.AddMySQLServiceOKBody{
				Mysql: &services.AddMySQLServiceOKBodyMysql{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: serviceName,
				},
			},
		}, res)
		defer removeServices(t, serviceID)

		// Check if the service saved in PMM-Managed.
		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.NotNil(t, serviceRes)
		assert.Equal(t, &services.GetServiceOK{
			Payload: &services.GetServiceOKBody{
				Mysql: &services.GetServiceOKBodyMysql{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: serviceName,
				},
			},
		}, serviceRes)

		// Check duplicates.
		params = &services.AddMySQLServiceParams{
			Body: services.AddMySQLServiceBody{
				NodeID:      genericNodeID,
				Address:     "127.0.0.1",
				Port:        3336,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err = client.Default.Services.AddMySQLService(params)
		assertEqualAPIError(t, err, ServerResponse{409, fmt.Sprintf("Service with name \"%s\" already exists.", serviceName)})
		if !assert.Nil(t, res) {
			removeServices(t, res.Payload.Mysql.ServiceID)
		}
	})

	t.Run("ChangeMySQLServiceName", func(t *testing.T) {
		t.Parallel()
		genericNodeOKBody := addGenericNode(t, withUUID(t, "Generic node for services test"))
		genericNodeID := genericNodeOKBody.Generic.NodeID
		defer removeNodes(t, genericNodeID)

		serviceName := withUUID(t, "MySQL Service to change name")
		body := services.AddMySQLServiceBody{
			NodeID:      genericNodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: serviceName,
		}
		service := addMySQLService(t, body)
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.Equal(t, &services.GetServiceOK{
			Payload: &services.GetServiceOKBody{
				Mysql: &services.GetServiceOKBodyMysql{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: serviceName,
				},
			},
		}, serviceRes)

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
		assert.Equal(t, &services.ChangeMySQLServiceOK{
			Payload: &services.ChangeMySQLServiceOKBody{
				Mysql: &services.ChangeMySQLServiceOKBodyMysql{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: changedServiceName,
				},
			},
		}, changeRes)

		// Check changes in backend.
		changedService, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.Equal(t, &services.GetServiceOK{
			Payload: &services.GetServiceOKBody{
				Mysql: &services.GetServiceOKBodyMysql{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: changedServiceName,
				},
			},
		}, changedService)
	})

	t.Run("ChangeMySQLServicePort", func(t *testing.T) {
		t.Skip("Not implemented yet.")

		genericNode := addGenericNode(t, withUUID(t, "Test Remote Node for List"))
		genericNodeID := genericNode.Generic.NodeID
		defer removeNodes(t, genericNodeID)

		serviceName := withUUID(t, "MySQL Service to change port")
		body := services.AddMySQLServiceBody{
			NodeID:      genericNodeID,
			Address:     "localhost",
			Port:        3306,
			ServiceName: serviceName,
		}
		service := addMySQLService(t, body)
		serviceID := service.Mysql.ServiceID
		defer removeServices(t, serviceID)

		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.Equal(t, &services.GetServiceOK{
			Payload: &services.GetServiceOKBody{
				Mysql: &services.GetServiceOKBodyMysql{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: serviceName,
				},
			},
		}, serviceRes)

		// Change MySQL service name.
		newPort := int64(3337)
		changeRes, err := client.Default.Services.ChangeMySQLService(&services.ChangeMySQLServiceParams{
			Body: services.ChangeMySQLServiceBody{
				ServiceID: serviceID,
				Port:      newPort,
			},
			Context: pmmapitests.Context,
		})
		assert.NoError(t, err)
		assert.Equal(t, &services.ChangeMySQLServiceOK{
			Payload: &services.ChangeMySQLServiceOKBody{
				Mysql: &services.ChangeMySQLServiceOKBodyMysql{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					Address:     "localhost",
					Port:        newPort,
					ServiceName: serviceName,
				},
			},
		}, changeRes)

		// Check changes in backend.
		changedService, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		assert.Equal(t, &services.GetServiceOK{
			Payload: &services.GetServiceOKBody{
				Mysql: &services.GetServiceOKBodyMysql{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					Address:     "localhost",
					Port:        newPort,
					ServiceName: serviceName,
				},
			},
		}, changedService)
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
		assertEqualAPIError(t, err, ServerResponse{400, "Empty Node ID."})
		if !assert.Nil(t, res) {
			removeServices(t, res.Payload.Mysql.ServiceID)
		}
	})

	t.Run("AddServiceNameEmpty", func(t *testing.T) {
		t.Parallel()

		genericNodeOKBody := addGenericNode(t, withUUID(t, "Generic node for services test"))
		genericNodeID := genericNodeOKBody.Generic.NodeID
		defer removeNodes(t, genericNodeID)

		params := &services.AddMySQLServiceParams{
			Body: services.AddMySQLServiceBody{
				NodeID:      genericNodeID,
				ServiceName: "",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddMySQLService(params)
		assertEqualAPIError(t, err, ServerResponse{400, ""})
		if !assert.Nil(t, res) {
			removeServices(t, res.Payload.Mysql.ServiceID)
		}
	})
}

func TestAmazonRDSMySQLService(t *testing.T) {
	t.Skip("Not implemented yet.")
	t.Run("Basic", func(t *testing.T) {
		remoteNodeOKBody := addRemoteNode(t, withUUID(t, "Remote node to check services filter"))
		remoteNodeID := remoteNodeOKBody.Remote.NodeID
		defer removeNodes(t, remoteNodeID)

		serviceName := withUUID(t, "Basic AmazonRDSMySQL Service")
		params := &services.AddAmazonRDSMySQLServiceParams{
			Body: services.AddAmazonRDSMySQLServiceBody{
				NodeID:      remoteNodeID,
				Address:     "localhost",
				Port:        3306,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddAmazonRDSMySQLService(params)
		assert.NoError(t, err)
		require.NotNil(t, res)
		serviceID := res.Payload.AmazonRDSMysql.ServiceID
		defer removeServices(t, serviceID)
		assert.Equal(t, &services.AddAmazonRDSMySQLServiceOK{
			Payload: &services.AddAmazonRDSMySQLServiceOKBody{
				AmazonRDSMysql: &services.AddAmazonRDSMySQLServiceOKBodyAmazonRDSMysql{
					ServiceID:   serviceID,
					NodeID:      remoteNodeID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: serviceName,
				},
			},
		}, res)

		// Check if the service saved in PMM-Managed.
		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, serviceRes)
		assert.Equal(t, &services.GetServiceOK{
			Payload: &services.GetServiceOKBody{
				AmazonRDSMysql: &services.GetServiceOKBodyAmazonRDSMysql{
					ServiceID:   serviceID,
					NodeID:      remoteNodeID,
					Address:     "localhost",
					Port:        3306,
					ServiceName: serviceName,
				},
			},
		}, serviceRes)

		// Check duplicates.
		params = &services.AddAmazonRDSMySQLServiceParams{
			Body: services.AddAmazonRDSMySQLServiceBody{
				NodeID:      remoteNodeID,
				Address:     "127.0.0.1",
				Port:        3336,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err = client.Default.Services.AddAmazonRDSMySQLService(params)
		assertEqualAPIError(t, err, ServerResponse{409, ""})
		if !assert.Nil(t, res) {
			removeServices(t, res.Payload.AmazonRDSMysql.ServiceID)
		}
	})
}

func TestMongoService(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()
		genericNodeOKBody := addGenericNode(t, withUUID(t, "Generic node for services test"))
		genericNodeID := genericNodeOKBody.Generic.NodeID
		defer removeNodes(t, genericNodeID)

		serviceName := withUUID(t, "Basic Mongo Service")
		params := &services.AddMongoDBServiceParams{
			Body: services.AddMongoDBServiceBody{
				NodeID:      genericNodeID,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddMongoDBService(params)
		assert.NoError(t, err)
		require.NotNil(t, res)
		serviceID := res.Payload.Mongodb.ServiceID
		assert.Equal(t, &services.AddMongoDBServiceOK{
			Payload: &services.AddMongoDBServiceOKBody{
				Mongodb: &services.AddMongoDBServiceOKBodyMongodb{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					ServiceName: serviceName,
				},
			},
		}, res)
		defer removeServices(t, serviceID)

		// Check if the service saved in PMM-Managed.
		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		require.NotNil(t, serviceRes)
		assert.Equal(t, &services.GetServiceOK{
			Payload: &services.GetServiceOKBody{
				Mongodb: &services.GetServiceOKBodyMongodb{
					ServiceID:   serviceID,
					NodeID:      genericNodeID,
					ServiceName: serviceName,
				},
			},
		}, serviceRes)

		// Check duplicates.
		params = &services.AddMongoDBServiceParams{
			Body: services.AddMongoDBServiceBody{
				NodeID:      genericNodeID,
				ServiceName: serviceName,
			},
			Context: pmmapitests.Context,
		}
		res, err = client.Default.Services.AddMongoDBService(params)
		assertEqualAPIError(t, err, ServerResponse{409, fmt.Sprintf("Service with name \"%s\" already exists.", serviceName)})
		if !assert.Nil(t, res) {
			removeServices(t, res.Payload.Mongodb.ServiceID)
		}
	})

	t.Run("AddNodeIDEmpty", func(t *testing.T) {
		t.Parallel()
		params := &services.AddMongoDBServiceParams{
			Body: services.AddMongoDBServiceBody{
				NodeID:      "",
				ServiceName: withUUID(t, "MongoDB Service with empty node id"),
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddMongoDBService(params)
		assertEqualAPIError(t, err, ServerResponse{400, "Empty Node ID."})
		if !assert.Nil(t, res) {
			removeServices(t, res.Payload.Mongodb.ServiceID)
		}
	})

	t.Run("AddServiceNameEmpty", func(t *testing.T) {
		t.Parallel()

		genericNodeOKBody := addGenericNode(t, withUUID(t, "Generic node for services test"))
		genericNodeID := genericNodeOKBody.Generic.NodeID
		defer removeNodes(t, genericNodeID)

		params := &services.AddMongoDBServiceParams{
			Body: services.AddMongoDBServiceBody{
				NodeID:      genericNodeID,
				ServiceName: "",
			},
			Context: pmmapitests.Context,
		}
		res, err := client.Default.Services.AddMongoDBService(params)
		assertEqualAPIError(t, err, ServerResponse{400, ""})
		if !assert.Nil(t, res) {
			removeServices(t, res.Payload.Mongodb.ServiceID)
		}
	})
}
