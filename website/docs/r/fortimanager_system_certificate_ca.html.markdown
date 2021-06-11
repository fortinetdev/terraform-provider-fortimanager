---
subcategory: "System Certificate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_certificate_ca"
description: |-
  CA certificate.
---

# fortimanager_system_certificate_ca
CA certificate.

## Example Usage

```hcl
resource "fortimanager_system_certificate_ca" "trname" {
  ca      = ["\"-----BEGIN CERTIFICATE-----\nMIID1TCCAr2gAwIBAgIJANr2NrRD1KWLMA0GCSqGSIb3DQEBCwUAMIGgMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTESMBAGA1UEBxMJU3Vubnl2YWxl\nMREwDwYDVQQKEwhGb3J0aW5ldDEeMBwGA1UECxMVQ2VydGlmaWNhdGUgQXV0aG9y\naXR5MRAwDgYDVQQDEwdzdXBwb3J0MSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0QGZv\ncnRpbmV0LmNvbTAeFw0xNTA3MTYyMjM0MzlaFw0zODAxMTkyMjM0MzlaMIGgMQsw\nCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTESMBAGA1UEBxMJU3Vubnl2\nYWxlMREwDwYDVQQKEwhGb3J0aW5ldDEeMBwGA1UECxMVQ2VydGlmaWNhdGUgQXV0\naG9yaXR5MRAwDgYDVQQDEwdzdXBwb3J0MSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0\nQGZvcnRpbmV0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANQ7\nUXPQNRISyMNOWUFI0TvBsKbVD4MhKYt4v+vDRJ3GZoKa7gQfoJ+r4M14xWzTNN26\nNrJtUutYSUmDmX/m5rbo24JjtGdx1FzKL6+3cDKHUDV8wZCI0DQZLamhgbNJ7fv0\nVergsuJXxp8urzFJA25pbqxH9X6u1lOs4NqCM6CTHvVp+IqmZGucVmFQTMh9fiE0\nL20IFY0K5wc/C3XP1Pa0CzXLu6smjtr0prwXisGSiTkxg0+HTSO24tv0q5ABPgfz\nOWKpX9b6gaMt17r3hSi5GhKElbCDGLtMbdKcldboxNBQZ5nxPRNFS26Lde5duB8j\noc6Rwdgv9dsxSS17HWMCAwEAAaMQMA4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B\nAQsFAAOCAQEAhxf7jexnSrTNshppXpiMmlK5C9HxuQQZQT/dq0FkmN/7cp6rtgl3\n6WUfBU95ia0q9GJnaxbSWpXO4jq4wT5vlGCIEaiajJLJx1BFRu7dMUZnkmoyIvKT\n893YUGlAZeGME/WXVBaAHsj4wIhYIOMvUm7VB3wUK9SjZtn0qGEz+OjxmWjnJ2+c\nk67pZbRLl8gOuRB6OkFIiNp8slO08g4lXIYquHIltC7RCadNQPz3qfYvW4npIW49\n04JH4CUAtzvbRdDs4MCvszKuPOICvm6pnqin1AIqJ+4cMwPrVVyBN5NsIhohAUjS\nELtlWJR1qQ7HT5Q9tVLCr/qOnEHEkoy4hw==\n-----END CERTIFICATE-----\""]
  comment = "terraform-comment"
  name    = "terraform-tefv-ca"
}
```

## Argument Reference


The following arguments are supported:


* `ca` - CA certificate.
* `comment` - CA certificate comment.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateCa can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_certificate_ca.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

