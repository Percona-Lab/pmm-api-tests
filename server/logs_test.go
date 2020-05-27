package server

import (
	"archive/zip"
	"bytes"
	"sort"
	"testing"

	serverClient "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestDownloadLogs(t *testing.T) {
	var buf bytes.Buffer
	res, err := serverClient.Default.Server.Logs(&server.LogsParams{
		Context: pmmapitests.Context,
	}, &buf)
	require.NoError(t, err)
	require.NotNil(t, res)

	r := bytes.NewReader(buf.Bytes())
	zipR, err := zip.NewReader(r, r.Size())
	assert.NoError(t, err)

	expected := []string{
		"alertmanager.log",
		"clickhouse-server.err.log",
		"clickhouse-server.log",
		"clickhouse-server.startup.log",
		"client/list.txt",
		"client/pmm-admin-version.txt",
		"client/pmm-agent-config.yaml",
		"client/pmm-agent-version.txt",
		"client/status.json",
		"cron.log",
		"dashboard-upgrade.log",
		"grafana.log",
		"installed.json",
		"nginx.access.log",
		"nginx.conf",
		"nginx.error.log",
		"nginx.startup.log",
		"pmm-agent.log",
		"pmm-agent.yaml",
		"pmm-managed.log",
		"pmm-ssl.conf",
		"pmm-version.txt",
		"pmm.conf",
		"pmm.ini",
		"postgresql.log",
		"postgresql.startup.log",
		"prometheus.ini",
		"prometheus.log",
		"prometheus.yml",
		"prometheus_targets.json",
		"qan-api2.ini",
		"qan-api2.log",
		"supervisorctl_status.log",
		"supervisord.conf",
		"supervisord.log",
		"systemctl_status.log",
	}

	actual := make([]string, len(zipR.File))
	for i, file := range zipR.File {
		actual[i] = file.Name
	}

	sort.Strings(actual)
	assert.Equal(t, expected, actual)
}
