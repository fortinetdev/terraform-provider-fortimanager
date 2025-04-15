// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Organization identifier list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListCreate,
		Read:   resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListRead,
		Update: resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListUpdate,
		Delete: resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListDelete,

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
			"anqp_roaming_consortium": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"oi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListCreate(d *schema.ResourceData, m interface{}) error {
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

	anqp_roaming_consortium := d.Get("anqp_roaming_consortium").(string)
	paradict["anqp_roaming_consortium"] = anqp_roaming_consortium

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListUpdate(d *schema.ResourceData, m interface{}) error {
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

	anqp_roaming_consortium := d.Get("anqp_roaming_consortium").(string)
	paradict["anqp_roaming_consortium"] = anqp_roaming_consortium

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListDelete(d *schema.ResourceData, m interface{}) error {
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

	anqp_roaming_consortium := d.Get("anqp_roaming_consortium").(string)
	paradict["anqp_roaming_consortium"] = anqp_roaming_consortium

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListRead(d *schema.ResourceData, m interface{}) error {
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

	anqp_roaming_consortium := d.Get("anqp_roaming_consortium").(string)
	if anqp_roaming_consortium == "" {
		anqp_roaming_consortium = importOptionChecking(m.(*FortiClient).Cfg, "anqp_roaming_consortium")
		if anqp_roaming_consortium == "" {
			return fmt.Errorf("Parameter anqp_roaming_consortium is missing")
		}
		if err = d.Set("anqp_roaming_consortium", anqp_roaming_consortium); err != nil {
			return fmt.Errorf("Error set params anqp_roaming_consortium: %v", err)
		}
	}
	paradict["anqp_roaming_consortium"] = anqp_roaming_consortium

	o, err := c.ReadObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comment", flattenObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment2edl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("index", flattenObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("oi", flattenObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi2edl(o["oi"], d, "oi")); err != nil {
		if vv, ok := fortiAPIPatch(o["oi"], "ObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList-Oi"); ok {
			if err = d.Set("oi", vv); err != nil {
				return fmt.Errorf("Error reading oi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oi: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment2edl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("oi"); ok || d.HasChange("oi") {
		t, err := expandObjectWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi2edl(d, v, "oi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oi"] = t
		}
	}

	return &obj, nil
}
