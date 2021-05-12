// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Import objects from device to ADOM, or from ADOM to Global.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSecurityconsoleImportDevObjs() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsoleImportDevObjsUpdate,
		Read:   resourceSecurityconsoleImportDevObjsRead,
		Update: resourceSecurityconsoleImportDevObjsUpdate,
		Delete: resourceSecurityconsoleImportDevObjsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"add_mappings": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_parent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"if_all_objs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"if_all_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"import_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"position": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSecurityconsoleImportDevObjsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectSecurityconsoleImportDevObjs(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleImportDevObjs resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsoleImportDevObjs(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleImportDevObjs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsoleImportDevObjs")

	return resourceSecurityconsoleImportDevObjsRead(d, m)
}

func resourceSecurityconsoleImportDevObjsDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsoleImportDevObjsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsoleImportDevObjsAddMappings(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSecurityconsoleImportDevObjsAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleImportDevObjsDstName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleImportDevObjsDstParent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleImportDevObjsIfAllObjs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "none",
			1: "all",
			2: "filter",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSecurityconsoleImportDevObjsIfAllPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSecurityconsoleImportDevObjsImportAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "do",
			1: "policy_search",
			2: "obj_search",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSecurityconsoleImportDevObjsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleImportDevObjsPosition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			-1: "bottom",
			0:  "top",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSecurityconsoleImportDevObjsVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSecurityconsoleImportDevObjs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("add_mappings", flattenSecurityconsoleImportDevObjsAddMappings(o["add_mappings"], d, "add_mappings")); err != nil {
		if vv, ok := fortiAPIPatch(o["add_mappings"], "SecurityconsoleImportDevObjs-AddMappings"); ok {
			if err = d.Set("add_mappings", vv); err != nil {
				return fmt.Errorf("Error reading add_mappings: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading add_mappings: %v", err)
		}
	}

	if err = d.Set("fmgadom", flattenSecurityconsoleImportDevObjsAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsoleImportDevObjs-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("dst_name", flattenSecurityconsoleImportDevObjsDstName(o["dst_name"], d, "dst_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst_name"], "SecurityconsoleImportDevObjs-DstName"); ok {
			if err = d.Set("dst_name", vv); err != nil {
				return fmt.Errorf("Error reading dst_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_name: %v", err)
		}
	}

	if err = d.Set("dst_parent", flattenSecurityconsoleImportDevObjsDstParent(o["dst_parent"], d, "dst_parent")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst_parent"], "SecurityconsoleImportDevObjs-DstParent"); ok {
			if err = d.Set("dst_parent", vv); err != nil {
				return fmt.Errorf("Error reading dst_parent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_parent: %v", err)
		}
	}

	if err = d.Set("if_all_objs", flattenSecurityconsoleImportDevObjsIfAllObjs(o["if_all_objs"], d, "if_all_objs")); err != nil {
		if vv, ok := fortiAPIPatch(o["if_all_objs"], "SecurityconsoleImportDevObjs-IfAllObjs"); ok {
			if err = d.Set("if_all_objs", vv); err != nil {
				return fmt.Errorf("Error reading if_all_objs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading if_all_objs: %v", err)
		}
	}

	if err = d.Set("if_all_policy", flattenSecurityconsoleImportDevObjsIfAllPolicy(o["if_all_policy"], d, "if_all_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["if_all_policy"], "SecurityconsoleImportDevObjs-IfAllPolicy"); ok {
			if err = d.Set("if_all_policy", vv); err != nil {
				return fmt.Errorf("Error reading if_all_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading if_all_policy: %v", err)
		}
	}

	if err = d.Set("import_action", flattenSecurityconsoleImportDevObjsImportAction(o["import_action"], d, "import_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["import_action"], "SecurityconsoleImportDevObjs-ImportAction"); ok {
			if err = d.Set("import_action", vv); err != nil {
				return fmt.Errorf("Error reading import_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_action: %v", err)
		}
	}

	if err = d.Set("name", flattenSecurityconsoleImportDevObjsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SecurityconsoleImportDevObjs-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("position", flattenSecurityconsoleImportDevObjsPosition(o["position"], d, "position")); err != nil {
		if vv, ok := fortiAPIPatch(o["position"], "SecurityconsoleImportDevObjs-Position"); ok {
			if err = d.Set("position", vv); err != nil {
				return fmt.Errorf("Error reading position: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading position: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSecurityconsoleImportDevObjsVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SecurityconsoleImportDevObjs-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSecurityconsoleImportDevObjsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsoleImportDevObjsAddMappings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsDstName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsDstParent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsIfAllObjs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsIfAllPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsImportAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsPosition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleImportDevObjsVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityconsoleImportDevObjs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_mappings"); ok {
		t, err := expandSecurityconsoleImportDevObjsAddMappings(d, v, "add_mappings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add_mappings"] = t
		}
	}

	if v, ok := d.GetOk("fmgadom"); ok {
		t, err := expandSecurityconsoleImportDevObjsAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("dst_name"); ok {
		t, err := expandSecurityconsoleImportDevObjsDstName(d, v, "dst_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst_name"] = t
		}
	}

	if v, ok := d.GetOk("dst_parent"); ok {
		t, err := expandSecurityconsoleImportDevObjsDstParent(d, v, "dst_parent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst_parent"] = t
		}
	}

	if v, ok := d.GetOk("if_all_objs"); ok {
		t, err := expandSecurityconsoleImportDevObjsIfAllObjs(d, v, "if_all_objs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["if_all_objs"] = t
		}
	}

	if v, ok := d.GetOk("if_all_policy"); ok {
		t, err := expandSecurityconsoleImportDevObjsIfAllPolicy(d, v, "if_all_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["if_all_policy"] = t
		}
	}

	if v, ok := d.GetOk("import_action"); ok {
		t, err := expandSecurityconsoleImportDevObjsImportAction(d, v, "import_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import_action"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSecurityconsoleImportDevObjsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("position"); ok {
		t, err := expandSecurityconsoleImportDevObjsPosition(d, v, "position")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["position"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSecurityconsoleImportDevObjsVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
