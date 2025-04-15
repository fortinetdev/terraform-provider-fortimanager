// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the FortiGate to connect to any available third-party NTP server.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemNtpNtpserver() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemNtpNtpserverCreate,
		Read:   resourceSystempSystemNtpNtpserverRead,
		Update: resourceSystempSystemNtpNtpserverUpdate,
		Delete: resourceSystempSystemNtpNtpserverDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"key_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntpv3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystempSystemNtpNtpserverCreate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemNtpNtpserver(d)
	if err != nil {
		return fmt.Errorf("Error creating SystempSystemNtpNtpserver resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystempSystemNtpNtpserver(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystempSystemNtpNtpserver resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystempSystemNtpNtpserverRead(d, m)
}

func resourceSystempSystemNtpNtpserverUpdate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemNtpNtpserver(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemNtpNtpserver resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempSystemNtpNtpserver(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemNtpNtpserver resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystempSystemNtpNtpserverRead(d, m)
}

func resourceSystempSystemNtpNtpserverDelete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	wsParams["adom"] = adomv

	err = c.DeleteSystempSystemNtpNtpserver(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemNtpNtpserver resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemNtpNtpserverRead(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	paradict["devprof"] = devprof

	o, err := c.ReadSystempSystemNtpNtpserver(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemNtpNtpserver resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemNtpNtpserver(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemNtpNtpserver resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemNtpNtpserverAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverIpType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverInterfaceSelectMethod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverKeyId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverKeyType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverNtpv32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverVrfSelect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempSystemNtpNtpserver(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("authentication", flattenSystempSystemNtpNtpserverAuthentication2edl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "SystempSystemNtpNtpserver-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystempSystemNtpNtpserverId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystempSystemNtpNtpserver-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_type", flattenSystempSystemNtpNtpserverIpType2edl(o["ip-type"], d, "ip_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-type"], "SystempSystemNtpNtpserver-IpType"); ok {
			if err = d.Set("ip_type", vv); err != nil {
				return fmt.Errorf("Error reading ip_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_type: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystempSystemNtpNtpserverInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystempSystemNtpNtpserver-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystempSystemNtpNtpserverInterfaceSelectMethod2edl(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "SystempSystemNtpNtpserver-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("key_id", flattenSystempSystemNtpNtpserverKeyId2edl(o["key-id"], d, "key_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-id"], "SystempSystemNtpNtpserver-KeyId"); ok {
			if err = d.Set("key_id", vv); err != nil {
				return fmt.Errorf("Error reading key_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_id: %v", err)
		}
	}

	if err = d.Set("key_type", flattenSystempSystemNtpNtpserverKeyType2edl(o["key-type"], d, "key_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-type"], "SystempSystemNtpNtpserver-KeyType"); ok {
			if err = d.Set("key_type", vv); err != nil {
				return fmt.Errorf("Error reading key_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if err = d.Set("ntpv3", flattenSystempSystemNtpNtpserverNtpv32edl(o["ntpv3"], d, "ntpv3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntpv3"], "SystempSystemNtpNtpserver-Ntpv3"); ok {
			if err = d.Set("ntpv3", vv); err != nil {
				return fmt.Errorf("Error reading ntpv3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntpv3: %v", err)
		}
	}

	if err = d.Set("server", flattenSystempSystemNtpNtpserverServer2edl(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystempSystemNtpNtpserver-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystempSystemNtpNtpserverVrfSelect2edl(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "SystempSystemNtpNtpserver-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemNtpNtpserverFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemNtpNtpserverAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverIpType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverInterfaceSelectMethod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemNtpNtpserverKeyId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverKeyType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverNtpv32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverVrfSelect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempSystemNtpNtpserver(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandSystempSystemNtpNtpserverAuthentication2edl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystempSystemNtpNtpserverId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_type"); ok || d.HasChange("ip_type") {
		t, err := expandSystempSystemNtpNtpserverIpType2edl(d, v, "ip_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-type"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystempSystemNtpNtpserverInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandSystempSystemNtpNtpserverInterfaceSelectMethod2edl(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandSystempSystemNtpNtpserverKey2edl(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("key_id"); ok || d.HasChange("key_id") {
		t, err := expandSystempSystemNtpNtpserverKeyId2edl(d, v, "key_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-id"] = t
		}
	}

	if v, ok := d.GetOk("key_type"); ok || d.HasChange("key_type") {
		t, err := expandSystempSystemNtpNtpserverKeyType2edl(d, v, "key_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-type"] = t
		}
	}

	if v, ok := d.GetOk("ntpv3"); ok || d.HasChange("ntpv3") {
		t, err := expandSystempSystemNtpNtpserverNtpv32edl(d, v, "ntpv3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntpv3"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystempSystemNtpNtpserverServer2edl(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandSystempSystemNtpNtpserverVrfSelect2edl(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
