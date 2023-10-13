// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IE validation.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpIeValidation() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpIeValidationUpdate,
		Read:   resourceObjectFirewallGtpIeValidationRead,
		Update: resourceObjectFirewallGtpIeValidationUpdate,
		Delete: resourceObjectFirewallGtpIeValidationDelete,

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
			"apn_restriction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"charging_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"charging_gateway_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_user_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gsn_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"imei": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"imsi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mm_context": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ms_tzone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ms_validated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"msisdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsapi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pdp_context": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"qos_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rai": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rat_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reordering_required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"selection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uli": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallGtpIeValidationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallGtpIeValidation(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpIeValidation resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpIeValidation(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpIeValidation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallGtpIeValidation")

	return resourceObjectFirewallGtpIeValidationRead(d, m)
}

func resourceObjectFirewallGtpIeValidationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallGtpIeValidation(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpIeValidation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpIeValidationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallGtpIeValidation(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpIeValidation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpIeValidation(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpIeValidation resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpIeValidationApnRestriction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationChargingId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationChargingGatewayAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationEndUserAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationGsnAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationImei2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationImsi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationMmContext2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationMsTzone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationMsValidated2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationMsisdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationNsapi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationPdpContext2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationQosProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationRai2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationRatType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationReorderingRequired2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationSelectionMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpIeValidationUli2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallGtpIeValidation(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("apn_restriction", flattenObjectFirewallGtpIeValidationApnRestriction2edl(o["apn-restriction"], d, "apn_restriction")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn-restriction"], "ObjectFirewallGtpIeValidation-ApnRestriction"); ok {
			if err = d.Set("apn_restriction", vv); err != nil {
				return fmt.Errorf("Error reading apn_restriction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn_restriction: %v", err)
		}
	}

	if err = d.Set("charging_id", flattenObjectFirewallGtpIeValidationChargingId2edl(o["charging-ID"], d, "charging_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["charging-ID"], "ObjectFirewallGtpIeValidation-ChargingId"); ok {
			if err = d.Set("charging_id", vv); err != nil {
				return fmt.Errorf("Error reading charging_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading charging_id: %v", err)
		}
	}

	if err = d.Set("charging_gateway_addr", flattenObjectFirewallGtpIeValidationChargingGatewayAddr2edl(o["charging-gateway-addr"], d, "charging_gateway_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["charging-gateway-addr"], "ObjectFirewallGtpIeValidation-ChargingGatewayAddr"); ok {
			if err = d.Set("charging_gateway_addr", vv); err != nil {
				return fmt.Errorf("Error reading charging_gateway_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading charging_gateway_addr: %v", err)
		}
	}

	if err = d.Set("end_user_addr", flattenObjectFirewallGtpIeValidationEndUserAddr2edl(o["end-user-addr"], d, "end_user_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-user-addr"], "ObjectFirewallGtpIeValidation-EndUserAddr"); ok {
			if err = d.Set("end_user_addr", vv); err != nil {
				return fmt.Errorf("Error reading end_user_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_user_addr: %v", err)
		}
	}

	if err = d.Set("gsn_addr", flattenObjectFirewallGtpIeValidationGsnAddr2edl(o["gsn-addr"], d, "gsn_addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["gsn-addr"], "ObjectFirewallGtpIeValidation-GsnAddr"); ok {
			if err = d.Set("gsn_addr", vv); err != nil {
				return fmt.Errorf("Error reading gsn_addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gsn_addr: %v", err)
		}
	}

	if err = d.Set("imei", flattenObjectFirewallGtpIeValidationImei2edl(o["imei"], d, "imei")); err != nil {
		if vv, ok := fortiAPIPatch(o["imei"], "ObjectFirewallGtpIeValidation-Imei"); ok {
			if err = d.Set("imei", vv); err != nil {
				return fmt.Errorf("Error reading imei: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imei: %v", err)
		}
	}

	if err = d.Set("imsi", flattenObjectFirewallGtpIeValidationImsi2edl(o["imsi"], d, "imsi")); err != nil {
		if vv, ok := fortiAPIPatch(o["imsi"], "ObjectFirewallGtpIeValidation-Imsi"); ok {
			if err = d.Set("imsi", vv); err != nil {
				return fmt.Errorf("Error reading imsi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading imsi: %v", err)
		}
	}

	if err = d.Set("mm_context", flattenObjectFirewallGtpIeValidationMmContext2edl(o["mm-context"], d, "mm_context")); err != nil {
		if vv, ok := fortiAPIPatch(o["mm-context"], "ObjectFirewallGtpIeValidation-MmContext"); ok {
			if err = d.Set("mm_context", vv); err != nil {
				return fmt.Errorf("Error reading mm_context: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mm_context: %v", err)
		}
	}

	if err = d.Set("ms_tzone", flattenObjectFirewallGtpIeValidationMsTzone2edl(o["ms-tzone"], d, "ms_tzone")); err != nil {
		if vv, ok := fortiAPIPatch(o["ms-tzone"], "ObjectFirewallGtpIeValidation-MsTzone"); ok {
			if err = d.Set("ms_tzone", vv); err != nil {
				return fmt.Errorf("Error reading ms_tzone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ms_tzone: %v", err)
		}
	}

	if err = d.Set("ms_validated", flattenObjectFirewallGtpIeValidationMsValidated2edl(o["ms-validated"], d, "ms_validated")); err != nil {
		if vv, ok := fortiAPIPatch(o["ms-validated"], "ObjectFirewallGtpIeValidation-MsValidated"); ok {
			if err = d.Set("ms_validated", vv); err != nil {
				return fmt.Errorf("Error reading ms_validated: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ms_validated: %v", err)
		}
	}

	if err = d.Set("msisdn", flattenObjectFirewallGtpIeValidationMsisdn2edl(o["msisdn"], d, "msisdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["msisdn"], "ObjectFirewallGtpIeValidation-Msisdn"); ok {
			if err = d.Set("msisdn", vv); err != nil {
				return fmt.Errorf("Error reading msisdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msisdn: %v", err)
		}
	}

	if err = d.Set("nsapi", flattenObjectFirewallGtpIeValidationNsapi2edl(o["nsapi"], d, "nsapi")); err != nil {
		if vv, ok := fortiAPIPatch(o["nsapi"], "ObjectFirewallGtpIeValidation-Nsapi"); ok {
			if err = d.Set("nsapi", vv); err != nil {
				return fmt.Errorf("Error reading nsapi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nsapi: %v", err)
		}
	}

	if err = d.Set("pdp_context", flattenObjectFirewallGtpIeValidationPdpContext2edl(o["pdp-context"], d, "pdp_context")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdp-context"], "ObjectFirewallGtpIeValidation-PdpContext"); ok {
			if err = d.Set("pdp_context", vv); err != nil {
				return fmt.Errorf("Error reading pdp_context: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdp_context: %v", err)
		}
	}

	if err = d.Set("qos_profile", flattenObjectFirewallGtpIeValidationQosProfile2edl(o["qos-profile"], d, "qos_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qos-profile"], "ObjectFirewallGtpIeValidation-QosProfile"); ok {
			if err = d.Set("qos_profile", vv); err != nil {
				return fmt.Errorf("Error reading qos_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qos_profile: %v", err)
		}
	}

	if err = d.Set("rai", flattenObjectFirewallGtpIeValidationRai2edl(o["rai"], d, "rai")); err != nil {
		if vv, ok := fortiAPIPatch(o["rai"], "ObjectFirewallGtpIeValidation-Rai"); ok {
			if err = d.Set("rai", vv); err != nil {
				return fmt.Errorf("Error reading rai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rai: %v", err)
		}
	}

	if err = d.Set("rat_type", flattenObjectFirewallGtpIeValidationRatType2edl(o["rat-type"], d, "rat_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["rat-type"], "ObjectFirewallGtpIeValidation-RatType"); ok {
			if err = d.Set("rat_type", vv); err != nil {
				return fmt.Errorf("Error reading rat_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rat_type: %v", err)
		}
	}

	if err = d.Set("reordering_required", flattenObjectFirewallGtpIeValidationReorderingRequired2edl(o["reordering-required"], d, "reordering_required")); err != nil {
		if vv, ok := fortiAPIPatch(o["reordering-required"], "ObjectFirewallGtpIeValidation-ReorderingRequired"); ok {
			if err = d.Set("reordering_required", vv); err != nil {
				return fmt.Errorf("Error reading reordering_required: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reordering_required: %v", err)
		}
	}

	if err = d.Set("selection_mode", flattenObjectFirewallGtpIeValidationSelectionMode2edl(o["selection-mode"], d, "selection_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["selection-mode"], "ObjectFirewallGtpIeValidation-SelectionMode"); ok {
			if err = d.Set("selection_mode", vv); err != nil {
				return fmt.Errorf("Error reading selection_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading selection_mode: %v", err)
		}
	}

	if err = d.Set("uli", flattenObjectFirewallGtpIeValidationUli2edl(o["uli"], d, "uli")); err != nil {
		if vv, ok := fortiAPIPatch(o["uli"], "ObjectFirewallGtpIeValidation-Uli"); ok {
			if err = d.Set("uli", vv); err != nil {
				return fmt.Errorf("Error reading uli: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uli: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpIeValidationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpIeValidationApnRestriction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationChargingId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationChargingGatewayAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationEndUserAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationGsnAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationImei2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationImsi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationMmContext2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationMsTzone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationMsValidated2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationMsisdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationNsapi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationPdpContext2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationQosProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationRai2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationRatType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationReorderingRequired2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationSelectionMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpIeValidationUli2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallGtpIeValidation(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apn_restriction"); ok || d.HasChange("apn_restriction") {
		t, err := expandObjectFirewallGtpIeValidationApnRestriction2edl(d, v, "apn_restriction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn-restriction"] = t
		}
	}

	if v, ok := d.GetOk("charging_id"); ok || d.HasChange("charging_id") {
		t, err := expandObjectFirewallGtpIeValidationChargingId2edl(d, v, "charging_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["charging-ID"] = t
		}
	}

	if v, ok := d.GetOk("charging_gateway_addr"); ok || d.HasChange("charging_gateway_addr") {
		t, err := expandObjectFirewallGtpIeValidationChargingGatewayAddr2edl(d, v, "charging_gateway_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["charging-gateway-addr"] = t
		}
	}

	if v, ok := d.GetOk("end_user_addr"); ok || d.HasChange("end_user_addr") {
		t, err := expandObjectFirewallGtpIeValidationEndUserAddr2edl(d, v, "end_user_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-user-addr"] = t
		}
	}

	if v, ok := d.GetOk("gsn_addr"); ok || d.HasChange("gsn_addr") {
		t, err := expandObjectFirewallGtpIeValidationGsnAddr2edl(d, v, "gsn_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gsn-addr"] = t
		}
	}

	if v, ok := d.GetOk("imei"); ok || d.HasChange("imei") {
		t, err := expandObjectFirewallGtpIeValidationImei2edl(d, v, "imei")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imei"] = t
		}
	}

	if v, ok := d.GetOk("imsi"); ok || d.HasChange("imsi") {
		t, err := expandObjectFirewallGtpIeValidationImsi2edl(d, v, "imsi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imsi"] = t
		}
	}

	if v, ok := d.GetOk("mm_context"); ok || d.HasChange("mm_context") {
		t, err := expandObjectFirewallGtpIeValidationMmContext2edl(d, v, "mm_context")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mm-context"] = t
		}
	}

	if v, ok := d.GetOk("ms_tzone"); ok || d.HasChange("ms_tzone") {
		t, err := expandObjectFirewallGtpIeValidationMsTzone2edl(d, v, "ms_tzone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ms-tzone"] = t
		}
	}

	if v, ok := d.GetOk("ms_validated"); ok || d.HasChange("ms_validated") {
		t, err := expandObjectFirewallGtpIeValidationMsValidated2edl(d, v, "ms_validated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ms-validated"] = t
		}
	}

	if v, ok := d.GetOk("msisdn"); ok || d.HasChange("msisdn") {
		t, err := expandObjectFirewallGtpIeValidationMsisdn2edl(d, v, "msisdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msisdn"] = t
		}
	}

	if v, ok := d.GetOk("nsapi"); ok || d.HasChange("nsapi") {
		t, err := expandObjectFirewallGtpIeValidationNsapi2edl(d, v, "nsapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nsapi"] = t
		}
	}

	if v, ok := d.GetOk("pdp_context"); ok || d.HasChange("pdp_context") {
		t, err := expandObjectFirewallGtpIeValidationPdpContext2edl(d, v, "pdp_context")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdp-context"] = t
		}
	}

	if v, ok := d.GetOk("qos_profile"); ok || d.HasChange("qos_profile") {
		t, err := expandObjectFirewallGtpIeValidationQosProfile2edl(d, v, "qos_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-profile"] = t
		}
	}

	if v, ok := d.GetOk("rai"); ok || d.HasChange("rai") {
		t, err := expandObjectFirewallGtpIeValidationRai2edl(d, v, "rai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rai"] = t
		}
	}

	if v, ok := d.GetOk("rat_type"); ok || d.HasChange("rat_type") {
		t, err := expandObjectFirewallGtpIeValidationRatType2edl(d, v, "rat_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rat-type"] = t
		}
	}

	if v, ok := d.GetOk("reordering_required"); ok || d.HasChange("reordering_required") {
		t, err := expandObjectFirewallGtpIeValidationReorderingRequired2edl(d, v, "reordering_required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reordering-required"] = t
		}
	}

	if v, ok := d.GetOk("selection_mode"); ok || d.HasChange("selection_mode") {
		t, err := expandObjectFirewallGtpIeValidationSelectionMode2edl(d, v, "selection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selection-mode"] = t
		}
	}

	if v, ok := d.GetOk("uli"); ok || d.HasChange("uli") {
		t, err := expandObjectFirewallGtpIeValidationUli2edl(d, v, "uli")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uli"] = t
		}
	}

	return &obj, nil
}
