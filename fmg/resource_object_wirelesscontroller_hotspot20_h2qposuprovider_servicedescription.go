// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSU service name.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionCreate,
		Read:   resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionRead,
		Update: resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionUpdate,
		Delete: resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionDelete,

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
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionCreate(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "service_id")))

	return resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionUpdate(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "service_id")))

	return resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionDelete(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider := d.Get("h2qp_osu_provider").(string)
	paradict["h2qp_osu_provider"] = h2qp_osu_provider

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("lang", flattenObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang2edl(o["lang"], d, "lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["lang"], "ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription-Lang"); ok {
			if err = d.Set("lang", vv); err != nil {
				return fmt.Errorf("Error reading lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lang: %v", err)
		}
	}

	if err = d.Set("service_description", flattenObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription2edl(o["service-description"], d, "service_description")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-description"], "ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription-ServiceDescription"); ok {
			if err = d.Set("service_description", vv); err != nil {
				return fmt.Errorf("Error reading service_description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_description: %v", err)
		}
	}

	if err = d.Set("service_id", flattenObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId2edl(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "ObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescription(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("lang"); ok || d.HasChange("lang") {
		t, err := expandObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionLang2edl(d, v, "lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lang"] = t
		}
	}

	if v, ok := d.GetOk("service_description"); ok || d.HasChange("service_description") {
		t, err := expandObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceDescription2edl(d, v, "service_description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-description"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandObjectWirelessControllerHotspot20H2QpOsuProviderServiceDescriptionServiceId2edl(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	return &obj, nil
}
