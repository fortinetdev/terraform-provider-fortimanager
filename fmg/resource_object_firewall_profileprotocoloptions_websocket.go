// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WebSocket protocol options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallProfileProtocolOptionsWebsocket() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallProfileProtocolOptionsWebsocketUpdate,
		Read:   resourceObjectFirewallProfileProtocolOptionsWebsocketRead,
		Update: resourceObjectFirewallProfileProtocolOptionsWebsocketUpdate,
		Delete: resourceObjectFirewallProfileProtocolOptionsWebsocketDelete,

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
			"scan_bzip2": &schema.Schema{
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
			"tunnel_non_websocket": &schema.Schema{
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

func resourceObjectFirewallProfileProtocolOptionsWebsocketUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallProfileProtocolOptionsWebsocket(d, false)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsWebsocket resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallProfileProtocolOptionsWebsocket(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsWebsocket resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallProfileProtocolOptionsWebsocket")

	return resourceObjectFirewallProfileProtocolOptionsWebsocketRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsWebsocketDelete(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallProfileProtocolOptionsWebsocket(d, true)

	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsWebsocket resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallProfileProtocolOptionsWebsocket(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing ObjectFirewallProfileProtocolOptionsWebsocket resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallProfileProtocolOptionsWebsocketRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallProfileProtocolOptionsWebsocket(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsWebsocket resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallProfileProtocolOptionsWebsocket(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsWebsocket resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketComfortAmount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketComfortInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketScanBzip22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketStreamBasedUncompressedLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketTcpWindowMaximum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketTcpWindowMinimum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketTcpWindowSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketTcpWindowType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketTunnelNonWebsocket2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketUncompressedNestLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketUncompressedOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallProfileProtocolOptionsWebsocket(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comfort_amount", flattenObjectFirewallProfileProtocolOptionsWebsocketComfortAmount2edl(o["comfort-amount"], d, "comfort_amount")); err != nil {
		if vv, ok := fortiAPIPatch(o["comfort-amount"], "ObjectFirewallProfileProtocolOptionsWebsocket-ComfortAmount"); ok {
			if err = d.Set("comfort_amount", vv); err != nil {
				return fmt.Errorf("Error reading comfort_amount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comfort_amount: %v", err)
		}
	}

	if err = d.Set("comfort_interval", flattenObjectFirewallProfileProtocolOptionsWebsocketComfortInterval2edl(o["comfort-interval"], d, "comfort_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["comfort-interval"], "ObjectFirewallProfileProtocolOptionsWebsocket-ComfortInterval"); ok {
			if err = d.Set("comfort_interval", vv); err != nil {
				return fmt.Errorf("Error reading comfort_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comfort_interval: %v", err)
		}
	}

	if err = d.Set("options", flattenObjectFirewallProfileProtocolOptionsWebsocketOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "ObjectFirewallProfileProtocolOptionsWebsocket-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("oversize_limit", flattenObjectFirewallProfileProtocolOptionsWebsocketOversizeLimit2edl(o["oversize-limit"], d, "oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-limit"], "ObjectFirewallProfileProtocolOptionsWebsocket-OversizeLimit"); ok {
			if err = d.Set("oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_limit: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", flattenObjectFirewallProfileProtocolOptionsWebsocketScanBzip22edl(o["scan-bzip2"], d, "scan_bzip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-bzip2"], "ObjectFirewallProfileProtocolOptionsWebsocket-ScanBzip2"); ok {
			if err = d.Set("scan_bzip2", vv); err != nil {
				return fmt.Errorf("Error reading scan_bzip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_bzip2: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallProfileProtocolOptionsWebsocketStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallProfileProtocolOptionsWebsocket-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("stream_based_uncompressed_limit", flattenObjectFirewallProfileProtocolOptionsWebsocketStreamBasedUncompressedLimit2edl(o["stream-based-uncompressed-limit"], d, "stream_based_uncompressed_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["stream-based-uncompressed-limit"], "ObjectFirewallProfileProtocolOptionsWebsocket-StreamBasedUncompressedLimit"); ok {
			if err = d.Set("stream_based_uncompressed_limit", vv); err != nil {
				return fmt.Errorf("Error reading stream_based_uncompressed_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stream_based_uncompressed_limit: %v", err)
		}
	}

	if err = d.Set("tcp_window_maximum", flattenObjectFirewallProfileProtocolOptionsWebsocketTcpWindowMaximum2edl(o["tcp-window-maximum"], d, "tcp_window_maximum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-maximum"], "ObjectFirewallProfileProtocolOptionsWebsocket-TcpWindowMaximum"); ok {
			if err = d.Set("tcp_window_maximum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
		}
	}

	if err = d.Set("tcp_window_minimum", flattenObjectFirewallProfileProtocolOptionsWebsocketTcpWindowMinimum2edl(o["tcp-window-minimum"], d, "tcp_window_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-minimum"], "ObjectFirewallProfileProtocolOptionsWebsocket-TcpWindowMinimum"); ok {
			if err = d.Set("tcp_window_minimum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
		}
	}

	if err = d.Set("tcp_window_size", flattenObjectFirewallProfileProtocolOptionsWebsocketTcpWindowSize2edl(o["tcp-window-size"], d, "tcp_window_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-size"], "ObjectFirewallProfileProtocolOptionsWebsocket-TcpWindowSize"); ok {
			if err = d.Set("tcp_window_size", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_size: %v", err)
		}
	}

	if err = d.Set("tcp_window_type", flattenObjectFirewallProfileProtocolOptionsWebsocketTcpWindowType2edl(o["tcp-window-type"], d, "tcp_window_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-type"], "ObjectFirewallProfileProtocolOptionsWebsocket-TcpWindowType"); ok {
			if err = d.Set("tcp_window_type", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_type: %v", err)
		}
	}

	if err = d.Set("tunnel_non_websocket", flattenObjectFirewallProfileProtocolOptionsWebsocketTunnelNonWebsocket2edl(o["tunnel-non-websocket"], d, "tunnel_non_websocket")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-non-websocket"], "ObjectFirewallProfileProtocolOptionsWebsocket-TunnelNonWebsocket"); ok {
			if err = d.Set("tunnel_non_websocket", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_non_websocket: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_non_websocket: %v", err)
		}
	}

	if err = d.Set("uncompressed_nest_limit", flattenObjectFirewallProfileProtocolOptionsWebsocketUncompressedNestLimit2edl(o["uncompressed-nest-limit"], d, "uncompressed_nest_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-nest-limit"], "ObjectFirewallProfileProtocolOptionsWebsocket-UncompressedNestLimit"); ok {
			if err = d.Set("uncompressed_nest_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
		}
	}

	if err = d.Set("uncompressed_oversize_limit", flattenObjectFirewallProfileProtocolOptionsWebsocketUncompressedOversizeLimit2edl(o["uncompressed-oversize-limit"], d, "uncompressed_oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-oversize-limit"], "ObjectFirewallProfileProtocolOptionsWebsocket-UncompressedOversizeLimit"); ok {
			if err = d.Set("uncompressed_oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallProfileProtocolOptionsWebsocketFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallProfileProtocolOptionsWebsocketComfortAmount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketComfortInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketScanBzip22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketStreamBasedUncompressedLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketTcpWindowMaximum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketTcpWindowMinimum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketTcpWindowSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketTcpWindowType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketTunnelNonWebsocket2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketUncompressedNestLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsWebsocketUncompressedOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallProfileProtocolOptionsWebsocket(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comfort_amount"); ok || d.HasChange("comfort_amount") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketComfortAmount2edl(d, v, "comfort_amount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comfort-amount"] = t
		}
	}

	if v, ok := d.GetOk("comfort_interval"); ok || d.HasChange("comfort_interval") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketComfortInterval2edl(d, v, "comfort_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comfort-interval"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("oversize_limit"); ok || d.HasChange("oversize_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketOversizeLimit2edl(d, v, "oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("scan_bzip2"); ok || d.HasChange("scan_bzip2") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketScanBzip22edl(d, v, "scan_bzip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-bzip2"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("stream_based_uncompressed_limit"); ok || d.HasChange("stream_based_uncompressed_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketStreamBasedUncompressedLimit2edl(d, v, "stream_based_uncompressed_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stream-based-uncompressed-limit"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_maximum"); ok || d.HasChange("tcp_window_maximum") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketTcpWindowMaximum2edl(d, v, "tcp_window_maximum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-maximum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_minimum"); ok || d.HasChange("tcp_window_minimum") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketTcpWindowMinimum2edl(d, v, "tcp_window_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-minimum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_size"); ok || d.HasChange("tcp_window_size") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketTcpWindowSize2edl(d, v, "tcp_window_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-size"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_type"); ok || d.HasChange("tcp_window_type") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketTcpWindowType2edl(d, v, "tcp_window_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-type"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_non_websocket"); ok || d.HasChange("tunnel_non_websocket") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketTunnelNonWebsocket2edl(d, v, "tunnel_non_websocket")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-non-websocket"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_nest_limit"); ok || d.HasChange("uncompressed_nest_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketUncompressedNestLimit2edl(d, v, "uncompressed_nest_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-nest-limit"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_oversize_limit"); ok || d.HasChange("uncompressed_oversize_limit") {
		t, err := expandObjectFirewallProfileProtocolOptionsWebsocketUncompressedOversizeLimit2edl(d, v, "uncompressed_oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-oversize-limit"] = t
		}
	}

	return &obj, nil
}
