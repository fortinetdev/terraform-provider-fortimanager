// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Add web forward servers to a list to form a server group. Optionally assign weights to each server.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebProxyForwardServerGroupServerList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebProxyForwardServerGroupServerListCreate,
		Read:   resourceObjectWebProxyForwardServerGroupServerListRead,
		Update: resourceObjectWebProxyForwardServerGroupServerListUpdate,
		Delete: resourceObjectWebProxyForwardServerGroupServerListDelete,

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
			"forward_server_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWebProxyForwardServerGroupServerListCreate(d *schema.ResourceData, m interface{}) error {
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

	forward_server_group := d.Get("forward_server_group").(string)
	paradict["forward_server_group"] = forward_server_group

	obj, err := getObjectObjectWebProxyForwardServerGroupServerList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyForwardServerGroupServerList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWebProxyForwardServerGroupServerList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyForwardServerGroupServerList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebProxyForwardServerGroupServerListRead(d, m)
}

func resourceObjectWebProxyForwardServerGroupServerListUpdate(d *schema.ResourceData, m interface{}) error {
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

	forward_server_group := d.Get("forward_server_group").(string)
	paradict["forward_server_group"] = forward_server_group

	obj, err := getObjectObjectWebProxyForwardServerGroupServerList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyForwardServerGroupServerList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebProxyForwardServerGroupServerList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyForwardServerGroupServerList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebProxyForwardServerGroupServerListRead(d, m)
}

func resourceObjectWebProxyForwardServerGroupServerListDelete(d *schema.ResourceData, m interface{}) error {
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

	forward_server_group := d.Get("forward_server_group").(string)
	paradict["forward_server_group"] = forward_server_group

	wsParams["adom"] = adomv

	err = c.DeleteObjectWebProxyForwardServerGroupServerList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebProxyForwardServerGroupServerList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebProxyForwardServerGroupServerListRead(d *schema.ResourceData, m interface{}) error {
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

	forward_server_group := d.Get("forward_server_group").(string)
	if forward_server_group == "" {
		forward_server_group = importOptionChecking(m.(*FortiClient).Cfg, "forward_server_group")
		if forward_server_group == "" {
			return fmt.Errorf("Parameter forward_server_group is missing")
		}
		if err = d.Set("forward_server_group", forward_server_group); err != nil {
			return fmt.Errorf("Error set params forward_server_group: %v", err)
		}
	}
	paradict["forward_server_group"] = forward_server_group

	o, err := c.ReadObjectWebProxyForwardServerGroupServerList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyForwardServerGroupServerList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebProxyForwardServerGroupServerList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyForwardServerGroupServerList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebProxyForwardServerGroupServerListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectWebProxyForwardServerGroupServerListWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWebProxyForwardServerGroupServerList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectWebProxyForwardServerGroupServerListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebProxyForwardServerGroupServerList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectWebProxyForwardServerGroupServerListWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectWebProxyForwardServerGroupServerList-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectWebProxyForwardServerGroupServerListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebProxyForwardServerGroupServerListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectWebProxyForwardServerGroupServerListWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebProxyForwardServerGroupServerList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWebProxyForwardServerGroupServerListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectWebProxyForwardServerGroupServerListWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
