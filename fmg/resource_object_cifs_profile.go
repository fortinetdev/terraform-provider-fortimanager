// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure CIFS profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectCifsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCifsProfileCreate,
		Read:   resourceObjectCifsProfileRead,
		Update: resourceObjectCifsProfileUpdate,
		Delete: resourceObjectCifsProfileDelete,

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
			"domain_controller": &schema.Schema{
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
			"server_credential_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_keytab": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keytab": &schema.Schema{
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
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectCifsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCifsProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCifsProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCifsProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCifsProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCifsProfileRead(d, m)
}

func resourceObjectCifsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCifsProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCifsProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCifsProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCifsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCifsProfileRead(d, m)
}

func resourceObjectCifsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectCifsProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCifsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCifsProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectCifsProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCifsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCifsProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCifsProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectCifsProfileDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCifsProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCifsProfileServerCredentialType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "none",
			2: "credential-replication",
			3: "credential-keytab",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectCifsProfileServerKeytab(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := i["keytab"]; ok {
			v := flattenObjectCifsProfileServerKeytabKeytab(i["keytab"], d, pre_append)
			tmp["keytab"] = fortiAPISubPartPatch(v, "ObjectCifsProfile-ServerKeytab-Keytab")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			v := flattenObjectCifsProfileServerKeytabPassword(i["password"], d, pre_append)
			tmp["password"] = fortiAPISubPartPatch(v, "ObjectCifsProfile-ServerKeytab-Password")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := i["principal"]; ok {
			v := flattenObjectCifsProfileServerKeytabPrincipal(i["principal"], d, pre_append)
			tmp["principal"] = fortiAPISubPartPatch(v, "ObjectCifsProfile-ServerKeytab-Principal")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectCifsProfileServerKeytabKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCifsProfileServerKeytabPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCifsProfileServerKeytabPrincipal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCifsProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("domain_controller", flattenObjectCifsProfileDomainController(o["domain-controller"], d, "domain_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-controller"], "ObjectCifsProfile-DomainController"); ok {
			if err = d.Set("domain_controller", vv); err != nil {
				return fmt.Errorf("Error reading domain_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_controller: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCifsProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCifsProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_credential_type", flattenObjectCifsProfileServerCredentialType(o["server-credential-type"], d, "server_credential_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-credential-type"], "ObjectCifsProfile-ServerCredentialType"); ok {
			if err = d.Set("server_credential_type", vv); err != nil {
				return fmt.Errorf("Error reading server_credential_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_credential_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_keytab", flattenObjectCifsProfileServerKeytab(o["server-keytab"], d, "server_keytab")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-keytab"], "ObjectCifsProfile-ServerKeytab"); ok {
				if err = d.Set("server_keytab", vv); err != nil {
					return fmt.Errorf("Error reading server_keytab: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_keytab: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_keytab"); ok {
			if err = d.Set("server_keytab", flattenObjectCifsProfileServerKeytab(o["server-keytab"], d, "server_keytab")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-keytab"], "ObjectCifsProfile-ServerKeytab"); ok {
					if err = d.Set("server_keytab", vv); err != nil {
						return fmt.Errorf("Error reading server_keytab: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_keytab: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectCifsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCifsProfileDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCifsProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCifsProfileServerCredentialType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCifsProfileServerKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keytab"], _ = expandObjectCifsProfileServerKeytabKeytab(d, i["keytab"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandObjectCifsProfileServerKeytabPassword(d, i["password"], pre_append)
		} else {
			tmp["password"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["principal"], _ = expandObjectCifsProfileServerKeytabPrincipal(d, i["principal"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectCifsProfileServerKeytabKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCifsProfileServerKeytabPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCifsProfileServerKeytabPrincipal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCifsProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain_controller"); ok {
		t, err := expandObjectCifsProfileDomainController(d, v, "domain_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-controller"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectCifsProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_credential_type"); ok {
		t, err := expandObjectCifsProfileServerCredentialType(d, v, "server_credential_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-credential-type"] = t
		}
	}

	if v, ok := d.GetOk("server_keytab"); ok {
		t, err := expandObjectCifsProfileServerKeytab(d, v, "server_keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-keytab"] = t
		}
	}

	return &obj, nil
}
