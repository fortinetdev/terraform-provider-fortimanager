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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectExtenderControllerDataplan() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerDataplanCreate,
		Read:   resourceObjectExtenderControllerDataplanRead,
		Update: resourceObjectExtenderControllerDataplanUpdate,
		Delete: resourceObjectExtenderControllerDataplanDelete,

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
				Computed: true,
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
				Computed: true,
			},
			"carrier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iccid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthly_fee": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
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
				Computed: true,
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
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtenderControllerDataplanCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectExtenderControllerDataplan(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerDataplan resource while getting object: %v", err)
	}

	_, err = c.CreateObjectExtenderControllerDataplan(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerDataplan resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerDataplanRead(d, m)
}

func resourceObjectExtenderControllerDataplanUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectExtenderControllerDataplan(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerDataplan resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtenderControllerDataplan(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerDataplan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerDataplanRead(d, m)
}

func resourceObjectExtenderControllerDataplanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectExtenderControllerDataplan(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerDataplan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerDataplanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectExtenderControllerDataplan(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerDataplan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerDataplan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerDataplan resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerDataplanApn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanBillingDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanIccid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanMonthlyFee(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanOverage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerDataplanPdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanPreferredSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanPrivateNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanSignalPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanSignalThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanSlot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerDataplanUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerDataplan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("apn", flattenObjectExtenderControllerDataplanApn(o["apn"], d, "apn")); err != nil {
		if vv, ok := fortiAPIPatch(o["apn"], "ObjectExtenderControllerDataplan-Apn"); ok {
			if err = d.Set("apn", vv); err != nil {
				return fmt.Errorf("Error reading apn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apn: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenObjectExtenderControllerDataplanAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "ObjectExtenderControllerDataplan-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("billing_date", flattenObjectExtenderControllerDataplanBillingDate(o["billing-date"], d, "billing_date")); err != nil {
		if vv, ok := fortiAPIPatch(o["billing-date"], "ObjectExtenderControllerDataplan-BillingDate"); ok {
			if err = d.Set("billing_date", vv); err != nil {
				return fmt.Errorf("Error reading billing_date: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading billing_date: %v", err)
		}
	}

	if err = d.Set("capacity", flattenObjectExtenderControllerDataplanCapacity(o["capacity"], d, "capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["capacity"], "ObjectExtenderControllerDataplan-Capacity"); ok {
			if err = d.Set("capacity", vv); err != nil {
				return fmt.Errorf("Error reading capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capacity: %v", err)
		}
	}

	if err = d.Set("carrier", flattenObjectExtenderControllerDataplanCarrier(o["carrier"], d, "carrier")); err != nil {
		if vv, ok := fortiAPIPatch(o["carrier"], "ObjectExtenderControllerDataplan-Carrier"); ok {
			if err = d.Set("carrier", vv); err != nil {
				return fmt.Errorf("Error reading carrier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading carrier: %v", err)
		}
	}

	if err = d.Set("iccid", flattenObjectExtenderControllerDataplanIccid(o["iccid"], d, "iccid")); err != nil {
		if vv, ok := fortiAPIPatch(o["iccid"], "ObjectExtenderControllerDataplan-Iccid"); ok {
			if err = d.Set("iccid", vv); err != nil {
				return fmt.Errorf("Error reading iccid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iccid: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenObjectExtenderControllerDataplanModemId(o["modem-id"], d, "modem_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-id"], "ObjectExtenderControllerDataplan-ModemId"); ok {
			if err = d.Set("modem_id", vv); err != nil {
				return fmt.Errorf("Error reading modem_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("monthly_fee", flattenObjectExtenderControllerDataplanMonthlyFee(o["monthly-fee"], d, "monthly_fee")); err != nil {
		if vv, ok := fortiAPIPatch(o["monthly-fee"], "ObjectExtenderControllerDataplan-MonthlyFee"); ok {
			if err = d.Set("monthly_fee", vv); err != nil {
				return fmt.Errorf("Error reading monthly_fee: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monthly_fee: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtenderControllerDataplanName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtenderControllerDataplan-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("overage", flattenObjectExtenderControllerDataplanOverage(o["overage"], d, "overage")); err != nil {
		if vv, ok := fortiAPIPatch(o["overage"], "ObjectExtenderControllerDataplan-Overage"); ok {
			if err = d.Set("overage", vv); err != nil {
				return fmt.Errorf("Error reading overage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overage: %v", err)
		}
	}

	if err = d.Set("pdn", flattenObjectExtenderControllerDataplanPdn(o["pdn"], d, "pdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["pdn"], "ObjectExtenderControllerDataplan-Pdn"); ok {
			if err = d.Set("pdn", vv); err != nil {
				return fmt.Errorf("Error reading pdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pdn: %v", err)
		}
	}

	if err = d.Set("preferred_subnet", flattenObjectExtenderControllerDataplanPreferredSubnet(o["preferred-subnet"], d, "preferred_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-subnet"], "ObjectExtenderControllerDataplan-PreferredSubnet"); ok {
			if err = d.Set("preferred_subnet", vv); err != nil {
				return fmt.Errorf("Error reading preferred_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_subnet: %v", err)
		}
	}

	if err = d.Set("private_network", flattenObjectExtenderControllerDataplanPrivateNetwork(o["private-network"], d, "private_network")); err != nil {
		if vv, ok := fortiAPIPatch(o["private-network"], "ObjectExtenderControllerDataplan-PrivateNetwork"); ok {
			if err = d.Set("private_network", vv); err != nil {
				return fmt.Errorf("Error reading private_network: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_network: %v", err)
		}
	}

	if err = d.Set("signal_period", flattenObjectExtenderControllerDataplanSignalPeriod(o["signal-period"], d, "signal_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-period"], "ObjectExtenderControllerDataplan-SignalPeriod"); ok {
			if err = d.Set("signal_period", vv); err != nil {
				return fmt.Errorf("Error reading signal_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_period: %v", err)
		}
	}

	if err = d.Set("signal_threshold", flattenObjectExtenderControllerDataplanSignalThreshold(o["signal-threshold"], d, "signal_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-threshold"], "ObjectExtenderControllerDataplan-SignalThreshold"); ok {
			if err = d.Set("signal_threshold", vv); err != nil {
				return fmt.Errorf("Error reading signal_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_threshold: %v", err)
		}
	}

	if err = d.Set("slot", flattenObjectExtenderControllerDataplanSlot(o["slot"], d, "slot")); err != nil {
		if vv, ok := fortiAPIPatch(o["slot"], "ObjectExtenderControllerDataplan-Slot"); ok {
			if err = d.Set("slot", vv); err != nil {
				return fmt.Errorf("Error reading slot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slot: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectExtenderControllerDataplanStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectExtenderControllerDataplan-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectExtenderControllerDataplanType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectExtenderControllerDataplan-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username", flattenObjectExtenderControllerDataplanUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "ObjectExtenderControllerDataplan-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerDataplanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerDataplanApn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanBillingDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanIccid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanMonthlyFee(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanOverage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerDataplanPdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanPreferredSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanPrivateNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanSignalPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanSignalThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanSlot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerDataplanUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerDataplan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apn"); ok {
		t, err := expandObjectExtenderControllerDataplanApn(d, v, "apn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apn"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandObjectExtenderControllerDataplanAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("billing_date"); ok {
		t, err := expandObjectExtenderControllerDataplanBillingDate(d, v, "billing_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-date"] = t
		}
	}

	if v, ok := d.GetOk("capacity"); ok {
		t, err := expandObjectExtenderControllerDataplanCapacity(d, v, "capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capacity"] = t
		}
	}

	if v, ok := d.GetOk("carrier"); ok {
		t, err := expandObjectExtenderControllerDataplanCarrier(d, v, "carrier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["carrier"] = t
		}
	}

	if v, ok := d.GetOk("iccid"); ok {
		t, err := expandObjectExtenderControllerDataplanIccid(d, v, "iccid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iccid"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok {
		t, err := expandObjectExtenderControllerDataplanModemId(d, v, "modem_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("monthly_fee"); ok {
		t, err := expandObjectExtenderControllerDataplanMonthlyFee(d, v, "monthly_fee")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monthly-fee"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectExtenderControllerDataplanName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("overage"); ok {
		t, err := expandObjectExtenderControllerDataplanOverage(d, v, "overage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overage"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandObjectExtenderControllerDataplanPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("pdn"); ok {
		t, err := expandObjectExtenderControllerDataplanPdn(d, v, "pdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pdn"] = t
		}
	}

	if v, ok := d.GetOk("preferred_subnet"); ok {
		t, err := expandObjectExtenderControllerDataplanPreferredSubnet(d, v, "preferred_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-subnet"] = t
		}
	}

	if v, ok := d.GetOk("private_network"); ok {
		t, err := expandObjectExtenderControllerDataplanPrivateNetwork(d, v, "private_network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-network"] = t
		}
	}

	if v, ok := d.GetOk("signal_period"); ok {
		t, err := expandObjectExtenderControllerDataplanSignalPeriod(d, v, "signal_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-period"] = t
		}
	}

	if v, ok := d.GetOk("signal_threshold"); ok {
		t, err := expandObjectExtenderControllerDataplanSignalThreshold(d, v, "signal_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-threshold"] = t
		}
	}

	if v, ok := d.GetOk("slot"); ok {
		t, err := expandObjectExtenderControllerDataplanSlot(d, v, "slot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slot"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandObjectExtenderControllerDataplanStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandObjectExtenderControllerDataplanType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandObjectExtenderControllerDataplanUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
