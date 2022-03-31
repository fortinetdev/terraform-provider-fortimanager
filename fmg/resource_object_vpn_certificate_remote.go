// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Remote certificate as a PEM file.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectVpnCertificateRemote() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnCertificateRemoteCreate,
		Read:   resourceObjectVpnCertificateRemoteRead,
		Update: resourceObjectVpnCertificateRemoteUpdate,
		Delete: resourceObjectVpnCertificateRemoteDelete,

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
			"remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVpnCertificateRemoteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnCertificateRemote(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnCertificateRemote resource while getting object: %v", err)
	}

	_, err = c.CreateObjectVpnCertificateRemote(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnCertificateRemote resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnCertificateRemoteRead(d, m)
}

func resourceObjectVpnCertificateRemoteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectVpnCertificateRemote(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnCertificateRemote resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVpnCertificateRemote(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnCertificateRemote resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnCertificateRemoteRead(d, m)
}

func resourceObjectVpnCertificateRemoteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectVpnCertificateRemote(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnCertificateRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnCertificateRemoteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectVpnCertificateRemote(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnCertificateRemote resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnCertificateRemote(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnCertificateRemote resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnCertificateRemoteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateRemoteRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateRemoteRemote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnCertificateRemoteSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnCertificateRemote(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectVpnCertificateRemoteName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnCertificateRemote-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenObjectVpnCertificateRemoteRange(o["range"], d, "range")); err != nil {
		if vv, ok := fortiAPIPatch(o["range"], "ObjectVpnCertificateRemote-Range"); ok {
			if err = d.Set("range", vv); err != nil {
				return fmt.Errorf("Error reading range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("remote", flattenObjectVpnCertificateRemoteRemote(o["remote"], d, "remote")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote"], "ObjectVpnCertificateRemote-Remote"); ok {
			if err = d.Set("remote", vv); err != nil {
				return fmt.Errorf("Error reading remote: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote: %v", err)
		}
	}

	if err = d.Set("source", flattenObjectVpnCertificateRemoteSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "ObjectVpnCertificateRemote-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnCertificateRemoteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnCertificateRemoteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateRemoteRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateRemoteRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnCertificateRemoteSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnCertificateRemote(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectVpnCertificateRemoteName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandObjectVpnCertificateRemoteRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("remote"); ok {
		t, err := expandObjectVpnCertificateRemoteRemote(d, v, "remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandObjectVpnCertificateRemoteSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
