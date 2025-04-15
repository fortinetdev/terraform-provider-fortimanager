// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: WTP, FortiAP, or AP platform.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerWtpProfilePlatform() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerWtpProfilePlatformUpdate,
		Read:   resourceObjectWirelessControllerWtpProfilePlatformRead,
		Update: resourceObjectWirelessControllerWtpProfilePlatformUpdate,
		Delete: resourceObjectWirelessControllerWtpProfilePlatformDelete,

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
			"wtp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"_local_platform_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddscan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerWtpProfilePlatformUpdate(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectObjectWirelessControllerWtpProfilePlatform(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfilePlatform resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerWtpProfilePlatform(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerWtpProfilePlatform resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWirelessControllerWtpProfilePlatform")

	return resourceObjectWirelessControllerWtpProfilePlatformRead(d, m)
}

func resourceObjectWirelessControllerWtpProfilePlatformDelete(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	paradict["wtp_profile"] = wtp_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerWtpProfilePlatform(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerWtpProfilePlatform resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerWtpProfilePlatformRead(d *schema.ResourceData, m interface{}) error {
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

	wtp_profile := d.Get("wtp_profile").(string)
	if wtp_profile == "" {
		wtp_profile = importOptionChecking(m.(*FortiClient).Cfg, "wtp_profile")
		if wtp_profile == "" {
			return fmt.Errorf("Parameter wtp_profile is missing")
		}
		if err = d.Set("wtp_profile", wtp_profile); err != nil {
			return fmt.Errorf("Error set params wtp_profile: %v", err)
		}
	}
	paradict["wtp_profile"] = wtp_profile

	o, err := c.ReadObjectWirelessControllerWtpProfilePlatform(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfilePlatform resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerWtpProfilePlatform(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerWtpProfilePlatform resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerWtpProfilePlatformLocalPlatformStr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfilePlatformDdscan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfilePlatformMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerWtpProfilePlatformType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerWtpProfilePlatform(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("_local_platform_str", flattenObjectWirelessControllerWtpProfilePlatformLocalPlatformStr2edl(o["_local_platform_str"], d, "_local_platform_str")); err != nil {
		if vv, ok := fortiAPIPatch(o["_local_platform_str"], "ObjectWirelessControllerWtpProfilePlatform-LocalPlatformStr"); ok {
			if err = d.Set("_local_platform_str", vv); err != nil {
				return fmt.Errorf("Error reading _local_platform_str: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _local_platform_str: %v", err)
		}
	}

	if err = d.Set("ddscan", flattenObjectWirelessControllerWtpProfilePlatformDdscan2edl(o["ddscan"], d, "ddscan")); err != nil {
		if vv, ok := fortiAPIPatch(o["ddscan"], "ObjectWirelessControllerWtpProfilePlatform-Ddscan"); ok {
			if err = d.Set("ddscan", vv); err != nil {
				return fmt.Errorf("Error reading ddscan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ddscan: %v", err)
		}
	}

	if err = d.Set("mode", flattenObjectWirelessControllerWtpProfilePlatformMode2edl(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "ObjectWirelessControllerWtpProfilePlatform-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectWirelessControllerWtpProfilePlatformType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectWirelessControllerWtpProfilePlatform-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerWtpProfilePlatformFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerWtpProfilePlatformLocalPlatformStr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfilePlatformDdscan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfilePlatformMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerWtpProfilePlatformType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerWtpProfilePlatform(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_local_platform_str"); ok || d.HasChange("_local_platform_str") {
		t, err := expandObjectWirelessControllerWtpProfilePlatformLocalPlatformStr2edl(d, v, "_local_platform_str")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_local_platform_str"] = t
		}
	}

	if v, ok := d.GetOk("ddscan"); ok || d.HasChange("ddscan") {
		t, err := expandObjectWirelessControllerWtpProfilePlatformDdscan2edl(d, v, "ddscan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddscan"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandObjectWirelessControllerWtpProfilePlatformMode2edl(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectWirelessControllerWtpProfilePlatformType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
