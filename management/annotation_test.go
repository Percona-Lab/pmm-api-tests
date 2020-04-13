package management

import (
	"testing"

	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/annotation"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAddAnnotation(t *testing.T) {
	t.Run("Add Basic Annotation", func(t *testing.T) {
		params := &annotation.AddAnnotationParams{
			Body: annotation.AddAnnotationBody{
				Text: "Annotation Text",
				Tags: []string{"tag1", "tag2"},
			},
			Context: pmmapitests.Context,
		}
		_, err := client.Default.Annotation.AddAnnotation(params)
		require.NoError(t, err)
	})

	t.Run("Add Empty Annotation", func(t *testing.T) {
		params := &annotation.AddAnnotationParams{
			Body: annotation.AddAnnotationBody{
				Text: "",
				Tags: []string{},
			},
			Context: pmmapitests.Context,
		}
		_, err := client.Default.Annotation.AddAnnotation(params)
		require.Error(t, err)
	})
}
