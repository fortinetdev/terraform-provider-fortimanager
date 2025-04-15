// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure URL Extraction

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileUrlExtraction() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileUrlExtractionUpdate,
		Read:   resourceObjectWebfilterProfileUrlExtractionRead,
		Update: resourceObjectWebfilterProfileUrlExtractionUpdate,
		Delete: resourceObjectWebfilterProfileUrlExtractionDelete,

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
			"redirect_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect_no_content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWebfilterProfileUrlExtractionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterProfileUrlExtraction(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileUrlExtraction resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterProfileUrlExtraction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileUrlExtraction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWebfilterProfileUrlExtraction")

	return resourceObjectWebfilterProfileUrlExtractionRead(d, m)
}

func resourceObjectWebfilterProfileUrlExtractionDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWebfilterProfileUrlExtraction(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileUrlExtraction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileUrlExtractionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebfilterProfileUrlExtraction(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileUrlExtraction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileUrlExtraction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileUrlExtraction resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileUrlExtractionRedirectHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtractionRedirectNoContent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtractionRedirectUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtractionServerFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileUrlExtractionStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebfilterProfileUrlExtraction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("redirect_header", flattenObjectWebfilterProfileUrlExtractionRedirectHeader2edl(o["redirect-header"], d, "redirect_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-header"], "ObjectWebfilterProfileUrlExtraction-RedirectHeader"); ok {
			if err = d.Set("redirect_header", vv); err != nil {
				return fmt.Errorf("Error reading redirect_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_header: %v", err)
		}
	}

	if err = d.Set("redirect_no_content", flattenObjectWebfilterProfileUrlExtractionRedirectNoContent2edl(o["redirect-no-content"], d, "redirect_no_content")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-no-content"], "ObjectWebfilterProfileUrlExtraction-RedirectNoContent"); ok {
			if err = d.Set("redirect_no_content", vv); err != nil {
				return fmt.Errorf("Error reading redirect_no_content: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_no_content: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenObjectWebfilterProfileUrlExtractionRedirectUrl2edl(o["redirect-url"], d, "redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-url"], "ObjectWebfilterProfileUrlExtraction-RedirectUrl"); ok {
			if err = d.Set("redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("server_fqdn", flattenObjectWebfilterProfileUrlExtractionServerFqdn2edl(o["server-fqdn"], d, "server_fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-fqdn"], "ObjectWebfilterProfileUrlExtraction-ServerFqdn"); ok {
			if err = d.Set("server_fqdn", vv); err != nil {
				return fmt.Errorf("Error reading server_fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_fqdn: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectWebfilterProfileUrlExtractionStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectWebfilterProfileUrlExtraction-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileUrlExtractionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileUrlExtractionRedirectHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtractionRedirectNoContent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtractionRedirectUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtractionServerFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileUrlExtractionStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebfilterProfileUrlExtraction(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("redirect_header"); ok || d.HasChange("redirect_header") {
		t, err := expandObjectWebfilterProfileUrlExtractionRedirectHeader2edl(d, v, "redirect_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-header"] = t
		}
	}

	if v, ok := d.GetOk("redirect_no_content"); ok || d.HasChange("redirect_no_content") {
		t, err := expandObjectWebfilterProfileUrlExtractionRedirectNoContent2edl(d, v, "redirect_no_content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-no-content"] = t
		}
	}

	if v, ok := d.GetOk("redirect_url"); ok || d.HasChange("redirect_url") {
		t, err := expandObjectWebfilterProfileUrlExtractionRedirectUrl2edl(d, v, "redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("server_fqdn"); ok || d.HasChange("server_fqdn") {
		t, err := expandObjectWebfilterProfileUrlExtractionServerFqdn2edl(d, v, "server_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectWebfilterProfileUrlExtractionStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
