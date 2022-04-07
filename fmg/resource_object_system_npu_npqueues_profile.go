// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure a NP7 class profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectSystemNpuNpQueuesProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpQueuesProfileCreate,
		Read:   resourceObjectSystemNpuNpQueuesProfileRead,
		Update: resourceObjectSystemNpuNpQueuesProfileUpdate,
		Delete: resourceObjectSystemNpuNpQueuesProfileDelete,

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
			"cos0": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp0": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp11": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp13": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp14": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp15": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp16": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp17": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp18": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp19": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp20": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp21": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp22": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp23": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp24": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp25": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp26": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp27": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp28": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp29": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp30": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp31": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp32": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp33": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp34": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp35": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp36": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp37": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp38": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp39": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp40": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp41": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp42": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp43": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp44": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp45": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp47": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp48": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp49": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp50": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp51": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp52": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp53": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp54": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp55": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp56": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp57": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp58": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp59": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp60": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp61": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp62": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp63": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuNpQueuesProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuNpQueuesProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpQueuesProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuNpQueuesProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpQueuesProfile resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemNpuNpQueuesProfileRead(d, m)
}

func resourceObjectSystemNpuNpQueuesProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuNpQueuesProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueuesProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuNpQueuesProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueuesProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemNpuNpQueuesProfileRead(d, m)
}

func resourceObjectSystemNpuNpQueuesProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuNpQueuesProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpQueuesProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpQueuesProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuNpQueuesProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueuesProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpQueuesProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpQueuesProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpQueuesProfileCos0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp11(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp13(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp14(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp15(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp16(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp17(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp18(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp19(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp20(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp21(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp22(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp23(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp24(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp25(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp26(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp27(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp28(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp29(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp30(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp31(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp32(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp33(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp35(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp36(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp37(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp38(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp39(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp40(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp41(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp42(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp43(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp44(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp45(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp47(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp48(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp49(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp50(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp51(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp52(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp53(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp54(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp55(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp56(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp57(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp58(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp59(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp60(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp61(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp62(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp63(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpQueuesProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cos0", flattenObjectSystemNpuNpQueuesProfileCos0(o["cos0"], d, "cos0")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos0"], "ObjectSystemNpuNpQueuesProfile-Cos0"); ok {
			if err = d.Set("cos0", vv); err != nil {
				return fmt.Errorf("Error reading cos0: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos0: %v", err)
		}
	}

	if err = d.Set("cos1", flattenObjectSystemNpuNpQueuesProfileCos1(o["cos1"], d, "cos1")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos1"], "ObjectSystemNpuNpQueuesProfile-Cos1"); ok {
			if err = d.Set("cos1", vv); err != nil {
				return fmt.Errorf("Error reading cos1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos1: %v", err)
		}
	}

	if err = d.Set("cos2", flattenObjectSystemNpuNpQueuesProfileCos2(o["cos2"], d, "cos2")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos2"], "ObjectSystemNpuNpQueuesProfile-Cos2"); ok {
			if err = d.Set("cos2", vv); err != nil {
				return fmt.Errorf("Error reading cos2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos2: %v", err)
		}
	}

	if err = d.Set("cos3", flattenObjectSystemNpuNpQueuesProfileCos3(o["cos3"], d, "cos3")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos3"], "ObjectSystemNpuNpQueuesProfile-Cos3"); ok {
			if err = d.Set("cos3", vv); err != nil {
				return fmt.Errorf("Error reading cos3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos3: %v", err)
		}
	}

	if err = d.Set("cos4", flattenObjectSystemNpuNpQueuesProfileCos4(o["cos4"], d, "cos4")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos4"], "ObjectSystemNpuNpQueuesProfile-Cos4"); ok {
			if err = d.Set("cos4", vv); err != nil {
				return fmt.Errorf("Error reading cos4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos4: %v", err)
		}
	}

	if err = d.Set("cos5", flattenObjectSystemNpuNpQueuesProfileCos5(o["cos5"], d, "cos5")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos5"], "ObjectSystemNpuNpQueuesProfile-Cos5"); ok {
			if err = d.Set("cos5", vv); err != nil {
				return fmt.Errorf("Error reading cos5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos5: %v", err)
		}
	}

	if err = d.Set("cos6", flattenObjectSystemNpuNpQueuesProfileCos6(o["cos6"], d, "cos6")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos6"], "ObjectSystemNpuNpQueuesProfile-Cos6"); ok {
			if err = d.Set("cos6", vv); err != nil {
				return fmt.Errorf("Error reading cos6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos6: %v", err)
		}
	}

	if err = d.Set("cos7", flattenObjectSystemNpuNpQueuesProfileCos7(o["cos7"], d, "cos7")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos7"], "ObjectSystemNpuNpQueuesProfile-Cos7"); ok {
			if err = d.Set("cos7", vv); err != nil {
				return fmt.Errorf("Error reading cos7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos7: %v", err)
		}
	}

	if err = d.Set("dscp0", flattenObjectSystemNpuNpQueuesProfileDscp0(o["dscp0"], d, "dscp0")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp0"], "ObjectSystemNpuNpQueuesProfile-Dscp0"); ok {
			if err = d.Set("dscp0", vv); err != nil {
				return fmt.Errorf("Error reading dscp0: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp0: %v", err)
		}
	}

	if err = d.Set("dscp1", flattenObjectSystemNpuNpQueuesProfileDscp1(o["dscp1"], d, "dscp1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp1"], "ObjectSystemNpuNpQueuesProfile-Dscp1"); ok {
			if err = d.Set("dscp1", vv); err != nil {
				return fmt.Errorf("Error reading dscp1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp1: %v", err)
		}
	}

	if err = d.Set("dscp10", flattenObjectSystemNpuNpQueuesProfileDscp10(o["dscp10"], d, "dscp10")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp10"], "ObjectSystemNpuNpQueuesProfile-Dscp10"); ok {
			if err = d.Set("dscp10", vv); err != nil {
				return fmt.Errorf("Error reading dscp10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp10: %v", err)
		}
	}

	if err = d.Set("dscp11", flattenObjectSystemNpuNpQueuesProfileDscp11(o["dscp11"], d, "dscp11")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp11"], "ObjectSystemNpuNpQueuesProfile-Dscp11"); ok {
			if err = d.Set("dscp11", vv); err != nil {
				return fmt.Errorf("Error reading dscp11: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp11: %v", err)
		}
	}

	if err = d.Set("dscp12", flattenObjectSystemNpuNpQueuesProfileDscp12(o["dscp12"], d, "dscp12")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp12"], "ObjectSystemNpuNpQueuesProfile-Dscp12"); ok {
			if err = d.Set("dscp12", vv); err != nil {
				return fmt.Errorf("Error reading dscp12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp12: %v", err)
		}
	}

	if err = d.Set("dscp13", flattenObjectSystemNpuNpQueuesProfileDscp13(o["dscp13"], d, "dscp13")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp13"], "ObjectSystemNpuNpQueuesProfile-Dscp13"); ok {
			if err = d.Set("dscp13", vv); err != nil {
				return fmt.Errorf("Error reading dscp13: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp13: %v", err)
		}
	}

	if err = d.Set("dscp14", flattenObjectSystemNpuNpQueuesProfileDscp14(o["dscp14"], d, "dscp14")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp14"], "ObjectSystemNpuNpQueuesProfile-Dscp14"); ok {
			if err = d.Set("dscp14", vv); err != nil {
				return fmt.Errorf("Error reading dscp14: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp14: %v", err)
		}
	}

	if err = d.Set("dscp15", flattenObjectSystemNpuNpQueuesProfileDscp15(o["dscp15"], d, "dscp15")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp15"], "ObjectSystemNpuNpQueuesProfile-Dscp15"); ok {
			if err = d.Set("dscp15", vv); err != nil {
				return fmt.Errorf("Error reading dscp15: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp15: %v", err)
		}
	}

	if err = d.Set("dscp16", flattenObjectSystemNpuNpQueuesProfileDscp16(o["dscp16"], d, "dscp16")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp16"], "ObjectSystemNpuNpQueuesProfile-Dscp16"); ok {
			if err = d.Set("dscp16", vv); err != nil {
				return fmt.Errorf("Error reading dscp16: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp16: %v", err)
		}
	}

	if err = d.Set("dscp17", flattenObjectSystemNpuNpQueuesProfileDscp17(o["dscp17"], d, "dscp17")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp17"], "ObjectSystemNpuNpQueuesProfile-Dscp17"); ok {
			if err = d.Set("dscp17", vv); err != nil {
				return fmt.Errorf("Error reading dscp17: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp17: %v", err)
		}
	}

	if err = d.Set("dscp18", flattenObjectSystemNpuNpQueuesProfileDscp18(o["dscp18"], d, "dscp18")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp18"], "ObjectSystemNpuNpQueuesProfile-Dscp18"); ok {
			if err = d.Set("dscp18", vv); err != nil {
				return fmt.Errorf("Error reading dscp18: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp18: %v", err)
		}
	}

	if err = d.Set("dscp19", flattenObjectSystemNpuNpQueuesProfileDscp19(o["dscp19"], d, "dscp19")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp19"], "ObjectSystemNpuNpQueuesProfile-Dscp19"); ok {
			if err = d.Set("dscp19", vv); err != nil {
				return fmt.Errorf("Error reading dscp19: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp19: %v", err)
		}
	}

	if err = d.Set("dscp2", flattenObjectSystemNpuNpQueuesProfileDscp2(o["dscp2"], d, "dscp2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp2"], "ObjectSystemNpuNpQueuesProfile-Dscp2"); ok {
			if err = d.Set("dscp2", vv); err != nil {
				return fmt.Errorf("Error reading dscp2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp2: %v", err)
		}
	}

	if err = d.Set("dscp20", flattenObjectSystemNpuNpQueuesProfileDscp20(o["dscp20"], d, "dscp20")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp20"], "ObjectSystemNpuNpQueuesProfile-Dscp20"); ok {
			if err = d.Set("dscp20", vv); err != nil {
				return fmt.Errorf("Error reading dscp20: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp20: %v", err)
		}
	}

	if err = d.Set("dscp21", flattenObjectSystemNpuNpQueuesProfileDscp21(o["dscp21"], d, "dscp21")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp21"], "ObjectSystemNpuNpQueuesProfile-Dscp21"); ok {
			if err = d.Set("dscp21", vv); err != nil {
				return fmt.Errorf("Error reading dscp21: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp21: %v", err)
		}
	}

	if err = d.Set("dscp22", flattenObjectSystemNpuNpQueuesProfileDscp22(o["dscp22"], d, "dscp22")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp22"], "ObjectSystemNpuNpQueuesProfile-Dscp22"); ok {
			if err = d.Set("dscp22", vv); err != nil {
				return fmt.Errorf("Error reading dscp22: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp22: %v", err)
		}
	}

	if err = d.Set("dscp23", flattenObjectSystemNpuNpQueuesProfileDscp23(o["dscp23"], d, "dscp23")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp23"], "ObjectSystemNpuNpQueuesProfile-Dscp23"); ok {
			if err = d.Set("dscp23", vv); err != nil {
				return fmt.Errorf("Error reading dscp23: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp23: %v", err)
		}
	}

	if err = d.Set("dscp24", flattenObjectSystemNpuNpQueuesProfileDscp24(o["dscp24"], d, "dscp24")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp24"], "ObjectSystemNpuNpQueuesProfile-Dscp24"); ok {
			if err = d.Set("dscp24", vv); err != nil {
				return fmt.Errorf("Error reading dscp24: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp24: %v", err)
		}
	}

	if err = d.Set("dscp25", flattenObjectSystemNpuNpQueuesProfileDscp25(o["dscp25"], d, "dscp25")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp25"], "ObjectSystemNpuNpQueuesProfile-Dscp25"); ok {
			if err = d.Set("dscp25", vv); err != nil {
				return fmt.Errorf("Error reading dscp25: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp25: %v", err)
		}
	}

	if err = d.Set("dscp26", flattenObjectSystemNpuNpQueuesProfileDscp26(o["dscp26"], d, "dscp26")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp26"], "ObjectSystemNpuNpQueuesProfile-Dscp26"); ok {
			if err = d.Set("dscp26", vv); err != nil {
				return fmt.Errorf("Error reading dscp26: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp26: %v", err)
		}
	}

	if err = d.Set("dscp27", flattenObjectSystemNpuNpQueuesProfileDscp27(o["dscp27"], d, "dscp27")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp27"], "ObjectSystemNpuNpQueuesProfile-Dscp27"); ok {
			if err = d.Set("dscp27", vv); err != nil {
				return fmt.Errorf("Error reading dscp27: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp27: %v", err)
		}
	}

	if err = d.Set("dscp28", flattenObjectSystemNpuNpQueuesProfileDscp28(o["dscp28"], d, "dscp28")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp28"], "ObjectSystemNpuNpQueuesProfile-Dscp28"); ok {
			if err = d.Set("dscp28", vv); err != nil {
				return fmt.Errorf("Error reading dscp28: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp28: %v", err)
		}
	}

	if err = d.Set("dscp29", flattenObjectSystemNpuNpQueuesProfileDscp29(o["dscp29"], d, "dscp29")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp29"], "ObjectSystemNpuNpQueuesProfile-Dscp29"); ok {
			if err = d.Set("dscp29", vv); err != nil {
				return fmt.Errorf("Error reading dscp29: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp29: %v", err)
		}
	}

	if err = d.Set("dscp3", flattenObjectSystemNpuNpQueuesProfileDscp3(o["dscp3"], d, "dscp3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp3"], "ObjectSystemNpuNpQueuesProfile-Dscp3"); ok {
			if err = d.Set("dscp3", vv); err != nil {
				return fmt.Errorf("Error reading dscp3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp3: %v", err)
		}
	}

	if err = d.Set("dscp30", flattenObjectSystemNpuNpQueuesProfileDscp30(o["dscp30"], d, "dscp30")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp30"], "ObjectSystemNpuNpQueuesProfile-Dscp30"); ok {
			if err = d.Set("dscp30", vv); err != nil {
				return fmt.Errorf("Error reading dscp30: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp30: %v", err)
		}
	}

	if err = d.Set("dscp31", flattenObjectSystemNpuNpQueuesProfileDscp31(o["dscp31"], d, "dscp31")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp31"], "ObjectSystemNpuNpQueuesProfile-Dscp31"); ok {
			if err = d.Set("dscp31", vv); err != nil {
				return fmt.Errorf("Error reading dscp31: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp31: %v", err)
		}
	}

	if err = d.Set("dscp32", flattenObjectSystemNpuNpQueuesProfileDscp32(o["dscp32"], d, "dscp32")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp32"], "ObjectSystemNpuNpQueuesProfile-Dscp32"); ok {
			if err = d.Set("dscp32", vv); err != nil {
				return fmt.Errorf("Error reading dscp32: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp32: %v", err)
		}
	}

	if err = d.Set("dscp33", flattenObjectSystemNpuNpQueuesProfileDscp33(o["dscp33"], d, "dscp33")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp33"], "ObjectSystemNpuNpQueuesProfile-Dscp33"); ok {
			if err = d.Set("dscp33", vv); err != nil {
				return fmt.Errorf("Error reading dscp33: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp33: %v", err)
		}
	}

	if err = d.Set("dscp34", flattenObjectSystemNpuNpQueuesProfileDscp34(o["dscp34"], d, "dscp34")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp34"], "ObjectSystemNpuNpQueuesProfile-Dscp34"); ok {
			if err = d.Set("dscp34", vv); err != nil {
				return fmt.Errorf("Error reading dscp34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp34: %v", err)
		}
	}

	if err = d.Set("dscp35", flattenObjectSystemNpuNpQueuesProfileDscp35(o["dscp35"], d, "dscp35")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp35"], "ObjectSystemNpuNpQueuesProfile-Dscp35"); ok {
			if err = d.Set("dscp35", vv); err != nil {
				return fmt.Errorf("Error reading dscp35: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp35: %v", err)
		}
	}

	if err = d.Set("dscp36", flattenObjectSystemNpuNpQueuesProfileDscp36(o["dscp36"], d, "dscp36")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp36"], "ObjectSystemNpuNpQueuesProfile-Dscp36"); ok {
			if err = d.Set("dscp36", vv); err != nil {
				return fmt.Errorf("Error reading dscp36: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp36: %v", err)
		}
	}

	if err = d.Set("dscp37", flattenObjectSystemNpuNpQueuesProfileDscp37(o["dscp37"], d, "dscp37")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp37"], "ObjectSystemNpuNpQueuesProfile-Dscp37"); ok {
			if err = d.Set("dscp37", vv); err != nil {
				return fmt.Errorf("Error reading dscp37: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp37: %v", err)
		}
	}

	if err = d.Set("dscp38", flattenObjectSystemNpuNpQueuesProfileDscp38(o["dscp38"], d, "dscp38")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp38"], "ObjectSystemNpuNpQueuesProfile-Dscp38"); ok {
			if err = d.Set("dscp38", vv); err != nil {
				return fmt.Errorf("Error reading dscp38: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp38: %v", err)
		}
	}

	if err = d.Set("dscp39", flattenObjectSystemNpuNpQueuesProfileDscp39(o["dscp39"], d, "dscp39")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp39"], "ObjectSystemNpuNpQueuesProfile-Dscp39"); ok {
			if err = d.Set("dscp39", vv); err != nil {
				return fmt.Errorf("Error reading dscp39: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp39: %v", err)
		}
	}

	if err = d.Set("dscp4", flattenObjectSystemNpuNpQueuesProfileDscp4(o["dscp4"], d, "dscp4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp4"], "ObjectSystemNpuNpQueuesProfile-Dscp4"); ok {
			if err = d.Set("dscp4", vv); err != nil {
				return fmt.Errorf("Error reading dscp4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp4: %v", err)
		}
	}

	if err = d.Set("dscp40", flattenObjectSystemNpuNpQueuesProfileDscp40(o["dscp40"], d, "dscp40")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp40"], "ObjectSystemNpuNpQueuesProfile-Dscp40"); ok {
			if err = d.Set("dscp40", vv); err != nil {
				return fmt.Errorf("Error reading dscp40: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp40: %v", err)
		}
	}

	if err = d.Set("dscp41", flattenObjectSystemNpuNpQueuesProfileDscp41(o["dscp41"], d, "dscp41")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp41"], "ObjectSystemNpuNpQueuesProfile-Dscp41"); ok {
			if err = d.Set("dscp41", vv); err != nil {
				return fmt.Errorf("Error reading dscp41: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp41: %v", err)
		}
	}

	if err = d.Set("dscp42", flattenObjectSystemNpuNpQueuesProfileDscp42(o["dscp42"], d, "dscp42")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp42"], "ObjectSystemNpuNpQueuesProfile-Dscp42"); ok {
			if err = d.Set("dscp42", vv); err != nil {
				return fmt.Errorf("Error reading dscp42: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp42: %v", err)
		}
	}

	if err = d.Set("dscp43", flattenObjectSystemNpuNpQueuesProfileDscp43(o["dscp43"], d, "dscp43")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp43"], "ObjectSystemNpuNpQueuesProfile-Dscp43"); ok {
			if err = d.Set("dscp43", vv); err != nil {
				return fmt.Errorf("Error reading dscp43: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp43: %v", err)
		}
	}

	if err = d.Set("dscp44", flattenObjectSystemNpuNpQueuesProfileDscp44(o["dscp44"], d, "dscp44")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp44"], "ObjectSystemNpuNpQueuesProfile-Dscp44"); ok {
			if err = d.Set("dscp44", vv); err != nil {
				return fmt.Errorf("Error reading dscp44: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp44: %v", err)
		}
	}

	if err = d.Set("dscp45", flattenObjectSystemNpuNpQueuesProfileDscp45(o["dscp45"], d, "dscp45")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp45"], "ObjectSystemNpuNpQueuesProfile-Dscp45"); ok {
			if err = d.Set("dscp45", vv); err != nil {
				return fmt.Errorf("Error reading dscp45: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp45: %v", err)
		}
	}

	if err = d.Set("dscp46", flattenObjectSystemNpuNpQueuesProfileDscp46(o["dscp46"], d, "dscp46")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp46"], "ObjectSystemNpuNpQueuesProfile-Dscp46"); ok {
			if err = d.Set("dscp46", vv); err != nil {
				return fmt.Errorf("Error reading dscp46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp46: %v", err)
		}
	}

	if err = d.Set("dscp47", flattenObjectSystemNpuNpQueuesProfileDscp47(o["dscp47"], d, "dscp47")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp47"], "ObjectSystemNpuNpQueuesProfile-Dscp47"); ok {
			if err = d.Set("dscp47", vv); err != nil {
				return fmt.Errorf("Error reading dscp47: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp47: %v", err)
		}
	}

	if err = d.Set("dscp48", flattenObjectSystemNpuNpQueuesProfileDscp48(o["dscp48"], d, "dscp48")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp48"], "ObjectSystemNpuNpQueuesProfile-Dscp48"); ok {
			if err = d.Set("dscp48", vv); err != nil {
				return fmt.Errorf("Error reading dscp48: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp48: %v", err)
		}
	}

	if err = d.Set("dscp49", flattenObjectSystemNpuNpQueuesProfileDscp49(o["dscp49"], d, "dscp49")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp49"], "ObjectSystemNpuNpQueuesProfile-Dscp49"); ok {
			if err = d.Set("dscp49", vv); err != nil {
				return fmt.Errorf("Error reading dscp49: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp49: %v", err)
		}
	}

	if err = d.Set("dscp5", flattenObjectSystemNpuNpQueuesProfileDscp5(o["dscp5"], d, "dscp5")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp5"], "ObjectSystemNpuNpQueuesProfile-Dscp5"); ok {
			if err = d.Set("dscp5", vv); err != nil {
				return fmt.Errorf("Error reading dscp5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp5: %v", err)
		}
	}

	if err = d.Set("dscp50", flattenObjectSystemNpuNpQueuesProfileDscp50(o["dscp50"], d, "dscp50")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp50"], "ObjectSystemNpuNpQueuesProfile-Dscp50"); ok {
			if err = d.Set("dscp50", vv); err != nil {
				return fmt.Errorf("Error reading dscp50: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp50: %v", err)
		}
	}

	if err = d.Set("dscp51", flattenObjectSystemNpuNpQueuesProfileDscp51(o["dscp51"], d, "dscp51")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp51"], "ObjectSystemNpuNpQueuesProfile-Dscp51"); ok {
			if err = d.Set("dscp51", vv); err != nil {
				return fmt.Errorf("Error reading dscp51: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp51: %v", err)
		}
	}

	if err = d.Set("dscp52", flattenObjectSystemNpuNpQueuesProfileDscp52(o["dscp52"], d, "dscp52")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp52"], "ObjectSystemNpuNpQueuesProfile-Dscp52"); ok {
			if err = d.Set("dscp52", vv); err != nil {
				return fmt.Errorf("Error reading dscp52: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp52: %v", err)
		}
	}

	if err = d.Set("dscp53", flattenObjectSystemNpuNpQueuesProfileDscp53(o["dscp53"], d, "dscp53")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp53"], "ObjectSystemNpuNpQueuesProfile-Dscp53"); ok {
			if err = d.Set("dscp53", vv); err != nil {
				return fmt.Errorf("Error reading dscp53: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp53: %v", err)
		}
	}

	if err = d.Set("dscp54", flattenObjectSystemNpuNpQueuesProfileDscp54(o["dscp54"], d, "dscp54")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp54"], "ObjectSystemNpuNpQueuesProfile-Dscp54"); ok {
			if err = d.Set("dscp54", vv); err != nil {
				return fmt.Errorf("Error reading dscp54: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp54: %v", err)
		}
	}

	if err = d.Set("dscp55", flattenObjectSystemNpuNpQueuesProfileDscp55(o["dscp55"], d, "dscp55")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp55"], "ObjectSystemNpuNpQueuesProfile-Dscp55"); ok {
			if err = d.Set("dscp55", vv); err != nil {
				return fmt.Errorf("Error reading dscp55: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp55: %v", err)
		}
	}

	if err = d.Set("dscp56", flattenObjectSystemNpuNpQueuesProfileDscp56(o["dscp56"], d, "dscp56")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp56"], "ObjectSystemNpuNpQueuesProfile-Dscp56"); ok {
			if err = d.Set("dscp56", vv); err != nil {
				return fmt.Errorf("Error reading dscp56: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp56: %v", err)
		}
	}

	if err = d.Set("dscp57", flattenObjectSystemNpuNpQueuesProfileDscp57(o["dscp57"], d, "dscp57")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp57"], "ObjectSystemNpuNpQueuesProfile-Dscp57"); ok {
			if err = d.Set("dscp57", vv); err != nil {
				return fmt.Errorf("Error reading dscp57: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp57: %v", err)
		}
	}

	if err = d.Set("dscp58", flattenObjectSystemNpuNpQueuesProfileDscp58(o["dscp58"], d, "dscp58")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp58"], "ObjectSystemNpuNpQueuesProfile-Dscp58"); ok {
			if err = d.Set("dscp58", vv); err != nil {
				return fmt.Errorf("Error reading dscp58: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp58: %v", err)
		}
	}

	if err = d.Set("dscp59", flattenObjectSystemNpuNpQueuesProfileDscp59(o["dscp59"], d, "dscp59")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp59"], "ObjectSystemNpuNpQueuesProfile-Dscp59"); ok {
			if err = d.Set("dscp59", vv); err != nil {
				return fmt.Errorf("Error reading dscp59: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp59: %v", err)
		}
	}

	if err = d.Set("dscp6", flattenObjectSystemNpuNpQueuesProfileDscp6(o["dscp6"], d, "dscp6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp6"], "ObjectSystemNpuNpQueuesProfile-Dscp6"); ok {
			if err = d.Set("dscp6", vv); err != nil {
				return fmt.Errorf("Error reading dscp6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp6: %v", err)
		}
	}

	if err = d.Set("dscp60", flattenObjectSystemNpuNpQueuesProfileDscp60(o["dscp60"], d, "dscp60")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp60"], "ObjectSystemNpuNpQueuesProfile-Dscp60"); ok {
			if err = d.Set("dscp60", vv); err != nil {
				return fmt.Errorf("Error reading dscp60: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp60: %v", err)
		}
	}

	if err = d.Set("dscp61", flattenObjectSystemNpuNpQueuesProfileDscp61(o["dscp61"], d, "dscp61")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp61"], "ObjectSystemNpuNpQueuesProfile-Dscp61"); ok {
			if err = d.Set("dscp61", vv); err != nil {
				return fmt.Errorf("Error reading dscp61: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp61: %v", err)
		}
	}

	if err = d.Set("dscp62", flattenObjectSystemNpuNpQueuesProfileDscp62(o["dscp62"], d, "dscp62")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp62"], "ObjectSystemNpuNpQueuesProfile-Dscp62"); ok {
			if err = d.Set("dscp62", vv); err != nil {
				return fmt.Errorf("Error reading dscp62: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp62: %v", err)
		}
	}

	if err = d.Set("dscp63", flattenObjectSystemNpuNpQueuesProfileDscp63(o["dscp63"], d, "dscp63")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp63"], "ObjectSystemNpuNpQueuesProfile-Dscp63"); ok {
			if err = d.Set("dscp63", vv); err != nil {
				return fmt.Errorf("Error reading dscp63: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp63: %v", err)
		}
	}

	if err = d.Set("dscp7", flattenObjectSystemNpuNpQueuesProfileDscp7(o["dscp7"], d, "dscp7")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp7"], "ObjectSystemNpuNpQueuesProfile-Dscp7"); ok {
			if err = d.Set("dscp7", vv); err != nil {
				return fmt.Errorf("Error reading dscp7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp7: %v", err)
		}
	}

	if err = d.Set("dscp8", flattenObjectSystemNpuNpQueuesProfileDscp8(o["dscp8"], d, "dscp8")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp8"], "ObjectSystemNpuNpQueuesProfile-Dscp8"); ok {
			if err = d.Set("dscp8", vv); err != nil {
				return fmt.Errorf("Error reading dscp8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp8: %v", err)
		}
	}

	if err = d.Set("dscp9", flattenObjectSystemNpuNpQueuesProfileDscp9(o["dscp9"], d, "dscp9")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp9"], "ObjectSystemNpuNpQueuesProfile-Dscp9"); ok {
			if err = d.Set("dscp9", vv); err != nil {
				return fmt.Errorf("Error reading dscp9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp9: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemNpuNpQueuesProfileId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemNpuNpQueuesProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSystemNpuNpQueuesProfileType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemNpuNpQueuesProfile-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectSystemNpuNpQueuesProfileWeight(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectSystemNpuNpQueuesProfile-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpQueuesProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpQueuesProfileCos0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp11(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp13(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp14(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp15(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp17(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp18(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp19(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp20(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp21(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp22(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp23(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp24(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp25(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp26(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp27(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp28(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp29(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp30(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp31(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp32(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp33(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp35(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp36(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp37(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp38(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp39(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp40(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp41(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp42(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp43(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp44(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp45(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp46(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp47(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp48(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp49(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp50(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp51(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp52(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp53(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp54(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp55(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp56(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp57(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp58(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp59(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp60(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp61(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp62(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp63(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpQueuesProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cos0"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileCos0(d, v, "cos0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos0"] = t
		}
	}

	if v, ok := d.GetOk("cos1"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileCos1(d, v, "cos1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos1"] = t
		}
	}

	if v, ok := d.GetOk("cos2"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileCos2(d, v, "cos2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos2"] = t
		}
	}

	if v, ok := d.GetOk("cos3"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileCos3(d, v, "cos3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos3"] = t
		}
	}

	if v, ok := d.GetOk("cos4"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileCos4(d, v, "cos4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos4"] = t
		}
	}

	if v, ok := d.GetOk("cos5"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileCos5(d, v, "cos5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos5"] = t
		}
	}

	if v, ok := d.GetOk("cos6"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileCos6(d, v, "cos6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos6"] = t
		}
	}

	if v, ok := d.GetOk("cos7"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileCos7(d, v, "cos7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos7"] = t
		}
	}

	if v, ok := d.GetOk("dscp0"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp0(d, v, "dscp0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp0"] = t
		}
	}

	if v, ok := d.GetOk("dscp1"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp1(d, v, "dscp1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp1"] = t
		}
	}

	if v, ok := d.GetOk("dscp10"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp10(d, v, "dscp10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp10"] = t
		}
	}

	if v, ok := d.GetOk("dscp11"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp11(d, v, "dscp11")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp11"] = t
		}
	}

	if v, ok := d.GetOk("dscp12"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp12(d, v, "dscp12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp12"] = t
		}
	}

	if v, ok := d.GetOk("dscp13"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp13(d, v, "dscp13")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp13"] = t
		}
	}

	if v, ok := d.GetOk("dscp14"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp14(d, v, "dscp14")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp14"] = t
		}
	}

	if v, ok := d.GetOk("dscp15"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp15(d, v, "dscp15")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp15"] = t
		}
	}

	if v, ok := d.GetOk("dscp16"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp16(d, v, "dscp16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp16"] = t
		}
	}

	if v, ok := d.GetOk("dscp17"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp17(d, v, "dscp17")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp17"] = t
		}
	}

	if v, ok := d.GetOk("dscp18"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp18(d, v, "dscp18")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp18"] = t
		}
	}

	if v, ok := d.GetOk("dscp19"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp19(d, v, "dscp19")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp19"] = t
		}
	}

	if v, ok := d.GetOk("dscp2"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp2(d, v, "dscp2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp2"] = t
		}
	}

	if v, ok := d.GetOk("dscp20"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp20(d, v, "dscp20")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp20"] = t
		}
	}

	if v, ok := d.GetOk("dscp21"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp21(d, v, "dscp21")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp21"] = t
		}
	}

	if v, ok := d.GetOk("dscp22"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp22(d, v, "dscp22")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp22"] = t
		}
	}

	if v, ok := d.GetOk("dscp23"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp23(d, v, "dscp23")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp23"] = t
		}
	}

	if v, ok := d.GetOk("dscp24"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp24(d, v, "dscp24")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp24"] = t
		}
	}

	if v, ok := d.GetOk("dscp25"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp25(d, v, "dscp25")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp25"] = t
		}
	}

	if v, ok := d.GetOk("dscp26"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp26(d, v, "dscp26")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp26"] = t
		}
	}

	if v, ok := d.GetOk("dscp27"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp27(d, v, "dscp27")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp27"] = t
		}
	}

	if v, ok := d.GetOk("dscp28"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp28(d, v, "dscp28")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp28"] = t
		}
	}

	if v, ok := d.GetOk("dscp29"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp29(d, v, "dscp29")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp29"] = t
		}
	}

	if v, ok := d.GetOk("dscp3"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp3(d, v, "dscp3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp3"] = t
		}
	}

	if v, ok := d.GetOk("dscp30"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp30(d, v, "dscp30")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp30"] = t
		}
	}

	if v, ok := d.GetOk("dscp31"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp31(d, v, "dscp31")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp31"] = t
		}
	}

	if v, ok := d.GetOk("dscp32"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp32(d, v, "dscp32")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp32"] = t
		}
	}

	if v, ok := d.GetOk("dscp33"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp33(d, v, "dscp33")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp33"] = t
		}
	}

	if v, ok := d.GetOk("dscp34"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp34(d, v, "dscp34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp34"] = t
		}
	}

	if v, ok := d.GetOk("dscp35"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp35(d, v, "dscp35")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp35"] = t
		}
	}

	if v, ok := d.GetOk("dscp36"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp36(d, v, "dscp36")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp36"] = t
		}
	}

	if v, ok := d.GetOk("dscp37"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp37(d, v, "dscp37")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp37"] = t
		}
	}

	if v, ok := d.GetOk("dscp38"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp38(d, v, "dscp38")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp38"] = t
		}
	}

	if v, ok := d.GetOk("dscp39"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp39(d, v, "dscp39")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp39"] = t
		}
	}

	if v, ok := d.GetOk("dscp4"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp4(d, v, "dscp4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp4"] = t
		}
	}

	if v, ok := d.GetOk("dscp40"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp40(d, v, "dscp40")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp40"] = t
		}
	}

	if v, ok := d.GetOk("dscp41"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp41(d, v, "dscp41")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp41"] = t
		}
	}

	if v, ok := d.GetOk("dscp42"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp42(d, v, "dscp42")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp42"] = t
		}
	}

	if v, ok := d.GetOk("dscp43"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp43(d, v, "dscp43")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp43"] = t
		}
	}

	if v, ok := d.GetOk("dscp44"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp44(d, v, "dscp44")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp44"] = t
		}
	}

	if v, ok := d.GetOk("dscp45"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp45(d, v, "dscp45")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp45"] = t
		}
	}

	if v, ok := d.GetOk("dscp46"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp46(d, v, "dscp46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp46"] = t
		}
	}

	if v, ok := d.GetOk("dscp47"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp47(d, v, "dscp47")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp47"] = t
		}
	}

	if v, ok := d.GetOk("dscp48"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp48(d, v, "dscp48")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp48"] = t
		}
	}

	if v, ok := d.GetOk("dscp49"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp49(d, v, "dscp49")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp49"] = t
		}
	}

	if v, ok := d.GetOk("dscp5"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp5(d, v, "dscp5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp5"] = t
		}
	}

	if v, ok := d.GetOk("dscp50"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp50(d, v, "dscp50")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp50"] = t
		}
	}

	if v, ok := d.GetOk("dscp51"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp51(d, v, "dscp51")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp51"] = t
		}
	}

	if v, ok := d.GetOk("dscp52"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp52(d, v, "dscp52")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp52"] = t
		}
	}

	if v, ok := d.GetOk("dscp53"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp53(d, v, "dscp53")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp53"] = t
		}
	}

	if v, ok := d.GetOk("dscp54"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp54(d, v, "dscp54")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp54"] = t
		}
	}

	if v, ok := d.GetOk("dscp55"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp55(d, v, "dscp55")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp55"] = t
		}
	}

	if v, ok := d.GetOk("dscp56"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp56(d, v, "dscp56")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp56"] = t
		}
	}

	if v, ok := d.GetOk("dscp57"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp57(d, v, "dscp57")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp57"] = t
		}
	}

	if v, ok := d.GetOk("dscp58"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp58(d, v, "dscp58")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp58"] = t
		}
	}

	if v, ok := d.GetOk("dscp59"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp59(d, v, "dscp59")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp59"] = t
		}
	}

	if v, ok := d.GetOk("dscp6"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp6(d, v, "dscp6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp6"] = t
		}
	}

	if v, ok := d.GetOk("dscp60"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp60(d, v, "dscp60")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp60"] = t
		}
	}

	if v, ok := d.GetOk("dscp61"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp61(d, v, "dscp61")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp61"] = t
		}
	}

	if v, ok := d.GetOk("dscp62"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp62(d, v, "dscp62")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp62"] = t
		}
	}

	if v, ok := d.GetOk("dscp63"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp63(d, v, "dscp63")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp63"] = t
		}
	}

	if v, ok := d.GetOk("dscp7"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp7(d, v, "dscp7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp7"] = t
		}
	}

	if v, ok := d.GetOk("dscp8"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp8(d, v, "dscp8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp8"] = t
		}
	}

	if v, ok := d.GetOk("dscp9"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp9(d, v, "dscp9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp9"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandObjectSystemNpuNpQueuesProfileWeight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
