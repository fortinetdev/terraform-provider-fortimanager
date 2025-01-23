// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Reclaim management tunnel to device.  Request without device name specified will reclaim tunnels for all managed devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExecFgfmReclaimDevTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceExecFgfmReclaimDevTunnelUpdate,
		Read:   resourceExecFgfmReclaimDevTunnelRead,
		Update: resourceExecFgfmReclaimDevTunnelUpdate,
		Delete: resourceExecFgfmReclaimDevTunnelDelete,

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
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExecFgfmReclaimDevTunnelUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	device_name := d.Get("device_name").(string)
	paradict["device_name"] = device_name

	obj, err := getObjectExecFgfmReclaimDevTunnel(d)
	if err != nil {
		return fmt.Errorf("Error updating ExecFgfmReclaimDevTunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateExecFgfmReclaimDevTunnel(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ExecFgfmReclaimDevTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ExecFgfmReclaimDevTunnel")

	return resourceExecFgfmReclaimDevTunnelRead(d, m)
}

func resourceExecFgfmReclaimDevTunnelDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceExecFgfmReclaimDevTunnelRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenExecFgfmReclaimDevTunnelFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectExecFgfmReclaimDevTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("flags", flattenExecFgfmReclaimDevTunnelFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "ExecFgfmReclaimDevTunnel-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	return nil
}

func flattenExecFgfmReclaimDevTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExecFgfmReclaimDevTunnelFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectExecFgfmReclaimDevTunnel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandExecFgfmReclaimDevTunnelFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	return &obj, nil
}
