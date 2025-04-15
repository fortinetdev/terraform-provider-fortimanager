// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectVpnmgr NodeProtectedSubnet

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnmgrNodeProtectedSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnmgrNodeProtectedSubnetCreate,
		Read:   resourceObjectVpnmgrNodeProtectedSubnetRead,
		Update: resourceObjectVpnmgrNodeProtectedSubnetUpdate,
		Delete: resourceObjectVpnmgrNodeProtectedSubnetDelete,

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
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seq": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectVpnmgrNodeProtectedSubnetCreate(d *schema.ResourceData, m interface{}) error {
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

	node := d.Get("node").(string)
	paradict["node"] = node

	obj, err := getObjectObjectVpnmgrNodeProtectedSubnet(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrNodeProtectedSubnet resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectVpnmgrNodeProtectedSubnet(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrNodeProtectedSubnet resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq")))

	return resourceObjectVpnmgrNodeProtectedSubnetRead(d, m)
}

func resourceObjectVpnmgrNodeProtectedSubnetUpdate(d *schema.ResourceData, m interface{}) error {
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

	node := d.Get("node").(string)
	paradict["node"] = node

	obj, err := getObjectObjectVpnmgrNodeProtectedSubnet(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrNodeProtectedSubnet resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVpnmgrNodeProtectedSubnet(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrNodeProtectedSubnet resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seq")))

	return resourceObjectVpnmgrNodeProtectedSubnetRead(d, m)
}

func resourceObjectVpnmgrNodeProtectedSubnetDelete(d *schema.ResourceData, m interface{}) error {
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

	node := d.Get("node").(string)
	paradict["node"] = node

	wsParams["adom"] = adomv

	err = c.DeleteObjectVpnmgrNodeProtectedSubnet(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnmgrNodeProtectedSubnet resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnmgrNodeProtectedSubnetRead(d *schema.ResourceData, m interface{}) error {
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

	node := d.Get("node").(string)
	if node == "" {
		node = importOptionChecking(m.(*FortiClient).Cfg, "node")
		if node == "" {
			return fmt.Errorf("Parameter node is missing")
		}
		if err = d.Set("node", node); err != nil {
			return fmt.Errorf("Error set params node: %v", err)
		}
	}
	paradict["node"] = node

	o, err := c.ReadObjectVpnmgrNodeProtectedSubnet(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrNodeProtectedSubnet resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnmgrNodeProtectedSubnet(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrNodeProtectedSubnet resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnmgrNodeProtectedSubnetAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectVpnmgrNodeProtectedSubnetSeq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnmgrNodeProtectedSubnet(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("addr", flattenObjectVpnmgrNodeProtectedSubnetAddr2edl(o["addr"], d, "addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr"], "ObjectVpnmgrNodeProtectedSubnet-Addr"); ok {
			if err = d.Set("addr", vv); err != nil {
				return fmt.Errorf("Error reading addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr: %v", err)
		}
	}

	if err = d.Set("seq", flattenObjectVpnmgrNodeProtectedSubnetSeq2edl(o["seq"], d, "seq")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq"], "ObjectVpnmgrNodeProtectedSubnet-Seq"); ok {
			if err = d.Set("seq", vv); err != nil {
				return fmt.Errorf("Error reading seq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnmgrNodeProtectedSubnetFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnmgrNodeProtectedSubnetAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectVpnmgrNodeProtectedSubnetSeq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnmgrNodeProtectedSubnet(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr"); ok || d.HasChange("addr") {
		t, err := expandObjectVpnmgrNodeProtectedSubnetAddr2edl(d, v, "addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr"] = t
		}
	}

	if v, ok := d.GetOk("seq"); ok || d.HasChange("seq") {
		t, err := expandObjectVpnmgrNodeProtectedSubnetSeq2edl(d, v, "seq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq"] = t
		}
	}

	return &obj, nil
}
