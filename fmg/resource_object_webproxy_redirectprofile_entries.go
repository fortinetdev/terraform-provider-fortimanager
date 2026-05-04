// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectWebProxy RedirectProfileEntries

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebProxyRedirectProfileEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebProxyRedirectProfileEntriesCreate,
		Read:   resourceObjectWebProxyRedirectProfileEntriesRead,
		Update: resourceObjectWebProxyRedirectProfileEntriesUpdate,
		Delete: resourceObjectWebProxyRedirectProfileEntriesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"redirect_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"redirect_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWebProxyRedirectProfileEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	redirect_profile := d.Get("redirect_profile").(string)
	paradict["redirect_profile"] = redirect_profile

	obj, err := getObjectObjectWebProxyRedirectProfileEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyRedirectProfileEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectWebProxyRedirectProfileEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectWebProxyRedirectProfileEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectWebProxyRedirectProfileEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectWebProxyRedirectProfileEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectWebProxyRedirectProfileEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebProxyRedirectProfileEntriesRead(d, m)
}

func resourceObjectWebProxyRedirectProfileEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	redirect_profile := d.Get("redirect_profile").(string)
	paradict["redirect_profile"] = redirect_profile

	obj, err := getObjectObjectWebProxyRedirectProfileEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyRedirectProfileEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebProxyRedirectProfileEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyRedirectProfileEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWebProxyRedirectProfileEntriesRead(d, m)
}

func resourceObjectWebProxyRedirectProfileEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	redirect_profile := d.Get("redirect_profile").(string)
	paradict["redirect_profile"] = redirect_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectWebProxyRedirectProfileEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebProxyRedirectProfileEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebProxyRedirectProfileEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	redirect_profile := d.Get("redirect_profile").(string)
	if redirect_profile == "" {
		redirect_profile = importOptionChecking(m.(*FortiClient).Cfg, "redirect_profile")
		if redirect_profile == "" {
			return fmt.Errorf("Parameter redirect_profile is missing")
		}
		if err = d.Set("redirect_profile", redirect_profile); err != nil {
			return fmt.Errorf("Error set params redirect_profile: %v", err)
		}
	}
	paradict["redirect_profile"] = redirect_profile

	o, err := c.ReadObjectWebProxyRedirectProfileEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectWebProxyRedirectProfileEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebProxyRedirectProfileEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyRedirectProfileEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebProxyRedirectProfileEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyRedirectProfileEntriesRedirectCode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyRedirectProfileEntriesRedirectUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyRedirectProfileEntriesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyRedirectProfileEntriesUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebProxyRedirectProfileEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectWebProxyRedirectProfileEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWebProxyRedirectProfileEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("redirect_code", flattenObjectWebProxyRedirectProfileEntriesRedirectCode2edl(o["redirect-code"], d, "redirect_code")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-code"], "ObjectWebProxyRedirectProfileEntries-RedirectCode"); ok {
			if err = d.Set("redirect_code", vv); err != nil {
				return fmt.Errorf("Error reading redirect_code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_code: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenObjectWebProxyRedirectProfileEntriesRedirectUrl2edl(o["redirect-url"], d, "redirect_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-url"], "ObjectWebProxyRedirectProfileEntries-RedirectUrl"); ok {
			if err = d.Set("redirect_url", vv); err != nil {
				return fmt.Errorf("Error reading redirect_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectWebProxyRedirectProfileEntriesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectWebProxyRedirectProfileEntries-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("url", flattenObjectWebProxyRedirectProfileEntriesUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ObjectWebProxyRedirectProfileEntries-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenObjectWebProxyRedirectProfileEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebProxyRedirectProfileEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyRedirectProfileEntriesRedirectCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyRedirectProfileEntriesRedirectUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyRedirectProfileEntriesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyRedirectProfileEntriesUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebProxyRedirectProfileEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWebProxyRedirectProfileEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("redirect_code"); ok || d.HasChange("redirect_code") {
		t, err := expandObjectWebProxyRedirectProfileEntriesRedirectCode2edl(d, v, "redirect_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-code"] = t
		}
	}

	if v, ok := d.GetOk("redirect_url"); ok || d.HasChange("redirect_url") {
		t, err := expandObjectWebProxyRedirectProfileEntriesRedirectUrl2edl(d, v, "redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectWebProxyRedirectProfileEntriesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandObjectWebProxyRedirectProfileEntriesUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
