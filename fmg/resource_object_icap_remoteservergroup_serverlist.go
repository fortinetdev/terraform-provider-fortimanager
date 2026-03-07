// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectIcap RemoteServerGroupServerList

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectIcapRemoteServerGroupServerList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIcapRemoteServerGroupServerListCreate,
		Read:   resourceObjectIcapRemoteServerGroupServerListRead,
		Update: resourceObjectIcapRemoteServerGroupServerListUpdate,
		Delete: resourceObjectIcapRemoteServerGroupServerListDelete,

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
			"remote_server_group": &schema.Schema{
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

func resourceObjectIcapRemoteServerGroupServerListCreate(d *schema.ResourceData, m interface{}) error {
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

	remote_server_group := d.Get("remote_server_group").(string)
	paradict["remote_server_group"] = remote_server_group

	obj, err := getObjectObjectIcapRemoteServerGroupServerList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapRemoteServerGroupServerList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectIcapRemoteServerGroupServerList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapRemoteServerGroupServerList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIcapRemoteServerGroupServerListRead(d, m)
}

func resourceObjectIcapRemoteServerGroupServerListUpdate(d *schema.ResourceData, m interface{}) error {
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

	remote_server_group := d.Get("remote_server_group").(string)
	paradict["remote_server_group"] = remote_server_group

	obj, err := getObjectObjectIcapRemoteServerGroupServerList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapRemoteServerGroupServerList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectIcapRemoteServerGroupServerList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapRemoteServerGroupServerList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectIcapRemoteServerGroupServerListRead(d, m)
}

func resourceObjectIcapRemoteServerGroupServerListDelete(d *schema.ResourceData, m interface{}) error {
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

	remote_server_group := d.Get("remote_server_group").(string)
	paradict["remote_server_group"] = remote_server_group

	wsParams["adom"] = adomv

	err = c.DeleteObjectIcapRemoteServerGroupServerList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIcapRemoteServerGroupServerList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIcapRemoteServerGroupServerListRead(d *schema.ResourceData, m interface{}) error {
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

	remote_server_group := d.Get("remote_server_group").(string)
	if remote_server_group == "" {
		remote_server_group = importOptionChecking(m.(*FortiClient).Cfg, "remote_server_group")
		if remote_server_group == "" {
			return fmt.Errorf("Parameter remote_server_group is missing")
		}
		if err = d.Set("remote_server_group", remote_server_group); err != nil {
			return fmt.Errorf("Error set params remote_server_group: %v", err)
		}
	}
	paradict["remote_server_group"] = remote_server_group

	o, err := c.ReadObjectIcapRemoteServerGroupServerList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapRemoteServerGroupServerList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIcapRemoteServerGroupServerList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapRemoteServerGroupServerList resource from API: %v", err)
	}
	return nil
}

func flattenObjectIcapRemoteServerGroupServerListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapRemoteServerGroupServerListWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectIcapRemoteServerGroupServerList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectIcapRemoteServerGroupServerListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectIcapRemoteServerGroupServerList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("weight", flattenObjectIcapRemoteServerGroupServerListWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ObjectIcapRemoteServerGroupServerList-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenObjectIcapRemoteServerGroupServerListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIcapRemoteServerGroupServerListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapRemoteServerGroupServerListWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIcapRemoteServerGroupServerList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectIcapRemoteServerGroupServerListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandObjectIcapRemoteServerGroupServerListWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
