package dbaas

import (
	"testing"

	dbaasClient "github.com/percona/pmm/api/managementpb/dbaas/json/client"
	"github.com/percona/pmm/api/managementpb/dbaas/json/client/kubernetes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func registerKubernetesCluster(t *testing.T, kubernetesClusterName string) {
	registerKubernetesClusterResponse, err := dbaasClient.Default.Kubernetes.RegisterKubernetesCluster(
		&kubernetes.RegisterKubernetesClusterParams{
			Body: kubernetes.RegisterKubernetesClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				KubeAuth:              &kubernetes.RegisterKubernetesClusterParamsBodyKubeAuth{Kubeconfig: "{}"},
			},
			Context: pmmapitests.Context,
		},
	)
	require.NoError(t, err)
	assert.NotNil(t, registerKubernetesClusterResponse)
}

func unregisterKubernetesCluster(kubernetesClusterName string) {
	_, _ = dbaasClient.Default.Kubernetes.UnregisterKubernetesCluster(
		&kubernetes.UnregisterKubernetesClusterParams{
			Body:    kubernetes.UnregisterKubernetesClusterBody{KubernetesClusterName: kubernetesClusterName},
			Context: pmmapitests.Context,
		},
	)
}

func containsKubernetesCluster(clusters []*kubernetes.KubernetesClustersItems0, name string) bool {
	for _, cluster := range clusters {
		if cluster.KubernetesClusterName == name {
			return true
		}
	}
	return false
}
