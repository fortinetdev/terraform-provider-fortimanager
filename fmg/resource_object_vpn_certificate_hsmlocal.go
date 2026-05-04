// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Local certificates whose keys are stored on HSM.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnCertificateHsmLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnCertificateHsmLocalCreate,
		Read:   resourceObjectVpnCertificateHsmLocalRead,
		Update: resourceObjectVpnCertificateHsmLocalUpdate,
		Delete: resourceObjectVpnCertificateHsmLocalDelete,

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
			"api_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_cloud_service_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"gch_cryptokey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_cryptokey_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gch_cryptokey_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_keyring": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gch_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tmp_cert_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVpnCertificateHsmLocalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnCertificateHsmLocal(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnCertificateHsmLocal resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectVpnCertificateHsmLocal(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectVpnCertificateHsmLocal(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectVpnCertificateHsmLocal resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectVpnCertificateHsmLocal(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectVpnCertificateHsmLocal resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnCertificateHsmLocalRead(d, m)
}

func resourceObjectVpnCertificateHsmLocalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnCertificateHsmLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnCertificateHsmLocal resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVpnCertificateHsmLocal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnCertificateHsmLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnCertificateHsmLocalRead(d, m)
}

func resourceObjectVpnCertificateHsmLocalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectVpnCertificateHsmLocal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnCertificateHsmLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnCertificateHsmLocalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectVpnCertificateHsmLocal(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectVpnCertificateHsmLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnCertificateHsmLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnCertificateHsmLocal resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnCertificateHsmLocalApiVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalGchCloudServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnCertificateHsmLocalGchCryptokey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalGchCryptokeyAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalGchCryptokeyVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalGchKeyring(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalGchLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalGchProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalGchUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalTmpCertFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateHsmLocalVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnCertificateHsmLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("api_version", flattenObjectVpnCertificateHsmLocalApiVersion(o["api-version"], d, "api_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["api-version"], "ObjectVpnCertificateHsmLocal-ApiVersion"); ok {
			if err = d.Set("api_version", vv); err != nil {
				return fmt.Errorf("Error reading api_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading api_version: %v", err)
		}
	}

	if err = d.Set("certificate", flattenObjectVpnCertificateHsmLocalCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "ObjectVpnCertificateHsmLocal-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("comments", flattenObjectVpnCertificateHsmLocalComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectVpnCertificateHsmLocal-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("gch_cloud_service_name", flattenObjectVpnCertificateHsmLocalGchCloudServiceName(o["gch-cloud-service-name"], d, "gch_cloud_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cloud-service-name"], "ObjectVpnCertificateHsmLocal-GchCloudServiceName"); ok {
			if err = d.Set("gch_cloud_service_name", vv); err != nil {
				return fmt.Errorf("Error reading gch_cloud_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cloud_service_name: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey", flattenObjectVpnCertificateHsmLocalGchCryptokey(o["gch-cryptokey"], d, "gch_cryptokey")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey"], "ObjectVpnCertificateHsmLocal-GchCryptokey"); ok {
			if err = d.Set("gch_cryptokey", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey_algorithm", flattenObjectVpnCertificateHsmLocalGchCryptokeyAlgorithm(o["gch-cryptokey-algorithm"], d, "gch_cryptokey_algorithm")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey-algorithm"], "ObjectVpnCertificateHsmLocal-GchCryptokeyAlgorithm"); ok {
			if err = d.Set("gch_cryptokey_algorithm", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey_algorithm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey_algorithm: %v", err)
		}
	}

	if err = d.Set("gch_cryptokey_version", flattenObjectVpnCertificateHsmLocalGchCryptokeyVersion(o["gch-cryptokey-version"], d, "gch_cryptokey_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-cryptokey-version"], "ObjectVpnCertificateHsmLocal-GchCryptokeyVersion"); ok {
			if err = d.Set("gch_cryptokey_version", vv); err != nil {
				return fmt.Errorf("Error reading gch_cryptokey_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_cryptokey_version: %v", err)
		}
	}

	if err = d.Set("gch_keyring", flattenObjectVpnCertificateHsmLocalGchKeyring(o["gch-keyring"], d, "gch_keyring")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-keyring"], "ObjectVpnCertificateHsmLocal-GchKeyring"); ok {
			if err = d.Set("gch_keyring", vv); err != nil {
				return fmt.Errorf("Error reading gch_keyring: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_keyring: %v", err)
		}
	}

	if err = d.Set("gch_location", flattenObjectVpnCertificateHsmLocalGchLocation(o["gch-location"], d, "gch_location")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-location"], "ObjectVpnCertificateHsmLocal-GchLocation"); ok {
			if err = d.Set("gch_location", vv); err != nil {
				return fmt.Errorf("Error reading gch_location: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_location: %v", err)
		}
	}

	if err = d.Set("gch_project", flattenObjectVpnCertificateHsmLocalGchProject(o["gch-project"], d, "gch_project")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-project"], "ObjectVpnCertificateHsmLocal-GchProject"); ok {
			if err = d.Set("gch_project", vv); err != nil {
				return fmt.Errorf("Error reading gch_project: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_project: %v", err)
		}
	}

	if err = d.Set("gch_url", flattenObjectVpnCertificateHsmLocalGchUrl(o["gch-url"], d, "gch_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["gch-url"], "ObjectVpnCertificateHsmLocal-GchUrl"); ok {
			if err = d.Set("gch_url", vv); err != nil {
				return fmt.Errorf("Error reading gch_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gch_url: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnCertificateHsmLocalName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnCertificateHsmLocal-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenObjectVpnCertificateHsmLocalRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "ObjectVpnCertificateHsmLocal-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenObjectVpnCertificateHsmLocalSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "ObjectVpnCertificateHsmLocal-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("tmp_cert_file", flattenObjectVpnCertificateHsmLocalTmpCertFile(o["tmp-cert-file"], d, "tmp_cert_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["tmp-cert-file"], "ObjectVpnCertificateHsmLocal-TmpCertFile"); ok {
			if err = d.Set("tmp_cert_file", vv); err != nil {
				return fmt.Errorf("Error reading tmp_cert_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tmp_cert_file: %v", err)
		}
	}

	if err = d.Set("vendor", flattenObjectVpnCertificateHsmLocalVendor(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "ObjectVpnCertificateHsmLocal-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnCertificateHsmLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnCertificateHsmLocalApiVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalGchCloudServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnCertificateHsmLocalGchCryptokey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalGchCryptokeyAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalGchCryptokeyVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalGchKeyring(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalGchLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalGchProject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalGchUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalTmpCertFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateHsmLocalVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnCertificateHsmLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("api_version"); ok || d.HasChange("api_version") {
		t, err := expandObjectVpnCertificateHsmLocalApiVersion(d, v, "api_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-version"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandObjectVpnCertificateHsmLocalCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectVpnCertificateHsmLocalComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("gch_cloud_service_name"); ok || d.HasChange("gch_cloud_service_name") {
		t, err := expandObjectVpnCertificateHsmLocalGchCloudServiceName(d, v, "gch_cloud_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cloud-service-name"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey"); ok || d.HasChange("gch_cryptokey") {
		t, err := expandObjectVpnCertificateHsmLocalGchCryptokey(d, v, "gch_cryptokey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey_algorithm"); ok || d.HasChange("gch_cryptokey_algorithm") {
		t, err := expandObjectVpnCertificateHsmLocalGchCryptokeyAlgorithm(d, v, "gch_cryptokey_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("gch_cryptokey_version"); ok || d.HasChange("gch_cryptokey_version") {
		t, err := expandObjectVpnCertificateHsmLocalGchCryptokeyVersion(d, v, "gch_cryptokey_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-cryptokey-version"] = t
		}
	}

	if v, ok := d.GetOk("gch_keyring"); ok || d.HasChange("gch_keyring") {
		t, err := expandObjectVpnCertificateHsmLocalGchKeyring(d, v, "gch_keyring")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-keyring"] = t
		}
	}

	if v, ok := d.GetOk("gch_location"); ok || d.HasChange("gch_location") {
		t, err := expandObjectVpnCertificateHsmLocalGchLocation(d, v, "gch_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-location"] = t
		}
	}

	if v, ok := d.GetOk("gch_project"); ok || d.HasChange("gch_project") {
		t, err := expandObjectVpnCertificateHsmLocalGchProject(d, v, "gch_project")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-project"] = t
		}
	}

	if v, ok := d.GetOk("gch_url"); ok || d.HasChange("gch_url") {
		t, err := expandObjectVpnCertificateHsmLocalGchUrl(d, v, "gch_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gch-url"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnCertificateHsmLocalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok || d.HasChange("range") {
		t, err := expandObjectVpnCertificateHsmLocalRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandObjectVpnCertificateHsmLocalSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("tmp_cert_file"); ok || d.HasChange("tmp_cert_file") {
		t, err := expandObjectVpnCertificateHsmLocalTmpCertFile(d, v, "tmp_cert_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tmp-cert-file"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandObjectVpnCertificateHsmLocalVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
