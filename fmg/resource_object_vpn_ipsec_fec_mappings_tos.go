// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FEC redundancy mapping table for specific type of service (TOS).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnIpsecFecMappingsTos() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnIpsecFecMappingsTosCreate,
		Read:   resourceObjectVpnIpsecFecMappingsTosRead,
		Update: resourceObjectVpnIpsecFecMappingsTosUpdate,
		Delete: resourceObjectVpnIpsecFecMappingsTosDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"fec": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mappings": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"base": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redundant": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"seqno": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectVpnIpsecFecMappingsTosCreate(d *schema.ResourceData, m interface{}) error {
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

	fec := d.Get("fec").(string)
	mappings := d.Get("mappings").(string)
	paradict["fec"] = fec
	paradict["mappings"] = mappings

	obj, err := getObjectObjectVpnIpsecFecMappingsTos(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnIpsecFecMappingsTos resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("seqno")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectVpnIpsecFecMappingsTos(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectVpnIpsecFecMappingsTos(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectVpnIpsecFecMappingsTos resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectVpnIpsecFecMappingsTos(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectVpnIpsecFecMappingsTos resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "seqno"))

	return resourceObjectVpnIpsecFecMappingsTosRead(d, m)
}

func resourceObjectVpnIpsecFecMappingsTosUpdate(d *schema.ResourceData, m interface{}) error {
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

	fec := d.Get("fec").(string)
	mappings := d.Get("mappings").(string)
	paradict["fec"] = fec
	paradict["mappings"] = mappings

	obj, err := getObjectObjectVpnIpsecFecMappingsTos(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnIpsecFecMappingsTos resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVpnIpsecFecMappingsTos(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnIpsecFecMappingsTos resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "seqno"))

	return resourceObjectVpnIpsecFecMappingsTosRead(d, m)
}

func resourceObjectVpnIpsecFecMappingsTosDelete(d *schema.ResourceData, m interface{}) error {
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

	fec := d.Get("fec").(string)
	mappings := d.Get("mappings").(string)
	paradict["fec"] = fec
	paradict["mappings"] = mappings

	wsParams["adom"] = adomv

	err = c.DeleteObjectVpnIpsecFecMappingsTos(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnIpsecFecMappingsTos resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnIpsecFecMappingsTosRead(d *schema.ResourceData, m interface{}) error {
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

	fec := d.Get("fec").(string)
	mappings := d.Get("mappings").(string)
	if fec == "" {
		fec = importOptionChecking(m.(*FortiClient).Cfg, "fec")
		if fec == "" {
			return fmt.Errorf("Parameter fec is missing")
		}
		if err = d.Set("fec", fec); err != nil {
			return fmt.Errorf("Error set params fec: %v", err)
		}
	}
	if mappings == "" {
		mappings = importOptionChecking(m.(*FortiClient).Cfg, "mappings")
		if mappings == "" {
			return fmt.Errorf("Parameter mappings is missing")
		}
		if err = d.Set("mappings", mappings); err != nil {
			return fmt.Errorf("Error set params mappings: %v", err)
		}
	}
	paradict["fec"] = fec
	paradict["mappings"] = mappings

	o, err := c.ReadObjectVpnIpsecFecMappingsTos(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectVpnIpsecFecMappingsTos resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnIpsecFecMappingsTos(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnIpsecFecMappingsTos resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnIpsecFecMappingsTosBase3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTosRedundant3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTosSeqno3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTosTos3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecFecMappingsTosTosMask3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnIpsecFecMappingsTos(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("base", flattenObjectVpnIpsecFecMappingsTosBase3rdl(o["base"], d, "base")); err != nil {
		if vv, ok := fortiAPIPatch(o["base"], "ObjectVpnIpsecFecMappingsTos-Base"); ok {
			if err = d.Set("base", vv); err != nil {
				return fmt.Errorf("Error reading base: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base: %v", err)
		}
	}

	if err = d.Set("redundant", flattenObjectVpnIpsecFecMappingsTosRedundant3rdl(o["redundant"], d, "redundant")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant"], "ObjectVpnIpsecFecMappingsTos-Redundant"); ok {
			if err = d.Set("redundant", vv); err != nil {
				return fmt.Errorf("Error reading redundant: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant: %v", err)
		}
	}

	if err = d.Set("seqno", flattenObjectVpnIpsecFecMappingsTosSeqno3rdl(o["seqno"], d, "seqno")); err != nil {
		if vv, ok := fortiAPIPatch(o["seqno"], "ObjectVpnIpsecFecMappingsTos-Seqno"); ok {
			if err = d.Set("seqno", vv); err != nil {
				return fmt.Errorf("Error reading seqno: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seqno: %v", err)
		}
	}

	if err = d.Set("tos", flattenObjectVpnIpsecFecMappingsTosTos3rdl(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "ObjectVpnIpsecFecMappingsTos-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenObjectVpnIpsecFecMappingsTosTosMask3rdl(o["tos-mask"], d, "tos_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos-mask"], "ObjectVpnIpsecFecMappingsTos-TosMask"); ok {
			if err = d.Set("tos_mask", vv); err != nil {
				return fmt.Errorf("Error reading tos_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnIpsecFecMappingsTosFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnIpsecFecMappingsTosBase3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTosRedundant3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTosSeqno3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTosTos3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecFecMappingsTosTosMask3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnIpsecFecMappingsTos(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("base"); ok || d.HasChange("base") {
		t, err := expandObjectVpnIpsecFecMappingsTosBase3rdl(d, v, "base")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base"] = t
		}
	}

	if v, ok := d.GetOk("redundant"); ok || d.HasChange("redundant") {
		t, err := expandObjectVpnIpsecFecMappingsTosRedundant3rdl(d, v, "redundant")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant"] = t
		}
	}

	if v, ok := d.GetOk("seqno"); ok || d.HasChange("seqno") {
		t, err := expandObjectVpnIpsecFecMappingsTosSeqno3rdl(d, v, "seqno")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seqno"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandObjectVpnIpsecFecMappingsTosTos3rdl(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok || d.HasChange("tos_mask") {
		t, err := expandObjectVpnIpsecFecMappingsTosTosMask3rdl(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	return &obj, nil
}
