// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Securityconsole PblockClone

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSecurityconsolePblockClone() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsolePblockCloneUpdate,
		Read:   resourceSecurityconsolePblockCloneRead,
		Update: resourceSecurityconsolePblockCloneUpdate,
		Delete: resourceSecurityconsolePblockCloneDelete,

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
			},
			"pblock": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSecurityconsolePblockCloneUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSecurityconsolePblockClone(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsolePblockClone resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsolePblockClone(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsolePblockClone resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsolePblockClone")

	return resourceSecurityconsolePblockCloneRead(d, m)
}

func resourceSecurityconsolePblockCloneDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsolePblockCloneRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsolePblockCloneAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsolePblockCloneDstName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsolePblockClonePblock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectSecurityconsolePblockClone(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenSecurityconsolePblockCloneAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsolePblockClone-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("dst_name", flattenSecurityconsolePblockCloneDstName(o["dst_name"], d, "dst_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst_name"], "SecurityconsolePblockClone-DstName"); ok {
			if err = d.Set("dst_name", vv); err != nil {
				return fmt.Errorf("Error reading dst_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_name: %v", err)
		}
	}

	if err = d.Set("pblock", flattenSecurityconsolePblockClonePblock(o["pblock"], d, "pblock")); err != nil {
		if vv, ok := fortiAPIPatch(o["pblock"], "SecurityconsolePblockClone-Pblock"); ok {
			if err = d.Set("pblock", vv); err != nil {
				return fmt.Errorf("Error reading pblock: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pblock: %v", err)
		}
	}

	return nil
}

func flattenSecurityconsolePblockCloneFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsolePblockCloneAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsolePblockCloneDstName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsolePblockClonePblock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectSecurityconsolePblockClone(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandSecurityconsolePblockCloneAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("dst_name"); ok || d.HasChange("dst_name") {
		t, err := expandSecurityconsolePblockCloneDstName(d, v, "dst_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst_name"] = t
		}
	}

	if v, ok := d.GetOk("pblock"); ok || d.HasChange("pblock") {
		t, err := expandSecurityconsolePblockClonePblock(d, v, "pblock")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pblock"] = t
		}
	}

	return &obj, nil
}
