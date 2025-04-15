// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: EAP auth param.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamCreate,
		Read:   resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamRead,
		Update: resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamUpdate,
		Delete: resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamDelete,

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
			"anqp_nai_realm": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nai_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"eap_method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamCreate(d *schema.ResourceData, m interface{}) error {
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

	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
	eap_method := d.Get("eap_method").(string)
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list
	paradict["eap_method"] = eap_method

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamUpdate(d *schema.ResourceData, m interface{}) error {
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

	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
	eap_method := d.Get("eap_method").(string)
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list
	paradict["eap_method"] = eap_method

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamDelete(d *schema.ResourceData, m interface{}) error {
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

	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
	eap_method := d.Get("eap_method").(string)
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list
	paradict["eap_method"] = eap_method

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamRead(d *schema.ResourceData, m interface{}) error {
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

	anqp_nai_realm := d.Get("anqp_nai_realm").(string)
	nai_list := d.Get("nai_list").(string)
	eap_method := d.Get("eap_method").(string)
	if anqp_nai_realm == "" {
		anqp_nai_realm = importOptionChecking(m.(*FortiClient).Cfg, "anqp_nai_realm")
		if anqp_nai_realm == "" {
			return fmt.Errorf("Parameter anqp_nai_realm is missing")
		}
		if err = d.Set("anqp_nai_realm", anqp_nai_realm); err != nil {
			return fmt.Errorf("Error set params anqp_nai_realm: %v", err)
		}
	}
	if nai_list == "" {
		nai_list = importOptionChecking(m.(*FortiClient).Cfg, "nai_list")
		if nai_list == "" {
			return fmt.Errorf("Parameter nai_list is missing")
		}
		if err = d.Set("nai_list", nai_list); err != nil {
			return fmt.Errorf("Error set params nai_list: %v", err)
		}
	}
	if eap_method == "" {
		eap_method = importOptionChecking(m.(*FortiClient).Cfg, "eap_method")
		if eap_method == "" {
			return fmt.Errorf("Parameter eap_method is missing")
		}
		if err = d.Set("eap_method", eap_method); err != nil {
			return fmt.Errorf("Error set params eap_method: %v", err)
		}
	}
	paradict["anqp_nai_realm"] = anqp_nai_realm
	paradict["nai_list"] = nai_list
	paradict["eap_method"] = eap_method

	o, err := c.ReadObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("index", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex4thl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("val", flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal4thl(o["val"], d, "val")); err != nil {
		if vv, ok := fortiAPIPatch(o["val"], "ObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam-Val"); ok {
			if err = d.Set("val", vv); err != nil {
				return fmt.Errorf("Error reading val: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading val: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParam(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamIndex4thl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("val"); ok || d.HasChange("val") {
		t, err := expandObjectWirelessControllerHotspot20AnqpNaiRealmNaiListEapMethodAuthParamVal4thl(d, v, "val")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["val"] = t
		}
	}

	return &obj, nil
}
