package server

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloadLogs(t *testing.T) {
	url := "http://localhost:7772/logs.zip"

	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	req.Header.Set("Accept", "application/zip")

	client := &http.Client{}

	resp, err := client.Do(req)
	assert.NoError(t, err)

	defer resp.Body.Close() //nolint:errcheck

	zipfile, err := ioutil.TempFile("", "*-test.zip")
	assert.NoError(t, err)

	defer zipfile.Close() //nolint:errcheck

	_, err = io.Copy(zipfile, resp.Body)
	require.NoError(t, err)

	reader, err := zip.OpenReader(zipfile.Name())
	assert.NoError(t, err)

	hasClientDir := false

	for _, file := range reader.File {
		if filepath.Dir(file.Name) == "client" {
			hasClientDir = true
			break
		}
	}

	assert.True(t, hasClientDir)
}
