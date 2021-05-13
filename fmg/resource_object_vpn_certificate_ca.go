// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CA certificate.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVpnCertificateCa() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnCertificateCaCreate,
		Read:   resourceObjectVpnCertificateCaRead,
		Update: resourceObjectVpnCertificateCaUpdate,
		Delete: resourceObjectVpnCertificateCaDelete,

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
			"_private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_update_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_update_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_inspection_trusted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVpnCertificateCaCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnCertificateCa resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnCertificateCa(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnCertificateCa resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnCertificateCaRead(d, m)
}

func resourceObjectVpnCertificateCaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnCertificateCa resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnCertificateCa(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnCertificateCa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnCertificateCaRead(d, m)
}

func resourceObjectVpnCertificateCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVpnCertificateCa(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnCertificateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnCertificateCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVpnCertificateCa(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnCertificateCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnCertificateCa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnCertificateCa resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnCertificateCaPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateCaAutoUpdateDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateCaAutoUpdateDaysWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateCaCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateCaLastUpdated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateCaName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateCaRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "global",
			1: "vdom",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnCertificateCaScepUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateCaSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "factory",
			1: "user",
			2: "bundle",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectVpnCertificateCaSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateCaSslInspectionTrusted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectVpnCertificateCa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("_private_key", flattenObjectVpnCertificateCaPrivateKey(o["_private_key"], d, "_private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["_private_key"], "ObjectVpnCertificateCa-PrivateKey"); ok {
			if err = d.Set("_private_key", vv); err != nil {
				return fmt.Errorf("Error reading _private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _private_key: %v", err)
		}
	}

	if err = d.Set("auto_update_days", flattenObjectVpnCertificateCaAutoUpdateDays(o["auto-update-days"], d, "auto_update_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-update-days"], "ObjectVpnCertificateCa-AutoUpdateDays"); ok {
			if err = d.Set("auto_update_days", vv); err != nil {
				return fmt.Errorf("Error reading auto_update_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_update_days: %v", err)
		}
	}

	if err = d.Set("auto_update_days_warning", flattenObjectVpnCertificateCaAutoUpdateDaysWarning(o["auto-update-days-warning"], d, "auto_update_days_warning")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-update-days-warning"], "ObjectVpnCertificateCa-AutoUpdateDaysWarning"); ok {
			if err = d.Set("auto_update_days_warning", vv); err != nil {
				return fmt.Errorf("Error reading auto_update_days_warning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_update_days_warning: %v", err)
		}
	}

	if err = d.Set("ca", flattenObjectVpnCertificateCaCa(o["ca"], d, "ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca"], "ObjectVpnCertificateCa-Ca"); ok {
			if err = d.Set("ca", vv); err != nil {
				return fmt.Errorf("Error reading ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenObjectVpnCertificateCaLastUpdated(o["last-updated"], d, "last_updated")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-updated"], "ObjectVpnCertificateCa-LastUpdated"); ok {
			if err = d.Set("last_updated", vv); err != nil {
				return fmt.Errorf("Error reading last_updated: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnCertificateCaName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnCertificateCa-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenObjectVpnCertificateCaRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "ObjectVpnCertificateCa-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenObjectVpnCertificateCaScepUrl(o["scep-url"], d, "scep_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["scep-url"], "ObjectVpnCertificateCa-ScepUrl"); ok {
			if err = d.Set("scep_url", vv); err != nil {
				return fmt.Errorf("Error reading scep_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("source", flattenObjectVpnCertificateCaSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "ObjectVpnCertificateCa-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenObjectVpnCertificateCaSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "ObjectVpnCertificateCa-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_inspection_trusted", flattenObjectVpnCertificateCaSslInspectionTrusted(o["ssl-inspection-trusted"], d, "ssl_inspection_trusted")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-inspection-trusted"], "ObjectVpnCertificateCa-SslInspectionTrusted"); ok {
			if err = d.Set("ssl_inspection_trusted", vv); err != nil {
				return fmt.Errorf("Error reading ssl_inspection_trusted: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_inspection_trusted: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnCertificateCaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnCertificateCaPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaAutoUpdateDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaAutoUpdateDaysWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaLastUpdated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaScepUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateCaSslInspectionTrusted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnCertificateCa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_private_key"); ok {
		t, err := expandObjectVpnCertificateCaPrivateKey(d, v, "_private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_private_key"] = t
		}
	}

	if v, ok := d.GetOk("auto_update_days"); ok {
		t, err := expandObjectVpnCertificateCaAutoUpdateDays(d, v, "auto_update_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days"] = t
		}
	}

	if v, ok := d.GetOk("auto_update_days_warning"); ok {
		t, err := expandObjectVpnCertificateCaAutoUpdateDaysWarning(d, v, "auto_update_days_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandObjectVpnCertificateCaCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("last_updated"); ok {
		t, err := expandObjectVpnCertificateCaLastUpdated(d, v, "last_updated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectVpnCertificateCaName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandObjectVpnCertificateCaRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandObjectVpnCertificateCaScepUrl(d, v, "scep_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandObjectVpnCertificateCaSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandObjectVpnCertificateCaSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ssl_inspection_trusted"); ok {
		t, err := expandObjectVpnCertificateCaSslInspectionTrusted(d, v, "ssl_inspection_trusted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-inspection-trusted"] = t
		}
	}

	return &obj, nil
}
