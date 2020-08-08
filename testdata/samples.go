package testdata

import "os"

func TestPath() (string, error) {
	return os.Getwd()
}

var SampleJSONState = `
{
  "format_version": "0.1",
  "terraform_version": "0.12.29",
  "values": {
    "root_module": {
      "resources": [
        {
          "address": "google_app_engine_application.default",
          "mode": "managed",
          "type": "google_app_engine_application",
          "name": "default",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "app_id": "di-test-blueprint-odrs-data-1",
            "auth_domain": "gmail.com",
            "code_bucket": "staging.di-test-blueprint-odrs-data-1.appspot.com",
            "database_type": "CLOUD_DATASTORE_COMPATIBILITY",
            "default_bucket": "di-test-blueprint-odrs-data-1.appspot.com",
            "default_hostname": "di-test-blueprint-odrs-data-1.ey.r.appspot.com",
            "feature_settings": [
              {
                "split_health_checks": true
              }
            ],
            "gcr_domain": "eu.gcr.io",
            "iap": [],
            "id": "di-test-blueprint-odrs-data-1",
            "location_id": "europe-west3",
            "name": "apps/di-test-blueprint-odrs-data-1",
            "project": "di-test-blueprint-odrs-data-1",
            "serving_status": "SERVING",
            "timeouts": null,
            "url_dispatch_rule": []
          },
          "depends_on": [
            "module.project.google_project.default",
            "module.project.google_project_iam_audit_config.audit_config",
            "module.project.google_project_iam_member.owner",
            "module.project.module.log_exporter.google_bigquery_dataset.dataset",
            "module.project.module.log_exporter.google_logging_project_sink.sink",
            "module.project.module.log_exporter.google_project_service.bigquery"
          ]
        },
        {
          "address": "google_project_iam_member.datastore_import_export_admin",
          "mode": "managed",
          "type": "google_project_iam_member",
          "name": "datastore_import_export_admin",
          "index": "serviceAccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "condition": [],
            "etag": "BwWsIPSdhKk=",
            "id": "di-test-blueprint-odrs-data-1/roles/datastore.importExportAdmin/serviceaccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
            "member": "serviceAccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
            "project": "di-test-blueprint-odrs-data-1",
            "role": "roles/datastore.importExportAdmin"
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        },
        {
          "address": "google_project_iam_member.datastore_user",
          "mode": "managed",
          "type": "google_project_iam_member",
          "name": "datastore_user",
          "index": "serviceAccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "condition": [],
            "etag": "BwWsIPSdhKk=",
            "id": "di-test-blueprint-odrs-data-1/roles/datastore.user/serviceaccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
            "member": "serviceAccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
            "project": "di-test-blueprint-odrs-data-1",
            "role": "roles/datastore.user"
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        },
        {
          "address": "google_project_iam_member.logging_config_writer",
          "mode": "managed",
          "type": "google_project_iam_member",
          "name": "logging_config_writer",
          "index": "serviceAccount:buildaccount@dre-configuration.iam.gserviceaccount.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "condition": [],
            "etag": "BwWsIPSdhKk=",
            "id": "di-test-blueprint-odrs-data-1/roles/logging.configWriter/serviceaccount:buildaccount@dre-configuration.iam.gserviceaccount.com",
            "member": "serviceAccount:buildaccount@dre-configuration.iam.gserviceaccount.com",
            "project": "di-test-blueprint-odrs-data-1",
            "role": "roles/logging.configWriter"
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        },
        {
          "address": "google_project_iam_member.spanner_admin",
          "mode": "managed",
          "type": "google_project_iam_member",
          "name": "spanner_admin",
          "index": "serviceAccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "condition": [],
            "etag": "BwWsIPSdhKk=",
            "id": "di-test-blueprint-odrs-data-1/roles/spanner.admin/serviceaccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
            "member": "serviceAccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
            "project": "di-test-blueprint-odrs-data-1",
            "role": "roles/spanner.admin"
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        },
        {
          "address": "google_project_iam_member.storage_admin",
          "mode": "managed",
          "type": "google_project_iam_member",
          "name": "storage_admin",
          "index": "serviceAccount:p-dre-eu-services@appspot.gserviceaccount.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "condition": [],
            "etag": "BwWsIPSdhKk=",
            "id": "di-test-blueprint-odrs-data-1/roles/storage.admin/serviceaccount:p-dre-eu-services@appspot.gserviceaccount.com",
            "member": "serviceAccount:p-dre-eu-services@appspot.gserviceaccount.com",
            "project": "di-test-blueprint-odrs-data-1",
            "role": "roles/storage.admin"
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        },
        {
          "address": "google_project_iam_member.storage_admin",
          "mode": "managed",
          "type": "google_project_iam_member",
          "name": "storage_admin",
          "index": "serviceAccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "condition": [],
            "etag": "BwWsIPSdhKk=",
            "id": "di-test-blueprint-odrs-data-1/roles/storage.admin/serviceaccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
            "member": "serviceAccount:gke-sa@p-dre-eu-services.iam.gserviceaccount.com",
            "project": "di-test-blueprint-odrs-data-1",
            "role": "roles/storage.admin"
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        },
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
            "id": "di-test-blueprint-odrs-data-1/datastore.googleapis.com",
            "project": "di-test-blueprint-odrs-data-1",
            "service": "datastore.googleapis.com",
            "timeouts": null
          },
          "depends_on": [
            "module.project.google_project.default",
            "module.project.google_project_iam_audit_config.audit_config",
            "module.project.google_project_iam_member.owner",
            "module.project.module.log_exporter.google_bigquery_dataset.dataset",
            "module.project.module.log_exporter.google_logging_project_sink.sink",
            "module.project.module.log_exporter.google_project_service.bigquery"
          ]
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
            "id": "di-test-blueprint-odrs-data-1/storage-component.googleapis.com",
            "project": "di-test-blueprint-odrs-data-1",
            "service": "storage-component.googleapis.com",
            "timeouts": null
          },
          "depends_on": [
            "module.project.google_project.default",
            "module.project.google_project_iam_audit_config.audit_config",
            "module.project.google_project_iam_member.owner",
            "module.project.module.log_exporter.google_bigquery_dataset.dataset",
            "module.project.module.log_exporter.google_logging_project_sink.sink",
            "module.project.module.log_exporter.google_project_service.bigquery"
          ]
        },
        {
          "address": "google_project_service.default",
          "mode": "managed",
          "type": "google_project_service",
          "name": "default",
          "index": "storage.googleapis.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "disable_dependent_services": null,
            "disable_on_destroy": false,
            "id": "di-test-blueprint-odrs-data-1/storage.googleapis.com",
            "project": "di-test-blueprint-odrs-data-1",
            "service": "storage.googleapis.com",
            "timeouts": null
          },
          "depends_on": [
            "module.project.google_project.default",
            "module.project.google_project_iam_audit_config.audit_config",
            "module.project.google_project_iam_member.owner",
            "module.project.module.log_exporter.google_bigquery_dataset.dataset",
            "module.project.module.log_exporter.google_logging_project_sink.sink",
            "module.project.module.log_exporter.google_project_service.bigquery"
          ]
        },
        {
          "address": "google_project_service.default",
          "mode": "managed",
          "type": "google_project_service",
          "name": "default",
          "index": "spanner.googleapis.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "disable_dependent_services": null,
            "disable_on_destroy": false,
            "id": "di-test-blueprint-odrs-data-1/spanner.googleapis.com",
            "project": "di-test-blueprint-odrs-data-1",
            "service": "spanner.googleapis.com",
            "timeouts": null
          },
          "depends_on": [
            "module.project.google_project.default",
            "module.project.google_project_iam_audit_config.audit_config",
            "module.project.google_project_iam_member.owner",
            "module.project.module.log_exporter.google_bigquery_dataset.dataset",
            "module.project.module.log_exporter.google_logging_project_sink.sink",
            "module.project.module.log_exporter.google_project_service.bigquery"
          ]
        },
        {
          "address": "google_project_service.default",
          "mode": "managed",
          "type": "google_project_service",
          "name": "default",
          "index": "appengine.googleapis.com",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "disable_dependent_services": null,
            "disable_on_destroy": false,
            "id": "di-test-blueprint-odrs-data-1/appengine.googleapis.com",
            "project": "di-test-blueprint-odrs-data-1",
            "service": "appengine.googleapis.com",
            "timeouts": null
          },
          "depends_on": [
            "module.project.google_project.default",
            "module.project.google_project_iam_audit_config.audit_config",
            "module.project.google_project_iam_member.owner",
            "module.project.module.log_exporter.google_bigquery_dataset.dataset",
            "module.project.module.log_exporter.google_logging_project_sink.sink",
            "module.project.module.log_exporter.google_project_service.bigquery"
          ]
        },
        {
          "address": "google_storage_bucket.default",
          "mode": "managed",
          "type": "google_storage_bucket",
          "name": "default",
          "index": "gridstore-v0-services-di-test-blueprint-odrs-data-1",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "bucket_policy_only": false,
            "cors": [],
            "default_event_based_hold": false,
            "encryption": [],
            "force_destroy": false,
            "id": "gridstore-v0-services-di-test-blueprint-odrs-data-1",
            "labels": null,
            "lifecycle_rule": [],
            "location": "EUROPE-WEST3",
            "logging": [],
            "name": "gridstore-v0-services-di-test-blueprint-odrs-data-1",
            "project": "di-test-blueprint-odrs-data-1",
            "requester_pays": false,
            "retention_policy": [],
            "self_link": "https://www.googleapis.com/storage/v1/b/gridstore-v0-services-di-test-blueprint-odrs-data-1",
            "storage_class": "STANDARD",
            "url": "gs://gridstore-v0-services-di-test-blueprint-odrs-data-1",
            "versioning": [],
            "website": []
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        },
        {
          "address": "google_storage_bucket.default",
          "mode": "managed",
          "type": "google_storage_bucket",
          "name": "default",
          "index": "digital-re-di-test-blueprint-odrs-data-1",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "bucket_policy_only": false,
            "cors": [],
            "default_event_based_hold": false,
            "encryption": [],
            "force_destroy": false,
            "id": "digital-re-di-test-blueprint-odrs-data-1",
            "labels": null,
            "lifecycle_rule": [],
            "location": "EUROPE-WEST3",
            "logging": [],
            "name": "digital-re-di-test-blueprint-odrs-data-1",
            "project": "di-test-blueprint-odrs-data-1",
            "requester_pays": false,
            "retention_policy": [],
            "self_link": "https://www.googleapis.com/storage/v1/b/digital-re-di-test-blueprint-odrs-data-1",
            "storage_class": "STANDARD",
            "url": "gs://digital-re-di-test-blueprint-odrs-data-1",
            "versioning": [],
            "website": []
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        },
        {
          "address": "google_storage_bucket.default",
          "mode": "managed",
          "type": "google_storage_bucket",
          "name": "default",
          "index": "gridstore-v0-di-test-blueprint-odrs-data-1",
          "provider_name": "google",
          "schema_version": 0,
          "values": {
            "bucket_policy_only": false,
            "cors": [],
            "default_event_based_hold": false,
            "encryption": [],
            "force_destroy": false,
            "id": "gridstore-v0-di-test-blueprint-odrs-data-1",
            "labels": null,
            "lifecycle_rule": [],
            "location": "EUROPE-WEST3",
            "logging": [],
            "name": "gridstore-v0-di-test-blueprint-odrs-data-1",
            "project": "di-test-blueprint-odrs-data-1",
            "requester_pays": false,
            "retention_policy": [],
            "self_link": "https://www.googleapis.com/storage/v1/b/gridstore-v0-di-test-blueprint-odrs-data-1",
            "storage_class": "STANDARD",
            "url": "gs://gridstore-v0-di-test-blueprint-odrs-data-1",
            "versioning": [],
            "website": []
          },
          "depends_on": [
            "module.project.google_project.default"
          ]
        }
      ],
      "child_modules": [
        {
          "resources": [
            {
              "address": "google_project.default",
              "mode": "managed",
              "type": "google_project",
              "name": "default",
              "provider_name": "google",
              "schema_version": 1,
              "values": {
                "auto_create_network": false,
                "billing_account": "002A75-ABB3C9-680D62",
                "folder_id": "746484843613",
                "id": "projects/di-test-blueprint-odrs-data-1",
                "labels": {
                  "ac-number": "us106088",
                  "env": "non-prod",
                  "financial-owner": "gpezzani",
                  "slb-org": "sdfc",
                  "ssr": "ear-aa-7542",
                  "technical-owner": "wdossantosjunior"
                },
                "name": "di-test-blueprint-odrs-data-1",
                "number": "1034955138102",
                "org_id": "",
                "project_id": "di-test-blueprint-odrs-data-1",
                "skip_delete": null,
                "timeouts": null
              }
            },
            {
              "address": "google_project_iam_audit_config.audit_config",
              "mode": "managed",
              "type": "google_project_iam_audit_config",
              "name": "audit_config",
              "provider_name": "google",
              "schema_version": 0,
              "values": {
                "audit_log_config": [
                  {
                    "exempted_members": [],
                    "log_type": "ADMIN_READ"
                  },
                  {
                    "exempted_members": [],
                    "log_type": "DATA_READ"
                  },
                  {
                    "exempted_members": [],
                    "log_type": "DATA_WRITE"
                  }
                ],
                "etag": "BwWsIPSdhKk=",
                "id": "di-test-blueprint-odrs-data-1/audit_config/allServices",
                "project": "di-test-blueprint-odrs-data-1",
                "service": "allServices"
              },
              "depends_on": [
                "module.project.google_project.default"
              ]
            },
            {
              "address": "google_project_iam_member.owner",
              "mode": "managed",
              "type": "google_project_iam_member",
              "name": "owner",
              "provider_name": "google",
              "schema_version": 0,
              "values": {
                "condition": [],
                "etag": "BwWsIPSdhKk=",
                "id": "di-test-blueprint-odrs-data-1/roles/owner/user:wdossantosjunior@slb.com",
                "member": "user:wdossantosjunior@slb.com",
                "project": "di-test-blueprint-odrs-data-1",
                "role": "roles/owner"
              },
              "depends_on": [
                "module.project.google_project.default"
              ]
            }
          ],
          "address": "module.project",
          "child_modules": [
            {
              "resources": [
                {
                  "address": "google_bigquery_dataset.dataset",
                  "mode": "managed",
                  "type": "google_bigquery_dataset",
                  "name": "dataset",
                  "provider_name": "google.dynamic",
                  "schema_version": 0,
                  "values": {
                    "access": [
                      {
                        "domain": "",
                        "group_by_email": "",
                        "role": "OWNER",
                        "special_group": "projectOwners",
                        "user_by_email": "",
                        "view": []
                      },
                      {
                        "domain": "",
                        "group_by_email": "",
                        "role": "READER",
                        "special_group": "projectReaders",
                        "user_by_email": "",
                        "view": []
                      },
                      {
                        "domain": "",
                        "group_by_email": "",
                        "role": "WRITER",
                        "special_group": "",
                        "user_by_email": "cloud-logs@system.gserviceaccount.com",
                        "view": []
                      }
                    ],
                    "creation_time": 1596632447640,
                    "dataset_id": "audit_logs",
                    "default_encryption_configuration": [],
                    "default_partition_expiration_ms": 0,
                    "default_table_expiration_ms": 0,
                    "delete_contents_on_destroy": false,
                    "description": "",
                    "etag": "oOFJzOdhpm7y2rejxTxBkw==",
                    "friendly_name": "",
                    "id": "projects/di-test-blueprint-odrs-data-1/datasets/audit_logs",
                    "labels": null,
                    "last_modified_time": 1596632447640,
                    "location": "US",
                    "project": "di-test-blueprint-odrs-data-1",
                    "self_link": "https://www.googleapis.com/bigquery/v2/projects/di-test-blueprint-odrs-data-1/datasets/audit_logs",
                    "timeouts": null
                  },
                  "depends_on": [
                    "module.project.google_project.default",
                    "module.project.module.log_exporter.google_project_service.bigquery"
                  ]
                },
                {
                  "address": "google_logging_project_sink.sink",
                  "mode": "managed",
                  "type": "google_logging_project_sink",
                  "name": "sink",
                  "provider_name": "google.dynamic",
                  "schema_version": 0,
                  "values": {
                    "bigquery_options": [],
                    "destination": "bigquery.googleapis.com/projects/di-test-blueprint-odrs-data-1/datasets/audit_logs",
                    "filter": "logName = (\n        \"projects/di-test-blueprint-odrs-data-1/logs/cloudaudit.googleapis.com%2Factivity\"\n        OR \"projects/di-test-blueprint-odrs-data-1/logs/cloudaudit.googleapis.com%2Fsystem_events\"\n        OR \"projects/di-test-blueprint-odrs-data-1/logs/cloudaudit.googleapis.com%2Fdata_access\"\n        OR \"projects/di-test-blueprint-odrs-data-1/logs/compute.googleapis.com%2Fvpc_flows\"\n      )",
                    "id": "projects/di-test-blueprint-odrs-data-1/sinks/audit_logs",
                    "name": "audit_logs",
                    "project": "di-test-blueprint-odrs-data-1",
                    "unique_writer_identity": false,
                    "writer_identity": "serviceAccount:cloud-logs@system.gserviceaccount.com"
                  },
                  "depends_on": [
                    "module.project.google_project.default",
                    "module.project.module.log_exporter.google_bigquery_dataset.dataset",
                    "module.project.module.log_exporter.google_project_service.bigquery"
                  ]
                },
                {
                  "address": "google_project_service.bigquery",
                  "mode": "managed",
                  "type": "google_project_service",
                  "name": "bigquery",
                  "provider_name": "google.dynamic",
                  "schema_version": 0,
                  "values": {
                    "disable_dependent_services": null,
                    "disable_on_destroy": false,
                    "id": "di-test-blueprint-odrs-data-1/bigquery.googleapis.com",
                    "project": "di-test-blueprint-odrs-data-1",
                    "service": "bigquery.googleapis.com",
                    "timeouts": null
                  },
                  "depends_on": [
                    "module.project.google_project.default"
                  ]
                }
              ],
              "address": "module.project.module.log_exporter"
            }
          ]
        }
      ]
    }
  }
}
`
