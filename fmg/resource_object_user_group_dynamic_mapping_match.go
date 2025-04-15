// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Group matches.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserGroupDynamicMappingMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserGroupDynamicMappingMatchCreate,
		Read:   resourceObjectUserGroupDynamicMappingMatchRead,
		Update: resourceObjectUserGroupDynamicMappingMatchUpdate,
		Delete: resourceObjectUserGroupDynamicMappingMatchDelete,

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
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dynamic_mapping_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_gui_meta": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectUserGroupDynamicMappingMatchCreate(d *schema.ResourceData, m interface{}) error {
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

	group := d.Get("group").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["group"] = group
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectObjectUserGroupDynamicMappingMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserGroupDynamicMappingMatch resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectUserGroupDynamicMappingMatch(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserGroupDynamicMappingMatch resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectUserGroupDynamicMappingMatchRead(d, m)
}

func resourceObjectUserGroupDynamicMappingMatchUpdate(d *schema.ResourceData, m interface{}) error {
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

	group := d.Get("group").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["group"] = group
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	obj, err := getObjectObjectUserGroupDynamicMappingMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGroupDynamicMappingMatch resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserGroupDynamicMappingMatch(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserGroupDynamicMappingMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectUserGroupDynamicMappingMatchRead(d, m)
}

func resourceObjectUserGroupDynamicMappingMatchDelete(d *schema.ResourceData, m interface{}) error {
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

	group := d.Get("group").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	paradict["group"] = group
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserGroupDynamicMappingMatch(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserGroupDynamicMappingMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserGroupDynamicMappingMatchRead(d *schema.ResourceData, m interface{}) error {
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

	group := d.Get("group").(string)
	dynamic_mapping_name := d.Get("dynamic_mapping_name").(string)
	dynamic_mapping_vdom := d.Get("dynamic_mapping_vdom").(string)
	if group == "" {
		group = importOptionChecking(m.(*FortiClient).Cfg, "group")
		if group == "" {
			return fmt.Errorf("Parameter group is missing")
		}
		if err = d.Set("group", group); err != nil {
			return fmt.Errorf("Error set params group: %v", err)
		}
	}
	if dynamic_mapping_name == "" {
		dynamic_mapping_name = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_name")
		if dynamic_mapping_name == "" {
			return fmt.Errorf("Parameter dynamic_mapping_name is missing")
		}
		if err = d.Set("dynamic_mapping_name", dynamic_mapping_name); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_name: %v", err)
		}
	}
	if dynamic_mapping_vdom == "" {
		dynamic_mapping_vdom = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_mapping_vdom")
		if dynamic_mapping_vdom == "" {
			return fmt.Errorf("Parameter dynamic_mapping_vdom is missing")
		}
		if err = d.Set("dynamic_mapping_vdom", dynamic_mapping_vdom); err != nil {
			return fmt.Errorf("Error set params dynamic_mapping_vdom: %v", err)
		}
	}
	paradict["group"] = group
	paradict["dynamic_mapping_name"] = dynamic_mapping_name
	paradict["dynamic_mapping_vdom"] = dynamic_mapping_vdom

	o, err := c.ReadObjectUserGroupDynamicMappingMatch(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserGroupDynamicMappingMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserGroupDynamicMappingMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserGroupDynamicMappingMatch resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserGroupDynamicMappingMatchGuiMeta3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchGroupName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserGroupDynamicMappingMatchServerName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectUserGroupDynamicMappingMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("_gui_meta", flattenObjectUserGroupDynamicMappingMatchGuiMeta3rdl(o["_gui_meta"], d, "_gui_meta")); err != nil {
		if vv, ok := fortiAPIPatch(o["_gui_meta"], "ObjectUserGroupDynamicMappingMatch-GuiMeta"); ok {
			if err = d.Set("_gui_meta", vv); err != nil {
				return fmt.Errorf("Error reading _gui_meta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _gui_meta: %v", err)
		}
	}

	if err = d.Set("group_name", flattenObjectUserGroupDynamicMappingMatchGroupName3rdl(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "ObjectUserGroupDynamicMappingMatch-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserGroupDynamicMappingMatchId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserGroupDynamicMappingMatch-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("server_name", flattenObjectUserGroupDynamicMappingMatchServerName3rdl(o["server-name"], d, "server_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-name"], "ObjectUserGroupDynamicMappingMatch-ServerName"); ok {
			if err = d.Set("server_name", vv); err != nil {
				return fmt.Errorf("Error reading server_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	return nil
}

func flattenObjectUserGroupDynamicMappingMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserGroupDynamicMappingMatchGuiMeta3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchGroupName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserGroupDynamicMappingMatchServerName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectUserGroupDynamicMappingMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_gui_meta"); ok || d.HasChange("_gui_meta") {
		t, err := expandObjectUserGroupDynamicMappingMatchGuiMeta3rdl(d, v, "_gui_meta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_gui_meta"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandObjectUserGroupDynamicMappingMatchGroupName3rdl(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserGroupDynamicMappingMatchId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok || d.HasChange("server_name") {
		t, err := expandObjectUserGroupDynamicMappingMatchServerName3rdl(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	return &obj, nil
}
