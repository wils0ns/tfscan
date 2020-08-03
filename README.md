# tfscan

Inspect Terraform resources in a state and plan JSON files

## Install

Using go:

```bash
git clone https://github.com/wilson-codeminus/tfscan.git
cd tfscan
go install
```

Or downloading the binary for a particular [release](https://github.com/wilson-codeminus/tfscan/releases).

## Use

Reading from `terraform` stdout:

```bash
terraform show -json | tfscan
```

Passing a JSON file:

```bash
tfscan -json state.json

```
