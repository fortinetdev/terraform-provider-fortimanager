---
subcategory: "Generic"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_exec_workspace_action"
description: |-
  FortiManager API Generic Interface.

---

# fortimanager_json_generic_api

FortiManager API Generic Interface.

## Example Usage

```hcl
provider "fortimanager" {
  hostname = "192.168.52.178"
  username = "admin"
  password = "admin"
  insecure = "true"

  scopetype = "adom"
  adom      = "root"
}

resource "fortimanager_json_generic_api" "test1" {
  json_content = <<JSON
{
    "method": "add",
    "params": [
        {
            "url": "/pm/config/adom/root/pkg/default/firewall/policy",
            "data": {
                "srcintf": [
                    "any"
                ],
                "dstintf": [
                    "any"
                ],
                "srcaddr": [
                    "all"
                ],
                "dstaddr": [
                    "all"
                ],
                "action": "accept",
                "status": "enable",
                "schedule": [
                    "always"
                ],
                "service": [
                    "ALL"
                ],
                "name": "sss1"
            }
        }
    ]
}
JSON
}

output name1 {
  value       = jsondecode(fortimanager_json_generic_api.test1.response)
}

resource "fortimanager_json_generic_api" "test2" {
  json_content = <<JSON
{
    "method": "get",
    "params": [
        {
            "url": "/pm/config/global/pkg/default/global/footer/consolidated/policy"
        }
    ]
}
JSON
}

output name2 {
  value       = jsondecode(fortimanager_json_generic_api.test2.response)
}

```

## Argument Reference


The following arguments are supported:

* `force_recreate` - The argument is optional, if it is set, when its value changes, the resource will be re-created. It is usually used when the return value needs to be forced to update.
* `json_content` - Body data in JSON format.
* `comment` - Comment.
* `response` - API returns results.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - The resource id.
