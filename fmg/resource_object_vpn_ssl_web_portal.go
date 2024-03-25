// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Portal.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnSslWebPortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnSslWebPortalCreate,
		Read:   resourceObjectVpnSslWebPortalRead,
		Update: resourceObjectVpnSslWebPortalUpdate,
		Delete: resourceObjectVpnSslWebPortalDelete,

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
			"allow_user_access": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_connect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bookmark_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Optional: true,
						},
					},
				},
			},
			"client_src_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clipboard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"customize_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_window_height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"default_window_width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhcp_ip_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_ra_giaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_ra_linkaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_connection_tools": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_history": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclusive_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"focus_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heading": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hide_sso_credential": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"host_check_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_pools": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_exclusive_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_pools": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_service_restriction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_split_tunneling_routing_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipv6_split_tunneling_routing_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keep_alive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"landing_page": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"logout_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_credential": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
					},
				},
			},
			"landing_page_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"limit_user_logins": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_addr_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_addr_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_addr_check_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mac_addr_list": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mac_addr_mask": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"macos_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"os_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_check_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"latest_patch_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tolerance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"prefer_ipv6_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redir_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rewrite_ip_uri_ui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"save_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_restriction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_check_for_unsupported_browser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_check_for_browser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_check_for_unsupported_os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_ntlmv1_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smbv1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"split_dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dns_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"domains": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipv6_dns_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_dns_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"split_tunneling_routing_address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"split_tunneling_routing_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transform_backward_slashes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_group_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"windows_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server2": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceObjectVpnSslWebPortalCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectVpnSslWebPortal(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortal resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnSslWebPortal(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnSslWebPortal resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalRead(d, m)
}

func resourceObjectVpnSslWebPortalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnSslWebPortal(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortal resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnSslWebPortal(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnSslWebPortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnSslWebPortalRead(d, m)
}

func resourceObjectVpnSslWebPortalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectVpnSslWebPortal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnSslWebPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnSslWebPortalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectVpnSslWebPortal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnSslWebPortal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnSslWebPortal resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnSslWebPortalAllowUserAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalAutoConnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bookmarks"
		if _, ok := i["bookmarks"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarks(i["bookmarks"], d, pre_append)
			tmp["bookmarks"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-BookmarkGroup-Bookmarks")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-BookmarkGroup-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarks(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(i["additional-params"], d, pre_append)
			tmp["additional_params"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-AdditionalParams")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := i["apptype"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksApptype(i["apptype"], d, pre_append)
			tmp["apptype"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Apptype")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := i["color-depth"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth(i["color-depth"], d, pre_append)
			tmp["color_depth"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-ColorDepth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDescription(i["description"], d, pre_append)
			tmp["description"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Description")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDomain(i["domain"], d, pre_append)
			tmp["domain"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Domain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := i["folder"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFolder(i["folder"], d, pre_append)
			tmp["folder"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Folder")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := i["form-data"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(i["form-data"], d, pre_append)
			tmp["form_data"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-FormData")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHeight(i["height"], d, pre_append)
			tmp["height"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Height")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHost(i["host"], d, pre_append)
			tmp["host"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Host")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := i["keyboard-layout"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(i["keyboard-layout"], d, pre_append)
			tmp["keyboard_layout"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-KeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := i["listening-port"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort(i["listening-port"], d, pre_append)
			tmp["listening_port"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-ListeningPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := i["load-balancing-info"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(i["load-balancing-info"], d, pre_append)
			tmp["load_balancing_info"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-LoadBalancingInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := i["logon-password"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword(i["logon-password"], d, pre_append)
			tmp["logon_password"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-LogonPassword")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := i["logon-user"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser(i["logon-user"], d, pre_append)
			tmp["logon_user"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-LogonUser")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := i["preconnection-blob"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(i["preconnection-blob"], d, pre_append)
			tmp["preconnection_blob"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-PreconnectionBlob")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := i["preconnection-id"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(i["preconnection-id"], d, pre_append)
			tmp["preconnection_id"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-PreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := i["remote-port"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort(i["remote-port"], d, pre_append)
			tmp["remote_port"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-RemotePort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := i["restricted-admin"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(i["restricted-admin"], d, pre_append)
			tmp["restricted_admin"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-RestrictedAdmin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity(i["security"], d, pre_append)
			tmp["security"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Security")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := i["send-preconnection-id"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(i["send-preconnection-id"], d, pre_append)
			tmp["send_preconnection_id"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SendPreconnectionId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := i["server-layout"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout(i["server-layout"], d, pre_append)
			tmp["server_layout"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-ServerLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := i["show-status-window"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(i["show-status-window"], d, pre_append)
			tmp["show_status_window"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-ShowStatusWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := i["sso"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSso(i["sso"], d, pre_append)
			tmp["sso"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Sso")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := i["sso-credential"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(i["sso-credential"], d, pre_append)
			tmp["sso_credential"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SsoCredential")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := i["sso-credential-sent-once"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(i["sso-credential-sent-once"], d, pre_append)
			tmp["sso_credential_sent_once"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SsoCredentialSentOnce")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := i["sso-password"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword(i["sso-password"], d, pre_append)
			tmp["sso_password"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SsoPassword")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := i["sso-username"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(i["sso-username"], d, pre_append)
			tmp["sso_username"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-SsoUsername")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksUrl(i["url"], d, pre_append)
			tmp["url"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Url")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := i["vnc-keyboard-layout"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(i["vnc-keyboard-layout"], d, pre_append)
			tmp["vnc_keyboard_layout"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-VncKeyboardLayout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksWidth(i["width"], d, pre_append)
			tmp["width"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroup-Bookmarks-Width")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksApptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroupBookmarks-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalBookmarkGroupBookmarks-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupBookmarksWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalBookmarkGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalClientSrcRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalClipboard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalCustomLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalCustomizeForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDefaultProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDefaultWindowHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDefaultWindowWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDhcpIpOverlap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDhcpRaGiaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDhcp6RaLinkaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDisplayBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDisplayConnectionTools(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDisplayHistory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDisplayStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalDnsSuffix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalExclusiveRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalFocusBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalForticlientDownload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalForticlientDownloadMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalHeading(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalHideSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalHostCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalHostCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalHostCheckPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectVpnSslWebPortalIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpPools(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectVpnSslWebPortalIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpv6ExclusiveRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpv6Pools(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectVpnSslWebPortalIpv6ServiceRestriction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpv6SplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpv6TunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpv6WinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalIpv6WinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalKeepAlive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPage(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "form_data"
	if _, ok := i["form-data"]; ok {
		result["form_data"] = flattenObjectVpnSslWebPortalLandingPageFormData(i["form-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "logout_url"
	if _, ok := i["logout-url"]; ok {
		result["logout_url"] = flattenObjectVpnSslWebPortalLandingPageLogoutUrl(i["logout-url"], d, pre_append)
	}

	pre_append = pre + ".0." + "sso"
	if _, ok := i["sso"]; ok {
		result["sso"] = flattenObjectVpnSslWebPortalLandingPageSso(i["sso"], d, pre_append)
	}

	pre_append = pre + ".0." + "sso_credential"
	if _, ok := i["sso-credential"]; ok {
		result["sso_credential"] = flattenObjectVpnSslWebPortalLandingPageSsoCredential(i["sso-credential"], d, pre_append)
	}

	pre_append = pre + ".0." + "sso_password"
	if _, ok := i["sso-password"]; ok {
		result["sso_password"] = flattenObjectVpnSslWebPortalLandingPageSsoPassword(i["sso-password"], d, pre_append)
	}

	pre_append = pre + ".0." + "sso_username"
	if _, ok := i["sso-username"]; ok {
		result["sso_username"] = flattenObjectVpnSslWebPortalLandingPageSsoUsername(i["sso-username"], d, pre_append)
	}

	pre_append = pre + ".0." + "url"
	if _, ok := i["url"]; ok {
		result["url"] = flattenObjectVpnSslWebPortalLandingPageUrl(i["url"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectVpnSslWebPortalLandingPageFormData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectVpnSslWebPortalLandingPageFormDataName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalLandingPage-FormData-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectVpnSslWebPortalLandingPageFormDataValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortalLandingPage-FormData-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalLandingPageFormDataName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageFormDataValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageSsoPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalLandingPageSsoUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLandingPageMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalLimitUserLogins(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalMacAddrAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalMacAddrCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalMacAddrCheckRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_list"
		if _, ok := i["mac-addr-list"]; ok {
			v := flattenObjectVpnSslWebPortalMacAddrCheckRuleMacAddrList(i["mac-addr-list"], d, pre_append)
			tmp["mac_addr_list"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-MacAddrCheckRule-MacAddrList")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_mask"
		if _, ok := i["mac-addr-mask"]; ok {
			v := flattenObjectVpnSslWebPortalMacAddrCheckRuleMacAddrMask(i["mac-addr-mask"], d, pre_append)
			tmp["mac_addr_mask"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-MacAddrCheckRule-MacAddrMask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectVpnSslWebPortalMacAddrCheckRuleName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-MacAddrCheckRule-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalMacAddrCheckRuleMacAddrList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalMacAddrCheckRuleMacAddrMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalMacAddrCheckRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalMacosForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalOsCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalOsCheckList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenObjectVpnSslWebPortalOsCheckListAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := i["latest-patch-level"]; ok {
		result["latest_patch_level"] = flattenObjectVpnSslWebPortalOsCheckListLatestPatchLevel(i["latest-patch-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenObjectVpnSslWebPortalOsCheckListName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "tolerance"
	if _, ok := i["tolerance"]; ok {
		result["tolerance"] = flattenObjectVpnSslWebPortalOsCheckListTolerance(i["tolerance"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectVpnSslWebPortalOsCheckListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalOsCheckListLatestPatchLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalOsCheckListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalOsCheckListTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalPreferIpv6Dns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalRedirUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalRewriteIpUriUi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSavePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalServiceRestriction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSkipCheckForUnsupportedBrowser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSkipCheckForBrowser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSkipCheckForUnsupportedOs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSmbMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSmbMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSmbNtlmv1Auth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSmbv1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSplitDns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server1"
		if _, ok := i["dns-server1"]; ok {
			v := flattenObjectVpnSslWebPortalSplitDnsDnsServer1(i["dns-server1"], d, pre_append)
			tmp["dns_server1"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-SplitDns-DnsServer1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server2"
		if _, ok := i["dns-server2"]; ok {
			v := flattenObjectVpnSslWebPortalSplitDnsDnsServer2(i["dns-server2"], d, pre_append)
			tmp["dns_server2"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-SplitDns-DnsServer2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := i["domains"]; ok {
			v := flattenObjectVpnSslWebPortalSplitDnsDomains(i["domains"], d, pre_append)
			tmp["domains"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-SplitDns-Domains")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectVpnSslWebPortalSplitDnsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-SplitDns-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server1"
		if _, ok := i["ipv6-dns-server1"]; ok {
			v := flattenObjectVpnSslWebPortalSplitDnsIpv6DnsServer1(i["ipv6-dns-server1"], d, pre_append)
			tmp["ipv6_dns_server1"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-SplitDns-Ipv6DnsServer1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server2"
		if _, ok := i["ipv6-dns-server2"]; ok {
			v := flattenObjectVpnSslWebPortalSplitDnsIpv6DnsServer2(i["ipv6-dns-server2"], d, pre_append)
			tmp["ipv6_dns_server2"] = fortiAPISubPartPatch(v, "ObjectVpnSslWebPortal-SplitDns-Ipv6DnsServer2")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVpnSslWebPortalSplitDnsDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSplitDnsDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSplitDnsDomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSplitDnsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSplitDnsIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSplitDnsIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalSplitTunnelingRoutingAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnSslWebPortalSplitTunnelingRoutingNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalTransformBackwardSlashes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalTunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalUserBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalUserGroupBookmark(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalWebMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalWindowsForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnSslWebPortalWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnSslWebPortal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("allow_user_access", flattenObjectVpnSslWebPortalAllowUserAccess(o["allow-user-access"], d, "allow_user_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-user-access"], "ObjectVpnSslWebPortal-AllowUserAccess"); ok {
			if err = d.Set("allow_user_access", vv); err != nil {
				return fmt.Errorf("Error reading allow_user_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_user_access: %v", err)
		}
	}

	if err = d.Set("auto_connect", flattenObjectVpnSslWebPortalAutoConnect(o["auto-connect"], d, "auto_connect")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-connect"], "ObjectVpnSslWebPortal-AutoConnect"); ok {
			if err = d.Set("auto_connect", vv); err != nil {
				return fmt.Errorf("Error reading auto_connect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_connect: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("bookmark_group", flattenObjectVpnSslWebPortalBookmarkGroup(o["bookmark-group"], d, "bookmark_group")); err != nil {
			if vv, ok := fortiAPIPatch(o["bookmark-group"], "ObjectVpnSslWebPortal-BookmarkGroup"); ok {
				if err = d.Set("bookmark_group", vv); err != nil {
					return fmt.Errorf("Error reading bookmark_group: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading bookmark_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmark_group"); ok {
			if err = d.Set("bookmark_group", flattenObjectVpnSslWebPortalBookmarkGroup(o["bookmark-group"], d, "bookmark_group")); err != nil {
				if vv, ok := fortiAPIPatch(o["bookmark-group"], "ObjectVpnSslWebPortal-BookmarkGroup"); ok {
					if err = d.Set("bookmark_group", vv); err != nil {
						return fmt.Errorf("Error reading bookmark_group: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading bookmark_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("client_src_range", flattenObjectVpnSslWebPortalClientSrcRange(o["client-src-range"], d, "client_src_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-src-range"], "ObjectVpnSslWebPortal-ClientSrcRange"); ok {
			if err = d.Set("client_src_range", vv); err != nil {
				return fmt.Errorf("Error reading client_src_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_src_range: %v", err)
		}
	}

	if err = d.Set("clipboard", flattenObjectVpnSslWebPortalClipboard(o["clipboard"], d, "clipboard")); err != nil {
		if vv, ok := fortiAPIPatch(o["clipboard"], "ObjectVpnSslWebPortal-Clipboard"); ok {
			if err = d.Set("clipboard", vv); err != nil {
				return fmt.Errorf("Error reading clipboard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clipboard: %v", err)
		}
	}

	if err = d.Set("custom_lang", flattenObjectVpnSslWebPortalCustomLang(o["custom-lang"], d, "custom_lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-lang"], "ObjectVpnSslWebPortal-CustomLang"); ok {
			if err = d.Set("custom_lang", vv); err != nil {
				return fmt.Errorf("Error reading custom_lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_lang: %v", err)
		}
	}

	if err = d.Set("customize_forticlient_download_url", flattenObjectVpnSslWebPortalCustomizeForticlientDownloadUrl(o["customize-forticlient-download-url"], d, "customize_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["customize-forticlient-download-url"], "ObjectVpnSslWebPortal-CustomizeForticlientDownloadUrl"); ok {
			if err = d.Set("customize_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("default_protocol", flattenObjectVpnSslWebPortalDefaultProtocol(o["default-protocol"], d, "default_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-protocol"], "ObjectVpnSslWebPortal-DefaultProtocol"); ok {
			if err = d.Set("default_protocol", vv); err != nil {
				return fmt.Errorf("Error reading default_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_protocol: %v", err)
		}
	}

	if err = d.Set("default_window_height", flattenObjectVpnSslWebPortalDefaultWindowHeight(o["default-window-height"], d, "default_window_height")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-window-height"], "ObjectVpnSslWebPortal-DefaultWindowHeight"); ok {
			if err = d.Set("default_window_height", vv); err != nil {
				return fmt.Errorf("Error reading default_window_height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_window_height: %v", err)
		}
	}

	if err = d.Set("default_window_width", flattenObjectVpnSslWebPortalDefaultWindowWidth(o["default-window-width"], d, "default_window_width")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-window-width"], "ObjectVpnSslWebPortal-DefaultWindowWidth"); ok {
			if err = d.Set("default_window_width", vv); err != nil {
				return fmt.Errorf("Error reading default_window_width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_window_width: %v", err)
		}
	}

	if err = d.Set("dhcp_ip_overlap", flattenObjectVpnSslWebPortalDhcpIpOverlap(o["dhcp-ip-overlap"], d, "dhcp_ip_overlap")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ip-overlap"], "ObjectVpnSslWebPortal-DhcpIpOverlap"); ok {
			if err = d.Set("dhcp_ip_overlap", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ip_overlap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ip_overlap: %v", err)
		}
	}

	if err = d.Set("dhcp_ra_giaddr", flattenObjectVpnSslWebPortalDhcpRaGiaddr(o["dhcp-ra-giaddr"], d, "dhcp_ra_giaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp-ra-giaddr"], "ObjectVpnSslWebPortal-DhcpRaGiaddr"); ok {
			if err = d.Set("dhcp_ra_giaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if err = d.Set("dhcp6_ra_linkaddr", flattenObjectVpnSslWebPortalDhcp6RaLinkaddr(o["dhcp6-ra-linkaddr"], d, "dhcp6_ra_linkaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dhcp6-ra-linkaddr"], "ObjectVpnSslWebPortal-Dhcp6RaLinkaddr"); ok {
			if err = d.Set("dhcp6_ra_linkaddr", vv); err != nil {
				return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
		}
	}

	if err = d.Set("display_bookmark", flattenObjectVpnSslWebPortalDisplayBookmark(o["display-bookmark"], d, "display_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-bookmark"], "ObjectVpnSslWebPortal-DisplayBookmark"); ok {
			if err = d.Set("display_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading display_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_bookmark: %v", err)
		}
	}

	if err = d.Set("display_connection_tools", flattenObjectVpnSslWebPortalDisplayConnectionTools(o["display-connection-tools"], d, "display_connection_tools")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-connection-tools"], "ObjectVpnSslWebPortal-DisplayConnectionTools"); ok {
			if err = d.Set("display_connection_tools", vv); err != nil {
				return fmt.Errorf("Error reading display_connection_tools: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_connection_tools: %v", err)
		}
	}

	if err = d.Set("display_history", flattenObjectVpnSslWebPortalDisplayHistory(o["display-history"], d, "display_history")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-history"], "ObjectVpnSslWebPortal-DisplayHistory"); ok {
			if err = d.Set("display_history", vv); err != nil {
				return fmt.Errorf("Error reading display_history: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_history: %v", err)
		}
	}

	if err = d.Set("display_status", flattenObjectVpnSslWebPortalDisplayStatus(o["display-status"], d, "display_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["display-status"], "ObjectVpnSslWebPortal-DisplayStatus"); ok {
			if err = d.Set("display_status", vv); err != nil {
				return fmt.Errorf("Error reading display_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading display_status: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenObjectVpnSslWebPortalDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server1"], "ObjectVpnSslWebPortal-DnsServer1"); ok {
			if err = d.Set("dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenObjectVpnSslWebPortalDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-server2"], "ObjectVpnSslWebPortal-DnsServer2"); ok {
			if err = d.Set("dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_suffix", flattenObjectVpnSslWebPortalDnsSuffix(o["dns-suffix"], d, "dns_suffix")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-suffix"], "ObjectVpnSslWebPortal-DnsSuffix"); ok {
			if err = d.Set("dns_suffix", vv); err != nil {
				return fmt.Errorf("Error reading dns_suffix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_suffix: %v", err)
		}
	}

	if err = d.Set("exclusive_routing", flattenObjectVpnSslWebPortalExclusiveRouting(o["exclusive-routing"], d, "exclusive_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclusive-routing"], "ObjectVpnSslWebPortal-ExclusiveRouting"); ok {
			if err = d.Set("exclusive_routing", vv); err != nil {
				return fmt.Errorf("Error reading exclusive_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclusive_routing: %v", err)
		}
	}

	if err = d.Set("focus_bookmark", flattenObjectVpnSslWebPortalFocusBookmark(o["focus-bookmark"], d, "focus_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["focus-bookmark"], "ObjectVpnSslWebPortal-FocusBookmark"); ok {
			if err = d.Set("focus_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading focus_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading focus_bookmark: %v", err)
		}
	}

	if err = d.Set("forticlient_download", flattenObjectVpnSslWebPortalForticlientDownload(o["forticlient-download"], d, "forticlient_download")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-download"], "ObjectVpnSslWebPortal-ForticlientDownload"); ok {
			if err = d.Set("forticlient_download", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_download: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_download: %v", err)
		}
	}

	if err = d.Set("forticlient_download_method", flattenObjectVpnSslWebPortalForticlientDownloadMethod(o["forticlient-download-method"], d, "forticlient_download_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["forticlient-download-method"], "ObjectVpnSslWebPortal-ForticlientDownloadMethod"); ok {
			if err = d.Set("forticlient_download_method", vv); err != nil {
				return fmt.Errorf("Error reading forticlient_download_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forticlient_download_method: %v", err)
		}
	}

	if err = d.Set("heading", flattenObjectVpnSslWebPortalHeading(o["heading"], d, "heading")); err != nil {
		if vv, ok := fortiAPIPatch(o["heading"], "ObjectVpnSslWebPortal-Heading"); ok {
			if err = d.Set("heading", vv); err != nil {
				return fmt.Errorf("Error reading heading: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading heading: %v", err)
		}
	}

	if err = d.Set("hide_sso_credential", flattenObjectVpnSslWebPortalHideSsoCredential(o["hide-sso-credential"], d, "hide_sso_credential")); err != nil {
		if vv, ok := fortiAPIPatch(o["hide-sso-credential"], "ObjectVpnSslWebPortal-HideSsoCredential"); ok {
			if err = d.Set("hide_sso_credential", vv); err != nil {
				return fmt.Errorf("Error reading hide_sso_credential: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hide_sso_credential: %v", err)
		}
	}

	if err = d.Set("host_check", flattenObjectVpnSslWebPortalHostCheck(o["host-check"], d, "host_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-check"], "ObjectVpnSslWebPortal-HostCheck"); ok {
			if err = d.Set("host_check", vv); err != nil {
				return fmt.Errorf("Error reading host_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_check: %v", err)
		}
	}

	if err = d.Set("host_check_interval", flattenObjectVpnSslWebPortalHostCheckInterval(o["host-check-interval"], d, "host_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-check-interval"], "ObjectVpnSslWebPortal-HostCheckInterval"); ok {
			if err = d.Set("host_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading host_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_check_interval: %v", err)
		}
	}

	if err = d.Set("host_check_policy", flattenObjectVpnSslWebPortalHostCheckPolicy(o["host-check-policy"], d, "host_check_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-check-policy"], "ObjectVpnSslWebPortal-HostCheckPolicy"); ok {
			if err = d.Set("host_check_policy", vv); err != nil {
				return fmt.Errorf("Error reading host_check_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_check_policy: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenObjectVpnSslWebPortalIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-mode"], "ObjectVpnSslWebPortal-IpMode"); ok {
			if err = d.Set("ip_mode", vv); err != nil {
				return fmt.Errorf("Error reading ip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if err = d.Set("ip_pools", flattenObjectVpnSslWebPortalIpPools(o["ip-pools"], d, "ip_pools")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-pools"], "ObjectVpnSslWebPortal-IpPools"); ok {
			if err = d.Set("ip_pools", vv); err != nil {
				return fmt.Errorf("Error reading ip_pools: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_pools: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenObjectVpnSslWebPortalIpv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server1"], "ObjectVpnSslWebPortal-Ipv6DnsServer1"); ok {
			if err = d.Set("ipv6_dns_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenObjectVpnSslWebPortalIpv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-dns-server2"], "ObjectVpnSslWebPortal-Ipv6DnsServer2"); ok {
			if err = d.Set("ipv6_dns_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_exclusive_routing", flattenObjectVpnSslWebPortalIpv6ExclusiveRouting(o["ipv6-exclusive-routing"], d, "ipv6_exclusive_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-exclusive-routing"], "ObjectVpnSslWebPortal-Ipv6ExclusiveRouting"); ok {
			if err = d.Set("ipv6_exclusive_routing", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_exclusive_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_exclusive_routing: %v", err)
		}
	}

	if err = d.Set("ipv6_pools", flattenObjectVpnSslWebPortalIpv6Pools(o["ipv6-pools"], d, "ipv6_pools")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-pools"], "ObjectVpnSslWebPortal-Ipv6Pools"); ok {
			if err = d.Set("ipv6_pools", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_pools: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_pools: %v", err)
		}
	}

	if err = d.Set("ipv6_service_restriction", flattenObjectVpnSslWebPortalIpv6ServiceRestriction(o["ipv6-service-restriction"], d, "ipv6_service_restriction")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-service-restriction"], "ObjectVpnSslWebPortal-Ipv6ServiceRestriction"); ok {
			if err = d.Set("ipv6_service_restriction", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_service_restriction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_service_restriction: %v", err)
		}
	}

	if err = d.Set("ipv6_split_tunneling", flattenObjectVpnSslWebPortalIpv6SplitTunneling(o["ipv6-split-tunneling"], d, "ipv6_split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-tunneling"], "ObjectVpnSslWebPortal-Ipv6SplitTunneling"); ok {
			if err = d.Set("ipv6_split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_tunneling: %v", err)
		}
	}

	if err = d.Set("ipv6_split_tunneling_routing_address", flattenObjectVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(o["ipv6-split-tunneling-routing-address"], d, "ipv6_split_tunneling_routing_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-tunneling-routing-address"], "ObjectVpnSslWebPortal-Ipv6SplitTunnelingRoutingAddress"); ok {
			if err = d.Set("ipv6_split_tunneling_routing_address", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_tunneling_routing_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_tunneling_routing_address: %v", err)
		}
	}

	if err = d.Set("ipv6_split_tunneling_routing_negate", flattenObjectVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(o["ipv6-split-tunneling-routing-negate"], d, "ipv6_split_tunneling_routing_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-split-tunneling-routing-negate"], "ObjectVpnSslWebPortal-Ipv6SplitTunnelingRoutingNegate"); ok {
			if err = d.Set("ipv6_split_tunneling_routing_negate", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_split_tunneling_routing_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_split_tunneling_routing_negate: %v", err)
		}
	}

	if err = d.Set("ipv6_tunnel_mode", flattenObjectVpnSslWebPortalIpv6TunnelMode(o["ipv6-tunnel-mode"], d, "ipv6_tunnel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-tunnel-mode"], "ObjectVpnSslWebPortal-Ipv6TunnelMode"); ok {
			if err = d.Set("ipv6_tunnel_mode", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_tunnel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_tunnel_mode: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server1", flattenObjectVpnSslWebPortalIpv6WinsServer1(o["ipv6-wins-server1"], d, "ipv6_wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-wins-server1"], "ObjectVpnSslWebPortal-Ipv6WinsServer1"); ok {
			if err = d.Set("ipv6_wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server2", flattenObjectVpnSslWebPortalIpv6WinsServer2(o["ipv6-wins-server2"], d, "ipv6_wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-wins-server2"], "ObjectVpnSslWebPortal-Ipv6WinsServer2"); ok {
			if err = d.Set("ipv6_wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
		}
	}

	if err = d.Set("keep_alive", flattenObjectVpnSslWebPortalKeepAlive(o["keep-alive"], d, "keep_alive")); err != nil {
		if vv, ok := fortiAPIPatch(o["keep-alive"], "ObjectVpnSslWebPortal-KeepAlive"); ok {
			if err = d.Set("keep_alive", vv); err != nil {
				return fmt.Errorf("Error reading keep_alive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keep_alive: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("landing_page", flattenObjectVpnSslWebPortalLandingPage(o["landing-page"], d, "landing_page")); err != nil {
			if vv, ok := fortiAPIPatch(o["landing-page"], "ObjectVpnSslWebPortal-LandingPage"); ok {
				if err = d.Set("landing_page", vv); err != nil {
					return fmt.Errorf("Error reading landing_page: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading landing_page: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("landing_page"); ok {
			if err = d.Set("landing_page", flattenObjectVpnSslWebPortalLandingPage(o["landing-page"], d, "landing_page")); err != nil {
				if vv, ok := fortiAPIPatch(o["landing-page"], "ObjectVpnSslWebPortal-LandingPage"); ok {
					if err = d.Set("landing_page", vv); err != nil {
						return fmt.Errorf("Error reading landing_page: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading landing_page: %v", err)
				}
			}
		}
	}

	if err = d.Set("landing_page_mode", flattenObjectVpnSslWebPortalLandingPageMode(o["landing-page-mode"], d, "landing_page_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["landing-page-mode"], "ObjectVpnSslWebPortal-LandingPageMode"); ok {
			if err = d.Set("landing_page_mode", vv); err != nil {
				return fmt.Errorf("Error reading landing_page_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading landing_page_mode: %v", err)
		}
	}

	if err = d.Set("limit_user_logins", flattenObjectVpnSslWebPortalLimitUserLogins(o["limit-user-logins"], d, "limit_user_logins")); err != nil {
		if vv, ok := fortiAPIPatch(o["limit-user-logins"], "ObjectVpnSslWebPortal-LimitUserLogins"); ok {
			if err = d.Set("limit_user_logins", vv); err != nil {
				return fmt.Errorf("Error reading limit_user_logins: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading limit_user_logins: %v", err)
		}
	}

	if err = d.Set("mac_addr_action", flattenObjectVpnSslWebPortalMacAddrAction(o["mac-addr-action"], d, "mac_addr_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-addr-action"], "ObjectVpnSslWebPortal-MacAddrAction"); ok {
			if err = d.Set("mac_addr_action", vv); err != nil {
				return fmt.Errorf("Error reading mac_addr_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_addr_action: %v", err)
		}
	}

	if err = d.Set("mac_addr_check", flattenObjectVpnSslWebPortalMacAddrCheck(o["mac-addr-check"], d, "mac_addr_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-addr-check"], "ObjectVpnSslWebPortal-MacAddrCheck"); ok {
			if err = d.Set("mac_addr_check", vv); err != nil {
				return fmt.Errorf("Error reading mac_addr_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_addr_check: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mac_addr_check_rule", flattenObjectVpnSslWebPortalMacAddrCheckRule(o["mac-addr-check-rule"], d, "mac_addr_check_rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["mac-addr-check-rule"], "ObjectVpnSslWebPortal-MacAddrCheckRule"); ok {
				if err = d.Set("mac_addr_check_rule", vv); err != nil {
					return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mac_addr_check_rule"); ok {
			if err = d.Set("mac_addr_check_rule", flattenObjectVpnSslWebPortalMacAddrCheckRule(o["mac-addr-check-rule"], d, "mac_addr_check_rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["mac-addr-check-rule"], "ObjectVpnSslWebPortal-MacAddrCheckRule"); ok {
					if err = d.Set("mac_addr_check_rule", vv); err != nil {
						return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("macos_forticlient_download_url", flattenObjectVpnSslWebPortalMacosForticlientDownloadUrl(o["macos-forticlient-download-url"], d, "macos_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["macos-forticlient-download-url"], "ObjectVpnSslWebPortal-MacosForticlientDownloadUrl"); ok {
			if err = d.Set("macos_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnSslWebPortalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnSslWebPortal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os_check", flattenObjectVpnSslWebPortalOsCheck(o["os-check"], d, "os_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["os-check"], "ObjectVpnSslWebPortal-OsCheck"); ok {
			if err = d.Set("os_check", vv); err != nil {
				return fmt.Errorf("Error reading os_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_check: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("os_check_list", flattenObjectVpnSslWebPortalOsCheckList(o["os-check-list"], d, "os_check_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["os-check-list"], "ObjectVpnSslWebPortal-OsCheckList"); ok {
				if err = d.Set("os_check_list", vv); err != nil {
					return fmt.Errorf("Error reading os_check_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading os_check_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("os_check_list"); ok {
			if err = d.Set("os_check_list", flattenObjectVpnSslWebPortalOsCheckList(o["os-check-list"], d, "os_check_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["os-check-list"], "ObjectVpnSslWebPortal-OsCheckList"); ok {
					if err = d.Set("os_check_list", vv); err != nil {
						return fmt.Errorf("Error reading os_check_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading os_check_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("prefer_ipv6_dns", flattenObjectVpnSslWebPortalPreferIpv6Dns(o["prefer-ipv6-dns"], d, "prefer_ipv6_dns")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefer-ipv6-dns"], "ObjectVpnSslWebPortal-PreferIpv6Dns"); ok {
			if err = d.Set("prefer_ipv6_dns", vv); err != nil {
				return fmt.Errorf("Error reading prefer_ipv6_dns: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefer_ipv6_dns: %v", err)
		}
	}

	if err = d.Set("redir_url", flattenObjectVpnSslWebPortalRedirUrl(o["redir-url"], d, "redir_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["redir-url"], "ObjectVpnSslWebPortal-RedirUrl"); ok {
			if err = d.Set("redir_url", vv); err != nil {
				return fmt.Errorf("Error reading redir_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redir_url: %v", err)
		}
	}

	if err = d.Set("rewrite_ip_uri_ui", flattenObjectVpnSslWebPortalRewriteIpUriUi(o["rewrite-ip-uri-ui"], d, "rewrite_ip_uri_ui")); err != nil {
		if vv, ok := fortiAPIPatch(o["rewrite-ip-uri-ui"], "ObjectVpnSslWebPortal-RewriteIpUriUi"); ok {
			if err = d.Set("rewrite_ip_uri_ui", vv); err != nil {
				return fmt.Errorf("Error reading rewrite_ip_uri_ui: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rewrite_ip_uri_ui: %v", err)
		}
	}

	if err = d.Set("save_password", flattenObjectVpnSslWebPortalSavePassword(o["save-password"], d, "save_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["save-password"], "ObjectVpnSslWebPortal-SavePassword"); ok {
			if err = d.Set("save_password", vv); err != nil {
				return fmt.Errorf("Error reading save_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading save_password: %v", err)
		}
	}

	if err = d.Set("service_restriction", flattenObjectVpnSslWebPortalServiceRestriction(o["service-restriction"], d, "service_restriction")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-restriction"], "ObjectVpnSslWebPortal-ServiceRestriction"); ok {
			if err = d.Set("service_restriction", vv); err != nil {
				return fmt.Errorf("Error reading service_restriction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_restriction: %v", err)
		}
	}

	if err = d.Set("skip_check_for_unsupported_browser", flattenObjectVpnSslWebPortalSkipCheckForUnsupportedBrowser(o["skip-check-for-unsupported-browser"], d, "skip_check_for_unsupported_browser")); err != nil {
		if vv, ok := fortiAPIPatch(o["skip-check-for-unsupported-browser"], "ObjectVpnSslWebPortal-SkipCheckForUnsupportedBrowser"); ok {
			if err = d.Set("skip_check_for_unsupported_browser", vv); err != nil {
				return fmt.Errorf("Error reading skip_check_for_unsupported_browser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skip_check_for_unsupported_browser: %v", err)
		}
	}

	if err = d.Set("skip_check_for_browser", flattenObjectVpnSslWebPortalSkipCheckForBrowser(o["skip-check-for-browser"], d, "skip_check_for_browser")); err != nil {
		if vv, ok := fortiAPIPatch(o["skip-check-for-browser"], "ObjectVpnSslWebPortal-SkipCheckForBrowser"); ok {
			if err = d.Set("skip_check_for_browser", vv); err != nil {
				return fmt.Errorf("Error reading skip_check_for_browser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skip_check_for_browser: %v", err)
		}
	}

	if err = d.Set("skip_check_for_unsupported_os", flattenObjectVpnSslWebPortalSkipCheckForUnsupportedOs(o["skip-check-for-unsupported-os"], d, "skip_check_for_unsupported_os")); err != nil {
		if vv, ok := fortiAPIPatch(o["skip-check-for-unsupported-os"], "ObjectVpnSslWebPortal-SkipCheckForUnsupportedOs"); ok {
			if err = d.Set("skip_check_for_unsupported_os", vv); err != nil {
				return fmt.Errorf("Error reading skip_check_for_unsupported_os: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading skip_check_for_unsupported_os: %v", err)
		}
	}

	if err = d.Set("smb_max_version", flattenObjectVpnSslWebPortalSmbMaxVersion(o["smb-max-version"], d, "smb_max_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["smb-max-version"], "ObjectVpnSslWebPortal-SmbMaxVersion"); ok {
			if err = d.Set("smb_max_version", vv); err != nil {
				return fmt.Errorf("Error reading smb_max_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smb_max_version: %v", err)
		}
	}

	if err = d.Set("smb_min_version", flattenObjectVpnSslWebPortalSmbMinVersion(o["smb-min-version"], d, "smb_min_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["smb-min-version"], "ObjectVpnSslWebPortal-SmbMinVersion"); ok {
			if err = d.Set("smb_min_version", vv); err != nil {
				return fmt.Errorf("Error reading smb_min_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smb_min_version: %v", err)
		}
	}

	if err = d.Set("smb_ntlmv1_auth", flattenObjectVpnSslWebPortalSmbNtlmv1Auth(o["smb-ntlmv1-auth"], d, "smb_ntlmv1_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["smb-ntlmv1-auth"], "ObjectVpnSslWebPortal-SmbNtlmv1Auth"); ok {
			if err = d.Set("smb_ntlmv1_auth", vv); err != nil {
				return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
		}
	}

	if err = d.Set("smbv1", flattenObjectVpnSslWebPortalSmbv1(o["smbv1"], d, "smbv1")); err != nil {
		if vv, ok := fortiAPIPatch(o["smbv1"], "ObjectVpnSslWebPortal-Smbv1"); ok {
			if err = d.Set("smbv1", vv); err != nil {
				return fmt.Errorf("Error reading smbv1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smbv1: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("split_dns", flattenObjectVpnSslWebPortalSplitDns(o["split-dns"], d, "split_dns")); err != nil {
			if vv, ok := fortiAPIPatch(o["split-dns"], "ObjectVpnSslWebPortal-SplitDns"); ok {
				if err = d.Set("split_dns", vv); err != nil {
					return fmt.Errorf("Error reading split_dns: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading split_dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_dns"); ok {
			if err = d.Set("split_dns", flattenObjectVpnSslWebPortalSplitDns(o["split-dns"], d, "split_dns")); err != nil {
				if vv, ok := fortiAPIPatch(o["split-dns"], "ObjectVpnSslWebPortal-SplitDns"); ok {
					if err = d.Set("split_dns", vv); err != nil {
						return fmt.Errorf("Error reading split_dns: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading split_dns: %v", err)
				}
			}
		}
	}

	if err = d.Set("split_tunneling", flattenObjectVpnSslWebPortalSplitTunneling(o["split-tunneling"], d, "split_tunneling")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling"], "ObjectVpnSslWebPortal-SplitTunneling"); ok {
			if err = d.Set("split_tunneling", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling: %v", err)
		}
	}

	if err = d.Set("split_tunneling_routing_address", flattenObjectVpnSslWebPortalSplitTunnelingRoutingAddress(o["split-tunneling-routing-address"], d, "split_tunneling_routing_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling-routing-address"], "ObjectVpnSslWebPortal-SplitTunnelingRoutingAddress"); ok {
			if err = d.Set("split_tunneling_routing_address", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling_routing_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling_routing_address: %v", err)
		}
	}

	if err = d.Set("split_tunneling_routing_negate", flattenObjectVpnSslWebPortalSplitTunnelingRoutingNegate(o["split-tunneling-routing-negate"], d, "split_tunneling_routing_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["split-tunneling-routing-negate"], "ObjectVpnSslWebPortal-SplitTunnelingRoutingNegate"); ok {
			if err = d.Set("split_tunneling_routing_negate", vv); err != nil {
				return fmt.Errorf("Error reading split_tunneling_routing_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading split_tunneling_routing_negate: %v", err)
		}
	}

	if err = d.Set("theme", flattenObjectVpnSslWebPortalTheme(o["theme"], d, "theme")); err != nil {
		if vv, ok := fortiAPIPatch(o["theme"], "ObjectVpnSslWebPortal-Theme"); ok {
			if err = d.Set("theme", vv); err != nil {
				return fmt.Errorf("Error reading theme: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading theme: %v", err)
		}
	}

	if err = d.Set("transform_backward_slashes", flattenObjectVpnSslWebPortalTransformBackwardSlashes(o["transform-backward-slashes"], d, "transform_backward_slashes")); err != nil {
		if vv, ok := fortiAPIPatch(o["transform-backward-slashes"], "ObjectVpnSslWebPortal-TransformBackwardSlashes"); ok {
			if err = d.Set("transform_backward_slashes", vv); err != nil {
				return fmt.Errorf("Error reading transform_backward_slashes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transform_backward_slashes: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenObjectVpnSslWebPortalTunnelMode(o["tunnel-mode"], d, "tunnel_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-mode"], "ObjectVpnSslWebPortal-TunnelMode"); ok {
			if err = d.Set("tunnel_mode", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	if err = d.Set("use_sdwan", flattenObjectVpnSslWebPortalUseSdwan(o["use-sdwan"], d, "use_sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-sdwan"], "ObjectVpnSslWebPortal-UseSdwan"); ok {
			if err = d.Set("use_sdwan", vv); err != nil {
				return fmt.Errorf("Error reading use_sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_sdwan: %v", err)
		}
	}

	if err = d.Set("user_bookmark", flattenObjectVpnSslWebPortalUserBookmark(o["user-bookmark"], d, "user_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-bookmark"], "ObjectVpnSslWebPortal-UserBookmark"); ok {
			if err = d.Set("user_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading user_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_bookmark: %v", err)
		}
	}

	if err = d.Set("user_group_bookmark", flattenObjectVpnSslWebPortalUserGroupBookmark(o["user-group-bookmark"], d, "user_group_bookmark")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-group-bookmark"], "ObjectVpnSslWebPortal-UserGroupBookmark"); ok {
			if err = d.Set("user_group_bookmark", vv); err != nil {
				return fmt.Errorf("Error reading user_group_bookmark: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_group_bookmark: %v", err)
		}
	}

	if err = d.Set("web_mode", flattenObjectVpnSslWebPortalWebMode(o["web-mode"], d, "web_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["web-mode"], "ObjectVpnSslWebPortal-WebMode"); ok {
			if err = d.Set("web_mode", vv); err != nil {
				return fmt.Errorf("Error reading web_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading web_mode: %v", err)
		}
	}

	if err = d.Set("windows_forticlient_download_url", flattenObjectVpnSslWebPortalWindowsForticlientDownloadUrl(o["windows-forticlient-download-url"], d, "windows_forticlient_download_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["windows-forticlient-download-url"], "ObjectVpnSslWebPortal-WindowsForticlientDownloadUrl"); ok {
			if err = d.Set("windows_forticlient_download_url", vv); err != nil {
				return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenObjectVpnSslWebPortalWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server1"], "ObjectVpnSslWebPortal-WinsServer1"); ok {
			if err = d.Set("wins_server1", vv); err != nil {
				return fmt.Errorf("Error reading wins_server1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenObjectVpnSslWebPortalWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wins-server2"], "ObjectVpnSslWebPortal-WinsServer2"); ok {
			if err = d.Set("wins_server2", vv); err != nil {
				return fmt.Errorf("Error reading wins_server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnSslWebPortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnSslWebPortalAllowUserAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalAutoConnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bookmarks"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarks(d, i["bookmarks"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["bookmarks"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectVpnSslWebPortalBookmarkGroupName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["additional-params"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(d, i["additional_params"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["apptype"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksApptype(d, i["apptype"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["color-depth"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth(d, i["color_depth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["description"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domain"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksDomain(d, i["domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["folder"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksFolder(d, i["folder"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(d, i["form_data"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["form-data"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["height"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksHeight(d, i["height"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["host"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keyboard-layout"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["listening-port"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort(d, i["listening_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["load-balancing-info"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-password"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword(d, i["logon_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["logon-user"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser(d, i["logon_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-blob"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preconnection-id"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(d, i["preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["remote-port"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort(d, i["remote_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["restricted-admin"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["send-preconnection-id"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-layout"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout(d, i["server_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["show-status-window"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(d, i["show_status_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSso(d, i["sso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(d, i["sso_credential"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-credential-sent-once"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(d, i["sso_credential_sent_once"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-password"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword(d, i["sso_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-username"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(d, i["sso_username"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["url"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vnc-keyboard-layout"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(d, i["vnc_keyboard_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["width"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksWidth(d, i["width"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksApptype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFolder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksListeningPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksRemotePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksServerLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupBookmarksWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalBookmarkGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalClientSrcRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalClipboard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalCustomLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalCustomizeForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDefaultProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDefaultWindowHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDefaultWindowWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDhcpIpOverlap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDhcpRaGiaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDhcp6RaLinkaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDisplayBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDisplayConnectionTools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDisplayHistory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDisplayStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalDnsSuffix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalExclusiveRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalFocusBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalForticlientDownload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalForticlientDownloadMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalHeading(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalHideSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalHostCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalHostCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalHostCheckPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectVpnSslWebPortalIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpPools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectVpnSslWebPortalIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpv6ExclusiveRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpv6Pools(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectVpnSslWebPortalIpv6ServiceRestriction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpv6SplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpv6TunnelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpv6WinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalIpv6WinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalKeepAlive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "form_data"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectVpnSslWebPortalLandingPageFormData(d, i["form_data"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["form-data"] = t
		}
	}
	pre_append = pre + ".0." + "logout_url"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["logout-url"], _ = expandObjectVpnSslWebPortalLandingPageLogoutUrl(d, i["logout_url"], pre_append)
	}
	pre_append = pre + ".0." + "sso"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sso"], _ = expandObjectVpnSslWebPortalLandingPageSso(d, i["sso"], pre_append)
	}
	pre_append = pre + ".0." + "sso_credential"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sso-credential"], _ = expandObjectVpnSslWebPortalLandingPageSsoCredential(d, i["sso_credential"], pre_append)
	}
	pre_append = pre + ".0." + "sso_password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sso-password"], _ = expandObjectVpnSslWebPortalLandingPageSsoPassword(d, i["sso_password"], pre_append)
	}
	pre_append = pre + ".0." + "sso_username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sso-username"], _ = expandObjectVpnSslWebPortalLandingPageSsoUsername(d, i["sso_username"], pre_append)
	}
	pre_append = pre + ".0." + "url"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["url"], _ = expandObjectVpnSslWebPortalLandingPageUrl(d, i["url"], pre_append)
	}

	return result, nil
}

func expandObjectVpnSslWebPortalLandingPageFormData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectVpnSslWebPortalLandingPageFormDataName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandObjectVpnSslWebPortalLandingPageFormDataValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalLandingPageFormDataName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageFormDataValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageSsoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalLandingPageSsoUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLandingPageMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalLimitUserLogins(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalMacAddrAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalMacAddrCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalMacAddrCheckRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-addr-list"], _ = expandObjectVpnSslWebPortalMacAddrCheckRuleMacAddrList(d, i["mac_addr_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-addr-mask"], _ = expandObjectVpnSslWebPortalMacAddrCheckRuleMacAddrMask(d, i["mac_addr_mask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectVpnSslWebPortalMacAddrCheckRuleName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalMacAddrCheckRuleMacAddrList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalMacAddrCheckRuleMacAddrMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalMacAddrCheckRuleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalMacosForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalOsCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalOsCheckList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["action"], _ = expandObjectVpnSslWebPortalOsCheckListAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "latest_patch_level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["latest-patch-level"], _ = expandObjectVpnSslWebPortalOsCheckListLatestPatchLevel(d, i["latest_patch_level"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandObjectVpnSslWebPortalOsCheckListName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "tolerance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tolerance"], _ = expandObjectVpnSslWebPortalOsCheckListTolerance(d, i["tolerance"], pre_append)
	}

	return result, nil
}

func expandObjectVpnSslWebPortalOsCheckListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalOsCheckListLatestPatchLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalOsCheckListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalOsCheckListTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalPreferIpv6Dns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalRedirUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalRewriteIpUriUi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSavePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalServiceRestriction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSkipCheckForUnsupportedBrowser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSkipCheckForBrowser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSkipCheckForUnsupportedOs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSmbMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSmbMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSmbNtlmv1Auth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSmbv1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSplitDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-server1"], _ = expandObjectVpnSslWebPortalSplitDnsDnsServer1(d, i["dns_server1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-server2"], _ = expandObjectVpnSslWebPortalSplitDnsDnsServer2(d, i["dns_server2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["domains"], _ = expandObjectVpnSslWebPortalSplitDnsDomains(d, i["domains"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectVpnSslWebPortalSplitDnsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-dns-server1"], _ = expandObjectVpnSslWebPortalSplitDnsIpv6DnsServer1(d, i["ipv6_dns_server1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ipv6-dns-server2"], _ = expandObjectVpnSslWebPortalSplitDnsIpv6DnsServer2(d, i["ipv6_dns_server2"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVpnSslWebPortalSplitDnsDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSplitDnsDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSplitDnsDomains(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSplitDnsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSplitDnsIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSplitDnsIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalSplitTunnelingRoutingAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnSslWebPortalSplitTunnelingRoutingNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalTheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalTransformBackwardSlashes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalTunnelMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalUseSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalUserBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalUserGroupBookmark(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalWebMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalWindowsForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnSslWebPortalWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnSslWebPortal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_user_access"); ok || d.HasChange("allow_user_access") {
		t, err := expandObjectVpnSslWebPortalAllowUserAccess(d, v, "allow_user_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-user-access"] = t
		}
	}

	if v, ok := d.GetOk("auto_connect"); ok || d.HasChange("auto_connect") {
		t, err := expandObjectVpnSslWebPortalAutoConnect(d, v, "auto_connect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-connect"] = t
		}
	}

	if v, ok := d.GetOk("bookmark_group"); ok || d.HasChange("bookmark_group") {
		t, err := expandObjectVpnSslWebPortalBookmarkGroup(d, v, "bookmark_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmark-group"] = t
		}
	}

	if v, ok := d.GetOk("client_src_range"); ok || d.HasChange("client_src_range") {
		t, err := expandObjectVpnSslWebPortalClientSrcRange(d, v, "client_src_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-src-range"] = t
		}
	}

	if v, ok := d.GetOk("clipboard"); ok || d.HasChange("clipboard") {
		t, err := expandObjectVpnSslWebPortalClipboard(d, v, "clipboard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clipboard"] = t
		}
	}

	if v, ok := d.GetOk("custom_lang"); ok || d.HasChange("custom_lang") {
		t, err := expandObjectVpnSslWebPortalCustomLang(d, v, "custom_lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-lang"] = t
		}
	}

	if v, ok := d.GetOk("customize_forticlient_download_url"); ok || d.HasChange("customize_forticlient_download_url") {
		t, err := expandObjectVpnSslWebPortalCustomizeForticlientDownloadUrl(d, v, "customize_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customize-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("default_protocol"); ok || d.HasChange("default_protocol") {
		t, err := expandObjectVpnSslWebPortalDefaultProtocol(d, v, "default_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-protocol"] = t
		}
	}

	if v, ok := d.GetOk("default_window_height"); ok || d.HasChange("default_window_height") {
		t, err := expandObjectVpnSslWebPortalDefaultWindowHeight(d, v, "default_window_height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-height"] = t
		}
	}

	if v, ok := d.GetOk("default_window_width"); ok || d.HasChange("default_window_width") {
		t, err := expandObjectVpnSslWebPortalDefaultWindowWidth(d, v, "default_window_width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-width"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ip_overlap"); ok || d.HasChange("dhcp_ip_overlap") {
		t, err := expandObjectVpnSslWebPortalDhcpIpOverlap(d, v, "dhcp_ip_overlap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ip-overlap"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ra_giaddr"); ok || d.HasChange("dhcp_ra_giaddr") {
		t, err := expandObjectVpnSslWebPortalDhcpRaGiaddr(d, v, "dhcp_ra_giaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ra-giaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_ra_linkaddr"); ok || d.HasChange("dhcp6_ra_linkaddr") {
		t, err := expandObjectVpnSslWebPortalDhcp6RaLinkaddr(d, v, "dhcp6_ra_linkaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-ra-linkaddr"] = t
		}
	}

	if v, ok := d.GetOk("display_bookmark"); ok || d.HasChange("display_bookmark") {
		t, err := expandObjectVpnSslWebPortalDisplayBookmark(d, v, "display_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("display_connection_tools"); ok || d.HasChange("display_connection_tools") {
		t, err := expandObjectVpnSslWebPortalDisplayConnectionTools(d, v, "display_connection_tools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-connection-tools"] = t
		}
	}

	if v, ok := d.GetOk("display_history"); ok || d.HasChange("display_history") {
		t, err := expandObjectVpnSslWebPortalDisplayHistory(d, v, "display_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-history"] = t
		}
	}

	if v, ok := d.GetOk("display_status"); ok || d.HasChange("display_status") {
		t, err := expandObjectVpnSslWebPortalDisplayStatus(d, v, "display_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-status"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok || d.HasChange("dns_server1") {
		t, err := expandObjectVpnSslWebPortalDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok || d.HasChange("dns_server2") {
		t, err := expandObjectVpnSslWebPortalDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_suffix"); ok || d.HasChange("dns_suffix") {
		t, err := expandObjectVpnSslWebPortalDnsSuffix(d, v, "dns_suffix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-suffix"] = t
		}
	}

	if v, ok := d.GetOk("exclusive_routing"); ok || d.HasChange("exclusive_routing") {
		t, err := expandObjectVpnSslWebPortalExclusiveRouting(d, v, "exclusive_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclusive-routing"] = t
		}
	}

	if v, ok := d.GetOk("focus_bookmark"); ok || d.HasChange("focus_bookmark") {
		t, err := expandObjectVpnSslWebPortalFocusBookmark(d, v, "focus_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["focus-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download"); ok || d.HasChange("forticlient_download") {
		t, err := expandObjectVpnSslWebPortalForticlientDownload(d, v, "forticlient_download")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download_method"); ok || d.HasChange("forticlient_download_method") {
		t, err := expandObjectVpnSslWebPortalForticlientDownloadMethod(d, v, "forticlient_download_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download-method"] = t
		}
	}

	if v, ok := d.GetOk("heading"); ok || d.HasChange("heading") {
		t, err := expandObjectVpnSslWebPortalHeading(d, v, "heading")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading"] = t
		}
	}

	if v, ok := d.GetOk("hide_sso_credential"); ok || d.HasChange("hide_sso_credential") {
		t, err := expandObjectVpnSslWebPortalHideSsoCredential(d, v, "hide_sso_credential")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hide-sso-credential"] = t
		}
	}

	if v, ok := d.GetOk("host_check"); ok || d.HasChange("host_check") {
		t, err := expandObjectVpnSslWebPortalHostCheck(d, v, "host_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check"] = t
		}
	}

	if v, ok := d.GetOk("host_check_interval"); ok || d.HasChange("host_check_interval") {
		t, err := expandObjectVpnSslWebPortalHostCheckInterval(d, v, "host_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("host_check_policy"); ok || d.HasChange("host_check_policy") {
		t, err := expandObjectVpnSslWebPortalHostCheckPolicy(d, v, "host_check_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check-policy"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok || d.HasChange("ip_mode") {
		t, err := expandObjectVpnSslWebPortalIpMode(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip_pools"); ok || d.HasChange("ip_pools") {
		t, err := expandObjectVpnSslWebPortalIpPools(d, v, "ip_pools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-pools"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok || d.HasChange("ipv6_dns_server1") {
		t, err := expandObjectVpnSslWebPortalIpv6DnsServer1(d, v, "ipv6_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok || d.HasChange("ipv6_dns_server2") {
		t, err := expandObjectVpnSslWebPortalIpv6DnsServer2(d, v, "ipv6_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exclusive_routing"); ok || d.HasChange("ipv6_exclusive_routing") {
		t, err := expandObjectVpnSslWebPortalIpv6ExclusiveRouting(d, v, "ipv6_exclusive_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exclusive-routing"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_pools"); ok || d.HasChange("ipv6_pools") {
		t, err := expandObjectVpnSslWebPortalIpv6Pools(d, v, "ipv6_pools")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-pools"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_service_restriction"); ok || d.HasChange("ipv6_service_restriction") {
		t, err := expandObjectVpnSslWebPortalIpv6ServiceRestriction(d, v, "ipv6_service_restriction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-service-restriction"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling"); ok || d.HasChange("ipv6_split_tunneling") {
		t, err := expandObjectVpnSslWebPortalIpv6SplitTunneling(d, v, "ipv6_split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling_routing_address"); ok || d.HasChange("ipv6_split_tunneling_routing_address") {
		t, err := expandObjectVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d, v, "ipv6_split_tunneling_routing_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling-routing-address"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling_routing_negate"); ok || d.HasChange("ipv6_split_tunneling_routing_negate") {
		t, err := expandObjectVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(d, v, "ipv6_split_tunneling_routing_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling-routing-negate"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_tunnel_mode"); ok || d.HasChange("ipv6_tunnel_mode") {
		t, err := expandObjectVpnSslWebPortalIpv6TunnelMode(d, v, "ipv6_tunnel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-tunnel-mode"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server1"); ok || d.HasChange("ipv6_wins_server1") {
		t, err := expandObjectVpnSslWebPortalIpv6WinsServer1(d, v, "ipv6_wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server2"); ok || d.HasChange("ipv6_wins_server2") {
		t, err := expandObjectVpnSslWebPortalIpv6WinsServer2(d, v, "ipv6_wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("keep_alive"); ok || d.HasChange("keep_alive") {
		t, err := expandObjectVpnSslWebPortalKeepAlive(d, v, "keep_alive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("landing_page"); ok || d.HasChange("landing_page") {
		t, err := expandObjectVpnSslWebPortalLandingPage(d, v, "landing_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["landing-page"] = t
		}
	}

	if v, ok := d.GetOk("landing_page_mode"); ok || d.HasChange("landing_page_mode") {
		t, err := expandObjectVpnSslWebPortalLandingPageMode(d, v, "landing_page_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["landing-page-mode"] = t
		}
	}

	if v, ok := d.GetOk("limit_user_logins"); ok || d.HasChange("limit_user_logins") {
		t, err := expandObjectVpnSslWebPortalLimitUserLogins(d, v, "limit_user_logins")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit-user-logins"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_action"); ok || d.HasChange("mac_addr_action") {
		t, err := expandObjectVpnSslWebPortalMacAddrAction(d, v, "mac_addr_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-action"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_check"); ok || d.HasChange("mac_addr_check") {
		t, err := expandObjectVpnSslWebPortalMacAddrCheck(d, v, "mac_addr_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-check"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_check_rule"); ok || d.HasChange("mac_addr_check_rule") {
		t, err := expandObjectVpnSslWebPortalMacAddrCheckRule(d, v, "mac_addr_check_rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-check-rule"] = t
		}
	}

	if v, ok := d.GetOk("macos_forticlient_download_url"); ok || d.HasChange("macos_forticlient_download_url") {
		t, err := expandObjectVpnSslWebPortalMacosForticlientDownloadUrl(d, v, "macos_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macos-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnSslWebPortalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os_check"); ok || d.HasChange("os_check") {
		t, err := expandObjectVpnSslWebPortalOsCheck(d, v, "os_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-check"] = t
		}
	}

	if v, ok := d.GetOk("os_check_list"); ok || d.HasChange("os_check_list") {
		t, err := expandObjectVpnSslWebPortalOsCheckList(d, v, "os_check_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-check-list"] = t
		}
	}

	if v, ok := d.GetOk("prefer_ipv6_dns"); ok || d.HasChange("prefer_ipv6_dns") {
		t, err := expandObjectVpnSslWebPortalPreferIpv6Dns(d, v, "prefer_ipv6_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefer-ipv6-dns"] = t
		}
	}

	if v, ok := d.GetOk("redir_url"); ok || d.HasChange("redir_url") {
		t, err := expandObjectVpnSslWebPortalRedirUrl(d, v, "redir_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redir-url"] = t
		}
	}

	if v, ok := d.GetOk("rewrite_ip_uri_ui"); ok || d.HasChange("rewrite_ip_uri_ui") {
		t, err := expandObjectVpnSslWebPortalRewriteIpUriUi(d, v, "rewrite_ip_uri_ui")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rewrite-ip-uri-ui"] = t
		}
	}

	if v, ok := d.GetOk("save_password"); ok || d.HasChange("save_password") {
		t, err := expandObjectVpnSslWebPortalSavePassword(d, v, "save_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-password"] = t
		}
	}

	if v, ok := d.GetOk("service_restriction"); ok || d.HasChange("service_restriction") {
		t, err := expandObjectVpnSslWebPortalServiceRestriction(d, v, "service_restriction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-restriction"] = t
		}
	}

	if v, ok := d.GetOk("skip_check_for_unsupported_browser"); ok || d.HasChange("skip_check_for_unsupported_browser") {
		t, err := expandObjectVpnSslWebPortalSkipCheckForUnsupportedBrowser(d, v, "skip_check_for_unsupported_browser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-check-for-unsupported-browser"] = t
		}
	}

	if v, ok := d.GetOk("skip_check_for_browser"); ok || d.HasChange("skip_check_for_browser") {
		t, err := expandObjectVpnSslWebPortalSkipCheckForBrowser(d, v, "skip_check_for_browser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-check-for-browser"] = t
		}
	}

	if v, ok := d.GetOk("skip_check_for_unsupported_os"); ok || d.HasChange("skip_check_for_unsupported_os") {
		t, err := expandObjectVpnSslWebPortalSkipCheckForUnsupportedOs(d, v, "skip_check_for_unsupported_os")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-check-for-unsupported-os"] = t
		}
	}

	if v, ok := d.GetOk("smb_max_version"); ok || d.HasChange("smb_max_version") {
		t, err := expandObjectVpnSslWebPortalSmbMaxVersion(d, v, "smb_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-max-version"] = t
		}
	}

	if v, ok := d.GetOk("smb_min_version"); ok || d.HasChange("smb_min_version") {
		t, err := expandObjectVpnSslWebPortalSmbMinVersion(d, v, "smb_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-min-version"] = t
		}
	}

	if v, ok := d.GetOk("smb_ntlmv1_auth"); ok || d.HasChange("smb_ntlmv1_auth") {
		t, err := expandObjectVpnSslWebPortalSmbNtlmv1Auth(d, v, "smb_ntlmv1_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-ntlmv1-auth"] = t
		}
	}

	if v, ok := d.GetOk("smbv1"); ok || d.HasChange("smbv1") {
		t, err := expandObjectVpnSslWebPortalSmbv1(d, v, "smbv1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smbv1"] = t
		}
	}

	if v, ok := d.GetOk("split_dns"); ok || d.HasChange("split_dns") {
		t, err := expandObjectVpnSslWebPortalSplitDns(d, v, "split_dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-dns"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling"); ok || d.HasChange("split_tunneling") {
		t, err := expandObjectVpnSslWebPortalSplitTunneling(d, v, "split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_routing_address"); ok || d.HasChange("split_tunneling_routing_address") {
		t, err := expandObjectVpnSslWebPortalSplitTunnelingRoutingAddress(d, v, "split_tunneling_routing_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-routing-address"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_routing_negate"); ok || d.HasChange("split_tunneling_routing_negate") {
		t, err := expandObjectVpnSslWebPortalSplitTunnelingRoutingNegate(d, v, "split_tunneling_routing_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-routing-negate"] = t
		}
	}

	if v, ok := d.GetOk("theme"); ok || d.HasChange("theme") {
		t, err := expandObjectVpnSslWebPortalTheme(d, v, "theme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["theme"] = t
		}
	}

	if v, ok := d.GetOk("transform_backward_slashes"); ok || d.HasChange("transform_backward_slashes") {
		t, err := expandObjectVpnSslWebPortalTransformBackwardSlashes(d, v, "transform_backward_slashes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transform-backward-slashes"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok || d.HasChange("tunnel_mode") {
		t, err := expandObjectVpnSslWebPortalTunnelMode(d, v, "tunnel_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mode"] = t
		}
	}

	if v, ok := d.GetOk("use_sdwan"); ok || d.HasChange("use_sdwan") {
		t, err := expandObjectVpnSslWebPortalUseSdwan(d, v, "use_sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-sdwan"] = t
		}
	}

	if v, ok := d.GetOk("user_bookmark"); ok || d.HasChange("user_bookmark") {
		t, err := expandObjectVpnSslWebPortalUserBookmark(d, v, "user_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("user_group_bookmark"); ok || d.HasChange("user_group_bookmark") {
		t, err := expandObjectVpnSslWebPortalUserGroupBookmark(d, v, "user_group_bookmark")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("web_mode"); ok || d.HasChange("web_mode") {
		t, err := expandObjectVpnSslWebPortalWebMode(d, v, "web_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-mode"] = t
		}
	}

	if v, ok := d.GetOk("windows_forticlient_download_url"); ok || d.HasChange("windows_forticlient_download_url") {
		t, err := expandObjectVpnSslWebPortalWindowsForticlientDownloadUrl(d, v, "windows_forticlient_download_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["windows-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok || d.HasChange("wins_server1") {
		t, err := expandObjectVpnSslWebPortalWinsServer1(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok || d.HasChange("wins_server2") {
		t, err := expandObjectVpnSslWebPortalWinsServer2(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	return &obj, nil
}
