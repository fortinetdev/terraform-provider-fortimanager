// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Portal bookmark group.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnSslWebPortalBookmarkGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebPortalBookmarkGroupCreate,
		Read:   resourceObjectVpnSslWebPortalBookmarkGroupRead,
		Update: resourceObjectVpnSslWebPortalBookmarkGroupUpdate,
		Delete: resourceObjectVpnSslWebPortalBookmarkGroupDelete,

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
			"bookmarks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceObjectVpnSslWebPortalBookmarkGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["portal"] = portal

	obj, err := getObjectObjectVpnSslWebPortalBookmarkGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalBookmarkGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnSslWebPortalBookmarkGroup(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortalBookmarkGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalBookmarkGroupRead(d, m)
}

func resourceObjectVpnSslWebPortalBookmarkGroupUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["portal"] = portal

	obj, err := getObjectObjectVpnSslWebPortalBookmarkGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalBookmarkGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebPortalBookmarkGroup(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortalBookmarkGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalBookmarkGroupRead(d, m)
}

func resourceObjectVpnSslWebPortalBookmarkGroupDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["portal"] = portal

	err = c.DeleteObjectVpnSslWebPortalBookmarkGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebPortalBookmarkGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebPortalBookmarkGroupRead(d *schema.ResourceData, m interface{}) error {
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
	if portal == "" {
		portal = importOptionChecking(m.(*FortiClient).Cfg, "portal")
		if portal == "" {
			return fmt.Errorf("Parameter portal is missing")
		}
		if err = d.Set("portal", portal); err != nil {
			return fmt.Errorf("Error set params portal: %v", err)
		}
	}
	paradict["portal"] = portal

	o, err := c.ReadObjectVpnSslWebPortalBookmarkGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalBookmarkGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebPortalBookmarkGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortalBookmarkGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarks2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := i["additional-params"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams2edl(i["additional-params"], d, pre_append)
			tmp["additional_params"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-AdditionalParams")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := i["apptype"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksApptype2edl(i["apptype"], d, pre_append)
			tmp["apptype"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Apptype")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := i["color-depth"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth2edl(i["color-depth"], d, pre_append)
			tmp["color_depth"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-ColorDepth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDescription2edl(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDomain2edl(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := i["folder"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFolder2edl(i["folder"], d, pre_append)
			tmp["folder"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Folder")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := i["form-data"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormData2edl(i["form-data"], d, pre_append)
			tmp["form_data"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-FormData")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHeight2edl(i["height"], d, pre_append)
			tmp["height"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Height")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHost2edl(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := i["keyboard-layout"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout2edl(i["keyboard-layout"], d, pre_append)
			tmp["keyboard_layout"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-KeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := i["listening-port"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort2edl(i["listening-port"], d, pre_append)
			tmp["listening_port"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-ListeningPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := i["load-balancing-info"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo2edl(i["load-balancing-info"], d, pre_append)
			tmp["load_balancing_info"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-LoadBalancingInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := i["logon-password"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword2edl(i["logon-password"], d, pre_append)
			tmp["logon_password"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-LogonPassword")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := i["logon-user"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser2edl(i["logon-user"], d, pre_append)
			tmp["logon_user"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-LogonUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPort2edl(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := i["preconnection-blob"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob2edl(i["preconnection-blob"], d, pre_append)
			tmp["preconnection_blob"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-PreconnectionBlob")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := i["preconnection-id"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId2edl(i["preconnection-id"], d, pre_append)
			tmp["preconnection_id"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-PreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := i["remote-port"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort2edl(i["remote-port"], d, pre_append)
			tmp["remote_port"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-RemotePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := i["restricted-admin"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin2edl(i["restricted-admin"], d, pre_append)
			tmp["restricted_admin"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-RestrictedAdmin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity2edl(i["security"], d, pre_append)
			tmp["security"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Security")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := i["send-preconnection-id"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId2edl(i["send-preconnection-id"], d, pre_append)
			tmp["send_preconnection_id"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SendPreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := i["server-layout"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout2edl(i["server-layout"], d, pre_append)
			tmp["server_layout"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-ServerLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := i["show-status-window"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow2edl(i["show-status-window"], d, pre_append)
			tmp["show_status_window"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-ShowStatusWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := i["sso"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSso2edl(i["sso"], d, pre_append)
			tmp["sso"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Sso")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := i["sso-credential"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential2edl(i["sso-credential"], d, pre_append)
			tmp["sso_credential"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SsoCredential")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := i["sso-credential-sent-once"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce2edl(i["sso-credential-sent-once"], d, pre_append)
			tmp["sso_credential_sent_once"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SsoCredentialSentOnce")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := i["sso-password"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword2edl(i["sso-password"], d, pre_append)
			tmp["sso_password"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SsoPassword")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := i["sso-username"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername2edl(i["sso-username"], d, pre_append)
			tmp["sso_username"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SsoUsername")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksUrl2edl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := i["vnc-keyboard-layout"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout2edl(i["vnc-keyboard-layout"], d, pre_append)
			tmp["vnc_keyboard_layout"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-VncKeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksWidth2edl(i["width"], d, pre_append)
			tmp["width"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Width")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksApptype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFolder2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormData2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroupBookmarks-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroupBookmarks-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSso2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksUrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksWidth2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebPortalBookmarkGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("bookmarks", flattenObjectVpnSslWebPortalBookmarkGroupBookmarks2edl(o["bookmarks"], d, "bookmarks")); err != nil {
			if vv, ok := fortiAPIPatch(o["bookmarks"], "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks"); ok {
				if err = d.Set("bookmarks", vv); err != nil {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmarks"); ok {
			if err = d.Set("bookmarks", flattenObjectVpnSslWebPortalBookmarkGroupBookmarks2edl(o["bookmarks"], d, "bookmarks")); err != nil {
				if vv, ok := fortiAPIPatch(o["bookmarks"], "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks"); ok {
					if err = d.Set("bookmarks", vv); err != nil {
						return fmt.Errorf("Error reading bookmarks: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectVpnSslWebPortalBookmarkGroupName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnSslWebPortalBookmarkGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebPortalBookmarkGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarks2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["additional-params"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams2edl(d, i["additional_params"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["apptype"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksApptype2edl(d, i["apptype"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color-depth"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth2edl(d, i["color_depth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksDescription2edl(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksDomain2edl(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["folder"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksFolder2edl(d, i["folder"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormData2edl(d, i["form_data"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["form-data"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["height"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksHeight2edl(d, i["height"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksHost2edl(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyboard-layout"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout2edl(d, i["keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["listening-port"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort2edl(d, i["listening_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["load-balancing-info"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo2edl(d, i["load_balancing_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-password"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword2edl(d, i["logon_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-user"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser2edl(d, i["logon_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksPort2edl(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-blob"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob2edl(d, i["preconnection_blob"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-id"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId2edl(d, i["preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-port"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort2edl(d, i["remote_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restricted-admin"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin2edl(d, i["restricted_admin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity2edl(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-preconnection-id"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId2edl(d, i["send_preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-layout"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout2edl(d, i["server_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["show-status-window"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow2edl(d, i["show_status_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSso2edl(d, i["sso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential2edl(d, i["sso_credential"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential-sent-once"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce2edl(d, i["sso_credential_sent_once"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-password"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword2edl(d, i["sso_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-username"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername2edl(d, i["sso_username"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksUrl2edl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vnc-keyboard-layout"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout2edl(d, i["vnc_keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["width"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksWidth2edl(d, i["width"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksApptype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFolder2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksHeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSso2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksUrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksWidth2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebPortalBookmarkGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bookmarks"); ok || d.HasChange("bookmarks") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarks2edl(d, v, "bookmarks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroupName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
