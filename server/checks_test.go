package server

import (
	"testing"

	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

)

func TestStartChecks(t *testing.T) {

	client := serverClient.Default.Server

	resp, err := client.StartSecurityChecks(nil)
	require.NoError(t, err)
	assert.NotNil(t, resp)
}
