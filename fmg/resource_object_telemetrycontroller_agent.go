// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiTelemetry agents managed by a FortiGate unit.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectTelemetryControllerAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectTelemetryControllerAgentCreate,
		Read:   resourceObjectTelemetryControllerAgentRead,
		Update: resourceObjectTelemetryControllerAgentUpdate,
		Delete: resourceObjectTelemetryControllerAgentDelete,

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
			"agent_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"agent_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"authz": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectTelemetryControllerAgentCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectTelemetryControllerAgent(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectTelemetryControllerAgent resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("agent_id")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectTelemetryControllerAgent(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectTelemetryControllerAgent(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectTelemetryControllerAgent resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectTelemetryControllerAgent(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectTelemetryControllerAgent resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "agent_id"))

	return resourceObjectTelemetryControllerAgentRead(d, m)
}

func resourceObjectTelemetryControllerAgentUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectTelemetryControllerAgent(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerAgent resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectTelemetryControllerAgent(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectTelemetryControllerAgent resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "agent_id"))

	return resourceObjectTelemetryControllerAgentRead(d, m)
}

func resourceObjectTelemetryControllerAgentDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectTelemetryControllerAgent(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectTelemetryControllerAgent resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectTelemetryControllerAgentRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectTelemetryControllerAgent(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectTelemetryControllerAgent resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectTelemetryControllerAgent(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectTelemetryControllerAgent resource from API: %v", err)
	}
	return nil
}

func flattenObjectTelemetryControllerAgentAgentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerAgentAgentProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectTelemetryControllerAgentAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerAgentAuthz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectTelemetryControllerAgentComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectTelemetryControllerAgent(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("agent_id", flattenObjectTelemetryControllerAgentAgentId(o["agent-id"], d, "agent_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["agent-id"], "ObjectTelemetryControllerAgent-AgentId"); ok {
			if err = d.Set("agent_id", vv); err != nil {
				return fmt.Errorf("Error reading agent_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agent_id: %v", err)
		}
	}

	if err = d.Set("agent_profile", flattenObjectTelemetryControllerAgentAgentProfile(o["agent-profile"], d, "agent_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["agent-profile"], "ObjectTelemetryControllerAgent-AgentProfile"); ok {
			if err = d.Set("agent_profile", vv); err != nil {
				return fmt.Errorf("Error reading agent_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading agent_profile: %v", err)
		}
	}

	if err = d.Set("alias", flattenObjectTelemetryControllerAgentAlias(o["alias"], d, "alias")); err != nil {
		if vv, ok := fortiAPIPatch(o["alias"], "ObjectTelemetryControllerAgent-Alias"); ok {
			if err = d.Set("alias", vv); err != nil {
				return fmt.Errorf("Error reading alias: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("authz", flattenObjectTelemetryControllerAgentAuthz(o["authz"], d, "authz")); err != nil {
		if vv, ok := fortiAPIPatch(o["authz"], "ObjectTelemetryControllerAgent-Authz"); ok {
			if err = d.Set("authz", vv); err != nil {
				return fmt.Errorf("Error reading authz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authz: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectTelemetryControllerAgentComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectTelemetryControllerAgent-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenObjectTelemetryControllerAgentFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectTelemetryControllerAgentAgentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerAgentAgentProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectTelemetryControllerAgentAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerAgentAuthz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectTelemetryControllerAgentComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectTelemetryControllerAgent(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("agent_id"); ok || d.HasChange("agent_id") {
		t, err := expandObjectTelemetryControllerAgentAgentId(d, v, "agent_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agent-id"] = t
		}
	}

	if v, ok := d.GetOk("agent_profile"); ok || d.HasChange("agent_profile") {
		t, err := expandObjectTelemetryControllerAgentAgentProfile(d, v, "agent_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agent-profile"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok || d.HasChange("alias") {
		t, err := expandObjectTelemetryControllerAgentAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("authz"); ok || d.HasChange("authz") {
		t, err := expandObjectTelemetryControllerAgentAuthz(d, v, "authz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authz"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectTelemetryControllerAgentComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
