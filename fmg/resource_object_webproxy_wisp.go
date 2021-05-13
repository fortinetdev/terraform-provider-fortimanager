// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Wireless Internet service provider (WISP) servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWebProxyWisp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebProxyWispCreate,
		Read:   resourceObjectWebProxyWispRead,
		Update: resourceObjectWebProxyWispUpdate,
		Delete: resourceObjectWebProxyWispDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_connections": &schema.Schema{
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
			"outgoing_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWebProxyWispCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebProxyWisp(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyWisp resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWebProxyWisp(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyWisp resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebProxyWispRead(d, m)
}

func resourceObjectWebProxyWispUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebProxyWisp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyWisp resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebProxyWisp(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyWisp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebProxyWispRead(d, m)
}

func resourceObjectWebProxyWispDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWebProxyWisp(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebProxyWisp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebProxyWispRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWebProxyWisp(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyWisp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebProxyWisp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyWisp resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebProxyWispComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyWispMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyWispName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyWispOutgoingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyWispServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyWispServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyWispTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebProxyWisp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenObjectWebProxyWispComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWebProxyWisp-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenObjectWebProxyWispMaxConnections(o["max-connections"], d, "max_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-connections"], "ObjectWebProxyWisp-MaxConnections"); ok {
			if err = d.Set("max_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWebProxyWispName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebProxyWisp-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("outgoing_ip", flattenObjectWebProxyWispOutgoingIp(o["outgoing-ip"], d, "outgoing_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["outgoing-ip"], "ObjectWebProxyWisp-OutgoingIp"); ok {
			if err = d.Set("outgoing_ip", vv); err != nil {
				return fmt.Errorf("Error reading outgoing_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outgoing_ip: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenObjectWebProxyWispServerIp(o["server-ip"], d, "server_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-ip"], "ObjectWebProxyWisp-ServerIp"); ok {
			if err = d.Set("server_ip", vv); err != nil {
				return fmt.Errorf("Error reading server_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	if err = d.Set("server_port", flattenObjectWebProxyWispServerPort(o["server-port"], d, "server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-port"], "ObjectWebProxyWisp-ServerPort"); ok {
			if err = d.Set("server_port", vv); err != nil {
				return fmt.Errorf("Error reading server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("timeout", flattenObjectWebProxyWispTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "ObjectWebProxyWisp-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func flattenObjectWebProxyWispFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebProxyWispComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyWispMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyWispName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyWispOutgoingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyWispServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyWispServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyWispTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebProxyWisp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectWebProxyWispComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok {
		t, err := expandObjectWebProxyWispMaxConnections(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWebProxyWispName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("outgoing_ip"); ok {
		t, err := expandObjectWebProxyWispOutgoingIp(d, v, "outgoing_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outgoing-ip"] = t
		}
	}

	if v, ok := d.GetOk("server_ip"); ok {
		t, err := expandObjectWebProxyWispServerIp(d, v, "server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ip"] = t
		}
	}

	if v, ok := d.GetOk("server_port"); ok {
		t, err := expandObjectWebProxyWispServerPort(d, v, "server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandObjectWebProxyWispTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	return &obj, nil
}
