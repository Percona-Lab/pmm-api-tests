package ia

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/brianvoe/gofakeit"
	"github.com/percona-platform/saas/pkg/alert"
	templatesClient "github.com/percona/pmm/api/managementpb/ia/json/client"
	"github.com/percona/pmm/api/managementpb/ia/json/client/rules"
	"github.com/percona/pmm/api/managementpb/ia/json/client/templates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"google.golang.org/grpc/codes"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

// Note: Even though the IA services check for alerting enabled or disabled before returning results
// we don't enable or disable IA explicit in our tests since it is enabled by default through
// ENABLE_ALERTING env var.
func TestAddTemplate(t *testing.T) {
	client := templatesClient.Default.Templates

	b, err := ioutil.ReadFile("../../testdata/ia/template.yaml")
	require.NoError(t, err)

	t.Run("normal", func(t *testing.T) {
		name := gofakeit.UUID()
		expr := gofakeit.UUID()
		yml := formatTemplateYaml(t, fmt.Sprintf(string(b), name, expr, "%", "s"))
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: yml,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		defer deleteTemplate(t, client, name)

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
				assert.Equal(t, yml, template.Yaml)
				assert.Equal(t, "Test summary", template.Summary)
				assert.Equal(t, expr, template.Expr)
				assert.Len(t, template.Params, 2)

				assert.Equal(t, "param1", template.Params[0].Name)
				assert.Equal(t, "first parameter with default value and defined range", template.Params[0].Summary)
				assert.Equal(t, "PERCENTAGE", *template.Params[0].Unit)
				assert.Equal(t, "FLOAT", *template.Params[0].Type)
				assert.True(t, template.Params[0].Float.HasDefault)
				assert.Equal(t, float32(80), template.Params[0].Float.Default)
				assert.True(t, template.Params[0].Float.HasMax)
				assert.Equal(t, float32(100), template.Params[0].Float.Max)
				assert.True(t, template.Params[0].Float.HasMin)
				assert.Equal(t, float32(0), template.Params[0].Float.Min)

				assert.Equal(t, "param2", template.Params[1].Name)
				assert.Equal(t, "second parameter without default value and defined range", template.Params[1].Summary)
				assert.Equal(t, "SECONDS", *template.Params[1].Unit)
				assert.Equal(t, "FLOAT", *template.Params[1].Type)
				assert.False(t, template.Params[1].Float.HasDefault)
				assert.Equal(t, float32(0), template.Params[1].Float.Default)
				assert.False(t, template.Params[1].Float.HasMax)
				assert.Equal(t, float32(0), template.Params[1].Float.Max)
				assert.False(t, template.Params[1].Float.HasMin)
				assert.Equal(t, float32(0), template.Params[1].Float.Min)
				found = true
			}
		}
		assert.Truef(t, found, "Template with id %s not found", name)
	})

	t.Run("duplicate", func(t *testing.T) {
		name := gofakeit.UUID()
		yaml := fmt.Sprintf(string(b), name, gofakeit.UUID(), "s", "%")
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: yaml,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		defer deleteTemplate(t, client, name)

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
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID(), "s", "%"),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		defer deleteTemplate(t, client, name)

		newExpr := gofakeit.UUID()
		yml := formatTemplateYaml(t, fmt.Sprintf(string(b), name, newExpr, "s", "%"))
		_, err = client.UpdateTemplate(&templates.UpdateTemplateParams{
			Body: templates.UpdateTemplateBody{
				Yaml: yml,
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
				assert.Equal(t, yml, template.Yaml)
				assert.Equal(t, "Test summary", template.Summary)
				assert.Len(t, template.Params, 2)

				assert.Equal(t, "param1", template.Params[0].Name)
				assert.Equal(t, "first parameter with default value and defined range", template.Params[0].Summary)
				assert.Equal(t, "SECONDS", *template.Params[0].Unit)
				assert.Equal(t, "FLOAT", *template.Params[0].Type)
				assert.True(t, template.Params[0].Float.HasDefault)
				assert.Equal(t, float32(80), template.Params[0].Float.Default)
				assert.True(t, template.Params[0].Float.HasMax)
				assert.Equal(t, float32(100), template.Params[0].Float.Max)
				assert.True(t, template.Params[0].Float.HasMin)
				assert.Equal(t, float32(0), template.Params[0].Float.Min)

				assert.Equal(t, "param2", template.Params[1].Name)
				assert.Equal(t, "second parameter without default value and defined range", template.Params[1].Summary)
				assert.Equal(t, "PERCENTAGE", *template.Params[1].Unit)
				assert.Equal(t, "FLOAT", *template.Params[1].Type)
				assert.False(t, template.Params[1].Float.HasDefault)
				assert.Equal(t, float32(0), template.Params[1].Float.Default)
				assert.False(t, template.Params[1].Float.HasMax)
				assert.Equal(t, float32(0), template.Params[1].Float.Max)
				assert.False(t, template.Params[1].Float.HasMin)
				assert.Equal(t, float32(0), template.Params[1].Float.Min)
				found = true
			}
		}
		assert.Truef(t, found, "Template with id %s not found", name)
	})

	t.Run("unknown template", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err = client.UpdateTemplate(&templates.UpdateTemplateParams{
			Body: templates.UpdateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID(), "s", "%"),
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 404, codes.NotFound, fmt.Sprintf("Template with name \"%s\" not found.", name))
	})

	t.Run("invalid yaml", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID(), "s", "%"),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		defer deleteTemplate(t, client, name)

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
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID(), "s", "%"),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		defer deleteTemplate(t, client, name)

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
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID(), "s", "%"),
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

	t.Run("template in use", func(t *testing.T) {
		name := gofakeit.UUID()
		_, err := client.CreateTemplate(&templates.CreateTemplateParams{
			Body: templates.CreateTemplateBody{
				Yaml: fmt.Sprintf(string(b), name, gofakeit.UUID(), "s", "%"),
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)
		defer deleteTemplate(t, templatesClient.Default.Templates, name)

		channelID := createChannel(t)
		defer deleteChannel(t, templatesClient.Default.Channels, channelID)

		params := createAlertRuleParams(name, channelID, "param2", &rules.FiltersItems0{
			Type:  pointer.ToString("EQUAL"),
			Key:   "threshold",
			Value: "12",
		})

		rule, err := templatesClient.Default.Rules.CreateAlertRule(params)
		require.NoError(t, err)
		defer deleteRule(t, templatesClient.Default.Rules, rule.Payload.RuleID)

		_, err = client.DeleteTemplate(&templates.DeleteTemplateParams{
			Body: templates.DeleteTemplateBody{
				Name: name,
			},
			Context: pmmapitests.Context,
		})
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, "Failed to delete rule template %s, as it is being used by some rule.", name)

		resp, err := client.ListTemplates(&templates.ListTemplatesParams{
			Body: templates.ListTemplatesBody{
				Reload: true,
			},
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		var found bool
		for _, template := range resp.Payload.Templates {
			if name == template.Name {
				found = true
			}
		}
		assert.Truef(t, found, "Template with id %s not found", name)
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
	yml := formatTemplateYaml(t, fmt.Sprintf(string(b), name, expr, "%", "s"))
	_, err = client.CreateTemplate(&templates.CreateTemplateParams{
		Body: templates.CreateTemplateBody{
			Yaml: yml,
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)
	defer deleteTemplate(t, client, name)

	t.Run("without pagination", func(t *testing.T) {
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
				assert.Len(t, template.Params, 2)

				assert.Equal(t, "param1", template.Params[0].Name)
				assert.Equal(t, "first parameter with default value and defined range", template.Params[0].Summary)
				assert.Equal(t, "FLOAT", *template.Params[0].Type)
				assert.Equal(t, "PERCENTAGE", *template.Params[0].Unit)
				assert.Nil(t, template.Params[0].Bool)
				assert.Nil(t, template.Params[0].String)
				assert.NotNil(t, template.Params[0].Float)
				assert.True(t, template.Params[0].Float.HasDefault)
				assert.Equal(t, float32(80), template.Params[0].Float.Default)
				assert.True(t, template.Params[0].Float.HasMax)
				assert.Equal(t, float32(100), template.Params[0].Float.Max)
				assert.True(t, template.Params[0].Float.HasMin)
				assert.Equal(t, float32(0), template.Params[0].Float.Min)

				assert.Equal(t, "param2", template.Params[1].Name)
				assert.Equal(t, "second parameter without default value and defined range", template.Params[1].Summary)
				assert.Equal(t, "FLOAT", *template.Params[1].Type)
				assert.Equal(t, "SECONDS", *template.Params[1].Unit)
				assert.Nil(t, template.Params[1].Bool)
				assert.Nil(t, template.Params[1].String)
				assert.NotNil(t, template.Params[1].Float)
				assert.False(t, template.Params[1].Float.HasDefault)
				assert.Equal(t, float32(0), template.Params[1].Float.Default)
				assert.False(t, template.Params[1].Float.HasMax)
				assert.Equal(t, float32(0), template.Params[1].Float.Max)
				assert.False(t, template.Params[1].Float.HasMin)
				assert.Equal(t, float32(00), template.Params[1].Float.Min)

				assert.Equal(t, map[string]string{"foo": "bar"}, template.Labels)
				assert.Equal(t, map[string]string{"description": "test description", "summary": "test summary"}, template.Annotations)
				assert.Equal(t, yml, template.Yaml)
				assert.NotEmpty(t, template.CreatedAt)
				found = true
			}
		}
		assert.Truef(t, found, "Template with id %s not found", name)
	})

	t.Run("with pagination", func(t *testing.T) {
		const templatesCount = 5

		templateNames := make(map[string]struct{})

		for i := 0; i < templatesCount; i++ {
			name := gofakeit.UUID()
			expr := gofakeit.UUID()
			yml := formatTemplateYaml(t, fmt.Sprintf(string(b), name, expr, "%", "s"))
			_, err = client.CreateTemplate(&templates.CreateTemplateParams{
				Body: templates.CreateTemplateBody{
					Yaml: yml,
				},
				Context: pmmapitests.Context,
			})

			templateNames[name] = struct{}{}
		}
		defer func() {
			for name := range templateNames {
				deleteTemplate(t, client, name)
			}
		}()

		// list rules, so they are all on the first page
		body := templates.ListTemplatesBody{
			PageParams: &templates.ListTemplatesParamsBodyPageParams{
				PageSize: 20,
				Index:    0,
			},
		}
		list1, err := client.ListTemplates(&templates.ListTemplatesParams{
			Body:    body,
			Context: pmmapitests.Context,
		})
		require.NoError(t, err)

		assert.GreaterOrEqual(t, len(list1.Payload.Templates), templatesCount)
		assert.Equal(t, int32(len(list1.Payload.Templates)), list1.Payload.Totals.TotalItems)
		assert.Equal(t, int32(1), list1.Payload.Totals.TotalPages)

		assertFindTemplate := func(list []*templates.TemplatesItems0, name string) func() bool {
			return func() bool {
				for _, tmpl := range list {
					if tmpl.Name == name {
						return true
					}
				}
				return false
			}
		}

		for name := range templateNames {
			assert.Conditionf(t, assertFindTemplate(list1.Payload.Templates, name), "template %s not found", name)
		}

		// paginate page over page with page size 1 and check the order - it should be the same as in list1.
		// last iteration checks that there is no elements for not existing page.
		for pageIndex := 0; pageIndex <= len(list1.Payload.Templates); pageIndex++ {
			body := templates.ListTemplatesBody{
				PageParams: &templates.ListTemplatesParamsBodyPageParams{
					PageSize: 1,
					Index:    int32(pageIndex),
				},
			}
			list2, err := client.ListTemplates(&templates.ListTemplatesParams{
				Body: body, Context: pmmapitests.Context,
			})
			require.NoError(t, err)

			assert.Equal(t, list2.Payload.Totals.TotalItems, list2.Payload.Totals.TotalItems)
			assert.GreaterOrEqual(t, list2.Payload.Totals.TotalPages, int32(templatesCount))

			if pageIndex != len(list1.Payload.Templates) {
				require.Len(t, list2.Payload.Templates, 1)
				assert.Equal(t, list1.Payload.Templates[pageIndex].Name, list2.Payload.Templates[0].Name)
			} else {
				assert.Len(t, list2.Payload.Templates, 0)
			}
		}

	})
}

func deleteTemplate(t *testing.T, client templates.ClientService, name string) {
	_, err := client.DeleteTemplate(&templates.DeleteTemplateParams{
		Body: templates.DeleteTemplateBody{
			Name: name,
		},
		Context: pmmapitests.Context,
	})
	assert.NoError(t, err)
}

func formatTemplateYaml(t *testing.T, yml string) string {
	params := &alert.ParseParams{
		DisallowUnknownFields:    true,
		DisallowInvalidTemplates: true,
	}
	r, err := alert.Parse(strings.NewReader(yml), params)
	require.NoError(t, err)
	type templates struct {
		Templates []alert.Template `yaml:"templates"`
	}

	s, err := yaml.Marshal(&templates{Templates: r})
	require.NoError(t, err)

	return string(s)
}
