// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Traffic from selected source or destination IP addresses is exempt from this signature.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectGlobalIpsSensorEntriesExemptIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGlobalIpsSensorEntriesExemptIpCreate,
		Read:   resourceObjectGlobalIpsSensorEntriesExemptIpRead,
		Update: resourceObjectGlobalIpsSensorEntriesExemptIpUpdate,
		Delete: resourceObjectGlobalIpsSensorEntriesExemptIpDelete,

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
			"sensor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dst_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectGlobalIpsSensorEntriesExemptIpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	obj, err := getObjectObjectGlobalIpsSensorEntriesExemptIp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGlobalIpsSensorEntriesExemptIp resource while getting object: %v", err)
	}

	_, err = c.CreateObjectGlobalIpsSensorEntriesExemptIp(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectGlobalIpsSensorEntriesExemptIp resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectGlobalIpsSensorEntriesExemptIpRead(d, m)
}

func resourceObjectGlobalIpsSensorEntriesExemptIpUpdate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	obj, err := getObjectObjectGlobalIpsSensorEntriesExemptIp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGlobalIpsSensorEntriesExemptIp resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectGlobalIpsSensorEntriesExemptIp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGlobalIpsSensorEntriesExemptIp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectGlobalIpsSensorEntriesExemptIpRead(d, m)
}

func resourceObjectGlobalIpsSensorEntriesExemptIpDelete(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	err = c.DeleteObjectGlobalIpsSensorEntriesExemptIp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGlobalIpsSensorEntriesExemptIp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGlobalIpsSensorEntriesExemptIpRead(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	if sensor == "" {
		sensor = importOptionChecking(m.(*FortiClient).Cfg, "sensor")
		if sensor == "" {
			return fmt.Errorf("Parameter sensor is missing")
		}
		if err = d.Set("sensor", sensor); err != nil {
			return fmt.Errorf("Error set params sensor: %v", err)
		}
	}
	if entries == "" {
		entries = importOptionChecking(m.(*FortiClient).Cfg, "entries")
		if entries == "" {
			return fmt.Errorf("Parameter entries is missing")
		}
		if err = d.Set("entries", entries); err != nil {
			return fmt.Errorf("Error set params entries: %v", err)
		}
	}
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	o, err := c.ReadObjectGlobalIpsSensorEntriesExemptIp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGlobalIpsSensorEntriesExemptIp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGlobalIpsSensorEntriesExemptIp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGlobalIpsSensorEntriesExemptIp resource from API: %v", err)
	}
	return nil
}

func flattenObjectGlobalIpsSensorEntriesExemptIpDstIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectGlobalIpsSensorEntriesExemptIpId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGlobalIpsSensorEntriesExemptIpSrcIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectGlobalIpsSensorEntriesExemptIp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dst_ip", flattenObjectGlobalIpsSensorEntriesExemptIpDstIp3rdl(o["dst-ip"], d, "dst_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-ip"], "ObjectGlobalIpsSensorEntriesExemptIp-DstIp"); ok {
			if err = d.Set("dst_ip", vv); err != nil {
				return fmt.Errorf("Error reading dst_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectGlobalIpsSensorEntriesExemptIpId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectGlobalIpsSensorEntriesExemptIp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("src_ip", flattenObjectGlobalIpsSensorEntriesExemptIpSrcIp3rdl(o["src-ip"], d, "src_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-ip"], "ObjectGlobalIpsSensorEntriesExemptIp-SrcIp"); ok {
			if err = d.Set("src_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_ip: %v", err)
		}
	}

	return nil
}

func flattenObjectGlobalIpsSensorEntriesExemptIpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGlobalIpsSensorEntriesExemptIpDstIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectGlobalIpsSensorEntriesExemptIpId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGlobalIpsSensorEntriesExemptIpSrcIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectGlobalIpsSensorEntriesExemptIp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst_ip"); ok || d.HasChange("dst_ip") {
		t, err := expandObjectGlobalIpsSensorEntriesExemptIpDstIp3rdl(d, v, "dst_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectGlobalIpsSensorEntriesExemptIpId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("src_ip"); ok || d.HasChange("src_ip") {
		t, err := expandObjectGlobalIpsSensorEntriesExemptIpSrcIp3rdl(d, v, "src_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip"] = t
		}
	}

	return &obj, nil
}
