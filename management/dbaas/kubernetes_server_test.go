package dbaas

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"

	dbaasClient "github.com/percona/pmm/api/managementpb/dbaas/json/client"
	"github.com/percona/pmm/api/managementpb/dbaas/json/client/kubernetes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestKubernetesServer(t *testing.T) {

	t.Run("Basic", func(t *testing.T) {
		kubernetesClusterName := pmmapitests.TestString(t, "api-test-cluster")
		clusters, err := dbaasClient.Default.Kubernetes.ListKubernetesClusters(nil)
		require.NoError(t, err)
		require.False(t, containsKubernetesCluster(clusters.Payload.KubernetesClusters, kubernetesClusterName))

		registerKubernetesCluster(t, kubernetesClusterName)
		defer unregisterKubernetesCluster(t, kubernetesClusterName)

		clusters, err = dbaasClient.Default.Kubernetes.ListKubernetesClusters(nil)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(clusters.Payload.KubernetesClusters), 1)
		require.True(t, containsKubernetesCluster(clusters.Payload.KubernetesClusters, kubernetesClusterName))

		unregisterKubernetesClusterResponse, err := dbaasClient.Default.Kubernetes.UnregisterKubernetesCluster(
			&kubernetes.UnregisterKubernetesClusterParams{
				Body: kubernetes.UnregisterKubernetesClusterBody{KubernetesClusterName: kubernetesClusterName},
			},
		)
		require.NoError(t, err)
		assert.NotNil(t, unregisterKubernetesClusterResponse)

		clusters, err = dbaasClient.Default.Kubernetes.ListKubernetesClusters(nil)
		assert.NoError(t, err)
		require.False(t, containsKubernetesCluster(clusters.Payload.KubernetesClusters, kubernetesClusterName))
	})

	t.Run("DuplicateClusterName", func(t *testing.T) {
		kubernetesClusterName := pmmapitests.TestString(t, "api-test-cluster-duplicate")
		registerKubernetesCluster(t, kubernetesClusterName)
		registerKubernetesClusterResponse, err := dbaasClient.Default.Kubernetes.RegisterKubernetesCluster(
			&kubernetes.RegisterKubernetesClusterParams{
				Body: kubernetes.RegisterKubernetesClusterBody{
					KubernetesClusterName: kubernetesClusterName,
					KubeAuth:              &kubernetes.RegisterKubernetesClusterParamsBodyKubeAuth{Kubeconfig: "{}"},
				},
			},
		)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, fmt.Sprintf("Cluster with Name %q already exists.", kubernetesClusterName))
		require.Nil(t, registerKubernetesClusterResponse)
	})
}
