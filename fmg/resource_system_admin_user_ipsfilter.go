// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPS filter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminUserIpsFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminUserIpsFilterCreate,
		Read:   resourceSystemAdminUserIpsFilterRead,
		Update: resourceSystemAdminUserIpsFilterUpdate,
		Delete: resourceSystemAdminUserIpsFilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ips_filter_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminUserIpsFilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	user := d.Get("user").(string)
	paradict["user"] = user

	obj, err := getObjectSystemAdminUserIpsFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminUserIpsFilter resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemAdminUserIpsFilter(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminUserIpsFilter resource: %v", err)
	}

	d.SetId(getStringKey(d, "ips_filter_name"))

	return resourceSystemAdminUserIpsFilterRead(d, m)
}

func resourceSystemAdminUserIpsFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	user := d.Get("user").(string)
	paradict["user"] = user

	obj, err := getObjectSystemAdminUserIpsFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminUserIpsFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAdminUserIpsFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminUserIpsFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "ips_filter_name"))

	return resourceSystemAdminUserIpsFilterRead(d, m)
}

func resourceSystemAdminUserIpsFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAdminUserIpsFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminUserIpsFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminUserIpsFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAdminUserIpsFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminUserIpsFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminUserIpsFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminUserIpsFilter resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminUserIpsFilterIpsFilterName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminUserIpsFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ips_filter_name", flattenSystemAdminUserIpsFilterIpsFilterName2edl(o["ips-filter-name"], d, "ips_filter_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-filter-name"], "SystemAdminUserIpsFilter-IpsFilterName"); ok {
			if err = d.Set("ips_filter_name", vv); err != nil {
				return fmt.Errorf("Error reading ips_filter_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_filter_name: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminUserIpsFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminUserIpsFilterIpsFilterName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminUserIpsFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ips_filter_name"); ok || d.HasChange("ips_filter_name") {
		t, err := expandSystemAdminUserIpsFilterIpsFilterName2edl(d, v, "ips_filter_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-filter-name"] = t
		}
	}

	return &obj, nil
}
