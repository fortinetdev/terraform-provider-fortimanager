// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPsec manual keys.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVpnIpsecManualkey() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVpnIpsecManualkeyCreate,
		Read:   resourceObjectVpnIpsecManualkeyRead,
		Update: resourceObjectVpnIpsecManualkeyUpdate,
		Delete: resourceObjectVpnIpsecManualkeyDelete,

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
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authkey": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"enckey": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localspi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"npu_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remotespi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVpnIpsecManualkeyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnIpsecManualkey(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectVpnIpsecManualkey resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectVpnIpsecManualkey(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectVpnIpsecManualkey(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectVpnIpsecManualkey resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectVpnIpsecManualkey(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectVpnIpsecManualkey resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnIpsecManualkeyRead(d, m)
}

func resourceObjectVpnIpsecManualkeyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVpnIpsecManualkey(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnIpsecManualkey resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVpnIpsecManualkey(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVpnIpsecManualkey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectVpnIpsecManualkeyRead(d, m)
}

func resourceObjectVpnIpsecManualkeyDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectVpnIpsecManualkey(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVpnIpsecManualkey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVpnIpsecManualkeyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectVpnIpsecManualkey(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectVpnIpsecManualkey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVpnIpsecManualkey(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVpnIpsecManualkey resource from API: %v", err)
	}
	return nil
}

func flattenObjectVpnIpsecManualkeyAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecManualkeyEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecManualkeyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectVpnIpsecManualkeyLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecManualkeyLocalspi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecManualkeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecManualkeyNpuOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecManualkeyRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVpnIpsecManualkeyRemotespi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVpnIpsecManualkey(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("authentication", flattenObjectVpnIpsecManualkeyAuthentication(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "ObjectVpnIpsecManualkey-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("encryption", flattenObjectVpnIpsecManualkeyEncryption(o["encryption"], d, "encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["encryption"], "ObjectVpnIpsecManualkey-Encryption"); ok {
			if err = d.Set("encryption", vv); err != nil {
				return fmt.Errorf("Error reading encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("interface", flattenObjectVpnIpsecManualkeyInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "ObjectVpnIpsecManualkey-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenObjectVpnIpsecManualkeyLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-gw"], "ObjectVpnIpsecManualkey-LocalGw"); ok {
			if err = d.Set("local_gw", vv); err != nil {
				return fmt.Errorf("Error reading local_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("localspi", flattenObjectVpnIpsecManualkeyLocalspi(o["localspi"], d, "localspi")); err != nil {
		if vv, ok := fortiAPIPatch(o["localspi"], "ObjectVpnIpsecManualkey-Localspi"); ok {
			if err = d.Set("localspi", vv); err != nil {
				return fmt.Errorf("Error reading localspi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading localspi: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectVpnIpsecManualkeyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectVpnIpsecManualkey-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("npu_offload", flattenObjectVpnIpsecManualkeyNpuOffload(o["npu-offload"], d, "npu_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["npu-offload"], "ObjectVpnIpsecManualkey-NpuOffload"); ok {
			if err = d.Set("npu_offload", vv); err != nil {
				return fmt.Errorf("Error reading npu_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading npu_offload: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenObjectVpnIpsecManualkeyRemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw"], "ObjectVpnIpsecManualkey-RemoteGw"); ok {
			if err = d.Set("remote_gw", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("remotespi", flattenObjectVpnIpsecManualkeyRemotespi(o["remotespi"], d, "remotespi")); err != nil {
		if vv, ok := fortiAPIPatch(o["remotespi"], "ObjectVpnIpsecManualkey-Remotespi"); ok {
			if err = d.Set("remotespi", vv); err != nil {
				return fmt.Errorf("Error reading remotespi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remotespi: %v", err)
		}
	}

	return nil
}

func flattenObjectVpnIpsecManualkeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVpnIpsecManualkeyAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecManualkeyAuthkey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnIpsecManualkeyEnckey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnIpsecManualkeyEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecManualkeyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectVpnIpsecManualkeyLocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecManualkeyLocalspi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecManualkeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecManualkeyNpuOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecManualkeyRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVpnIpsecManualkeyRemotespi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVpnIpsecManualkey(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandObjectVpnIpsecManualkeyAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("authkey"); ok || d.HasChange("authkey") {
		t, err := expandObjectVpnIpsecManualkeyAuthkey(d, v, "authkey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authkey"] = t
		}
	}

	if v, ok := d.GetOk("enckey"); ok || d.HasChange("enckey") {
		t, err := expandObjectVpnIpsecManualkeyEnckey(d, v, "enckey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enckey"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok || d.HasChange("encryption") {
		t, err := expandObjectVpnIpsecManualkeyEncryption(d, v, "encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandObjectVpnIpsecManualkeyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok || d.HasChange("local_gw") {
		t, err := expandObjectVpnIpsecManualkeyLocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("localspi"); ok || d.HasChange("localspi") {
		t, err := expandObjectVpnIpsecManualkeyLocalspi(d, v, "localspi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localspi"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectVpnIpsecManualkeyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("npu_offload"); ok || d.HasChange("npu_offload") {
		t, err := expandObjectVpnIpsecManualkeyNpuOffload(d, v, "npu_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offload"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok || d.HasChange("remote_gw") {
		t, err := expandObjectVpnIpsecManualkeyRemoteGw(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("remotespi"); ok || d.HasChange("remotespi") {
		t, err := expandObjectVpnIpsecManualkeyRemotespi(d, v, "remotespi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remotespi"] = t
		}
	}

	return &obj, nil
}
