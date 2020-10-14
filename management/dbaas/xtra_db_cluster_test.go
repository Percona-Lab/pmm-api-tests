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
			Context: pmmapitests.Context,
			Body: xtra_db_cluster.CreateXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "first-pxc-test",
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

		_, err := dbaasClient.Default.XtraDBCluster.CreateXtraDBCluster(&paramsFirstPXC)
		assert.NoError(t, err)

		// Create one more XtraDB Cluster.
		paramsSecondPXC := xtra_db_cluster.CreateXtraDBClusterParams{
			Context: pmmapitests.Context,
			Body: xtra_db_cluster.CreateXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "second-pxc-test",
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
		_, err = dbaasClient.Default.XtraDBCluster.CreateXtraDBCluster(&paramsSecondPXC)
		assert.NoError(t, err)

		listXtraDBClustersParamsParam := xtra_db_cluster.ListXtraDBClustersParams{
			Context: pmmapitests.Context,
			Body: xtra_db_cluster.ListXtraDBClustersBody{
				KubernetesClusterName: kubernetesClusterName,
			},
		}
		xtraDBClusters, err := dbaasClient.Default.XtraDBCluster.ListXtraDBClusters(&listXtraDBClustersParamsParam)
		assert.NoError(t, err)

		for _, name := range []string{"first-pxc-test", "second-pxc-test"} {
			foundPXC := false
			for _, pxc := range xtraDBClusters.Payload.Clusters {
				if name == pxc.Name {
					foundPXC = true

					break
				}
			}
			assert.True(t, foundPXC, "Cannot find PXC with name %s in cluster list", name)
		}

		showXtraDBClusterParamsParam := xtra_db_cluster.ShowXtraDBClusterParams{
			Context: pmmapitests.Context,
			Body: xtra_db_cluster.ShowXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "first-pxc-test",
			},
		}
		xtraDBCluster, err := dbaasClient.Default.XtraDBCluster.ShowXtraDBCluster(&showXtraDBClusterParamsParam)
		assert.NoError(t, err)
		assert.Equal(t, xtraDBCluster.Payload.Name, "first-pxc-test")
		assert.Equal(t, xtraDBCluster.Payload.Username, "root")
		assert.Equal(t, xtraDBCluster.Payload.Host, "first-pxc-test-proxysql")
		assert.Equal(t, xtraDBCluster.Payload.Port, 3306)

		paramsUpdatePXC := xtra_db_cluster.UpdateXtraDBClusterParams{
			Context: pmmapitests.Context,
			Body: xtra_db_cluster.UpdateXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "second-pxc-test",
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

		_, err = dbaasClient.Default.XtraDBCluster.UpdateXtraDBCluster(&paramsUpdatePXC)
		pmmapitests.AssertAPIErrorf(t, err, 501, codes.Unimplemented, `This method is not implemented yet.`)

		for _, pxc := range xtraDBClusters.Payload.Clusters {
			if pxc.Name == "" {
				continue
			}
			deleteXtraDBClusterParamsParam := xtra_db_cluster.DeleteXtraDBClusterParams{
				Context: pmmapitests.Context,
				Body: xtra_db_cluster.DeleteXtraDBClusterBody{
					KubernetesClusterName: kubernetesClusterName,
					Name:                  pxc.Name,
				},
			}
			_, err := dbaasClient.Default.XtraDBCluster.DeleteXtraDBCluster(&deleteXtraDBClusterParamsParam)
			assert.NoError(t, err)
		}
	})

	t.Run("CreateXtraDBClusterEmptyName", func(t *testing.T) {
		paramsPXCEmptyName := xtra_db_cluster.CreateXtraDBClusterParams{
			Context: pmmapitests.Context,
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
		_, err := dbaasClient.Default.XtraDBCluster.CreateXtraDBCluster(&paramsPXCEmptyName)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `invalid field Name: value '' must be a string conforming to regex "^[a-z]([-a-z0-9]*[a-z0-9])?$"`)
	})

	t.Run("CreateXtraDBClusterInvalidName", func(t *testing.T) {
		paramsPXCInvalidName := xtra_db_cluster.CreateXtraDBClusterParams{
			Context: pmmapitests.Context,
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
		_, err := dbaasClient.Default.XtraDBCluster.CreateXtraDBCluster(&paramsPXCInvalidName)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, `invalid field Name: value '123_asd' must be a string conforming to regex "^[a-z]([-a-z0-9]*[a-z0-9])?$"`)
	})

	t.Run("ListUnknownCluster", func(t *testing.T) {
		listXtraDBClustersParamsParam := xtra_db_cluster.ListXtraDBClustersParams{
			Context: pmmapitests.Context,
			Body: xtra_db_cluster.ListXtraDBClustersBody{
				KubernetesClusterName: "Unknown-kubernetes-cluster-name",
			},
		}
		_, err := dbaasClient.Default.XtraDBCluster.ListXtraDBClusters(&listXtraDBClustersParamsParam)
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, `Kubernetes Cluster with name "Unknown-kubernetes-cluster-name" not found.`)
	})

	t.Run("DeleteUnknownXtraDBCluster", func(t *testing.T) {
		deleteXtraDBClusterParamsParam := xtra_db_cluster.DeleteXtraDBClusterParams{
			Context: pmmapitests.Context,
			Body: xtra_db_cluster.DeleteXtraDBClusterBody{
				KubernetesClusterName: kubernetesClusterName,
				Name:                  "Unknown-pxc-name",
			},
		}
		_, err := dbaasClient.Default.XtraDBCluster.DeleteXtraDBCluster(&deleteXtraDBClusterParamsParam)
		require.Error(t, err)
		assert.Equal(t, 500, err.(pmmapitests.ErrorResponse).Code())
	})
}
