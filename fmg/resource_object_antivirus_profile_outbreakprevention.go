// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Virus Outbreak Prevention settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectAntivirusProfileOutbreakPrevention() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectAntivirusProfileOutbreakPreventionUpdate,
		Read:   resourceObjectAntivirusProfileOutbreakPreventionRead,
		Update: resourceObjectAntivirusProfileOutbreakPreventionUpdate,
		Delete: resourceObjectAntivirusProfileOutbreakPreventionDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_blocklist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftgd_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectAntivirusProfileOutbreakPreventionUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectAntivirusProfileOutbreakPrevention(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfileOutbreakPrevention resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectAntivirusProfileOutbreakPrevention(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectAntivirusProfileOutbreakPrevention resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectAntivirusProfileOutbreakPrevention")

	return resourceObjectAntivirusProfileOutbreakPreventionRead(d, m)
}

func resourceObjectAntivirusProfileOutbreakPreventionDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectAntivirusProfileOutbreakPrevention(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectAntivirusProfileOutbreakPrevention resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectAntivirusProfileOutbreakPreventionRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectAntivirusProfileOutbreakPrevention(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAntivirusProfileOutbreakPrevention resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectAntivirusProfileOutbreakPrevention(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectAntivirusProfileOutbreakPrevention resource from API: %v", err)
	}
	return nil
}

func flattenObjectAntivirusProfileOutbreakPreventionExternalBlocklist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectAntivirusProfileOutbreakPreventionFtgdService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectAntivirusProfileOutbreakPrevention(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("external_blocklist", flattenObjectAntivirusProfileOutbreakPreventionExternalBlocklist2edl(o["external-blocklist"], d, "external_blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist"], "ObjectAntivirusProfileOutbreakPrevention-ExternalBlocklist"); ok {
			if err = d.Set("external_blocklist", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist: %v", err)
		}
	}

	if err = d.Set("ftgd_service", flattenObjectAntivirusProfileOutbreakPreventionFtgdService2edl(o["ftgd-service"], d, "ftgd_service")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftgd-service"], "ObjectAntivirusProfileOutbreakPrevention-FtgdService"); ok {
			if err = d.Set("ftgd_service", vv); err != nil {
				return fmt.Errorf("Error reading ftgd_service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftgd_service: %v", err)
		}
	}

	return nil
}

func flattenObjectAntivirusProfileOutbreakPreventionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectAntivirusProfileOutbreakPreventionExternalBlocklist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectAntivirusProfileOutbreakPreventionFtgdService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectAntivirusProfileOutbreakPrevention(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("external_blocklist"); ok || d.HasChange("external_blocklist") {
		t, err := expandObjectAntivirusProfileOutbreakPreventionExternalBlocklist2edl(d, v, "external_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_service"); ok || d.HasChange("ftgd_service") {
		t, err := expandObjectAntivirusProfileOutbreakPreventionFtgdService2edl(d, v, "ftgd_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-service"] = t
		}
	}

	return &obj, nil
}
