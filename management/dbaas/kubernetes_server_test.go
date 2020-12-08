package dbaas

import (
	"fmt"
	"os"
	"testing"
	"time"

	dbaasClient "github.com/percona/pmm/api/managementpb/dbaas/json/client"
	"github.com/percona/pmm/api/managementpb/dbaas/json/client/kubernetes"
	psmdbcluster "github.com/percona/pmm/api/managementpb/dbaas/json/client/psmdb_cluster"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestKubernetesServer(t *testing.T) {
	if os.Getenv("PERCONA_TEST_DBAAS") != "1" {
		t.Skip("PERCONA_TEST_DBAAS env variable is not passed, skipping")
	}
	kubeConfig := os.Getenv("PERCONA_TEST_DBAAS_KUBECONFIG")
	if kubeConfig == "" {
		t.Skip("PERCONA_TEST_DBAAS_KUBECONFIG env variable is not provided")
	}
	t.Run("Basic", func(t *testing.T) {
		kubernetesClusterName := pmmapitests.TestString(t, "api-test-cluster")
		clusters, err := dbaasClient.Default.Kubernetes.ListKubernetesClusters(nil)
		require.NoError(t, err)
		require.NotContains(t, clusters.Payload.KubernetesClusters, &kubernetes.KubernetesClustersItems0{KubernetesClusterName: kubernetesClusterName})

		registerKubernetesCluster(t, kubernetesClusterName, kubeConfig)
		clusters, err = dbaasClient.Default.Kubernetes.ListKubernetesClusters(nil)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(clusters.Payload.KubernetesClusters), 1)
		assert.Contains(t, clusters.Payload.KubernetesClusters, &kubernetes.KubernetesClustersItems0{KubernetesClusterName: kubernetesClusterName})

		unregisterKubernetesClusterResponse, err := dbaasClient.Default.Kubernetes.UnregisterKubernetesCluster(
			&kubernetes.UnregisterKubernetesClusterParams{
				Body:    kubernetes.UnregisterKubernetesClusterBody{KubernetesClusterName: kubernetesClusterName},
				Context: pmmapitests.Context,
			},
		)
		require.NoError(t, err)
		assert.NotNil(t, unregisterKubernetesClusterResponse)

		clusters, err = dbaasClient.Default.Kubernetes.ListKubernetesClusters(nil)
		assert.NoError(t, err)
		require.NotContains(t, clusters.Payload.KubernetesClusters, &kubernetes.KubernetesClustersItems0{KubernetesClusterName: kubernetesClusterName})
	})

	t.Run("DuplicateClusterName", func(t *testing.T) {
		kubernetesClusterName := pmmapitests.TestString(t, "api-test-cluster-duplicate")

		registerKubernetesCluster(t, kubernetesClusterName, kubeConfig)
		registerKubernetesClusterResponse, err := dbaasClient.Default.Kubernetes.RegisterKubernetesCluster(
			&kubernetes.RegisterKubernetesClusterParams{
				Body: kubernetes.RegisterKubernetesClusterBody{
					KubernetesClusterName: kubernetesClusterName,
					KubeAuth:              &kubernetes.RegisterKubernetesClusterParamsBodyKubeAuth{Kubeconfig: kubeConfig},
				},
				Context: pmmapitests.Context,
			},
		)
		pmmapitests.AssertAPIErrorf(t, err, 409, codes.AlreadyExists, fmt.Sprintf("Kubernetes Cluster with Name %q already exists.", kubernetesClusterName))
		require.Nil(t, registerKubernetesClusterResponse)
	})

	t.Run("EmptyKubernetesClusterName", func(t *testing.T) {
		registerKubernetesClusterResponse, err := dbaasClient.Default.Kubernetes.RegisterKubernetesCluster(
			&kubernetes.RegisterKubernetesClusterParams{
				Body: kubernetes.RegisterKubernetesClusterBody{
					KubernetesClusterName: "",
					KubeAuth:              &kubernetes.RegisterKubernetesClusterParamsBodyKubeAuth{Kubeconfig: kubeConfig},
				},
				Context: pmmapitests.Context,
			},
		)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field KubernetesClusterName: value '' must not be an empty string")
		require.Nil(t, registerKubernetesClusterResponse)
	})

	t.Run("EmptyKubeConfig", func(t *testing.T) {
		registerKubernetesClusterResponse, err := dbaasClient.Default.Kubernetes.RegisterKubernetesCluster(
			&kubernetes.RegisterKubernetesClusterParams{
				Body: kubernetes.RegisterKubernetesClusterBody{
					KubernetesClusterName: "empty-kube-config",
					KubeAuth:              &kubernetes.RegisterKubernetesClusterParamsBodyKubeAuth{},
				},
				Context: pmmapitests.Context,
			},
		)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field KubeAuth.Kubeconfig: value '' must not be an empty string")
		require.Nil(t, registerKubernetesClusterResponse)
	})

	t.Run("UnregisterNotExistCluster", func(t *testing.T) {
		unregisterKubernetesClusterOK, err := unregisterKubernetesCluster("not-exist-cluster")
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, "Kubernetes Cluster with name \"not-exist-cluster\" not found.")
		require.Nil(t, unregisterKubernetesClusterOK)
	})

	t.Run("UnregisterEmptyClusterName", func(t *testing.T) {
		unregisterKubernetesClusterOK, err := unregisterKubernetesCluster("")
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "invalid field KubernetesClusterName: value '' must not be an empty string")
		require.Nil(t, unregisterKubernetesClusterOK)
	})

	t.Run("UnregisterWithoutAndWithForce", func(t *testing.T) {
		kubernetesClusterName := pmmapitests.TestString(t, "api-test-cluster")
		dbClusterName := "first-psmdb-test"
		clusters, err := dbaasClient.Default.Kubernetes.ListKubernetesClusters(nil)
		require.NoError(t, err)
		require.NotContains(t, clusters.Payload.KubernetesClusters, &kubernetes.KubernetesClustersItems0{KubernetesClusterName: kubernetesClusterName})
		registerKubernetesCluster(t, kubernetesClusterName, kubeConfig)

		paramsFirstPSMDB := psmdbcluster.CreatePSMDBClusterParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.CreatePSMDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  dbClusterName,
				Params: &psmdbcluster.CreatePSMDBClusterParamsBodyParams{
					ClusterSize: 3,
					Replicaset: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicaset{
						ComputeResources: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources{
							CPUm:        500,
							MemoryBytes: "1000000000",
						},
						DiskSize: "1000000000",
					},
				},
			},
		}
		_, err = dbaasClient.Default.PSMDBCluster.CreatePSMDBCluster(&paramsFirstPSMDB)
		assert.NoError(t, err)

		clusters, err = dbaasClient.Default.Kubernetes.ListKubernetesClusters(nil)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(clusters.Payload.KubernetesClusters), 1)
		assert.Contains(t, clusters.Payload.KubernetesClusters, &kubernetes.KubernetesClustersItems0{KubernetesClusterName: kubernetesClusterName})

		unregisterKubernetesClusterResponse, err := dbaasClient.Default.Kubernetes.UnregisterKubernetesCluster(
			&kubernetes.UnregisterKubernetesClusterParams{
				Body: kubernetes.UnregisterKubernetesClusterBody{
					KubernetesClusterName: kubernetesClusterName,
				},
				Context: pmmapitests.Context,
			},
		)
		require.Error(t, err)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, fmt.Sprintf(`Kubernetes cluster %s has PSMDB clusters`, kubernetesClusterName))

		unregisterKubernetesClusterResponse, err = dbaasClient.Default.Kubernetes.UnregisterKubernetesCluster(
			&kubernetes.UnregisterKubernetesClusterParams{
				Body: kubernetes.UnregisterKubernetesClusterBody{
					KubernetesClusterName: kubernetesClusterName,
					Force:                 true,
				},
				Context: pmmapitests.Context,
			},
		)
		require.NoError(t, err)
		assert.NotNil(t, unregisterKubernetesClusterResponse)

		unregisterKubernetesClusterResponse, err = dbaasClient.Default.Kubernetes.UnregisterKubernetesCluster(
			&kubernetes.UnregisterKubernetesClusterParams{
				Body: kubernetes.UnregisterKubernetesClusterBody{
					KubernetesClusterName: kubernetesClusterName,
				},
				Context: pmmapitests.Context,
			},
		)
		require.Error(t, err)
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, fmt.Sprintf(`Kubernetes Cluster with name "%s" not found.`, kubernetesClusterName))

		registerKubernetesCluster(t, kubernetesClusterName, kubeConfig)
		deletePSMDBClusterParamsParam := psmdbcluster.DeletePSMDBClusterParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.DeletePSMDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  dbClusterName,
			},
		}
		_, err = dbaasClient.Default.PSMDBCluster.DeletePSMDBCluster(&deletePSMDBClusterParamsParam)
		assert.NoError(t, err)

		time.Sleep(10 * time.Second)

		unregisterKubernetesClusterResponse, err = dbaasClient.Default.Kubernetes.UnregisterKubernetesCluster(
			&kubernetes.UnregisterKubernetesClusterParams{
				Body: kubernetes.UnregisterKubernetesClusterBody{
					KubernetesClusterName: kubernetesClusterName,
				},
				Context: pmmapitests.Context,
			},
		)
		assert.NoError(t, err)
	})
}
