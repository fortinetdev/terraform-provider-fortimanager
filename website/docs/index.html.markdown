---
layout: "fortimanager"
page_title: "Provider: FortiManager"
sidebar_current: "docs-fortimanager-index"
description: |-
  The FortiManager provider interacts with FortiManager.
---

# FortiManager Provider

The FortiManager provider is used to interact with the resources supported by FortiManager. We need to configure the provider with the proper credentials before it can be used. Please use the navigation on the left to read more details about the available resources.

## Example Usage

```hcl
# Configure the Provider for FortiManager
provider "fortimanager" {
  hostname     = "192.168.52.178"
  username     = "admin"
  password     = "admin"
  insecure     = "false"
  cabundlefile = "/path/yourCA.crt"

  scopetype    = "adom"
  adom         = "root"
}

# Create a firewall vip object
resource "fortimanager_object_firewall_vip" "trname1" {
  scopetype = "inherit"
  adom      = "root"
  extintf   = "any"
  extip     = "1.1.1.1-2.1.1.1"
  mappedip  = ["12.1.1.1-13.1.1.1"]
  name      = "viptest"
}

```

Before using this provider, the permission level for rpc-permit need to be set. See `Guides->To Set the Permission Level for RPC-Permit` for details.

If it is used for testing, you can set `insecure` to "true" and unset `cabundlefile` to quickly set the provider up, for example:

```hcl
provider "fortimanager" {
  hostname     = "192.168.52.178"
  username     = "admin"
  password     = "admin"
  insecure     = "true"

  scopetype    = "adom"
  adom         = "root"
}
```

Please refer to the Argument Reference below for more help on `insecure` and `cabundlefile`.


## Authentication

The FortiManager provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by `username` and `password` parameters in the FortiManager provider block.

Usage:

```hcl
provider "fortimanager" {
  hostname     = "192.168.52.178"
  username     = "admin"
  password     = "admin"
  insecure     = "true"

  scopetype    = "adom"
  adom         = "root"
}
```

### Environment variables

You can provide your credentials via the `FORTIMANAGER_ACCESS_HOSTNAME`, `FORTIMANAGER_ACCESS_USERNAME`, `FORTIMANAGER_ACCESS_PASSWORD`, `FORTIMANAGER_INSECURE` and `FORTIMANAGER_CA_CABUNDLE` environment variables. Note that setting your FortiManager credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIMANAGER_ACCESS_HOSTNAME"="192.168.52.178"
$ export "FORTIMANAGER_ACCESS_USERNAME"="admin"
$ export "FORTIMANAGER_ACCESS_PASSWORD"="admin"
$ export "FORTIMANAGER_INSECURE"="false"
$ export "FORTIMANAGER_CA_CABUNDLE"="/path/yourCA.crt"
```

Then configure the FortiManager Provider as following:

```hcl
provider "fortimanager" {
  scopetype = "adom"
  adom      = "root"
}
```

## Argument Reference

The following arguments are supported:

* `hostname` - (Optional) The hostname or IP address of FortiManager unit. It must be provided, but it can also be sourced from the `FORTIMANAGER_ACCESS_HOSTNAME` environment variable.

* `username` - (Optional) Your username. It must be provided, but it can also be sourced from the `FORTIMANAGER_ACCESS_USERNAME` environment variable.

* `password` - (Optional) Your password. It must be provided, but it can also be sourced from the `FORTIMANAGER_ACCESS_PASSWORD` environment variable.

* `insecure` - (Optional) Control whether the Provider to perform insecure SSL requests. If omitted, the `FORTIMANAGER_INSECURE` environment variable is used. If neither is set, default value is `false`.

* `cabundlefile` - (Optional) The path of a custom CA bundle file. You can specify a path to the file, or you can specify it by the `FORTIMANAGER_CA_CABUNDLE` environment variable.

* `scopetype` - (Optional) The option is used to set the default scope of application of those resources managed by the provider. Valid values: `adom`, `global`. The default value is `adom`. Each resource can also set its own scope as needed, see the description of each resource for details.

* `adom` - (Optional) Adom. This value is valid only when the `scopetype` is set to `adom`. The option is used to set the default adom of the resources managed by the provider. The default value is `root`. Each resource can also set its own adom as needed, see the description of each resource for details.

* `import_options` - (Optional) This parameter is only used for import in some special cases. When the resource to be imported includes pkg parameter, you need to assign a value to the parameter here, for example:

    ```hcl
    provider "fortimanager" {
      hostname       = "192.168.52.178"
      username       = "admin"
      password       = "admin"
      insecure       = "true"

      scopetype      = "adom"
      adom           = "root"

      import_options = ["pkg=default"]
    }
    ```

* `logsession` - (Optional) Save the session to a local file. Used to assist fortimanager_exec_workspace_action resource. Valid values: `true`: log to file, `false`: do not log to file. Default is `false`. See `Guides -> To Lock for Restricting Configuration Changes` for details.

* `presession` - (Optional) The session saved earlier and within the validity period, used to reuse the previous session and assist fortimanager_exec_workspace_action resource. See `Guides -> To Lock for Restricting Configuration Changes` for details. Default is empty.


## Release
Check out the FortiManager provider release notes and additional information from: [the FortiManager provider releases](https://github.com/fortinetdev/terraform-provider-fortimanager/releases).


## FortiManager best practices

FortiManager is an integrated platform for the centralized management of products in a Fortinet security infrastructure, including FortiGates.

Once FortiGates are managed by a FortiManager that is operating in Normal Mode, whenever possible, configuration changes should be made on the FortiManager and not the FortiGate. This recommendation also applies when the configuration of FortiGates and FortiManager is executed through the [FortiOS](https://registry.terraform.io/providers/fortinetdev/fortios/latest) and [FortiManager](https://registry.terraform.io/providers/fortinetdev/fortimanager/latest) providers.

To help you get the most out of your FortiManager products, maximize performance, and avoid potential problems, please refer to the [FortiManager documentation](https://docs.fortinet.com/product/fortimanager), including the Administration Guide and Best Practices documents.

Fortinet also provides a developer community to help administrators and advanced users enhance and increase the effectiveness of Fortinet products. The [Fortinet Developer Network (FNDN)](https://fndn.fortinet.net/) provides the official documentation and advanced tools for developing custom solutions using Fortinet products.


## Versioning

The provider can cover FortiManager 6.4 to 7.2 versions, the configuration of all parameters should be based on the relevant FortiManager version manual and FortiManager API guide.
