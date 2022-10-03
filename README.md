# Terraform Provider for Tanzu Service Mesh

This will hopefully be a start to a terraform provider for VMware Tanzu Service Mesh.

This is in a pre-alpha state.  DO NOT USE IN PRODUCTION, unless you are ready to fork and maintain this code yourself.

# Terraform Provider Tanzu

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-tsm
```

## Local release build

```shell
$ go install github.com/goreleaser/goreleaser@latest
```

```shell
$ make release
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init -plugin-dir=~/.terraform.d
$ terraform init && terraform apply
```

# Unit testing
In vscode, you can right click and choose "Go: Test Function At Cursor" or run the tests manually.
```
go test terraform-provider-tsm/plugin/provider -v
```
