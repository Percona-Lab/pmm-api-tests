package pmmapitests

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/agents"
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"
	"github.com/percona/pmm/api/inventorypb/json/client/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

type ErrorResponse interface {
	Code() int
}

// A minimal subset of *testing.T that we use that is also should be implemented by *expectedFailureTestingT.
type TestingT interface {
	Helper()
	Name() string
	Errorf(format string, args ...interface{})
	FailNow()
}

// TestString returns semi-random string that can be used as a test data.
func TestString(t TestingT, name string) string {
	t.Helper()

	n := rand.Int() //nolint:gosec
	return fmt.Sprintf("pmm-api-tests/%s/%s/%s/%d", Hostname, t.Name(), name, n)
}

// AssertAPIErrorf check that actual API error equals expected.
func AssertAPIErrorf(t TestingT, actual error, httpStatus int, grpcCode codes.Code, format string, a ...interface{}) {
	t.Helper()

	require.Error(t, actual)

	require.Implementsf(t, new(ErrorResponse), actual, "Wrong response type. Expected %T, got %T.\nError message: %v", new(ErrorResponse), actual, actual)

	assert.Equal(t, httpStatus, actual.(ErrorResponse).Code())

	// Have to use reflect because there are a lot of types with the same structure and different names.
	payload := reflect.ValueOf(actual).Elem().FieldByName("Payload")
	require.True(t, payload.IsValid(), "Wrong response structure. There is no field Payload.")

	codeField := payload.Elem().FieldByName("Code")
	require.True(t, codeField.IsValid(), "Wrong response structure. There is no field Code in Payload.")
	assert.Equal(t, int64(grpcCode), codeField.Int(), "gRPC status codes are not equal")

	errorField := payload.Elem().FieldByName("Error")
	require.True(t, errorField.IsValid(), "Wrong response structure. There is no field Error in Payload.")
	if len(a) > 0 {
		format = fmt.Sprintf(format, a...)
	}
	assert.Equal(t, format, errorField.String())
}

// AssertAPIErrorContains check that actual API error code equals expected and error text contains a string
func AssertAPIErrorContains(t TestingT, actual error, httpStatus int, grpcCode codes.Code, msg string) {
	t.Helper()

	require.Error(t, actual)

	require.Implementsf(t, new(ErrorResponse), actual, "Wrong response type. Expected %T, got %T.\nError message: %v", new(ErrorResponse), actual, actual)

	assert.Equal(t, httpStatus, actual.(ErrorResponse).Code())

	// Have to use reflect because there are a lot of types with the same structure and different names.
	payload := reflect.ValueOf(actual).Elem().FieldByName("Payload")
	require.True(t, payload.IsValid(), "Wrong response structure. There is no field Payload.")

	codeField := payload.Elem().FieldByName("Code")
	require.True(t, codeField.IsValid(), "Wrong response structure. There is no field Code in Payload.")
	assert.Equal(t, int64(grpcCode), codeField.Int(), "gRPC status codes are not equal")

	errorField := payload.Elem().FieldByName("Error")
	require.True(t, errorField.IsValid(), "Wrong response structure. There is no field Error in Payload.")
	assert.Contains(t, errorField.String(), msg)
}

func ExpectFailure(t *testing.T, link string) *expectedFailureTestingT {
	return &expectedFailureTestingT{
		t:    t,
		link: link,
	}
}

// expectedFailureTestingT expects that test will fail.
// if test is failed we skip it
// if it doesn't we call Fail
type expectedFailureTestingT struct {
	t      *testing.T
	errors []string
	failed bool
	link   string
}

func (tt *expectedFailureTestingT) Helper()      { tt.t.Helper() }
func (tt *expectedFailureTestingT) Name() string { return tt.t.Name() }

func (tt *expectedFailureTestingT) Errorf(format string, args ...interface{}) {
	tt.errors = append(tt.errors, fmt.Sprintf(format, args...))
	tt.failed = true
}

func (tt *expectedFailureTestingT) FailNow() {
	tt.failed = true

	// We have to set unexported testing.T.finished = true to make everything work,
	// but we can't call tt.t.FailNow() as it calls Fail().
	tt.t.SkipNow()
}

func (tt *expectedFailureTestingT) Check() {
	tt.t.Helper()

	if tt.failed {
		for _, v := range tt.errors {
			tt.t.Log(v)
		}
		tt.t.Skipf("Expected failure: %s", tt.link)
		return
	}

	tt.t.Fatalf("%s expected to fail, but didn't: %s", tt.Name(), tt.link)
}

func RemoveNodes(t TestingT, nodeIDs ...string) {
	t.Helper()

	for _, nodeID := range nodeIDs {
		params := &nodes.RemoveNodeParams{
			Body: nodes.RemoveNodeBody{
				NodeID: nodeID,
			},
			Context: context.Background(),
		}
		res, err := client.Default.Nodes.RemoveNode(params)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

func RemoveServices(t TestingT, serviceIDs ...string) {
	t.Helper()

	for _, serviceID := range serviceIDs {
		params := &services.RemoveServiceParams{
			Body: services.RemoveServiceBody{
				ServiceID: serviceID,
				//Force:     true,
			},
			Context: context.Background(),
		}
		res, err := client.Default.Services.RemoveService(params)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

func RemoveAgents(t TestingT, agentIDs ...string) {
	t.Helper()

	for _, agentID := range agentIDs {
		params := &agents.RemoveAgentParams{
			Body: agents.RemoveAgentBody{
				AgentID: agentID,
			},
			Context: context.Background(),
		}
		res, err := client.Default.Agents.RemoveAgent(params)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

// check interfaces
var (
	_ assert.TestingT  = (*expectedFailureTestingT)(nil)
	_ require.TestingT = (*expectedFailureTestingT)(nil)
	_ TestingT         = (*expectedFailureTestingT)(nil)
)
