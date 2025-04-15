// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: User list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminProfileWritePasswdUserList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminProfileWritePasswdUserListCreate,
		Read:   resourceSystemAdminProfileWritePasswdUserListRead,
		Update: resourceSystemAdminProfileWritePasswdUserListUpdate,
		Delete: resourceSystemAdminProfileWritePasswdUserListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"userid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminProfileWritePasswdUserListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectSystemAdminProfileWritePasswdUserList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminProfileWritePasswdUserList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemAdminProfileWritePasswdUserList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminProfileWritePasswdUserList resource: %v", err)
	}

	d.SetId(getStringKey(d, "userid"))

	return resourceSystemAdminProfileWritePasswdUserListRead(d, m)
}

func resourceSystemAdminProfileWritePasswdUserListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectSystemAdminProfileWritePasswdUserList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminProfileWritePasswdUserList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAdminProfileWritePasswdUserList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminProfileWritePasswdUserList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "userid"))

	return resourceSystemAdminProfileWritePasswdUserListRead(d, m)
}

func resourceSystemAdminProfileWritePasswdUserListDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAdminProfileWritePasswdUserList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminProfileWritePasswdUserList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminProfileWritePasswdUserListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAdminProfileWritePasswdUserList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminProfileWritePasswdUserList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminProfileWritePasswdUserList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminProfileWritePasswdUserList resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminProfileWritePasswdUserListUserid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminProfileWritePasswdUserList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("userid", flattenSystemAdminProfileWritePasswdUserListUserid2edl(o["userid"], d, "userid")); err != nil {
		if vv, ok := fortiAPIPatch(o["userid"], "SystemAdminProfileWritePasswdUserList-Userid"); ok {
			if err = d.Set("userid", vv); err != nil {
				return fmt.Errorf("Error reading userid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading userid: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminProfileWritePasswdUserListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminProfileWritePasswdUserListUserid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminProfileWritePasswdUserList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("userid"); ok || d.HasChange("userid") {
		t, err := expandSystemAdminProfileWritePasswdUserListUserid2edl(d, v, "userid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["userid"] = t
		}
	}

	return &obj, nil
}
