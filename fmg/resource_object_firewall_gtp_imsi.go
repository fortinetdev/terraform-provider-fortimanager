// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IMSI.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpImsi() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpImsiCreate,
		Read:   resourceObjectFirewallGtpImsiRead,
		Update: resourceObjectFirewallGtpImsiUpdate,
		Delete: resourceObjectFirewallGtpImsiDelete,

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
			"gtp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apnmember": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"mcc_mnc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"msisdn_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"selection_mode": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallGtpImsiCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpImsi(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpImsi resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallGtpImsi(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallGtpImsi resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpImsiRead(d, m)
}

func resourceObjectFirewallGtpImsiUpdate(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpImsi(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpImsi resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpImsi(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpImsi resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallGtpImsiRead(d, m)
}

func resourceObjectFirewallGtpImsiDelete(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	err = c.DeleteObjectFirewallGtpImsi(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpImsi resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpImsiRead(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	if gtp == "" {
		gtp = importOptionChecking(m.(*FortiClient).Cfg, "gtp")
		if gtp == "" {
			return fmt.Errorf("Parameter gtp is missing")
		}
		if err = d.Set("gtp", gtp); err != nil {
			return fmt.Errorf("Error set params gtp: %v", err)
		}
	}
	paradict["gtp"] = gtp

	o, err := c.ReadObjectFirewallGtpImsi(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpImsi resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpImsi(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpImsi resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpImsiAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiApnmember2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiMccMnc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiMsisdnPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpImsiSelectionMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallGtpImsi(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectFirewallGtpImsiAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectFirewallGtpImsi-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("apnmember", flattenObjectFirewallGtpImsiApnmember2edl(o["apnmember"], d, "apnmember")); err != nil {
		if vv, ok := fortiAPIPatch(o["apnmember"], "ObjectFirewallGtpImsi-Apnmember"); ok {
			if err = d.Set("apnmember", vv); err != nil {
				return fmt.Errorf("Error reading apnmember: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apnmember: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallGtpImsiId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallGtpImsi-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("mcc_mnc", flattenObjectFirewallGtpImsiMccMnc2edl(o["mcc-mnc"], d, "mcc_mnc")); err != nil {
		if vv, ok := fortiAPIPatch(o["mcc-mnc"], "ObjectFirewallGtpImsi-MccMnc"); ok {
			if err = d.Set("mcc_mnc", vv); err != nil {
				return fmt.Errorf("Error reading mcc_mnc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mcc_mnc: %v", err)
		}
	}

	if err = d.Set("msisdn_prefix", flattenObjectFirewallGtpImsiMsisdnPrefix2edl(o["msisdn-prefix"], d, "msisdn_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["msisdn-prefix"], "ObjectFirewallGtpImsi-MsisdnPrefix"); ok {
			if err = d.Set("msisdn_prefix", vv); err != nil {
				return fmt.Errorf("Error reading msisdn_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msisdn_prefix: %v", err)
		}
	}

	if err = d.Set("selection_mode", flattenObjectFirewallGtpImsiSelectionMode2edl(o["selection-mode"], d, "selection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["selection-mode"], "ObjectFirewallGtpImsi-SelectionMode"); ok {
			if err = d.Set("selection_mode", vv); err != nil {
				return fmt.Errorf("Error reading selection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selection_mode: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpImsiFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpImsiAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiApnmember2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiMccMnc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiMsisdnPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpImsiSelectionMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFirewallGtpImsi(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectFirewallGtpImsiAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("apnmember"); ok || d.HasChange("apnmember") {
		t, err := expandObjectFirewallGtpImsiApnmember2edl(d, v, "apnmember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apnmember"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallGtpImsiId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mcc_mnc"); ok || d.HasChange("mcc_mnc") {
		t, err := expandObjectFirewallGtpImsiMccMnc2edl(d, v, "mcc_mnc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcc-mnc"] = t
		}
	}

	if v, ok := d.GetOk("msisdn_prefix"); ok || d.HasChange("msisdn_prefix") {
		t, err := expandObjectFirewallGtpImsiMsisdnPrefix2edl(d, v, "msisdn_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msisdn-prefix"] = t
		}
	}

	if v, ok := d.GetOk("selection_mode"); ok || d.HasChange("selection_mode") {
		t, err := expandObjectFirewallGtpImsiSelectionMode2edl(d, v, "selection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selection-mode"] = t
		}
	}

	return &obj, nil
}
