# Terraform Provider Tanzu

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-tanzu
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

