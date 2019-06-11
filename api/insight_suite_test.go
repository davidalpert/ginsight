package api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"gopkg.in/resty.v1"

	api "github.com/davidalpert/ginsight/api"
)

func TestInsight(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Insight Suite")
}

var MockBaseURL = "https://jira.mydomain.com"

func MockURLFor(endpoint string) string {
	return MockBaseURL + endpoint
}

var testClient *api.Client

var _ = BeforeSuite(func() {
	// build the api client
	clientConfiguration := api.ClientConfiguration{
		BaseURL:  MockBaseURL,
		Username: "mal",
		Password: "serenity",
		Debug:    false, // toggle true to see Resty and other logs
	}

	if client, err := api.BuildClient(&clientConfiguration); err == nil {
		testClient = client
	}

	// block all HTTP requests
	httpmock.Activate()

	// and wire up the httpmock transport to resty
	httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
})

var _ = BeforeEach(func() {
	// remove any mocks
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})

// all the Insight API responses speak application/json so we need
// to feed that back to the resty library so that it knows how to
// unmarshall the responses.
func InsightApiResponder(statusCode int, body string) httpmock.Responder {
	response := httpmock.NewStringResponse(statusCode, body)
	response.Header.Add("content-type", "application/json")
	return httpmock.ResponderFromResponse(response)
}
