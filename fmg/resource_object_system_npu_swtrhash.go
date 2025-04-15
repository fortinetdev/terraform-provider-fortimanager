// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure switch traditional hashing.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuSwTrHash() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuSwTrHashUpdate,
		Read:   resourceObjectSystemNpuSwTrHashRead,
		Update: resourceObjectSystemNpuSwTrHashUpdate,
		Delete: resourceObjectSystemNpuSwTrHashDelete,

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
			"draco15": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_udp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuSwTrHashUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuSwTrHash(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuSwTrHash resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpuSwTrHash(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuSwTrHash resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuSwTrHash")

	return resourceObjectSystemNpuSwTrHashRead(d, m)
}

func resourceObjectSystemNpuSwTrHashDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuSwTrHash(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuSwTrHash resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuSwTrHashRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuSwTrHash(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuSwTrHash resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuSwTrHash(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuSwTrHash resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuSwTrHashDraco152edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuSwTrHashTcpUdpPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuSwTrHash(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("draco15", flattenObjectSystemNpuSwTrHashDraco152edl(o["draco15"], d, "draco15")); err != nil {
		if vv, ok := fortiAPIPatch(o["draco15"], "ObjectSystemNpuSwTrHash-Draco15"); ok {
			if err = d.Set("draco15", vv); err != nil {
				return fmt.Errorf("Error reading draco15: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading draco15: %v", err)
		}
	}

	if err = d.Set("tcp_udp_port", flattenObjectSystemNpuSwTrHashTcpUdpPort2edl(o["tcp-udp-port"], d, "tcp_udp_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-udp-port"], "ObjectSystemNpuSwTrHash-TcpUdpPort"); ok {
			if err = d.Set("tcp_udp_port", vv); err != nil {
				return fmt.Errorf("Error reading tcp_udp_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_udp_port: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuSwTrHashFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuSwTrHashDraco152edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuSwTrHashTcpUdpPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuSwTrHash(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("draco15"); ok || d.HasChange("draco15") {
		t, err := expandObjectSystemNpuSwTrHashDraco152edl(d, v, "draco15")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["draco15"] = t
		}
	}

	if v, ok := d.GetOk("tcp_udp_port"); ok || d.HasChange("tcp_udp_port") {
		t, err := expandObjectSystemNpuSwTrHashTcpUdpPort2edl(d, v, "tcp_udp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-udp-port"] = t
		}
	}

	return &obj, nil
}
