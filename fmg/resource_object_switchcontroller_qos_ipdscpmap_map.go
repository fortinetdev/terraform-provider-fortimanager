// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Maps between IP-DSCP value to COS queue.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerQosIpDscpMapMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerQosIpDscpMapMapCreate,
		Read:   resourceObjectSwitchControllerQosIpDscpMapMapRead,
		Update: resourceObjectSwitchControllerQosIpDscpMapMapUpdate,
		Delete: resourceObjectSwitchControllerQosIpDscpMapMapDelete,

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
			"ip_dscp_map": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cos_queue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_precedence": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerQosIpDscpMapMapCreate(d *schema.ResourceData, m interface{}) error {
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

	ip_dscp_map := d.Get("ip_dscp_map").(string)
	paradict["ip_dscp_map"] = ip_dscp_map

	obj, err := getObjectObjectSwitchControllerQosIpDscpMapMap(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosIpDscpMapMap resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerQosIpDscpMapMap(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerQosIpDscpMapMap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosIpDscpMapMapRead(d, m)
}

func resourceObjectSwitchControllerQosIpDscpMapMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	ip_dscp_map := d.Get("ip_dscp_map").(string)
	paradict["ip_dscp_map"] = ip_dscp_map

	obj, err := getObjectObjectSwitchControllerQosIpDscpMapMap(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosIpDscpMapMap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerQosIpDscpMapMap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerQosIpDscpMapMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerQosIpDscpMapMapRead(d, m)
}

func resourceObjectSwitchControllerQosIpDscpMapMapDelete(d *schema.ResourceData, m interface{}) error {
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

	ip_dscp_map := d.Get("ip_dscp_map").(string)
	paradict["ip_dscp_map"] = ip_dscp_map

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerQosIpDscpMapMap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerQosIpDscpMapMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerQosIpDscpMapMapRead(d *schema.ResourceData, m interface{}) error {
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

	ip_dscp_map := d.Get("ip_dscp_map").(string)
	if ip_dscp_map == "" {
		ip_dscp_map = importOptionChecking(m.(*FortiClient).Cfg, "ip_dscp_map")
		if ip_dscp_map == "" {
			return fmt.Errorf("Parameter ip_dscp_map is missing")
		}
		if err = d.Set("ip_dscp_map", ip_dscp_map); err != nil {
			return fmt.Errorf("Error set params ip_dscp_map: %v", err)
		}
	}
	paradict["ip_dscp_map"] = ip_dscp_map

	o, err := c.ReadObjectSwitchControllerQosIpDscpMapMap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosIpDscpMapMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerQosIpDscpMapMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerQosIpDscpMapMap resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerQosIpDscpMapMapCosQueue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosIpDscpMapMapDiffserv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerQosIpDscpMapMapIpPrecedence2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerQosIpDscpMapMapName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerQosIpDscpMapMapValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerQosIpDscpMapMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cos_queue", flattenObjectSwitchControllerQosIpDscpMapMapCosQueue2edl(o["cos-queue"], d, "cos_queue")); err != nil {
		if vv, ok := fortiAPIPatch(o["cos-queue"], "ObjectSwitchControllerQosIpDscpMapMap-CosQueue"); ok {
			if err = d.Set("cos_queue", vv); err != nil {
				return fmt.Errorf("Error reading cos_queue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cos_queue: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenObjectSwitchControllerQosIpDscpMapMapDiffserv2edl(o["diffserv"], d, "diffserv")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv"], "ObjectSwitchControllerQosIpDscpMapMap-Diffserv"); ok {
			if err = d.Set("diffserv", vv); err != nil {
				return fmt.Errorf("Error reading diffserv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("ip_precedence", flattenObjectSwitchControllerQosIpDscpMapMapIpPrecedence2edl(o["ip-precedence"], d, "ip_precedence")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-precedence"], "ObjectSwitchControllerQosIpDscpMapMap-IpPrecedence"); ok {
			if err = d.Set("ip_precedence", vv); err != nil {
				return fmt.Errorf("Error reading ip_precedence: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_precedence: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerQosIpDscpMapMapName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerQosIpDscpMapMap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenObjectSwitchControllerQosIpDscpMapMapValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ObjectSwitchControllerQosIpDscpMapMap-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerQosIpDscpMapMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerQosIpDscpMapMapCosQueue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosIpDscpMapMapDiffserv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerQosIpDscpMapMapIpPrecedence2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerQosIpDscpMapMapName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerQosIpDscpMapMapValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerQosIpDscpMapMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cos_queue"); ok || d.HasChange("cos_queue") {
		t, err := expandObjectSwitchControllerQosIpDscpMapMapCosQueue2edl(d, v, "cos_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok || d.HasChange("diffserv") {
		t, err := expandObjectSwitchControllerQosIpDscpMapMapDiffserv2edl(d, v, "diffserv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("ip_precedence"); ok || d.HasChange("ip_precedence") {
		t, err := expandObjectSwitchControllerQosIpDscpMapMapIpPrecedence2edl(d, v, "ip_precedence")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-precedence"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerQosIpDscpMapMapName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandObjectSwitchControllerQosIpDscpMapMapValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
