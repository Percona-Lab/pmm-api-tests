package inventory

import (
	"context"
	"testing"

	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/services"
	"github.com/stretchr/testify/require"

	_ "github.com/Percona-Lab/pmm-api-tests" // init default client
)

func TestServices(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		t.Parallel()
		body := services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: "Some MySQL Service",
		}
		serviceID := addMySQLService(t, body)
		defer removeServices(t, serviceID)
		res, err := client.Default.Services.ListServices(&services.ListServicesParams{Context: context.TODO()})
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
	})
}

func TestMySQLService(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t.Parallel()
		params := &services.AddMySQLServiceParams{
			Body: services.AddMySQLServiceBody{
				NodeID:      "pmm-server",
				Address:     "localhost",
				Port:        3306,
				ServiceName: "Basic MySQL Service",
			},
			Context: context.TODO(),
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
		require.Equal(t, "Basic MySQL Service", res.Payload.Mysql.ServiceName)

		// Check if the service saved in PMM-Managed.
		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: context.TODO(),
		})
		require.NoError(t, err)
		require.NotNil(t, serviceRes)
		require.NotNil(t, serviceRes.Payload.Mysql)
		require.Nil(t, serviceRes.Payload.AmazonRDSMysql)
		require.NotEmpty(t, serviceRes.Payload.Mysql.ServiceID)
		require.Equal(t, "pmm-server", serviceRes.Payload.Mysql.NodeID)
		require.Equal(t, "localhost", serviceRes.Payload.Mysql.Address)
		require.Equal(t, int64(3306), serviceRes.Payload.Mysql.Port)
		require.Equal(t, "Basic MySQL Service", serviceRes.Payload.Mysql.ServiceName)

		// Check duplicates.
		params = &services.AddMySQLServiceParams{
			Body: services.AddMySQLServiceBody{
				NodeID:      "pmm-server",
				Address:     "127.0.0.1",
				Port:        3336,
				ServiceName: "Basic MySQL Service",
			},
			Context: context.TODO(),
		}
		res, err = client.Default.Services.AddMySQLService(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 409)")
		require.Nil(t, res)
	})

	t.Run("ChangeMySQLServiceName", func(t *testing.T) {
		t.Parallel()
		body := services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: "MySQL Service to change name",
		}
		serviceID := addMySQLService(t, body)
		defer removeServices(t, serviceID)

		service, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: context.TODO(),
		})

		// Change MySQL service name.
		changeRes, err := client.Default.Services.ChangeMySQLService(&services.ChangeMySQLServiceParams{
			Body: services.ChangeMySQLServiceBody{
				ServiceID:   serviceID,
				ServiceName: "Changed MySQL Service",
			},
			Context: context.TODO(),
		})
		require.NoError(t, err)
		require.NotNil(t, changeRes.Payload.Mysql)
		require.Equal(t, serviceID, changeRes.Payload.Mysql.ServiceID)
		require.Equal(t, "Changed MySQL Service", changeRes.Payload.Mysql.ServiceName)

		// Check changes in backend.
		changedService, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: context.TODO(),
		})
		require.Equal(t, "Changed MySQL Service", changedService.Payload.Mysql.ServiceName)
		// Check that other fields isn't changed.
		require.Equal(t, service.Payload.Mysql.Port, changedService.Payload.Mysql.Port)
		require.Equal(t, service.Payload.Mysql.Address, changedService.Payload.Mysql.Address)
		require.Equal(t, service.Payload.Mysql.NodeID, changedService.Payload.Mysql.NodeID)
	})

	t.Run("ChangeMySQLServicePort", func(t *testing.T) { //TODO: will we implement this?
		t.Skip("Not implemented yet.")
		body := services.AddMySQLServiceBody{
			NodeID:      "pmm-server",
			Address:     "localhost",
			Port:        3306,
			ServiceName: "MySQL Service to change port",
		}
		serviceID := addMySQLService(t, body)
		defer removeServices(t, serviceID)

		service, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: context.TODO(),
		})

		// Change MySQL service name.
		newPort := int64(3337)
		changeRes, err := client.Default.Services.ChangeMySQLService(&services.ChangeMySQLServiceParams{
			Body: services.ChangeMySQLServiceBody{
				ServiceID: serviceID,
				Port:      newPort,
			},
			Context: context.TODO(),
		})
		require.NoError(t, err)
		require.NotNil(t, changeRes.Payload.Mysql)
		require.Equal(t, serviceID, changeRes.Payload.Mysql.ServiceID)
		require.Equal(t, newPort, changeRes.Payload.Mysql.Port)

		// Check changes in backend.
		changedService, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: context.TODO(),
		})
		require.Equal(t, newPort, changedService.Payload.Mysql.Port)
		// Check that other fields isn't changed.
		require.Equal(t, service.Payload.Mysql.ServiceName, changedService.Payload.Mysql.ServiceName)
		require.Equal(t, service.Payload.Mysql.Address, changedService.Payload.Mysql.Address)
		require.Equal(t, service.Payload.Mysql.NodeID, changedService.Payload.Mysql.NodeID)
	})
}

func TestAmazonRDSMySQLService(t *testing.T) {
	t.Skip("Not implemented yet.")
	t.Run("Basic", func(t *testing.T) {
		params := &services.AddAmazonRDSMySQLServiceParams{
			Body: services.AddAmazonRDSMySQLServiceBody{
				NodeID:      "pmm-server",
				Address:     "localhost",
				Port:        3306,
				ServiceName: "Basic AmazonRDSMySQL Service",
			},
			Context: context.TODO(),
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
		require.Equal(t, "Basic AmazonRDSMySQL Service", res.Payload.AmazonRDSMysql.ServiceName)

		// Check if the service saved in PMM-Managed.
		serviceRes, err := client.Default.Services.GetService(&services.GetServiceParams{
			Body:    services.GetServiceBody{ServiceID: serviceID},
			Context: context.TODO(),
		})
		require.NoError(t, err)
		require.NotNil(t, serviceRes)
		require.NotNil(t, serviceRes.Payload.AmazonRDSMysql)
		require.Nil(t, serviceRes.Payload.Mysql)
		require.NotEmpty(t, serviceRes.Payload.AmazonRDSMysql.ServiceID)
		require.Equal(t, "pmm-server", serviceRes.Payload.AmazonRDSMysql.NodeID)
		require.Equal(t, "localhost", serviceRes.Payload.AmazonRDSMysql.Address)
		require.Equal(t, int64(3306), serviceRes.Payload.AmazonRDSMysql.Port)
		require.Equal(t, "Basic AmazonRDSMySQL Service", serviceRes.Payload.AmazonRDSMysql.ServiceName)

		// Check duplicates.
		params = &services.AddAmazonRDSMySQLServiceParams{
			Body: services.AddAmazonRDSMySQLServiceBody{
				NodeID:      "pmm-server",
				Address:     "127.0.0.1",
				Port:        3336,
				ServiceName: "Basic AmazonRDSMySQL Service",
			},
			Context: context.TODO(),
		}
		res, err = client.Default.Services.AddAmazonRDSMySQLService(params)
		require.Error(t, err) // Can't use EqualError because it returns different references each time.
		require.Contains(t, err.Error(), "unknown error (status 409)")
		require.Nil(t, res)
	})
}
