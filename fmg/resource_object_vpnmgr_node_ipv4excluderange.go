// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectVpnmgr NodeIpv4ExcludeRange

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnmgrNodeIpv4ExcludeRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnmgrNodeIpv4ExcludeRangeCreate,
		Read:   resourceObjectVpnmgrNodeIpv4ExcludeRangeRead,
		Update: resourceObjectVpnmgrNodeIpv4ExcludeRangeUpdate,
		Delete: resourceObjectVpnmgrNodeIpv4ExcludeRangeDelete,

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
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectVpnmgrNodeIpv4ExcludeRangeCreate(d *schema.ResourceData, m interface{}) error {
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

	node := d.Get("node").(string)
	paradict["node"] = node

	obj, err := getObjectObjectVpnmgrNodeIpv4ExcludeRange(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrNodeIpv4ExcludeRange resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectVpnmgrNodeIpv4ExcludeRange(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrNodeIpv4ExcludeRange resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVpnmgrNodeIpv4ExcludeRangeRead(d, m)
}

func resourceObjectVpnmgrNodeIpv4ExcludeRangeUpdate(d *schema.ResourceData, m interface{}) error {
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

	node := d.Get("node").(string)
	paradict["node"] = node

	obj, err := getObjectObjectVpnmgrNodeIpv4ExcludeRange(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrNodeIpv4ExcludeRange resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVpnmgrNodeIpv4ExcludeRange(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrNodeIpv4ExcludeRange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectVpnmgrNodeIpv4ExcludeRangeRead(d, m)
}

func resourceObjectVpnmgrNodeIpv4ExcludeRangeDelete(d *schema.ResourceData, m interface{}) error {
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

	node := d.Get("node").(string)
	paradict["node"] = node

	wsParams["adom"] = adomv

	err = c.DeleteObjectVpnmgrNodeIpv4ExcludeRange(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnmgrNodeIpv4ExcludeRange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnmgrNodeIpv4ExcludeRangeRead(d *schema.ResourceData, m interface{}) error {
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

	node := d.Get("node").(string)
	if node == "" {
		node = importOptionChecking(m.(*FortiClient).Cfg, "node")
		if node == "" {
			return fmt.Errorf("Parameter node is missing")
		}
		if err = d.Set("node", node); err != nil {
			return fmt.Errorf("Error set params node: %v", err)
		}
	}
	paradict["node"] = node

	o, err := c.ReadObjectVpnmgrNodeIpv4ExcludeRange(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrNodeIpv4ExcludeRange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnmgrNodeIpv4ExcludeRange(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrNodeIpv4ExcludeRange resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnmgrNodeIpv4ExcludeRangeEndIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4ExcludeRangeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeIpv4ExcludeRangeStartIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnmgrNodeIpv4ExcludeRange(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("end_ip", flattenObjectVpnmgrNodeIpv4ExcludeRangeEndIp2edl(o["end-ip"], d, "end_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-ip"], "ObjectVpnmgrNodeIpv4ExcludeRange-EndIp"); ok {
			if err = d.Set("end_ip", vv); err != nil {
				return fmt.Errorf("Error reading end_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectVpnmgrNodeIpv4ExcludeRangeId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectVpnmgrNodeIpv4ExcludeRange-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenObjectVpnmgrNodeIpv4ExcludeRangeStartIp2edl(o["start-ip"], d, "start_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-ip"], "ObjectVpnmgrNodeIpv4ExcludeRange-StartIp"); ok {
			if err = d.Set("start_ip", vv); err != nil {
				return fmt.Errorf("Error reading start_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnmgrNodeIpv4ExcludeRangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnmgrNodeIpv4ExcludeRangeEndIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4ExcludeRangeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeIpv4ExcludeRangeStartIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnmgrNodeIpv4ExcludeRange(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("end_ip"); ok || d.HasChange("end_ip") {
		t, err := expandObjectVpnmgrNodeIpv4ExcludeRangeEndIp2edl(d, v, "end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectVpnmgrNodeIpv4ExcludeRangeId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok || d.HasChange("start_ip") {
		t, err := expandObjectVpnmgrNodeIpv4ExcludeRangeStartIp2edl(d, v, "start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	return &obj, nil
}
