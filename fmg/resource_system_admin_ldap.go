// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: LDAP server entry configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdminLdap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminLdapCreate,
		Read:   resourceSystemAdminLdapRead,
		Update: resourceSystemAdminLdapUpdate,
		Delete: resourceSystemAdminLdapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fmgadom": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adom_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"adom_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adom_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attributes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cnid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"memberof_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"profile_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemAdminLdapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemAdminLdap(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminLdap resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAdminLdap(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAdminLdap resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminLdapRead(d, m)
}

func resourceSystemAdminLdapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemAdminLdap(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminLdap resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdminLdap(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminLdap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminLdapRead(d, m)
}

func resourceSystemAdminLdapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemAdminLdap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdminLdap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminLdapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemAdminLdap(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminLdap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminLdap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminLdap resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminLdapAdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom_name"
		if _, ok := i["adom-name"]; ok {
			v := flattenSystemAdminLdapAdomAdomName(i["adom-name"], d, pre_append)
			tmp["adom_name"] = fortiAPISubPartPatch(v, "SystemAdminLdap-Adom-AdomName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAdminLdapAdomAdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapAdomAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapAdomAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapAttributes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapCnid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapConnectTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapDn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapMemberofAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminLdapPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapProfileAttr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapSecure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLdapUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdminLdap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("fmgadom", flattenSystemAdminLdapAdom(o["adom"], d, "fmgadom")); err != nil {
			if vv, ok := fortiAPIPatch(o["adom"], "SystemAdminLdap-Adom"); ok {
				if err = d.Set("fmgadom", vv); err != nil {
					return fmt.Errorf("Error reading fmgadom: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fmgadom"); ok {
			if err = d.Set("fmgadom", flattenSystemAdminLdapAdom(o["adom"], d, "fmgadom")); err != nil {
				if vv, ok := fortiAPIPatch(o["adom"], "SystemAdminLdap-Adom"); ok {
					if err = d.Set("fmgadom", vv); err != nil {
						return fmt.Errorf("Error reading fmgadom: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fmgadom: %v", err)
				}
			}
		}
	}

	if err = d.Set("adom_access", flattenSystemAdminLdapAdomAccess(o["adom-access"], d, "adom_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-access"], "SystemAdminLdap-AdomAccess"); ok {
			if err = d.Set("adom_access", vv); err != nil {
				return fmt.Errorf("Error reading adom_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_access: %v", err)
		}
	}

	if err = d.Set("adom_attr", flattenSystemAdminLdapAdomAttr(o["adom-attr"], d, "adom_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom-attr"], "SystemAdminLdap-AdomAttr"); ok {
			if err = d.Set("adom_attr", vv); err != nil {
				return fmt.Errorf("Error reading adom_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adom_attr: %v", err)
		}
	}

	if err = d.Set("attributes", flattenSystemAdminLdapAttributes(o["attributes"], d, "attributes")); err != nil {
		if vv, ok := fortiAPIPatch(o["attributes"], "SystemAdminLdap-Attributes"); ok {
			if err = d.Set("attributes", vv); err != nil {
				return fmt.Errorf("Error reading attributes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attributes: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenSystemAdminLdapCaCert(o["ca-cert"], d, "ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cert"], "SystemAdminLdap-CaCert"); ok {
			if err = d.Set("ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("cnid", flattenSystemAdminLdapCnid(o["cnid"], d, "cnid")); err != nil {
		if vv, ok := fortiAPIPatch(o["cnid"], "SystemAdminLdap-Cnid"); ok {
			if err = d.Set("cnid", vv); err != nil {
				return fmt.Errorf("Error reading cnid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cnid: %v", err)
		}
	}

	if err = d.Set("connect_timeout", flattenSystemAdminLdapConnectTimeout(o["connect-timeout"], d, "connect_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["connect-timeout"], "SystemAdminLdap-ConnectTimeout"); ok {
			if err = d.Set("connect_timeout", vv); err != nil {
				return fmt.Errorf("Error reading connect_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connect_timeout: %v", err)
		}
	}

	if err = d.Set("dn", flattenSystemAdminLdapDn(o["dn"], d, "dn")); err != nil {
		if vv, ok := fortiAPIPatch(o["dn"], "SystemAdminLdap-Dn"); ok {
			if err = d.Set("dn", vv); err != nil {
				return fmt.Errorf("Error reading dn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dn: %v", err)
		}
	}

	if err = d.Set("filter", flattenSystemAdminLdapFilter(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "SystemAdminLdap-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("group", flattenSystemAdminLdapGroup(o["group"], d, "group")); err != nil {
		if vv, ok := fortiAPIPatch(o["group"], "SystemAdminLdap-Group"); ok {
			if err = d.Set("group", vv); err != nil {
				return fmt.Errorf("Error reading group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("memberof_attr", flattenSystemAdminLdapMemberofAttr(o["memberof-attr"], d, "memberof_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["memberof-attr"], "SystemAdminLdap-MemberofAttr"); ok {
			if err = d.Set("memberof_attr", vv); err != nil {
				return fmt.Errorf("Error reading memberof_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading memberof_attr: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAdminLdapName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAdminLdap-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAdminLdapPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemAdminLdap-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("profile_attr", flattenSystemAdminLdapProfileAttr(o["profile-attr"], d, "profile_attr")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-attr"], "SystemAdminLdap-ProfileAttr"); ok {
			if err = d.Set("profile_attr", vv); err != nil {
				return fmt.Errorf("Error reading profile_attr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_attr: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenSystemAdminLdapSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "SystemAdminLdap-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("secure", flattenSystemAdminLdapSecure(o["secure"], d, "secure")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure"], "SystemAdminLdap-Secure"); ok {
			if err = d.Set("secure", vv); err != nil {
				return fmt.Errorf("Error reading secure: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemAdminLdapServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemAdminLdap-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenSystemAdminLdapTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "SystemAdminLdap-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemAdminLdapType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemAdminLdap-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemAdminLdapUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemAdminLdap-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminLdapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminLdapAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adom_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adom-name"], _ = expandSystemAdminLdapAdomAdomName(d, i["adom_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAdminLdapAdomAdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapAdomAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapAdomAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapAttributes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapCnid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapConnectTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapDn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapMemberofAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminLdapPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapProfileAttr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapSecure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLdapUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdminLdap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandSystemAdminLdapAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("adom_access"); ok || d.HasChange("adom_access") {
		t, err := expandSystemAdminLdapAdomAccess(d, v, "adom_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-access"] = t
		}
	}

	if v, ok := d.GetOk("adom_attr"); ok || d.HasChange("adom_attr") {
		t, err := expandSystemAdminLdapAdomAttr(d, v, "adom_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom-attr"] = t
		}
	}

	if v, ok := d.GetOk("attributes"); ok || d.HasChange("attributes") {
		t, err := expandSystemAdminLdapAttributes(d, v, "attributes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attributes"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok || d.HasChange("ca_cert") {
		t, err := expandSystemAdminLdapCaCert(d, v, "ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("cnid"); ok || d.HasChange("cnid") {
		t, err := expandSystemAdminLdapCnid(d, v, "cnid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cnid"] = t
		}
	}

	if v, ok := d.GetOk("connect_timeout"); ok || d.HasChange("connect_timeout") {
		t, err := expandSystemAdminLdapConnectTimeout(d, v, "connect_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dn"); ok || d.HasChange("dn") {
		t, err := expandSystemAdminLdapDn(d, v, "dn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dn"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandSystemAdminLdapFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("group"); ok || d.HasChange("group") {
		t, err := expandSystemAdminLdapGroup(d, v, "group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("memberof_attr"); ok || d.HasChange("memberof_attr") {
		t, err := expandSystemAdminLdapMemberofAttr(d, v, "memberof_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memberof-attr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAdminLdapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemAdminLdapPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemAdminLdapPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("profile_attr"); ok || d.HasChange("profile_attr") {
		t, err := expandSystemAdminLdapProfileAttr(d, v, "profile_attr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-attr"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandSystemAdminLdapSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok || d.HasChange("secure") {
		t, err := expandSystemAdminLdapSecure(d, v, "secure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemAdminLdapServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok || d.HasChange("tertiary_server") {
		t, err := expandSystemAdminLdapTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemAdminLdapType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemAdminLdapUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
