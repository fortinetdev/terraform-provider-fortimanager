// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: configure server info.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectLogNpuServerServerInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectLogNpuServerServerInfoCreate,
		Read:   resourceObjectLogNpuServerServerInfoRead,
		Update: resourceObjectLogNpuServerServerInfoUpdate,
		Delete: resourceObjectLogNpuServerServerInfoDelete,

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
			"dest_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ip_family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"template_tx_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectLogNpuServerServerInfoCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectLogNpuServerServerInfo(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectLogNpuServerServerInfo resource while getting object: %v", err)
	}

	_, err = c.CreateObjectLogNpuServerServerInfo(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectLogNpuServerServerInfo resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectLogNpuServerServerInfoRead(d, m)
}

func resourceObjectLogNpuServerServerInfoUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectLogNpuServerServerInfo(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectLogNpuServerServerInfo resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectLogNpuServerServerInfo(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectLogNpuServerServerInfo resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectLogNpuServerServerInfoRead(d, m)
}

func resourceObjectLogNpuServerServerInfoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectLogNpuServerServerInfo(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectLogNpuServerServerInfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectLogNpuServerServerInfoRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectLogNpuServerServerInfo(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectLogNpuServerServerInfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectLogNpuServerServerInfo(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectLogNpuServerServerInfo resource from API: %v", err)
	}
	return nil
}

func flattenObjectLogNpuServerServerInfoDestPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpv4Server(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoIpv6Server(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoTemplateTxTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectLogNpuServerServerInfoVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectLogNpuServerServerInfo(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dest_port", flattenObjectLogNpuServerServerInfoDestPort(o["dest-port"], d, "dest_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["dest-port"], "ObjectLogNpuServerServerInfo-DestPort"); ok {
			if err = d.Set("dest_port", vv); err != nil {
				return fmt.Errorf("Error reading dest_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dest_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectLogNpuServerServerInfoId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectLogNpuServerServerInfo-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_family", flattenObjectLogNpuServerServerInfoIpFamily(o["ip-family"], d, "ip_family")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-family"], "ObjectLogNpuServerServerInfo-IpFamily"); ok {
			if err = d.Set("ip_family", vv); err != nil {
				return fmt.Errorf("Error reading ip_family: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_family: %v", err)
		}
	}

	if err = d.Set("ipv4_server", flattenObjectLogNpuServerServerInfoIpv4Server(o["ipv4-server"], d, "ipv4_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-server"], "ObjectLogNpuServerServerInfo-Ipv4Server"); ok {
			if err = d.Set("ipv4_server", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_server: %v", err)
		}
	}

	if err = d.Set("ipv6_server", flattenObjectLogNpuServerServerInfoIpv6Server(o["ipv6-server"], d, "ipv6_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-server"], "ObjectLogNpuServerServerInfo-Ipv6Server"); ok {
			if err = d.Set("ipv6_server", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_server: %v", err)
		}
	}

	if err = d.Set("source_port", flattenObjectLogNpuServerServerInfoSourcePort(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "ObjectLogNpuServerServerInfo-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("template_tx_timeout", flattenObjectLogNpuServerServerInfoTemplateTxTimeout(o["template-tx-timeout"], d, "template_tx_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["template-tx-timeout"], "ObjectLogNpuServerServerInfo-TemplateTxTimeout"); ok {
			if err = d.Set("template_tx_timeout", vv); err != nil {
				return fmt.Errorf("Error reading template_tx_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template_tx_timeout: %v", err)
		}
	}

	if err = d.Set("vdom", flattenObjectLogNpuServerServerInfoVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "ObjectLogNpuServerServerInfo-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenObjectLogNpuServerServerInfoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectLogNpuServerServerInfoDestPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpv4Server(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoIpv6Server(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoSourcePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoTemplateTxTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectLogNpuServerServerInfoVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectLogNpuServerServerInfo(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dest_port"); ok || d.HasChange("dest_port") {
		t, err := expandObjectLogNpuServerServerInfoDestPort(d, v, "dest_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dest-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandObjectLogNpuServerServerInfoId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_family"); ok || d.HasChange("ip_family") {
		t, err := expandObjectLogNpuServerServerInfoIpFamily(d, v, "ip_family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-family"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_server"); ok || d.HasChange("ipv4_server") {
		t, err := expandObjectLogNpuServerServerInfoIpv4Server(d, v, "ipv4_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-server"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_server"); ok || d.HasChange("ipv6_server") {
		t, err := expandObjectLogNpuServerServerInfoIpv6Server(d, v, "ipv6_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-server"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandObjectLogNpuServerServerInfoSourcePort(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	if v, ok := d.GetOk("template_tx_timeout"); ok || d.HasChange("template_tx_timeout") {
		t, err := expandObjectLogNpuServerServerInfoTemplateTxTimeout(d, v, "template_tx_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template-tx-timeout"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandObjectLogNpuServerServerInfoVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
