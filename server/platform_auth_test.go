package server

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

// Tests in this file cover Percona Platform authentication.

func TestPlatformAuth(t *testing.T) {
	client := serverClient.Default.Server
	login := gofakeit.Email()
	password := gofakeit.Password(true, true, true, false, false, 14)

	_, err := client.PlatformSignUp(&server.PlatformSignUpParams{
		Body: server.PlatformSignUpBody{
			Email:    login,
			Password: password,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)

	_, err = client.PlatformSignIn(&server.PlatformSignInParams{
		Body: server.PlatformSignInBody{
			Email:    login,
			Password: password,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)
}
