// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Server keytab.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallProfileProtocolOptionsCifsServerKeytab() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabCreate,
		Read:   resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabRead,
		Update: resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabUpdate,
		Delete: resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabDelete,

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"keytab": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"principal": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectObjectFirewallProfileProtocolOptionsCifsServerKeytab(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallProfileProtocolOptionsCifsServerKeytab resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallProfileProtocolOptionsCifsServerKeytab(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
	}

	d.SetId(getStringKey(d, "principal"))

	return resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectObjectFirewallProfileProtocolOptionsCifsServerKeytab(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsCifsServerKeytab resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallProfileProtocolOptionsCifsServerKeytab(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "principal"))

	return resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabDelete(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	err = c.DeleteObjectFirewallProfileProtocolOptionsCifsServerKeytab(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallProfileProtocolOptionsCifsServerKeytabRead(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadObjectFirewallProfileProtocolOptionsCifsServerKeytab(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsCifsServerKeytab resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallProfileProtocolOptionsCifsServerKeytab(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsCifsServerKeytab resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabKeytab3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabPassword3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabPrincipal3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallProfileProtocolOptionsCifsServerKeytab(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("keytab", flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabKeytab3rdl(o["keytab"], d, "keytab")); err != nil {
		if vv, ok := fortiAPIPatch(o["keytab"], "ObjectFirewallProfileProtocolOptionsCifsServerKeytab-Keytab"); ok {
			if err = d.Set("keytab", vv); err != nil {
				return fmt.Errorf("Error reading keytab: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keytab: %v", err)
		}
	}

	if err = d.Set("password", flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabPassword3rdl(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "ObjectFirewallProfileProtocolOptionsCifsServerKeytab-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("principal", flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabPrincipal3rdl(o["principal"], d, "principal")); err != nil {
		if vv, ok := fortiAPIPatch(o["principal"], "ObjectFirewallProfileProtocolOptionsCifsServerKeytab-Principal"); ok {
			if err = d.Set("principal", vv); err != nil {
				return fmt.Errorf("Error reading principal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading principal: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallProfileProtocolOptionsCifsServerKeytabFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallProfileProtocolOptionsCifsServerKeytabKeytab3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsCifsServerKeytabPassword3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsCifsServerKeytabPrincipal3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallProfileProtocolOptionsCifsServerKeytab(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("keytab"); ok || d.HasChange("keytab") {
		t, err := expandObjectFirewallProfileProtocolOptionsCifsServerKeytabKeytab3rdl(d, v, "keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keytab"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectFirewallProfileProtocolOptionsCifsServerKeytabPassword3rdl(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("principal"); ok || d.HasChange("principal") {
		t, err := expandObjectFirewallProfileProtocolOptionsCifsServerKeytabPrincipal3rdl(d, v, "principal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["principal"] = t
		}
	}

	return &obj, nil
}
