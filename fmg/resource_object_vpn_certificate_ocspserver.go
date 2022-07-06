// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OCSP server configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVpnCertificateOcspServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnCertificateOcspServerCreate,
		Read:   resourceObjectVpnCertificateOcspServerRead,
		Update: resourceObjectVpnCertificateOcspServerUpdate,
		Delete: resourceObjectVpnCertificateOcspServerDelete,

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
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"secondary_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unavail_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVpnCertificateOcspServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnCertificateOcspServer(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnCertificateOcspServer resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnCertificateOcspServer(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnCertificateOcspServer resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnCertificateOcspServerRead(d, m)
}

func resourceObjectVpnCertificateOcspServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnCertificateOcspServer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnCertificateOcspServer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnCertificateOcspServer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnCertificateOcspServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnCertificateOcspServerRead(d, m)
}

func resourceObjectVpnCertificateOcspServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVpnCertificateOcspServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnCertificateOcspServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnCertificateOcspServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVpnCertificateOcspServer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnCertificateOcspServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnCertificateOcspServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnCertificateOcspServer resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnCertificateOcspServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateOcspServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateOcspServerSecondaryCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateOcspServerSecondaryUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateOcspServerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateOcspServerUnavailAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateOcspServerUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnCertificateOcspServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("cert", flattenObjectVpnCertificateOcspServerCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "ObjectVpnCertificateOcspServer-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnCertificateOcspServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnCertificateOcspServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("secondary_cert", flattenObjectVpnCertificateOcspServerSecondaryCert(o["secondary-cert"], d, "secondary_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-cert"], "ObjectVpnCertificateOcspServer-SecondaryCert"); ok {
			if err = d.Set("secondary_cert", vv); err != nil {
				return fmt.Errorf("Error reading secondary_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_cert: %v", err)
		}
	}

	if err = d.Set("secondary_url", flattenObjectVpnCertificateOcspServerSecondaryUrl(o["secondary-url"], d, "secondary_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-url"], "ObjectVpnCertificateOcspServer-SecondaryUrl"); ok {
			if err = d.Set("secondary_url", vv); err != nil {
				return fmt.Errorf("Error reading secondary_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_url: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectVpnCertificateOcspServerSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectVpnCertificateOcspServer-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("unavail_action", flattenObjectVpnCertificateOcspServerUnavailAction(o["unavail-action"], d, "unavail_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["unavail-action"], "ObjectVpnCertificateOcspServer-UnavailAction"); ok {
			if err = d.Set("unavail_action", vv); err != nil {
				return fmt.Errorf("Error reading unavail_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unavail_action: %v", err)
		}
	}

	if err = d.Set("url", flattenObjectVpnCertificateOcspServerUrl(o["url"], d, "url")); err != nil {
		if vv, ok := fortiAPIPatch(o["url"], "ObjectVpnCertificateOcspServer-Url"); ok {
			if err = d.Set("url", vv); err != nil {
				return fmt.Errorf("Error reading url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnCertificateOcspServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnCertificateOcspServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateOcspServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateOcspServerSecondaryCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateOcspServerSecondaryUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateOcspServerSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateOcspServerUnavailAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateOcspServerUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnCertificateOcspServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandObjectVpnCertificateOcspServerCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnCertificateOcspServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("secondary_cert"); ok || d.HasChange("secondary_cert") {
		t, err := expandObjectVpnCertificateOcspServerSecondaryCert(d, v, "secondary_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-cert"] = t
		}
	}

	if v, ok := d.GetOk("secondary_url"); ok || d.HasChange("secondary_url") {
		t, err := expandObjectVpnCertificateOcspServerSecondaryUrl(d, v, "secondary_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-url"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandObjectVpnCertificateOcspServerSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("unavail_action"); ok || d.HasChange("unavail_action") {
		t, err := expandObjectVpnCertificateOcspServerUnavailAction(d, v, "unavail_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unavail-action"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok || d.HasChange("url") {
		t, err := expandObjectVpnCertificateOcspServerUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
