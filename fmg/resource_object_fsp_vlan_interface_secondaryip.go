// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Second IP address of interface.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanInterfaceSecondaryip() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanInterfaceSecondaryipCreate,
		Read:   resourceObjectFspVlanInterfaceSecondaryipRead,
		Update: resourceObjectFspVlanInterfaceSecondaryipUpdate,
		Delete: resourceObjectFspVlanInterfaceSecondaryipDelete,

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
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"detectprotocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"detectserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gwdetect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secip_relay_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ping_serv_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"seq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectFspVlanInterfaceSecondaryipCreate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceSecondaryip(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceSecondaryip resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFspVlanInterfaceSecondaryip(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceSecondaryip resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanInterfaceSecondaryipRead(d, m)
}

func resourceObjectFspVlanInterfaceSecondaryipUpdate(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceSecondaryip(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceSecondaryip resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFspVlanInterfaceSecondaryip(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceSecondaryip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFspVlanInterfaceSecondaryipRead(d, m)
}

func resourceObjectFspVlanInterfaceSecondaryipDelete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	wsParams["adom"] = adomv

	err = c.DeleteObjectFspVlanInterfaceSecondaryip(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanInterfaceSecondaryip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanInterfaceSecondaryipRead(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanInterfaceSecondaryip(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceSecondaryip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanInterfaceSecondaryip(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceSecondaryip resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanInterfaceSecondaryipAllowaccess3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceSecondaryipDetectprotocol3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectFspVlanInterfaceSecondaryipDetectserver3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipGwdetect3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipHaPriority3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipSecipRelayIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipPingServStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceSecondaryipSeq3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanInterfaceSecondaryip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("allowaccess", flattenObjectFspVlanInterfaceSecondaryipAllowaccess3rdl(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ObjectFspVlanInterfaceSecondaryip-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("detectprotocol", flattenObjectFspVlanInterfaceSecondaryipDetectprotocol3rdl(o["detectprotocol"], d, "detectprotocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["detectprotocol"], "ObjectFspVlanInterfaceSecondaryip-Detectprotocol"); ok {
			if err = d.Set("detectprotocol", vv); err != nil {
				return fmt.Errorf("Error reading detectprotocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detectprotocol: %v", err)
		}
	}

	if err = d.Set("detectserver", flattenObjectFspVlanInterfaceSecondaryipDetectserver3rdl(o["detectserver"], d, "detectserver")); err != nil {
		if vv, ok := fortiAPIPatch(o["detectserver"], "ObjectFspVlanInterfaceSecondaryip-Detectserver"); ok {
			if err = d.Set("detectserver", vv); err != nil {
				return fmt.Errorf("Error reading detectserver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading detectserver: %v", err)
		}
	}

	if err = d.Set("gwdetect", flattenObjectFspVlanInterfaceSecondaryipGwdetect3rdl(o["gwdetect"], d, "gwdetect")); err != nil {
		if vv, ok := fortiAPIPatch(o["gwdetect"], "ObjectFspVlanInterfaceSecondaryip-Gwdetect"); ok {
			if err = d.Set("gwdetect", vv); err != nil {
				return fmt.Errorf("Error reading gwdetect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gwdetect: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenObjectFspVlanInterfaceSecondaryipHaPriority3rdl(o["ha-priority"], d, "ha_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-priority"], "ObjectFspVlanInterfaceSecondaryip-HaPriority"); ok {
			if err = d.Set("ha_priority", vv); err != nil {
				return fmt.Errorf("Error reading ha_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFspVlanInterfaceSecondaryipId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFspVlanInterfaceSecondaryip-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectFspVlanInterfaceSecondaryipIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectFspVlanInterfaceSecondaryip-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("secip_relay_ip", flattenObjectFspVlanInterfaceSecondaryipSecipRelayIp3rdl(o["secip-relay-ip"], d, "secip_relay_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["secip-relay-ip"], "ObjectFspVlanInterfaceSecondaryip-SecipRelayIp"); ok {
			if err = d.Set("secip_relay_ip", vv); err != nil {
				return fmt.Errorf("Error reading secip_relay_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secip_relay_ip: %v", err)
		}
	}

	if err = d.Set("ping_serv_status", flattenObjectFspVlanInterfaceSecondaryipPingServStatus3rdl(o["ping-serv-status"], d, "ping_serv_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["ping-serv-status"], "ObjectFspVlanInterfaceSecondaryip-PingServStatus"); ok {
			if err = d.Set("ping_serv_status", vv); err != nil {
				return fmt.Errorf("Error reading ping_serv_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ping_serv_status: %v", err)
		}
	}

	if err = d.Set("seq", flattenObjectFspVlanInterfaceSecondaryipSeq3rdl(o["seq"], d, "seq")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq"], "ObjectFspVlanInterfaceSecondaryip-Seq"); ok {
			if err = d.Set("seq", vv); err != nil {
				return fmt.Errorf("Error reading seq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanInterfaceSecondaryipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanInterfaceSecondaryipAllowaccess3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceSecondaryipDetectprotocol3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectFspVlanInterfaceSecondaryipDetectserver3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipGwdetect3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipHaPriority3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipSecipRelayIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipPingServStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceSecondaryipSeq3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanInterfaceSecondaryip(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandObjectFspVlanInterfaceSecondaryipAllowaccess3rdl(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("detectprotocol"); ok || d.HasChange("detectprotocol") {
		t, err := expandObjectFspVlanInterfaceSecondaryipDetectprotocol3rdl(d, v, "detectprotocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detectprotocol"] = t
		}
	}

	if v, ok := d.GetOk("detectserver"); ok || d.HasChange("detectserver") {
		t, err := expandObjectFspVlanInterfaceSecondaryipDetectserver3rdl(d, v, "detectserver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detectserver"] = t
		}
	}

	if v, ok := d.GetOk("gwdetect"); ok || d.HasChange("gwdetect") {
		t, err := expandObjectFspVlanInterfaceSecondaryipGwdetect3rdl(d, v, "gwdetect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gwdetect"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok || d.HasChange("ha_priority") {
		t, err := expandObjectFspVlanInterfaceSecondaryipHaPriority3rdl(d, v, "ha_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFspVlanInterfaceSecondaryipId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectFspVlanInterfaceSecondaryipIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("secip_relay_ip"); ok || d.HasChange("secip_relay_ip") {
		t, err := expandObjectFspVlanInterfaceSecondaryipSecipRelayIp3rdl(d, v, "secip_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secip-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("ping_serv_status"); ok || d.HasChange("ping_serv_status") {
		t, err := expandObjectFspVlanInterfaceSecondaryipPingServStatus3rdl(d, v, "ping_serv_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-serv-status"] = t
		}
	}

	if v, ok := d.GetOk("seq"); ok || d.HasChange("seq") {
		t, err := expandObjectFspVlanInterfaceSecondaryipSeq3rdl(d, v, "seq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq"] = t
		}
	}

	return &obj, nil
}
