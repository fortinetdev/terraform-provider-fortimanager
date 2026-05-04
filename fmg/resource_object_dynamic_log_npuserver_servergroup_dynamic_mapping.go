// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic LogNpuServerServerGroupDynamicMapping

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDynamicLogNpuServerServerGroupDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicLogNpuServerServerGroupDynamicMappingCreate,
		Read:   resourceObjectDynamicLogNpuServerServerGroupDynamicMappingRead,
		Update: resourceObjectDynamicLogNpuServerServerGroupDynamicMappingUpdate,
		Delete: resourceObjectDynamicLogNpuServerServerGroupDynamicMappingDelete,

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
			"server_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_scope": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"group_name": &schema.Schema{
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

func resourceObjectDynamicLogNpuServerServerGroupDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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

	server_group := d.Get("server_group").(string)
	paradict["server_group"] = server_group

	obj, err := getObjectObjectDynamicLogNpuServerServerGroupDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicLogNpuServerServerGroupDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("_scope")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectDynamicLogNpuServerServerGroupDynamicMapping(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectDynamicLogNpuServerServerGroupDynamicMapping(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectDynamicLogNpuServerServerGroupDynamicMapping resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectDynamicLogNpuServerServerGroupDynamicMapping(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectDynamicLogNpuServerServerGroupDynamicMapping resource: %v", err)
		}

	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectDynamicLogNpuServerServerGroupDynamicMappingRead(d, m)
}

func resourceObjectDynamicLogNpuServerServerGroupDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	server_group := d.Get("server_group").(string)
	paradict["server_group"] = server_group

	obj, err := getObjectObjectDynamicLogNpuServerServerGroupDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicLogNpuServerServerGroupDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDynamicLogNpuServerServerGroupDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicLogNpuServerServerGroupDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectDynamicLogNpuServerServerGroupDynamicMappingRead(d, m)
}

func resourceObjectDynamicLogNpuServerServerGroupDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	server_group := d.Get("server_group").(string)
	paradict["server_group"] = server_group

	wsParams["adom"] = adomv

	err = c.DeleteObjectDynamicLogNpuServerServerGroupDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicLogNpuServerServerGroupDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicLogNpuServerServerGroupDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	server_group := d.Get("server_group").(string)
	if server_group == "" {
		server_group = importOptionChecking(m.(*FortiClient).Cfg, "server_group")
		if server_group == "" {
			return fmt.Errorf("Parameter server_group is missing")
		}
		if err = d.Set("server_group", server_group); err != nil {
			return fmt.Errorf("Error set params server_group: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["server_group"] = server_group

	o, err := c.ReadObjectDynamicLogNpuServerServerGroupDynamicMapping(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectDynamicLogNpuServerServerGroupDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicLogNpuServerServerGroupDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicLogNpuServerServerGroupDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicLogNpuServerServerGroupDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectDynamicLogNpuServerServerGroupDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectDynamicLogNpuServerServerGroupDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectDynamicLogNpuServerServerGroupDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectDynamicLogNpuServerServerGroupDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectDynamicLogNpuServerServerGroupDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicLogNpuServerServerGroupDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicLogNpuServerServerGroupDynamicMappingGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDynamicLogNpuServerServerGroupDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectDynamicLogNpuServerServerGroupDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectDynamicLogNpuServerServerGroupDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectDynamicLogNpuServerServerGroupDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectDynamicLogNpuServerServerGroupDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_name", flattenObjectDynamicLogNpuServerServerGroupDynamicMappingGroupName2edl(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "ObjectDynamicLogNpuServerServerGroupDynamicMapping-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicLogNpuServerServerGroupDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicLogNpuServerServerGroupDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectDynamicLogNpuServerServerGroupDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectDynamicLogNpuServerServerGroupDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectDynamicLogNpuServerServerGroupDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicLogNpuServerServerGroupDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicLogNpuServerServerGroupDynamicMappingGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDynamicLogNpuServerServerGroupDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectDynamicLogNpuServerServerGroupDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandObjectDynamicLogNpuServerServerGroupDynamicMappingGroupName2edl(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	return &obj, nil
}
