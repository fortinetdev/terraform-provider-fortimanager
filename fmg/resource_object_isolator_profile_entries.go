// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectIsolator ProfileEntries

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectIsolatorProfileEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIsolatorProfileEntriesCreate,
		Read:   resourceObjectIsolatorProfileEntriesRead,
		Update: resourceObjectIsolatorProfileEntriesUpdate,
		Delete: resourceObjectIsolatorProfileEntriesDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"copy_paste": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"proxy_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"right_click": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectIsolatorProfileEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectIsolatorProfileEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIsolatorProfileEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectIsolatorProfileEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIsolatorProfileEntries resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIsolatorProfileEntriesRead(d, m)
}

func resourceObjectIsolatorProfileEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectIsolatorProfileEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIsolatorProfileEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectIsolatorProfileEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIsolatorProfileEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIsolatorProfileEntriesRead(d, m)
}

func resourceObjectIsolatorProfileEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectIsolatorProfileEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIsolatorProfileEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIsolatorProfileEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectIsolatorProfileEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIsolatorProfileEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIsolatorProfileEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIsolatorProfileEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectIsolatorProfileEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIsolatorProfileEntriesCopyPaste2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIsolatorProfileEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIsolatorProfileEntriesProxyAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectIsolatorProfileEntriesRightClick2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIsolatorProfileEntriesStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectIsolatorProfileEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectIsolatorProfileEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectIsolatorProfileEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("copy_paste", flattenObjectIsolatorProfileEntriesCopyPaste2edl(o["copy-paste"], d, "copy_paste")); err != nil {
		if vv, ok := fortiAPIPatch(o["copy-paste"], "ObjectIsolatorProfileEntries-CopyPaste"); ok {
			if err = d.Set("copy_paste", vv); err != nil {
				return fmt.Errorf("Error reading copy_paste: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading copy_paste: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectIsolatorProfileEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectIsolatorProfileEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("proxy_address", flattenObjectIsolatorProfileEntriesProxyAddress2edl(o["proxy-address"], d, "proxy_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["proxy-address"], "ObjectIsolatorProfileEntries-ProxyAddress"); ok {
			if err = d.Set("proxy_address", vv); err != nil {
				return fmt.Errorf("Error reading proxy_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading proxy_address: %v", err)
		}
	}

	if err = d.Set("right_click", flattenObjectIsolatorProfileEntriesRightClick2edl(o["right-click"], d, "right_click")); err != nil {
		if vv, ok := fortiAPIPatch(o["right-click"], "ObjectIsolatorProfileEntries-RightClick"); ok {
			if err = d.Set("right_click", vv); err != nil {
				return fmt.Errorf("Error reading right_click: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading right_click: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectIsolatorProfileEntriesStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectIsolatorProfileEntries-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectIsolatorProfileEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIsolatorProfileEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIsolatorProfileEntriesCopyPaste2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIsolatorProfileEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIsolatorProfileEntriesProxyAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectIsolatorProfileEntriesRightClick2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIsolatorProfileEntriesStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIsolatorProfileEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectIsolatorProfileEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("copy_paste"); ok || d.HasChange("copy_paste") {
		t, err := expandObjectIsolatorProfileEntriesCopyPaste2edl(d, v, "copy_paste")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["copy-paste"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectIsolatorProfileEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("proxy_address"); ok || d.HasChange("proxy_address") {
		t, err := expandObjectIsolatorProfileEntriesProxyAddress2edl(d, v, "proxy_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-address"] = t
		}
	}

	if v, ok := d.GetOk("right_click"); ok || d.HasChange("right_click") {
		t, err := expandObjectIsolatorProfileEntriesRightClick2edl(d, v, "right_click")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["right-click"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectIsolatorProfileEntriesStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
