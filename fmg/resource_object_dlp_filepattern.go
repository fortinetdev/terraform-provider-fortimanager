// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure file patterns used by DLP blocking.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectDlpFilepattern() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDlpFilepatternCreate,
		Read:   resourceObjectDlpFilepatternRead,
		Update: resourceObjectDlpFilepatternUpdate,
		Delete: resourceObjectDlpFilepatternDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"filter_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectDlpFilepatternCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDlpFilepattern(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpFilepattern resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDlpFilepattern(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDlpFilepattern resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpFilepatternRead(d, m)
}

func resourceObjectDlpFilepatternUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectDlpFilepattern(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpFilepattern resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDlpFilepattern(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDlpFilepattern resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDlpFilepatternRead(d, m)
}

func resourceObjectDlpFilepatternDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectDlpFilepattern(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDlpFilepattern resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDlpFilepatternRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectDlpFilepattern(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpFilepattern resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDlpFilepattern(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDlpFilepattern resource from API: %v", err)
	}
	return nil
}

func flattenObjectDlpFilepatternComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpFilepatternEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			v := flattenObjectDlpFilepatternEntriesFileType(i["file-type"], d, pre_append)
			tmp["file_type"] = fortiAPISubPartPatch(v, "ObjectDlpFilepattern-Entries-FileType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {
			v := flattenObjectDlpFilepatternEntriesFilterType(i["filter-type"], d, pre_append)
			tmp["filter_type"] = fortiAPISubPartPatch(v, "ObjectDlpFilepattern-Entries-FilterType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenObjectDlpFilepatternEntriesPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "ObjectDlpFilepattern-Entries-Pattern")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectDlpFilepatternEntriesFileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:    "unknown",
			2:    "exe",
			4:    "elf",
			5:    "bat",
			7:    "javascript",
			8:    "html",
			9:    "hta",
			10:   "msoffice",
			11:   "gzip",
			12:   "rar",
			13:   "tar",
			14:   "lzh",
			15:   "upx",
			16:   "zip",
			17:   "cab",
			18:   "bzip2",
			19:   "bzip",
			20:   "activemime",
			21:   "mime",
			22:   "hlp",
			23:   "arj",
			24:   "base64",
			25:   "binhex",
			26:   "uue",
			27:   "fsg",
			28:   "aspack",
			32:   "msc",
			33:   "petite",
			39:   "jpeg",
			40:   "gif",
			41:   "tiff",
			42:   "png",
			43:   "bmp",
			44:   "msi",
			45:   "mpeg",
			46:   "mov",
			47:   "mp3",
			48:   "wma",
			49:   "wav",
			50:   "pdf",
			51:   "avi",
			52:   "rm",
			53:   "torrent",
			54:   "hibun",
			55:   "7z",
			56:   "xz",
			57:   "msofficex",
			58:   "mach-o",
			59:   "dmg",
			60:   ".net",
			61:   "xar",
			62:   "chm",
			63:   "iso",
			64:   "crx",
			65:   "flac",
			1000: "sis",
			1003: "class",
			1004: "jad",
			1005: "cod",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDlpFilepatternEntriesFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pattern",
			1: "type",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectDlpFilepatternEntriesPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpFilepatternId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDlpFilepatternName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDlpFilepattern(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenObjectDlpFilepatternComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectDlpFilepattern-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenObjectDlpFilepatternEntries(o["entries"], d, "entries")); err != nil {
			if vv, ok := fortiAPIPatch(o["entries"], "ObjectDlpFilepattern-Entries"); ok {
				if err = d.Set("entries", vv); err != nil {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenObjectDlpFilepatternEntries(o["entries"], d, "entries")); err != nil {
				if vv, ok := fortiAPIPatch(o["entries"], "ObjectDlpFilepattern-Entries"); ok {
					if err = d.Set("entries", vv); err != nil {
						return fmt.Errorf("Error reading entries: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("fosid", flattenObjectDlpFilepatternId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectDlpFilepattern-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectDlpFilepatternName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectDlpFilepattern-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectDlpFilepatternFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDlpFilepatternComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpFilepatternEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["file-type"], _ = expandObjectDlpFilepatternEntriesFileType(d, i["file_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-type"], _ = expandObjectDlpFilepatternEntriesFilterType(d, i["filter_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandObjectDlpFilepatternEntriesPattern(d, i["pattern"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectDlpFilepatternEntriesFileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpFilepatternEntriesFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpFilepatternEntriesPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpFilepatternId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDlpFilepatternName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDlpFilepattern(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectDlpFilepatternComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {
		t, err := expandObjectDlpFilepatternEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandObjectDlpFilepatternId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectDlpFilepatternName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
