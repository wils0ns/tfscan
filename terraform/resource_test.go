package terraform_test

import (
	"encoding/json"
	"testing"

	"github.com/wils0ns/tfscan/terraform"
	"github.com/wils0ns/tfscan/testdata"
)

func TestResourceEquals(t *testing.T) {
	testdata.StateReader.Seek(0, 0)
	state, err := terraform.NewState(testdata.StateReader)
	if err != nil {
		t.Fatal(err)
	}

	resList, err := state.ResourceLookup("google_app_engine_application.default")
	if err != nil {
		t.Fatal(err)
	}

	resCopy := *resList[0]
	if !resCopy.Equals(resList[0]) {
		t.Error("Expected resource copy to be equal to original")
	}

	testdata.StateReader.Seek(0, 0)
	stateCopy, err := terraform.NewState(testdata.StateReader)
	if err != nil {
		t.Fatal(err)
	}
	resListCopy, err := stateCopy.ResourceLookup("google_app_engine_application.default")
	if err != nil {
		t.Fatal(err)
	}

	resListCopy[0].Values["app_id"] = "different"

	if resListCopy[0].Equals(resList[0]) {
		t.Error("Expected modified resource to be different than original")
	}

	resJSON := `
	{
    "address": "google_app_engine_application.default",
    "mode": "managed",
    "type": "google_app_engine_application",
    "name": "default",
    "index": "",
    "provider_name": "google",
    "schema_version": 0,
    "values": {
      "app_id": "myproject-example-1.NOT-EQUAL",
      "auth_domain": "gmail.com",
      "code_bucket": "staging.myproject-example-1.appspot.com",
      "database_type": "CLOUD_DATASTORE_COMPATIBILITY",
      "default_bucket": "myproject-example-1.appspot.com",
      "default_hostname": "myproject-example-1.ey.r.appspot.com",
      "feature_settings": [
        {
          "split_health_checks": true
        }
      ],
      "gcr_domain": "eu.gcr.io",
      "iap": [],
      "id": "myproject-example-1",
      "location_id": "europe-west3",
      "name": "apps/myproject-example-1",
      "project": "myproject-example-1",
      "serving_status": "SERVING",
      "timeouts": null,
      "url_dispatch_rule": []
    }
	}
	`

	res := &terraform.Resource{}
	json.Unmarshal([]byte(resJSON), res)

	if res.Equals(resList[0]) {
		t.Error("Expected modified resource to NOT be equal to original")
	}

}
