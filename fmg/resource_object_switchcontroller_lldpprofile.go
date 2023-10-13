// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch LLDP profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerLldpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerLldpProfileCreate,
		Read:   resourceObjectSwitchControllerLldpProfileRead,
		Update: resourceObjectSwitchControllerLldpProfileUpdate,
		Delete: resourceObjectSwitchControllerLldpProfileDelete,

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
			"n8021_tlvs": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"n8023_tlvs": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_isl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth_encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_isl_auth_macsec_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_isl_auth_reauth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_isl_hello_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_isl_port_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_isl_receive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_mclag_icl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_tlvs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"information_string": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"oui": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subtype": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"med_location_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sys_location_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"med_network_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vlan_intf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"med_tlvs": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
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

func resourceObjectSwitchControllerLldpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerLldpProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerLldpProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerLldpProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerLldpProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerLldpProfileRead(d, m)
}

func resourceObjectSwitchControllerLldpProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerLldpProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerLldpProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerLldpProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerLldpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerLldpProfileRead(d, m)
}

func resourceObjectSwitchControllerLldpProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerLldpProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerLldpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerLldpProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerLldpProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerLldpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerLldpProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerLldpProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerLldpProfile8021Tlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerLldpProfile8023Tlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerLldpProfileAutoIsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslAuthEncrypt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslAuthIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslAuthMacsecProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslAuthReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslAuthUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslHelloTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslPortGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoIslReceiveTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileAutoMclagIcl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileCustomTlvs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := i["information-string"]; ok {
			v := flattenObjectSwitchControllerLldpProfileCustomTlvsInformationString(i["information-string"], d, pre_append)
			tmp["information_string"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-CustomTlvs-InformationString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSwitchControllerLldpProfileCustomTlvsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-CustomTlvs-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := i["oui"]; ok {
			v := flattenObjectSwitchControllerLldpProfileCustomTlvsOui(i["oui"], d, pre_append)
			tmp["oui"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-CustomTlvs-Oui")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := i["subtype"]; ok {
			v := flattenObjectSwitchControllerLldpProfileCustomTlvsSubtype(i["subtype"], d, pre_append)
			tmp["subtype"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-CustomTlvs-Subtype")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsInformationString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsOui(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileCustomTlvsSubtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedLocationService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectSwitchControllerLldpProfileMedLocationServiceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedLocationService-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedLocationServiceStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedLocationService-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sys_location_id"
		if _, ok := i["sys-location-id"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedLocationServiceSysLocationId(i["sys-location-id"], d, pre_append)
			tmp["sys_location_id"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedLocationService-SysLocationId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerLldpProfileMedLocationServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedLocationServiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedLocationServiceSysLocationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedNetworkPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_vlan"
		if _, ok := i["assign-vlan"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(i["assign-vlan"], d, pre_append)
			tmp["assign_vlan"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedNetworkPolicy-AssignVlan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := i["dscp"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedNetworkPolicyDscp(i["dscp"], d, pre_append)
			tmp["dscp"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedNetworkPolicy-Dscp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedNetworkPolicyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedNetworkPolicy-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedNetworkPolicyPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedNetworkPolicy-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedNetworkPolicyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedNetworkPolicy-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedNetworkPolicyVlan(i["vlan"], d, pre_append)
			tmp["vlan"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedNetworkPolicy-Vlan")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_intf"
		if _, ok := i["vlan-intf"]; ok {
			v := flattenObjectSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(i["vlan-intf"], d, pre_append)
			tmp["vlan_intf"] = fortiAPISubPartPatch(v, "ObjectSwitchControllerLldpProfile-MedNetworkPolicy-VlanIntf")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedNetworkPolicyDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedNetworkPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedNetworkPolicyPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedNetworkPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedNetworkPolicyVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerLldpProfileMedTlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerLldpProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerLldpProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("n8021_tlvs", flattenObjectSwitchControllerLldpProfile8021Tlvs(o["802.1-tlvs"], d, "n8021_tlvs")); err != nil {
		if vv, ok := fortiAPIPatch(o["802.1-tlvs"], "ObjectSwitchControllerLldpProfile-8021Tlvs"); ok {
			if err = d.Set("n8021_tlvs", vv); err != nil {
				return fmt.Errorf("Error reading n8021_tlvs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n8021_tlvs: %v", err)
		}
	}

	if err = d.Set("n8023_tlvs", flattenObjectSwitchControllerLldpProfile8023Tlvs(o["802.3-tlvs"], d, "n8023_tlvs")); err != nil {
		if vv, ok := fortiAPIPatch(o["802.3-tlvs"], "ObjectSwitchControllerLldpProfile-8023Tlvs"); ok {
			if err = d.Set("n8023_tlvs", vv); err != nil {
				return fmt.Errorf("Error reading n8023_tlvs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading n8023_tlvs: %v", err)
		}
	}

	if err = d.Set("auto_isl", flattenObjectSwitchControllerLldpProfileAutoIsl(o["auto-isl"], d, "auto_isl")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl"], "ObjectSwitchControllerLldpProfile-AutoIsl"); ok {
			if err = d.Set("auto_isl", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth", flattenObjectSwitchControllerLldpProfileAutoIslAuth(o["auto-isl-auth"], d, "auto_isl_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth"], "ObjectSwitchControllerLldpProfile-AutoIslAuth"); ok {
			if err = d.Set("auto_isl_auth", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_encrypt", flattenObjectSwitchControllerLldpProfileAutoIslAuthEncrypt(o["auto-isl-auth-encrypt"], d, "auto_isl_auth_encrypt")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-encrypt"], "ObjectSwitchControllerLldpProfile-AutoIslAuthEncrypt"); ok {
			if err = d.Set("auto_isl_auth_encrypt", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_encrypt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_encrypt: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_identity", flattenObjectSwitchControllerLldpProfileAutoIslAuthIdentity(o["auto-isl-auth-identity"], d, "auto_isl_auth_identity")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-identity"], "ObjectSwitchControllerLldpProfile-AutoIslAuthIdentity"); ok {
			if err = d.Set("auto_isl_auth_identity", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_identity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_identity: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_macsec_profile", flattenObjectSwitchControllerLldpProfileAutoIslAuthMacsecProfile(o["auto-isl-auth-macsec-profile"], d, "auto_isl_auth_macsec_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-macsec-profile"], "ObjectSwitchControllerLldpProfile-AutoIslAuthMacsecProfile"); ok {
			if err = d.Set("auto_isl_auth_macsec_profile", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_macsec_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_macsec_profile: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_reauth", flattenObjectSwitchControllerLldpProfileAutoIslAuthReauth(o["auto-isl-auth-reauth"], d, "auto_isl_auth_reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-reauth"], "ObjectSwitchControllerLldpProfile-AutoIslAuthReauth"); ok {
			if err = d.Set("auto_isl_auth_reauth", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_reauth: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_user", flattenObjectSwitchControllerLldpProfileAutoIslAuthUser(o["auto-isl-auth-user"], d, "auto_isl_auth_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-auth-user"], "ObjectSwitchControllerLldpProfile-AutoIslAuthUser"); ok {
			if err = d.Set("auto_isl_auth_user", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_auth_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_auth_user: %v", err)
		}
	}

	if err = d.Set("auto_isl_hello_timer", flattenObjectSwitchControllerLldpProfileAutoIslHelloTimer(o["auto-isl-hello-timer"], d, "auto_isl_hello_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-hello-timer"], "ObjectSwitchControllerLldpProfile-AutoIslHelloTimer"); ok {
			if err = d.Set("auto_isl_hello_timer", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_hello_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_hello_timer: %v", err)
		}
	}

	if err = d.Set("auto_isl_port_group", flattenObjectSwitchControllerLldpProfileAutoIslPortGroup(o["auto-isl-port-group"], d, "auto_isl_port_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-port-group"], "ObjectSwitchControllerLldpProfile-AutoIslPortGroup"); ok {
			if err = d.Set("auto_isl_port_group", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_port_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_port_group: %v", err)
		}
	}

	if err = d.Set("auto_isl_receive_timeout", flattenObjectSwitchControllerLldpProfileAutoIslReceiveTimeout(o["auto-isl-receive-timeout"], d, "auto_isl_receive_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-isl-receive-timeout"], "ObjectSwitchControllerLldpProfile-AutoIslReceiveTimeout"); ok {
			if err = d.Set("auto_isl_receive_timeout", vv); err != nil {
				return fmt.Errorf("Error reading auto_isl_receive_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_isl_receive_timeout: %v", err)
		}
	}

	if err = d.Set("auto_mclag_icl", flattenObjectSwitchControllerLldpProfileAutoMclagIcl(o["auto-mclag-icl"], d, "auto_mclag_icl")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-mclag-icl"], "ObjectSwitchControllerLldpProfile-AutoMclagIcl"); ok {
			if err = d.Set("auto_mclag_icl", vv); err != nil {
				return fmt.Errorf("Error reading auto_mclag_icl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_mclag_icl: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_tlvs", flattenObjectSwitchControllerLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-tlvs"], "ObjectSwitchControllerLldpProfile-CustomTlvs"); ok {
				if err = d.Set("custom_tlvs", vv); err != nil {
					return fmt.Errorf("Error reading custom_tlvs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_tlvs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_tlvs"); ok {
			if err = d.Set("custom_tlvs", flattenObjectSwitchControllerLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-tlvs"], "ObjectSwitchControllerLldpProfile-CustomTlvs"); ok {
					if err = d.Set("custom_tlvs", vv); err != nil {
						return fmt.Errorf("Error reading custom_tlvs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_tlvs: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("med_location_service", flattenObjectSwitchControllerLldpProfileMedLocationService(o["med-location-service"], d, "med_location_service")); err != nil {
			if vv, ok := fortiAPIPatch(o["med-location-service"], "ObjectSwitchControllerLldpProfile-MedLocationService"); ok {
				if err = d.Set("med_location_service", vv); err != nil {
					return fmt.Errorf("Error reading med_location_service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading med_location_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_location_service"); ok {
			if err = d.Set("med_location_service", flattenObjectSwitchControllerLldpProfileMedLocationService(o["med-location-service"], d, "med_location_service")); err != nil {
				if vv, ok := fortiAPIPatch(o["med-location-service"], "ObjectSwitchControllerLldpProfile-MedLocationService"); ok {
					if err = d.Set("med_location_service", vv); err != nil {
						return fmt.Errorf("Error reading med_location_service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading med_location_service: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("med_network_policy", flattenObjectSwitchControllerLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy")); err != nil {
			if vv, ok := fortiAPIPatch(o["med-network-policy"], "ObjectSwitchControllerLldpProfile-MedNetworkPolicy"); ok {
				if err = d.Set("med_network_policy", vv); err != nil {
					return fmt.Errorf("Error reading med_network_policy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading med_network_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_network_policy"); ok {
			if err = d.Set("med_network_policy", flattenObjectSwitchControllerLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy")); err != nil {
				if vv, ok := fortiAPIPatch(o["med-network-policy"], "ObjectSwitchControllerLldpProfile-MedNetworkPolicy"); ok {
					if err = d.Set("med_network_policy", vv); err != nil {
						return fmt.Errorf("Error reading med_network_policy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading med_network_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("med_tlvs", flattenObjectSwitchControllerLldpProfileMedTlvs(o["med-tlvs"], d, "med_tlvs")); err != nil {
		if vv, ok := fortiAPIPatch(o["med-tlvs"], "ObjectSwitchControllerLldpProfile-MedTlvs"); ok {
			if err = d.Set("med_tlvs", vv); err != nil {
				return fmt.Errorf("Error reading med_tlvs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading med_tlvs: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerLldpProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerLldpProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerLldpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerLldpProfile8021Tlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerLldpProfile8023Tlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerLldpProfileAutoIsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslAuthEncrypt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslAuthIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslAuthMacsecProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslAuthReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslAuthUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslHelloTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslPortGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoIslReceiveTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileAutoMclagIcl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["information-string"], _ = expandObjectSwitchControllerLldpProfileCustomTlvsInformationString(d, i["information_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSwitchControllerLldpProfileCustomTlvsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["oui"], _ = expandObjectSwitchControllerLldpProfileCustomTlvsOui(d, i["oui"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["subtype"], _ = expandObjectSwitchControllerLldpProfileCustomTlvsSubtype(d, i["subtype"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerLldpProfileCustomTlvsInformationString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileCustomTlvsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileCustomTlvsOui(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileCustomTlvsSubtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedLocationService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandObjectSwitchControllerLldpProfileMedLocationServiceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectSwitchControllerLldpProfileMedLocationServiceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sys_location_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sys-location-id"], _ = expandObjectSwitchControllerLldpProfileMedLocationServiceSysLocationId(d, i["sys_location_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerLldpProfileMedLocationServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedLocationServiceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedLocationServiceSysLocationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedNetworkPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["assign-vlan"], _ = expandObjectSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(d, i["assign_vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp"], _ = expandObjectSwitchControllerLldpProfileMedNetworkPolicyDscp(d, i["dscp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectSwitchControllerLldpProfileMedNetworkPolicyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandObjectSwitchControllerLldpProfileMedNetworkPolicyPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectSwitchControllerLldpProfileMedNetworkPolicyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan"], _ = expandObjectSwitchControllerLldpProfileMedNetworkPolicyVlan(d, i["vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_intf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vlan-intf"], _ = expandObjectSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(d, i["vlan_intf"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedNetworkPolicyDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedNetworkPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedNetworkPolicyPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedNetworkPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedNetworkPolicyVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerLldpProfileMedTlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerLldpProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerLldpProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("n8021_tlvs"); ok || d.HasChange("n8021_tlvs") {
		t, err := expandObjectSwitchControllerLldpProfile8021Tlvs(d, v, "n8021_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.1-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("n8023_tlvs"); ok || d.HasChange("n8023_tlvs") {
		t, err := expandObjectSwitchControllerLldpProfile8023Tlvs(d, v, "n8023_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.3-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl"); ok || d.HasChange("auto_isl") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIsl(d, v, "auto_isl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth"); ok || d.HasChange("auto_isl_auth") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslAuth(d, v, "auto_isl_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_encrypt"); ok || d.HasChange("auto_isl_auth_encrypt") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslAuthEncrypt(d, v, "auto_isl_auth_encrypt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-encrypt"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_identity"); ok || d.HasChange("auto_isl_auth_identity") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslAuthIdentity(d, v, "auto_isl_auth_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-identity"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_macsec_profile"); ok || d.HasChange("auto_isl_auth_macsec_profile") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslAuthMacsecProfile(d, v, "auto_isl_auth_macsec_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-macsec-profile"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_reauth"); ok || d.HasChange("auto_isl_auth_reauth") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslAuthReauth(d, v, "auto_isl_auth_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-reauth"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_user"); ok || d.HasChange("auto_isl_auth_user") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslAuthUser(d, v, "auto_isl_auth_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-user"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_hello_timer"); ok || d.HasChange("auto_isl_hello_timer") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslHelloTimer(d, v, "auto_isl_hello_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-hello-timer"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_port_group"); ok || d.HasChange("auto_isl_port_group") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslPortGroup(d, v, "auto_isl_port_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-port-group"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_receive_timeout"); ok || d.HasChange("auto_isl_receive_timeout") {
		t, err := expandObjectSwitchControllerLldpProfileAutoIslReceiveTimeout(d, v, "auto_isl_receive_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOk("auto_mclag_icl"); ok || d.HasChange("auto_mclag_icl") {
		t, err := expandObjectSwitchControllerLldpProfileAutoMclagIcl(d, v, "auto_mclag_icl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-mclag-icl"] = t
		}
	}

	if v, ok := d.GetOk("custom_tlvs"); ok || d.HasChange("custom_tlvs") {
		t, err := expandObjectSwitchControllerLldpProfileCustomTlvs(d, v, "custom_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("med_location_service"); ok || d.HasChange("med_location_service") {
		t, err := expandObjectSwitchControllerLldpProfileMedLocationService(d, v, "med_location_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-location-service"] = t
		}
	}

	if v, ok := d.GetOk("med_network_policy"); ok || d.HasChange("med_network_policy") {
		t, err := expandObjectSwitchControllerLldpProfileMedNetworkPolicy(d, v, "med_network_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-network-policy"] = t
		}
	}

	if v, ok := d.GetOk("med_tlvs"); ok || d.HasChange("med_tlvs") {
		t, err := expandObjectSwitchControllerLldpProfileMedTlvs(d, v, "med_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerLldpProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
