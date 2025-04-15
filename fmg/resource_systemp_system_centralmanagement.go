// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure central management.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemCentralManagementUpdate,
		Read:   resourceSystempSystemCentralManagementRead,
		Update: resourceSystempSystemCentralManagementUpdate,
		Delete: resourceSystempSystemCentralManagementDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"include_default_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_address6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceSystempSystemCentralManagementUpdate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemCentralManagement(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemCentralManagement resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystempSystemCentralManagement(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemCentralManagement resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempSystemCentralManagement")

	return resourceSystempSystemCentralManagementRead(d, m)
}

func resourceSystempSystemCentralManagementDelete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	wsParams["adom"] = adomv

	err = c.DeleteSystempSystemCentralManagement(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemCentralManagement resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemCentralManagementRead(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	paradict["devprof"] = devprof

	o, err := c.ReadSystempSystemCentralManagement(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemCentralManagement resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemCentralManagement(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemCentralManagement resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemCentralManagementIncludeDefaultServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			v := flattenSystempSystemCentralManagementServerListAddrType(i["addr-type"], d, pre_append)
			tmp["addr_type"] = fortiAPISubPartPatch(v, "SystempSystemCentralManagement-ServerList-AddrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			v := flattenSystempSystemCentralManagementServerListFqdn(i["fqdn"], d, pre_append)
			tmp["fqdn"] = fortiAPISubPartPatch(v, "SystempSystemCentralManagement-ServerList-Fqdn")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystempSystemCentralManagementServerListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystempSystemCentralManagement-ServerList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := i["server-address"]; ok {
			v := flattenSystempSystemCentralManagementServerListServerAddress(i["server-address"], d, pre_append)
			tmp["server_address"] = fortiAPISubPartPatch(v, "SystempSystemCentralManagement-ServerList-ServerAddress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := i["server-address6"]; ok {
			v := flattenSystempSystemCentralManagementServerListServerAddress6(i["server-address6"], d, pre_append)
			tmp["server_address6"] = fortiAPISubPartPatch(v, "SystempSystemCentralManagement-ServerList-ServerAddress6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			v := flattenSystempSystemCentralManagementServerListServerType(i["server-type"], d, pre_append)
			tmp["server_type"] = fortiAPISubPartPatch(v, "SystempSystemCentralManagement-ServerList-ServerType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystempSystemCentralManagementServerListAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListServerAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystempSystemCentralManagement(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("include_default_servers", flattenSystempSystemCentralManagementIncludeDefaultServers(o["include-default-servers"], d, "include_default_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["include-default-servers"], "SystempSystemCentralManagement-IncludeDefaultServers"); ok {
			if err = d.Set("include_default_servers", vv); err != nil {
				return fmt.Errorf("Error reading include_default_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading include_default_servers: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_list", flattenSystempSystemCentralManagementServerList(o["server-list"], d, "server_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-list"], "SystempSystemCentralManagement-ServerList"); ok {
				if err = d.Set("server_list", vv); err != nil {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_list"); ok {
			if err = d.Set("server_list", flattenSystempSystemCentralManagementServerList(o["server-list"], d, "server_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-list"], "SystempSystemCentralManagement-ServerList"); ok {
					if err = d.Set("server_list", vv); err != nil {
						return fmt.Errorf("Error reading server_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystempSystemCentralManagementFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemCentralManagementIncludeDefaultServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-type"], _ = expandSystempSystemCentralManagementServerListAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fqdn"], _ = expandSystempSystemCentralManagementServerListFqdn(d, i["fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystempSystemCentralManagementServerListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-address"], _ = expandSystempSystemCentralManagementServerListServerAddress(d, i["server_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-address6"], _ = expandSystempSystemCentralManagementServerListServerAddress6(d, i["server_address6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-type"], _ = expandSystempSystemCentralManagementServerListServerType(d, i["server_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystempSystemCentralManagementServerListAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListServerAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListServerAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystempSystemCentralManagement(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("include_default_servers"); ok || d.HasChange("include_default_servers") {
		t, err := expandSystempSystemCentralManagementIncludeDefaultServers(d, v, "include_default_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-default-servers"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok || d.HasChange("server_list") {
		t, err := expandSystempSystemCentralManagementServerList(d, v, "server_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	return &obj, nil
}
