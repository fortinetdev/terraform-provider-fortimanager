// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure web proxy address.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallProxyAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallProxyAddressCreate,
		Read:   resourceObjectFirewallProxyAddressRead,
		Update: resourceObjectFirewallProxyAddressUpdate,
		Delete: resourceObjectFirewallProxyAddressDelete,

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
			"_image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"header_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"referrer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ua": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectFirewallProxyAddressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallProxyAddress(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallProxyAddress resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallProxyAddress(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallProxyAddress resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallProxyAddressRead(d, m)
}

func resourceObjectFirewallProxyAddressUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallProxyAddress(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProxyAddress resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallProxyAddress(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProxyAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallProxyAddressRead(d, m)
}

func resourceObjectFirewallProxyAddressDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallProxyAddress(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallProxyAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallProxyAddressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallProxyAddress(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProxyAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallProxyAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProxyAddress resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallProxyAddressImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressHeaderGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			v := flattenObjectFirewallProxyAddressHeaderGroupCaseSensitivity(i["case-sensitivity"], d, pre_append)
			tmp["case_sensitivity"] = fortiAPISubPartPatch(v, "ObjectFirewallProxyAddress-HeaderGroup-CaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenObjectFirewallProxyAddressHeaderGroupHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "ObjectFirewallProxyAddress-HeaderGroup-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenObjectFirewallProxyAddressHeaderGroupHeaderName(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "ObjectFirewallProxyAddress-HeaderGroup-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallProxyAddressHeaderGroupId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallProxyAddress-HeaderGroup-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallProxyAddressHeaderGroupCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressHeaderGroupHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressHeaderGroupHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressHeaderGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressHostRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProxyAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressReferrer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			v := flattenObjectFirewallProxyAddressTaggingCategory(i["category"], d, pre_append)
			tmp["category"] = fortiAPISubPartPatch(v, "ObjectFirewallProxyAddress-Tagging-Category")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectFirewallProxyAddressTaggingName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectFirewallProxyAddress-Tagging-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			v := flattenObjectFirewallProxyAddressTaggingTags(i["tags"], d, pre_append)
			tmp["tags"] = fortiAPISubPartPatch(v, "ObjectFirewallProxyAddress-Tagging-Tags")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallProxyAddressTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressTaggingTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProxyAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressUa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProxyAddressUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProxyAddressVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallProxyAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("_image_base64", flattenObjectFirewallProxyAddressImageBase64(o["_image-base64"], d, "_image_base64")); err != nil {
		if vv, ok := fortiAPIPatch(o["_image-base64"], "ObjectFirewallProxyAddress-ImageBase64"); ok {
			if err = d.Set("_image_base64", vv); err != nil {
				return fmt.Errorf("Error reading _image_base64: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _image_base64: %v", err)
		}
	}

	if err = d.Set("case_sensitivity", flattenObjectFirewallProxyAddressCaseSensitivity(o["case-sensitivity"], d, "case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitivity"], "ObjectFirewallProxyAddress-CaseSensitivity"); ok {
			if err = d.Set("case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitivity: %v", err)
		}
	}

	if err = d.Set("category", flattenObjectFirewallProxyAddressCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ObjectFirewallProxyAddress-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("color", flattenObjectFirewallProxyAddressColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallProxyAddress-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectFirewallProxyAddressComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectFirewallProxyAddress-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("header", flattenObjectFirewallProxyAddressHeader(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "ObjectFirewallProxyAddress-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("header_group", flattenObjectFirewallProxyAddressHeaderGroup(o["header-group"], d, "header_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["header-group"], "ObjectFirewallProxyAddress-HeaderGroup"); ok {
				if err = d.Set("header_group", vv); err != nil {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading header_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header_group"); ok {
			if err = d.Set("header_group", flattenObjectFirewallProxyAddressHeaderGroup(o["header-group"], d, "header_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["header-group"], "ObjectFirewallProxyAddress-HeaderGroup"); ok {
					if err = d.Set("header_group", vv); err != nil {
						return fmt.Errorf("Error reading header_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("header_name", flattenObjectFirewallProxyAddressHeaderName(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "ObjectFirewallProxyAddress-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("host", flattenObjectFirewallProxyAddressHost(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ObjectFirewallProxyAddress-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_regex", flattenObjectFirewallProxyAddressHostRegex(o["host-regex"], d, "host_regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-regex"], "ObjectFirewallProxyAddress-HostRegex"); ok {
			if err = d.Set("host_regex", vv); err != nil {
				return fmt.Errorf("Error reading host_regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_regex: %v", err)
		}
	}

	if err = d.Set("method", flattenObjectFirewallProxyAddressMethod(o["method"], d, "method")); err != nil {
		if vv, ok := fortiAPIPatch(o["method"], "ObjectFirewallProxyAddress-Method"); ok {
			if err = d.Set("method", vv); err != nil {
				return fmt.Errorf("Error reading method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallProxyAddressName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallProxyAddress-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("path", flattenObjectFirewallProxyAddressPath(o["path"], d, "path")); err != nil {
		if vv, ok := fortiAPIPatch(o["path"], "ObjectFirewallProxyAddress-Path"); ok {
			if err = d.Set("path", vv); err != nil {
				return fmt.Errorf("Error reading path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading path: %v", err)
		}
	}

	if err = d.Set("query", flattenObjectFirewallProxyAddressQuery(o["query"], d, "query")); err != nil {
		if vv, ok := fortiAPIPatch(o["query"], "ObjectFirewallProxyAddress-Query"); ok {
			if err = d.Set("query", vv); err != nil {
				return fmt.Errorf("Error reading query: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query: %v", err)
		}
	}

	if err = d.Set("referrer", flattenObjectFirewallProxyAddressReferrer(o["referrer"], d, "referrer")); err != nil {
		if vv, ok := fortiAPIPatch(o["referrer"], "ObjectFirewallProxyAddress-Referrer"); ok {
			if err = d.Set("referrer", vv); err != nil {
				return fmt.Errorf("Error reading referrer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading referrer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenObjectFirewallProxyAddressTagging(o["tagging"], d, "tagging")); err != nil {
			if vv, ok := fortiAPIPatch(o["tagging"], "ObjectFirewallProxyAddress-Tagging"); ok {
				if err = d.Set("tagging", vv); err != nil {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenObjectFirewallProxyAddressTagging(o["tagging"], d, "tagging")); err != nil {
				if vv, ok := fortiAPIPatch(o["tagging"], "ObjectFirewallProxyAddress-Tagging"); ok {
					if err = d.Set("tagging", vv); err != nil {
						return fmt.Errorf("Error reading tagging: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenObjectFirewallProxyAddressType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallProxyAddress-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ua", flattenObjectFirewallProxyAddressUa(o["ua"], d, "ua")); err != nil {
		if vv, ok := fortiAPIPatch(o["ua"], "ObjectFirewallProxyAddress-Ua"); ok {
			if err = d.Set("ua", vv); err != nil {
				return fmt.Errorf("Error reading ua: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ua: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallProxyAddressUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallProxyAddress-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("visibility", flattenObjectFirewallProxyAddressVisibility(o["visibility"], d, "visibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["visibility"], "ObjectFirewallProxyAddress-Visibility"); ok {
			if err = d.Set("visibility", vv); err != nil {
				return fmt.Errorf("Error reading visibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallProxyAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallProxyAddressImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressHeaderGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitivity"], _ = expandObjectFirewallProxyAddressHeaderGroupCaseSensitivity(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandObjectFirewallProxyAddressHeaderGroupHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandObjectFirewallProxyAddressHeaderGroupHeaderName(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallProxyAddressHeaderGroupId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallProxyAddressHeaderGroupCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressHeaderGroupHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressHeaderGroupHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressHeaderGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressHeaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressHostRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProxyAddressName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressQuery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressReferrer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category"], _ = expandObjectFirewallProxyAddressTaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectFirewallProxyAddressTaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tags"], _ = expandObjectFirewallProxyAddressTaggingTags(d, i["tags"], pre_append)
		} else {
			tmp["tags"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallProxyAddressTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressTaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProxyAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressUa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProxyAddressUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProxyAddressVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallProxyAddress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_image_base64"); ok || d.HasChange("_image_base64") {
		t, err := expandObjectFirewallProxyAddressImageBase64(d, v, "_image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_image-base64"] = t
		}
	}

	if v, ok := d.GetOk("case_sensitivity"); ok || d.HasChange("case_sensitivity") {
		t, err := expandObjectFirewallProxyAddressCaseSensitivity(d, v, "case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandObjectFirewallProxyAddressCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandObjectFirewallProxyAddressColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectFirewallProxyAddressComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandObjectFirewallProxyAddressHeader(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("header_group"); ok || d.HasChange("header_group") {
		t, err := expandObjectFirewallProxyAddressHeaderGroup(d, v, "header_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-group"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandObjectFirewallProxyAddressHeaderName(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandObjectFirewallProxyAddressHost(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("host_regex"); ok || d.HasChange("host_regex") {
		t, err := expandObjectFirewallProxyAddressHostRegex(d, v, "host_regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-regex"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok || d.HasChange("method") {
		t, err := expandObjectFirewallProxyAddressMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallProxyAddressName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("path"); ok || d.HasChange("path") {
		t, err := expandObjectFirewallProxyAddressPath(d, v, "path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["path"] = t
		}
	}

	if v, ok := d.GetOk("query"); ok || d.HasChange("query") {
		t, err := expandObjectFirewallProxyAddressQuery(d, v, "query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query"] = t
		}
	}

	if v, ok := d.GetOk("referrer"); ok || d.HasChange("referrer") {
		t, err := expandObjectFirewallProxyAddressReferrer(d, v, "referrer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referrer"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandObjectFirewallProxyAddressTagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallProxyAddressType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("ua"); ok || d.HasChange("ua") {
		t, err := expandObjectFirewallProxyAddressUa(d, v, "ua")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectFirewallProxyAddressUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok || d.HasChange("visibility") {
		t, err := expandObjectFirewallProxyAddressVisibility(d, v, "visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	return &obj, nil
}
