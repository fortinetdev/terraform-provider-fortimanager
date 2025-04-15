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

func resourceObjectIpsSensorEntriesExemptIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIpsSensorEntriesExemptIpCreate,
		Read:   resourceObjectIpsSensorEntriesExemptIpRead,
		Update: resourceObjectIpsSensorEntriesExemptIpUpdate,
		Delete: resourceObjectIpsSensorEntriesExemptIpDelete,

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

func resourceObjectIpsSensorEntriesExemptIpCreate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	obj, err := getObjectObjectIpsSensorEntriesExemptIp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIpsSensorEntriesExemptIp resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectIpsSensorEntriesExemptIp(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIpsSensorEntriesExemptIp resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIpsSensorEntriesExemptIpRead(d, m)
}

func resourceObjectIpsSensorEntriesExemptIpUpdate(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	obj, err := getObjectObjectIpsSensorEntriesExemptIp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIpsSensorEntriesExemptIp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectIpsSensorEntriesExemptIp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIpsSensorEntriesExemptIp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIpsSensorEntriesExemptIpRead(d, m)
}

func resourceObjectIpsSensorEntriesExemptIpDelete(d *schema.ResourceData, m interface{}) error {
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

	sensor := d.Get("sensor").(string)
	entries := d.Get("entries").(string)
	paradict["sensor"] = sensor
	paradict["entries"] = entries

	wsParams["adom"] = adomv

	err = c.DeleteObjectIpsSensorEntriesExemptIp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIpsSensorEntriesExemptIp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIpsSensorEntriesExemptIpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectIpsSensorEntriesExemptIp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIpsSensorEntriesExemptIp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIpsSensorEntriesExemptIp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIpsSensorEntriesExemptIp resource from API: %v", err)
	}
	return nil
}

func flattenObjectIpsSensorEntriesExemptIpDstIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectIpsSensorEntriesExemptIpId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIpsSensorEntriesExemptIpSrcIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectObjectIpsSensorEntriesExemptIp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dst_ip", flattenObjectIpsSensorEntriesExemptIpDstIp3rdl(o["dst-ip"], d, "dst_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-ip"], "ObjectIpsSensorEntriesExemptIp-DstIp"); ok {
			if err = d.Set("dst_ip", vv); err != nil {
				return fmt.Errorf("Error reading dst_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectIpsSensorEntriesExemptIpId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectIpsSensorEntriesExemptIp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("src_ip", flattenObjectIpsSensorEntriesExemptIpSrcIp3rdl(o["src-ip"], d, "src_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-ip"], "ObjectIpsSensorEntriesExemptIp-SrcIp"); ok {
			if err = d.Set("src_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_ip: %v", err)
		}
	}

	return nil
}

func flattenObjectIpsSensorEntriesExemptIpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIpsSensorEntriesExemptIpDstIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectIpsSensorEntriesExemptIpId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIpsSensorEntriesExemptIpSrcIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectObjectIpsSensorEntriesExemptIp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst_ip"); ok || d.HasChange("dst_ip") {
		t, err := expandObjectIpsSensorEntriesExemptIpDstIp3rdl(d, v, "dst_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectIpsSensorEntriesExemptIpId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("src_ip"); ok || d.HasChange("src_ip") {
		t, err := expandObjectIpsSensorEntriesExemptIpSrcIp3rdl(d, v, "src_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip"] = t
		}
	}

	return &obj, nil
}
