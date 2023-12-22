---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_filepattern_entries"
description: |-
  Configure file patterns used by DLP blocking.
---

# fortimanager_object_dlp_filepattern_entries
Configure file patterns used by DLP blocking.

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_dlp_filepattern`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_dlp_filepattern_entries" "trname" {
  filepattern = fortimanager_object_dlp_filepattern.trname2.fosid
  file_type   = "unknown"
  filter_type = "pattern"
  pattern     = "*.bat"
  depends_on  = [fortimanager_object_dlp_filepattern.trname2]
}

resource "fortimanager_object_dlp_filepattern" "trname2" {
  name  = "terr-dlp-filepattern"
  fosid = 5
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `filepattern` - Filepattern.

* `file_type` - Select a file type. Valid values: `unknown`, `ignored`, `exe`, `elf`, `bat`, `javascript`, `html`, `hta`, `msoffice`, `gzip`, `rar`, `tar`, `lzh`, `upx`, `zip`, `cab`, `bzip2`, `bzip`, `activemime`, `mime`, `hlp`, `arj`, `base64`, `binhex`, `uue`, `fsg`, `aspack`, `msc`, `petite`, `jpeg`, `gif`, `tiff`, `png`, `bmp`, `msi`, `mpeg`, `mov`, `mp3`, `wma`, `wav`, `pdf`, `avi`, `rm`, `torrent`, `hibun`, `7z`, `xz`, `msofficex`, `mach-o`, `dmg`, `.net`, `xar`, `chm`, `iso`, `crx`, `flac`, `sis`, `prc`, `class`, `jad`, `cod`.

* `filter_type` - Filter by file name pattern or by file type. Valid values: `pattern`, `type`.

* `pattern` - Add a file name pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{pattern}}.

## Import

ObjectDlp FilepatternEntries can be imported using any of these accepted formats:
```
Set import_options = ["filepattern=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_filepattern_entries.labelname {{pattern}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
