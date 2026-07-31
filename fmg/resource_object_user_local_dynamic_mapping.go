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

func resourceObjectUserLocalDynamicMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserLocalDynamicMappingCreate,
		Read:   resourceObjectUserLocalDynamicMappingRead,
		Update: resourceObjectUserLocalDynamicMappingUpdate,
		Delete: resourceObjectUserLocalDynamicMappingDelete,

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
			"local": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"fosid": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectUserLocalDynamicMappingCreate(d *schema.ResourceData, m interface{}) error {
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

	local := d.Get("local").(string)
	paradict["local"] = local

	obj, err := getObjectObjectUserLocalDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserLocalDynamicMapping resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("_scope")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectUserLocalDynamicMapping(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectUserLocalDynamicMapping(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectUserLocalDynamicMapping resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectUserLocalDynamicMapping(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectUserLocalDynamicMapping resource: %v", err)
		}

	}

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserLocalDynamicMappingRead(d, m)
}

func resourceObjectUserLocalDynamicMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	local := d.Get("local").(string)
	paradict["local"] = local

	obj, err := getObjectObjectUserLocalDynamicMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLocalDynamicMapping resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectUserLocalDynamicMapping(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserLocalDynamicMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getScopeKey(d, "_scope"))

	return resourceObjectUserLocalDynamicMappingRead(d, m)
}

func resourceObjectUserLocalDynamicMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	local := d.Get("local").(string)
	paradict["local"] = local

	wsParams["adom"] = adomv

	err = c.DeleteObjectUserLocalDynamicMapping(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserLocalDynamicMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserLocalDynamicMappingRead(d *schema.ResourceData, m interface{}) error {
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

	local := d.Get("local").(string)
	if local == "" {
		local = importOptionChecking(m.(*FortiClient).Cfg, "local")
		if local == "" {
			return fmt.Errorf("Parameter local is missing")
		}
		if err = d.Set("local", local); err != nil {
			return fmt.Errorf("Error set params local: %v", err)
		}
	}
	if mkey, err = checkScopeId(mkey); err != nil {
		return fmt.Errorf("Error set ID : %v", err)
	}
	paradict["local"] = local

	o, err := c.ReadObjectUserLocalDynamicMapping(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectUserLocalDynamicMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserLocalDynamicMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserLocalDynamicMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserLocalDynamicMappingScope2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectUserLocalDynamicMappingScopeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectUserLocalDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectUserLocalDynamicMappingScopeVdom2edl(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectUserLocalDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectUserLocalDynamicMappingScopeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingScopeVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingAuthConcurrentOverride2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingAuthConcurrentValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingAuthtimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingEmailTo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingFabricForceSync2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingFabricObject2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingFabricObjectSource2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingFortitoken2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingLdapServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingPasswdPolicy2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingPasswdTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingPpkIdentity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingQkdProfile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingRadiusServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingSamlServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingSmsCustomServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingSmsPhone2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingSmsProvider2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingSmsServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingTacacsServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserLocalDynamicMappingTwoFactor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingTwoFactorAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingTwoFactorNotification2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingUsernameCaseInsensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingUsernameCaseSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingUsernameSensitivity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingUuid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserLocalDynamicMappingWorkstation2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserLocalDynamicMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("_scope", flattenObjectUserLocalDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserLocalDynamicMapping-Scope"); ok {
				if err = d.Set("_scope", vv); err != nil {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading _scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("_scope"); ok {
			if err = d.Set("_scope", flattenObjectUserLocalDynamicMappingScope2edl(o["_scope"], d, "_scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["_scope"], "ObjectUserLocalDynamicMapping-Scope"); ok {
					if err = d.Set("_scope", vv); err != nil {
						return fmt.Errorf("Error reading _scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading _scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_concurrent_override", flattenObjectUserLocalDynamicMappingAuthConcurrentOverride2edl(o["auth-concurrent-override"], d, "auth_concurrent_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-override"], "ObjectUserLocalDynamicMapping-AuthConcurrentOverride"); ok {
			if err = d.Set("auth_concurrent_override", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenObjectUserLocalDynamicMappingAuthConcurrentValue2edl(o["auth-concurrent-value"], d, "auth_concurrent_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-concurrent-value"], "ObjectUserLocalDynamicMapping-AuthConcurrentValue"); ok {
			if err = d.Set("auth_concurrent_value", vv); err != nil {
				return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenObjectUserLocalDynamicMappingAuthtimeout2edl(o["authtimeout"], d, "authtimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["authtimeout"], "ObjectUserLocalDynamicMapping-Authtimeout"); ok {
			if err = d.Set("authtimeout", vv); err != nil {
				return fmt.Errorf("Error reading authtimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if err = d.Set("email_to", flattenObjectUserLocalDynamicMappingEmailTo2edl(o["email-to"], d, "email_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-to"], "ObjectUserLocalDynamicMapping-EmailTo"); ok {
			if err = d.Set("email_to", vv); err != nil {
				return fmt.Errorf("Error reading email_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("fabric_force_sync", flattenObjectUserLocalDynamicMappingFabricForceSync2edl(o["fabric-force-sync"], d, "fabric_force_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-force-sync"], "ObjectUserLocalDynamicMapping-FabricForceSync"); ok {
			if err = d.Set("fabric_force_sync", vv); err != nil {
				return fmt.Errorf("Error reading fabric_force_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_force_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenObjectUserLocalDynamicMappingFabricObject2edl(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "ObjectUserLocalDynamicMapping-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fabric_object_source", flattenObjectUserLocalDynamicMappingFabricObjectSource2edl(o["fabric-object-source"], d, "fabric_object_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object-source"], "ObjectUserLocalDynamicMapping-FabricObjectSource"); ok {
			if err = d.Set("fabric_object_source", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object_source: %v", err)
		}
	}

	if err = d.Set("fortitoken", flattenObjectUserLocalDynamicMappingFortitoken2edl(o["fortitoken"], d, "fortitoken")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken"], "ObjectUserLocalDynamicMapping-Fortitoken"); ok {
			if err = d.Set("fortitoken", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserLocalDynamicMappingId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserLocalDynamicMapping-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserLocalDynamicMappingLdapServer2edl(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserLocalDynamicMapping-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("passwd_policy", flattenObjectUserLocalDynamicMappingPasswdPolicy2edl(o["passwd-policy"], d, "passwd_policy")); err != nil {
		if vv, ok := fortiAPIPatch(o["passwd-policy"], "ObjectUserLocalDynamicMapping-PasswdPolicy"); ok {
			if err = d.Set("passwd_policy", vv); err != nil {
				return fmt.Errorf("Error reading passwd_policy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passwd_policy: %v", err)
		}
	}

	if err = d.Set("passwd_time", flattenObjectUserLocalDynamicMappingPasswdTime2edl(o["passwd-time"], d, "passwd_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["passwd-time"], "ObjectUserLocalDynamicMapping-PasswdTime"); ok {
			if err = d.Set("passwd_time", vv); err != nil {
				return fmt.Errorf("Error reading passwd_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passwd_time: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenObjectUserLocalDynamicMappingPpkIdentity2edl(o["ppk-identity"], d, "ppk_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["ppk-identity"], "ObjectUserLocalDynamicMapping-PpkIdentity"); ok {
			if err = d.Set("ppk_identity", vv); err != nil {
				return fmt.Errorf("Error reading ppk_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("qkd_profile", flattenObjectUserLocalDynamicMappingQkdProfile2edl(o["qkd-profile"], d, "qkd_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["qkd-profile"], "ObjectUserLocalDynamicMapping-QkdProfile"); ok {
			if err = d.Set("qkd_profile", vv); err != nil {
				return fmt.Errorf("Error reading qkd_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading qkd_profile: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenObjectUserLocalDynamicMappingRadiusServer2edl(o["radius-server"], d, "radius_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-server"], "ObjectUserLocalDynamicMapping-RadiusServer"); ok {
			if err = d.Set("radius_server", vv); err != nil {
				return fmt.Errorf("Error reading radius_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenObjectUserLocalDynamicMappingSamlServer2edl(o["saml-server"], d, "saml_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["saml-server"], "ObjectUserLocalDynamicMapping-SamlServer"); ok {
			if err = d.Set("saml_server", vv); err != nil {
				return fmt.Errorf("Error reading saml_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenObjectUserLocalDynamicMappingSmsCustomServer2edl(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-custom-server"], "ObjectUserLocalDynamicMapping-SmsCustomServer"); ok {
			if err = d.Set("sms_custom_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_custom_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_phone", flattenObjectUserLocalDynamicMappingSmsPhone2edl(o["sms-phone"], d, "sms_phone")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-phone"], "ObjectUserLocalDynamicMapping-SmsPhone"); ok {
			if err = d.Set("sms_phone", vv); err != nil {
				return fmt.Errorf("Error reading sms_phone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_phone: %v", err)
		}
	}

	if err = d.Set("sms_provider", flattenObjectUserLocalDynamicMappingSmsProvider2edl(o["sms-provider"], d, "sms_provider")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-provider"], "ObjectUserLocalDynamicMapping-SmsProvider"); ok {
			if err = d.Set("sms_provider", vv); err != nil {
				return fmt.Errorf("Error reading sms_provider: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_provider: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenObjectUserLocalDynamicMappingSmsServer2edl(o["sms-server"], d, "sms_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-server"], "ObjectUserLocalDynamicMapping-SmsServer"); ok {
			if err = d.Set("sms_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectUserLocalDynamicMappingStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectUserLocalDynamicMapping-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tacacs_server", flattenObjectUserLocalDynamicMappingTacacsServer2edl(o["tacacs+-server"], d, "tacacs_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tacacs+-server"], "ObjectUserLocalDynamicMapping-TacacsServer"); ok {
			if err = d.Set("tacacs_server", vv); err != nil {
				return fmt.Errorf("Error reading tacacs_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tacacs_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenObjectUserLocalDynamicMappingTwoFactor2edl(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "ObjectUserLocalDynamicMapping-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenObjectUserLocalDynamicMappingTwoFactorAuthentication2edl(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-authentication"], "ObjectUserLocalDynamicMapping-TwoFactorAuthentication"); ok {
			if err = d.Set("two_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenObjectUserLocalDynamicMappingTwoFactorNotification2edl(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-notification"], "ObjectUserLocalDynamicMapping-TwoFactorNotification"); ok {
			if err = d.Set("two_factor_notification", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectUserLocalDynamicMappingType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectUserLocalDynamicMapping-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username_case_insensitivity", flattenObjectUserLocalDynamicMappingUsernameCaseInsensitivity2edl(o["username-case-insensitivity"], d, "username_case_insensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-insensitivity"], "ObjectUserLocalDynamicMapping-UsernameCaseInsensitivity"); ok {
			if err = d.Set("username_case_insensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
		}
	}

	if err = d.Set("username_case_sensitivity", flattenObjectUserLocalDynamicMappingUsernameCaseSensitivity2edl(o["username-case-sensitivity"], d, "username_case_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-sensitivity"], "ObjectUserLocalDynamicMapping-UsernameCaseSensitivity"); ok {
			if err = d.Set("username_case_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
		}
	}

	if err = d.Set("username_sensitivity", flattenObjectUserLocalDynamicMappingUsernameSensitivity2edl(o["username-sensitivity"], d, "username_sensitivity")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-sensitivity"], "ObjectUserLocalDynamicMapping-UsernameSensitivity"); ok {
			if err = d.Set("username_sensitivity", vv); err != nil {
				return fmt.Errorf("Error reading username_sensitivity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_sensitivity: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectUserLocalDynamicMappingUuid2edl(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectUserLocalDynamicMapping-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("workstation", flattenObjectUserLocalDynamicMappingWorkstation2edl(o["workstation"], d, "workstation")); err != nil {
		if vv, ok := fortiAPIPatch(o["workstation"], "ObjectUserLocalDynamicMapping-Workstation"); ok {
			if err = d.Set("workstation", vv); err != nil {
				return fmt.Errorf("Error reading workstation: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading workstation: %v", err)
		}
	}

	return nil
}

func flattenObjectUserLocalDynamicMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserLocalDynamicMappingScope2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectUserLocalDynamicMappingScopeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectUserLocalDynamicMappingScopeVdom2edl(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectUserLocalDynamicMappingScopeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingScopeVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingAuthConcurrentOverride2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingAuthConcurrentValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingAuthtimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingEmailTo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingFabricForceSync2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingFabricObject2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingFabricObjectSource2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingFortitoken2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory02edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory102edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory112edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory122edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory132edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory142edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory152edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory162edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory172edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory182edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory192edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory72edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory82edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingHistory92edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingLdapServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingPasswd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingPasswdPolicy2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingPasswdTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingPpkIdentity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingPpkSecret2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingQkdProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingRadiusServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingSamlServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingSmsCustomServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingSmsPhone2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingSmsProvider2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingSmsServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingTacacsServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserLocalDynamicMappingTwoFactor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingTwoFactorAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingTwoFactorNotification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingUsernameCaseInsensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingUsernameCaseSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingUsernameSensitivity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingUuid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserLocalDynamicMappingWorkstation2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserLocalDynamicMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_scope"); ok || d.HasChange("_scope") {
		t, err := expandObjectUserLocalDynamicMappingScope2edl(d, v, "_scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_scope"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_override"); ok || d.HasChange("auth_concurrent_override") {
		t, err := expandObjectUserLocalDynamicMappingAuthConcurrentOverride2edl(d, v, "auth_concurrent_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_value"); ok || d.HasChange("auth_concurrent_value") {
		t, err := expandObjectUserLocalDynamicMappingAuthConcurrentValue2edl(d, v, "auth_concurrent_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("authtimeout"); ok || d.HasChange("authtimeout") {
		t, err := expandObjectUserLocalDynamicMappingAuthtimeout2edl(d, v, "authtimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok || d.HasChange("email_to") {
		t, err := expandObjectUserLocalDynamicMappingEmailTo2edl(d, v, "email_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("fabric_force_sync"); ok || d.HasChange("fabric_force_sync") {
		t, err := expandObjectUserLocalDynamicMappingFabricForceSync2edl(d, v, "fabric_force_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-force-sync"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandObjectUserLocalDynamicMappingFabricObject2edl(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_source"); ok || d.HasChange("fabric_object_source") {
		t, err := expandObjectUserLocalDynamicMappingFabricObjectSource2edl(d, v, "fabric_object_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-source"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken"); ok || d.HasChange("fortitoken") {
		t, err := expandObjectUserLocalDynamicMappingFortitoken2edl(d, v, "fortitoken")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken"] = t
		}
	}

	if v, ok := d.GetOk("history0"); ok || d.HasChange("history0") {
		t, err := expandObjectUserLocalDynamicMappingHistory02edl(d, v, "history0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history0"] = t
		}
	}

	if v, ok := d.GetOk("history1"); ok || d.HasChange("history1") {
		t, err := expandObjectUserLocalDynamicMappingHistory12edl(d, v, "history1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history1"] = t
		}
	}

	if v, ok := d.GetOk("history10"); ok || d.HasChange("history10") {
		t, err := expandObjectUserLocalDynamicMappingHistory102edl(d, v, "history10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history10"] = t
		}
	}

	if v, ok := d.GetOk("history11"); ok || d.HasChange("history11") {
		t, err := expandObjectUserLocalDynamicMappingHistory112edl(d, v, "history11")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history11"] = t
		}
	}

	if v, ok := d.GetOk("history12"); ok || d.HasChange("history12") {
		t, err := expandObjectUserLocalDynamicMappingHistory122edl(d, v, "history12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history12"] = t
		}
	}

	if v, ok := d.GetOk("history13"); ok || d.HasChange("history13") {
		t, err := expandObjectUserLocalDynamicMappingHistory132edl(d, v, "history13")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history13"] = t
		}
	}

	if v, ok := d.GetOk("history14"); ok || d.HasChange("history14") {
		t, err := expandObjectUserLocalDynamicMappingHistory142edl(d, v, "history14")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history14"] = t
		}
	}

	if v, ok := d.GetOk("history15"); ok || d.HasChange("history15") {
		t, err := expandObjectUserLocalDynamicMappingHistory152edl(d, v, "history15")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history15"] = t
		}
	}

	if v, ok := d.GetOk("history16"); ok || d.HasChange("history16") {
		t, err := expandObjectUserLocalDynamicMappingHistory162edl(d, v, "history16")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history16"] = t
		}
	}

	if v, ok := d.GetOk("history17"); ok || d.HasChange("history17") {
		t, err := expandObjectUserLocalDynamicMappingHistory172edl(d, v, "history17")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history17"] = t
		}
	}

	if v, ok := d.GetOk("history18"); ok || d.HasChange("history18") {
		t, err := expandObjectUserLocalDynamicMappingHistory182edl(d, v, "history18")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history18"] = t
		}
	}

	if v, ok := d.GetOk("history19"); ok || d.HasChange("history19") {
		t, err := expandObjectUserLocalDynamicMappingHistory192edl(d, v, "history19")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history19"] = t
		}
	}

	if v, ok := d.GetOk("history2"); ok || d.HasChange("history2") {
		t, err := expandObjectUserLocalDynamicMappingHistory22edl(d, v, "history2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history2"] = t
		}
	}

	if v, ok := d.GetOk("history3"); ok || d.HasChange("history3") {
		t, err := expandObjectUserLocalDynamicMappingHistory32edl(d, v, "history3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history3"] = t
		}
	}

	if v, ok := d.GetOk("history4"); ok || d.HasChange("history4") {
		t, err := expandObjectUserLocalDynamicMappingHistory42edl(d, v, "history4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history4"] = t
		}
	}

	if v, ok := d.GetOk("history5"); ok || d.HasChange("history5") {
		t, err := expandObjectUserLocalDynamicMappingHistory52edl(d, v, "history5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history5"] = t
		}
	}

	if v, ok := d.GetOk("history6"); ok || d.HasChange("history6") {
		t, err := expandObjectUserLocalDynamicMappingHistory62edl(d, v, "history6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history6"] = t
		}
	}

	if v, ok := d.GetOk("history7"); ok || d.HasChange("history7") {
		t, err := expandObjectUserLocalDynamicMappingHistory72edl(d, v, "history7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history7"] = t
		}
	}

	if v, ok := d.GetOk("history8"); ok || d.HasChange("history8") {
		t, err := expandObjectUserLocalDynamicMappingHistory82edl(d, v, "history8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history8"] = t
		}
	}

	if v, ok := d.GetOk("history9"); ok || d.HasChange("history9") {
		t, err := expandObjectUserLocalDynamicMappingHistory92edl(d, v, "history9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history9"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserLocalDynamicMappingId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandObjectUserLocalDynamicMappingLdapServer2edl(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok || d.HasChange("passwd") {
		t, err := expandObjectUserLocalDynamicMappingPasswd2edl(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("passwd_policy"); ok || d.HasChange("passwd_policy") {
		t, err := expandObjectUserLocalDynamicMappingPasswdPolicy2edl(d, v, "passwd_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd-policy"] = t
		}
	}

	if v, ok := d.GetOk("passwd_time"); ok || d.HasChange("passwd_time") {
		t, err := expandObjectUserLocalDynamicMappingPasswdTime2edl(d, v, "passwd_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd-time"] = t
		}
	}

	if v, ok := d.GetOk("ppk_identity"); ok || d.HasChange("ppk_identity") {
		t, err := expandObjectUserLocalDynamicMappingPpkIdentity2edl(d, v, "ppk_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok || d.HasChange("ppk_secret") {
		t, err := expandObjectUserLocalDynamicMappingPpkSecret2edl(d, v, "ppk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	}

	if v, ok := d.GetOk("qkd_profile"); ok || d.HasChange("qkd_profile") {
		t, err := expandObjectUserLocalDynamicMappingQkdProfile2edl(d, v, "qkd_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd-profile"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok || d.HasChange("radius_server") {
		t, err := expandObjectUserLocalDynamicMappingRadiusServer2edl(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("saml_server"); ok || d.HasChange("saml_server") {
		t, err := expandObjectUserLocalDynamicMappingSamlServer2edl(d, v, "saml_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok || d.HasChange("sms_custom_server") {
		t, err := expandObjectUserLocalDynamicMappingSmsCustomServer2edl(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_phone"); ok || d.HasChange("sms_phone") {
		t, err := expandObjectUserLocalDynamicMappingSmsPhone2edl(d, v, "sms_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-phone"] = t
		}
	}

	if v, ok := d.GetOk("sms_provider"); ok || d.HasChange("sms_provider") {
		t, err := expandObjectUserLocalDynamicMappingSmsProvider2edl(d, v, "sms_provider")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-provider"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok || d.HasChange("sms_server") {
		t, err := expandObjectUserLocalDynamicMappingSmsServer2edl(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectUserLocalDynamicMappingStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tacacs_server"); ok || d.HasChange("tacacs_server") {
		t, err := expandObjectUserLocalDynamicMappingTacacsServer2edl(d, v, "tacacs_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tacacs+-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandObjectUserLocalDynamicMappingTwoFactor2edl(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok || d.HasChange("two_factor_authentication") {
		t, err := expandObjectUserLocalDynamicMappingTwoFactorAuthentication2edl(d, v, "two_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok || d.HasChange("two_factor_notification") {
		t, err := expandObjectUserLocalDynamicMappingTwoFactorNotification2edl(d, v, "two_factor_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectUserLocalDynamicMappingType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username_case_insensitivity"); ok || d.HasChange("username_case_insensitivity") {
		t, err := expandObjectUserLocalDynamicMappingUsernameCaseInsensitivity2edl(d, v, "username_case_insensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-insensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitivity"); ok || d.HasChange("username_case_sensitivity") {
		t, err := expandObjectUserLocalDynamicMappingUsernameCaseSensitivity2edl(d, v, "username_case_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_sensitivity"); ok || d.HasChange("username_sensitivity") {
		t, err := expandObjectUserLocalDynamicMappingUsernameSensitivity2edl(d, v, "username_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandObjectUserLocalDynamicMappingUuid2edl(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("workstation"); ok || d.HasChange("workstation") {
		t, err := expandObjectUserLocalDynamicMappingWorkstation2edl(d, v, "workstation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workstation"] = t
		}
	}

	return &obj, nil
}
