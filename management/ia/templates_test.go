package ia

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/brianvoe/gofakeit"
	templatesClient "github.com/percona/pmm/api/managementpb/ia/json/client"
	"github.com/percona/pmm/api/managementpb/ia/json/client/templates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestAddTemplate(t *testing.T) {
	client := templatesClient.Default.Templates

	b, err := ioutil.ReadFile("../../testdata/ia/template.yaml")
	require.NoError(t, err)

	t.Run("normal", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID()),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		resp, err := client.ListTemplates(&templates.ListTemplatesParams{
			Body: templates.ListTemplatesBody{
				Reload: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		var found bool
		for _, template := range resp.Payload.Templates {
			if template.Name == name {
				found = true
			}
		}
		assert.Truef(t, found, "Template with id %s not found", name)
	})

	t.Run("duplicate", func(t *testing.T) {
		name := gofakeit.UUID()
		yaml := fmt.Sprintf(string(b), name, gofakeit.UUID())
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: yaml,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		_, err = client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: yaml,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 409, codes.AlreadyExists, fmt.Sprintf("Template with name \"%s\" already exists.", name))
	})

	t.Run("invalid yaml", func(t *testing.T) {
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: "not a yaml",
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "Failed to parse rule template.")
	})

	t.Run("invalid template", func(t *testing.T) {
		b, err := ioutil.ReadFile("../../testdata/ia/invalid-template.yaml")
		require.NoError(t, err)
		name := gofakeit.UUID()
		_, err = client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID()),
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "Failed to parse rule template.")
	})
}

func TestChangeTemplate(t *testing.T) {
	client := templatesClient.Default.Templates

	b, err := ioutil.ReadFile("../../testdata/ia/template.yaml")
	require.NoError(t, err)

	t.Run("normal", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID()),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		newExpr := gofakeit.UUID()
		_, err = client.UpdateTemplate(&templates.UpdateTemplateParams{
			Body: templates.UpdateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, newExpr),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		resp, err := client.ListTemplates(&templates.ListTemplatesParams{
			Body: templates.ListTemplatesBody{
				Reload: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		var found bool
		for _, template := range resp.Payload.Templates {
			if template.Name == name {
				assert.Equal(t, newExpr, template.Expr)
				found = true
			}
		}
		assert.Truef(t, found, "Template with id %s not found", name)
	})

	t.Run("unknown template", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err = client.UpdateTemplate(&templates.UpdateTemplateParams{
			Body: templates.UpdateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID()),
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, fmt.Sprintf("Template with name \"%s\" not found.", name))
	})

	t.Run("invalid yaml", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID()),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		_, err = client.UpdateTemplate(&templates.UpdateTemplateParams{
			Body: templates.UpdateTemplateBody{
				Yaml: "not a yaml",
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "Failed to parse rule template.")
	})

	t.Run("invalid template", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err = client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID()),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		b, err = ioutil.ReadFile("../../testdata/ia/invalid-template.yaml")
		_, err = client.UpdateTemplate(&templates.UpdateTemplateParams{
			Body: templates.UpdateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID()),
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.InvalidArgument, "Failed to parse rule template.")
	})
}

func TestDeleteTemplate(t *testing.T) {
	client := templatesClient.Default.Templates

	b, err := ioutil.ReadFile("../../testdata/ia/template.yaml")
	require.NoError(t, err)

	t.Run("normal", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID()),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		_, err = client.DeleteTemplate(&templates.DeleteTemplateParams{
			Body: templates.DeleteTemplateBody{
				Name: name,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		resp, err := client.ListTemplates(&templates.ListTemplatesParams{
			Body: templates.ListTemplatesBody{
				Reload: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		for _, template := range resp.Payload.Templates {
			assert.NotEqual(t, name, template.Name)
		}
	})

	t.Run("unknown template", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err = client.DeleteTemplate(&templates.DeleteTemplateParams{
			Body: templates.DeleteTemplateBody{
				Name: name,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, fmt.Sprintf("Template with name \"%s\" not found.", name))
	})
}

func TestListTemplate(t *testing.T) {
	client := templatesClient.Default.Templates

	b, err := ioutil.ReadFile("../../testdata/ia/template.yaml")
	require.NoError(t, err)

	name := gofakeit.UUID()
	expr := gofakeit.UUID()
	file := fmt.Sprintf(string(b), name, expr)
	_, err = client.CreateTemplate(&templates.CreateTemplateParams{
		Body: templates.CreateTemplateBody{
			Yaml: file,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)
	resp, err := client.ListTemplates(&templates.ListTemplatesParams{
		Body: templates.ListTemplatesBody{
			Reload: true,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)

	var found bool
	for _, template := range resp.Payload.Templates {
		if template.Name == name {
			assert.Equal(t, expr, template.Expr)
			assert.Equal(t, "Test summary", template.Summary)
			assert.Equal(t, "USER_API", *template.Source)
			assert.Equal(t, "SEVERITY_WARNING", *template.Severity)
			assert.Equal(t, "300s", template.For)
			assert.Len(t, template.Params, 1)

			param := template.Params[0]
			assert.Equal(t, "threshold", param.Name)
			assert.Equal(t, "test param summary", param.Summary)
			assert.Equal(t, "FLOAT", *param.Type)
			assert.Equal(t, "PERCENTAGE", *param.Unit)
			assert.Nil(t, param.Bool)
			assert.Nil(t, param.String)
			assert.NotNil(t, param.Float)

			float := param.Float
			assert.True(t, float.HasDefault)
			assert.Equal(t, float32(80), float.Default)
			assert.True(t, float.HasMax)
			assert.Equal(t, float32(100), float.Max)
			assert.True(t, float.HasMin)
			assert.Equal(t, float32(0), float.Min)

			assert.Equal(t, map[string]string{"foo": "bar"}, template.Labels)
			assert.Equal(t, map[string]string{"description": "test description", "summary": "test summary"}, template.Annotations)
			assert.Equal(t, file, template.Yaml)
			assert.NotEmpty(t, template.CreatedAt)
			found = true
		}
	}
	assert.Truef(t, found, "Template with id %s not found", name)
}
