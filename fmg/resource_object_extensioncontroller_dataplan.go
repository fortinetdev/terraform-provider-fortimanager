// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender dataplan configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerDataplan() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerDataplanCreate,
		Read:   resourceObjectExtensionControllerDataplanRead,
		Update: resourceObjectExtensionControllerDataplanUpdate,
		Delete: resourceObjectExtensionControllerDataplanDelete,

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
			"apn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"billing_date": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"carrier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iccid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"modem_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthly_fee": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"overage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"pdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferred_subnet": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"private_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signal_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signal_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"slot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectExtensionControllerDataplanCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectExtensionControllerDataplan(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerDataplan resource while getting object: %v", err)
	}

	_, err = c.CreateObjectExtensionControllerDataplan(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerDataplan resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerDataplanRead(d, m)
}

func resourceObjectExtensionControllerDataplanUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerDataplan(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerDataplan resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtensionControllerDataplan(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerDataplan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerDataplanRead(d, m)
}

func resourceObjectExtensionControllerDataplanDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerDataplan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerDataplan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerDataplanRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerDataplan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerDataplan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerDataplan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerDataplan resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerDataplanApn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanBillingDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanIccid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanMonthlyFee(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanOverage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanPdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanPreferredSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanPrivateNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanSignalPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanSignalThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanSlot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerDataplanUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerDataplan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("apn", flattenObjectExtensionControllerDataplanApn(o["apn"], d, "apn")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn"], "ObjectExtensionControllerDataplan-Apn"); ok {
			if err = d.Set("apn", vv); err != nil {
				return fmt.Errorf("Error reading apn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenObjectExtensionControllerDataplanAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "ObjectExtensionControllerDataplan-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("billing_date", flattenObjectExtensionControllerDataplanBillingDate(o["billing-date"], d, "billing_date")); err != nil {
		if vv, ok := fortiAPIPatch(o["billing-date"], "ObjectExtensionControllerDataplan-BillingDate"); ok {
			if err = d.Set("billing_date", vv); err != nil {
				return fmt.Errorf("Error reading billing_date: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading billing_date: %v", err)
		}
	}

	if err = d.Set("capacity", flattenObjectExtensionControllerDataplanCapacity(o["capacity"], d, "capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["capacity"], "ObjectExtensionControllerDataplan-Capacity"); ok {
			if err = d.Set("capacity", vv); err != nil {
				return fmt.Errorf("Error reading capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capacity: %v", err)
		}
	}

	if err = d.Set("carrier", flattenObjectExtensionControllerDataplanCarrier(o["carrier"], d, "carrier")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrier"], "ObjectExtensionControllerDataplan-Carrier"); ok {
			if err = d.Set("carrier", vv); err != nil {
				return fmt.Errorf("Error reading carrier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrier: %v", err)
		}
	}

	if err = d.Set("iccid", flattenObjectExtensionControllerDataplanIccid(o["iccid"], d, "iccid")); err != nil {
		if vv, ok := fortiAPIPatch(o["iccid"], "ObjectExtensionControllerDataplan-Iccid"); ok {
			if err = d.Set("iccid", vv); err != nil {
				return fmt.Errorf("Error reading iccid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iccid: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenObjectExtensionControllerDataplanModemId(o["modem-id"], d, "modem_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-id"], "ObjectExtensionControllerDataplan-ModemId"); ok {
			if err = d.Set("modem_id", vv); err != nil {
				return fmt.Errorf("Error reading modem_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("monthly_fee", flattenObjectExtensionControllerDataplanMonthlyFee(o["monthly-fee"], d, "monthly_fee")); err != nil {
		if vv, ok := fortiAPIPatch(o["monthly-fee"], "ObjectExtensionControllerDataplan-MonthlyFee"); ok {
			if err = d.Set("monthly_fee", vv); err != nil {
				return fmt.Errorf("Error reading monthly_fee: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monthly_fee: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtensionControllerDataplanName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtensionControllerDataplan-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("overage", flattenObjectExtensionControllerDataplanOverage(o["overage"], d, "overage")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage"], "ObjectExtensionControllerDataplan-Overage"); ok {
			if err = d.Set("overage", vv); err != nil {
				return fmt.Errorf("Error reading overage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage: %v", err)
		}
	}

	if err = d.Set("pdn", flattenObjectExtensionControllerDataplanPdn(o["pdn"], d, "pdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdn"], "ObjectExtensionControllerDataplan-Pdn"); ok {
			if err = d.Set("pdn", vv); err != nil {
				return fmt.Errorf("Error reading pdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdn: %v", err)
		}
	}

	if err = d.Set("preferred_subnet", flattenObjectExtensionControllerDataplanPreferredSubnet(o["preferred-subnet"], d, "preferred_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-subnet"], "ObjectExtensionControllerDataplan-PreferredSubnet"); ok {
			if err = d.Set("preferred_subnet", vv); err != nil {
				return fmt.Errorf("Error reading preferred_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_subnet: %v", err)
		}
	}

	if err = d.Set("private_network", flattenObjectExtensionControllerDataplanPrivateNetwork(o["private-network"], d, "private_network")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-network"], "ObjectExtensionControllerDataplan-PrivateNetwork"); ok {
			if err = d.Set("private_network", vv); err != nil {
				return fmt.Errorf("Error reading private_network: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_network: %v", err)
		}
	}

	if err = d.Set("signal_period", flattenObjectExtensionControllerDataplanSignalPeriod(o["signal-period"], d, "signal_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-period"], "ObjectExtensionControllerDataplan-SignalPeriod"); ok {
			if err = d.Set("signal_period", vv); err != nil {
				return fmt.Errorf("Error reading signal_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_period: %v", err)
		}
	}

	if err = d.Set("signal_threshold", flattenObjectExtensionControllerDataplanSignalThreshold(o["signal-threshold"], d, "signal_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-threshold"], "ObjectExtensionControllerDataplan-SignalThreshold"); ok {
			if err = d.Set("signal_threshold", vv); err != nil {
				return fmt.Errorf("Error reading signal_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_threshold: %v", err)
		}
	}

	if err = d.Set("slot", flattenObjectExtensionControllerDataplanSlot(o["slot"], d, "slot")); err != nil {
		if vv, ok := fortiAPIPatch(o["slot"], "ObjectExtensionControllerDataplan-Slot"); ok {
			if err = d.Set("slot", vv); err != nil {
				return fmt.Errorf("Error reading slot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slot: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectExtensionControllerDataplanType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectExtensionControllerDataplan-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectExtensionControllerDataplanUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectExtensionControllerDataplan-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerDataplanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerDataplanApn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanBillingDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanIccid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanMonthlyFee(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanOverage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerDataplanPdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanPreferredSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanPrivateNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanSignalPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanSignalThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanSlot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerDataplanUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerDataplan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apn"); ok || d.HasChange("apn") {
		t, err := expandObjectExtensionControllerDataplanApn(d, v, "apn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandObjectExtensionControllerDataplanAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("billing_date"); ok || d.HasChange("billing_date") {
		t, err := expandObjectExtensionControllerDataplanBillingDate(d, v, "billing_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-date"] = t
		}
	}

	if v, ok := d.GetOk("capacity"); ok || d.HasChange("capacity") {
		t, err := expandObjectExtensionControllerDataplanCapacity(d, v, "capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capacity"] = t
		}
	}

	if v, ok := d.GetOk("carrier"); ok || d.HasChange("carrier") {
		t, err := expandObjectExtensionControllerDataplanCarrier(d, v, "carrier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier"] = t
		}
	}

	if v, ok := d.GetOk("iccid"); ok || d.HasChange("iccid") {
		t, err := expandObjectExtensionControllerDataplanIccid(d, v, "iccid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iccid"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok || d.HasChange("modem_id") {
		t, err := expandObjectExtensionControllerDataplanModemId(d, v, "modem_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("monthly_fee"); ok || d.HasChange("monthly_fee") {
		t, err := expandObjectExtensionControllerDataplanMonthlyFee(d, v, "monthly_fee")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monthly-fee"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtensionControllerDataplanName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("overage"); ok || d.HasChange("overage") {
		t, err := expandObjectExtensionControllerDataplanOverage(d, v, "overage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandObjectExtensionControllerDataplanPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("pdn"); ok || d.HasChange("pdn") {
		t, err := expandObjectExtensionControllerDataplanPdn(d, v, "pdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdn"] = t
		}
	}

	if v, ok := d.GetOk("preferred_subnet"); ok || d.HasChange("preferred_subnet") {
		t, err := expandObjectExtensionControllerDataplanPreferredSubnet(d, v, "preferred_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-subnet"] = t
		}
	}

	if v, ok := d.GetOk("private_network"); ok || d.HasChange("private_network") {
		t, err := expandObjectExtensionControllerDataplanPrivateNetwork(d, v, "private_network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-network"] = t
		}
	}

	if v, ok := d.GetOk("signal_period"); ok || d.HasChange("signal_period") {
		t, err := expandObjectExtensionControllerDataplanSignalPeriod(d, v, "signal_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-period"] = t
		}
	}

	if v, ok := d.GetOk("signal_threshold"); ok || d.HasChange("signal_threshold") {
		t, err := expandObjectExtensionControllerDataplanSignalThreshold(d, v, "signal_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-threshold"] = t
		}
	}

	if v, ok := d.GetOk("slot"); ok || d.HasChange("slot") {
		t, err := expandObjectExtensionControllerDataplanSlot(d, v, "slot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slot"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectExtensionControllerDataplanType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandObjectExtensionControllerDataplanUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
