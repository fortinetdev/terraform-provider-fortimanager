// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Kerberos keytab entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserKrbKeytab() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserKrbKeytabCreate,
		Read:   resourceObjectUserKrbKeytabRead,
		Update: resourceObjectUserKrbKeytabUpdate,
		Delete: resourceObjectUserKrbKeytabDelete,

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
			"keytab": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"pac_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"principal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectUserKrbKeytabCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserKrbKeytab(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserKrbKeytab resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserKrbKeytab(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserKrbKeytab resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserKrbKeytabRead(d, m)
}

func resourceObjectUserKrbKeytabUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserKrbKeytab(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserKrbKeytab resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserKrbKeytab(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserKrbKeytab resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserKrbKeytabRead(d, m)
}

func resourceObjectUserKrbKeytabDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserKrbKeytab(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserKrbKeytab resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserKrbKeytabRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserKrbKeytab(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserKrbKeytab resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserKrbKeytab(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserKrbKeytab resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserKrbKeytabKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserKrbKeytabLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserKrbKeytabName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserKrbKeytabPacData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserKrbKeytabPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectUserKrbKeytabPrincipal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserKrbKeytab(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("keytab", flattenObjectUserKrbKeytabKeytab(o["keytab"], d, "keytab")); err != nil {
		if vv, ok := fortiAPIPatch(o["keytab"], "ObjectUserKrbKeytab-Keytab"); ok {
			if err = d.Set("keytab", vv); err != nil {
				return fmt.Errorf("Error reading keytab: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keytab: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenObjectUserKrbKeytabLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["ldap-server"], "ObjectUserKrbKeytab-LdapServer"); ok {
			if err = d.Set("ldap_server", vv); err != nil {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserKrbKeytabName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserKrbKeytab-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pac_data", flattenObjectUserKrbKeytabPacData(o["pac-data"], d, "pac_data")); err != nil {
		if vv, ok := fortiAPIPatch(o["pac-data"], "ObjectUserKrbKeytab-PacData"); ok {
			if err = d.Set("pac_data", vv); err != nil {
				return fmt.Errorf("Error reading pac_data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pac_data: %v", err)
		}
	}

	if err = d.Set("password", flattenObjectUserKrbKeytabPassword(o["password"], d, "password")); err != nil {
		if vv, ok := fortiAPIPatch(o["password"], "ObjectUserKrbKeytab-Password"); ok {
			if err = d.Set("password", vv); err != nil {
				return fmt.Errorf("Error reading password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("principal", flattenObjectUserKrbKeytabPrincipal(o["principal"], d, "principal")); err != nil {
		if vv, ok := fortiAPIPatch(o["principal"], "ObjectUserKrbKeytab-Principal"); ok {
			if err = d.Set("principal", vv); err != nil {
				return fmt.Errorf("Error reading principal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading principal: %v", err)
		}
	}

	return nil
}

func flattenObjectUserKrbKeytabFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserKrbKeytabKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserKrbKeytabLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandObjectUserKrbKeytabName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserKrbKeytabPacData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserKrbKeytabPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectUserKrbKeytabPrincipal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserKrbKeytab(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("keytab"); ok {
		t, err := expandObjectUserKrbKeytabKeytab(d, v, "keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keytab"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandObjectUserKrbKeytabLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectUserKrbKeytabName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pac_data"); ok {
		t, err := expandObjectUserKrbKeytabPacData(d, v, "pac_data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-data"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandObjectUserKrbKeytabPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("principal"); ok {
		t, err := expandObjectUserKrbKeytabPrincipal(d, v, "principal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["principal"] = t
		}
	}

	return &obj, nil
}
