// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Bookmark table.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectZtnaWebPortalBookmarkBookmarks() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectZtnaWebPortalBookmarkBookmarksCreate,
		Read:   resourceObjectZtnaWebPortalBookmarkBookmarksRead,
		Update: resourceObjectZtnaWebPortalBookmarkBookmarksUpdate,
		Delete: resourceObjectZtnaWebPortalBookmarkBookmarksDelete,

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
			"web_portal_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apptype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"color_depth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"folder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keyboard_layout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"load_balancing_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logon_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"logon_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"preconnection_blob": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"preconnection_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"restricted_admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_preconnection_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnc_keyboard_layout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectZtnaWebPortalBookmarkBookmarksCreate(d *schema.ResourceData, m interface{}) error {
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

	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["web_portal_bookmark"] = web_portal_bookmark

	obj, err := getObjectObjectZtnaWebPortalBookmarkBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectZtnaWebPortalBookmarkBookmarks resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectZtnaWebPortalBookmarkBookmarks(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectZtnaWebPortalBookmarkBookmarks(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectZtnaWebPortalBookmarkBookmarks resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectZtnaWebPortalBookmarkBookmarks(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectZtnaWebPortalBookmarkBookmarks resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaWebPortalBookmarkBookmarksRead(d, m)
}

func resourceObjectZtnaWebPortalBookmarkBookmarksUpdate(d *schema.ResourceData, m interface{}) error {
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

	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["web_portal_bookmark"] = web_portal_bookmark

	obj, err := getObjectObjectZtnaWebPortalBookmarkBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaWebPortalBookmarkBookmarks resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectZtnaWebPortalBookmarkBookmarks(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectZtnaWebPortalBookmarkBookmarks resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectZtnaWebPortalBookmarkBookmarksRead(d, m)
}

func resourceObjectZtnaWebPortalBookmarkBookmarksDelete(d *schema.ResourceData, m interface{}) error {
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

	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	paradict["web_portal_bookmark"] = web_portal_bookmark

	wsParams["adom"] = adomv

	err = c.DeleteObjectZtnaWebPortalBookmarkBookmarks(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectZtnaWebPortalBookmarkBookmarks resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectZtnaWebPortalBookmarkBookmarksRead(d *schema.ResourceData, m interface{}) error {
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

	web_portal_bookmark := d.Get("web_portal_bookmark").(string)
	if web_portal_bookmark == "" {
		web_portal_bookmark = importOptionChecking(m.(*FortiClient).Cfg, "web_portal_bookmark")
		if web_portal_bookmark == "" {
			return fmt.Errorf("Parameter web_portal_bookmark is missing")
		}
		if err = d.Set("web_portal_bookmark", web_portal_bookmark); err != nil {
			return fmt.Errorf("Error set params web_portal_bookmark: %v", err)
		}
	}
	paradict["web_portal_bookmark"] = web_portal_bookmark

	o, err := c.ReadObjectZtnaWebPortalBookmarkBookmarks(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectZtnaWebPortalBookmarkBookmarks resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectZtnaWebPortalBookmarkBookmarks(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectZtnaWebPortalBookmarkBookmarks resource from API: %v", err)
	}
	return nil
}

func flattenObjectZtnaWebPortalBookmarkBookmarksApptype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksColorDepth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksFolder2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksHeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksLoadBalancingInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksLogonUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksPreconnectionBlob2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksRestrictedAdmin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksSecurity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksSendPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksSso2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksVncKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectZtnaWebPortalBookmarkBookmarksWidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectZtnaWebPortalBookmarkBookmarks(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("apptype", flattenObjectZtnaWebPortalBookmarkBookmarksApptype2edl(o["apptype"], d, "apptype")); err != nil {
		if vv, ok := fortiAPIPatch(o["apptype"], "ObjectZtnaWebPortalBookmarkBookmarks-Apptype"); ok {
			if err = d.Set("apptype", vv); err != nil {
				return fmt.Errorf("Error reading apptype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apptype: %v", err)
		}
	}

	if err = d.Set("color_depth", flattenObjectZtnaWebPortalBookmarkBookmarksColorDepth2edl(o["color-depth"], d, "color_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["color-depth"], "ObjectZtnaWebPortalBookmarkBookmarks-ColorDepth"); ok {
			if err = d.Set("color_depth", vv); err != nil {
				return fmt.Errorf("Error reading color_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color_depth: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectZtnaWebPortalBookmarkBookmarksDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectZtnaWebPortalBookmarkBookmarks-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectZtnaWebPortalBookmarkBookmarksDomain2edl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectZtnaWebPortalBookmarkBookmarks-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("folder", flattenObjectZtnaWebPortalBookmarkBookmarksFolder2edl(o["folder"], d, "folder")); err != nil {
		if vv, ok := fortiAPIPatch(o["folder"], "ObjectZtnaWebPortalBookmarkBookmarks-Folder"); ok {
			if err = d.Set("folder", vv); err != nil {
				return fmt.Errorf("Error reading folder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if err = d.Set("height", flattenObjectZtnaWebPortalBookmarkBookmarksHeight2edl(o["height"], d, "height")); err != nil {
		if vv, ok := fortiAPIPatch(o["height"], "ObjectZtnaWebPortalBookmarkBookmarks-Height"); ok {
			if err = d.Set("height", vv); err != nil {
				return fmt.Errorf("Error reading height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading height: %v", err)
		}
	}

	if err = d.Set("host", flattenObjectZtnaWebPortalBookmarkBookmarksHost2edl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ObjectZtnaWebPortalBookmarkBookmarks-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("keyboard_layout", flattenObjectZtnaWebPortalBookmarkBookmarksKeyboardLayout2edl(o["keyboard-layout"], d, "keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyboard-layout"], "ObjectZtnaWebPortalBookmarkBookmarks-KeyboardLayout"); ok {
			if err = d.Set("keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyboard_layout: %v", err)
		}
	}

	if err = d.Set("load_balancing_info", flattenObjectZtnaWebPortalBookmarkBookmarksLoadBalancingInfo2edl(o["load-balancing-info"], d, "load_balancing_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balancing-info"], "ObjectZtnaWebPortalBookmarkBookmarks-LoadBalancingInfo"); ok {
			if err = d.Set("load_balancing_info", vv); err != nil {
				return fmt.Errorf("Error reading load_balancing_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balancing_info: %v", err)
		}
	}

	if err = d.Set("logon_user", flattenObjectZtnaWebPortalBookmarkBookmarksLogonUser2edl(o["logon-user"], d, "logon_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-user"], "ObjectZtnaWebPortalBookmarkBookmarks-LogonUser"); ok {
			if err = d.Set("logon_user", vv); err != nil {
				return fmt.Errorf("Error reading logon_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_user: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectZtnaWebPortalBookmarkBookmarksName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectZtnaWebPortalBookmarkBookmarks-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectZtnaWebPortalBookmarkBookmarksPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectZtnaWebPortalBookmarkBookmarks-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("preconnection_blob", flattenObjectZtnaWebPortalBookmarkBookmarksPreconnectionBlob2edl(o["preconnection-blob"], d, "preconnection_blob")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-blob"], "ObjectZtnaWebPortalBookmarkBookmarks-PreconnectionBlob"); ok {
			if err = d.Set("preconnection_blob", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_blob: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_blob: %v", err)
		}
	}

	if err = d.Set("preconnection_id", flattenObjectZtnaWebPortalBookmarkBookmarksPreconnectionId2edl(o["preconnection-id"], d, "preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-id"], "ObjectZtnaWebPortalBookmarkBookmarks-PreconnectionId"); ok {
			if err = d.Set("preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_id: %v", err)
		}
	}

	if err = d.Set("restricted_admin", flattenObjectZtnaWebPortalBookmarkBookmarksRestrictedAdmin2edl(o["restricted-admin"], d, "restricted_admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["restricted-admin"], "ObjectZtnaWebPortalBookmarkBookmarks-RestrictedAdmin"); ok {
			if err = d.Set("restricted_admin", vv); err != nil {
				return fmt.Errorf("Error reading restricted_admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restricted_admin: %v", err)
		}
	}

	if err = d.Set("security", flattenObjectZtnaWebPortalBookmarkBookmarksSecurity2edl(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "ObjectZtnaWebPortalBookmarkBookmarks-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("send_preconnection_id", flattenObjectZtnaWebPortalBookmarkBookmarksSendPreconnectionId2edl(o["send-preconnection-id"], d, "send_preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-preconnection-id"], "ObjectZtnaWebPortalBookmarkBookmarks-SendPreconnectionId"); ok {
			if err = d.Set("send_preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading send_preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_preconnection_id: %v", err)
		}
	}

	if err = d.Set("sso", flattenObjectZtnaWebPortalBookmarkBookmarksSso2edl(o["sso"], d, "sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso"], "ObjectZtnaWebPortalBookmarkBookmarks-Sso"); ok {
			if err = d.Set("sso", vv); err != nil {
				return fmt.Errorf("Error reading sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso: %v", err)
		}
	}

	if err = d.Set("url", flattenObjectZtnaWebPortalBookmarkBookmarksUrl2edl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ObjectZtnaWebPortalBookmarkBookmarks-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("vnc_keyboard_layout", flattenObjectZtnaWebPortalBookmarkBookmarksVncKeyboardLayout2edl(o["vnc-keyboard-layout"], d, "vnc_keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["vnc-keyboard-layout"], "ObjectZtnaWebPortalBookmarkBookmarks-VncKeyboardLayout"); ok {
			if err = d.Set("vnc_keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
		}
	}

	if err = d.Set("width", flattenObjectZtnaWebPortalBookmarkBookmarksWidth2edl(o["width"], d, "width")); err != nil {
		if vv, ok := fortiAPIPatch(o["width"], "ObjectZtnaWebPortalBookmarkBookmarks-Width"); ok {
			if err = d.Set("width", vv); err != nil {
				return fmt.Errorf("Error reading width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading width: %v", err)
		}
	}

	return nil
}

func flattenObjectZtnaWebPortalBookmarkBookmarksFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectZtnaWebPortalBookmarkBookmarksApptype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksColorDepth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksFolder2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksHeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksLoadBalancingInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksLogonPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksLogonUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksPreconnectionBlob2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksRestrictedAdmin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksSecurity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksSendPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksSso2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksVncKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectZtnaWebPortalBookmarkBookmarksWidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectZtnaWebPortalBookmarkBookmarks(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("apptype"); ok || d.HasChange("apptype") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksApptype2edl(d, v, "apptype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apptype"] = t
		}
	}

	if v, ok := d.GetOk("color_depth"); ok || d.HasChange("color_depth") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksColorDepth2edl(d, v, "color_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color-depth"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksDomain2edl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("folder"); ok || d.HasChange("folder") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksFolder2edl(d, v, "folder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folder"] = t
		}
	}

	if v, ok := d.GetOk("height"); ok || d.HasChange("height") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksHeight2edl(d, v, "height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["height"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksHost2edl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("keyboard_layout"); ok || d.HasChange("keyboard_layout") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksKeyboardLayout2edl(d, v, "keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("load_balancing_info"); ok || d.HasChange("load_balancing_info") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksLoadBalancingInfo2edl(d, v, "load_balancing_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balancing-info"] = t
		}
	}

	if v, ok := d.GetOk("logon_password"); ok || d.HasChange("logon_password") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksLogonPassword2edl(d, v, "logon_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-password"] = t
		}
	}

	if v, ok := d.GetOk("logon_user"); ok || d.HasChange("logon_user") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksLogonUser2edl(d, v, "logon_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-user"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_blob"); ok || d.HasChange("preconnection_blob") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksPreconnectionBlob2edl(d, v, "preconnection_blob")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-blob"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_id"); ok || d.HasChange("preconnection_id") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksPreconnectionId2edl(d, v, "preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("restricted_admin"); ok || d.HasChange("restricted_admin") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksRestrictedAdmin2edl(d, v, "restricted_admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restricted-admin"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksSecurity2edl(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("send_preconnection_id"); ok || d.HasChange("send_preconnection_id") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksSendPreconnectionId2edl(d, v, "send_preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("sso"); ok || d.HasChange("sso") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksSso2edl(d, v, "sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksUrl2edl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("vnc_keyboard_layout"); ok || d.HasChange("vnc_keyboard_layout") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksVncKeyboardLayout2edl(d, v, "vnc_keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vnc-keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("width"); ok || d.HasChange("width") {
		t, err := expandObjectZtnaWebPortalBookmarkBookmarksWidth2edl(d, v, "width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["width"] = t
		}
	}

	return &obj, nil
}
