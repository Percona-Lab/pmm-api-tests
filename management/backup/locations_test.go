package backup

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	backupClient "github.com/percona/pmm/api/managementpb/backup/json/client"
	"github.com/percona/pmm/api/managementpb/backup/json/client/locations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAddLocation(t *testing.T) {
	client := backupClient.Default.Locations

	t.Run("normal fs config", func(t *testing.T) {
		resp, err := client.AddLocation(&locations.AddLocationParams{
			Body: locations.AddLocationBody{
				Name:        gofakeit.Name(),
				Description: gofakeit.Question(),
				FsConfig: &locations.AddLocationParamsBodyFsConfig{
					Path: "/tmp",
				},
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		defer deleteLocation(t, client, resp.Payload.LocationID)

		assert.NotEmpty(t, resp.Payload.LocationID)
	})

}

func deleteLocation(t *testing.T, client locations.ClientService, id string) {
	// @TODO
}
