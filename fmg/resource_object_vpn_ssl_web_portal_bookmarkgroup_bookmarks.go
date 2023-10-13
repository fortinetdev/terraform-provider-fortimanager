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

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarks() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebPortalBookmarkGroupBookmarksCreate,
		Read:   resourceObjectVpnSslWebPortalBookmarkGroupBookmarksRead,
		Update: resourceObjectVpnSslWebPortalBookmarkGroupBookmarksUpdate,
		Delete: resourceObjectVpnSslWebPortalBookmarkGroupBookmarksDelete,

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
			"additional_params": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apptype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"color_depth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"form_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			},
			"listening_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"load_balancing_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logon_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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
			"remote_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"restricted_admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"send_preconnection_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_layout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"show_status_window": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_credential": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_credential_sent_once": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sso_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnc_keyboard_layout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["portal"] = portal
	paradict["bookmark_group"] = bookmark_group

	obj, err := getObjectObjectVpnSslWebPortalBookmarkGroupBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalBookmarkGroupBookmarks resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnSslWebPortalBookmarkGroupBookmarks(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalBookmarkGroupBookmarks resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalBookmarkGroupBookmarksRead(d, m)
}

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["portal"] = portal
	paradict["bookmark_group"] = bookmark_group

	obj, err := getObjectObjectVpnSslWebPortalBookmarkGroupBookmarks(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalBookmarkGroupBookmarks resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebPortalBookmarkGroupBookmarks(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalBookmarkGroupBookmarks resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalBookmarkGroupBookmarksRead(d, m)
}

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["portal"] = portal
	paradict["bookmark_group"] = bookmark_group

	err = c.DeleteObjectVpnSslWebPortalBookmarkGroupBookmarks(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebPortalBookmarkGroupBookmarks resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebPortalBookmarkGroupBookmarksRead(d *schema.ResourceData, m interface{}) error {
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
	paradict["portal"] = portal
	paradict["bookmark_group"] = bookmark_group

	o, err := c.ReadObjectVpnSslWebPortalBookmarkGroupBookmarks(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalBookmarkGroupBookmarks resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebPortalBookmarkGroupBookmarks(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalBookmarkGroupBookmarks resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksApptype3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDescription3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDomain3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFolder3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormData3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName3rdl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroupBookmarks-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue3rdl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroupBookmarks-FormData-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSso3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksUrl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksWidth3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebPortalBookmarkGroupBookmarks(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("additional_params", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams3rdl(o["additional-params"], d, "additional_params")); err != nil {
		if vv, ok := fortiAPIPatch(o["additional-params"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-AdditionalParams"); ok {
			if err = d.Set("additional_params", vv); err != nil {
				return fmt.Errorf("Error reading additional_params: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading additional_params: %v", err)
		}
	}

	if err = d.Set("apptype", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksApptype3rdl(o["apptype"], d, "apptype")); err != nil {
		if vv, ok := fortiAPIPatch(o["apptype"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Apptype"); ok {
			if err = d.Set("apptype", vv); err != nil {
				return fmt.Errorf("Error reading apptype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading apptype: %v", err)
		}
	}

	if err = d.Set("color_depth", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth3rdl(o["color-depth"], d, "color_depth")); err != nil {
		if vv, ok := fortiAPIPatch(o["color-depth"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-ColorDepth"); ok {
			if err = d.Set("color_depth", vv); err != nil {
				return fmt.Errorf("Error reading color_depth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color_depth: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDescription3rdl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("domain", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDomain3rdl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("folder", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFolder3rdl(o["folder"], d, "folder")); err != nil {
		if vv, ok := fortiAPIPatch(o["folder"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Folder"); ok {
			if err = d.Set("folder", vv); err != nil {
				return fmt.Errorf("Error reading folder: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading folder: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("form_data", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormData3rdl(o["form-data"], d, "form_data")); err != nil {
			if vv, ok := fortiAPIPatch(o["form-data"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-FormData"); ok {
				if err = d.Set("form_data", vv); err != nil {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading form_data: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("form_data"); ok {
			if err = d.Set("form_data", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormData3rdl(o["form-data"], d, "form_data")); err != nil {
				if vv, ok := fortiAPIPatch(o["form-data"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-FormData"); ok {
					if err = d.Set("form_data", vv); err != nil {
						return fmt.Errorf("Error reading form_data: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			}
		}
	}

	if err = d.Set("height", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHeight3rdl(o["height"], d, "height")); err != nil {
		if vv, ok := fortiAPIPatch(o["height"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Height"); ok {
			if err = d.Set("height", vv); err != nil {
				return fmt.Errorf("Error reading height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading height: %v", err)
		}
	}

	if err = d.Set("host", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHost3rdl(o["host"], d, "host")); err != nil {
		if vv, ok := fortiAPIPatch(o["host"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Host"); ok {
			if err = d.Set("host", vv); err != nil {
				return fmt.Errorf("Error reading host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("keyboard_layout", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout3rdl(o["keyboard-layout"], d, "keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["keyboard-layout"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-KeyboardLayout"); ok {
			if err = d.Set("keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keyboard_layout: %v", err)
		}
	}

	if err = d.Set("listening_port", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort3rdl(o["listening-port"], d, "listening_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["listening-port"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-ListeningPort"); ok {
			if err = d.Set("listening_port", vv); err != nil {
				return fmt.Errorf("Error reading listening_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading listening_port: %v", err)
		}
	}

	if err = d.Set("load_balancing_info", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo3rdl(o["load-balancing-info"], d, "load_balancing_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balancing-info"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-LoadBalancingInfo"); ok {
			if err = d.Set("load_balancing_info", vv); err != nil {
				return fmt.Errorf("Error reading load_balancing_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balancing_info: %v", err)
		}
	}

	if err = d.Set("logon_password", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword3rdl(o["logon-password"], d, "logon_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-password"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-LogonPassword"); ok {
			if err = d.Set("logon_password", vv); err != nil {
				return fmt.Errorf("Error reading logon_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_password: %v", err)
		}
	}

	if err = d.Set("logon_user", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser3rdl(o["logon-user"], d, "logon_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["logon-user"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-LogonUser"); ok {
			if err = d.Set("logon_user", vv); err != nil {
				return fmt.Errorf("Error reading logon_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logon_user: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("preconnection_blob", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob3rdl(o["preconnection-blob"], d, "preconnection_blob")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-blob"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-PreconnectionBlob"); ok {
			if err = d.Set("preconnection_blob", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_blob: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_blob: %v", err)
		}
	}

	if err = d.Set("preconnection_id", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId3rdl(o["preconnection-id"], d, "preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["preconnection-id"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-PreconnectionId"); ok {
			if err = d.Set("preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preconnection_id: %v", err)
		}
	}

	if err = d.Set("remote_port", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort3rdl(o["remote-port"], d, "remote_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-port"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-RemotePort"); ok {
			if err = d.Set("remote_port", vv); err != nil {
				return fmt.Errorf("Error reading remote_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_port: %v", err)
		}
	}

	if err = d.Set("restricted_admin", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin3rdl(o["restricted-admin"], d, "restricted_admin")); err != nil {
		if vv, ok := fortiAPIPatch(o["restricted-admin"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-RestrictedAdmin"); ok {
			if err = d.Set("restricted_admin", vv); err != nil {
				return fmt.Errorf("Error reading restricted_admin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restricted_admin: %v", err)
		}
	}

	if err = d.Set("security", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity3rdl(o["security"], d, "security")); err != nil {
		if vv, ok := fortiAPIPatch(o["security"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Security"); ok {
			if err = d.Set("security", vv); err != nil {
				return fmt.Errorf("Error reading security: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("send_preconnection_id", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId3rdl(o["send-preconnection-id"], d, "send_preconnection_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-preconnection-id"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-SendPreconnectionId"); ok {
			if err = d.Set("send_preconnection_id", vv); err != nil {
				return fmt.Errorf("Error reading send_preconnection_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_preconnection_id: %v", err)
		}
	}

	if err = d.Set("server_layout", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout3rdl(o["server-layout"], d, "server_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-layout"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-ServerLayout"); ok {
			if err = d.Set("server_layout", vv); err != nil {
				return fmt.Errorf("Error reading server_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_layout: %v", err)
		}
	}

	if err = d.Set("show_status_window", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow3rdl(o["show-status-window"], d, "show_status_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["show-status-window"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-ShowStatusWindow"); ok {
			if err = d.Set("show_status_window", vv); err != nil {
				return fmt.Errorf("Error reading show_status_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading show_status_window: %v", err)
		}
	}

	if err = d.Set("sso", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSso3rdl(o["sso"], d, "sso")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Sso"); ok {
			if err = d.Set("sso", vv); err != nil {
				return fmt.Errorf("Error reading sso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso: %v", err)
		}
	}

	if err = d.Set("sso_credential", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential3rdl(o["sso-credential"], d, "sso_credential")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-credential"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-SsoCredential"); ok {
			if err = d.Set("sso_credential", vv); err != nil {
				return fmt.Errorf("Error reading sso_credential: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_credential: %v", err)
		}
	}

	if err = d.Set("sso_credential_sent_once", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce3rdl(o["sso-credential-sent-once"], d, "sso_credential_sent_once")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-credential-sent-once"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-SsoCredentialSentOnce"); ok {
			if err = d.Set("sso_credential_sent_once", vv); err != nil {
				return fmt.Errorf("Error reading sso_credential_sent_once: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_credential_sent_once: %v", err)
		}
	}

	if err = d.Set("sso_password", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword3rdl(o["sso-password"], d, "sso_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-password"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-SsoPassword"); ok {
			if err = d.Set("sso_password", vv); err != nil {
				return fmt.Errorf("Error reading sso_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_password: %v", err)
		}
	}

	if err = d.Set("sso_username", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername3rdl(o["sso-username"], d, "sso_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-username"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-SsoUsername"); ok {
			if err = d.Set("sso_username", vv); err != nil {
				return fmt.Errorf("Error reading sso_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_username: %v", err)
		}
	}

	if err = d.Set("url", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksUrl3rdl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("vnc_keyboard_layout", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout3rdl(o["vnc-keyboard-layout"], d, "vnc_keyboard_layout")); err != nil {
		if vv, ok := fortiAPIPatch(o["vnc-keyboard-layout"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-VncKeyboardLayout"); ok {
			if err = d.Set("vnc_keyboard_layout", vv); err != nil {
				return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vnc_keyboard_layout: %v", err)
		}
	}

	if err = d.Set("width", flattenObjectVpnSslWebPortalBookmarkGroupBookmarksWidth3rdl(o["width"], d, "width")); err != nil {
		if vv, ok := fortiAPIPatch(o["width"], "ObjectVpnSslWebPortalBookmarkGroupBookmarks-Width"); ok {
			if err = d.Set("width", vv); err != nil {
				return fmt.Errorf("Error reading width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading width: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksApptype3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksDescription3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksDomain3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFolder3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormData3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName3rdl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue3rdl(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksHeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSso3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksUrl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksWidth3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebPortalBookmarkGroupBookmarks(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("additional_params"); ok || d.HasChange("additional_params") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams3rdl(d, v, "additional_params")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-params"] = t
		}
	}

	if v, ok := d.GetOk("apptype"); ok || d.HasChange("apptype") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksApptype3rdl(d, v, "apptype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apptype"] = t
		}
	}

	if v, ok := d.GetOk("color_depth"); ok || d.HasChange("color_depth") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth3rdl(d, v, "color_depth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color-depth"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksDescription3rdl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksDomain3rdl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("folder"); ok || d.HasChange("folder") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksFolder3rdl(d, v, "folder")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folder"] = t
		}
	}

	if v, ok := d.GetOk("form_data"); ok || d.HasChange("form_data") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormData3rdl(d, v, "form_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["form-data"] = t
		}
	}

	if v, ok := d.GetOk("height"); ok || d.HasChange("height") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksHeight3rdl(d, v, "height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["height"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok || d.HasChange("host") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksHost3rdl(d, v, "host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("keyboard_layout"); ok || d.HasChange("keyboard_layout") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout3rdl(d, v, "keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("listening_port"); ok || d.HasChange("listening_port") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort3rdl(d, v, "listening_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["listening-port"] = t
		}
	}

	if v, ok := d.GetOk("load_balancing_info"); ok || d.HasChange("load_balancing_info") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo3rdl(d, v, "load_balancing_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balancing-info"] = t
		}
	}

	if v, ok := d.GetOk("logon_password"); ok || d.HasChange("logon_password") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword3rdl(d, v, "logon_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-password"] = t
		}
	}

	if v, ok := d.GetOk("logon_user"); ok || d.HasChange("logon_user") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser3rdl(d, v, "logon_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-user"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_blob"); ok || d.HasChange("preconnection_blob") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob3rdl(d, v, "preconnection_blob")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-blob"] = t
		}
	}

	if v, ok := d.GetOk("preconnection_id"); ok || d.HasChange("preconnection_id") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId3rdl(d, v, "preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_port"); ok || d.HasChange("remote_port") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort3rdl(d, v, "remote_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-port"] = t
		}
	}

	if v, ok := d.GetOk("restricted_admin"); ok || d.HasChange("restricted_admin") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin3rdl(d, v, "restricted_admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restricted-admin"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok || d.HasChange("security") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity3rdl(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("send_preconnection_id"); ok || d.HasChange("send_preconnection_id") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId3rdl(d, v, "send_preconnection_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-preconnection-id"] = t
		}
	}

	if v, ok := d.GetOk("server_layout"); ok || d.HasChange("server_layout") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout3rdl(d, v, "server_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-layout"] = t
		}
	}

	if v, ok := d.GetOk("show_status_window"); ok || d.HasChange("show_status_window") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow3rdl(d, v, "show_status_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-status-window"] = t
		}
	}

	if v, ok := d.GetOk("sso"); ok || d.HasChange("sso") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksSso3rdl(d, v, "sso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso"] = t
		}
	}

	if v, ok := d.GetOk("sso_credential"); ok || d.HasChange("sso_credential") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential3rdl(d, v, "sso_credential")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-credential"] = t
		}
	}

	if v, ok := d.GetOk("sso_credential_sent_once"); ok || d.HasChange("sso_credential_sent_once") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce3rdl(d, v, "sso_credential_sent_once")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-credential-sent-once"] = t
		}
	}

	if v, ok := d.GetOk("sso_password"); ok || d.HasChange("sso_password") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword3rdl(d, v, "sso_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-password"] = t
		}
	}

	if v, ok := d.GetOk("sso_username"); ok || d.HasChange("sso_username") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername3rdl(d, v, "sso_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-username"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksUrl3rdl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("vnc_keyboard_layout"); ok || d.HasChange("vnc_keyboard_layout") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout3rdl(d, v, "vnc_keyboard_layout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vnc-keyboard-layout"] = t
		}
	}

	if v, ok := d.GetOk("width"); ok || d.HasChange("width") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksWidth3rdl(d, v, "width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["width"] = t
		}
	}

	return &obj, nil
}
