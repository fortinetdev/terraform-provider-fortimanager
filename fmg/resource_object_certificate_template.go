// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCertificate Template

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCertificateTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCertificateTemplateCreate,
		Read:   resourceObjectCertificateTemplateRead,
		Update: resourceObjectCertificateTemplateUpdate,
		Delete: resourceObjectCertificateTemplateDelete,

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
			"city": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"curve_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digest_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enroll_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_ca_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_client_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"est_http_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"est_http_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_regeneration_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"est_server_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"est_server_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"est_srp_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"est_srp_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_unit": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scep_ca_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scep_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"scep_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_alt_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"validity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCertificateTemplateCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCertificateTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCertificateTemplate resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectCertificateTemplate(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectCertificateTemplate(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectCertificateTemplate resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectCertificateTemplate(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectCertificateTemplate resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCertificateTemplateRead(d, m)
}

func resourceObjectCertificateTemplateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCertificateTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCertificateTemplate resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCertificateTemplate(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCertificateTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCertificateTemplateRead(d, m)
}

func resourceObjectCertificateTemplateDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCertificateTemplate(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCertificateTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCertificateTemplateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCertificateTemplate(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectCertificateTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCertificateTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCertificateTemplate resource from API: %v", err)
	}
	return nil
}

func flattenObjectCertificateTemplateCity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateCurveName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateDigestType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateEnrollProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateEstCaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateEstClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCertificateTemplateEstHttpUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateEstRegenerationMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateEstServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCertificateTemplateEstServerUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateEstSrpUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateKeySize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateOrganization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateOrganizationUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCertificateTemplateScepCaIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateScepServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateSubjectAltName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateSubjectName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCertificateTemplateValidity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCertificateTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("city", flattenObjectCertificateTemplateCity(o["city"], d, "city")); err != nil {
		if vv, ok := fortiAPIPatch(o["city"], "ObjectCertificateTemplate-City"); ok {
			if err = d.Set("city", vv); err != nil {
				return fmt.Errorf("Error reading city: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading city: %v", err)
		}
	}

	if err = d.Set("country", flattenObjectCertificateTemplateCountry(o["country"], d, "country")); err != nil {
		if vv, ok := fortiAPIPatch(o["country"], "ObjectCertificateTemplate-Country"); ok {
			if err = d.Set("country", vv); err != nil {
				return fmt.Errorf("Error reading country: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("curve_name", flattenObjectCertificateTemplateCurveName(o["curve-name"], d, "curve_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["curve-name"], "ObjectCertificateTemplate-CurveName"); ok {
			if err = d.Set("curve_name", vv); err != nil {
				return fmt.Errorf("Error reading curve_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading curve_name: %v", err)
		}
	}

	if err = d.Set("digest_type", flattenObjectCertificateTemplateDigestType(o["digest-type"], d, "digest_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-type"], "ObjectCertificateTemplate-DigestType"); ok {
			if err = d.Set("digest_type", vv); err != nil {
				return fmt.Errorf("Error reading digest_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_type: %v", err)
		}
	}

	if err = d.Set("email", flattenObjectCertificateTemplateEmail(o["email"], d, "email")); err != nil {
		if vv, ok := fortiAPIPatch(o["email"], "ObjectCertificateTemplate-Email"); ok {
			if err = d.Set("email", vv); err != nil {
				return fmt.Errorf("Error reading email: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email: %v", err)
		}
	}

	if err = d.Set("enroll_protocol", flattenObjectCertificateTemplateEnrollProtocol(o["enroll-protocol"], d, "enroll_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["enroll-protocol"], "ObjectCertificateTemplate-EnrollProtocol"); ok {
			if err = d.Set("enroll_protocol", vv); err != nil {
				return fmt.Errorf("Error reading enroll_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enroll_protocol: %v", err)
		}
	}

	if err = d.Set("est_ca_id", flattenObjectCertificateTemplateEstCaId(o["est-ca-id"], d, "est_ca_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-ca-id"], "ObjectCertificateTemplate-EstCaId"); ok {
			if err = d.Set("est_ca_id", vv); err != nil {
				return fmt.Errorf("Error reading est_ca_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_ca_id: %v", err)
		}
	}

	if err = d.Set("est_client_cert", flattenObjectCertificateTemplateEstClientCert(o["est-client-cert"], d, "est_client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-client-cert"], "ObjectCertificateTemplate-EstClientCert"); ok {
			if err = d.Set("est_client_cert", vv); err != nil {
				return fmt.Errorf("Error reading est_client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_client_cert: %v", err)
		}
	}

	if err = d.Set("est_http_username", flattenObjectCertificateTemplateEstHttpUsername(o["est-http-username"], d, "est_http_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-http-username"], "ObjectCertificateTemplate-EstHttpUsername"); ok {
			if err = d.Set("est_http_username", vv); err != nil {
				return fmt.Errorf("Error reading est_http_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_http_username: %v", err)
		}
	}

	if err = d.Set("est_regeneration_method", flattenObjectCertificateTemplateEstRegenerationMethod(o["est-regeneration-method"], d, "est_regeneration_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-regeneration-method"], "ObjectCertificateTemplate-EstRegenerationMethod"); ok {
			if err = d.Set("est_regeneration_method", vv); err != nil {
				return fmt.Errorf("Error reading est_regeneration_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_regeneration_method: %v", err)
		}
	}

	if err = d.Set("est_server_cert", flattenObjectCertificateTemplateEstServerCert(o["est-server-cert"], d, "est_server_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-server-cert"], "ObjectCertificateTemplate-EstServerCert"); ok {
			if err = d.Set("est_server_cert", vv); err != nil {
				return fmt.Errorf("Error reading est_server_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_server_cert: %v", err)
		}
	}

	if err = d.Set("est_server_url", flattenObjectCertificateTemplateEstServerUrl(o["est-server-url"], d, "est_server_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-server-url"], "ObjectCertificateTemplate-EstServerUrl"); ok {
			if err = d.Set("est_server_url", vv); err != nil {
				return fmt.Errorf("Error reading est_server_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_server_url: %v", err)
		}
	}

	if err = d.Set("est_srp_username", flattenObjectCertificateTemplateEstSrpUsername(o["est-srp-username"], d, "est_srp_username")); err != nil {
		if vv, ok := fortiAPIPatch(o["est-srp-username"], "ObjectCertificateTemplate-EstSrpUsername"); ok {
			if err = d.Set("est_srp_username", vv); err != nil {
				return fmt.Errorf("Error reading est_srp_username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading est_srp_username: %v", err)
		}
	}

	if err = d.Set("id_type", flattenObjectCertificateTemplateIdType(o["id-type"], d, "id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["id-type"], "ObjectCertificateTemplate-IdType"); ok {
			if err = d.Set("id_type", vv); err != nil {
				return fmt.Errorf("Error reading id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading id_type: %v", err)
		}
	}

	if err = d.Set("key_size", flattenObjectCertificateTemplateKeySize(o["key-size"], d, "key_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-size"], "ObjectCertificateTemplate-KeySize"); ok {
			if err = d.Set("key_size", vv); err != nil {
				return fmt.Errorf("Error reading key_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_size: %v", err)
		}
	}

	if err = d.Set("key_type", flattenObjectCertificateTemplateKeyType(o["key-type"], d, "key_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-type"], "ObjectCertificateTemplate-KeyType"); ok {
			if err = d.Set("key_type", vv); err != nil {
				return fmt.Errorf("Error reading key_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCertificateTemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCertificateTemplate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("organization", flattenObjectCertificateTemplateOrganization(o["organization"], d, "organization")); err != nil {
		if vv, ok := fortiAPIPatch(o["organization"], "ObjectCertificateTemplate-Organization"); ok {
			if err = d.Set("organization", vv); err != nil {
				return fmt.Errorf("Error reading organization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading organization: %v", err)
		}
	}

	if err = d.Set("organization_unit", flattenObjectCertificateTemplateOrganizationUnit(o["organization-unit"], d, "organization_unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["organization-unit"], "ObjectCertificateTemplate-OrganizationUnit"); ok {
			if err = d.Set("organization_unit", vv); err != nil {
				return fmt.Errorf("Error reading organization_unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading organization_unit: %v", err)
		}
	}

	if err = d.Set("scep_ca_identifier", flattenObjectCertificateTemplateScepCaIdentifier(o["scep-ca-identifier"], d, "scep_ca_identifier")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-ca-identifier"], "ObjectCertificateTemplate-ScepCaIdentifier"); ok {
			if err = d.Set("scep_ca_identifier", vv); err != nil {
				return fmt.Errorf("Error reading scep_ca_identifier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_ca_identifier: %v", err)
		}
	}

	if err = d.Set("scep_server", flattenObjectCertificateTemplateScepServer(o["scep-server"], d, "scep_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-server"], "ObjectCertificateTemplate-ScepServer"); ok {
			if err = d.Set("scep_server", vv); err != nil {
				return fmt.Errorf("Error reading scep_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectCertificateTemplateSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectCertificateTemplate-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("state", flattenObjectCertificateTemplateState(o["state"], d, "state")); err != nil {
		if vv, ok := fortiAPIPatch(o["state"], "ObjectCertificateTemplate-State"); ok {
			if err = d.Set("state", vv); err != nil {
				return fmt.Errorf("Error reading state: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading state: %v", err)
		}
	}

	if err = d.Set("subject_alt_name", flattenObjectCertificateTemplateSubjectAltName(o["subject-alt-name"], d, "subject_alt_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["subject-alt-name"], "ObjectCertificateTemplate-SubjectAltName"); ok {
			if err = d.Set("subject_alt_name", vv); err != nil {
				return fmt.Errorf("Error reading subject_alt_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subject_alt_name: %v", err)
		}
	}

	if err = d.Set("subject_name", flattenObjectCertificateTemplateSubjectName(o["subject-name"], d, "subject_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["subject-name"], "ObjectCertificateTemplate-SubjectName"); ok {
			if err = d.Set("subject_name", vv); err != nil {
				return fmt.Errorf("Error reading subject_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subject_name: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCertificateTemplateType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCertificateTemplate-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("validity", flattenObjectCertificateTemplateValidity(o["validity"], d, "validity")); err != nil {
		if vv, ok := fortiAPIPatch(o["validity"], "ObjectCertificateTemplate-Validity"); ok {
			if err = d.Set("validity", vv); err != nil {
				return fmt.Errorf("Error reading validity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading validity: %v", err)
		}
	}

	return nil
}

func flattenObjectCertificateTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCertificateTemplateCity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateCurveName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateDigestType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateEnrollProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateEstCaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateEstClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCertificateTemplateEstHttpPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCertificateTemplateEstHttpUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateEstRegenerationMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateEstServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCertificateTemplateEstServerUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateEstSrpPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCertificateTemplateEstSrpUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateKeySize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateKeyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateOrganization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateOrganizationUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCertificateTemplateScepCaIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateScepPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCertificateTemplateScepServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateSubjectAltName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateSubjectName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCertificateTemplateValidity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCertificateTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("city"); ok || d.HasChange("city") {
		t, err := expandObjectCertificateTemplateCity(d, v, "city")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok || d.HasChange("country") {
		t, err := expandObjectCertificateTemplateCountry(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("curve_name"); ok || d.HasChange("curve_name") {
		t, err := expandObjectCertificateTemplateCurveName(d, v, "curve_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["curve-name"] = t
		}
	}

	if v, ok := d.GetOk("digest_type"); ok || d.HasChange("digest_type") {
		t, err := expandObjectCertificateTemplateDigestType(d, v, "digest_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-type"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok || d.HasChange("email") {
		t, err := expandObjectCertificateTemplateEmail(d, v, "email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}

	if v, ok := d.GetOk("enroll_protocol"); ok || d.HasChange("enroll_protocol") {
		t, err := expandObjectCertificateTemplateEnrollProtocol(d, v, "enroll_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enroll-protocol"] = t
		}
	}

	if v, ok := d.GetOk("est_ca_id"); ok || d.HasChange("est_ca_id") {
		t, err := expandObjectCertificateTemplateEstCaId(d, v, "est_ca_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-ca-id"] = t
		}
	}

	if v, ok := d.GetOk("est_client_cert"); ok || d.HasChange("est_client_cert") {
		t, err := expandObjectCertificateTemplateEstClientCert(d, v, "est_client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-client-cert"] = t
		}
	}

	if v, ok := d.GetOk("est_http_password"); ok || d.HasChange("est_http_password") {
		t, err := expandObjectCertificateTemplateEstHttpPassword(d, v, "est_http_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-http-password"] = t
		}
	}

	if v, ok := d.GetOk("est_http_username"); ok || d.HasChange("est_http_username") {
		t, err := expandObjectCertificateTemplateEstHttpUsername(d, v, "est_http_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-http-username"] = t
		}
	}

	if v, ok := d.GetOk("est_regeneration_method"); ok || d.HasChange("est_regeneration_method") {
		t, err := expandObjectCertificateTemplateEstRegenerationMethod(d, v, "est_regeneration_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-regeneration-method"] = t
		}
	}

	if v, ok := d.GetOk("est_server_cert"); ok || d.HasChange("est_server_cert") {
		t, err := expandObjectCertificateTemplateEstServerCert(d, v, "est_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("est_server_url"); ok || d.HasChange("est_server_url") {
		t, err := expandObjectCertificateTemplateEstServerUrl(d, v, "est_server_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-server-url"] = t
		}
	}

	if v, ok := d.GetOk("est_srp_password"); ok || d.HasChange("est_srp_password") {
		t, err := expandObjectCertificateTemplateEstSrpPassword(d, v, "est_srp_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-srp-password"] = t
		}
	}

	if v, ok := d.GetOk("est_srp_username"); ok || d.HasChange("est_srp_username") {
		t, err := expandObjectCertificateTemplateEstSrpUsername(d, v, "est_srp_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-srp-username"] = t
		}
	}

	if v, ok := d.GetOk("id_type"); ok || d.HasChange("id_type") {
		t, err := expandObjectCertificateTemplateIdType(d, v, "id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id-type"] = t
		}
	}

	if v, ok := d.GetOk("key_size"); ok || d.HasChange("key_size") {
		t, err := expandObjectCertificateTemplateKeySize(d, v, "key_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-size"] = t
		}
	}

	if v, ok := d.GetOk("key_type"); ok || d.HasChange("key_type") {
		t, err := expandObjectCertificateTemplateKeyType(d, v, "key_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCertificateTemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("organization"); ok || d.HasChange("organization") {
		t, err := expandObjectCertificateTemplateOrganization(d, v, "organization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization"] = t
		}
	}

	if v, ok := d.GetOk("organization_unit"); ok || d.HasChange("organization_unit") {
		t, err := expandObjectCertificateTemplateOrganizationUnit(d, v, "organization_unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization-unit"] = t
		}
	}

	if v, ok := d.GetOk("scep_ca_identifier"); ok || d.HasChange("scep_ca_identifier") {
		t, err := expandObjectCertificateTemplateScepCaIdentifier(d, v, "scep_ca_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-ca-identifier"] = t
		}
	}

	if v, ok := d.GetOk("scep_password"); ok || d.HasChange("scep_password") {
		t, err := expandObjectCertificateTemplateScepPassword(d, v, "scep_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-password"] = t
		}
	}

	if v, ok := d.GetOk("scep_server"); ok || d.HasChange("scep_server") {
		t, err := expandObjectCertificateTemplateScepServer(d, v, "scep_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectCertificateTemplateSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("state"); ok || d.HasChange("state") {
		t, err := expandObjectCertificateTemplateState(d, v, "state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	}

	if v, ok := d.GetOk("subject_alt_name"); ok || d.HasChange("subject_alt_name") {
		t, err := expandObjectCertificateTemplateSubjectAltName(d, v, "subject_alt_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject-alt-name"] = t
		}
	}

	if v, ok := d.GetOk("subject_name"); ok || d.HasChange("subject_name") {
		t, err := expandObjectCertificateTemplateSubjectName(d, v, "subject_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject-name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectCertificateTemplateType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("validity"); ok || d.HasChange("validity") {
		t, err := expandObjectCertificateTemplateValidity(d, v, "validity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["validity"] = t
		}
	}

	return &obj, nil
}
