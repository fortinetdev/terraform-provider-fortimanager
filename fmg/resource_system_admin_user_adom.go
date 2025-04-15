// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Admin domain.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminUserAdom() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminUserAdomCreate,
		Read:   resourceSystemAdminUserAdomRead,
		Update: resourceSystemAdminUserAdomUpdate,
		Delete: resourceSystemAdminUserAdomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"adom_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminUserAdomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	user := d.Get("user").(string)
	paradict["user"] = user

	obj, err := getObjectSystemAdminUserAdom(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminUserAdom resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemAdminUserAdom(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminUserAdom resource: %v", err)
	}

	d.SetId(getStringKey(d, "adom_name"))

	return resourceSystemAdminUserAdomRead(d, m)
}

func resourceSystemAdminUserAdomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	user := d.Get("user").(string)
	paradict["user"] = user

	obj, err := getObjectSystemAdminUserAdom(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminUserAdom resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAdminUserAdom(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminUserAdom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "adom_name"))

	return resourceSystemAdminUserAdomRead(d, m)
}

func resourceSystemAdminUserAdomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	user := d.Get("user").(string)
	paradict["user"] = user

	wsParams["adom"] = adomv

	err = c.DeleteSystemAdminUserAdom(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminUserAdom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminUserAdomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	user := d.Get("user").(string)
	if user == "" {
		user = importOptionChecking(m.(*FortiClient).Cfg, "user")
		if user == "" {
			return fmt.Errorf("Parameter user is missing")
		}
		if err = d.Set("user", user); err != nil {
			return fmt.Errorf("Error set params user: %v", err)
		}
	}
	paradict["user"] = user

	o, err := c.ReadSystemAdminUserAdom(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminUserAdom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminUserAdom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminUserAdom resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminUserAdomAdomName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminUserAdom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("adom_name", flattenSystemAdminUserAdomAdomName2edl(o["adom-name"], d, "adom_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-name"], "SystemAdminUserAdom-AdomName"); ok {
			if err = d.Set("adom_name", vv); err != nil {
				return fmt.Errorf("Error reading adom_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_name: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminUserAdomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminUserAdomAdomName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminUserAdom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adom_name"); ok || d.HasChange("adom_name") {
		t, err := expandSystemAdminUserAdomAdomName2edl(d, v, "adom_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-name"] = t
		}
	}

	return &obj, nil
}
