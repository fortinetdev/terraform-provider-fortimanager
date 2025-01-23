// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure system NTP information.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemNtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemNtpUpdate,
		Read:   resourceSystempSystemNtpRead,
		Update: resourceSystempSystemNtpUpdate,
		Delete: resourceSystempSystemNtpDelete,

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
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"ntpserver": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
					},
				},
			},
			"ntpsync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystempSystemNtpUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemNtp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemNtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystempSystemNtp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemNtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempSystemNtp")

	return resourceSystempSystemNtpRead(d, m)
}

func resourceSystempSystemNtpDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["devprof"] = devprof

	err = c.DeleteSystempSystemNtp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemNtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemNtpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystempSystemNtp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemNtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemNtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemNtp resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemNtpAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempSystemNtpKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserver(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			v := flattenSystempSystemNtpNtpserverAuthentication(i["authentication"], d, pre_append)
			tmp["authentication"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-Authentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystempSystemNtpNtpserverId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_type"
		if _, ok := i["ip-type"]; ok {
			v := flattenSystempSystemNtpNtpserverIpType(i["ip-type"], d, pre_append)
			tmp["ip_type"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-IpType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystempSystemNtpNtpserverInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenSystempSystemNtpNtpserverInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := i["key-id"]; ok {
			v := flattenSystempSystemNtpNtpserverKeyId(i["key-id"], d, pre_append)
			tmp["key_id"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-KeyId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := i["key-type"]; ok {
			v := flattenSystempSystemNtpNtpserverKeyType(i["key-type"], d, pre_append)
			tmp["key_type"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-KeyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := i["ntpv3"]; ok {
			v := flattenSystempSystemNtpNtpserverNtpv3(i["ntpv3"], d, pre_append)
			tmp["ntpv3"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-Ntpv3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenSystempSystemNtpNtpserverServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "SystempSystemNtp-Ntpserver-Server")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystempSystemNtpNtpserverAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverIpType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverNtpv3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpserverServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpNtpsync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpServerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpSyncinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemNtpType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempSystemNtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("authentication", flattenSystempSystemNtpAuthentication(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "SystempSystemNtp-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystempSystemNtpInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystempSystemNtp-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("key_id", flattenSystempSystemNtpKeyId(o["key-id"], d, "key_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-id"], "SystempSystemNtp-KeyId"); ok {
			if err = d.Set("key_id", vv); err != nil {
				return fmt.Errorf("Error reading key_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_id: %v", err)
		}
	}

	if err = d.Set("key_type", flattenSystempSystemNtpKeyType(o["key-type"], d, "key_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-type"], "SystempSystemNtp-KeyType"); ok {
			if err = d.Set("key_type", vv); err != nil {
				return fmt.Errorf("Error reading key_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ntpserver", flattenSystempSystemNtpNtpserver(o["ntpserver"], d, "ntpserver")); err != nil {
			if vv, ok := fortiAPIPatch(o["ntpserver"], "SystempSystemNtp-Ntpserver"); ok {
				if err = d.Set("ntpserver", vv); err != nil {
					return fmt.Errorf("Error reading ntpserver: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ntpserver: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ntpserver"); ok {
			if err = d.Set("ntpserver", flattenSystempSystemNtpNtpserver(o["ntpserver"], d, "ntpserver")); err != nil {
				if vv, ok := fortiAPIPatch(o["ntpserver"], "SystempSystemNtp-Ntpserver"); ok {
					if err = d.Set("ntpserver", vv); err != nil {
						return fmt.Errorf("Error reading ntpserver: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ntpserver: %v", err)
				}
			}
		}
	}

	if err = d.Set("ntpsync", flattenSystempSystemNtpNtpsync(o["ntpsync"], d, "ntpsync")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntpsync"], "SystempSystemNtp-Ntpsync"); ok {
			if err = d.Set("ntpsync", vv); err != nil {
				return fmt.Errorf("Error reading ntpsync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntpsync: %v", err)
		}
	}

	if err = d.Set("server_mode", flattenSystempSystemNtpServerMode(o["server-mode"], d, "server_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-mode"], "SystempSystemNtp-ServerMode"); ok {
			if err = d.Set("server_mode", vv); err != nil {
				return fmt.Errorf("Error reading server_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_mode: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystempSystemNtpSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystempSystemNtp-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystempSystemNtpSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "SystempSystemNtp-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("syncinterval", flattenSystempSystemNtpSyncinterval(o["syncinterval"], d, "syncinterval")); err != nil {
		if vv, ok := fortiAPIPatch(o["syncinterval"], "SystempSystemNtp-Syncinterval"); ok {
			if err = d.Set("syncinterval", vv); err != nil {
				return fmt.Errorf("Error reading syncinterval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syncinterval: %v", err)
		}
	}

	if err = d.Set("type", flattenSystempSystemNtpType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystempSystemNtp-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemNtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemNtpAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemNtpKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemNtpKeyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpKeyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authentication"], _ = expandSystempSystemNtpNtpserverAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystempSystemNtpNtpserverId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip-type"], _ = expandSystempSystemNtpNtpserverIpType(d, i["ip_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystempSystemNtpNtpserverInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandSystempSystemNtpNtpserverInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key"], _ = expandSystempSystemNtpNtpserverKey(d, i["key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-id"], _ = expandSystempSystemNtpNtpserverKeyId(d, i["key_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-type"], _ = expandSystempSystemNtpNtpserverKeyType(d, i["key_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ntpv3"], _ = expandSystempSystemNtpNtpserverNtpv3(d, i["ntpv3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandSystempSystemNtpNtpserverServer(d, i["server"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystempSystemNtpNtpserverAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverIpType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemNtpNtpserverKeyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverKeyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverNtpv3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpserverServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpNtpsync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpServerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpSyncinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemNtpType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempSystemNtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandSystempSystemNtpAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystempSystemNtpInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandSystempSystemNtpKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("key_id"); ok || d.HasChange("key_id") {
		t, err := expandSystempSystemNtpKeyId(d, v, "key_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-id"] = t
		}
	}

	if v, ok := d.GetOk("key_type"); ok || d.HasChange("key_type") {
		t, err := expandSystempSystemNtpKeyType(d, v, "key_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-type"] = t
		}
	}

	if v, ok := d.GetOk("ntpserver"); ok || d.HasChange("ntpserver") {
		t, err := expandSystempSystemNtpNtpserver(d, v, "ntpserver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntpserver"] = t
		}
	}

	if v, ok := d.GetOk("ntpsync"); ok || d.HasChange("ntpsync") {
		t, err := expandSystempSystemNtpNtpsync(d, v, "ntpsync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntpsync"] = t
		}
	}

	if v, ok := d.GetOk("server_mode"); ok || d.HasChange("server_mode") {
		t, err := expandSystempSystemNtpServerMode(d, v, "server_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-mode"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystempSystemNtpSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandSystempSystemNtpSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("syncinterval"); ok || d.HasChange("syncinterval") {
		t, err := expandSystempSystemNtpSyncinterval(d, v, "syncinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syncinterval"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystempSystemNtpType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
