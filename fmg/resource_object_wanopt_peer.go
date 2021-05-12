// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WAN optimization peers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWanoptPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWanoptPeerCreate,
		Read:   resourceObjectWanoptPeerRead,
		Update: resourceObjectWanoptPeerUpdate,
		Delete: resourceObjectWanoptPeerDelete,

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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_host_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWanoptPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWanoptPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWanoptPeer resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWanoptPeer(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWanoptPeer resource: %v", err)
	}

	d.SetId(getStringKey(d, "peer-host-id"))

	return resourceObjectWanoptPeerRead(d, m)
}

func resourceObjectWanoptPeerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWanoptPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWanoptPeer resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWanoptPeer(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWanoptPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "peer-host-id"))

	return resourceObjectWanoptPeerRead(d, m)
}

func resourceObjectWanoptPeerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWanoptPeer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWanoptPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWanoptPeerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWanoptPeer(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWanoptPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWanoptPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWanoptPeer resource from API: %v", err)
	}
	return nil
}

func flattenObjectWanoptPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptPeerPeerHostId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWanoptPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", flattenObjectWanoptPeerIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ObjectWanoptPeer-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("peer_host_id", flattenObjectWanoptPeerPeerHostId(o["peer-host-id"], d, "peer_host_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-host-id"], "ObjectWanoptPeer-PeerHostId"); ok {
			if err = d.Set("peer_host_id", vv); err != nil {
				return fmt.Errorf("Error reading peer_host_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_host_id: %v", err)
		}
	}

	return nil
}

func flattenObjectWanoptPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWanoptPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptPeerPeerHostId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWanoptPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandObjectWanoptPeerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("peer_host_id"); ok {
		t, err := expandObjectWanoptPeerPeerHostId(d, v, "peer_host_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-host-id"] = t
		}
	}

	return &obj, nil
}
