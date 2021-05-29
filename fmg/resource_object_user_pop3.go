// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: POP3 server entry configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserPop3() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserPop3Create,
		Read:   resourceObjectUserPop3Read,
		Update: resourceObjectUserPop3Update,
		Delete: resourceObjectUserPop3Delete,

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
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserPop3Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserPop3(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserPop3 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserPop3(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserPop3 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserPop3Read(d, m)
}

func resourceObjectUserPop3Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserPop3(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserPop3 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserPop3(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserPop3 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserPop3Read(d, m)
}

func resourceObjectUserPop3Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserPop3(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserPop3 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserPop3Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserPop3(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserPop3 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserPop3(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserPop3 resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserPop3Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPop3Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPop3Secure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "none",
			1: "starttls",
			4: "pop3s",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectUserPop3Server(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserPop3SslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "default",
			1: "TLSv1",
			2: "TLSv1-1",
			4: "TLSv1-2",
			8: "SSLv3",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectObjectUserPop3(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenObjectUserPop3Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserPop3-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectUserPop3Port(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectUserPop3-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secure", flattenObjectUserPop3Secure(o["secure"], d, "secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure"], "ObjectUserPop3-Secure"); ok {
			if err = d.Set("secure", vv); err != nil {
				return fmt.Errorf("Error reading secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("server", flattenObjectUserPop3Server(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "ObjectUserPop3-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenObjectUserPop3SslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-min-proto-version"], "ObjectUserPop3-SslMinProtoVersion"); ok {
			if err = d.Set("ssl_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	return nil
}

func flattenObjectUserPop3FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserPop3Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPop3Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPop3Secure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPop3Server(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserPop3SslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserPop3(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserPop3Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandObjectUserPop3Port(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok {
		t, err := expandObjectUserPop3Secure(d, v, "secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandObjectUserPop3Server(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandObjectUserPop3SslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	return &obj, nil
}
