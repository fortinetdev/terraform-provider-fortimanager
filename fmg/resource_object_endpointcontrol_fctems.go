// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiClient Enterprise Management Server (EMS) entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectEndpointControlFctems() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectEndpointControlFctemsCreate,
		Read:   resourceObjectEndpointControlFctemsRead,
		Update: resourceObjectEndpointControlFctemsUpdate,
		Delete: resourceObjectEndpointControlFctemsDelete,

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
			"ca_cn_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"admin_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"call_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"capabilities": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certificate_fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dirty_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ems_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fortinetone_cloud_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"out_of_sync_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"preserve_ssl_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_avatars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_malware_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_sysinfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_tags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_vulnerabilities": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_tags_to_all_vdoms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ca_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verified_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verifying_ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"websocket_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectEndpointControlFctemsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectEndpointControlFctems(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectEndpointControlFctems resource while getting object: %v", err)
	}

	_, err = c.CreateObjectEndpointControlFctems(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectEndpointControlFctems resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectEndpointControlFctemsRead(d, m)
}

func resourceObjectEndpointControlFctemsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectEndpointControlFctems(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEndpointControlFctems resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectEndpointControlFctems(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectEndpointControlFctems resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectEndpointControlFctemsRead(d, m)
}

func resourceObjectEndpointControlFctemsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectEndpointControlFctems(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectEndpointControlFctems resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectEndpointControlFctemsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectEndpointControlFctems(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEndpointControlFctems resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectEndpointControlFctems(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectEndpointControlFctems resource from API: %v", err)
	}
	return nil
}

func flattenObjectEndpointControlFctemsCaCnInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsAdminPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectEndpointControlFctemsAdminUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsCallTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsCapabilities(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectEndpointControlFctemsCertificateFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsCloudServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsDirtyReason(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsEmsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsFortinetoneCloudAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsOutOfSyncThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsPreserveSslSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsPullAvatars(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsPullMalwareHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsPullSysinfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsPullTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsPullVulnerabilities(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsSendTagsToAllVdoms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsTrustCaCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsVerifiedCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsVerifyingCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsTenantId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsStatusCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectEndpointControlFctemsWebsocketOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectEndpointControlFctems(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ca_cn_info", flattenObjectEndpointControlFctemsCaCnInfo(o["ca-cn-info"], d, "ca_cn_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cn-info"], "ObjectEndpointControlFctems-CaCnInfo"); ok {
			if err = d.Set("ca_cn_info", vv); err != nil {
				return fmt.Errorf("Error reading ca_cn_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cn_info: %v", err)
		}
	}

	if err = d.Set("admin_password", flattenObjectEndpointControlFctemsAdminPassword(o["admin-password"], d, "admin_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-password"], "ObjectEndpointControlFctems-AdminPassword"); ok {
			if err = d.Set("admin_password", vv); err != nil {
				return fmt.Errorf("Error reading admin_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_password: %v", err)
		}
	}

	if err = d.Set("admin_username", flattenObjectEndpointControlFctemsAdminUsername(o["admin-username"], d, "admin_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["admin-username"], "ObjectEndpointControlFctems-AdminUsername"); ok {
			if err = d.Set("admin_username", vv); err != nil {
				return fmt.Errorf("Error reading admin_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admin_username: %v", err)
		}
	}

	if err = d.Set("call_timeout", flattenObjectEndpointControlFctemsCallTimeout(o["call-timeout"], d, "call_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-timeout"], "ObjectEndpointControlFctems-CallTimeout"); ok {
			if err = d.Set("call_timeout", vv); err != nil {
				return fmt.Errorf("Error reading call_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_timeout: %v", err)
		}
	}

	if err = d.Set("certificate", flattenObjectEndpointControlFctemsCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "ObjectEndpointControlFctems-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("capabilities", flattenObjectEndpointControlFctemsCapabilities(o["capabilities"], d, "capabilities")); err != nil {
		if vv, ok := fortiAPIPatch(o["capabilities"], "ObjectEndpointControlFctems-Capabilities"); ok {
			if err = d.Set("capabilities", vv); err != nil {
				return fmt.Errorf("Error reading capabilities: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capabilities: %v", err)
		}
	}

	if err = d.Set("certificate_fingerprint", flattenObjectEndpointControlFctemsCertificateFingerprint(o["certificate-fingerprint"], d, "certificate_fingerprint")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate-fingerprint"], "ObjectEndpointControlFctems-CertificateFingerprint"); ok {
			if err = d.Set("certificate_fingerprint", vv); err != nil {
				return fmt.Errorf("Error reading certificate_fingerprint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate_fingerprint: %v", err)
		}
	}

	if err = d.Set("cloud_server_type", flattenObjectEndpointControlFctemsCloudServerType(o["cloud-server-type"], d, "cloud_server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["cloud-server-type"], "ObjectEndpointControlFctems-CloudServerType"); ok {
			if err = d.Set("cloud_server_type", vv); err != nil {
				return fmt.Errorf("Error reading cloud_server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cloud_server_type: %v", err)
		}
	}

	if err = d.Set("dirty_reason", flattenObjectEndpointControlFctemsDirtyReason(o["dirty-reason"], d, "dirty_reason")); err != nil {
		if vv, ok := fortiAPIPatch(o["dirty-reason"], "ObjectEndpointControlFctems-DirtyReason"); ok {
			if err = d.Set("dirty_reason", vv); err != nil {
				return fmt.Errorf("Error reading dirty_reason: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dirty_reason: %v", err)
		}
	}

	if err = d.Set("ems_id", flattenObjectEndpointControlFctemsEmsId(o["ems-id"], d, "ems_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-id"], "ObjectEndpointControlFctems-EmsId"); ok {
			if err = d.Set("ems_id", vv); err != nil {
				return fmt.Errorf("Error reading ems_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_id: %v", err)
		}
	}

	if err = d.Set("fortinetone_cloud_authentication", flattenObjectEndpointControlFctemsFortinetoneCloudAuthentication(o["fortinetone-cloud-authentication"], d, "fortinetone_cloud_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortinetone-cloud-authentication"], "ObjectEndpointControlFctems-FortinetoneCloudAuthentication"); ok {
			if err = d.Set("fortinetone_cloud_authentication", vv); err != nil {
				return fmt.Errorf("Error reading fortinetone_cloud_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortinetone_cloud_authentication: %v", err)
		}
	}

	if err = d.Set("https_port", flattenObjectEndpointControlFctemsHttpsPort(o["https-port"], d, "https_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-port"], "ObjectEndpointControlFctems-HttpsPort"); ok {
			if err = d.Set("https_port", vv); err != nil {
				return fmt.Errorf("Error reading https_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectEndpointControlFctemsInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectEndpointControlFctems-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenObjectEndpointControlFctemsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "ObjectEndpointControlFctems-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectEndpointControlFctemsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectEndpointControlFctems-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("out_of_sync_threshold", flattenObjectEndpointControlFctemsOutOfSyncThreshold(o["out-of-sync-threshold"], d, "out_of_sync_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["out-of-sync-threshold"], "ObjectEndpointControlFctems-OutOfSyncThreshold"); ok {
			if err = d.Set("out_of_sync_threshold", vv); err != nil {
				return fmt.Errorf("Error reading out_of_sync_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading out_of_sync_threshold: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenObjectEndpointControlFctemsSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial-number"], "ObjectEndpointControlFctems-SerialNumber"); ok {
			if err = d.Set("serial_number", vv); err != nil {
				return fmt.Errorf("Error reading serial_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("preserve_ssl_session", flattenObjectEndpointControlFctemsPreserveSslSession(o["preserve-ssl-session"], d, "preserve_ssl_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["preserve-ssl-session"], "ObjectEndpointControlFctems-PreserveSslSession"); ok {
			if err = d.Set("preserve_ssl_session", vv); err != nil {
				return fmt.Errorf("Error reading preserve_ssl_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preserve_ssl_session: %v", err)
		}
	}

	if err = d.Set("pull_avatars", flattenObjectEndpointControlFctemsPullAvatars(o["pull-avatars"], d, "pull_avatars")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-avatars"], "ObjectEndpointControlFctems-PullAvatars"); ok {
			if err = d.Set("pull_avatars", vv); err != nil {
				return fmt.Errorf("Error reading pull_avatars: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_avatars: %v", err)
		}
	}

	if err = d.Set("pull_malware_hash", flattenObjectEndpointControlFctemsPullMalwareHash(o["pull-malware-hash"], d, "pull_malware_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-malware-hash"], "ObjectEndpointControlFctems-PullMalwareHash"); ok {
			if err = d.Set("pull_malware_hash", vv); err != nil {
				return fmt.Errorf("Error reading pull_malware_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_malware_hash: %v", err)
		}
	}

	if err = d.Set("pull_sysinfo", flattenObjectEndpointControlFctemsPullSysinfo(o["pull-sysinfo"], d, "pull_sysinfo")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-sysinfo"], "ObjectEndpointControlFctems-PullSysinfo"); ok {
			if err = d.Set("pull_sysinfo", vv); err != nil {
				return fmt.Errorf("Error reading pull_sysinfo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_sysinfo: %v", err)
		}
	}

	if err = d.Set("pull_tags", flattenObjectEndpointControlFctemsPullTags(o["pull-tags"], d, "pull_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-tags"], "ObjectEndpointControlFctems-PullTags"); ok {
			if err = d.Set("pull_tags", vv); err != nil {
				return fmt.Errorf("Error reading pull_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_tags: %v", err)
		}
	}

	if err = d.Set("pull_vulnerabilities", flattenObjectEndpointControlFctemsPullVulnerabilities(o["pull-vulnerabilities"], d, "pull_vulnerabilities")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-vulnerabilities"], "ObjectEndpointControlFctems-PullVulnerabilities"); ok {
			if err = d.Set("pull_vulnerabilities", vv); err != nil {
				return fmt.Errorf("Error reading pull_vulnerabilities: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_vulnerabilities: %v", err)
		}
	}

	if err = d.Set("send_tags_to_all_vdoms", flattenObjectEndpointControlFctemsSendTagsToAllVdoms(o["send-tags-to-all-vdoms"], d, "send_tags_to_all_vdoms")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-tags-to-all-vdoms"], "ObjectEndpointControlFctems-SendTagsToAllVdoms"); ok {
			if err = d.Set("send_tags_to_all_vdoms", vv); err != nil {
				return fmt.Errorf("Error reading send_tags_to_all_vdoms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_tags_to_all_vdoms: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectEndpointControlFctemsServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectEndpointControlFctems-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectEndpointControlFctemsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectEndpointControlFctems-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectEndpointControlFctemsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectEndpointControlFctems-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trust_ca_cn", flattenObjectEndpointControlFctemsTrustCaCn(o["trust-ca-cn"], d, "trust_ca_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ca-cn"], "ObjectEndpointControlFctems-TrustCaCn"); ok {
			if err = d.Set("trust_ca_cn", vv); err != nil {
				return fmt.Errorf("Error reading trust_ca_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ca_cn: %v", err)
		}
	}

	if err = d.Set("verified_cn", flattenObjectEndpointControlFctemsVerifiedCn(o["verified-cn"], d, "verified_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["verified-cn"], "ObjectEndpointControlFctems-VerifiedCn"); ok {
			if err = d.Set("verified_cn", vv); err != nil {
				return fmt.Errorf("Error reading verified_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verified_cn: %v", err)
		}
	}

	if err = d.Set("verifying_ca", flattenObjectEndpointControlFctemsVerifyingCa(o["verifying-ca"], d, "verifying_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["verifying-ca"], "ObjectEndpointControlFctems-VerifyingCa"); ok {
			if err = d.Set("verifying_ca", vv); err != nil {
				return fmt.Errorf("Error reading verifying_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verifying_ca: %v", err)
		}
	}

	if err = d.Set("tenant_id", flattenObjectEndpointControlFctemsTenantId(o["tenant-id"], d, "tenant_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-id"], "ObjectEndpointControlFctems-TenantId"); ok {
			if err = d.Set("tenant_id", vv); err != nil {
				return fmt.Errorf("Error reading tenant_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("status_check_interval", flattenObjectEndpointControlFctemsStatusCheckInterval(o["status-check-interval"], d, "status_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-check-interval"], "ObjectEndpointControlFctems-StatusCheckInterval"); ok {
			if err = d.Set("status_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading status_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_check_interval: %v", err)
		}
	}

	if err = d.Set("websocket_override", flattenObjectEndpointControlFctemsWebsocketOverride(o["websocket-override"], d, "websocket_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["websocket-override"], "ObjectEndpointControlFctems-WebsocketOverride"); ok {
			if err = d.Set("websocket_override", vv); err != nil {
				return fmt.Errorf("Error reading websocket_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading websocket_override: %v", err)
		}
	}

	return nil
}

func flattenObjectEndpointControlFctemsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectEndpointControlFctemsCaCnInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsAdminPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectEndpointControlFctemsAdminUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsCallTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsCapabilities(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectEndpointControlFctemsCertificateFingerprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsCloudServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsDirtyReason(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsEmsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsFortinetoneCloudAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsOutOfSyncThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsPreserveSslSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsPullAvatars(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsPullMalwareHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsPullSysinfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsPullTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsPullVulnerabilities(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsSendTagsToAllVdoms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsTrustCaCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsVerifiedCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsVerifyingCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsTenantId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsStatusCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectEndpointControlFctemsWebsocketOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectEndpointControlFctems(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca_cn_info"); ok || d.HasChange("ca_cn_info") {
		t, err := expandObjectEndpointControlFctemsCaCnInfo(d, v, "ca_cn_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cn-info"] = t
		}
	}

	if v, ok := d.GetOk("admin_password"); ok || d.HasChange("admin_password") {
		t, err := expandObjectEndpointControlFctemsAdminPassword(d, v, "admin_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-password"] = t
		}
	}

	if v, ok := d.GetOk("admin_username"); ok || d.HasChange("admin_username") {
		t, err := expandObjectEndpointControlFctemsAdminUsername(d, v, "admin_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-username"] = t
		}
	}

	if v, ok := d.GetOk("call_timeout"); ok || d.HasChange("call_timeout") {
		t, err := expandObjectEndpointControlFctemsCallTimeout(d, v, "call_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-timeout"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandObjectEndpointControlFctemsCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("capabilities"); ok || d.HasChange("capabilities") {
		t, err := expandObjectEndpointControlFctemsCapabilities(d, v, "capabilities")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capabilities"] = t
		}
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok || d.HasChange("certificate_fingerprint") {
		t, err := expandObjectEndpointControlFctemsCertificateFingerprint(d, v, "certificate_fingerprint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-fingerprint"] = t
		}
	}

	if v, ok := d.GetOk("cloud_server_type"); ok || d.HasChange("cloud_server_type") {
		t, err := expandObjectEndpointControlFctemsCloudServerType(d, v, "cloud_server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cloud-server-type"] = t
		}
	}

	if v, ok := d.GetOk("dirty_reason"); ok || d.HasChange("dirty_reason") {
		t, err := expandObjectEndpointControlFctemsDirtyReason(d, v, "dirty_reason")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dirty-reason"] = t
		}
	}

	if v, ok := d.GetOk("ems_id"); ok || d.HasChange("ems_id") {
		t, err := expandObjectEndpointControlFctemsEmsId(d, v, "ems_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-id"] = t
		}
	}

	if v, ok := d.GetOk("fortinetone_cloud_authentication"); ok || d.HasChange("fortinetone_cloud_authentication") {
		t, err := expandObjectEndpointControlFctemsFortinetoneCloudAuthentication(d, v, "fortinetone_cloud_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinetone-cloud-authentication"] = t
		}
	}

	if v, ok := d.GetOk("https_port"); ok || d.HasChange("https_port") {
		t, err := expandObjectEndpointControlFctemsHttpsPort(d, v, "https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-port"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectEndpointControlFctemsInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandObjectEndpointControlFctemsInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectEndpointControlFctemsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("out_of_sync_threshold"); ok || d.HasChange("out_of_sync_threshold") {
		t, err := expandObjectEndpointControlFctemsOutOfSyncThreshold(d, v, "out_of_sync_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-of-sync-threshold"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok || d.HasChange("serial_number") {
		t, err := expandObjectEndpointControlFctemsSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("preserve_ssl_session"); ok || d.HasChange("preserve_ssl_session") {
		t, err := expandObjectEndpointControlFctemsPreserveSslSession(d, v, "preserve_ssl_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preserve-ssl-session"] = t
		}
	}

	if v, ok := d.GetOk("pull_avatars"); ok || d.HasChange("pull_avatars") {
		t, err := expandObjectEndpointControlFctemsPullAvatars(d, v, "pull_avatars")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-avatars"] = t
		}
	}

	if v, ok := d.GetOk("pull_malware_hash"); ok || d.HasChange("pull_malware_hash") {
		t, err := expandObjectEndpointControlFctemsPullMalwareHash(d, v, "pull_malware_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-malware-hash"] = t
		}
	}

	if v, ok := d.GetOk("pull_sysinfo"); ok || d.HasChange("pull_sysinfo") {
		t, err := expandObjectEndpointControlFctemsPullSysinfo(d, v, "pull_sysinfo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("pull_tags"); ok || d.HasChange("pull_tags") {
		t, err := expandObjectEndpointControlFctemsPullTags(d, v, "pull_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-tags"] = t
		}
	}

	if v, ok := d.GetOk("pull_vulnerabilities"); ok || d.HasChange("pull_vulnerabilities") {
		t, err := expandObjectEndpointControlFctemsPullVulnerabilities(d, v, "pull_vulnerabilities")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-vulnerabilities"] = t
		}
	}

	if v, ok := d.GetOk("send_tags_to_all_vdoms"); ok || d.HasChange("send_tags_to_all_vdoms") {
		t, err := expandObjectEndpointControlFctemsSendTagsToAllVdoms(d, v, "send_tags_to_all_vdoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-tags-to-all-vdoms"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandObjectEndpointControlFctemsServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectEndpointControlFctemsSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectEndpointControlFctemsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trust_ca_cn"); ok || d.HasChange("trust_ca_cn") {
		t, err := expandObjectEndpointControlFctemsTrustCaCn(d, v, "trust_ca_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ca-cn"] = t
		}
	}

	if v, ok := d.GetOk("verified_cn"); ok || d.HasChange("verified_cn") {
		t, err := expandObjectEndpointControlFctemsVerifiedCn(d, v, "verified_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verified-cn"] = t
		}
	}

	if v, ok := d.GetOk("verifying_ca"); ok || d.HasChange("verifying_ca") {
		t, err := expandObjectEndpointControlFctemsVerifyingCa(d, v, "verifying_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verifying-ca"] = t
		}
	}

	if v, ok := d.GetOk("tenant_id"); ok || d.HasChange("tenant_id") {
		t, err := expandObjectEndpointControlFctemsTenantId(d, v, "tenant_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-id"] = t
		}
	}

	if v, ok := d.GetOk("status_check_interval"); ok || d.HasChange("status_check_interval") {
		t, err := expandObjectEndpointControlFctemsStatusCheckInterval(d, v, "status_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("websocket_override"); ok || d.HasChange("websocket_override") {
		t, err := expandObjectEndpointControlFctemsWebsocketOverride(d, v, "websocket_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websocket-override"] = t
		}
	}

	return &obj, nil
}
