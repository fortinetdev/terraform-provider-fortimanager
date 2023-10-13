// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Servers to exempt from SSL inspection.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallSslSshProfileSslExempt() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallSslSshProfileSslExemptCreate,
		Read:   resourceObjectFirewallSslSshProfileSslExemptRead,
		Update: resourceObjectFirewallSslSshProfileSslExemptUpdate,
		Delete: resourceObjectFirewallSslSshProfileSslExemptDelete,

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
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortiguard_category": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"regex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_fqdn": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallSslSshProfileSslExemptCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectObjectFirewallSslSshProfileSslExempt(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallSslSshProfileSslExempt resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallSslSshProfileSslExempt(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallSslSshProfileSslExempt resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallSslSshProfileSslExemptRead(d, m)
}

func resourceObjectFirewallSslSshProfileSslExemptUpdate(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	obj, err := getObjectObjectFirewallSslSshProfileSslExempt(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfileSslExempt resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallSslSshProfileSslExempt(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallSslSshProfileSslExempt resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallSslSshProfileSslExemptRead(d, m)
}

func resourceObjectFirewallSslSshProfileSslExemptDelete(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	err = c.DeleteObjectFirewallSslSshProfileSslExempt(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallSslSshProfileSslExempt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallSslSshProfileSslExemptRead(d *schema.ResourceData, m interface{}) error {
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

	ssl_ssh_profile := d.Get("ssl_ssh_profile").(string)
	if ssl_ssh_profile == "" {
		ssl_ssh_profile = importOptionChecking(m.(*FortiClient).Cfg, "ssl_ssh_profile")
		if ssl_ssh_profile == "" {
			return fmt.Errorf("Parameter ssl_ssh_profile is missing")
		}
		if err = d.Set("ssl_ssh_profile", ssl_ssh_profile); err != nil {
			return fmt.Errorf("Error set params ssl_ssh_profile: %v", err)
		}
	}
	paradict["ssl_ssh_profile"] = ssl_ssh_profile

	o, err := c.ReadObjectFirewallSslSshProfileSslExempt(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfileSslExempt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallSslSshProfileSslExempt(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallSslSshProfileSslExempt resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallSslSshProfileSslExemptAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptAddress62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptFortiguardCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallSslSshProfileSslExemptId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptRegex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallSslSshProfileSslExemptWildcardFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallSslSshProfileSslExempt(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("address", flattenObjectFirewallSslSshProfileSslExemptAddress2edl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ObjectFirewallSslSshProfileSslExempt-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("address6", flattenObjectFirewallSslSshProfileSslExemptAddress62edl(o["address6"], d, "address6")); err != nil {
		if vv, ok := fortiAPIPatch(o["address6"], "ObjectFirewallSslSshProfileSslExempt-Address6"); ok {
			if err = d.Set("address6", vv); err != nil {
				return fmt.Errorf("Error reading address6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("fortiguard_category", flattenObjectFirewallSslSshProfileSslExemptFortiguardCategory2edl(o["fortiguard-category"], d, "fortiguard_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiguard-category"], "ObjectFirewallSslSshProfileSslExempt-FortiguardCategory"); ok {
			if err = d.Set("fortiguard_category", vv); err != nil {
				return fmt.Errorf("Error reading fortiguard_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiguard_category: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallSslSshProfileSslExemptId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallSslSshProfileSslExempt-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("regex", flattenObjectFirewallSslSshProfileSslExemptRegex2edl(o["regex"], d, "regex")); err != nil {
		if vv, ok := fortiAPIPatch(o["regex"], "ObjectFirewallSslSshProfileSslExempt-Regex"); ok {
			if err = d.Set("regex", vv); err != nil {
				return fmt.Errorf("Error reading regex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regex: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallSslSshProfileSslExemptType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallSslSshProfileSslExempt-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("wildcard_fqdn", flattenObjectFirewallSslSshProfileSslExemptWildcardFqdn2edl(o["wildcard-fqdn"], d, "wildcard_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard-fqdn"], "ObjectFirewallSslSshProfileSslExempt-WildcardFqdn"); ok {
			if err = d.Set("wildcard_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallSslSshProfileSslExemptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallSslSshProfileSslExemptAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptAddress62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptFortiguardCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectFirewallSslSshProfileSslExemptId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptRegex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallSslSshProfileSslExemptWildcardFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectObjectFirewallSslSshProfileSslExempt(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandObjectFirewallSslSshProfileSslExemptAddress2edl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("address6"); ok || d.HasChange("address6") {
		t, err := expandObjectFirewallSslSshProfileSslExemptAddress62edl(d, v, "address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address6"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_category"); ok || d.HasChange("fortiguard_category") {
		t, err := expandObjectFirewallSslSshProfileSslExemptFortiguardCategory2edl(d, v, "fortiguard_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-category"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallSslSshProfileSslExemptId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("regex"); ok || d.HasChange("regex") {
		t, err := expandObjectFirewallSslSshProfileSslExemptRegex2edl(d, v, "regex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regex"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallSslSshProfileSslExemptType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok || d.HasChange("wildcard_fqdn") {
		t, err := expandObjectFirewallSslSshProfileSslExemptWildcardFqdn2edl(d, v, "wildcard_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	}

	return &obj, nil
}
