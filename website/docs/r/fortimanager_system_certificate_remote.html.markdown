---
subcategory: "System Certificate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_certificate_remote"
description: |-
  Remote certificate.
---

# fortimanager_system_certificate_remote
Remote certificate.

## Example Usage

```hcl
resource "fortimanager_system_certificate_remote" "trname" {
  cert    = ["-----BEGIN CERTIFICATE-----\nMIIDuTCCAqGgAwIBAgIJAKCr2aCM9Je5MA0GCSqGSIb3DQEBCwUAMHMxCzAJBgNV\nBAYTAkdCMRMwEQYDVQQIDApTb21lLVN0YXRlMQ8wDQYDVQQHDAZMb25kb24xDDAK\nBgNVBAoMA0ZETDESMBAGA1UEAwwJbG9jYWxob3N0MRwwGgYJKoZIhvcNAQkBFg1z\nc0Bzc3Nzc3MuY29tMB4XDTE5MDUyOTE1MDIzMFoXDTIwMDUyODE1MDIzMFowczEL\nMAkGA1UEBhMCR0IxEzARBgNVBAgMClNvbWUtU3RhdGUxDzANBgNVBAcMBkxvbmRv\nbjEMMAoGA1UECgwDRkRMMRIwEAYDVQQDDAlsb2NhbGhvc3QxHDAaBgkqhkiG9w0B\nCQEWDXNzQHNzc3Nzcy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\nAQDLtWP/2ozguJYmN34YxSKNKsgwkuagcNMso+JPgIBC+Cf4wOjUn/KCRUklLCJ9\nafyi4ZLJxqyOaAUkkiK7ay1DObgRDe7zl++prOijzTqkE5CbWBZgRjlSDfIyl113\nvFRA7yzrRUP9lDyk3OwpmVFZ6Uac3D1PNlcAGL8UHEsVDhM6fDnN+JtnV2e8xXeZ\n2DQxjaqh73rCmJJkq3iCRKbM/BFHtCdjMG0Xy9UuUnYjl+omjj+e32Pcp3ZfbBRp\nnX3hoiuu4IVToIXEBvZlMkvDj/3sRdHzCqgN+FTnmqTulRHbZRy5+2htZgIys7t7\nnVZ0FgRYiV0C1xmRSyTttRTNAgMBAAGjUDBOMB0GA1UdDgQWBBTinUGXSHLwDOVm\nOMdVbk0NdJNcRzAfBgNVHSMEGDAWgBTinUGXSHLwDOVmOMdVbk0NdJNcRzAMBgNV\nHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBX7ZHWH3N8EN+XcmUak7RG9qzy\nmnWPvyBiM8YI11rs87SkD+L8Q/ylxdmoI57cfPHpnqtkGAseRdN1EtzqILpyOo4Q\nQ2aF3ZHKUOEPBbblWqv+AbyXPHhODrm+Nlyu42axcqfLwLGAIRhVkVerX24lV/u2\ns3W/G5cse7RfNtxWPVUah7jAmsIv1Yc7Yi4TEIlQDImQI5TAoGTQOpPjYZXCtHXS\nxUIGOKDTds9X2wWb3lM7ANecrINh6CNB/tg3GNdGV8SCJvJnYtEgfipjS7cQoc5C\npBmnz+IlqzrwBZBxmB+1xFrYATm/hZ3HMFrLKQVoTJgTP74/PIpCaO/mjis4\n-----END CERTIFICATE-----"]
  comment = "This is a Terraform example"
  name    = "terr-sys-cer-remote"
}
```

## Argument Reference


The following arguments are supported:


* `cert` - Remote certificate.
* `comment` - Remote certificate comment.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateRemote can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_certificate_remote.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

