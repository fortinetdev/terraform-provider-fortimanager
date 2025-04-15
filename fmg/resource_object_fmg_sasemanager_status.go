// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFmg SaseManagerStatus

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFmgSaseManagerStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFmgSaseManagerStatusUpdate,
		Read:   resourceObjectFmgSaseManagerStatusRead,
		Update: resourceObjectFmgSaseManagerStatusUpdate,
		Delete: resourceObjectFmgSaseManagerStatusDelete,

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
			"forticlient_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forticloud_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spa_hubs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectFmgSaseManagerStatusUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFmgSaseManagerStatus(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgSaseManagerStatus resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFmgSaseManagerStatus(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgSaseManagerStatus resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFmgSaseManagerStatus")

	return resourceObjectFmgSaseManagerStatusRead(d, m)
}

func resourceObjectFmgSaseManagerStatusDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFmgSaseManagerStatus(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFmgSaseManagerStatus resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFmgSaseManagerStatusRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFmgSaseManagerStatus(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgSaseManagerStatus resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFmgSaseManagerStatus(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgSaseManagerStatus resource from API: %v", err)
	}
	return nil
}

func flattenObjectFmgSaseManagerStatusForticlientVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgSaseManagerStatusForticloudId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgSaseManagerStatusLicenseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgSaseManagerStatusSpaHubs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFmgSaseManagerStatus(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("forticlient_ver", flattenObjectFmgSaseManagerStatusForticlientVer(o["forticlient-ver"], d, "forticlient_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-ver"], "ObjectFmgSaseManagerStatus-ForticlientVer"); ok {
			if err = d.Set("forticlient_ver", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_ver: %v", err)
		}
	}

	if err = d.Set("forticloud_id", flattenObjectFmgSaseManagerStatusForticloudId(o["forticloud-id"], d, "forticloud_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticloud-id"], "ObjectFmgSaseManagerStatus-ForticloudId"); ok {
			if err = d.Set("forticloud_id", vv); err != nil {
				return fmt.Errorf("Error reading forticloud_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticloud_id: %v", err)
		}
	}

	if err = d.Set("license_type", flattenObjectFmgSaseManagerStatusLicenseType(o["license-type"], d, "license_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["license-type"], "ObjectFmgSaseManagerStatus-LicenseType"); ok {
			if err = d.Set("license_type", vv); err != nil {
				return fmt.Errorf("Error reading license_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading license_type: %v", err)
		}
	}

	if err = d.Set("spa_hubs", flattenObjectFmgSaseManagerStatusSpaHubs(o["spa-hubs"], d, "spa_hubs")); err != nil {
		if vv, ok := fortiAPIPatch(o["spa-hubs"], "ObjectFmgSaseManagerStatus-SpaHubs"); ok {
			if err = d.Set("spa_hubs", vv); err != nil {
				return fmt.Errorf("Error reading spa_hubs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spa_hubs: %v", err)
		}
	}

	return nil
}

func flattenObjectFmgSaseManagerStatusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFmgSaseManagerStatusForticlientVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgSaseManagerStatusForticloudId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgSaseManagerStatusLicenseType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgSaseManagerStatusSpaHubs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFmgSaseManagerStatus(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("forticlient_ver"); ok || d.HasChange("forticlient_ver") {
		t, err := expandObjectFmgSaseManagerStatusForticlientVer(d, v, "forticlient_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-ver"] = t
		}
	}

	if v, ok := d.GetOk("forticloud_id"); ok || d.HasChange("forticloud_id") {
		t, err := expandObjectFmgSaseManagerStatusForticloudId(d, v, "forticloud_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticloud-id"] = t
		}
	}

	if v, ok := d.GetOk("license_type"); ok || d.HasChange("license_type") {
		t, err := expandObjectFmgSaseManagerStatusLicenseType(d, v, "license_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license-type"] = t
		}
	}

	if v, ok := d.GetOk("spa_hubs"); ok || d.HasChange("spa_hubs") {
		t, err := expandObjectFmgSaseManagerStatusSpaHubs(d, v, "spa_hubs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spa-hubs"] = t
		}
	}

	return &obj, nil
}
