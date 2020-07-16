package server

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestPerconaAuth(t *testing.T) {
	client := serverClient.Default.Server
	login := gofakeit.Email()
	password := "Password12345"

	_, err := client.SignUp(&server.SignUpParams{
		Body: server.SignUpBody{
			Email:    login,
			Password: password,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)

	_, err = client.SignIn(&server.SignInParams{
		Body: server.SignInBody{
			Email:    login,
			Password: password,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)
}

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}
