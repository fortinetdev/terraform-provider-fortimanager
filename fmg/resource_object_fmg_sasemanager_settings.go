// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFmg SaseManagerSettings

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFmgSaseManagerSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFmgSaseManagerSettingsUpdate,
		Read:   resourceObjectFmgSaseManagerSettingsRead,
		Update: resourceObjectFmgSaseManagerSettingsUpdate,
		Delete: resourceObjectFmgSaseManagerSettingsDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sync_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_profile_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFmgSaseManagerSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFmgSaseManagerSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgSaseManagerSettings resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFmgSaseManagerSettings(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgSaseManagerSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFmgSaseManagerSettings")

	return resourceObjectFmgSaseManagerSettingsRead(d, m)
}

func resourceObjectFmgSaseManagerSettingsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFmgSaseManagerSettings(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFmgSaseManagerSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFmgSaseManagerSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFmgSaseManagerSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgSaseManagerSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFmgSaseManagerSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgSaseManagerSettings resource from API: %v", err)
	}
	return nil
}

func flattenObjectFmgSaseManagerSettingsAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFmgSaseManagerSettingsProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFmgSaseManagerSettingsSyncAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgSaseManagerSettingsSyncProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgSaseManagerSettingsSyncUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgSaseManagerSettingsUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFmgSaseManagerSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("address", flattenObjectFmgSaseManagerSettingsAddress(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ObjectFmgSaseManagerSettings-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenObjectFmgSaseManagerSettingsProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-group"], "ObjectFmgSaseManagerSettings-ProfileGroup"); ok {
			if err = d.Set("profile_group", vv); err != nil {
				return fmt.Errorf("Error reading profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("sync_address", flattenObjectFmgSaseManagerSettingsSyncAddress(o["sync-address"], d, "sync_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-address"], "ObjectFmgSaseManagerSettings-SyncAddress"); ok {
			if err = d.Set("sync_address", vv); err != nil {
				return fmt.Errorf("Error reading sync_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_address: %v", err)
		}
	}

	if err = d.Set("sync_profile_group", flattenObjectFmgSaseManagerSettingsSyncProfileGroup(o["sync-profile-group"], d, "sync_profile_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-profile-group"], "ObjectFmgSaseManagerSettings-SyncProfileGroup"); ok {
			if err = d.Set("sync_profile_group", vv); err != nil {
				return fmt.Errorf("Error reading sync_profile_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_profile_group: %v", err)
		}
	}

	if err = d.Set("sync_user", flattenObjectFmgSaseManagerSettingsSyncUser(o["sync-user"], d, "sync_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-user"], "ObjectFmgSaseManagerSettings-SyncUser"); ok {
			if err = d.Set("sync_user", vv); err != nil {
				return fmt.Errorf("Error reading sync_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_user: %v", err)
		}
	}

	if err = d.Set("user", flattenObjectFmgSaseManagerSettingsUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "ObjectFmgSaseManagerSettings-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenObjectFmgSaseManagerSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFmgSaseManagerSettingsAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFmgSaseManagerSettingsProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFmgSaseManagerSettingsSyncAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgSaseManagerSettingsSyncProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgSaseManagerSettingsSyncUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgSaseManagerSettingsUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFmgSaseManagerSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandObjectFmgSaseManagerSettingsAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok || d.HasChange("profile_group") {
		t, err := expandObjectFmgSaseManagerSettingsProfileGroup(d, v, "profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("sync_address"); ok || d.HasChange("sync_address") {
		t, err := expandObjectFmgSaseManagerSettingsSyncAddress(d, v, "sync_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-address"] = t
		}
	}

	if v, ok := d.GetOk("sync_profile_group"); ok || d.HasChange("sync_profile_group") {
		t, err := expandObjectFmgSaseManagerSettingsSyncProfileGroup(d, v, "sync_profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-profile-group"] = t
		}
	}

	if v, ok := d.GetOk("sync_user"); ok || d.HasChange("sync_user") {
		t, err := expandObjectFmgSaseManagerSettingsSyncUser(d, v, "sync_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-user"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandObjectFmgSaseManagerSettingsUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
