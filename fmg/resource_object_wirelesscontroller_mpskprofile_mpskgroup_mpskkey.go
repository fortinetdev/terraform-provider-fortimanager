// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of multiple PSK entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerMpskProfileMpskGroupMpskKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyCreate,
		Read:   resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyRead,
		Update: resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyUpdate,
		Delete: resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyDelete,

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
			"mpsk_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mpsk_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"concurrent_client_limit_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"concurrent_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk_schedules": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"passphrase": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"pmk": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
		},
	}
}

func resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	mpsk_profile := d.Get("mpsk_profile").(string)
	mpsk_group := d.Get("mpsk_group").(string)
	paradict["mpsk_profile"] = mpsk_profile
	paradict["mpsk_group"] = mpsk_group

	obj, err := getObjectObjectWirelessControllerMpskProfileMpskGroupMpskKey(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerMpskProfileMpskGroupMpskKey resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerMpskProfileMpskGroupMpskKey(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerMpskProfileMpskGroupMpskKey resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyRead(d, m)
}

func resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyUpdate(d *schema.ResourceData, m interface{}) error {
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

	mpsk_profile := d.Get("mpsk_profile").(string)
	mpsk_group := d.Get("mpsk_group").(string)
	paradict["mpsk_profile"] = mpsk_profile
	paradict["mpsk_group"] = mpsk_group

	obj, err := getObjectObjectWirelessControllerMpskProfileMpskGroupMpskKey(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerMpskProfileMpskGroupMpskKey resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerMpskProfileMpskGroupMpskKey(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerMpskProfileMpskGroupMpskKey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyRead(d, m)
}

func resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyDelete(d *schema.ResourceData, m interface{}) error {
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

	mpsk_profile := d.Get("mpsk_profile").(string)
	mpsk_group := d.Get("mpsk_group").(string)
	paradict["mpsk_profile"] = mpsk_profile
	paradict["mpsk_group"] = mpsk_group

	err = c.DeleteObjectWirelessControllerMpskProfileMpskGroupMpskKey(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerMpskProfileMpskGroupMpskKey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerMpskProfileMpskGroupMpskKeyRead(d *schema.ResourceData, m interface{}) error {
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

	mpsk_profile := d.Get("mpsk_profile").(string)
	mpsk_group := d.Get("mpsk_group").(string)
	if mpsk_profile == "" {
		mpsk_profile = importOptionChecking(m.(*FortiClient).Cfg, "mpsk_profile")
		if mpsk_profile == "" {
			return fmt.Errorf("Parameter mpsk_profile is missing")
		}
		if err = d.Set("mpsk_profile", mpsk_profile); err != nil {
			return fmt.Errorf("Error set params mpsk_profile: %v", err)
		}
	}
	if mpsk_group == "" {
		mpsk_group = importOptionChecking(m.(*FortiClient).Cfg, "mpsk_group")
		if mpsk_group == "" {
			return fmt.Errorf("Parameter mpsk_group is missing")
		}
		if err = d.Set("mpsk_group", mpsk_group); err != nil {
			return fmt.Errorf("Error set params mpsk_group: %v", err)
		}
	}
	paradict["mpsk_profile"] = mpsk_profile
	paradict["mpsk_group"] = mpsk_group

	o, err := c.ReadObjectWirelessControllerMpskProfileMpskGroupMpskKey(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerMpskProfileMpskGroupMpskKey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerMpskProfileMpskGroupMpskKey(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerMpskProfileMpskGroupMpskKey resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerMpskProfileMpskGroupMpskKey(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment3rdl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerMpskProfileMpskGroupMpskKey-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("concurrent_client_limit_type", flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType3rdl(o["concurrent-client-limit-type"], d, "concurrent_client_limit_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["concurrent-client-limit-type"], "ObjectWirelessControllerMpskProfileMpskGroupMpskKey-ConcurrentClientLimitType"); ok {
			if err = d.Set("concurrent_client_limit_type", vv); err != nil {
				return fmt.Errorf("Error reading concurrent_client_limit_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading concurrent_client_limit_type: %v", err)
		}
	}

	if err = d.Set("concurrent_clients", flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients3rdl(o["concurrent-clients"], d, "concurrent_clients")); err != nil {
		if vv, ok := fortiAPIPatch(o["concurrent-clients"], "ObjectWirelessControllerMpskProfileMpskGroupMpskKey-ConcurrentClients"); ok {
			if err = d.Set("concurrent_clients", vv); err != nil {
				return fmt.Errorf("Error reading concurrent_clients: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading concurrent_clients: %v", err)
		}
	}

	if err = d.Set("mac", flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac3rdl(o["mac"], d, "mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac"], "ObjectWirelessControllerMpskProfileMpskGroupMpskKey-Mac"); ok {
			if err = d.Set("mac", vv); err != nil {
				return fmt.Errorf("Error reading mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("mpsk_schedules", flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules3rdl(o["mpsk-schedules"], d, "mpsk_schedules")); err != nil {
		if vv, ok := fortiAPIPatch(o["mpsk-schedules"], "ObjectWirelessControllerMpskProfileMpskGroupMpskKey-MpskSchedules"); ok {
			if err = d.Set("mpsk_schedules", vv); err != nil {
				return fmt.Errorf("Error reading mpsk_schedules: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mpsk_schedules: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerMpskProfileMpskGroupMpskKey-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerMpskProfileMpskGroupMpskKeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPmk3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectWirelessControllerMpskProfileMpskGroupMpskKey(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyComment3rdl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("concurrent_client_limit_type"); ok || d.HasChange("concurrent_client_limit_type") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClientLimitType3rdl(d, v, "concurrent_client_limit_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concurrent-client-limit-type"] = t
		}
	}

	if v, ok := d.GetOk("concurrent_clients"); ok || d.HasChange("concurrent_clients") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyConcurrentClients3rdl(d, v, "concurrent_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concurrent-clients"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok || d.HasChange("mac") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMac3rdl(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_schedules"); ok || d.HasChange("mpsk_schedules") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyMpskSchedules3rdl(d, v, "mpsk_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-schedules"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("passphrase"); ok || d.HasChange("passphrase") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPassphrase3rdl(d, v, "passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	}

	if v, ok := d.GetOk("pmk"); ok || d.HasChange("pmk") {
		t, err := expandObjectWirelessControllerMpskProfileMpskGroupMpskKeyPmk3rdl(d, v, "pmk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmk"] = t
		}
	}

	return &obj, nil
}
