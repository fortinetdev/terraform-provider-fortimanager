// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Peer.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFmgClusterPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFmgClusterPeerCreate,
		Read:   resourceSystemFmgClusterPeerRead,
		Update: resourceSystemFmgClusterPeerUpdate,
		Delete: resourceSystemFmgClusterPeerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sn": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemFmgClusterPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemFmgClusterPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemFmgClusterPeer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemFmgClusterPeer(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemFmgClusterPeer resource: %v", err)
	}

	d.SetId(getStringKey(d, "sn"))

	return resourceSystemFmgClusterPeerRead(d, m)
}

func resourceSystemFmgClusterPeerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemFmgClusterPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFmgClusterPeer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemFmgClusterPeer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemFmgClusterPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "sn"))

	return resourceSystemFmgClusterPeerRead(d, m)
}

func resourceSystemFmgClusterPeerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemFmgClusterPeer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFmgClusterPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFmgClusterPeerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemFmgClusterPeer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemFmgClusterPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFmgClusterPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFmgClusterPeer resource from API: %v", err)
	}
	return nil
}

func flattenSystemFmgClusterPeerAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterPeerFqdn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterPeerName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmgClusterPeerSn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFmgClusterPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr", flattenSystemFmgClusterPeerAddr2edl(o["addr"], d, "addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr"], "SystemFmgClusterPeer-Addr"); ok {
			if err = d.Set("addr", vv); err != nil {
				return fmt.Errorf("Error reading addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenSystemFmgClusterPeerFqdn2edl(o["fqdn"], d, "fqdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["fqdn"], "SystemFmgClusterPeer-Fqdn"); ok {
			if err = d.Set("fqdn", vv); err != nil {
				return fmt.Errorf("Error reading fqdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemFmgClusterPeerName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemFmgClusterPeer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("sn", flattenSystemFmgClusterPeerSn2edl(o["sn"], d, "sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sn"], "SystemFmgClusterPeer-Sn"); ok {
			if err = d.Set("sn", vv); err != nil {
				return fmt.Errorf("Error reading sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sn: %v", err)
		}
	}

	return nil
}

func flattenSystemFmgClusterPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFmgClusterPeerAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterPeerFqdn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterPeerName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmgClusterPeerSn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFmgClusterPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr"); ok || d.HasChange("addr") {
		t, err := expandSystemFmgClusterPeerAddr2edl(d, v, "addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok || d.HasChange("fqdn") {
		t, err := expandSystemFmgClusterPeerFqdn2edl(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemFmgClusterPeerName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("sn"); ok || d.HasChange("sn") {
		t, err := expandSystemFmgClusterPeerSn2edl(d, v, "sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sn"] = t
		}
	}

	return &obj, nil
}
