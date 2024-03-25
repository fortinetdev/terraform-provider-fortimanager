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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuNpQueuesProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpQueuesProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuNpQueuesProfile(obj, paradict)

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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuNpQueuesProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpQueuesProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuNpQueuesProfile(obj, mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	err = c.DeleteObjectSystemNpuNpQueuesProfile(mkey, paradict)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectSystemNpuNpQueuesProfile(mkey, paradict)
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

func flattenObjectSystemNpuNpQueuesProfileCos03rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos13rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos23rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos33rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos43rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos53rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileCos73rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp03rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp13rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp103rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp113rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp123rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp133rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp143rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp153rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp163rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp173rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp183rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp193rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp23rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp203rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp213rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp223rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp233rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp243rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp253rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp263rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp273rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp283rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp293rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp33rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp303rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp313rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp323rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp333rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp343rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp353rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp363rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp373rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp383rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp393rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp43rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp403rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp413rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp423rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp433rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp443rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp453rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp463rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp473rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp483rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp493rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp53rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp503rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp513rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp523rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp533rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp543rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp553rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp563rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp573rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp583rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp593rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp603rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp613rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp623rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp633rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp73rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp83rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileDscp93rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpQueuesProfileWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpQueuesProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cos0", flattenObjectSystemNpuNpQueuesProfileCos03rdl(o["cos0"], d, "cos0")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos0"], "ObjectSystemNpuNpQueuesProfile-Cos0"); ok {
			if err = d.Set("cos0", vv); err != nil {
				return fmt.Errorf("Error reading cos0: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos0: %v", err)
		}
	}

	if err = d.Set("cos1", flattenObjectSystemNpuNpQueuesProfileCos13rdl(o["cos1"], d, "cos1")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos1"], "ObjectSystemNpuNpQueuesProfile-Cos1"); ok {
			if err = d.Set("cos1", vv); err != nil {
				return fmt.Errorf("Error reading cos1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos1: %v", err)
		}
	}

	if err = d.Set("cos2", flattenObjectSystemNpuNpQueuesProfileCos23rdl(o["cos2"], d, "cos2")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos2"], "ObjectSystemNpuNpQueuesProfile-Cos2"); ok {
			if err = d.Set("cos2", vv); err != nil {
				return fmt.Errorf("Error reading cos2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos2: %v", err)
		}
	}

	if err = d.Set("cos3", flattenObjectSystemNpuNpQueuesProfileCos33rdl(o["cos3"], d, "cos3")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos3"], "ObjectSystemNpuNpQueuesProfile-Cos3"); ok {
			if err = d.Set("cos3", vv); err != nil {
				return fmt.Errorf("Error reading cos3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos3: %v", err)
		}
	}

	if err = d.Set("cos4", flattenObjectSystemNpuNpQueuesProfileCos43rdl(o["cos4"], d, "cos4")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos4"], "ObjectSystemNpuNpQueuesProfile-Cos4"); ok {
			if err = d.Set("cos4", vv); err != nil {
				return fmt.Errorf("Error reading cos4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos4: %v", err)
		}
	}

	if err = d.Set("cos5", flattenObjectSystemNpuNpQueuesProfileCos53rdl(o["cos5"], d, "cos5")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos5"], "ObjectSystemNpuNpQueuesProfile-Cos5"); ok {
			if err = d.Set("cos5", vv); err != nil {
				return fmt.Errorf("Error reading cos5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos5: %v", err)
		}
	}

	if err = d.Set("cos6", flattenObjectSystemNpuNpQueuesProfileCos63rdl(o["cos6"], d, "cos6")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos6"], "ObjectSystemNpuNpQueuesProfile-Cos6"); ok {
			if err = d.Set("cos6", vv); err != nil {
				return fmt.Errorf("Error reading cos6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos6: %v", err)
		}
	}

	if err = d.Set("cos7", flattenObjectSystemNpuNpQueuesProfileCos73rdl(o["cos7"], d, "cos7")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos7"], "ObjectSystemNpuNpQueuesProfile-Cos7"); ok {
			if err = d.Set("cos7", vv); err != nil {
				return fmt.Errorf("Error reading cos7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos7: %v", err)
		}
	}

	if err = d.Set("dscp0", flattenObjectSystemNpuNpQueuesProfileDscp03rdl(o["dscp0"], d, "dscp0")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp0"], "ObjectSystemNpuNpQueuesProfile-Dscp0"); ok {
			if err = d.Set("dscp0", vv); err != nil {
				return fmt.Errorf("Error reading dscp0: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp0: %v", err)
		}
	}

	if err = d.Set("dscp1", flattenObjectSystemNpuNpQueuesProfileDscp13rdl(o["dscp1"], d, "dscp1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp1"], "ObjectSystemNpuNpQueuesProfile-Dscp1"); ok {
			if err = d.Set("dscp1", vv); err != nil {
				return fmt.Errorf("Error reading dscp1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp1: %v", err)
		}
	}

	if err = d.Set("dscp10", flattenObjectSystemNpuNpQueuesProfileDscp103rdl(o["dscp10"], d, "dscp10")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp10"], "ObjectSystemNpuNpQueuesProfile-Dscp10"); ok {
			if err = d.Set("dscp10", vv); err != nil {
				return fmt.Errorf("Error reading dscp10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp10: %v", err)
		}
	}

	if err = d.Set("dscp11", flattenObjectSystemNpuNpQueuesProfileDscp113rdl(o["dscp11"], d, "dscp11")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp11"], "ObjectSystemNpuNpQueuesProfile-Dscp11"); ok {
			if err = d.Set("dscp11", vv); err != nil {
				return fmt.Errorf("Error reading dscp11: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp11: %v", err)
		}
	}

	if err = d.Set("dscp12", flattenObjectSystemNpuNpQueuesProfileDscp123rdl(o["dscp12"], d, "dscp12")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp12"], "ObjectSystemNpuNpQueuesProfile-Dscp12"); ok {
			if err = d.Set("dscp12", vv); err != nil {
				return fmt.Errorf("Error reading dscp12: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp12: %v", err)
		}
	}

	if err = d.Set("dscp13", flattenObjectSystemNpuNpQueuesProfileDscp133rdl(o["dscp13"], d, "dscp13")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp13"], "ObjectSystemNpuNpQueuesProfile-Dscp13"); ok {
			if err = d.Set("dscp13", vv); err != nil {
				return fmt.Errorf("Error reading dscp13: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp13: %v", err)
		}
	}

	if err = d.Set("dscp14", flattenObjectSystemNpuNpQueuesProfileDscp143rdl(o["dscp14"], d, "dscp14")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp14"], "ObjectSystemNpuNpQueuesProfile-Dscp14"); ok {
			if err = d.Set("dscp14", vv); err != nil {
				return fmt.Errorf("Error reading dscp14: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp14: %v", err)
		}
	}

	if err = d.Set("dscp15", flattenObjectSystemNpuNpQueuesProfileDscp153rdl(o["dscp15"], d, "dscp15")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp15"], "ObjectSystemNpuNpQueuesProfile-Dscp15"); ok {
			if err = d.Set("dscp15", vv); err != nil {
				return fmt.Errorf("Error reading dscp15: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp15: %v", err)
		}
	}

	if err = d.Set("dscp16", flattenObjectSystemNpuNpQueuesProfileDscp163rdl(o["dscp16"], d, "dscp16")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp16"], "ObjectSystemNpuNpQueuesProfile-Dscp16"); ok {
			if err = d.Set("dscp16", vv); err != nil {
				return fmt.Errorf("Error reading dscp16: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp16: %v", err)
		}
	}

	if err = d.Set("dscp17", flattenObjectSystemNpuNpQueuesProfileDscp173rdl(o["dscp17"], d, "dscp17")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp17"], "ObjectSystemNpuNpQueuesProfile-Dscp17"); ok {
			if err = d.Set("dscp17", vv); err != nil {
				return fmt.Errorf("Error reading dscp17: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp17: %v", err)
		}
	}

	if err = d.Set("dscp18", flattenObjectSystemNpuNpQueuesProfileDscp183rdl(o["dscp18"], d, "dscp18")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp18"], "ObjectSystemNpuNpQueuesProfile-Dscp18"); ok {
			if err = d.Set("dscp18", vv); err != nil {
				return fmt.Errorf("Error reading dscp18: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp18: %v", err)
		}
	}

	if err = d.Set("dscp19", flattenObjectSystemNpuNpQueuesProfileDscp193rdl(o["dscp19"], d, "dscp19")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp19"], "ObjectSystemNpuNpQueuesProfile-Dscp19"); ok {
			if err = d.Set("dscp19", vv); err != nil {
				return fmt.Errorf("Error reading dscp19: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp19: %v", err)
		}
	}

	if err = d.Set("dscp2", flattenObjectSystemNpuNpQueuesProfileDscp23rdl(o["dscp2"], d, "dscp2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp2"], "ObjectSystemNpuNpQueuesProfile-Dscp2"); ok {
			if err = d.Set("dscp2", vv); err != nil {
				return fmt.Errorf("Error reading dscp2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp2: %v", err)
		}
	}

	if err = d.Set("dscp20", flattenObjectSystemNpuNpQueuesProfileDscp203rdl(o["dscp20"], d, "dscp20")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp20"], "ObjectSystemNpuNpQueuesProfile-Dscp20"); ok {
			if err = d.Set("dscp20", vv); err != nil {
				return fmt.Errorf("Error reading dscp20: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp20: %v", err)
		}
	}

	if err = d.Set("dscp21", flattenObjectSystemNpuNpQueuesProfileDscp213rdl(o["dscp21"], d, "dscp21")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp21"], "ObjectSystemNpuNpQueuesProfile-Dscp21"); ok {
			if err = d.Set("dscp21", vv); err != nil {
				return fmt.Errorf("Error reading dscp21: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp21: %v", err)
		}
	}

	if err = d.Set("dscp22", flattenObjectSystemNpuNpQueuesProfileDscp223rdl(o["dscp22"], d, "dscp22")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp22"], "ObjectSystemNpuNpQueuesProfile-Dscp22"); ok {
			if err = d.Set("dscp22", vv); err != nil {
				return fmt.Errorf("Error reading dscp22: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp22: %v", err)
		}
	}

	if err = d.Set("dscp23", flattenObjectSystemNpuNpQueuesProfileDscp233rdl(o["dscp23"], d, "dscp23")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp23"], "ObjectSystemNpuNpQueuesProfile-Dscp23"); ok {
			if err = d.Set("dscp23", vv); err != nil {
				return fmt.Errorf("Error reading dscp23: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp23: %v", err)
		}
	}

	if err = d.Set("dscp24", flattenObjectSystemNpuNpQueuesProfileDscp243rdl(o["dscp24"], d, "dscp24")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp24"], "ObjectSystemNpuNpQueuesProfile-Dscp24"); ok {
			if err = d.Set("dscp24", vv); err != nil {
				return fmt.Errorf("Error reading dscp24: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp24: %v", err)
		}
	}

	if err = d.Set("dscp25", flattenObjectSystemNpuNpQueuesProfileDscp253rdl(o["dscp25"], d, "dscp25")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp25"], "ObjectSystemNpuNpQueuesProfile-Dscp25"); ok {
			if err = d.Set("dscp25", vv); err != nil {
				return fmt.Errorf("Error reading dscp25: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp25: %v", err)
		}
	}

	if err = d.Set("dscp26", flattenObjectSystemNpuNpQueuesProfileDscp263rdl(o["dscp26"], d, "dscp26")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp26"], "ObjectSystemNpuNpQueuesProfile-Dscp26"); ok {
			if err = d.Set("dscp26", vv); err != nil {
				return fmt.Errorf("Error reading dscp26: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp26: %v", err)
		}
	}

	if err = d.Set("dscp27", flattenObjectSystemNpuNpQueuesProfileDscp273rdl(o["dscp27"], d, "dscp27")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp27"], "ObjectSystemNpuNpQueuesProfile-Dscp27"); ok {
			if err = d.Set("dscp27", vv); err != nil {
				return fmt.Errorf("Error reading dscp27: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp27: %v", err)
		}
	}

	if err = d.Set("dscp28", flattenObjectSystemNpuNpQueuesProfileDscp283rdl(o["dscp28"], d, "dscp28")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp28"], "ObjectSystemNpuNpQueuesProfile-Dscp28"); ok {
			if err = d.Set("dscp28", vv); err != nil {
				return fmt.Errorf("Error reading dscp28: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp28: %v", err)
		}
	}

	if err = d.Set("dscp29", flattenObjectSystemNpuNpQueuesProfileDscp293rdl(o["dscp29"], d, "dscp29")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp29"], "ObjectSystemNpuNpQueuesProfile-Dscp29"); ok {
			if err = d.Set("dscp29", vv); err != nil {
				return fmt.Errorf("Error reading dscp29: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp29: %v", err)
		}
	}

	if err = d.Set("dscp3", flattenObjectSystemNpuNpQueuesProfileDscp33rdl(o["dscp3"], d, "dscp3")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp3"], "ObjectSystemNpuNpQueuesProfile-Dscp3"); ok {
			if err = d.Set("dscp3", vv); err != nil {
				return fmt.Errorf("Error reading dscp3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp3: %v", err)
		}
	}

	if err = d.Set("dscp30", flattenObjectSystemNpuNpQueuesProfileDscp303rdl(o["dscp30"], d, "dscp30")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp30"], "ObjectSystemNpuNpQueuesProfile-Dscp30"); ok {
			if err = d.Set("dscp30", vv); err != nil {
				return fmt.Errorf("Error reading dscp30: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp30: %v", err)
		}
	}

	if err = d.Set("dscp31", flattenObjectSystemNpuNpQueuesProfileDscp313rdl(o["dscp31"], d, "dscp31")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp31"], "ObjectSystemNpuNpQueuesProfile-Dscp31"); ok {
			if err = d.Set("dscp31", vv); err != nil {
				return fmt.Errorf("Error reading dscp31: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp31: %v", err)
		}
	}

	if err = d.Set("dscp32", flattenObjectSystemNpuNpQueuesProfileDscp323rdl(o["dscp32"], d, "dscp32")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp32"], "ObjectSystemNpuNpQueuesProfile-Dscp32"); ok {
			if err = d.Set("dscp32", vv); err != nil {
				return fmt.Errorf("Error reading dscp32: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp32: %v", err)
		}
	}

	if err = d.Set("dscp33", flattenObjectSystemNpuNpQueuesProfileDscp333rdl(o["dscp33"], d, "dscp33")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp33"], "ObjectSystemNpuNpQueuesProfile-Dscp33"); ok {
			if err = d.Set("dscp33", vv); err != nil {
				return fmt.Errorf("Error reading dscp33: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp33: %v", err)
		}
	}

	if err = d.Set("dscp34", flattenObjectSystemNpuNpQueuesProfileDscp343rdl(o["dscp34"], d, "dscp34")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp34"], "ObjectSystemNpuNpQueuesProfile-Dscp34"); ok {
			if err = d.Set("dscp34", vv); err != nil {
				return fmt.Errorf("Error reading dscp34: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp34: %v", err)
		}
	}

	if err = d.Set("dscp35", flattenObjectSystemNpuNpQueuesProfileDscp353rdl(o["dscp35"], d, "dscp35")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp35"], "ObjectSystemNpuNpQueuesProfile-Dscp35"); ok {
			if err = d.Set("dscp35", vv); err != nil {
				return fmt.Errorf("Error reading dscp35: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp35: %v", err)
		}
	}

	if err = d.Set("dscp36", flattenObjectSystemNpuNpQueuesProfileDscp363rdl(o["dscp36"], d, "dscp36")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp36"], "ObjectSystemNpuNpQueuesProfile-Dscp36"); ok {
			if err = d.Set("dscp36", vv); err != nil {
				return fmt.Errorf("Error reading dscp36: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp36: %v", err)
		}
	}

	if err = d.Set("dscp37", flattenObjectSystemNpuNpQueuesProfileDscp373rdl(o["dscp37"], d, "dscp37")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp37"], "ObjectSystemNpuNpQueuesProfile-Dscp37"); ok {
			if err = d.Set("dscp37", vv); err != nil {
				return fmt.Errorf("Error reading dscp37: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp37: %v", err)
		}
	}

	if err = d.Set("dscp38", flattenObjectSystemNpuNpQueuesProfileDscp383rdl(o["dscp38"], d, "dscp38")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp38"], "ObjectSystemNpuNpQueuesProfile-Dscp38"); ok {
			if err = d.Set("dscp38", vv); err != nil {
				return fmt.Errorf("Error reading dscp38: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp38: %v", err)
		}
	}

	if err = d.Set("dscp39", flattenObjectSystemNpuNpQueuesProfileDscp393rdl(o["dscp39"], d, "dscp39")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp39"], "ObjectSystemNpuNpQueuesProfile-Dscp39"); ok {
			if err = d.Set("dscp39", vv); err != nil {
				return fmt.Errorf("Error reading dscp39: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp39: %v", err)
		}
	}

	if err = d.Set("dscp4", flattenObjectSystemNpuNpQueuesProfileDscp43rdl(o["dscp4"], d, "dscp4")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp4"], "ObjectSystemNpuNpQueuesProfile-Dscp4"); ok {
			if err = d.Set("dscp4", vv); err != nil {
				return fmt.Errorf("Error reading dscp4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp4: %v", err)
		}
	}

	if err = d.Set("dscp40", flattenObjectSystemNpuNpQueuesProfileDscp403rdl(o["dscp40"], d, "dscp40")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp40"], "ObjectSystemNpuNpQueuesProfile-Dscp40"); ok {
			if err = d.Set("dscp40", vv); err != nil {
				return fmt.Errorf("Error reading dscp40: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp40: %v", err)
		}
	}

	if err = d.Set("dscp41", flattenObjectSystemNpuNpQueuesProfileDscp413rdl(o["dscp41"], d, "dscp41")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp41"], "ObjectSystemNpuNpQueuesProfile-Dscp41"); ok {
			if err = d.Set("dscp41", vv); err != nil {
				return fmt.Errorf("Error reading dscp41: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp41: %v", err)
		}
	}

	if err = d.Set("dscp42", flattenObjectSystemNpuNpQueuesProfileDscp423rdl(o["dscp42"], d, "dscp42")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp42"], "ObjectSystemNpuNpQueuesProfile-Dscp42"); ok {
			if err = d.Set("dscp42", vv); err != nil {
				return fmt.Errorf("Error reading dscp42: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp42: %v", err)
		}
	}

	if err = d.Set("dscp43", flattenObjectSystemNpuNpQueuesProfileDscp433rdl(o["dscp43"], d, "dscp43")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp43"], "ObjectSystemNpuNpQueuesProfile-Dscp43"); ok {
			if err = d.Set("dscp43", vv); err != nil {
				return fmt.Errorf("Error reading dscp43: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp43: %v", err)
		}
	}

	if err = d.Set("dscp44", flattenObjectSystemNpuNpQueuesProfileDscp443rdl(o["dscp44"], d, "dscp44")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp44"], "ObjectSystemNpuNpQueuesProfile-Dscp44"); ok {
			if err = d.Set("dscp44", vv); err != nil {
				return fmt.Errorf("Error reading dscp44: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp44: %v", err)
		}
	}

	if err = d.Set("dscp45", flattenObjectSystemNpuNpQueuesProfileDscp453rdl(o["dscp45"], d, "dscp45")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp45"], "ObjectSystemNpuNpQueuesProfile-Dscp45"); ok {
			if err = d.Set("dscp45", vv); err != nil {
				return fmt.Errorf("Error reading dscp45: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp45: %v", err)
		}
	}

	if err = d.Set("dscp46", flattenObjectSystemNpuNpQueuesProfileDscp463rdl(o["dscp46"], d, "dscp46")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp46"], "ObjectSystemNpuNpQueuesProfile-Dscp46"); ok {
			if err = d.Set("dscp46", vv); err != nil {
				return fmt.Errorf("Error reading dscp46: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp46: %v", err)
		}
	}

	if err = d.Set("dscp47", flattenObjectSystemNpuNpQueuesProfileDscp473rdl(o["dscp47"], d, "dscp47")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp47"], "ObjectSystemNpuNpQueuesProfile-Dscp47"); ok {
			if err = d.Set("dscp47", vv); err != nil {
				return fmt.Errorf("Error reading dscp47: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp47: %v", err)
		}
	}

	if err = d.Set("dscp48", flattenObjectSystemNpuNpQueuesProfileDscp483rdl(o["dscp48"], d, "dscp48")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp48"], "ObjectSystemNpuNpQueuesProfile-Dscp48"); ok {
			if err = d.Set("dscp48", vv); err != nil {
				return fmt.Errorf("Error reading dscp48: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp48: %v", err)
		}
	}

	if err = d.Set("dscp49", flattenObjectSystemNpuNpQueuesProfileDscp493rdl(o["dscp49"], d, "dscp49")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp49"], "ObjectSystemNpuNpQueuesProfile-Dscp49"); ok {
			if err = d.Set("dscp49", vv); err != nil {
				return fmt.Errorf("Error reading dscp49: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp49: %v", err)
		}
	}

	if err = d.Set("dscp5", flattenObjectSystemNpuNpQueuesProfileDscp53rdl(o["dscp5"], d, "dscp5")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp5"], "ObjectSystemNpuNpQueuesProfile-Dscp5"); ok {
			if err = d.Set("dscp5", vv); err != nil {
				return fmt.Errorf("Error reading dscp5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp5: %v", err)
		}
	}

	if err = d.Set("dscp50", flattenObjectSystemNpuNpQueuesProfileDscp503rdl(o["dscp50"], d, "dscp50")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp50"], "ObjectSystemNpuNpQueuesProfile-Dscp50"); ok {
			if err = d.Set("dscp50", vv); err != nil {
				return fmt.Errorf("Error reading dscp50: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp50: %v", err)
		}
	}

	if err = d.Set("dscp51", flattenObjectSystemNpuNpQueuesProfileDscp513rdl(o["dscp51"], d, "dscp51")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp51"], "ObjectSystemNpuNpQueuesProfile-Dscp51"); ok {
			if err = d.Set("dscp51", vv); err != nil {
				return fmt.Errorf("Error reading dscp51: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp51: %v", err)
		}
	}

	if err = d.Set("dscp52", flattenObjectSystemNpuNpQueuesProfileDscp523rdl(o["dscp52"], d, "dscp52")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp52"], "ObjectSystemNpuNpQueuesProfile-Dscp52"); ok {
			if err = d.Set("dscp52", vv); err != nil {
				return fmt.Errorf("Error reading dscp52: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp52: %v", err)
		}
	}

	if err = d.Set("dscp53", flattenObjectSystemNpuNpQueuesProfileDscp533rdl(o["dscp53"], d, "dscp53")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp53"], "ObjectSystemNpuNpQueuesProfile-Dscp53"); ok {
			if err = d.Set("dscp53", vv); err != nil {
				return fmt.Errorf("Error reading dscp53: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp53: %v", err)
		}
	}

	if err = d.Set("dscp54", flattenObjectSystemNpuNpQueuesProfileDscp543rdl(o["dscp54"], d, "dscp54")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp54"], "ObjectSystemNpuNpQueuesProfile-Dscp54"); ok {
			if err = d.Set("dscp54", vv); err != nil {
				return fmt.Errorf("Error reading dscp54: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp54: %v", err)
		}
	}

	if err = d.Set("dscp55", flattenObjectSystemNpuNpQueuesProfileDscp553rdl(o["dscp55"], d, "dscp55")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp55"], "ObjectSystemNpuNpQueuesProfile-Dscp55"); ok {
			if err = d.Set("dscp55", vv); err != nil {
				return fmt.Errorf("Error reading dscp55: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp55: %v", err)
		}
	}

	if err = d.Set("dscp56", flattenObjectSystemNpuNpQueuesProfileDscp563rdl(o["dscp56"], d, "dscp56")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp56"], "ObjectSystemNpuNpQueuesProfile-Dscp56"); ok {
			if err = d.Set("dscp56", vv); err != nil {
				return fmt.Errorf("Error reading dscp56: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp56: %v", err)
		}
	}

	if err = d.Set("dscp57", flattenObjectSystemNpuNpQueuesProfileDscp573rdl(o["dscp57"], d, "dscp57")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp57"], "ObjectSystemNpuNpQueuesProfile-Dscp57"); ok {
			if err = d.Set("dscp57", vv); err != nil {
				return fmt.Errorf("Error reading dscp57: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp57: %v", err)
		}
	}

	if err = d.Set("dscp58", flattenObjectSystemNpuNpQueuesProfileDscp583rdl(o["dscp58"], d, "dscp58")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp58"], "ObjectSystemNpuNpQueuesProfile-Dscp58"); ok {
			if err = d.Set("dscp58", vv); err != nil {
				return fmt.Errorf("Error reading dscp58: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp58: %v", err)
		}
	}

	if err = d.Set("dscp59", flattenObjectSystemNpuNpQueuesProfileDscp593rdl(o["dscp59"], d, "dscp59")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp59"], "ObjectSystemNpuNpQueuesProfile-Dscp59"); ok {
			if err = d.Set("dscp59", vv); err != nil {
				return fmt.Errorf("Error reading dscp59: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp59: %v", err)
		}
	}

	if err = d.Set("dscp6", flattenObjectSystemNpuNpQueuesProfileDscp63rdl(o["dscp6"], d, "dscp6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp6"], "ObjectSystemNpuNpQueuesProfile-Dscp6"); ok {
			if err = d.Set("dscp6", vv); err != nil {
				return fmt.Errorf("Error reading dscp6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp6: %v", err)
		}
	}

	if err = d.Set("dscp60", flattenObjectSystemNpuNpQueuesProfileDscp603rdl(o["dscp60"], d, "dscp60")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp60"], "ObjectSystemNpuNpQueuesProfile-Dscp60"); ok {
			if err = d.Set("dscp60", vv); err != nil {
				return fmt.Errorf("Error reading dscp60: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp60: %v", err)
		}
	}

	if err = d.Set("dscp61", flattenObjectSystemNpuNpQueuesProfileDscp613rdl(o["dscp61"], d, "dscp61")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp61"], "ObjectSystemNpuNpQueuesProfile-Dscp61"); ok {
			if err = d.Set("dscp61", vv); err != nil {
				return fmt.Errorf("Error reading dscp61: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp61: %v", err)
		}
	}

	if err = d.Set("dscp62", flattenObjectSystemNpuNpQueuesProfileDscp623rdl(o["dscp62"], d, "dscp62")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp62"], "ObjectSystemNpuNpQueuesProfile-Dscp62"); ok {
			if err = d.Set("dscp62", vv); err != nil {
				return fmt.Errorf("Error reading dscp62: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp62: %v", err)
		}
	}

	if err = d.Set("dscp63", flattenObjectSystemNpuNpQueuesProfileDscp633rdl(o["dscp63"], d, "dscp63")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp63"], "ObjectSystemNpuNpQueuesProfile-Dscp63"); ok {
			if err = d.Set("dscp63", vv); err != nil {
				return fmt.Errorf("Error reading dscp63: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp63: %v", err)
		}
	}

	if err = d.Set("dscp7", flattenObjectSystemNpuNpQueuesProfileDscp73rdl(o["dscp7"], d, "dscp7")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp7"], "ObjectSystemNpuNpQueuesProfile-Dscp7"); ok {
			if err = d.Set("dscp7", vv); err != nil {
				return fmt.Errorf("Error reading dscp7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp7: %v", err)
		}
	}

	if err = d.Set("dscp8", flattenObjectSystemNpuNpQueuesProfileDscp83rdl(o["dscp8"], d, "dscp8")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp8"], "ObjectSystemNpuNpQueuesProfile-Dscp8"); ok {
			if err = d.Set("dscp8", vv); err != nil {
				return fmt.Errorf("Error reading dscp8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp8: %v", err)
		}
	}

	if err = d.Set("dscp9", flattenObjectSystemNpuNpQueuesProfileDscp93rdl(o["dscp9"], d, "dscp9")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp9"], "ObjectSystemNpuNpQueuesProfile-Dscp9"); ok {
			if err = d.Set("dscp9", vv); err != nil {
				return fmt.Errorf("Error reading dscp9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp9: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemNpuNpQueuesProfileId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemNpuNpQueuesProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSystemNpuNpQueuesProfileType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemNpuNpQueuesProfile-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectSystemNpuNpQueuesProfileWeight3rdl(o["weight"], d, "weight")); err != nil {
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

func expandObjectSystemNpuNpQueuesProfileCos03rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos13rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos23rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos33rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos43rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos53rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileCos73rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp03rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp13rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp103rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp113rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp123rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp133rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp143rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp153rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp163rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp173rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp183rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp193rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp23rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp203rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp213rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp223rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp233rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp243rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp253rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp263rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp273rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp283rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp293rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp33rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp303rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp313rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp323rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp333rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp343rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp353rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp363rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp373rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp383rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp393rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp43rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp403rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp413rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp423rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp433rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp443rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp453rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp463rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp473rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp483rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp493rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp53rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp503rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp513rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp523rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp533rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp543rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp553rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp563rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp573rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp583rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp593rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp603rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp613rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp623rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp633rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp73rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp83rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileDscp93rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpQueuesProfileWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpQueuesProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cos0"); ok || d.HasChange("cos0") {
		t, err := expandObjectSystemNpuNpQueuesProfileCos03rdl(d, v, "cos0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos0"] = t
		}
	}

	if v, ok := d.GetOk("cos1"); ok || d.HasChange("cos1") {
		t, err := expandObjectSystemNpuNpQueuesProfileCos13rdl(d, v, "cos1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos1"] = t
		}
	}

	if v, ok := d.GetOk("cos2"); ok || d.HasChange("cos2") {
		t, err := expandObjectSystemNpuNpQueuesProfileCos23rdl(d, v, "cos2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos2"] = t
		}
	}

	if v, ok := d.GetOk("cos3"); ok || d.HasChange("cos3") {
		t, err := expandObjectSystemNpuNpQueuesProfileCos33rdl(d, v, "cos3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos3"] = t
		}
	}

	if v, ok := d.GetOk("cos4"); ok || d.HasChange("cos4") {
		t, err := expandObjectSystemNpuNpQueuesProfileCos43rdl(d, v, "cos4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos4"] = t
		}
	}

	if v, ok := d.GetOk("cos5"); ok || d.HasChange("cos5") {
		t, err := expandObjectSystemNpuNpQueuesProfileCos53rdl(d, v, "cos5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos5"] = t
		}
	}

	if v, ok := d.GetOk("cos6"); ok || d.HasChange("cos6") {
		t, err := expandObjectSystemNpuNpQueuesProfileCos63rdl(d, v, "cos6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos6"] = t
		}
	}

	if v, ok := d.GetOk("cos7"); ok || d.HasChange("cos7") {
		t, err := expandObjectSystemNpuNpQueuesProfileCos73rdl(d, v, "cos7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos7"] = t
		}
	}

	if v, ok := d.GetOk("dscp0"); ok || d.HasChange("dscp0") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp03rdl(d, v, "dscp0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp0"] = t
		}
	}

	if v, ok := d.GetOk("dscp1"); ok || d.HasChange("dscp1") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp13rdl(d, v, "dscp1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp1"] = t
		}
	}

	if v, ok := d.GetOk("dscp10"); ok || d.HasChange("dscp10") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp103rdl(d, v, "dscp10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp10"] = t
		}
	}

	if v, ok := d.GetOk("dscp11"); ok || d.HasChange("dscp11") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp113rdl(d, v, "dscp11")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp11"] = t
		}
	}

	if v, ok := d.GetOk("dscp12"); ok || d.HasChange("dscp12") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp123rdl(d, v, "dscp12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp12"] = t
		}
	}

	if v, ok := d.GetOk("dscp13"); ok || d.HasChange("dscp13") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp133rdl(d, v, "dscp13")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp13"] = t
		}
	}

	if v, ok := d.GetOk("dscp14"); ok || d.HasChange("dscp14") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp143rdl(d, v, "dscp14")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp14"] = t
		}
	}

	if v, ok := d.GetOk("dscp15"); ok || d.HasChange("dscp15") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp153rdl(d, v, "dscp15")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp15"] = t
		}
	}

	if v, ok := d.GetOk("dscp16"); ok || d.HasChange("dscp16") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp163rdl(d, v, "dscp16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp16"] = t
		}
	}

	if v, ok := d.GetOk("dscp17"); ok || d.HasChange("dscp17") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp173rdl(d, v, "dscp17")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp17"] = t
		}
	}

	if v, ok := d.GetOk("dscp18"); ok || d.HasChange("dscp18") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp183rdl(d, v, "dscp18")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp18"] = t
		}
	}

	if v, ok := d.GetOk("dscp19"); ok || d.HasChange("dscp19") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp193rdl(d, v, "dscp19")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp19"] = t
		}
	}

	if v, ok := d.GetOk("dscp2"); ok || d.HasChange("dscp2") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp23rdl(d, v, "dscp2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp2"] = t
		}
	}

	if v, ok := d.GetOk("dscp20"); ok || d.HasChange("dscp20") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp203rdl(d, v, "dscp20")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp20"] = t
		}
	}

	if v, ok := d.GetOk("dscp21"); ok || d.HasChange("dscp21") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp213rdl(d, v, "dscp21")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp21"] = t
		}
	}

	if v, ok := d.GetOk("dscp22"); ok || d.HasChange("dscp22") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp223rdl(d, v, "dscp22")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp22"] = t
		}
	}

	if v, ok := d.GetOk("dscp23"); ok || d.HasChange("dscp23") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp233rdl(d, v, "dscp23")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp23"] = t
		}
	}

	if v, ok := d.GetOk("dscp24"); ok || d.HasChange("dscp24") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp243rdl(d, v, "dscp24")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp24"] = t
		}
	}

	if v, ok := d.GetOk("dscp25"); ok || d.HasChange("dscp25") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp253rdl(d, v, "dscp25")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp25"] = t
		}
	}

	if v, ok := d.GetOk("dscp26"); ok || d.HasChange("dscp26") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp263rdl(d, v, "dscp26")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp26"] = t
		}
	}

	if v, ok := d.GetOk("dscp27"); ok || d.HasChange("dscp27") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp273rdl(d, v, "dscp27")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp27"] = t
		}
	}

	if v, ok := d.GetOk("dscp28"); ok || d.HasChange("dscp28") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp283rdl(d, v, "dscp28")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp28"] = t
		}
	}

	if v, ok := d.GetOk("dscp29"); ok || d.HasChange("dscp29") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp293rdl(d, v, "dscp29")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp29"] = t
		}
	}

	if v, ok := d.GetOk("dscp3"); ok || d.HasChange("dscp3") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp33rdl(d, v, "dscp3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp3"] = t
		}
	}

	if v, ok := d.GetOk("dscp30"); ok || d.HasChange("dscp30") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp303rdl(d, v, "dscp30")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp30"] = t
		}
	}

	if v, ok := d.GetOk("dscp31"); ok || d.HasChange("dscp31") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp313rdl(d, v, "dscp31")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp31"] = t
		}
	}

	if v, ok := d.GetOk("dscp32"); ok || d.HasChange("dscp32") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp323rdl(d, v, "dscp32")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp32"] = t
		}
	}

	if v, ok := d.GetOk("dscp33"); ok || d.HasChange("dscp33") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp333rdl(d, v, "dscp33")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp33"] = t
		}
	}

	if v, ok := d.GetOk("dscp34"); ok || d.HasChange("dscp34") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp343rdl(d, v, "dscp34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp34"] = t
		}
	}

	if v, ok := d.GetOk("dscp35"); ok || d.HasChange("dscp35") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp353rdl(d, v, "dscp35")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp35"] = t
		}
	}

	if v, ok := d.GetOk("dscp36"); ok || d.HasChange("dscp36") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp363rdl(d, v, "dscp36")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp36"] = t
		}
	}

	if v, ok := d.GetOk("dscp37"); ok || d.HasChange("dscp37") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp373rdl(d, v, "dscp37")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp37"] = t
		}
	}

	if v, ok := d.GetOk("dscp38"); ok || d.HasChange("dscp38") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp383rdl(d, v, "dscp38")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp38"] = t
		}
	}

	if v, ok := d.GetOk("dscp39"); ok || d.HasChange("dscp39") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp393rdl(d, v, "dscp39")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp39"] = t
		}
	}

	if v, ok := d.GetOk("dscp4"); ok || d.HasChange("dscp4") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp43rdl(d, v, "dscp4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp4"] = t
		}
	}

	if v, ok := d.GetOk("dscp40"); ok || d.HasChange("dscp40") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp403rdl(d, v, "dscp40")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp40"] = t
		}
	}

	if v, ok := d.GetOk("dscp41"); ok || d.HasChange("dscp41") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp413rdl(d, v, "dscp41")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp41"] = t
		}
	}

	if v, ok := d.GetOk("dscp42"); ok || d.HasChange("dscp42") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp423rdl(d, v, "dscp42")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp42"] = t
		}
	}

	if v, ok := d.GetOk("dscp43"); ok || d.HasChange("dscp43") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp433rdl(d, v, "dscp43")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp43"] = t
		}
	}

	if v, ok := d.GetOk("dscp44"); ok || d.HasChange("dscp44") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp443rdl(d, v, "dscp44")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp44"] = t
		}
	}

	if v, ok := d.GetOk("dscp45"); ok || d.HasChange("dscp45") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp453rdl(d, v, "dscp45")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp45"] = t
		}
	}

	if v, ok := d.GetOk("dscp46"); ok || d.HasChange("dscp46") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp463rdl(d, v, "dscp46")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp46"] = t
		}
	}

	if v, ok := d.GetOk("dscp47"); ok || d.HasChange("dscp47") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp473rdl(d, v, "dscp47")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp47"] = t
		}
	}

	if v, ok := d.GetOk("dscp48"); ok || d.HasChange("dscp48") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp483rdl(d, v, "dscp48")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp48"] = t
		}
	}

	if v, ok := d.GetOk("dscp49"); ok || d.HasChange("dscp49") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp493rdl(d, v, "dscp49")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp49"] = t
		}
	}

	if v, ok := d.GetOk("dscp5"); ok || d.HasChange("dscp5") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp53rdl(d, v, "dscp5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp5"] = t
		}
	}

	if v, ok := d.GetOk("dscp50"); ok || d.HasChange("dscp50") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp503rdl(d, v, "dscp50")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp50"] = t
		}
	}

	if v, ok := d.GetOk("dscp51"); ok || d.HasChange("dscp51") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp513rdl(d, v, "dscp51")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp51"] = t
		}
	}

	if v, ok := d.GetOk("dscp52"); ok || d.HasChange("dscp52") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp523rdl(d, v, "dscp52")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp52"] = t
		}
	}

	if v, ok := d.GetOk("dscp53"); ok || d.HasChange("dscp53") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp533rdl(d, v, "dscp53")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp53"] = t
		}
	}

	if v, ok := d.GetOk("dscp54"); ok || d.HasChange("dscp54") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp543rdl(d, v, "dscp54")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp54"] = t
		}
	}

	if v, ok := d.GetOk("dscp55"); ok || d.HasChange("dscp55") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp553rdl(d, v, "dscp55")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp55"] = t
		}
	}

	if v, ok := d.GetOk("dscp56"); ok || d.HasChange("dscp56") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp563rdl(d, v, "dscp56")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp56"] = t
		}
	}

	if v, ok := d.GetOk("dscp57"); ok || d.HasChange("dscp57") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp573rdl(d, v, "dscp57")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp57"] = t
		}
	}

	if v, ok := d.GetOk("dscp58"); ok || d.HasChange("dscp58") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp583rdl(d, v, "dscp58")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp58"] = t
		}
	}

	if v, ok := d.GetOk("dscp59"); ok || d.HasChange("dscp59") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp593rdl(d, v, "dscp59")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp59"] = t
		}
	}

	if v, ok := d.GetOk("dscp6"); ok || d.HasChange("dscp6") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp63rdl(d, v, "dscp6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp6"] = t
		}
	}

	if v, ok := d.GetOk("dscp60"); ok || d.HasChange("dscp60") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp603rdl(d, v, "dscp60")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp60"] = t
		}
	}

	if v, ok := d.GetOk("dscp61"); ok || d.HasChange("dscp61") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp613rdl(d, v, "dscp61")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp61"] = t
		}
	}

	if v, ok := d.GetOk("dscp62"); ok || d.HasChange("dscp62") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp623rdl(d, v, "dscp62")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp62"] = t
		}
	}

	if v, ok := d.GetOk("dscp63"); ok || d.HasChange("dscp63") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp633rdl(d, v, "dscp63")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp63"] = t
		}
	}

	if v, ok := d.GetOk("dscp7"); ok || d.HasChange("dscp7") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp73rdl(d, v, "dscp7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp7"] = t
		}
	}

	if v, ok := d.GetOk("dscp8"); ok || d.HasChange("dscp8") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp83rdl(d, v, "dscp8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp8"] = t
		}
	}

	if v, ok := d.GetOk("dscp9"); ok || d.HasChange("dscp9") {
		t, err := expandObjectSystemNpuNpQueuesProfileDscp93rdl(d, v, "dscp9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp9"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSystemNpuNpQueuesProfileId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSystemNpuNpQueuesProfileType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectSystemNpuNpQueuesProfileWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
