package terraform_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wils0ns/tfscan/terraform"
	"github.com/wils0ns/tfscan/testdata"
)

// TODO: use a better testing package

func TestNewState(t *testing.T) {
	testdata.SampleReader.Seek(0, 0)
	_, err := terraform.NewState(testdata.SampleReader)
	if err != nil {
		t.Error(err)
	}

	r := strings.NewReader("")
	_, err = terraform.NewState(r)
	expectedError := "*json.SyntaxError"
	actualError := fmt.Sprintf("%T", err)
	if actualError != expectedError {
		t.Error("Expected error:", expectedError, "Got: ", actualError)
	}
}

func TestGetResourcesByFullAddress(t *testing.T) {
	testdata.SampleReader.Seek(0, 0)
	state, err := terraform.NewState(testdata.SampleReader)
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []struct {
		AddressRegExp     string
		ExpectedAddresses []string
		ExpectedError     error
	}{
		{
			AddressRegExp: "module.project.google",
			ExpectedAddresses: []string{
				"google_project.default",
				"google_project_iam_audit_config.audit_config",
				"google_project_iam_member.owner",
			},
			ExpectedError: nil,
		},
		{
			AddressRegExp:     "google_project_service.default[\"storage.googleapis.com\"]",
			ExpectedAddresses: []string{"google_project_service.default"},
			ExpectedError:     nil,
		},
		{
			AddressRegExp:     "no.resource",
			ExpectedAddresses: []string{},
			ExpectedError:     &terraform.ResourceNotFoundError{Address: "no.resource"},
		},
	}

	for _, s := range scenarios {
		resources, err := state.ResourceLookup(s.AddressRegExp)

		errMsg := ""
		expectedErrMsg := ""

		if err != nil && s.ExpectedError == nil {
			t.Fatal(err)
		}

		if err != nil {
			errMsg = err.Error()
		}

		if s.ExpectedError != nil {
			expectedErrMsg = s.ExpectedError.Error()
		}

		if errMsg != expectedErrMsg {
			t.Fatal("Expected:", expectedErrMsg, "Got:", errMsg)
		}

		if s.ExpectedError == nil {
			addresses := []string{}
			for _, res := range resources {
				addresses = append(addresses, res.Address)
			}
			for i := range addresses {
				if addresses[i] != s.ExpectedAddresses[i] {
					t.Error("Expected:", s.ExpectedAddresses[i], "Got:", addresses[i])
				}
			}
		}

	}
}

func TestResourceTypes(t *testing.T) {
	testdata.SampleReader.Seek(0, 0)
	state, err := terraform.NewState(testdata.SampleReader)
	if err != nil {
		t.Fatal(err)
	}

	expectedTypes := []string{
		"google_app_engine_application",
		"google_bigquery_dataset",
		"google_logging_project_sink",
		"google_project",
		"google_project_iam_audit_config",
		"google_project_iam_member",
		"google_project_service",
		"google_storage_bucket",
	}

	resTypes, err := state.ResourceTypes()
	if err != nil {
		t.Fatal(err)
	}

	numResTypes := len(resTypes)
	expectedNumResTypes := len(expectedTypes)

	if numResTypes == expectedNumResTypes {
		for i := range resTypes {
			if resTypes[i] != expectedTypes[i] {
				t.Errorf("Expected: %v, Got: %v", expectedTypes[i], resTypes[i])
			}
		}
	} else {
		t.Errorf("Expected number of resource types to be: %v, got: %v", expectedNumResTypes, numResTypes)
	}
}
