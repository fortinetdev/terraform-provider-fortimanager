// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectDynamic AddressDynamicAddrMapping

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectDynamicAddressDynamicAddrMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectDynamicAddressDynamicAddrMappingCreate,
		Read:   resourceObjectDynamicAddressDynamicAddrMappingRead,
		Update: resourceObjectDynamicAddressDynamicAddrMappingUpdate,
		Delete: resourceObjectDynamicAddressDynamicAddrMappingDelete,

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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectDynamicAddressDynamicAddrMappingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	address := d.Get("address").(string)
	paradict["address"] = address

	obj, err := getObjectObjectDynamicAddressDynamicAddrMapping(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicAddressDynamicAddrMapping resource while getting object: %v", err)
	}

	_, err = c.CreateObjectDynamicAddressDynamicAddrMapping(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectDynamicAddressDynamicAddrMapping resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDynamicAddressDynamicAddrMappingRead(d, m)
}

func resourceObjectDynamicAddressDynamicAddrMappingUpdate(d *schema.ResourceData, m interface{}) error {
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

	address := d.Get("address").(string)
	paradict["address"] = address

	obj, err := getObjectObjectDynamicAddressDynamicAddrMapping(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicAddressDynamicAddrMapping resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectDynamicAddressDynamicAddrMapping(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectDynamicAddressDynamicAddrMapping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectDynamicAddressDynamicAddrMappingRead(d, m)
}

func resourceObjectDynamicAddressDynamicAddrMappingDelete(d *schema.ResourceData, m interface{}) error {
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

	address := d.Get("address").(string)
	paradict["address"] = address

	err = c.DeleteObjectDynamicAddressDynamicAddrMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectDynamicAddressDynamicAddrMapping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectDynamicAddressDynamicAddrMappingRead(d *schema.ResourceData, m interface{}) error {
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

	address := d.Get("address").(string)
	if address == "" {
		address = importOptionChecking(m.(*FortiClient).Cfg, "address")
		if address == "" {
			return fmt.Errorf("Parameter address is missing")
		}
		if err = d.Set("address", address); err != nil {
			return fmt.Errorf("Error set params address: %v", err)
		}
	}
	paradict["address"] = address

	o, err := c.ReadObjectDynamicAddressDynamicAddrMapping(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicAddressDynamicAddrMapping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectDynamicAddressDynamicAddrMapping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDynamicAddressDynamicAddrMapping resource from API: %v", err)
	}
	return nil
}

func flattenObjectDynamicAddressDynamicAddrMappingAddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectDynamicAddressDynamicAddrMappingId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectDynamicAddressDynamicAddrMapping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("addr", flattenObjectDynamicAddressDynamicAddrMappingAddr2edl(o["addr"], d, "addr")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr"], "ObjectDynamicAddressDynamicAddrMapping-Addr"); ok {
			if err = d.Set("addr", vv); err != nil {
				return fmt.Errorf("Error reading addr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectDynamicAddressDynamicAddrMappingId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectDynamicAddressDynamicAddrMapping-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectDynamicAddressDynamicAddrMappingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectDynamicAddressDynamicAddrMappingAddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectDynamicAddressDynamicAddrMappingId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectDynamicAddressDynamicAddrMapping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr"); ok || d.HasChange("addr") {
		t, err := expandObjectDynamicAddressDynamicAddrMappingAddr2edl(d, v, "addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectDynamicAddressDynamicAddrMappingId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
