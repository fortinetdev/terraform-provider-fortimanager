// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSU NAI list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListCreate,
		Read:   resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListRead,
		Update: resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListUpdate,
		Delete: resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListDelete,

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
			"h2qp_osu_provider_nai": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"osu_nai": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListCreate(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider_nai := d.Get("h2qp_osu_provider_nai").(string)
	paradict["h2qp_osu_provider_nai"] = h2qp_osu_provider_nai

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListUpdate(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider_nai := d.Get("h2qp_osu_provider_nai").(string)
	paradict["h2qp_osu_provider_nai"] = h2qp_osu_provider_nai

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListDelete(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider_nai := d.Get("h2qp_osu_provider_nai").(string)
	paradict["h2qp_osu_provider_nai"] = h2qp_osu_provider_nai

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListRead(d *schema.ResourceData, m interface{}) error {
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

	h2qp_osu_provider_nai := d.Get("h2qp_osu_provider_nai").(string)
	if h2qp_osu_provider_nai == "" {
		h2qp_osu_provider_nai = importOptionChecking(m.(*FortiClient).Cfg, "h2qp_osu_provider_nai")
		if h2qp_osu_provider_nai == "" {
			return fmt.Errorf("Parameter h2qp_osu_provider_nai is missing")
		}
		if err = d.Set("h2qp_osu_provider_nai", h2qp_osu_provider_nai); err != nil {
			return fmt.Errorf("Error set params h2qp_osu_provider_nai: %v", err)
		}
	}
	paradict["h2qp_osu_provider_nai"] = h2qp_osu_provider_nai

	o, err := c.ReadObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("osu_nai", flattenObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai2edl(o["osu-nai"], d, "osu_nai")); err != nil {
		if vv, ok := fortiAPIPatch(o["osu-nai"], "ObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList-OsuNai"); ok {
			if err = d.Set("osu_nai", vv); err != nil {
				return fmt.Errorf("Error reading osu_nai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading osu_nai: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("osu_nai"); ok || d.HasChange("osu_nai") {
		t, err := expandObjectWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai2edl(d, v, "osu_nai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-nai"] = t
		}
	}

	return &obj, nil
}
