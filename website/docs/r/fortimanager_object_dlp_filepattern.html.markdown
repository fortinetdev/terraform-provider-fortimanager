---
subcategory: "Object DLP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_filepattern"
description: |-
  Configure file patterns used by DLP blocking.
---

# fortimanager_object_dlp_filepattern
Configure file patterns used by DLP blocking.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fortimanager_object_dlp_filepattern_entries`



## Example Usage

```hcl
resource "fortimanager_object_dlp_filepattern" "trname1" {
  comment = "This is a Terraform example"
  fosid   = 3
  name    = "terr-dlp-filepattern"
}



resource "fortimanager_object_dlp_filepattern" "trname2" {
  fosid = 2
  name  = "stefv1"

  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.bat"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.com"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.dll"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.doc"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.exe"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.gz"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.hta"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.ppt"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.rar"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.scr"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.tar"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.tgz"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.vb?"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.wps"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.xl?"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.zip"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.pif"
  }
  entries {
    file_type   = "unknown"
    filter_type = "pattern"
    pattern     = "*.cpl"
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `name` - Name of table containing the file pattern list.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `file_type` - Select a file type. Valid values: `unknown`, `ignored`, `exe`, `elf`, `bat`, `javascript`, `html`, `hta`, `msoffice`, `gzip`, `rar`, `tar`, `lzh`, `upx`, `zip`, `cab`, `bzip2`, `bzip`, `activemime`, `mime`, `hlp`, `arj`, `base64`, `binhex`, `uue`, `fsg`, `aspack`, `msc`, `petite`, `jpeg`, `gif`, `tiff`, `png`, `bmp`, `msi`, `mpeg`, `mov`, `mp3`, `wma`, `wav`, `pdf`, `avi`, `rm`, `torrent`, `hibun`, `7z`, `xz`, `msofficex`, `mach-o`, `dmg`, `.net`, `xar`, `chm`, `iso`, `crx`, `flac`, `sis`, `prc`, `class`, `jad`, `cod`.

* `filter_type` - Filter by file name pattern or by file type. Valid values: `pattern`, `type`.

* `pattern` - Add a file name pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectDlp Filepattern can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_filepattern.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
