// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure decrypted traffic mirror.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallDecryptedTrafficMirror() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallDecryptedTrafficMirrorCreate,
		Read:   resourceObjectFirewallDecryptedTrafficMirrorRead,
		Update: resourceObjectFirewallDecryptedTrafficMirrorUpdate,
		Delete: resourceObjectFirewallDecryptedTrafficMirrorDelete,

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
			"dstmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"traffic_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallDecryptedTrafficMirrorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectFirewallDecryptedTrafficMirror(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallDecryptedTrafficMirror resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallDecryptedTrafficMirror(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallDecryptedTrafficMirror resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallDecryptedTrafficMirrorRead(d, m)
}

func resourceObjectFirewallDecryptedTrafficMirrorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallDecryptedTrafficMirror(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallDecryptedTrafficMirror resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallDecryptedTrafficMirror(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallDecryptedTrafficMirror resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallDecryptedTrafficMirrorRead(d, m)
}

func resourceObjectFirewallDecryptedTrafficMirrorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallDecryptedTrafficMirror(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallDecryptedTrafficMirror resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallDecryptedTrafficMirrorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallDecryptedTrafficMirror(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallDecryptedTrafficMirror resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallDecryptedTrafficMirror(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallDecryptedTrafficMirror resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallDecryptedTrafficMirrorDstmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallDecryptedTrafficMirrorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallDecryptedTrafficMirrorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallDecryptedTrafficMirrorTrafficSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallDecryptedTrafficMirrorTrafficType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallDecryptedTrafficMirror(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dstmac", flattenObjectFirewallDecryptedTrafficMirrorDstmac(o["dstmac"], d, "dstmac")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstmac"], "ObjectFirewallDecryptedTrafficMirror-Dstmac"); ok {
			if err = d.Set("dstmac", vv); err != nil {
				return fmt.Errorf("Error reading dstmac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstmac: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectFirewallDecryptedTrafficMirrorInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectFirewallDecryptedTrafficMirror-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallDecryptedTrafficMirrorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallDecryptedTrafficMirror-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("traffic_source", flattenObjectFirewallDecryptedTrafficMirrorTrafficSource(o["traffic-source"], d, "traffic_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-source"], "ObjectFirewallDecryptedTrafficMirror-TrafficSource"); ok {
			if err = d.Set("traffic_source", vv); err != nil {
				return fmt.Errorf("Error reading traffic_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_source: %v", err)
		}
	}

	if err = d.Set("traffic_type", flattenObjectFirewallDecryptedTrafficMirrorTrafficType(o["traffic-type"], d, "traffic_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-type"], "ObjectFirewallDecryptedTrafficMirror-TrafficType"); ok {
			if err = d.Set("traffic_type", vv); err != nil {
				return fmt.Errorf("Error reading traffic_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_type: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallDecryptedTrafficMirrorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallDecryptedTrafficMirrorDstmac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallDecryptedTrafficMirrorInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallDecryptedTrafficMirrorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallDecryptedTrafficMirrorTrafficSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallDecryptedTrafficMirrorTrafficType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFirewallDecryptedTrafficMirror(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dstmac"); ok || d.HasChange("dstmac") {
		t, err := expandObjectFirewallDecryptedTrafficMirrorDstmac(d, v, "dstmac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstmac"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectFirewallDecryptedTrafficMirrorInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallDecryptedTrafficMirrorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("traffic_source"); ok || d.HasChange("traffic_source") {
		t, err := expandObjectFirewallDecryptedTrafficMirrorTrafficSource(d, v, "traffic_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-source"] = t
		}
	}

	if v, ok := d.GetOk("traffic_type"); ok || d.HasChange("traffic_type") {
		t, err := expandObjectFirewallDecryptedTrafficMirrorTrafficType(d, v, "traffic_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-type"] = t
		}
	}

	return &obj, nil
}
