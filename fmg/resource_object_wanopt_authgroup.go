// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WAN optimization authentication groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWanoptAuthGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWanoptAuthGroupCreate,
		Read:   resourceObjectWanoptAuthGroupRead,
		Update: resourceObjectWanoptAuthGroupUpdate,
		Delete: resourceObjectWanoptAuthGroupDelete,

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
			"auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert": &schema.Schema{
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
			"peer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_accept": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psk": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWanoptAuthGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWanoptAuthGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWanoptAuthGroup resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWanoptAuthGroup(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWanoptAuthGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWanoptAuthGroupRead(d, m)
}

func resourceObjectWanoptAuthGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWanoptAuthGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWanoptAuthGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWanoptAuthGroup(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWanoptAuthGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWanoptAuthGroupRead(d, m)
}

func resourceObjectWanoptAuthGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWanoptAuthGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWanoptAuthGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWanoptAuthGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWanoptAuthGroup(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWanoptAuthGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWanoptAuthGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWanoptAuthGroup resource from API: %v", err)
	}
	return nil
}

func flattenObjectWanoptAuthGroupAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptAuthGroupCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptAuthGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptAuthGroupPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptAuthGroupPeerAccept(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWanoptAuthGroupPsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectWanoptAuthGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auth_method", flattenObjectWanoptAuthGroupAuthMethod(o["auth-method"], d, "auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-method"], "ObjectWanoptAuthGroup-AuthMethod"); ok {
			if err = d.Set("auth_method", vv); err != nil {
				return fmt.Errorf("Error reading auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_method: %v", err)
		}
	}

	if err = d.Set("cert", flattenObjectWanoptAuthGroupCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "ObjectWanoptAuthGroup-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWanoptAuthGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWanoptAuthGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("peer", flattenObjectWanoptAuthGroupPeer(o["peer"], d, "peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer"], "ObjectWanoptAuthGroup-Peer"); ok {
			if err = d.Set("peer", vv); err != nil {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("peer_accept", flattenObjectWanoptAuthGroupPeerAccept(o["peer-accept"], d, "peer_accept")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-accept"], "ObjectWanoptAuthGroup-PeerAccept"); ok {
			if err = d.Set("peer_accept", vv); err != nil {
				return fmt.Errorf("Error reading peer_accept: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_accept: %v", err)
		}
	}

	if err = d.Set("psk", flattenObjectWanoptAuthGroupPsk(o["psk"], d, "psk")); err != nil {
		if vv, ok := fortiAPIPatch(o["psk"], "ObjectWanoptAuthGroup-Psk"); ok {
			if err = d.Set("psk", vv); err != nil {
				return fmt.Errorf("Error reading psk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading psk: %v", err)
		}
	}

	return nil
}

func flattenObjectWanoptAuthGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWanoptAuthGroupAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptAuthGroupCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptAuthGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptAuthGroupPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptAuthGroupPeerAccept(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWanoptAuthGroupPsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectWanoptAuthGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_method"); ok || d.HasChange("auth_method") {
		t, err := expandObjectWanoptAuthGroupAuthMethod(d, v, "auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-method"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandObjectWanoptAuthGroupCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWanoptAuthGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandObjectWanoptAuthGroupPeer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("peer_accept"); ok || d.HasChange("peer_accept") {
		t, err := expandObjectWanoptAuthGroupPeerAccept(d, v, "peer_accept")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-accept"] = t
		}
	}

	if v, ok := d.GetOk("psk"); ok || d.HasChange("psk") {
		t, err := expandObjectWanoptAuthGroupPsk(d, v, "psk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psk"] = t
		}
	}

	return &obj, nil
}
