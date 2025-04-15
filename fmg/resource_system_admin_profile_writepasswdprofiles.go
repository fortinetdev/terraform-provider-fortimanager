// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Profile list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminProfileWritePasswdProfiles() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminProfileWritePasswdProfilesCreate,
		Read:   resourceSystemAdminProfileWritePasswdProfilesRead,
		Update: resourceSystemAdminProfileWritePasswdProfilesUpdate,
		Delete: resourceSystemAdminProfileWritePasswdProfilesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"profileid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminProfileWritePasswdProfilesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectSystemAdminProfileWritePasswdProfiles(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminProfileWritePasswdProfiles resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemAdminProfileWritePasswdProfiles(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminProfileWritePasswdProfiles resource: %v", err)
	}

	d.SetId(getStringKey(d, "profileid"))

	return resourceSystemAdminProfileWritePasswdProfilesRead(d, m)
}

func resourceSystemAdminProfileWritePasswdProfilesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectSystemAdminProfileWritePasswdProfiles(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminProfileWritePasswdProfiles resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAdminProfileWritePasswdProfiles(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminProfileWritePasswdProfiles resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "profileid"))

	return resourceSystemAdminProfileWritePasswdProfilesRead(d, m)
}

func resourceSystemAdminProfileWritePasswdProfilesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteSystemAdminProfileWritePasswdProfiles(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminProfileWritePasswdProfiles resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminProfileWritePasswdProfilesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadSystemAdminProfileWritePasswdProfiles(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminProfileWritePasswdProfiles resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminProfileWritePasswdProfiles(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminProfileWritePasswdProfiles resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminProfileWritePasswdProfilesProfileid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminProfileWritePasswdProfiles(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("profileid", flattenSystemAdminProfileWritePasswdProfilesProfileid2edl(o["profileid"], d, "profileid")); err != nil {
		if vv, ok := fortiAPIPatch(o["profileid"], "SystemAdminProfileWritePasswdProfiles-Profileid"); ok {
			if err = d.Set("profileid", vv); err != nil {
				return fmt.Errorf("Error reading profileid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profileid: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminProfileWritePasswdProfilesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminProfileWritePasswdProfilesProfileid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminProfileWritePasswdProfiles(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("profileid"); ok || d.HasChange("profileid") {
		t, err := expandSystemAdminProfileWritePasswdProfilesProfileid2edl(d, v, "profileid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profileid"] = t
		}
	}

	return &obj, nil
}
