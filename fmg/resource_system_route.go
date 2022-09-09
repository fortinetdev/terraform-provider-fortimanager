// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Routing table configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemRouteCreate,
		Read:   resourceSystemRouteRead,
		Update: resourceSystemRouteUpdate,
		Delete: resourceSystemRouteDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemRouteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemRoute(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemRoute resource while getting object: %v", err)
	}

	_, err = c.CreateSystemRoute(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemRoute resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceSystemRouteRead(d, m)
}

func resourceSystemRouteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemRoute(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemRoute resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemRoute(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemRoute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seq_num")))

	return resourceSystemRouteRead(d, m)
}

func resourceSystemRouteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemRoute(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemRoute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemRouteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemRoute(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemRoute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemRoute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemRoute resource from API: %v", err)
	}
	return nil
}

func flattenSystemRouteDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemRouteDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemRouteGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemRouteSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemRoute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device", flattenSystemRouteDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemRoute-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemRouteDst(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SystemRoute-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemRouteGateway(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "SystemRoute-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("seq_num", flattenSystemRouteSeqNum(o["seq_num"], d, "seq_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq_num"], "SystemRoute-SeqNum"); ok {
			if err = d.Set("seq_num", vv); err != nil {
				return fmt.Errorf("Error reading seq_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	return nil
}

func flattenSystemRouteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemRouteDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemRouteDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemRouteGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemRouteSeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemRoute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemRouteDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemRouteDst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandSystemRouteGateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("seq_num"); ok || d.HasChange("seq_num") {
		t, err := expandSystemRouteSeqNum(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq_num"] = t
		}
	}

	return &obj, nil
}
