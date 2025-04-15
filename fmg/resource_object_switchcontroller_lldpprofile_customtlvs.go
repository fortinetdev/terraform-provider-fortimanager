// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit custom TLV entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerLldpProfileCustomTlvs() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerLldpProfileCustomTlvsCreate,
		Read:   resourceObjectSwitchControllerLldpProfileCustomTlvsRead,
		Update: resourceObjectSwitchControllerLldpProfileCustomTlvsUpdate,
		Delete: resourceObjectSwitchControllerLldpProfileCustomTlvsDelete,

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
			"lldp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"information_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"oui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subtype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerLldpProfileCustomTlvsCreate(d *schema.ResourceData, m interface{}) error {
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

	lldp_profile := d.Get("lldp_profile").(string)
	paradict["lldp_profile"] = lldp_profile

	obj, err := getObjectObjectSwitchControllerLldpProfileCustomTlvs(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerLldpProfileCustomTlvs resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerLldpProfileCustomTlvs(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerLldpProfileCustomTlvs resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerLldpProfileCustomTlvsRead(d, m)
}

func resourceObjectSwitchControllerLldpProfileCustomTlvsUpdate(d *schema.ResourceData, m interface{}) error {
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

	lldp_profile := d.Get("lldp_profile").(string)
	paradict["lldp_profile"] = lldp_profile

	obj, err := getObjectObjectSwitchControllerLldpProfileCustomTlvs(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerLldpProfileCustomTlvs resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerLldpProfileCustomTlvs(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerLldpProfileCustomTlvs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerLldpProfileCustomTlvsRead(d, m)
}

func resourceObjectSwitchControllerLldpProfileCustomTlvsDelete(d *schema.ResourceData, m interface{}) error {
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

	lldp_profile := d.Get("lldp_profile").(string)
	paradict["lldp_profile"] = lldp_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerLldpProfileCustomTlvs(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerLldpProfileCustomTlvs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerLldpProfileCustomTlvsRead(d *schema.ResourceData, m interface{}) error {
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

	lldp_profile := d.Get("lldp_profile").(string)
	if lldp_profile == "" {
		lldp_profile = importOptionChecking(m.(*FortiClient).Cfg, "lldp_profile")
		if lldp_profile == "" {
			return fmt.Errorf("Parameter lldp_profile is missing")
		}
		if err = d.Set("lldp_profile", lldp_profile); err != nil {
			return fmt.Errorf("Error set params lldp_profile: %v", err)
		}
	}
	paradict["lldp_profile"] = lldp_profile

	o, err := c.ReadObjectSwitchControllerLldpProfileCustomTlvs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerLldpProfileCustomTlvs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerLldpProfileCustomTlvs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerLldpProfileCustomTlvs resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsInformationString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsOui2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsSubtype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("information_string", flattenObjectSwitchControllerLldpProfileCustomTlvsInformationString2edl(o["information-string"], d, "information_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["information-string"], "ObjectSwitchControllerLldpProfileCustomTlvs-InformationString"); ok {
			if err = d.Set("information_string", vv); err != nil {
				return fmt.Errorf("Error reading information_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading information_string: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerLldpProfileCustomTlvsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerLldpProfileCustomTlvs-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("oui", flattenObjectSwitchControllerLldpProfileCustomTlvsOui2edl(o["oui"], d, "oui")); err != nil {
		if vv, ok := fortiAPIPatch(o["oui"], "ObjectSwitchControllerLldpProfileCustomTlvs-Oui"); ok {
			if err = d.Set("oui", vv); err != nil {
				return fmt.Errorf("Error reading oui: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oui: %v", err)
		}
	}

	if err = d.Set("subtype", flattenObjectSwitchControllerLldpProfileCustomTlvsSubtype2edl(o["subtype"], d, "subtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["subtype"], "ObjectSwitchControllerLldpProfileCustomTlvs-Subtype"); ok {
			if err = d.Set("subtype", vv); err != nil {
				return fmt.Errorf("Error reading subtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subtype: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerLldpProfileCustomTlvsInformationString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileCustomTlvsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileCustomTlvsOui2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileCustomTlvsSubtype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("information_string"); ok || d.HasChange("information_string") {
		t, err := expandObjectSwitchControllerLldpProfileCustomTlvsInformationString2edl(d, v, "information_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["information-string"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerLldpProfileCustomTlvsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("oui"); ok || d.HasChange("oui") {
		t, err := expandObjectSwitchControllerLldpProfileCustomTlvsOui2edl(d, v, "oui")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oui"] = t
		}
	}

	if v, ok := d.GetOk("subtype"); ok || d.HasChange("subtype") {
		t, err := expandObjectSwitchControllerLldpProfileCustomTlvsSubtype2edl(d, v, "subtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subtype"] = t
		}
	}

	return &obj, nil
}
