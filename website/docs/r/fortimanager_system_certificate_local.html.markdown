---
subcategory: "System Certificate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_certificate_local"
description: |-
  Local keys and certificates.
---

# fortimanager_system_certificate_local
Local keys and certificates.

## Example Usage

```hcl
resource "fortimanager_system_certificate_local" "trname" {
  certificate = ["-----BEGIN CERTIFICATE-----\nMIIDuTCCAqGgAwIBAgIJAKCr2aCM9Je5MA0GCSqGSIb3DQEBCwUAMHMxCzAJBgNV\nBAYTAkdCMRMwEQYDVQQIDApTb21lLVN0YXRlMQ8wDQYDVQQHDAZMb25kb24xDDAK\nBgNVBAoMA0ZETDESMBAGA1UEAwwJbG9jYWxob3N0MRwwGgYJKoZIhvcNAQkBFg1z\nc0Bzc3Nzc3MuY29tMB4XDTE5MDUyOTE1MDIzMFoXDTIwMDUyODE1MDIzMFowczEL\nMAkGA1UEBhMCR0IxEzARBgNVBAgMClNvbWUtU3RhdGUxDzANBgNVBAcMBkxvbmRv\nbjEMMAoGA1UECgwDRkRMMRIwEAYDVQQDDAlsb2NhbGhvc3QxHDAaBgkqhkiG9w0B\nCQEWDXNzQHNzc3Nzcy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\nAQDLtWP/2ozguJYmN34YxSKNKsgwkuagcNMso+JPgIBC+Cf4wOjUn/KCRUklLCJ9\nafyi4ZLJxqyOaAUkkiK7ay1DObgRDe7zl++prOijzTqkE5CbWBZgRjlSDfIyl113\nvFRA7yzrRUP9lDyk3OwpmVFZ6Uac3D1PNlcAGL8UHEsVDhM6fDnN+JtnV2e8xXeZ\n2DQxjaqh73rCmJJkq3iCRKbM/BFHtCdjMG0Xy9UuUnYjl+omjj+e32Pcp3ZfbBRp\nnX3hoiuu4IVToIXEBvZlMkvDj/3sRdHzCqgN+FTnmqTulRHbZRy5+2htZgIys7t7\nnVZ0FgRYiV0C1xmRSyTttRTNAgMBAAGjUDBOMB0GA1UdDgQWBBTinUGXSHLwDOVm\nOMdVbk0NdJNcRzAfBgNVHSMEGDAWgBTinUGXSHLwDOVmOMdVbk0NdJNcRzAMBgNV\nHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBX7ZHWH3N8EN+XcmUak7RG9qzy\nmnWPvyBiM8YI11rs87SkD+L8Q/ylxdmoI57cfPHpnqtkGAseRdN1EtzqILpyOo4Q\nQ2aF3ZHKUOEPBbblWqv+AbyXPHhODrm+Nlyu42axcqfLwLGAIRhVkVerX24lV/u2\ns3W/G5cse7RfNtxWPVUah7jAmsIv1Yc7Yi4TEIlQDImQI5TAoGTQOpPjYZXCtHXS\nxUIGOKDTds9X2wWb3lM7ANecrINh6CNB/tg3GNdGV8SCJvJnYtEgfipjS7cQoc5C\npBmnz+IlqzrwBZBxmB+1xFrYATm/hZ3HMFrLKQVoTJgTP74/PIpCaO/mjis4\n-----END CERTIFICATE-----"]
  comment     = "This is a Terraform example"
  name        = "terr-sys-cer-local"
  password    = ["fortinet"]
  private_key = ["-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAy7Vj/9qM4LiWJjd+GMUijSrIMJLmoHDTLKPiT4CAQvgn+MDo\n1J/ygkVJJSwifWn8ouGSycasjmgFJJIiu2stQzm4EQ3u85fvqazoo806pBOQm1gW\nYEY5Ug3yMpddd7xUQO8s60VD/ZQ8pNzsKZlRWelGnNw9TzZXABi/FBxLFQ4TOnw5\nzfibZ1dnvMV3mdg0MY2qoe96wpiSZKt4gkSmzPwRR7QnYzBtF8vVLlJ2I5fqJo4/\nnt9j3Kd2X2wUaZ194aIrruCFU6CFxAb2ZTJLw4/97EXR8wqoDfhU55qk7pUR22Uc\nuftobWYCMrO7e51WdBYEWIldAtcZkUsk7bUUzQIDAQABAoIBABlV8x0EOpdMfeg8\n6KL+CcES/BkGfEaiIbGgpGoM6mbp5FbM72haiFfpdCJ6bcO5ZeGAOrh7zERd7Z3R\nyx4SQ2vkBt+gIwMK95Tb24db5Bo6ELcxan8I3OI2t9PQ/aABvVziIm0UjVNBl5VN\noNW/qt2K5Oxne/yZHpL1gPZoWnJAuBNcDZDNI5qQfT1JTWhmbu9pkjiNg3h0l5O4\nboETBdb3W2jlvCIegIQJ+xPkChS3I4cMoZ4LBRWMLpzK+Q57+zml75drsQ7PA0XH\nlx2nzUFCByu27pM6kiajXleUSGVH2VCUSNysQAYBSa77g5t7O+m+o3iNUslUDor3\nLY5eiKUCgYEA6dqJJxF28Wt6UbMgywQuv5koo8v8nyUhR4hZhvy5qge5v5Sh82UE\nRyVfSvMDR9oWnXs8tBZaf1sgsUEaFl2I/5kmFWTe7aGj1eXprOxOuMNk51AN2w9T\njHWici/rj0JEjvAteDvLjY6CTAi8Zg9OxuJWNyV2gZ1LpZ2cIlULzLsCgYEA3wAH\nJQ0jVeJ70v2I01m16d/klTNcqv9sTIPwowz0RFkOG8v482SSQ7Q43xkSYJvxKg12\nBO6qq+RzYwDPgA7/7kLrq1Ye2VobhrsRey6dEWGdrgA/TfoCgSjK0LEBh4Gn5h7l\nDycvfNRbum1D+uheyTPC94fJChwutihUsrXuEBcCgYAaoUczCrsXvNx+Bz75v20v\nZlqJZIZM/SZwBefkBk2CPkT5uwxCMkOtcmUKnOfHu98NaeY8v7roe9EaPkahO1+J\nc8AxeX4lY13L0tWsWnCQe7e225foVTN3cEHibPCPLMWv3UvgQDbq1Mqjq+8AVEft\nQAL/XqXDFs1xe6Q3CKZCVwKBgQCaCBbnTM/ffvUwo9dixVCWHwRw2m1j39Iad/g7\nZ7NBkpHgOV/YHtu40D+IOnUrLgvClFG0znYtDTt2YxTwy2uUU70dN/tO/qKMyaIl\nh+kOHHMhwSH45nvcYyTUSa9YvgIPPb/SW6q9eqFxgA+4u9DdAVfmSnBe/2B0ih8W\n4ftyOQKBgF0puFMyA7ySE2tQ5quiPBO95Vls4SMl59ofhEgMghmUErEFGvTHPxff\n2UF7AWahrhNq02cF8iTU/lS77D0W01TpEFl8xC5LyqewKzLatgFTFhFPPGt/wK0s\nuIAliRuV1Iyv2vDYmYaugpeZakogVBpkVPqvEzfEPgn9VEZKLQ7y\n-----END RSA PRIVATE KEY-----"]
}
```

## Argument Reference


The following arguments are supported:


* `certificate` - Certificate.
* `comment` - Local certificate comment.
* `csr` - Certificate Signing Request.
* `name` - Name of local certificate.
* `password` - Local certificate password.
* `private_key` - Local certificate private-key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateLocal can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_certificate_local.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

