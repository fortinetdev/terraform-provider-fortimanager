// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic InterfaceDynamicMapping

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDynamicInterfaceDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicInterfaceDynamicMappingCreate,
		Read:   resourceObjectDynamicInterfaceDynamicMappingRead,
		Update: resourceObjectDynamicInterfaceDynamicMappingUpdate,
		Delete: resourceObjectDynamicInterfaceDynamicMappingDelete,

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
			"interface": &schema.Schema{
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
			"egress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ingress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"intrazone_deny": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_intf": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceObjectDynamicInterfaceDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	obj, err := getObjectObjectDynamicInterfaceDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicInterfaceDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectDynamicInterfaceDynamicMapping(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicInterfaceDynamicMapping resource: %v", err)
	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectDynamicInterfaceDynamicMappingRead(d, m)
}

func resourceObjectDynamicInterfaceDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	obj, err := getObjectObjectDynamicInterfaceDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicInterfaceDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectDynamicInterfaceDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicInterfaceDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectDynamicInterfaceDynamicMappingRead(d, m)
}

func resourceObjectDynamicInterfaceDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	var_interface := d.Get("interface").(string)
	paradict["interface"] = var_interface

	wsParams["adom"] = adomv

	err = c.DeleteObjectDynamicInterfaceDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicInterfaceDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicInterfaceDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	var_interface := d.Get("interface").(string)
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["interface"] = var_interface

	o, err := c.ReadObjectDynamicInterfaceDynamicMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicInterfaceDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicInterfaceDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicInterfaceDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicInterfaceDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectDynamicInterfaceDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectDynamicInterfaceDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectDynamicInterfaceDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectDynamicInterfaceDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectDynamicInterfaceDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMappingEgressShapingProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectDynamicInterfaceDynamicMappingIngressShapingProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectDynamicInterfaceDynamicMappingIntrazoneDeny2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicInterfaceDynamicMappingLocalIntf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectDynamicInterfaceDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectDynamicInterfaceDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectDynamicInterfaceDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectDynamicInterfaceDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectDynamicInterfaceDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("egress_shaping_profile", flattenObjectDynamicInterfaceDynamicMappingEgressShapingProfile2edl(o["egress-shaping-profile"], d, "egress_shaping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["egress-shaping-profile"], "ObjectDynamicInterfaceDynamicMapping-EgressShapingProfile"); ok {
			if err = d.Set("egress_shaping_profile", vv); err != nil {
				return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("ingress_shaping_profile", flattenObjectDynamicInterfaceDynamicMappingIngressShapingProfile2edl(o["ingress-shaping-profile"], d, "ingress_shaping_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ingress-shaping-profile"], "ObjectDynamicInterfaceDynamicMapping-IngressShapingProfile"); ok {
			if err = d.Set("ingress_shaping_profile", vv); err != nil {
				return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("intrazone_deny", flattenObjectDynamicInterfaceDynamicMappingIntrazoneDeny2edl(o["intrazone-deny"], d, "intrazone_deny")); err != nil {
		if vv, ok := fortiAPIPatch(o["intrazone-deny"], "ObjectDynamicInterfaceDynamicMapping-IntrazoneDeny"); ok {
			if err = d.Set("intrazone_deny", vv); err != nil {
				return fmt.Errorf("Error reading intrazone_deny: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intrazone_deny: %v", err)
		}
	}

	if err = d.Set("local_intf", flattenObjectDynamicInterfaceDynamicMappingLocalIntf2edl(o["local-intf"], d, "local_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-intf"], "ObjectDynamicInterfaceDynamicMapping-LocalIntf"); ok {
			if err = d.Set("local_intf", vv); err != nil {
				return fmt.Errorf("Error reading local_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_intf: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicInterfaceDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicInterfaceDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectDynamicInterfaceDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectDynamicInterfaceDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectDynamicInterfaceDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMappingEgressShapingProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectDynamicInterfaceDynamicMappingIngressShapingProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectDynamicInterfaceDynamicMappingIntrazoneDeny2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicInterfaceDynamicMappingLocalIntf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectDynamicInterfaceDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectDynamicInterfaceDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("egress_shaping_profile"); ok || d.HasChange("egress_shaping_profile") {
		t, err := expandObjectDynamicInterfaceDynamicMappingEgressShapingProfile2edl(d, v, "egress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("ingress_shaping_profile"); ok || d.HasChange("ingress_shaping_profile") {
		t, err := expandObjectDynamicInterfaceDynamicMappingIngressShapingProfile2edl(d, v, "ingress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("intrazone_deny"); ok || d.HasChange("intrazone_deny") {
		t, err := expandObjectDynamicInterfaceDynamicMappingIntrazoneDeny2edl(d, v, "intrazone_deny")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intrazone-deny"] = t
		}
	}

	if v, ok := d.GetOk("local_intf"); ok || d.HasChange("local_intf") {
		t, err := expandObjectDynamicInterfaceDynamicMappingLocalIntf2edl(d, v, "local_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-intf"] = t
		}
	}

	return &obj, nil
}
