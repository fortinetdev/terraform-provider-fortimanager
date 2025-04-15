// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Azure route.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemSdnConnectorRouteTableRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemSdnConnectorRouteTableRouteCreate,
		Read:   resourceObjectSystemSdnConnectorRouteTableRouteRead,
		Update: resourceObjectSystemSdnConnectorRouteTableRouteUpdate,
		Delete: resourceObjectSystemSdnConnectorRouteTableRouteDelete,

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
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route_table": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"next_hop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemSdnConnectorRouteTableRouteCreate(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	route_table := d.Get("route_table").(string)
	paradict["sdn_connector"] = sdn_connector
	paradict["route_table"] = route_table

	obj, err := getObjectObjectSystemSdnConnectorRouteTableRoute(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorRouteTableRoute resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSystemSdnConnectorRouteTableRoute(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemSdnConnectorRouteTableRoute resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemSdnConnectorRouteTableRouteRead(d, m)
}

func resourceObjectSystemSdnConnectorRouteTableRouteUpdate(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	route_table := d.Get("route_table").(string)
	paradict["sdn_connector"] = sdn_connector
	paradict["route_table"] = route_table

	obj, err := getObjectObjectSystemSdnConnectorRouteTableRoute(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorRouteTableRoute resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemSdnConnectorRouteTableRoute(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemSdnConnectorRouteTableRoute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemSdnConnectorRouteTableRouteRead(d, m)
}

func resourceObjectSystemSdnConnectorRouteTableRouteDelete(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	route_table := d.Get("route_table").(string)
	paradict["sdn_connector"] = sdn_connector
	paradict["route_table"] = route_table

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemSdnConnectorRouteTableRoute(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemSdnConnectorRouteTableRoute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemSdnConnectorRouteTableRouteRead(d *schema.ResourceData, m interface{}) error {
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

	sdn_connector := d.Get("sdn_connector").(string)
	route_table := d.Get("route_table").(string)
	if sdn_connector == "" {
		sdn_connector = importOptionChecking(m.(*FortiClient).Cfg, "sdn_connector")
		if sdn_connector == "" {
			return fmt.Errorf("Parameter sdn_connector is missing")
		}
		if err = d.Set("sdn_connector", sdn_connector); err != nil {
			return fmt.Errorf("Error set params sdn_connector: %v", err)
		}
	}
	if route_table == "" {
		route_table = importOptionChecking(m.(*FortiClient).Cfg, "route_table")
		if route_table == "" {
			return fmt.Errorf("Parameter route_table is missing")
		}
		if err = d.Set("route_table", route_table); err != nil {
			return fmt.Errorf("Error set params route_table: %v", err)
		}
	}
	paradict["sdn_connector"] = sdn_connector
	paradict["route_table"] = route_table

	o, err := c.ReadObjectSystemSdnConnectorRouteTableRoute(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorRouteTableRoute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemSdnConnectorRouteTableRoute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemSdnConnectorRouteTableRoute resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemSdnConnectorRouteTableRouteName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemSdnConnectorRouteTableRouteNextHop3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemSdnConnectorRouteTableRoute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectSystemSdnConnectorRouteTableRouteName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemSdnConnectorRouteTableRoute-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("next_hop", flattenObjectSystemSdnConnectorRouteTableRouteNextHop3rdl(o["next-hop"], d, "next_hop")); err != nil {
		if vv, ok := fortiAPIPatch(o["next-hop"], "ObjectSystemSdnConnectorRouteTableRoute-NextHop"); ok {
			if err = d.Set("next_hop", vv); err != nil {
				return fmt.Errorf("Error reading next_hop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading next_hop: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemSdnConnectorRouteTableRouteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemSdnConnectorRouteTableRouteName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemSdnConnectorRouteTableRouteNextHop3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemSdnConnectorRouteTableRoute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemSdnConnectorRouteTableRouteName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("next_hop"); ok || d.HasChange("next_hop") {
		t, err := expandObjectSystemSdnConnectorRouteTableRouteNextHop3rdl(d, v, "next_hop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop"] = t
		}
	}

	return &obj, nil
}
