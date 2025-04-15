// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FTP protocol options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallProfileProtocolOptionsFtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallProfileProtocolOptionsFtpUpdate,
		Read:   resourceObjectFirewallProfileProtocolOptionsFtpRead,
		Update: resourceObjectFirewallProfileProtocolOptionsFtpUpdate,
		Delete: resourceObjectFirewallProfileProtocolOptionsFtpDelete,

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comfort_amount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comfort_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"explicit_ftp_tls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inspect_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"scan_bzip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_offloaded": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stream_based_uncompressed_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_window_maximum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_window_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_window_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uncompressed_nest_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uncompressed_oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallProfileProtocolOptionsFtpUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectObjectFirewallProfileProtocolOptionsFtp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsFtp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallProfileProtocolOptionsFtp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsFtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallProfileProtocolOptionsFtp")

	return resourceObjectFirewallProfileProtocolOptionsFtpRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsFtpDelete(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallProfileProtocolOptionsFtp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallProfileProtocolOptionsFtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallProfileProtocolOptionsFtpRead(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadObjectFirewallProfileProtocolOptionsFtp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsFtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallProfileProtocolOptionsFtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsFtp resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallProfileProtocolOptionsFtpComfortAmount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpComfortInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpExplicitFtpTls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpInspectAll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsFtpOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsFtpScanBzip22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpSslOffloaded2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowMaximum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowMinimum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpUncompressedNestLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallProfileProtocolOptionsFtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comfort_amount", flattenObjectFirewallProfileProtocolOptionsFtpComfortAmount2edl(o["comfort-amount"], d, "comfort_amount")); err != nil {
		if vv, ok := fortiAPIPatch(o["comfort-amount"], "ObjectFirewallProfileProtocolOptionsFtp-ComfortAmount"); ok {
			if err = d.Set("comfort_amount", vv); err != nil {
				return fmt.Errorf("Error reading comfort_amount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comfort_amount: %v", err)
		}
	}

	if err = d.Set("comfort_interval", flattenObjectFirewallProfileProtocolOptionsFtpComfortInterval2edl(o["comfort-interval"], d, "comfort_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["comfort-interval"], "ObjectFirewallProfileProtocolOptionsFtp-ComfortInterval"); ok {
			if err = d.Set("comfort_interval", vv); err != nil {
				return fmt.Errorf("Error reading comfort_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comfort_interval: %v", err)
		}
	}

	if err = d.Set("explicit_ftp_tls", flattenObjectFirewallProfileProtocolOptionsFtpExplicitFtpTls2edl(o["explicit-ftp-tls"], d, "explicit_ftp_tls")); err != nil {
		if vv, ok := fortiAPIPatch(o["explicit-ftp-tls"], "ObjectFirewallProfileProtocolOptionsFtp-ExplicitFtpTls"); ok {
			if err = d.Set("explicit_ftp_tls", vv); err != nil {
				return fmt.Errorf("Error reading explicit_ftp_tls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading explicit_ftp_tls: %v", err)
		}
	}

	if err = d.Set("inspect_all", flattenObjectFirewallProfileProtocolOptionsFtpInspectAll2edl(o["inspect-all"], d, "inspect_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["inspect-all"], "ObjectFirewallProfileProtocolOptionsFtp-InspectAll"); ok {
			if err = d.Set("inspect_all", vv); err != nil {
				return fmt.Errorf("Error reading inspect_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading inspect_all: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectFirewallProfileProtocolOptionsFtpOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectFirewallProfileProtocolOptionsFtp-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("oversize_limit", flattenObjectFirewallProfileProtocolOptionsFtpOversizeLimit2edl(o["oversize-limit"], d, "oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-limit"], "ObjectFirewallProfileProtocolOptionsFtp-OversizeLimit"); ok {
			if err = d.Set("oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_limit: %v", err)
		}
	}

	if err = d.Set("ports", flattenObjectFirewallProfileProtocolOptionsFtpPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "ObjectFirewallProfileProtocolOptionsFtp-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", flattenObjectFirewallProfileProtocolOptionsFtpScanBzip22edl(o["scan-bzip2"], d, "scan_bzip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-bzip2"], "ObjectFirewallProfileProtocolOptionsFtp-ScanBzip2"); ok {
			if err = d.Set("scan_bzip2", vv); err != nil {
				return fmt.Errorf("Error reading scan_bzip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_bzip2: %v", err)
		}
	}

	if err = d.Set("ssl_offloaded", flattenObjectFirewallProfileProtocolOptionsFtpSslOffloaded2edl(o["ssl-offloaded"], d, "ssl_offloaded")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-offloaded"], "ObjectFirewallProfileProtocolOptionsFtp-SslOffloaded"); ok {
			if err = d.Set("ssl_offloaded", vv); err != nil {
				return fmt.Errorf("Error reading ssl_offloaded: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_offloaded: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallProfileProtocolOptionsFtpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallProfileProtocolOptionsFtp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("stream_based_uncompressed_limit", flattenObjectFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit2edl(o["stream-based-uncompressed-limit"], d, "stream_based_uncompressed_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["stream-based-uncompressed-limit"], "ObjectFirewallProfileProtocolOptionsFtp-StreamBasedUncompressedLimit"); ok {
			if err = d.Set("stream_based_uncompressed_limit", vv); err != nil {
				return fmt.Errorf("Error reading stream_based_uncompressed_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stream_based_uncompressed_limit: %v", err)
		}
	}

	if err = d.Set("tcp_window_maximum", flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowMaximum2edl(o["tcp-window-maximum"], d, "tcp_window_maximum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-maximum"], "ObjectFirewallProfileProtocolOptionsFtp-TcpWindowMaximum"); ok {
			if err = d.Set("tcp_window_maximum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
		}
	}

	if err = d.Set("tcp_window_minimum", flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowMinimum2edl(o["tcp-window-minimum"], d, "tcp_window_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-minimum"], "ObjectFirewallProfileProtocolOptionsFtp-TcpWindowMinimum"); ok {
			if err = d.Set("tcp_window_minimum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
		}
	}

	if err = d.Set("tcp_window_size", flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowSize2edl(o["tcp-window-size"], d, "tcp_window_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-size"], "ObjectFirewallProfileProtocolOptionsFtp-TcpWindowSize"); ok {
			if err = d.Set("tcp_window_size", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_size: %v", err)
		}
	}

	if err = d.Set("tcp_window_type", flattenObjectFirewallProfileProtocolOptionsFtpTcpWindowType2edl(o["tcp-window-type"], d, "tcp_window_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-type"], "ObjectFirewallProfileProtocolOptionsFtp-TcpWindowType"); ok {
			if err = d.Set("tcp_window_type", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_type: %v", err)
		}
	}

	if err = d.Set("uncompressed_nest_limit", flattenObjectFirewallProfileProtocolOptionsFtpUncompressedNestLimit2edl(o["uncompressed-nest-limit"], d, "uncompressed_nest_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-nest-limit"], "ObjectFirewallProfileProtocolOptionsFtp-UncompressedNestLimit"); ok {
			if err = d.Set("uncompressed_nest_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
		}
	}

	if err = d.Set("uncompressed_oversize_limit", flattenObjectFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit2edl(o["uncompressed-oversize-limit"], d, "uncompressed_oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-oversize-limit"], "ObjectFirewallProfileProtocolOptionsFtp-UncompressedOversizeLimit"); ok {
			if err = d.Set("uncompressed_oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallProfileProtocolOptionsFtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallProfileProtocolOptionsFtpComfortAmount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpComfortInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpExplicitFtpTls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpInspectAll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsFtpOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsFtpScanBzip22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpSslOffloaded2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpTcpWindowMaximum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpTcpWindowMinimum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpTcpWindowSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpTcpWindowType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpUncompressedNestLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallProfileProtocolOptionsFtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comfort_amount"); ok || d.HasChange("comfort_amount") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpComfortAmount2edl(d, v, "comfort_amount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comfort-amount"] = t
		}
	}

	if v, ok := d.GetOk("comfort_interval"); ok || d.HasChange("comfort_interval") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpComfortInterval2edl(d, v, "comfort_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comfort-interval"] = t
		}
	}

	if v, ok := d.GetOk("explicit_ftp_tls"); ok || d.HasChange("explicit_ftp_tls") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpExplicitFtpTls2edl(d, v, "explicit_ftp_tls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-ftp-tls"] = t
		}
	}

	if v, ok := d.GetOk("inspect_all"); ok || d.HasChange("inspect_all") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpInspectAll2edl(d, v, "inspect_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspect-all"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("oversize_limit"); ok || d.HasChange("oversize_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpOversizeLimit2edl(d, v, "oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("scan_bzip2"); ok || d.HasChange("scan_bzip2") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpScanBzip22edl(d, v, "scan_bzip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-bzip2"] = t
		}
	}

	if v, ok := d.GetOk("ssl_offloaded"); ok || d.HasChange("ssl_offloaded") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpSslOffloaded2edl(d, v, "ssl_offloaded")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-offloaded"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("stream_based_uncompressed_limit"); ok || d.HasChange("stream_based_uncompressed_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit2edl(d, v, "stream_based_uncompressed_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stream-based-uncompressed-limit"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_maximum"); ok || d.HasChange("tcp_window_maximum") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpTcpWindowMaximum2edl(d, v, "tcp_window_maximum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-maximum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_minimum"); ok || d.HasChange("tcp_window_minimum") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpTcpWindowMinimum2edl(d, v, "tcp_window_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-minimum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_size"); ok || d.HasChange("tcp_window_size") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpTcpWindowSize2edl(d, v, "tcp_window_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-size"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_type"); ok || d.HasChange("tcp_window_type") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpTcpWindowType2edl(d, v, "tcp_window_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-type"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_nest_limit"); ok || d.HasChange("uncompressed_nest_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpUncompressedNestLimit2edl(d, v, "uncompressed_nest_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-nest-limit"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_oversize_limit"); ok || d.HasChange("uncompressed_oversize_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit2edl(d, v, "uncompressed_oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-oversize-limit"] = t
		}
	}

	return &obj, nil
}
