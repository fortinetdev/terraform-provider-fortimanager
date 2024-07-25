// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the rate of ICMP messages generated by this FortiGate.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuIcmpRateCtrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuIcmpRateCtrlUpdate,
		Read:   resourceObjectSystemNpuIcmpRateCtrlRead,
		Update: resourceObjectSystemNpuIcmpRateCtrlUpdate,
		Delete: resourceObjectSystemNpuIcmpRateCtrlDelete,

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
			"icmp_v4_bucket_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"icmp_v4_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmp_v6_bucket_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmp_v6_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuIcmpRateCtrlUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuIcmpRateCtrl(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuIcmpRateCtrl resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuIcmpRateCtrl(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuIcmpRateCtrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuIcmpRateCtrl")

	return resourceObjectSystemNpuIcmpRateCtrlRead(d, m)
}

func resourceObjectSystemNpuIcmpRateCtrlDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuIcmpRateCtrl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuIcmpRateCtrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuIcmpRateCtrlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuIcmpRateCtrl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuIcmpRateCtrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuIcmpRateCtrl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuIcmpRateCtrl resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuIcmpRateCtrlIcmpV4BucketSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpRateCtrlIcmpV4Rate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpRateCtrlIcmpV6BucketSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuIcmpRateCtrlIcmpV6Rate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuIcmpRateCtrl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("icmp_v4_bucket_size", flattenObjectSystemNpuIcmpRateCtrlIcmpV4BucketSize2edl(o["icmp-v4-bucket-size"], d, "icmp_v4_bucket_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-v4-bucket-size"], "ObjectSystemNpuIcmpRateCtrl-IcmpV4BucketSize"); ok {
			if err = d.Set("icmp_v4_bucket_size", vv); err != nil {
				return fmt.Errorf("Error reading icmp_v4_bucket_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_v4_bucket_size: %v", err)
		}
	}

	if err = d.Set("icmp_v4_rate", flattenObjectSystemNpuIcmpRateCtrlIcmpV4Rate2edl(o["icmp-v4-rate"], d, "icmp_v4_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-v4-rate"], "ObjectSystemNpuIcmpRateCtrl-IcmpV4Rate"); ok {
			if err = d.Set("icmp_v4_rate", vv); err != nil {
				return fmt.Errorf("Error reading icmp_v4_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_v4_rate: %v", err)
		}
	}

	if err = d.Set("icmp_v6_bucket_size", flattenObjectSystemNpuIcmpRateCtrlIcmpV6BucketSize2edl(o["icmp-v6-bucket-size"], d, "icmp_v6_bucket_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-v6-bucket-size"], "ObjectSystemNpuIcmpRateCtrl-IcmpV6BucketSize"); ok {
			if err = d.Set("icmp_v6_bucket_size", vv); err != nil {
				return fmt.Errorf("Error reading icmp_v6_bucket_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_v6_bucket_size: %v", err)
		}
	}

	if err = d.Set("icmp_v6_rate", flattenObjectSystemNpuIcmpRateCtrlIcmpV6Rate2edl(o["icmp-v6-rate"], d, "icmp_v6_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["icmp-v6-rate"], "ObjectSystemNpuIcmpRateCtrl-IcmpV6Rate"); ok {
			if err = d.Set("icmp_v6_rate", vv); err != nil {
				return fmt.Errorf("Error reading icmp_v6_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icmp_v6_rate: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuIcmpRateCtrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuIcmpRateCtrlIcmpV4BucketSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpRateCtrlIcmpV4Rate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpRateCtrlIcmpV6BucketSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuIcmpRateCtrlIcmpV6Rate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuIcmpRateCtrl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("icmp_v4_bucket_size"); ok || d.HasChange("icmp_v4_bucket_size") {
		t, err := expandObjectSystemNpuIcmpRateCtrlIcmpV4BucketSize2edl(d, v, "icmp_v4_bucket_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-v4-bucket-size"] = t
		}
	}

	if v, ok := d.GetOk("icmp_v4_rate"); ok || d.HasChange("icmp_v4_rate") {
		t, err := expandObjectSystemNpuIcmpRateCtrlIcmpV4Rate2edl(d, v, "icmp_v4_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-v4-rate"] = t
		}
	}

	if v, ok := d.GetOk("icmp_v6_bucket_size"); ok || d.HasChange("icmp_v6_bucket_size") {
		t, err := expandObjectSystemNpuIcmpRateCtrlIcmpV6BucketSize2edl(d, v, "icmp_v6_bucket_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-v6-bucket-size"] = t
		}
	}

	if v, ok := d.GetOk("icmp_v6_rate"); ok || d.HasChange("icmp_v6_rate") {
		t, err := expandObjectSystemNpuIcmpRateCtrlIcmpV6Rate2edl(d, v, "icmp_v6_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-v6-rate"] = t
		}
	}

	return &obj, nil
}