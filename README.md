# tfscan

Inspect Terraform resources in a state and plan JSON files

## Install

Using go:

```bash
git clone https://github.com/wils0ns/tfscan.git
cd tfscan
go install
```

Or download the binary for a particular [release](https://github.com/wils0ns/tfscan/releases).

## Examples

### Reading from terraform stdout

Command:

```bash
terraform show -json | tfscan
```

Output:

```bash
root_module:
├─ google_app_engine_application.default
├─ google_project_iam_member.datastore_import_export_admin["serviceAccount:service01@project-example.iam.gserviceaccount.com"]
├─ google_project_iam_member.datastore_user["serviceAccount:service01@project-example.iam.gserviceaccount.com"]
├─ google_project_iam_member.logging_config_writer["serviceAccount:service02@project-example.iam.gserviceaccount.com"]
├─ google_project_iam_member.spanner_admin["serviceAccount:service01@project-example.iam.gserviceaccount.com"]
├─ google_project_iam_member.storage_admin["serviceAccount:service01@project-example.iam.gserviceaccount.com"]
├─ google_project_iam_member.storage_admin["serviceAccount:project-example@appspot.gserviceaccount.com"]
├─ google_project_service.default["spanner.googleapis.com"]
├─ google_project_service.default["datastore.googleapis.com"]
├─ google_project_service.default["appengine.googleapis.com"]
├─ google_project_service.default["storage.googleapis.com"]
├─ google_project_service.default["storage-component.googleapis.com"]
├─ google_storage_bucket.default["bucket01"]
├─ google_storage_bucket.default["bucket02"]
├─ google_storage_bucket.default["bucket03"]
├─ module.project:
│  ├─ google_project.default
│  ├─ google_project_iam_audit_config.audit_config
│  ├─ google_project_iam_member.owner
│  ├─ module.project.module.log_exporter:
│  │  ├─ google_bigquery_dataset.dataset
│  │  ├─ google_logging_project_sink.sink
│  │  ├─ google_project_service.bigquery
```

### List all unique resource types

Command:

```bash
tfscan -json state.json -types
```

Output:

```bash
[
  "google_app_engine_application",
  "google_bigquery_dataset",
  "google_logging_project_sink",
  "google_project",
  "google_project_iam_audit_config",
  "google_project_iam_member",
  "google_project_service",
  "google_storage_bucket"
]

```

### Get resources by regular expression

Command:

```bash
tfscan -json state1.json -get google_project_service.default
```

Output:

```bash
[
  {
    "address": "google_project_service.default",
    "mode": "managed",
    "type": "google_project_service",
    "name": "default",
    "index": "datastore.googleapis.com",
    "provider_name": "google",
    "schema_version": 0,
    "values": {
      "disable_dependent_services": null,
      "disable_on_destroy": false,
      "id": "myproject-example-1/datastore.googleapis.com",
      "project": "myproject-example-1",
      "service": "datastore.googleapis.com",
      "timeouts": null
    }
  },
  {
    "address": "google_project_service.default",
    "mode": "managed",
    "type": "google_project_service",
    "name": "default",
    "index": "storage-component.googleapis.com",
    "provider_name": "google",
    "schema_version": 0,
    "values": {
      "disable_dependent_services": null,
      "disable_on_destroy": false,
      "id": "myproject-example-1/storage-component.googleapis.com",
      "project": "myproject-example-1",
      "service": "storage-component.googleapis.com",
      "timeouts": null
    }
  },
]
```

### Diff state resources

Command:

```bash
tfscan -json state2.json -diff state1.json -get google_app_engine_application.default
```

Output:

```bash
 [
   {
     "address": "google_app_engine_application.default",
     "mode": "managed",
     "type": "google_app_engine_application",
     "name": "default",
     "index": "",
     "provider_name": "google",
     "schema_version": 0,
     "values": {
-      "app_id": "myproject-example-1",
+      "app_id": "myproject-example",
       "auth_domain": "gmail.com",
-      "code_bucket": "staging.myproject-example-1.appspot.com",
+      "code_bucket": "staging.myproject-example.appspot.com",
       "database_type": "CLOUD_DATASTORE_COMPATIBILITY",
-      "default_bucket": "myproject-example-1.appspot.com",
-      "default_hostname": "myproject-example-1.ey.r.appspot.com",
+      "default_bucket": "myproject-example.appspot.com",
+      "default_hostname": "myproject-example.wl.r.appspot.com",
       "feature_settings": [
         {
           "split_health_checks": true
         }
       ],
-      "gcr_domain": "eu.gcr.io",
+      "gcr_domain": "us.gcr.io",
       "iap": [],
-      "id": "myproject-example-1",
-      "location_id": "europe-west3",
-      "name": "apps/myproject-example-1",
-      "project": "myproject-example-1",
+      "id": "myproject-example",
+      "location_id": "us-west2",
+      "name": "apps/myproject-example",
+      "project": "myproject-example",
       "serving_status": "SERVING",
       "timeouts": null,
       "url_dispatch_rule": []
     }
   }
 ]
```

### Diff resource types between states

Command:

```bash
tfscan -json testdata/state2.json -diff testdata/state1.json -types
```

Output:

```bash
 [
   "google_app_engine_application",
   "google_bigquery_dataset",
+  "google_cloud_scheduler_job",
+  "google_cloudfunctions_function",
+  "google_datastore_index",
+  "google_logging_metric",
   "google_logging_project_sink",
   "google_project",
   "google_project_iam_audit_config",
   "google_project_iam_member",
   "google_project_service",
-  "google_storage_bucket"
+  "google_pubsub_topic",
+  "google_service_account",
+  "google_storage_bucket",
+  "google_storage_bucket_object",
+  "null_resource"
 ]
```

Diff Legend:

* `-` : `state1.json` does not have this content and `state2.json` does.
* `+` : `state1.json` has this content and `state2.json` does not.
* ` ` : `state1.json` has the same content as `state2.json`
