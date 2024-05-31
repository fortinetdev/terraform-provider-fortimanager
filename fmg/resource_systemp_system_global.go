// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure global attributes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemGlobalUpdate,
		Read:   resourceSystempSystemGlobalRead,
		Update: resourceSystempSystemGlobalUpdate,
		Delete: resourceSystempSystemGlobalDelete,

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
			"admin_https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_scp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_sport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_v1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_telnet_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admintimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gui_device_latitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_device_longitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_lines_per_page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gui_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystempSystemGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystempSystemGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystempSystemGlobal(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempSystemGlobal")

	return resourceSystempSystemGlobalRead(d, m)
}

func resourceSystempSystemGlobalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystempSystemGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystempSystemGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemGlobalAdminHttpsRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalAdminPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalAdminScp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalAdminSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalAdminSshPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalAdminSshV1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalAdminTelnetPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalAdmintimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalGuiDeviceLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalGuiDeviceLongitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalGuiIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalGuiLinesPerPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalGuiTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemGlobalTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convstr2list(v, d.Get(pre))
}

func flattenSystempSystemGlobalSwitchController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempSystemGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("admin_https_redirect", flattenSystempSystemGlobalAdminHttpsRedirect(o["admin-https-redirect"], d, "admin_https_redirect")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-https-redirect"], "SystempSystemGlobal-AdminHttpsRedirect"); ok {
			if err = d.Set("admin_https_redirect", vv); err != nil {
				return fmt.Errorf("Error reading admin_https_redirect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_https_redirect: %v", err)
		}
	}

	if err = d.Set("admin_port", flattenSystempSystemGlobalAdminPort(o["admin-port"], d, "admin_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-port"], "SystempSystemGlobal-AdminPort"); ok {
			if err = d.Set("admin_port", vv); err != nil {
				return fmt.Errorf("Error reading admin_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_port: %v", err)
		}
	}

	if err = d.Set("admin_scp", flattenSystempSystemGlobalAdminScp(o["admin-scp"], d, "admin_scp")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-scp"], "SystempSystemGlobal-AdminScp"); ok {
			if err = d.Set("admin_scp", vv); err != nil {
				return fmt.Errorf("Error reading admin_scp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_scp: %v", err)
		}
	}

	if err = d.Set("admin_sport", flattenSystempSystemGlobalAdminSport(o["admin-sport"], d, "admin_sport")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-sport"], "SystempSystemGlobal-AdminSport"); ok {
			if err = d.Set("admin_sport", vv); err != nil {
				return fmt.Errorf("Error reading admin_sport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_sport: %v", err)
		}
	}

	if err = d.Set("admin_ssh_port", flattenSystempSystemGlobalAdminSshPort(o["admin-ssh-port"], d, "admin_ssh_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-ssh-port"], "SystempSystemGlobal-AdminSshPort"); ok {
			if err = d.Set("admin_ssh_port", vv); err != nil {
				return fmt.Errorf("Error reading admin_ssh_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_ssh_port: %v", err)
		}
	}

	if err = d.Set("admin_ssh_v1", flattenSystempSystemGlobalAdminSshV1(o["admin-ssh-v1"], d, "admin_ssh_v1")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-ssh-v1"], "SystempSystemGlobal-AdminSshV1"); ok {
			if err = d.Set("admin_ssh_v1", vv); err != nil {
				return fmt.Errorf("Error reading admin_ssh_v1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_ssh_v1: %v", err)
		}
	}

	if err = d.Set("admin_telnet_port", flattenSystempSystemGlobalAdminTelnetPort(o["admin-telnet-port"], d, "admin_telnet_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-telnet-port"], "SystempSystemGlobal-AdminTelnetPort"); ok {
			if err = d.Set("admin_telnet_port", vv); err != nil {
				return fmt.Errorf("Error reading admin_telnet_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_telnet_port: %v", err)
		}
	}

	if err = d.Set("admintimeout", flattenSystempSystemGlobalAdmintimeout(o["admintimeout"], d, "admintimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["admintimeout"], "SystempSystemGlobal-Admintimeout"); ok {
			if err = d.Set("admintimeout", vv); err != nil {
				return fmt.Errorf("Error reading admintimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("gui_device_latitude", flattenSystempSystemGlobalGuiDeviceLatitude(o["gui-device-latitude"], d, "gui_device_latitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-device-latitude"], "SystempSystemGlobal-GuiDeviceLatitude"); ok {
			if err = d.Set("gui_device_latitude", vv); err != nil {
				return fmt.Errorf("Error reading gui_device_latitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_device_latitude: %v", err)
		}
	}

	if err = d.Set("gui_device_longitude", flattenSystempSystemGlobalGuiDeviceLongitude(o["gui-device-longitude"], d, "gui_device_longitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-device-longitude"], "SystempSystemGlobal-GuiDeviceLongitude"); ok {
			if err = d.Set("gui_device_longitude", vv); err != nil {
				return fmt.Errorf("Error reading gui_device_longitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_device_longitude: %v", err)
		}
	}

	if err = d.Set("gui_ipv6", flattenSystempSystemGlobalGuiIpv6(o["gui-ipv6"], d, "gui_ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ipv6"], "SystempSystemGlobal-GuiIpv6"); ok {
			if err = d.Set("gui_ipv6", vv); err != nil {
				return fmt.Errorf("Error reading gui_ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ipv6: %v", err)
		}
	}

	if err = d.Set("gui_lines_per_page", flattenSystempSystemGlobalGuiLinesPerPage(o["gui-lines-per-page"], d, "gui_lines_per_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-lines-per-page"], "SystempSystemGlobal-GuiLinesPerPage"); ok {
			if err = d.Set("gui_lines_per_page", vv); err != nil {
				return fmt.Errorf("Error reading gui_lines_per_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_lines_per_page: %v", err)
		}
	}

	if err = d.Set("gui_theme", flattenSystempSystemGlobalGuiTheme(o["gui-theme"], d, "gui_theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-theme"], "SystempSystemGlobal-GuiTheme"); ok {
			if err = d.Set("gui_theme", vv); err != nil {
				return fmt.Errorf("Error reading gui_theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_theme: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystempSystemGlobalHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "SystempSystemGlobal-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("language", flattenSystempSystemGlobalLanguage(o["language"], d, "language")); err != nil {
		if vv, ok := fortiAPIPatch(o["language"], "SystempSystemGlobal-Language"); ok {
			if err = d.Set("language", vv); err != nil {
				return fmt.Errorf("Error reading language: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystempSystemGlobalTimezone(o["timezone"], d, "timezone")); err != nil {
		if vv, ok := fortiAPIPatch(o["timezone"], "SystempSystemGlobal-Timezone"); ok {
			if err = d.Set("timezone", vv); err != nil {
				return fmt.Errorf("Error reading timezone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("switch_controller", flattenSystempSystemGlobalSwitchController(o["switch-controller"], d, "switch_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller"], "SystempSystemGlobal-SwitchController"); ok {
			if err = d.Set("switch_controller", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemGlobalAdminHttpsRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalAdminPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalAdminScp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalAdminSport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalAdminSshPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalAdminSshV1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalAdminTelnetPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalAdmintimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalGuiDeviceLatitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalGuiDeviceLongitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalGuiIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalGuiLinesPerPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalGuiTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalLanguage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemGlobalTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempSystemGlobalSwitchController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempSystemGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("admin_https_redirect"); ok || d.HasChange("admin_https_redirect") {
		t, err := expandSystempSystemGlobalAdminHttpsRedirect(d, v, "admin_https_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-redirect"] = t
		}
	}

	if v, ok := d.GetOk("admin_port"); ok || d.HasChange("admin_port") {
		t, err := expandSystempSystemGlobalAdminPort(d, v, "admin_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_scp"); ok || d.HasChange("admin_scp") {
		t, err := expandSystempSystemGlobalAdminScp(d, v, "admin_scp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-scp"] = t
		}
	}

	if v, ok := d.GetOk("admin_sport"); ok || d.HasChange("admin_sport") {
		t, err := expandSystempSystemGlobalAdminSport(d, v, "admin_sport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-sport"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_port"); ok || d.HasChange("admin_ssh_port") {
		t, err := expandSystempSystemGlobalAdminSshPort(d, v, "admin_ssh_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_v1"); ok || d.HasChange("admin_ssh_v1") {
		t, err := expandSystempSystemGlobalAdminSshV1(d, v, "admin_ssh_v1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-v1"] = t
		}
	}

	if v, ok := d.GetOk("admin_telnet_port"); ok || d.HasChange("admin_telnet_port") {
		t, err := expandSystempSystemGlobalAdminTelnetPort(d, v, "admin_telnet_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-telnet-port"] = t
		}
	}

	if v, ok := d.GetOk("admintimeout"); ok || d.HasChange("admintimeout") {
		t, err := expandSystempSystemGlobalAdmintimeout(d, v, "admintimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admintimeout"] = t
		}
	}

	if v, ok := d.GetOk("gui_device_latitude"); ok || d.HasChange("gui_device_latitude") {
		t, err := expandSystempSystemGlobalGuiDeviceLatitude(d, v, "gui_device_latitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-device-latitude"] = t
		}
	}

	if v, ok := d.GetOk("gui_device_longitude"); ok || d.HasChange("gui_device_longitude") {
		t, err := expandSystempSystemGlobalGuiDeviceLongitude(d, v, "gui_device_longitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-device-longitude"] = t
		}
	}

	if v, ok := d.GetOk("gui_ipv6"); ok || d.HasChange("gui_ipv6") {
		t, err := expandSystempSystemGlobalGuiIpv6(d, v, "gui_ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("gui_lines_per_page"); ok || d.HasChange("gui_lines_per_page") {
		t, err := expandSystempSystemGlobalGuiLinesPerPage(d, v, "gui_lines_per_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-lines-per-page"] = t
		}
	}

	if v, ok := d.GetOk("gui_theme"); ok || d.HasChange("gui_theme") {
		t, err := expandSystempSystemGlobalGuiTheme(d, v, "gui_theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-theme"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandSystempSystemGlobalHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("language"); ok || d.HasChange("language") {
		t, err := expandSystempSystemGlobalLanguage(d, v, "language")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok || d.HasChange("timezone") {
		t, err := expandSystempSystemGlobalTimezone(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller"); ok || d.HasChange("switch_controller") {
		t, err := expandSystempSystemGlobalSwitchController(d, v, "switch_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller"] = t
		}
	}

	return &obj, nil
}
