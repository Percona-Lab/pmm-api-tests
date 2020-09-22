package dbaas

import (
	"testing"

	dbaasClient "github.com/percona/pmm/api/managementpb/dbaas/json/client"
	"github.com/percona/pmm/api/managementpb/dbaas/json/client/xtra_db_cluster"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

const (
	kubernetesClusterName = "api-test-k8s-cluster"
)

//nolint:funlen
func TestXtraDBClusterServer(t *testing.T) {
	if pmmapitests.Kubeconfig == "" {
		t.Skip("Skip tests of XtraDBClusterServer without kubeconfig")
	}
	registerKubernetesCluster(t, kubernetesClusterName, pmmapitests.Kubeconfig)

	t.Run("BasicXtraDBCluster", func(t *testing.T) {
		paramsFirstPXC := xtra_db_cluster.CreateXtraDBClusterParams{
			Body: xtra_db_cluster.CreateXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "first.pxc.test.percona.com",
				Params: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParams{
					ClusterSize: 3,
					Proxysql: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsProxysql{
						ComputeResources: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
					Pxc: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsPxc{
						ComputeResources: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsPxcComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
				},
			},
		}

		paramsFirstPXC.WithContext(pmmapitests.Context)

		_, err := dbaasClient.Default.XtraDBCluster.CreateXtraDBCluster(&paramsFirstPXC)
		assert.NoError(t, err)

		// Create one more XtraDB Cluster.
		paramsSecondPXC := xtra_db_cluster.CreateXtraDBClusterParams{
			Body: xtra_db_cluster.CreateXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "second.pxc.test.percona.com",
				Params: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParams{
					ClusterSize: 1,
					Proxysql: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsProxysql{
						ComputeResources: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
					Pxc: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsPxc{
						ComputeResources: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsPxcComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
				},
			},
		}
		paramsSecondPXC.WithContext(pmmapitests.Context)
		_, err = dbaasClient.Default.XtraDBCluster.CreateXtraDBCluster(&paramsSecondPXC)
		assert.NoError(t, err)

		listXtraDBClustersParamsParam := xtra_db_cluster.ListXtraDBClustersParams{
			Body: xtra_db_cluster.ListXtraDBClustersBody{
				KubernetesClusterName: kubernetesClusterName,
			},
		}
		listXtraDBClustersParamsParam.WithContext(pmmapitests.Context)
		xtraDBClusters, err := dbaasClient.Default.XtraDBCluster.ListXtraDBClusters(&listXtraDBClustersParamsParam)
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

		paramsUpdatePXC := xtra_db_cluster.UpdateXtraDBClusterParams{
			Body: xtra_db_cluster.UpdateXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "second.pxc.test.percona.com",
				Params: &xtra_db_cluster.UpdateXtraDBClusterParamsBodyParams{
					ClusterSize: 2,
					Proxysql: &xtra_db_cluster.UpdateXtraDBClusterParamsBodyParamsProxysql{
						ComputeResources: &xtra_db_cluster.UpdateXtraDBClusterParamsBodyParamsProxysqlComputeResources{
							CPUm:        2,
							MemoryBytes: "128",
						},
					},
					Pxc: &xtra_db_cluster.UpdateXtraDBClusterParamsBodyParamsPxc{
						ComputeResources: &xtra_db_cluster.UpdateXtraDBClusterParamsBodyParamsPxcComputeResources{
							CPUm:        2,
							MemoryBytes: "128",
						},
					},
				},
			},
		}

		paramsUpdatePXC.WithContext(pmmapitests.Context)
		_, err = dbaasClient.Default.XtraDBCluster.UpdateXtraDBCluster(&paramsUpdatePXC)
		pmmapitests.AssertAPIErrorf(t, err, 501, codes.Unimplemented, `This method is not implemented yet.`)

		for _, pxc := range xtraDBClusters.Payload.Clusters {
			if pxc.Name == "" {
				continue
			}
			deleteXtraDBClusterParamsParam := xtra_db_cluster.DeleteXtraDBClusterParams{
				Body: xtra_db_cluster.DeleteXtraDBClusterBody{
					KubernetesClusterName: kubernetesClusterName,
					Name:                  pxc.Name,
				},
			}
			deleteXtraDBClusterParamsParam.WithContext(pmmapitests.Context)
			_, err := dbaasClient.Default.XtraDBCluster.DeleteXtraDBCluster(&deleteXtraDBClusterParamsParam)
			assert.NoError(t, err)
		}
	})

	t.Run("Create_XtraDBCluster_Invalid_Name", func(t *testing.T) {
		paramsPXCEmptyName := xtra_db_cluster.CreateXtraDBClusterParams{
			Body: xtra_db_cluster.CreateXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "",
				Params: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParams{
					ClusterSize: 1,
					Proxysql: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsProxysql{
						ComputeResources: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
					Pxc: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsPxc{
						ComputeResources: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsPxcComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
				},
			},
		}
		paramsPXCEmptyName.WithContext(pmmapitests.Context)
		_, err := dbaasClient.Default.XtraDBCluster.CreateXtraDBCluster(&paramsPXCEmptyName)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `invalid field Name: value '' must not be an empty string`)

		paramsPXCInvalidName := xtra_db_cluster.CreateXtraDBClusterParams{
			Body: xtra_db_cluster.CreateXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "123_asd",
				Params: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParams{
					ClusterSize: 1,
					Proxysql: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsProxysql{
						ComputeResources: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
					Pxc: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsPxc{
						ComputeResources: &xtra_db_cluster.CreateXtraDBClusterParamsBodyParamsPxcComputeResources{
							CPUm:        1,
							MemoryBytes: "64",
						},
					},
				},
			},
		}
		paramsPXCInvalidName.WithContext(pmmapitests.Context)
		_, err = dbaasClient.Default.XtraDBCluster.CreateXtraDBCluster(&paramsPXCInvalidName)
		assert.Error(t, err)
	})

	t.Run("List_Unknown_Cluster", func(t *testing.T) {
		listXtraDBClustersParamsParam := xtra_db_cluster.ListXtraDBClustersParams{
			Body: xtra_db_cluster.ListXtraDBClustersBody{
				KubernetesClusterName: "Unknown-kubernetes-cluster-name",
			},
		}
		listXtraDBClustersParamsParam.WithContext(pmmapitests.Context)
		_, err := dbaasClient.Default.XtraDBCluster.ListXtraDBClusters(&listXtraDBClustersParamsParam)
		require.Error(t, err)
	})

	t.Run("Delete_Unknown_XtraDBCluster", func(t *testing.T) {
		deleteXtraDBClusterParamsParam := xtra_db_cluster.DeleteXtraDBClusterParams{
			Body: xtra_db_cluster.DeleteXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "Unknown-pxc-name",
			},
		}
		deleteXtraDBClusterParamsParam.WithContext(pmmapitests.Context)
		_, err := dbaasClient.Default.XtraDBCluster.DeleteXtraDBCluster(&deleteXtraDBClusterParamsParam)
		require.Error(t, err)
	})
}
