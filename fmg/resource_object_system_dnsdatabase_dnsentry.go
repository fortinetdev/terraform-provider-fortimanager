// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: DNS entry.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemDnsDatabaseDnsEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemDnsDatabaseDnsEntryCreate,
		Read:   resourceObjectSystemDnsDatabaseDnsEntryRead,
		Update: resourceObjectSystemDnsDatabaseDnsEntryUpdate,
		Delete: resourceObjectSystemDnsDatabaseDnsEntryDelete,

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
			"dns_database": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"canonical_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preference": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemDnsDatabaseDnsEntryCreate(d *schema.ResourceData, m interface{}) error {
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

	dns_database := d.Get("dns_database").(string)
	paradict["dns_database"] = dns_database

	obj, err := getObjectObjectSystemDnsDatabaseDnsEntry(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemDnsDatabaseDnsEntry resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectSystemDnsDatabaseDnsEntry(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectSystemDnsDatabaseDnsEntry(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectSystemDnsDatabaseDnsEntry resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateObjectSystemDnsDatabaseDnsEntry(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectSystemDnsDatabaseDnsEntry resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceObjectSystemDnsDatabaseDnsEntryRead(d, m)
			} else {
				return fmt.Errorf("Error creating ObjectSystemDnsDatabaseDnsEntry resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemDnsDatabaseDnsEntryRead(d, m)
}

func resourceObjectSystemDnsDatabaseDnsEntryUpdate(d *schema.ResourceData, m interface{}) error {
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

	dns_database := d.Get("dns_database").(string)
	paradict["dns_database"] = dns_database

	obj, err := getObjectObjectSystemDnsDatabaseDnsEntry(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemDnsDatabaseDnsEntry resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemDnsDatabaseDnsEntry(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemDnsDatabaseDnsEntry resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectSystemDnsDatabaseDnsEntryRead(d, m)
}

func resourceObjectSystemDnsDatabaseDnsEntryDelete(d *schema.ResourceData, m interface{}) error {
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

	dns_database := d.Get("dns_database").(string)
	paradict["dns_database"] = dns_database

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemDnsDatabaseDnsEntry(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemDnsDatabaseDnsEntry resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemDnsDatabaseDnsEntryRead(d *schema.ResourceData, m interface{}) error {
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

	dns_database := d.Get("dns_database").(string)
	if dns_database == "" {
		dns_database = importOptionChecking(m.(*FortiClient).Cfg, "dns_database")
		if dns_database == "" {
			return fmt.Errorf("Parameter dns_database is missing")
		}
		if err = d.Set("dns_database", dns_database); err != nil {
			return fmt.Errorf("Error set params dns_database: %v", err)
		}
	}
	paradict["dns_database"] = dns_database

	o, err := c.ReadObjectSystemDnsDatabaseDnsEntry(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectSystemDnsDatabaseDnsEntry resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemDnsDatabaseDnsEntry(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemDnsDatabaseDnsEntry resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemDnsDatabaseDnsEntryCanonicalName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryHostname2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryIpv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryPreference2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemDnsDatabaseDnsEntryType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemDnsDatabaseDnsEntry(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("canonical_name", flattenObjectSystemDnsDatabaseDnsEntryCanonicalName2edl(o["canonical-name"], d, "canonical_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["canonical-name"], "ObjectSystemDnsDatabaseDnsEntry-CanonicalName"); ok {
			if err = d.Set("canonical_name", vv); err != nil {
				return fmt.Errorf("Error reading canonical_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading canonical_name: %v", err)
		}
	}

	if err = d.Set("hostname", flattenObjectSystemDnsDatabaseDnsEntryHostname2edl(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "ObjectSystemDnsDatabaseDnsEntry-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectSystemDnsDatabaseDnsEntryId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemDnsDatabaseDnsEntry-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenObjectSystemDnsDatabaseDnsEntryIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectSystemDnsDatabaseDnsEntry-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenObjectSystemDnsDatabaseDnsEntryIpv62edl(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "ObjectSystemDnsDatabaseDnsEntry-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("preference", flattenObjectSystemDnsDatabaseDnsEntryPreference2edl(o["preference"], d, "preference")); err != nil {
		if vv, ok := fortiAPIPatch(o["preference"], "ObjectSystemDnsDatabaseDnsEntry-Preference"); ok {
			if err = d.Set("preference", vv); err != nil {
				return fmt.Errorf("Error reading preference: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preference: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectSystemDnsDatabaseDnsEntryStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectSystemDnsDatabaseDnsEntry-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ttl", flattenObjectSystemDnsDatabaseDnsEntryTtl2edl(o["ttl"], d, "ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttl"], "ObjectSystemDnsDatabaseDnsEntry-Ttl"); ok {
			if err = d.Set("ttl", vv); err != nil {
				return fmt.Errorf("Error reading ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectSystemDnsDatabaseDnsEntryType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemDnsDatabaseDnsEntry-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemDnsDatabaseDnsEntryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemDnsDatabaseDnsEntryCanonicalName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryHostname2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryIpv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryPreference2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemDnsDatabaseDnsEntryType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemDnsDatabaseDnsEntry(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("canonical_name"); ok || d.HasChange("canonical_name") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryCanonicalName2edl(d, v, "canonical_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["canonical-name"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok || d.HasChange("hostname") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryHostname2edl(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryIpv62edl(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("preference"); ok || d.HasChange("preference") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryPreference2edl(d, v, "preference")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preference"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok || d.HasChange("ttl") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryTtl2edl(d, v, "ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSystemDnsDatabaseDnsEntryType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
