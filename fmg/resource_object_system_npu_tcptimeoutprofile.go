// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure TCP timeout profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuTcpTimeoutProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuTcpTimeoutProfileCreate,
		Read:   resourceObjectSystemNpuTcpTimeoutProfileRead,
		Update: resourceObjectSystemNpuTcpTimeoutProfileUpdate,
		Delete: resourceObjectSystemNpuTcpTimeoutProfileDelete,

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
			"close_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fin_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"syn_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"syn_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"time_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuTcpTimeoutProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuTcpTimeoutProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuTcpTimeoutProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuTcpTimeoutProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuTcpTimeoutProfile resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemNpuTcpTimeoutProfileRead(d, m)
}

func resourceObjectSystemNpuTcpTimeoutProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuTcpTimeoutProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuTcpTimeoutProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuTcpTimeoutProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuTcpTimeoutProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemNpuTcpTimeoutProfileRead(d, m)
}

func resourceObjectSystemNpuTcpTimeoutProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuTcpTimeoutProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuTcpTimeoutProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuTcpTimeoutProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuTcpTimeoutProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuTcpTimeoutProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuTcpTimeoutProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuTcpTimeoutProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuTcpTimeoutProfileCloseWait2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileFinWait2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileSynSent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileSynWait2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileTcpIdle2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuTcpTimeoutProfileTimeWait2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuTcpTimeoutProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("close_wait", flattenObjectSystemNpuTcpTimeoutProfileCloseWait2edl(o["close-wait"], d, "close_wait")); err != nil {
		if vv, ok := fortiAPIPatch(o["close-wait"], "ObjectSystemNpuTcpTimeoutProfile-CloseWait"); ok {
			if err = d.Set("close_wait", vv); err != nil {
				return fmt.Errorf("Error reading close_wait: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading close_wait: %v", err)
		}
	}

	if err = d.Set("fin_wait", flattenObjectSystemNpuTcpTimeoutProfileFinWait2edl(o["fin-wait"], d, "fin_wait")); err != nil {
		if vv, ok := fortiAPIPatch(o["fin-wait"], "ObjectSystemNpuTcpTimeoutProfile-FinWait"); ok {
			if err = d.Set("fin_wait", vv); err != nil {
				return fmt.Errorf("Error reading fin_wait: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fin_wait: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemNpuTcpTimeoutProfileId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemNpuTcpTimeoutProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("syn_sent", flattenObjectSystemNpuTcpTimeoutProfileSynSent2edl(o["syn-sent"], d, "syn_sent")); err != nil {
		if vv, ok := fortiAPIPatch(o["syn-sent"], "ObjectSystemNpuTcpTimeoutProfile-SynSent"); ok {
			if err = d.Set("syn_sent", vv); err != nil {
				return fmt.Errorf("Error reading syn_sent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syn_sent: %v", err)
		}
	}

	if err = d.Set("syn_wait", flattenObjectSystemNpuTcpTimeoutProfileSynWait2edl(o["syn-wait"], d, "syn_wait")); err != nil {
		if vv, ok := fortiAPIPatch(o["syn-wait"], "ObjectSystemNpuTcpTimeoutProfile-SynWait"); ok {
			if err = d.Set("syn_wait", vv); err != nil {
				return fmt.Errorf("Error reading syn_wait: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syn_wait: %v", err)
		}
	}

	if err = d.Set("tcp_idle", flattenObjectSystemNpuTcpTimeoutProfileTcpIdle2edl(o["tcp-idle"], d, "tcp_idle")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-idle"], "ObjectSystemNpuTcpTimeoutProfile-TcpIdle"); ok {
			if err = d.Set("tcp_idle", vv); err != nil {
				return fmt.Errorf("Error reading tcp_idle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_idle: %v", err)
		}
	}

	if err = d.Set("time_wait", flattenObjectSystemNpuTcpTimeoutProfileTimeWait2edl(o["time-wait"], d, "time_wait")); err != nil {
		if vv, ok := fortiAPIPatch(o["time-wait"], "ObjectSystemNpuTcpTimeoutProfile-TimeWait"); ok {
			if err = d.Set("time_wait", vv); err != nil {
				return fmt.Errorf("Error reading time_wait: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time_wait: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuTcpTimeoutProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuTcpTimeoutProfileCloseWait2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileFinWait2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileSynSent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileSynWait2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileTcpIdle2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuTcpTimeoutProfileTimeWait2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuTcpTimeoutProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("close_wait"); ok || d.HasChange("close_wait") {
		t, err := expandObjectSystemNpuTcpTimeoutProfileCloseWait2edl(d, v, "close_wait")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["close-wait"] = t
		}
	}

	if v, ok := d.GetOk("fin_wait"); ok || d.HasChange("fin_wait") {
		t, err := expandObjectSystemNpuTcpTimeoutProfileFinWait2edl(d, v, "fin_wait")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fin-wait"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSystemNpuTcpTimeoutProfileId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("syn_sent"); ok || d.HasChange("syn_sent") {
		t, err := expandObjectSystemNpuTcpTimeoutProfileSynSent2edl(d, v, "syn_sent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syn-sent"] = t
		}
	}

	if v, ok := d.GetOk("syn_wait"); ok || d.HasChange("syn_wait") {
		t, err := expandObjectSystemNpuTcpTimeoutProfileSynWait2edl(d, v, "syn_wait")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syn-wait"] = t
		}
	}

	if v, ok := d.GetOk("tcp_idle"); ok || d.HasChange("tcp_idle") {
		t, err := expandObjectSystemNpuTcpTimeoutProfileTcpIdle2edl(d, v, "tcp_idle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-idle"] = t
		}
	}

	if v, ok := d.GetOk("time_wait"); ok || d.HasChange("time_wait") {
		t, err := expandObjectSystemNpuTcpTimeoutProfileTimeWait2edl(d, v, "time_wait")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time-wait"] = t
		}
	}

	return &obj, nil
}
