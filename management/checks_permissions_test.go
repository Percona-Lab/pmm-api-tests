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

const pmmAddr = "localhost:443"

func createUserWithRole(login, role string) error {
	userID, err := createUser(login)
	if err != nil {
		return err
	}

	if err = setRole(userID, role); err != nil {
		return err
	}

	return nil
}

func createUser(login string) (int, error) {
	// https://grafana.com/docs/http_api/admin/#global-users
	data, err := json.Marshal(map[string]string{
		"name":     login,
		"email":    login + "@percona.invalid",
		"login":    login,
		"password": login,
	})
	if err != nil {
		return 0, err
	}

	u := url.URL{
		Scheme: "https",
		Host:   pmmAddr,
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

func setRole(userID int, role string) error {
	// https://grafana.com/docs/http_api/org/#updates-the-given-user
	data, err := json.Marshal(map[string]string{
		"role": role,
	})
	if err != nil {
		return err
	}

	u := url.URL{
		Scheme: "https",
		Host:   pmmAddr,
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
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	none := "none-" + ts
	viewer := "viewer-" + ts
	editor := "editor-" + ts
	admin := "admin-" + ts

	_, err := createUser(none)
	require.NoError(t, err)

	err = createUserWithRole(viewer, "Viewer")
	require.NoError(t, err)

	err = createUserWithRole(editor, "Editor")
	require.NoError(t, err)

	err = createUserWithRole(admin, "Admin")
	require.NoError(t, err)

	tests := []struct {
		name       string
		login      string
		statusCode int
	}{
		{name: "default", login: none, statusCode: http.StatusUnauthorized},
		{name: "viewer", login: viewer, statusCode: http.StatusUnauthorized},
		{name: "editor", login: editor, statusCode: http.StatusUnauthorized},
		{name: "admin", login: admin, statusCode: http.StatusOK},
	}

	for _, test := range tests {
		test := test
		t.Run("get settings/"+test.name, func(t *testing.T) {
			// make a BaseURL without authentication
			u, err := url.Parse(pmmapitests.BaseURL.String())
			require.NoError(t, err)
			u.User = url.UserPassword(test.login, test.login)
			u.Path = "/v1/Settings/Get"

			resp, err := http.Post(u.String(), "", nil)
			require.NoError(t, err)
			defer resp.Body.Close() //nolint:errcheck

			assert.Equal(t, test.statusCode, resp.StatusCode)
		})

		t.Run("get alerts/"+test.name, func(t *testing.T) {
			// make a BaseURL without authentication
			u, err := url.Parse(pmmapitests.BaseURL.String())
			require.NoError(t, err)
			u.User = url.UserPassword(test.login, test.login)
			u.Path = "/alertmanager/api/v2/alerts"

			resp, err := http.Get(u.String())
			require.NoError(t, err)
			defer resp.Body.Close() //nolint:errcheck

			assert.Equal(t, test.statusCode, resp.StatusCode)
		})
	}
}
