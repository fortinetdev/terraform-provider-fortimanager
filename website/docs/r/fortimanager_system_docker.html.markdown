---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_docker"
description: |-
  Docker host.
---

# fortimanager_system_docker
Docker host.

## Example Usage

```hcl
resource "fortimanager_system_docker" "trname" {
  cpu = 0
  default_address_pool_base = [
    "172.17.0.0",
    "255.255.0.0",
  ]
  default_address_pool_size = 24
  docker_user_login_max     = 0
  fortiauthenticator        = "disable"
  fortiportal               = "disable"
  fortisigconverter         = "disable"
  fortiwlm                  = "disable"
  mem                       = 0
  sdwancontroller           = "disable"
  status                    = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `cpu` - Max CPU usage.
* `default_address_pool_base` - Set default-address-pool CIDR.
* `default_address_pool_size` - Set default-address-pool size.
* `docker_user_login_max` - Max login session for docker users.
* `fortiauthenticator` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortiportal` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortisigconverter` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortisoar` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortiwlm` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `mem` - Max % RAM usage.
* `sdwancontroller` - Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `status` - Enable and set registry. disable - Disable docker host service. enable - Enable production registry. qa - Enable QA test registry. dev - Enable QA test registry (without signature). Valid values: `disable`, `enable`, `qa`, `dev`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Docker can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_docker.labelname SystemDocker
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

