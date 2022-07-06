// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Move and/or rename a policy package within the same ADOM.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSecurityconsolePackageMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsolePackageMoveUpdate,
		Read:   resourceSecurityconsolePackageMoveRead,
		Update: resourceSecurityconsolePackageMoveUpdate,
		Delete: resourceSecurityconsolePackageMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_parent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSecurityconsolePackageMoveUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectSecurityconsolePackageMove(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsolePackageMove resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsolePackageMove(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsolePackageMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsolePackageMove")

	return resourceSecurityconsolePackageMoveRead(d, m)
}

func resourceSecurityconsolePackageMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsolePackageMoveRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsolePackageMoveAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsolePackageMoveDstName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsolePackageMoveDstParent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsolePackageMovePkg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSecurityconsolePackageMove(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenSecurityconsolePackageMoveAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsolePackageMove-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("dst_name", flattenSecurityconsolePackageMoveDstName(o["dst_name"], d, "dst_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst_name"], "SecurityconsolePackageMove-DstName"); ok {
			if err = d.Set("dst_name", vv); err != nil {
				return fmt.Errorf("Error reading dst_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_name: %v", err)
		}
	}

	if err = d.Set("dst_parent", flattenSecurityconsolePackageMoveDstParent(o["dst_parent"], d, "dst_parent")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst_parent"], "SecurityconsolePackageMove-DstParent"); ok {
			if err = d.Set("dst_parent", vv); err != nil {
				return fmt.Errorf("Error reading dst_parent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_parent: %v", err)
		}
	}

	if err = d.Set("pkg", flattenSecurityconsolePackageMovePkg(o["pkg"], d, "pkg")); err != nil {
		if vv, ok := fortiAPIPatch(o["pkg"], "SecurityconsolePackageMove-Pkg"); ok {
			if err = d.Set("pkg", vv); err != nil {
				return fmt.Errorf("Error reading pkg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pkg: %v", err)
		}
	}

	return nil
}

func flattenSecurityconsolePackageMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsolePackageMoveAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsolePackageMoveDstName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsolePackageMoveDstParent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsolePackageMovePkg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityconsolePackageMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("adom") {
		t, err := expandSecurityconsolePackageMoveAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("dst_name"); ok || d.HasChange("dst_name") {
		t, err := expandSecurityconsolePackageMoveDstName(d, v, "dst_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst_name"] = t
		}
	}

	if v, ok := d.GetOk("dst_parent"); ok || d.HasChange("dst_parent") {
		t, err := expandSecurityconsolePackageMoveDstParent(d, v, "dst_parent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst_parent"] = t
		}
	}

	if v, ok := d.GetOk("pkg"); ok || d.HasChange("pkg") {
		t, err := expandSecurityconsolePackageMovePkg(d, v, "pkg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pkg"] = t
		}
	}

	return &obj, nil
}
