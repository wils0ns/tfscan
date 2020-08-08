package terraform_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wils0ns/tfscan/terraform"
	"github.com/wils0ns/tfscan/testdata"
)

// TODO: use a better testing package

var sampleReader = strings.NewReader(testdata.SampleJSONState)

func TestNewState(t *testing.T) {
	_, err := terraform.NewState(sampleReader)
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

func TestGetResourceByFullAddress(t *testing.T) {
	sampleReader.Seek(0, 0)
	state, err := terraform.NewState(sampleReader)
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []struct {
		FullAddress     string
		ExpectedAddress string
		ExpectedError   error
	}{
		{
			FullAddress:     "module.project.google_project.default",
			ExpectedAddress: "google_project.default",
			ExpectedError:   nil,
		},
		{
			FullAddress:     "google_project_service.default[\"storage.googleapis.com\"]",
			ExpectedAddress: "google_project_service.default",
			ExpectedError:   nil,
		},
		{
			FullAddress:     "no.resource",
			ExpectedAddress: "",
			ExpectedError:   &terraform.ResourceNotFoundError{Address: "no.resource"},
		},
	}

	for _, s := range scenarios {
		res, err := state.GetResourceByFullAddress(s.FullAddress)

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
			actual := res.Address
			if actual != s.ExpectedAddress {
				t.Error("Expected:", s.ExpectedAddress, "Got:", actual)
			}
		}

	}
}

func TestResourceTypes(t *testing.T) {
	sampleReader.Seek(0, 0)
	state, err := terraform.NewState(sampleReader)
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
