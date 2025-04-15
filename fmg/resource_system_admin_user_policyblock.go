// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Policy block write access.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminUserPolicyBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminUserPolicyBlockCreate,
		Read:   resourceSystemAdminUserPolicyBlockRead,
		Update: resourceSystemAdminUserPolicyBlockUpdate,
		Delete: resourceSystemAdminUserPolicyBlockDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy_block_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminUserPolicyBlockCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	user := d.Get("user").(string)
	paradict["user"] = user

	obj, err := getObjectSystemAdminUserPolicyBlock(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminUserPolicyBlock resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemAdminUserPolicyBlock(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminUserPolicyBlock resource: %v", err)
	}

	d.SetId(getStringKey(d, "policy_block_name"))

	return resourceSystemAdminUserPolicyBlockRead(d, m)
}

func resourceSystemAdminUserPolicyBlockUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	user := d.Get("user").(string)
	paradict["user"] = user

	obj, err := getObjectSystemAdminUserPolicyBlock(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminUserPolicyBlock resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemAdminUserPolicyBlock(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminUserPolicyBlock resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "policy_block_name"))

	return resourceSystemAdminUserPolicyBlockRead(d, m)
}

func resourceSystemAdminUserPolicyBlockDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAdminUserPolicyBlock(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminUserPolicyBlock resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminUserPolicyBlockRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAdminUserPolicyBlock(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminUserPolicyBlock resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminUserPolicyBlock(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminUserPolicyBlock resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminUserPolicyBlockPolicyBlockName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminUserPolicyBlock(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("policy_block_name", flattenSystemAdminUserPolicyBlockPolicyBlockName2edl(o["policy-block-name"], d, "policy_block_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-block-name"], "SystemAdminUserPolicyBlock-PolicyBlockName"); ok {
			if err = d.Set("policy_block_name", vv); err != nil {
				return fmt.Errorf("Error reading policy_block_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_block_name: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminUserPolicyBlockFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminUserPolicyBlockPolicyBlockName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminUserPolicyBlock(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("policy_block_name"); ok || d.HasChange("policy_block_name") {
		t, err := expandSystemAdminUserPolicyBlockPolicyBlockName2edl(d, v, "policy_block_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-block-name"] = t
		}
	}

	return &obj, nil
}
