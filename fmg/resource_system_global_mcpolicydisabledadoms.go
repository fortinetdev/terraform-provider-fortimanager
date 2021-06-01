// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Multicast policy disabled adoms.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemGlobalMcPolicyDisabledAdoms() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGlobalMcPolicyDisabledAdomsCreate,
		Read:   resourceSystemGlobalMcPolicyDisabledAdomsRead,
		Update: resourceSystemGlobalMcPolicyDisabledAdomsUpdate,
		Delete: resourceSystemGlobalMcPolicyDisabledAdomsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"adom_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemGlobalMcPolicyDisabledAdomsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemGlobalMcPolicyDisabledAdoms(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemGlobalMcPolicyDisabledAdoms resource while getting object: %v", err)
	}

	_, err = c.CreateSystemGlobalMcPolicyDisabledAdoms(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemGlobalMcPolicyDisabledAdoms resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "adom_name")))

	return resourceSystemGlobalMcPolicyDisabledAdomsRead(d, m)
}

func resourceSystemGlobalMcPolicyDisabledAdomsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemGlobalMcPolicyDisabledAdoms(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobalMcPolicyDisabledAdoms resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGlobalMcPolicyDisabledAdoms(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobalMcPolicyDisabledAdoms resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "adom_name")))

	return resourceSystemGlobalMcPolicyDisabledAdomsRead(d, m)
}

func resourceSystemGlobalMcPolicyDisabledAdomsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemGlobalMcPolicyDisabledAdoms(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGlobalMcPolicyDisabledAdoms resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGlobalMcPolicyDisabledAdomsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemGlobalMcPolicyDisabledAdoms(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobalMcPolicyDisabledAdoms resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGlobalMcPolicyDisabledAdoms(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobalMcPolicyDisabledAdoms resource from API: %v", err)
	}
	return nil
}

func flattenSystemGlobalMcPolicyDisabledAdomsAdomNameSgma(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGlobalMcPolicyDisabledAdoms(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("adom_name", flattenSystemGlobalMcPolicyDisabledAdomsAdomNameSgma(o["adom-name"], d, "adom_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-name"], "SystemGlobalMcPolicyDisabledAdoms-AdomName"); ok {
			if err = d.Set("adom_name", vv); err != nil {
				return fmt.Errorf("Error reading adom_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_name: %v", err)
		}
	}

	return nil
}

func flattenSystemGlobalMcPolicyDisabledAdomsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGlobalMcPolicyDisabledAdomsAdomNameSgma(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGlobalMcPolicyDisabledAdoms(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adom_name"); ok {
		t, err := expandSystemGlobalMcPolicyDisabledAdomsAdomNameSgma(d, v, "adom_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-name"] = t
		}
	}

	return &obj, nil
}
