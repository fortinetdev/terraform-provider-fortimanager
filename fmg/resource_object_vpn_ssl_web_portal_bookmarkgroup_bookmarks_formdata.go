// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Form data.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormData() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataCreate,
		Read:   resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataRead,
		Update: resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataUpdate,
		Delete: resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataDelete,

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
			"portal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bookmark_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bookmarks": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	portal := d.Get("portal").(string)
	bookmark_group := d.Get("bookmark_group").(string)
	bookmarks := d.Get("bookmarks").(string)
	paradict["portal"] = portal
	paradict["bookmark_group"] = bookmark_group
	paradict["bookmarks"] = bookmarks

	obj, err := getObjectObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataRead(d, m)
}

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataUpdate(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	bookmark_group := d.Get("bookmark_group").(string)
	bookmarks := d.Get("bookmarks").(string)
	paradict["portal"] = portal
	paradict["bookmark_group"] = bookmark_group
	paradict["bookmarks"] = bookmarks

	obj, err := getObjectObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataRead(d, m)
}

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataDelete(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	bookmark_group := d.Get("bookmark_group").(string)
	bookmarks := d.Get("bookmarks").(string)
	paradict["portal"] = portal
	paradict["bookmark_group"] = bookmark_group
	paradict["bookmarks"] = bookmarks

	err = c.DeleteObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataRead(d *schema.ResourceData, m interface{}) error {
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

	portal := d.Get("portal").(string)
	bookmark_group := d.Get("bookmark_group").(string)
	bookmarks := d.Get("bookmarks").(string)
	if portal == "" {
		portal = importOptionChecking(m.(*FortiClient).Cfg, "portal")
		if portal == "" {
			return fmt.Errorf("Parameter portal is missing")
		}
		if err = d.Set("portal", portal); err != nil {
			return fmt.Errorf("Error set params portal: %v", err)
		}
	}
	if bookmark_group == "" {
		bookmark_group = importOptionChecking(m.(*FortiClient).Cfg, "bookmark_group")
		if bookmark_group == "" {
			return fmt.Errorf("Parameter bookmark_group is missing")
		}
		if err = d.Set("bookmark_group", bookmark_group); err != nil {
			return fmt.Errorf("Error set params bookmark_group: %v", err)
		}
	}
	if bookmarks == "" {
		bookmarks = importOptionChecking(m.(*FortiClient).Cfg, "bookmarks")
		if bookmarks == "" {
			return fmt.Errorf("Parameter bookmarks is missing")
		}
		if err = d.Set("bookmarks", bookmarks); err != nil {
			return fmt.Errorf("Error set params bookmarks: %v", err)
		}
	}
	paradict["portal"] = portal
	paradict["bookmark_group"] = bookmark_group
	paradict["bookmarks"] = bookmarks

	o, err := c.ReadObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("value", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue4thl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ObjectVpnSslWebPortalBookmarkGroupBookmarksFormData-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue4thl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
