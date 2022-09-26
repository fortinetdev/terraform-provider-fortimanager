// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Generate install preview for a device.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecurityconsoleInstallPreview() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsoleInstallPreviewUpdate,
		Read:   resourceSecurityconsoleInstallPreviewRead,
		Update: resourceSecurityconsoleInstallPreviewUpdate,
		Delete: resourceSecurityconsoleInstallPreviewDelete,

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
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vdoms": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSecurityconsoleInstallPreviewUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectSecurityconsoleInstallPreview(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleInstallPreview resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsoleInstallPreview(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsoleInstallPreview resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsoleInstallPreview")

	return resourceSecurityconsoleInstallPreviewRead(d, m)
}

func resourceSecurityconsoleInstallPreviewDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsoleInstallPreviewRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsoleInstallPreviewAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleInstallPreviewDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsoleInstallPreviewFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSecurityconsoleInstallPreviewVdoms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSecurityconsoleInstallPreview(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenSecurityconsoleInstallPreviewAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsoleInstallPreview-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("device", flattenSecurityconsoleInstallPreviewDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SecurityconsoleInstallPreview-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("flags", flattenSecurityconsoleInstallPreviewFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "SecurityconsoleInstallPreview-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("vdoms", flattenSecurityconsoleInstallPreviewVdoms(o["vdoms"], d, "vdoms")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdoms"], "SecurityconsoleInstallPreview-Vdoms"); ok {
			if err = d.Set("vdoms", vv); err != nil {
				return fmt.Errorf("Error reading vdoms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdoms: %v", err)
		}
	}

	return nil
}

func flattenSecurityconsoleInstallPreviewFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsoleInstallPreviewAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleInstallPreviewDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsoleInstallPreviewFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSecurityconsoleInstallPreviewVdoms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSecurityconsoleInstallPreview(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandSecurityconsoleInstallPreviewAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSecurityconsoleInstallPreviewDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandSecurityconsoleInstallPreviewFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("vdoms"); ok || d.HasChange("vdoms") {
		t, err := expandSecurityconsoleInstallPreviewVdoms(d, v, "vdoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdoms"] = t
		}
	}

	return &obj, nil
}
