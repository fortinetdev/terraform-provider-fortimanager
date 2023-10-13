// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemCentralManagementServerList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemCentralManagementServerListCreate,
		Read:   resourceSystempSystemCentralManagementServerListRead,
		Update: resourceSystempSystemCentralManagementServerListUpdate,
		Delete: resourceSystempSystemCentralManagementServerListDelete,

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
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
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
	}
}

func resourceSystempSystemCentralManagementServerListCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemCentralManagementServerList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystempSystemCentralManagementServerList resource while getting object: %v", err)
	}

	_, err = c.CreateSystempSystemCentralManagementServerList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystempSystemCentralManagementServerList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystempSystemCentralManagementServerListRead(d, m)
}

func resourceSystempSystemCentralManagementServerListUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemCentralManagementServerList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemCentralManagementServerList resource while getting object: %v", err)
	}

	_, err = c.UpdateSystempSystemCentralManagementServerList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemCentralManagementServerList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystempSystemCentralManagementServerListRead(d, m)
}

func resourceSystempSystemCentralManagementServerListDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["devprof"] = devprof

	err = c.DeleteSystempSystemCentralManagementServerList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemCentralManagementServerList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemCentralManagementServerListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystempSystemCentralManagementServerList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemCentralManagementServerList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemCentralManagementServerList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemCentralManagementServerList resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemCentralManagementServerListAddrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListServerAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListServerAddress62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemCentralManagementServerListServerType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystempSystemCentralManagementServerList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("addr_type", flattenSystempSystemCentralManagementServerListAddrType2edl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "SystempSystemCentralManagementServerList-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenSystempSystemCentralManagementServerListFqdn2edl(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "SystempSystemCentralManagementServerList-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystempSystemCentralManagementServerListId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystempSystemCentralManagementServerList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("server_address", flattenSystempSystemCentralManagementServerListServerAddress2edl(o["server-address"], d, "server_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-address"], "SystempSystemCentralManagementServerList-ServerAddress"); ok {
			if err = d.Set("server_address", vv); err != nil {
				return fmt.Errorf("Error reading server_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_address: %v", err)
		}
	}

	if err = d.Set("server_address6", flattenSystempSystemCentralManagementServerListServerAddress62edl(o["server-address6"], d, "server_address6")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-address6"], "SystempSystemCentralManagementServerList-ServerAddress6"); ok {
			if err = d.Set("server_address6", vv); err != nil {
				return fmt.Errorf("Error reading server_address6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_address6: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystempSystemCentralManagementServerListServerType2edl(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystempSystemCentralManagementServerList-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemCentralManagementServerListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemCentralManagementServerListAddrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListServerAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListServerAddress62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemCentralManagementServerListServerType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystempSystemCentralManagementServerList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandSystempSystemCentralManagementServerListAddrType2edl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandSystempSystemCentralManagementServerListFqdn2edl(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystempSystemCentralManagementServerListId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("server_address"); ok || d.HasChange("server_address") {
		t, err := expandSystempSystemCentralManagementServerListServerAddress2edl(d, v, "server_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-address"] = t
		}
	}

	if v, ok := d.GetOk("server_address6"); ok || d.HasChange("server_address6") {
		t, err := expandSystempSystemCentralManagementServerListServerAddress62edl(d, v, "server_address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-address6"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystempSystemCentralManagementServerListServerType2edl(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	return &obj, nil
}
