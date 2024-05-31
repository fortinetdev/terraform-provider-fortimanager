# FortiManager Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://www.datocms-assets.com/2885/1629941242-logo-terraform-main.svg" width="600px">

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x +
- [Go](https://golang.org/doc/install) 1.21.x (to build the provider plugin)
- The provider can cover FortiManager 6.0, 6.2, 6.4, 7.0 versions.

## Building the Provider

1. Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-fortimanager`.

    ```sh
    $ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
    $ git clone git@github.com:fortinetdev/terraform-provider-fortimanager
    ```

2. Enter the provider directory and build the provider.

    ```sh
    $ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-fortimanager
    $ make build
    ```

## Using the Provider

If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory,  run `terraform init` to initialize it.

```sh
$ terraform init
```

## Developing the Provider

If you wish to work on the provider, you'll first need Go installed on your machine (version 1.13+ is required). You'll also need to correctly setup a GOPATH, as well as adding $GOPATH/bin to your $PATH.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the $GOPATH/bin directory.

```sh
$ make build
...
```
