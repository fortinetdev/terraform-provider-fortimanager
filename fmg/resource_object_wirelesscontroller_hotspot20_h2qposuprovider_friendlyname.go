// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSU provider friendly name.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameCreate,
		Read:   resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameRead,
		Update: resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameUpdate,
		Delete: resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameDelete,

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
			"h2qp_osu_provider": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameUpdate(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "index")))

	return resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameDelete(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	err = c.DeleteObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameRead(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	if h2qp_osu_provider == "" {
		h2qp_osu_provider = importOptionChecking(m.(*FortiClient).Cfg, "h2qp_osu_provider")
		if h2qp_osu_provider == "" {
			return fmt.Errorf("Parameter h2qp_osu_provider is missing")
		}
		if err = d.Set("h2qp_osu_provider", h2qp_osu_provider); err != nil {
			return fmt.Errorf("Error set params h2qp_osu_provider: %v", err)
		}
	}
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	o, err := c.ReadObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFriendlyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("friendly_name", flattenObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFriendlyName2edl(o["friendly-name"], d, "friendly_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["friendly-name"], "ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName-FriendlyName"); ok {
			if err = d.Set("friendly_name", vv); err != nil {
				return fmt.Errorf("Error reading friendly_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading friendly_name: %v", err)
		}
	}

	if err = d.Set("index", flattenObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("lang", flattenObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameLang2edl(o["lang"], d, "lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["lang"], "ObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName-Lang"); ok {
			if err = d.Set("lang", vv); err != nil {
				return fmt.Errorf("Error reading lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lang: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFriendlyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("friendly_name"); ok || d.HasChange("friendly_name") {
		t, err := expandObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameFriendlyName2edl(d, v, "friendly_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["friendly-name"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("lang"); ok || d.HasChange("lang") {
		t, err := expandObjectWirelessControllerHotspot20H2QpOsuProviderFriendlyNameLang2edl(d, v, "lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lang"] = t
		}
	}

	return &obj, nil
}
