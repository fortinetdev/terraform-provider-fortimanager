// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure local users.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserLocalCreate,
		Read:   resourceObjectUserLocalRead,
		Update: resourceObjectUserLocalUpdate,
		Delete: resourceObjectUserLocalDelete,

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
			"auth_concurrent_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_concurrent_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"auth_concurrent_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_concurrent_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"authtimeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"email_to": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fabric_force_sync": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fabric_object": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fabric_object_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortitoken": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"history0": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history1": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history10": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history11": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history12": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history13": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history14": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history15": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history16": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history17": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history18": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history19": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history2": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history3": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history4": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history5": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history6": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history7": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history8": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"history9": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ldap_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"passwd": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"passwd_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"passwd_time": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ppk_identity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ppk_secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"qkd_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"radius_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"saml_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sms_custom_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sms_phone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sms_provider": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sms_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tacacs_server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"two_factor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"two_factor_authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"two_factor_notification": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"username_case_insensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"username_case_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"username_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"workstation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_force_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortitoken": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"history0": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history1": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history10": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history11": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history12": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history13": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history14": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history15": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history16": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history17": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history18": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history19": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history2": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history3": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history4": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history5": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history6": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history7": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history8": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history9": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"passwd_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ppk_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ppk_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"qkd_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"radius_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tacacs_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_insensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"workstation": &schema.Schema{
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

func resourceObjectUserLocalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserLocal(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserLocal resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectUserLocal(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectUserLocal(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectUserLocal resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectUserLocal(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectUserLocal resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserLocalRead(d, m)
}

func resourceObjectUserLocalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectUserLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLocal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserLocal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserLocalRead(d, m)
}

func resourceObjectUserLocalDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserLocal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserLocalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectUserLocal(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectUserLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserLocal resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserLocalAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenObjectUserLocalDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_override"
		if _, ok := i["auth-concurrent-override"]; ok {
			v := flattenObjectUserLocalDynamicMappingAuthConcurrentOverride(i["auth-concurrent-override"], d, pre_append)
			tmp["auth_concurrent_override"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-AuthConcurrentOverride")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_value"
		if _, ok := i["auth-concurrent-value"]; ok {
			v := flattenObjectUserLocalDynamicMappingAuthConcurrentValue(i["auth-concurrent-value"], d, pre_append)
			tmp["auth_concurrent_value"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-AuthConcurrentValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authtimeout"
		if _, ok := i["authtimeout"]; ok {
			v := flattenObjectUserLocalDynamicMappingAuthtimeout(i["authtimeout"], d, pre_append)
			tmp["authtimeout"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-Authtimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email_to"
		if _, ok := i["email-to"]; ok {
			v := flattenObjectUserLocalDynamicMappingEmailTo(i["email-to"], d, pre_append)
			tmp["email_to"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-EmailTo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_force_sync"
		if _, ok := i["fabric-force-sync"]; ok {
			v := flattenObjectUserLocalDynamicMappingFabricForceSync(i["fabric-force-sync"], d, pre_append)
			tmp["fabric_force_sync"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-FabricForceSync")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := i["fabric-object"]; ok {
			v := flattenObjectUserLocalDynamicMappingFabricObject(i["fabric-object"], d, pre_append)
			tmp["fabric_object"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-FabricObject")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object_source"
		if _, ok := i["fabric-object-source"]; ok {
			v := flattenObjectUserLocalDynamicMappingFabricObjectSource(i["fabric-object-source"], d, pre_append)
			tmp["fabric_object_source"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-FabricObjectSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortitoken"
		if _, ok := i["fortitoken"]; ok {
			v := flattenObjectUserLocalDynamicMappingFortitoken(i["fortitoken"], d, pre_append)
			tmp["fortitoken"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-Fortitoken")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserLocalDynamicMappingId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_server"
		if _, ok := i["ldap-server"]; ok {
			v := flattenObjectUserLocalDynamicMappingLdapServer(i["ldap-server"], d, pre_append)
			tmp["ldap_server"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-LdapServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_policy"
		if _, ok := i["passwd-policy"]; ok {
			v := flattenObjectUserLocalDynamicMappingPasswdPolicy(i["passwd-policy"], d, pre_append)
			tmp["passwd_policy"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-PasswdPolicy")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_time"
		if _, ok := i["passwd-time"]; ok {
			v := flattenObjectUserLocalDynamicMappingPasswdTime(i["passwd-time"], d, pre_append)
			tmp["passwd_time"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-PasswdTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ppk_identity"
		if _, ok := i["ppk-identity"]; ok {
			v := flattenObjectUserLocalDynamicMappingPpkIdentity(i["ppk-identity"], d, pre_append)
			tmp["ppk_identity"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-PpkIdentity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qkd_profile"
		if _, ok := i["qkd-profile"]; ok {
			v := flattenObjectUserLocalDynamicMappingQkdProfile(i["qkd-profile"], d, pre_append)
			tmp["qkd_profile"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-QkdProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_server"
		if _, ok := i["radius-server"]; ok {
			v := flattenObjectUserLocalDynamicMappingRadiusServer(i["radius-server"], d, pre_append)
			tmp["radius_server"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-RadiusServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {
			v := flattenObjectUserLocalDynamicMappingSamlServer(i["saml-server"], d, pre_append)
			tmp["saml_server"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-SamlServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_custom_server"
		if _, ok := i["sms-custom-server"]; ok {
			v := flattenObjectUserLocalDynamicMappingSmsCustomServer(i["sms-custom-server"], d, pre_append)
			tmp["sms_custom_server"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-SmsCustomServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_phone"
		if _, ok := i["sms-phone"]; ok {
			v := flattenObjectUserLocalDynamicMappingSmsPhone(i["sms-phone"], d, pre_append)
			tmp["sms_phone"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-SmsPhone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_provider"
		if _, ok := i["sms-provider"]; ok {
			v := flattenObjectUserLocalDynamicMappingSmsProvider(i["sms-provider"], d, pre_append)
			tmp["sms_provider"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-SmsProvider")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_server"
		if _, ok := i["sms-server"]; ok {
			v := flattenObjectUserLocalDynamicMappingSmsServer(i["sms-server"], d, pre_append)
			tmp["sms_server"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-SmsServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectUserLocalDynamicMappingStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tacacs_server"
		if _, ok := i["tacacs+-server"]; ok {
			v := flattenObjectUserLocalDynamicMappingTacacsServer(i["tacacs+-server"], d, pre_append)
			tmp["tacacs_server"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-TacacsServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor"
		if _, ok := i["two-factor"]; ok {
			v := flattenObjectUserLocalDynamicMappingTwoFactor(i["two-factor"], d, pre_append)
			tmp["two_factor"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-TwoFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_authentication"
		if _, ok := i["two-factor-authentication"]; ok {
			v := flattenObjectUserLocalDynamicMappingTwoFactorAuthentication(i["two-factor-authentication"], d, pre_append)
			tmp["two_factor_authentication"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-TwoFactorAuthentication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_notification"
		if _, ok := i["two-factor-notification"]; ok {
			v := flattenObjectUserLocalDynamicMappingTwoFactorNotification(i["two-factor-notification"], d, pre_append)
			tmp["two_factor_notification"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-TwoFactorNotification")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectUserLocalDynamicMappingType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_case_insensitivity"
		if _, ok := i["username-case-insensitivity"]; ok {
			v := flattenObjectUserLocalDynamicMappingUsernameCaseInsensitivity(i["username-case-insensitivity"], d, pre_append)
			tmp["username_case_insensitivity"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-UsernameCaseInsensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_case_sensitivity"
		if _, ok := i["username-case-sensitivity"]; ok {
			v := flattenObjectUserLocalDynamicMappingUsernameCaseSensitivity(i["username-case-sensitivity"], d, pre_append)
			tmp["username_case_sensitivity"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-UsernameCaseSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_sensitivity"
		if _, ok := i["username-sensitivity"]; ok {
			v := flattenObjectUserLocalDynamicMappingUsernameSensitivity(i["username-sensitivity"], d, pre_append)
			tmp["username_sensitivity"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-UsernameSensitivity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := i["uuid"]; ok {
			v := flattenObjectUserLocalDynamicMappingUuid(i["uuid"], d, pre_append)
			tmp["uuid"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-Uuid")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "workstation"
		if _, ok := i["workstation"]; ok {
			v := flattenObjectUserLocalDynamicMappingWorkstation(i["workstation"], d, pre_append)
			tmp["workstation"] = fortiAPISubPartPatch(v, "ObjectUserLocal-DynamicMapping-Workstation")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserLocalDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserLocalDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserLocalDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserLocalDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserLocalDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserLocalDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingAuthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingEmailTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingFabricForceSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingFabricObjectSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingFortitoken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingPasswdPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingPasswdTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingPpkIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingQkdProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingSamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingSmsPhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingSmsProvider(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingTacacsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingUsernameCaseInsensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingUsernameCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingUsernameSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingWorkstation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalEmailTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalFabricForceSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalFabricObjectSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalFortitoken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLocalId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLocalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalPasswdPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLocalPpkIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalQkdProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLocalRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLocalSamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLocalSmsPhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalTacacsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectUserLocalTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalUsernameCaseInsensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalUsernameCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalUsernameSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalWorkstation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("auth_concurrent_override", flattenObjectUserLocalAuthConcurrentOverride(o["auth-concurrent-override"], d, "auth_concurrent_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-override"], "ObjectUserLocal-AuthConcurrentOverride"); ok {
			if err = d.Set("auth_concurrent_override", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenObjectUserLocalAuthConcurrentValue(o["auth-concurrent-value"], d, "auth_concurrent_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-value"], "ObjectUserLocal-AuthConcurrentValue"); ok {
			if err = d.Set("auth_concurrent_value", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenObjectUserLocalAuthtimeout(o["authtimeout"], d, "authtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["authtimeout"], "ObjectUserLocal-Authtimeout"); ok {
			if err = d.Set("authtimeout", vv); err != nil {
				return fmt.Errorf("Error reading authtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenObjectUserLocalDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserLocal-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenObjectUserLocalDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "ObjectUserLocal-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("email_to", flattenObjectUserLocalEmailTo(o["email-to"], d, "email_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-to"], "ObjectUserLocal-EmailTo"); ok {
			if err = d.Set("email_to", vv); err != nil {
				return fmt.Errorf("Error reading email_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("fabric_force_sync", flattenObjectUserLocalFabricForceSync(o["fabric-force-sync"], d, "fabric_force_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-force-sync"], "ObjectUserLocal-FabricForceSync"); ok {
			if err = d.Set("fabric_force_sync", vv); err != nil {
				return fmt.Errorf("Error reading fabric_force_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_force_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectUserLocalFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectUserLocal-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fabric_object_source", flattenObjectUserLocalFabricObjectSource(o["fabric-object-source"], d, "fabric_object_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object-source"], "ObjectUserLocal-FabricObjectSource"); ok {
			if err = d.Set("fabric_object_source", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object_source: %v", err)
		}
	}

	if err = d.Set("fortitoken", flattenObjectUserLocalFortitoken(o["fortitoken"], d, "fortitoken")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken"], "ObjectUserLocal-Fortitoken"); ok {
			if err = d.Set("fortitoken", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserLocalId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserLocal-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserLocalLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserLocal-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserLocalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserLocal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("passwd_policy", flattenObjectUserLocalPasswdPolicy(o["passwd-policy"], d, "passwd_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["passwd-policy"], "ObjectUserLocal-PasswdPolicy"); ok {
			if err = d.Set("passwd_policy", vv); err != nil {
				return fmt.Errorf("Error reading passwd_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passwd_policy: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenObjectUserLocalPpkIdentity(o["ppk-identity"], d, "ppk_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppk-identity"], "ObjectUserLocal-PpkIdentity"); ok {
			if err = d.Set("ppk_identity", vv); err != nil {
				return fmt.Errorf("Error reading ppk_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("qkd_profile", flattenObjectUserLocalQkdProfile(o["qkd-profile"], d, "qkd_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qkd-profile"], "ObjectUserLocal-QkdProfile"); ok {
			if err = d.Set("qkd_profile", vv); err != nil {
				return fmt.Errorf("Error reading qkd_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qkd_profile: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenObjectUserLocalRadiusServer(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "ObjectUserLocal-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenObjectUserLocalSamlServer(o["saml-server"], d, "saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-server"], "ObjectUserLocal-SamlServer"); ok {
			if err = d.Set("saml_server", vv); err != nil {
				return fmt.Errorf("Error reading saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenObjectUserLocalSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-custom-server"], "ObjectUserLocal-SmsCustomServer"); ok {
			if err = d.Set("sms_custom_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_custom_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_phone", flattenObjectUserLocalSmsPhone(o["sms-phone"], d, "sms_phone")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-phone"], "ObjectUserLocal-SmsPhone"); ok {
			if err = d.Set("sms_phone", vv); err != nil {
				return fmt.Errorf("Error reading sms_phone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_phone: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenObjectUserLocalSmsServer(o["sms-server"], d, "sms_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-server"], "ObjectUserLocal-SmsServer"); ok {
			if err = d.Set("sms_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserLocalStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserLocal-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tacacs_server", flattenObjectUserLocalTacacsServer(o["tacacs+-server"], d, "tacacs_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tacacs+-server"], "ObjectUserLocal-TacacsServer"); ok {
			if err = d.Set("tacacs_server", vv); err != nil {
				return fmt.Errorf("Error reading tacacs_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tacacs_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenObjectUserLocalTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "ObjectUserLocal-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenObjectUserLocalTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-authentication"], "ObjectUserLocal-TwoFactorAuthentication"); ok {
			if err = d.Set("two_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenObjectUserLocalTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-notification"], "ObjectUserLocal-TwoFactorNotification"); ok {
			if err = d.Set("two_factor_notification", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserLocalType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserLocal-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username_case_insensitivity", flattenObjectUserLocalUsernameCaseInsensitivity(o["username-case-insensitivity"], d, "username_case_insensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-insensitivity"], "ObjectUserLocal-UsernameCaseInsensitivity"); ok {
			if err = d.Set("username_case_insensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
		}
	}

	if err = d.Set("username_case_sensitivity", flattenObjectUserLocalUsernameCaseSensitivity(o["username-case-sensitivity"], d, "username_case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-sensitivity"], "ObjectUserLocal-UsernameCaseSensitivity"); ok {
			if err = d.Set("username_case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
		}
	}

	if err = d.Set("username_sensitivity", flattenObjectUserLocalUsernameSensitivity(o["username-sensitivity"], d, "username_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-sensitivity"], "ObjectUserLocal-UsernameSensitivity"); ok {
			if err = d.Set("username_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_sensitivity: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectUserLocalUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectUserLocal-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("workstation", flattenObjectUserLocalWorkstation(o["workstation"], d, "workstation")); err != nil {
		if vv, ok := fortiAPIPatch(o["workstation"], "ObjectUserLocal-Workstation"); ok {
			if err = d.Set("workstation", vv); err != nil {
				return fmt.Errorf("Error reading workstation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workstation: %v", err)
		}
	}

	return nil
}

func flattenObjectUserLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserLocalAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalAuthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandObjectUserLocalDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-concurrent-override"], _ = expandObjectUserLocalDynamicMappingAuthConcurrentOverride(d, i["auth_concurrent_override"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_concurrent_value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-concurrent-value"], _ = expandObjectUserLocalDynamicMappingAuthConcurrentValue(d, i["auth_concurrent_value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authtimeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authtimeout"], _ = expandObjectUserLocalDynamicMappingAuthtimeout(d, i["authtimeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email_to"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["email-to"], _ = expandObjectUserLocalDynamicMappingEmailTo(d, i["email_to"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_force_sync"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-force-sync"], _ = expandObjectUserLocalDynamicMappingFabricForceSync(d, i["fabric_force_sync"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-object"], _ = expandObjectUserLocalDynamicMappingFabricObject(d, i["fabric_object"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_object_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fabric-object-source"], _ = expandObjectUserLocalDynamicMappingFabricObjectSource(d, i["fabric_object_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortitoken"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortitoken"], _ = expandObjectUserLocalDynamicMappingFortitoken(d, i["fortitoken"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history0"], _ = expandObjectUserLocalDynamicMappingHistory0(d, i["history0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history1"], _ = expandObjectUserLocalDynamicMappingHistory1(d, i["history1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history10"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history10"], _ = expandObjectUserLocalDynamicMappingHistory10(d, i["history10"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history11"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history11"], _ = expandObjectUserLocalDynamicMappingHistory11(d, i["history11"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history12"], _ = expandObjectUserLocalDynamicMappingHistory12(d, i["history12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history13"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history13"], _ = expandObjectUserLocalDynamicMappingHistory13(d, i["history13"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history14"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history14"], _ = expandObjectUserLocalDynamicMappingHistory14(d, i["history14"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history15"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history15"], _ = expandObjectUserLocalDynamicMappingHistory15(d, i["history15"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history16"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history16"], _ = expandObjectUserLocalDynamicMappingHistory16(d, i["history16"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history17"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history17"], _ = expandObjectUserLocalDynamicMappingHistory17(d, i["history17"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history18"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history18"], _ = expandObjectUserLocalDynamicMappingHistory18(d, i["history18"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history19"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history19"], _ = expandObjectUserLocalDynamicMappingHistory19(d, i["history19"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history2"], _ = expandObjectUserLocalDynamicMappingHistory2(d, i["history2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history3"], _ = expandObjectUserLocalDynamicMappingHistory3(d, i["history3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history4"], _ = expandObjectUserLocalDynamicMappingHistory4(d, i["history4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history5"], _ = expandObjectUserLocalDynamicMappingHistory5(d, i["history5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history6"], _ = expandObjectUserLocalDynamicMappingHistory6(d, i["history6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history7"], _ = expandObjectUserLocalDynamicMappingHistory7(d, i["history7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history8"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history8"], _ = expandObjectUserLocalDynamicMappingHistory8(d, i["history8"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "history9"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["history9"], _ = expandObjectUserLocalDynamicMappingHistory9(d, i["history9"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserLocalDynamicMappingId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldap_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ldap-server"], _ = expandObjectUserLocalDynamicMappingLdapServer(d, i["ldap_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passwd"], _ = expandObjectUserLocalDynamicMappingPasswd(d, i["passwd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_policy"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passwd-policy"], _ = expandObjectUserLocalDynamicMappingPasswdPolicy(d, i["passwd_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passwd-time"], _ = expandObjectUserLocalDynamicMappingPasswdTime(d, i["passwd_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ppk_identity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ppk-identity"], _ = expandObjectUserLocalDynamicMappingPpkIdentity(d, i["ppk_identity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ppk_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ppk-secret"], _ = expandObjectUserLocalDynamicMappingPpkSecret(d, i["ppk_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qkd_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["qkd-profile"], _ = expandObjectUserLocalDynamicMappingQkdProfile(d, i["qkd_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-server"], _ = expandObjectUserLocalDynamicMappingRadiusServer(d, i["radius_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["saml-server"], _ = expandObjectUserLocalDynamicMappingSamlServer(d, i["saml_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_custom_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sms-custom-server"], _ = expandObjectUserLocalDynamicMappingSmsCustomServer(d, i["sms_custom_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_phone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sms-phone"], _ = expandObjectUserLocalDynamicMappingSmsPhone(d, i["sms_phone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_provider"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sms-provider"], _ = expandObjectUserLocalDynamicMappingSmsProvider(d, i["sms_provider"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sms_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sms-server"], _ = expandObjectUserLocalDynamicMappingSmsServer(d, i["sms_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectUserLocalDynamicMappingStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tacacs_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tacacs+-server"], _ = expandObjectUserLocalDynamicMappingTacacsServer(d, i["tacacs_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor"], _ = expandObjectUserLocalDynamicMappingTwoFactor(d, i["two_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_authentication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor-authentication"], _ = expandObjectUserLocalDynamicMappingTwoFactorAuthentication(d, i["two_factor_authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "two_factor_notification"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["two-factor-notification"], _ = expandObjectUserLocalDynamicMappingTwoFactorNotification(d, i["two_factor_notification"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectUserLocalDynamicMappingType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_case_insensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["username-case-insensitivity"], _ = expandObjectUserLocalDynamicMappingUsernameCaseInsensitivity(d, i["username_case_insensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["username-case-sensitivity"], _ = expandObjectUserLocalDynamicMappingUsernameCaseSensitivity(d, i["username_case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_sensitivity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["username-sensitivity"], _ = expandObjectUserLocalDynamicMappingUsernameSensitivity(d, i["username_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uuid"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["uuid"], _ = expandObjectUserLocalDynamicMappingUuid(d, i["uuid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "workstation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["workstation"], _ = expandObjectUserLocalDynamicMappingWorkstation(d, i["workstation"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserLocalDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserLocalDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserLocalDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserLocalDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingAuthtimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingEmailTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingFabricForceSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingFabricObjectSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingFortitoken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory11(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory13(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory14(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory15(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory17(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory18(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory19(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingPasswdPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingPasswdTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingPpkIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingPpkSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingQkdProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingSamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingSmsPhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingSmsProvider(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingTacacsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingUsernameCaseInsensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingUsernameCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingUsernameSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingWorkstation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalEmailTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalFabricForceSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalFabricObjectSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalFortitoken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLocalHistory0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory11(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory13(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory14(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory15(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory16(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory17(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory18(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory19(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalHistory9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLocalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalPasswdPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLocalPpkIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalPpkSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalQkdProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLocalRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLocalSamlServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLocalSmsPhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalTacacsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectUserLocalTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalUsernameCaseInsensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalUsernameCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalUsernameSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalWorkstation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_concurrent_override"); ok || d.HasChange("auth_concurrent_override") {
		t, err := expandObjectUserLocalAuthConcurrentOverride(d, v, "auth_concurrent_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_value"); ok || d.HasChange("auth_concurrent_value") {
		t, err := expandObjectUserLocalAuthConcurrentValue(d, v, "auth_concurrent_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("authtimeout"); ok || d.HasChange("authtimeout") {
		t, err := expandObjectUserLocalAuthtimeout(d, v, "authtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandObjectUserLocalDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok || d.HasChange("email_to") {
		t, err := expandObjectUserLocalEmailTo(d, v, "email_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("fabric_force_sync"); ok || d.HasChange("fabric_force_sync") {
		t, err := expandObjectUserLocalFabricForceSync(d, v, "fabric_force_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-force-sync"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectUserLocalFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_source"); ok || d.HasChange("fabric_object_source") {
		t, err := expandObjectUserLocalFabricObjectSource(d, v, "fabric_object_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-source"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken"); ok || d.HasChange("fortitoken") {
		t, err := expandObjectUserLocalFortitoken(d, v, "fortitoken")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken"] = t
		}
	}

	if v, ok := d.GetOk("history0"); ok || d.HasChange("history0") {
		t, err := expandObjectUserLocalHistory0(d, v, "history0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history0"] = t
		}
	}

	if v, ok := d.GetOk("history1"); ok || d.HasChange("history1") {
		t, err := expandObjectUserLocalHistory1(d, v, "history1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history1"] = t
		}
	}

	if v, ok := d.GetOk("history10"); ok || d.HasChange("history10") {
		t, err := expandObjectUserLocalHistory10(d, v, "history10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history10"] = t
		}
	}

	if v, ok := d.GetOk("history11"); ok || d.HasChange("history11") {
		t, err := expandObjectUserLocalHistory11(d, v, "history11")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history11"] = t
		}
	}

	if v, ok := d.GetOk("history12"); ok || d.HasChange("history12") {
		t, err := expandObjectUserLocalHistory12(d, v, "history12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history12"] = t
		}
	}

	if v, ok := d.GetOk("history13"); ok || d.HasChange("history13") {
		t, err := expandObjectUserLocalHistory13(d, v, "history13")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history13"] = t
		}
	}

	if v, ok := d.GetOk("history14"); ok || d.HasChange("history14") {
		t, err := expandObjectUserLocalHistory14(d, v, "history14")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history14"] = t
		}
	}

	if v, ok := d.GetOk("history15"); ok || d.HasChange("history15") {
		t, err := expandObjectUserLocalHistory15(d, v, "history15")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history15"] = t
		}
	}

	if v, ok := d.GetOk("history16"); ok || d.HasChange("history16") {
		t, err := expandObjectUserLocalHistory16(d, v, "history16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history16"] = t
		}
	}

	if v, ok := d.GetOk("history17"); ok || d.HasChange("history17") {
		t, err := expandObjectUserLocalHistory17(d, v, "history17")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history17"] = t
		}
	}

	if v, ok := d.GetOk("history18"); ok || d.HasChange("history18") {
		t, err := expandObjectUserLocalHistory18(d, v, "history18")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history18"] = t
		}
	}

	if v, ok := d.GetOk("history19"); ok || d.HasChange("history19") {
		t, err := expandObjectUserLocalHistory19(d, v, "history19")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history19"] = t
		}
	}

	if v, ok := d.GetOk("history2"); ok || d.HasChange("history2") {
		t, err := expandObjectUserLocalHistory2(d, v, "history2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history2"] = t
		}
	}

	if v, ok := d.GetOk("history3"); ok || d.HasChange("history3") {
		t, err := expandObjectUserLocalHistory3(d, v, "history3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history3"] = t
		}
	}

	if v, ok := d.GetOk("history4"); ok || d.HasChange("history4") {
		t, err := expandObjectUserLocalHistory4(d, v, "history4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history4"] = t
		}
	}

	if v, ok := d.GetOk("history5"); ok || d.HasChange("history5") {
		t, err := expandObjectUserLocalHistory5(d, v, "history5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history5"] = t
		}
	}

	if v, ok := d.GetOk("history6"); ok || d.HasChange("history6") {
		t, err := expandObjectUserLocalHistory6(d, v, "history6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history6"] = t
		}
	}

	if v, ok := d.GetOk("history7"); ok || d.HasChange("history7") {
		t, err := expandObjectUserLocalHistory7(d, v, "history7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history7"] = t
		}
	}

	if v, ok := d.GetOk("history8"); ok || d.HasChange("history8") {
		t, err := expandObjectUserLocalHistory8(d, v, "history8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history8"] = t
		}
	}

	if v, ok := d.GetOk("history9"); ok || d.HasChange("history9") {
		t, err := expandObjectUserLocalHistory9(d, v, "history9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history9"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserLocalId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandObjectUserLocalLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserLocalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok || d.HasChange("passwd") {
		t, err := expandObjectUserLocalPasswd(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("passwd_policy"); ok || d.HasChange("passwd_policy") {
		t, err := expandObjectUserLocalPasswdPolicy(d, v, "passwd_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd-policy"] = t
		}
	}

	if v, ok := d.GetOk("ppk_identity"); ok || d.HasChange("ppk_identity") {
		t, err := expandObjectUserLocalPpkIdentity(d, v, "ppk_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok || d.HasChange("ppk_secret") {
		t, err := expandObjectUserLocalPpkSecret(d, v, "ppk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	}

	if v, ok := d.GetOk("qkd_profile"); ok || d.HasChange("qkd_profile") {
		t, err := expandObjectUserLocalQkdProfile(d, v, "qkd_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd-profile"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandObjectUserLocalRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("saml_server"); ok || d.HasChange("saml_server") {
		t, err := expandObjectUserLocalSamlServer(d, v, "saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok || d.HasChange("sms_custom_server") {
		t, err := expandObjectUserLocalSmsCustomServer(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_phone"); ok || d.HasChange("sms_phone") {
		t, err := expandObjectUserLocalSmsPhone(d, v, "sms_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-phone"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok || d.HasChange("sms_server") {
		t, err := expandObjectUserLocalSmsServer(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserLocalStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tacacs_server"); ok || d.HasChange("tacacs_server") {
		t, err := expandObjectUserLocalTacacsServer(d, v, "tacacs_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tacacs+-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandObjectUserLocalTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok || d.HasChange("two_factor_authentication") {
		t, err := expandObjectUserLocalTwoFactorAuthentication(d, v, "two_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok || d.HasChange("two_factor_notification") {
		t, err := expandObjectUserLocalTwoFactorNotification(d, v, "two_factor_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserLocalType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username_case_insensitivity"); ok || d.HasChange("username_case_insensitivity") {
		t, err := expandObjectUserLocalUsernameCaseInsensitivity(d, v, "username_case_insensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-insensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitivity"); ok || d.HasChange("username_case_sensitivity") {
		t, err := expandObjectUserLocalUsernameCaseSensitivity(d, v, "username_case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_sensitivity"); ok || d.HasChange("username_sensitivity") {
		t, err := expandObjectUserLocalUsernameSensitivity(d, v, "username_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectUserLocalUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("workstation"); ok || d.HasChange("workstation") {
		t, err := expandObjectUserLocalWorkstation(d, v, "workstation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workstation"] = t
		}
	}

	return &obj, nil
}
