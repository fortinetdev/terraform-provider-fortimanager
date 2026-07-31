// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure CASB SaaS application.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbSaasApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbSaasApplicationCreate,
		Read:   resourceObjectCasbSaasApplicationRead,
		Update: resourceObjectCasbSaasApplicationUpdate,
		Delete: resourceObjectCasbSaasApplicationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"casb_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domains": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"icon_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"input_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fallback_input": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"required": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"output_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"optional": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"required": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"popularity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
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

func resourceObjectCasbSaasApplicationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectCasbSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbSaasApplication resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectCasbSaasApplication(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectCasbSaasApplication(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectCasbSaasApplication resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectCasbSaasApplication(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectCasbSaasApplication resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbSaasApplicationRead(d, m)
}

func resourceObjectCasbSaasApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectCasbSaasApplication(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbSaasApplication resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbSaasApplication(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbSaasApplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbSaasApplicationRead(d, m)
}

func resourceObjectCasbSaasApplicationDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbSaasApplication(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbSaasApplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbSaasApplicationRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectCasbSaasApplication(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectCasbSaasApplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbSaasApplication(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbSaasApplication resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbSaasApplicationCasbName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationDisplayName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationDomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbSaasApplicationIconId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributes(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attr_type"
		if _, ok := i["attr-type"]; ok {
			v := flattenObjectCasbSaasApplicationInputAttributesAttrType(i["attr-type"], d, pre_append)
			tmp["attr_type"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-InputAttributes-AttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := i["default"]; ok {
			v := flattenObjectCasbSaasApplicationInputAttributesDefault(i["default"], d, pre_append)
			tmp["default"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-InputAttributes-Default")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectCasbSaasApplicationInputAttributesDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-InputAttributes-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fallback_input"
		if _, ok := i["fallback-input"]; ok {
			v := flattenObjectCasbSaasApplicationInputAttributesFallbackInput(i["fallback-input"], d, pre_append)
			tmp["fallback_input"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-InputAttributes-FallbackInput")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbSaasApplicationInputAttributesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-InputAttributes-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := i["required"]; ok {
			v := flattenObjectCasbSaasApplicationInputAttributesRequired(i["required"], d, pre_append)
			tmp["required"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-InputAttributes-Required")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectCasbSaasApplicationInputAttributesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-InputAttributes-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbSaasApplicationInputAttributesAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesFallbackInput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributes(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attr_type"
		if _, ok := i["attr-type"]; ok {
			v := flattenObjectCasbSaasApplicationOutputAttributesAttrType(i["attr-type"], d, pre_append)
			tmp["attr_type"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-OutputAttributes-AttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectCasbSaasApplicationOutputAttributesDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-OutputAttributes-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbSaasApplicationOutputAttributesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-OutputAttributes-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := i["optional"]; ok {
			v := flattenObjectCasbSaasApplicationOutputAttributesOptional(i["optional"], d, pre_append)
			tmp["optional"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-OutputAttributes-Optional")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := i["required"]; ok {
			v := flattenObjectCasbSaasApplicationOutputAttributesRequired(i["required"], d, pre_append)
			tmp["required"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-OutputAttributes-Required")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectCasbSaasApplicationOutputAttributesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectCasbSaasApplication-OutputAttributes-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbSaasApplicationOutputAttributesAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesOptional(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationPopularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbSaasApplication(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("casb_name", flattenObjectCasbSaasApplicationCasbName(o["casb-name"], d, "casb_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-name"], "ObjectCasbSaasApplication-CasbName"); ok {
			if err = d.Set("casb_name", vv); err != nil {
				return fmt.Errorf("Error reading casb_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_name: %v", err)
		}
	}

	if err = d.Set("category", flattenObjectCasbSaasApplicationCategory(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ObjectCasbSaasApplication-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectCasbSaasApplicationDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectCasbSaasApplication-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("display_name", flattenObjectCasbSaasApplicationDisplayName(o["display-name"], d, "display_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-name"], "ObjectCasbSaasApplication-DisplayName"); ok {
			if err = d.Set("display_name", vv); err != nil {
				return fmt.Errorf("Error reading display_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_name: %v", err)
		}
	}

	if err = d.Set("domains", flattenObjectCasbSaasApplicationDomains(o["domains"], d, "domains")); err != nil {
		if vv, ok := fortiAPIPatch(o["domains"], "ObjectCasbSaasApplication-Domains"); ok {
			if err = d.Set("domains", vv); err != nil {
				return fmt.Errorf("Error reading domains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domains: %v", err)
		}
	}

	if err = d.Set("icon_id", flattenObjectCasbSaasApplicationIconId(o["icon-id"], d, "icon_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["icon-id"], "ObjectCasbSaasApplication-IconId"); ok {
			if err = d.Set("icon_id", vv); err != nil {
				return fmt.Errorf("Error reading icon_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icon_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("input_attributes", flattenObjectCasbSaasApplicationInputAttributes(o["input-attributes"], d, "input_attributes")); err != nil {
			if vv, ok := fortiAPIPatch(o["input-attributes"], "ObjectCasbSaasApplication-InputAttributes"); ok {
				if err = d.Set("input_attributes", vv); err != nil {
					return fmt.Errorf("Error reading input_attributes: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading input_attributes: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("input_attributes"); ok {
			if err = d.Set("input_attributes", flattenObjectCasbSaasApplicationInputAttributes(o["input-attributes"], d, "input_attributes")); err != nil {
				if vv, ok := fortiAPIPatch(o["input-attributes"], "ObjectCasbSaasApplication-InputAttributes"); ok {
					if err = d.Set("input_attributes", vv); err != nil {
						return fmt.Errorf("Error reading input_attributes: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading input_attributes: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectCasbSaasApplicationName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbSaasApplication-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("output_attributes", flattenObjectCasbSaasApplicationOutputAttributes(o["output-attributes"], d, "output_attributes")); err != nil {
			if vv, ok := fortiAPIPatch(o["output-attributes"], "ObjectCasbSaasApplication-OutputAttributes"); ok {
				if err = d.Set("output_attributes", vv); err != nil {
					return fmt.Errorf("Error reading output_attributes: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading output_attributes: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("output_attributes"); ok {
			if err = d.Set("output_attributes", flattenObjectCasbSaasApplicationOutputAttributes(o["output-attributes"], d, "output_attributes")); err != nil {
				if vv, ok := fortiAPIPatch(o["output-attributes"], "ObjectCasbSaasApplication-OutputAttributes"); ok {
					if err = d.Set("output_attributes", vv); err != nil {
						return fmt.Errorf("Error reading output_attributes: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading output_attributes: %v", err)
				}
			}
		}
	}

	if err = d.Set("popularity", flattenObjectCasbSaasApplicationPopularity(o["popularity"], d, "popularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["popularity"], "ObjectCasbSaasApplication-Popularity"); ok {
			if err = d.Set("popularity", vv); err != nil {
				return fmt.Errorf("Error reading popularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading popularity: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectCasbSaasApplicationStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectCasbSaasApplication-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCasbSaasApplicationType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCasbSaasApplication-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectCasbSaasApplicationUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectCasbSaasApplication-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbSaasApplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbSaasApplicationCasbName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationDisplayName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationDomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbSaasApplicationIconId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attr-type"], _ = expandObjectCasbSaasApplicationInputAttributesAttrType(d, i["attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default"], _ = expandObjectCasbSaasApplicationInputAttributesDefault(d, i["default"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectCasbSaasApplicationInputAttributesDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fallback_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fallback-input"], _ = expandObjectCasbSaasApplicationInputAttributesFallbackInput(d, i["fallback_input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbSaasApplicationInputAttributesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["required"], _ = expandObjectCasbSaasApplicationInputAttributesRequired(d, i["required"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectCasbSaasApplicationInputAttributesType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbSaasApplicationInputAttributesAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesFallbackInput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["attr-type"], _ = expandObjectCasbSaasApplicationOutputAttributesAttrType(d, i["attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectCasbSaasApplicationOutputAttributesDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbSaasApplicationOutputAttributesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["optional"], _ = expandObjectCasbSaasApplicationOutputAttributesOptional(d, i["optional"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["required"], _ = expandObjectCasbSaasApplicationOutputAttributesRequired(d, i["required"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectCasbSaasApplicationOutputAttributesType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbSaasApplicationOutputAttributesAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesOptional(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationPopularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbSaasApplication(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("casb_name"); ok || d.HasChange("casb_name") {
		t, err := expandObjectCasbSaasApplicationCasbName(d, v, "casb_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-name"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandObjectCasbSaasApplicationCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectCasbSaasApplicationDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("display_name"); ok || d.HasChange("display_name") {
		t, err := expandObjectCasbSaasApplicationDisplayName(d, v, "display_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-name"] = t
		}
	}

	if v, ok := d.GetOk("domains"); ok || d.HasChange("domains") {
		t, err := expandObjectCasbSaasApplicationDomains(d, v, "domains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domains"] = t
		}
	}

	if v, ok := d.GetOk("icon_id"); ok || d.HasChange("icon_id") {
		t, err := expandObjectCasbSaasApplicationIconId(d, v, "icon_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-id"] = t
		}
	}

	if v, ok := d.GetOk("input_attributes"); ok || d.HasChange("input_attributes") {
		t, err := expandObjectCasbSaasApplicationInputAttributes(d, v, "input_attributes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-attributes"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbSaasApplicationName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("output_attributes"); ok || d.HasChange("output_attributes") {
		t, err := expandObjectCasbSaasApplicationOutputAttributes(d, v, "output_attributes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-attributes"] = t
		}
	}

	if v, ok := d.GetOk("popularity"); ok || d.HasChange("popularity") {
		t, err := expandObjectCasbSaasApplicationPopularity(d, v, "popularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["popularity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectCasbSaasApplicationStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectCasbSaasApplicationType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectCasbSaasApplicationUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
