// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectVpnmgr NodeSummaryAddr

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnmgrNodeSummaryAddr() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnmgrNodeSummaryAddrCreate,
		Read:   resourceObjectVpnmgrNodeSummaryAddrRead,
		Update: resourceObjectVpnmgrNodeSummaryAddrUpdate,
		Delete: resourceObjectVpnmgrNodeSummaryAddrDelete,

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
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceObjectVpnmgrNodeSummaryAddrCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnmgrNodeSummaryAddr(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrNodeSummaryAddr resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectVpnmgrNodeSummaryAddr(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnmgrNodeSummaryAddr resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "seq")))

	return resourceObjectVpnmgrNodeSummaryAddrRead(d, m)
}

func resourceObjectVpnmgrNodeSummaryAddrUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnmgrNodeSummaryAddr(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrNodeSummaryAddr resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVpnmgrNodeSummaryAddr(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnmgrNodeSummaryAddr resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "seq")))

	return resourceObjectVpnmgrNodeSummaryAddrRead(d, m)
}

func resourceObjectVpnmgrNodeSummaryAddrDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectVpnmgrNodeSummaryAddr(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnmgrNodeSummaryAddr resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnmgrNodeSummaryAddrRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectVpnmgrNodeSummaryAddr(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrNodeSummaryAddr resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnmgrNodeSummaryAddr(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnmgrNodeSummaryAddr resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnmgrNodeSummaryAddrAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectVpnmgrNodeSummaryAddrPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnmgrNodeSummaryAddrSeq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnmgrNodeSummaryAddr(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("addr", flattenObjectVpnmgrNodeSummaryAddrAddr2edl(o["addr"], d, "addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr"], "ObjectVpnmgrNodeSummaryAddr-Addr"); ok {
			if err = d.Set("addr", vv); err != nil {
				return fmt.Errorf("Error reading addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr: %v", err)
		}
	}

	if err = d.Set("priority", flattenObjectVpnmgrNodeSummaryAddrPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectVpnmgrNodeSummaryAddr-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("seq", flattenObjectVpnmgrNodeSummaryAddrSeq2edl(o["seq"], d, "seq")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq"], "ObjectVpnmgrNodeSummaryAddr-Seq"); ok {
			if err = d.Set("seq", vv); err != nil {
				return fmt.Errorf("Error reading seq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnmgrNodeSummaryAddrFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnmgrNodeSummaryAddrAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectVpnmgrNodeSummaryAddrPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnmgrNodeSummaryAddrSeq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnmgrNodeSummaryAddr(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr"); ok || d.HasChange("addr") {
		t, err := expandObjectVpnmgrNodeSummaryAddrAddr2edl(d, v, "addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandObjectVpnmgrNodeSummaryAddrPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("seq"); ok || d.HasChange("seq") {
		t, err := expandObjectVpnmgrNodeSummaryAddrSeq2edl(d, v, "seq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq"] = t
		}
	}

	return &obj, nil
}
