package management

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

const grafanaHost = "localhost"

func createUserWithRole(t *testing.T, login, role string) error {
	userID, err := createUser(t, login)
	if err != nil {
		return err
	}

	if err = setRole(t, userID, role); err != nil {
		return err
	}

	return nil
}

func createUser(t *testing.T, login string) (int, error) {
	// https://grafana.com/docs/http_api/admin/#global-users
	data, err := json.Marshal(map[string]string{
		"name":     login,
		"email":    login + "@percona.invalid",
		"login":    login,
		"password": login,
	})

	u := url.URL{
		Scheme: "http",
		Host:   grafanaHost,
		Path:   "/graph/api/admin/users",
		User:   url.UserPassword("admin", "admin"),
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(data))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close() //nolint:errcheck

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to create user, status code: %d, response: %s", resp.StatusCode, b)
	}

	var m map[string]interface{}
	if resp.Body != nil {
		if err = json.Unmarshal(b, &m); err != nil {
			return 0, err
		}
	}

	return int(m["id"].(float64)), nil
}

func setRole(t *testing.T, userID int, role string) error {
	// https://grafana.com/docs/http_api/org/#updates-the-given-user
	data, err := json.Marshal(map[string]string{
		"role": role,
	})

	u := url.URL{
		Scheme: "http",
		Host:   grafanaHost,
		Path:   "/graph/api/org/users/" + strconv.Itoa(userID),
		User:   url.UserPassword("admin", "admin"),
	}

	req, err := http.NewRequest(http.MethodPatch, u.String(), bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to set role for user, status code: %d, response: %s", resp.StatusCode, b)
	}

	return nil
}

func TestPermissions(t *testing.T) {
	viewer := "viewer" + strconv.FormatInt(time.Now().Unix(), 10)
	admin := "admin" + strconv.FormatInt(time.Now().Unix(), 10)

	err := createUserWithRole(t, viewer, "Viewer")
	require.NoError(t, err)

	err = createUserWithRole(t, admin, "Admin")
	require.NoError(t, err)

	tests := []struct {
		login      string
		statusCode int
	}{
		{login: viewer, statusCode: http.StatusUnauthorized},
		{login: admin, statusCode: http.StatusOK},
	}

	for _, test := range tests {
		test := test
		t.Run("get settings", func(t *testing.T) {
			// make a BaseURL without authentication
			u, err := url.Parse(pmmapitests.BaseURL.String())
			require.NoError(t, err)
			u.User = url.UserPassword(test.login, test.login)
			u.Path = "/v1/Settings/Get"

			resp, err := http.Post(u.String(), "", nil)
			require.NoError(t, err)
			assert.Equal(t, test.statusCode, resp.StatusCode)
			// b,err := ioutil.ReadAll(resp.Body)
			// require.NoError(t, err)
			// t.Log(string(b))
		})

		t.Run("get alerts", func(t *testing.T) {
			// make a BaseURL without authentication
			u, err := url.Parse(pmmapitests.BaseURL.String())
			require.NoError(t, err)
			u.User = url.UserPassword(test.login, test.login)
			u.Path = "/v1/Settings/Get"

			resp, err := http.Post(u.String(), "", nil)
			require.NoError(t, err)
			assert.Equal(t, test.statusCode, resp.StatusCode)
			// b,err := ioutil.ReadAll(resp.Body)
			// require.NoError(t, err)
			// t.Log(string(b))
		})
	}

}
