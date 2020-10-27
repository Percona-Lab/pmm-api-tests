package dbaas

import (
	"testing"

	dbaasClient "github.com/percona/pmm/api/managementpb/dbaas/json/client"
	psmdbcluster "github.com/percona/pmm/api/managementpb/dbaas/json/client/p_s_m_db_cluster"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

const (
	psmdbKubernetesClusterName = "api-test-k8s-mongodb-cluster"
)

//nolint:funlen
func TestPSMDBClusterServer(t *testing.T) {
	if pmmapitests.Kubeconfig == "" {
		t.Skip("Skip tests of PSMDBClusterServer without kubeconfig")
	}
	registerKubernetesCluster(t, psmdbKubernetesClusterName, pmmapitests.Kubeconfig)

	t.Run("BasicPSMDBCluster", func(t *testing.T) {
		paramsFirstPXC := psmdbcluster.CreatePSMDBClusterParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.CreatePSMDBClusterBody{
				KubernetesClusterName: psmdbKubernetesClusterName,
				Name:                  "first.pxc.test.percona.com",
				Params: &psmdbcluster.CreatePSMDBClusterParamsBodyParams{
					ClusterSize: 3,
					Replicaset: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicaset{
						ComputeResources: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
				},
			},
		}

		_, err := dbaasClient.Default.PsmDBCluster.CreatePSMDBCluster(&paramsFirstPXC)
		assert.NoError(t, err)

		// Create one more PSMDB Cluster.
		paramsSecondPXC := psmdbcluster.CreatePSMDBClusterParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.CreatePSMDBClusterBody{
				KubernetesClusterName: psmdbKubernetesClusterName,
				Name:                  "second.pxc.test.percona.com",
				Params: &psmdbcluster.CreatePSMDBClusterParamsBodyParams{
					ClusterSize: 1,
					Replicaset: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicaset{
						ComputeResources: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
				},
			},
		}
		_, err = dbaasClient.Default.PsmDBCluster.CreatePSMDBCluster(&paramsSecondPXC)
		assert.NoError(t, err)

		listPSMDBClustersParamsParam := psmdbcluster.ListPSMDBClustersParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.ListPSMDBClustersBody{
				KubernetesClusterName: psmdbKubernetesClusterName,
			},
		}
		xtraDBClusters, err := dbaasClient.Default.PsmDBCluster.ListPSMDBClusters(&listPSMDBClustersParamsParam)
		assert.NoError(t, err)

		for _, name := range []string{"first.pxc.test.percona.com", "second.pxc.test.percona.com"} {
			foundPXC := false
			for _, pxc := range xtraDBClusters.Payload.Clusters {
				if name == pxc.Name {
					foundPXC = true

					break
				}
			}
			assert.True(t, foundPXC, "Cannot find PXC with name %s in cluster list", name)
		}

		paramsUpdatePXC := psmdbcluster.UpdatePSMDBClusterParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.UpdatePSMDBClusterBody{
				KubernetesClusterName: psmdbKubernetesClusterName,
				Name:                  "second.pxc.test.percona.com",
				Params: &psmdbcluster.UpdatePSMDBClusterParamsBodyParams{
					ClusterSize: 2,
					Replicaset: &psmdbcluster.UpdatePSMDBClusterParamsBodyParamsReplicaset{
						ComputeResources: &psmdbcluster.UpdatePSMDBClusterParamsBodyParamsReplicasetComputeResources{
							CPUm:        2,
							MemoryBytes: "128",
						},
					},
				},
			},
		}

		_, err = dbaasClient.Default.PsmDBCluster.UpdatePSMDBCluster(&paramsUpdatePXC)
		pmmapitests.AssertAPIErrorf(t, err, 501, codes.Unimplemented, `This method is not implemented yet.`)

		for _, pxc := range xtraDBClusters.Payload.Clusters {
			if pxc.Name == "" {
				continue
			}
			deletePSMDBClusterParamsParam := psmdbcluster.DeletePSMDBClusterParams{
				Context: pmmapitests.Context,
				Body: psmdbcluster.DeletePSMDBClusterBody{
					KubernetesClusterName: psmdbKubernetesClusterName,
					Name:                  pxc.Name,
				},
			}
			_, err := dbaasClient.Default.PsmDBCluster.DeletePSMDBCluster(&deletePSMDBClusterParamsParam)
			assert.NoError(t, err)
		}
	})

	t.Run("CreatePSMDBClusterEmptyName", func(t *testing.T) {
		paramsPXCEmptyName := psmdbcluster.CreatePSMDBClusterParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.CreatePSMDBClusterBody{
				KubernetesClusterName: psmdbKubernetesClusterName,
				Name:                  "",
				Params: &psmdbcluster.CreatePSMDBClusterParamsBodyParams{
					ClusterSize: 3,
					Replicaset: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicaset{
						ComputeResources: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
				},
			},
		}
		_, err := dbaasClient.Default.PsmDBCluster.CreatePSMDBCluster(&paramsPXCEmptyName)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `invalid field Name: value '' must not be an empty string`)
	})

	t.Run("CreatePSMDBClusterInvalidName", func(t *testing.T) {
		paramsPXCInvalidName := psmdbcluster.CreatePSMDBClusterParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.CreatePSMDBClusterBody{
				KubernetesClusterName: psmdbKubernetesClusterName,
				Name:                  "123_asd",
				Params: &psmdbcluster.CreatePSMDBClusterParamsBodyParams{
					ClusterSize: 3,
					Replicaset: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicaset{
						ComputeResources: &psmdbcluster.CreatePSMDBClusterParamsBodyParamsReplicasetComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
				},
			},
		}
		_, err := dbaasClient.Default.PsmDBCluster.CreatePSMDBCluster(&paramsPXCInvalidName)
		assert.Error(t, err)
		assert.Equal(t, 500, err.(pmmapitests.ErrorResponse).Code())
	})

	t.Run("ListUnknownCluster", func(t *testing.T) {
		listPSMDBClustersParamsParam := psmdbcluster.ListPSMDBClustersParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.ListPSMDBClustersBody{
				KubernetesClusterName: "Unknown-kubernetes-cluster-name",
			},
		}
		_, err := dbaasClient.Default.PsmDBCluster.ListPSMDBClusters(&listPSMDBClustersParamsParam)
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, `Kubernetes Cluster with name "Unknown-kubernetes-cluster-name" not found.`)
	})

	t.Run("DeleteUnknownPSMDBCluster", func(t *testing.T) {
		deletePSMDBClusterParamsParam := psmdbcluster.DeletePSMDBClusterParams{
			Context: pmmapitests.Context,
			Body: psmdbcluster.DeletePSMDBClusterBody{
				KubernetesClusterName: psmdbKubernetesClusterName,
				Name:                  "Unknown-pxc-name",
			},
		}
		_, err := dbaasClient.Default.PsmDBCluster.DeletePSMDBCluster(&deletePSMDBClusterParamsParam)
		require.Error(t, err)
		assert.Equal(t, 500, err.(pmmapitests.ErrorResponse).Code())
	})
}
