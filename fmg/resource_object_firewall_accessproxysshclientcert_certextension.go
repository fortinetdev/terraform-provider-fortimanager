// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure certificate extension for user certificate.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallAccessProxySshClientCertCertExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallAccessProxySshClientCertCertExtensionCreate,
		Read:   resourceObjectFirewallAccessProxySshClientCertCertExtensionRead,
		Update: resourceObjectFirewallAccessProxySshClientCertCertExtensionUpdate,
		Delete: resourceObjectFirewallAccessProxySshClientCertCertExtensionDelete,

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
			"access_proxy_ssh_client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"critical": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallAccessProxySshClientCertCertExtensionCreate(d *schema.ResourceData, m interface{}) error {
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

	access_proxy_ssh_client_cert := d.Get("access_proxy_ssh_client_cert").(string)
	paradict["access_proxy_ssh_client_cert"] = access_proxy_ssh_client_cert

	obj, err := getObjectObjectFirewallAccessProxySshClientCertCertExtension(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAccessProxySshClientCertCertExtension resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallAccessProxySshClientCertCertExtension(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallAccessProxySshClientCertCertExtension resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAccessProxySshClientCertCertExtensionRead(d, m)
}

func resourceObjectFirewallAccessProxySshClientCertCertExtensionUpdate(d *schema.ResourceData, m interface{}) error {
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

	access_proxy_ssh_client_cert := d.Get("access_proxy_ssh_client_cert").(string)
	paradict["access_proxy_ssh_client_cert"] = access_proxy_ssh_client_cert

	obj, err := getObjectObjectFirewallAccessProxySshClientCertCertExtension(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAccessProxySshClientCertCertExtension resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallAccessProxySshClientCertCertExtension(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallAccessProxySshClientCertCertExtension resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallAccessProxySshClientCertCertExtensionRead(d, m)
}

func resourceObjectFirewallAccessProxySshClientCertCertExtensionDelete(d *schema.ResourceData, m interface{}) error {
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

	access_proxy_ssh_client_cert := d.Get("access_proxy_ssh_client_cert").(string)
	paradict["access_proxy_ssh_client_cert"] = access_proxy_ssh_client_cert

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallAccessProxySshClientCertCertExtension(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallAccessProxySshClientCertCertExtension resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallAccessProxySshClientCertCertExtensionRead(d *schema.ResourceData, m interface{}) error {
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

	access_proxy_ssh_client_cert := d.Get("access_proxy_ssh_client_cert").(string)
	if access_proxy_ssh_client_cert == "" {
		access_proxy_ssh_client_cert = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy_ssh_client_cert")
		if access_proxy_ssh_client_cert == "" {
			return fmt.Errorf("Parameter access_proxy_ssh_client_cert is missing")
		}
		if err = d.Set("access_proxy_ssh_client_cert", access_proxy_ssh_client_cert); err != nil {
			return fmt.Errorf("Error set params access_proxy_ssh_client_cert: %v", err)
		}
	}
	paradict["access_proxy_ssh_client_cert"] = access_proxy_ssh_client_cert

	o, err := c.ReadObjectFirewallAccessProxySshClientCertCertExtension(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAccessProxySshClientCertCertExtension resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallAccessProxySshClientCertCertExtension(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallAccessProxySshClientCertCertExtension resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionCritical2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionData2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallAccessProxySshClientCertCertExtension(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("critical", flattenObjectFirewallAccessProxySshClientCertCertExtensionCritical2edl(o["critical"], d, "critical")); err != nil {
		if vv, ok := fortiAPIPatch(o["critical"], "ObjectFirewallAccessProxySshClientCertCertExtension-Critical"); ok {
			if err = d.Set("critical", vv); err != nil {
				return fmt.Errorf("Error reading critical: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading critical: %v", err)
		}
	}

	if err = d.Set("data", flattenObjectFirewallAccessProxySshClientCertCertExtensionData2edl(o["data"], d, "data")); err != nil {
		if vv, ok := fortiAPIPatch(o["data"], "ObjectFirewallAccessProxySshClientCertCertExtension-Data"); ok {
			if err = d.Set("data", vv); err != nil {
				return fmt.Errorf("Error reading data: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallAccessProxySshClientCertCertExtensionName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallAccessProxySshClientCertCertExtension-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallAccessProxySshClientCertCertExtensionType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallAccessProxySshClientCertCertExtension-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallAccessProxySshClientCertCertExtensionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallAccessProxySshClientCertCertExtensionCritical2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertCertExtensionData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertCertExtensionName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallAccessProxySshClientCertCertExtensionType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallAccessProxySshClientCertCertExtension(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("critical"); ok || d.HasChange("critical") {
		t, err := expandObjectFirewallAccessProxySshClientCertCertExtensionCritical2edl(d, v, "critical")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["critical"] = t
		}
	}

	if v, ok := d.GetOk("data"); ok || d.HasChange("data") {
		t, err := expandObjectFirewallAccessProxySshClientCertCertExtensionData2edl(d, v, "data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallAccessProxySshClientCertCertExtensionName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallAccessProxySshClientCertCertExtensionType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
