// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Pre-authorized and blocked security fabric nodes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsfTrustedList() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfTrustedListCreate,
		Read:   resourceSystemCsfTrustedListRead,
		Update: resourceSystemCsfTrustedListUpdate,
		Delete: resourceSystemCsfTrustedListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"downstream_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_members": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemCsfTrustedListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemCsfTrustedList(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCsfTrustedList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateSystemCsfTrustedList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemCsfTrustedList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCsfTrustedListRead(d, m)
}

func resourceSystemCsfTrustedListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemCsfTrustedList(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfTrustedList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemCsfTrustedList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsfTrustedList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemCsfTrustedListRead(d, m)
}

func resourceSystemCsfTrustedListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemCsfTrustedList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsfTrustedList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfTrustedListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemCsfTrustedList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfTrustedList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsfTrustedList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsfTrustedList resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfTrustedListAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListAuthorizationType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListCertificate2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListDownstreamAuthorization2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListHaMembers2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListIndex2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListSerial2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCsfTrustedList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemCsfTrustedListAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemCsfTrustedList-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("authorization_type", flattenSystemCsfTrustedListAuthorizationType2edl(o["authorization-type"], d, "authorization_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorization-type"], "SystemCsfTrustedList-AuthorizationType"); ok {
			if err = d.Set("authorization_type", vv); err != nil {
				return fmt.Errorf("Error reading authorization_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorization_type: %v", err)
		}
	}

	if err = d.Set("certificate", flattenSystemCsfTrustedListCertificate2edl(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "SystemCsfTrustedList-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("downstream_authorization", flattenSystemCsfTrustedListDownstreamAuthorization2edl(o["downstream-authorization"], d, "downstream_authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["downstream-authorization"], "SystemCsfTrustedList-DownstreamAuthorization"); ok {
			if err = d.Set("downstream_authorization", vv); err != nil {
				return fmt.Errorf("Error reading downstream_authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downstream_authorization: %v", err)
		}
	}

	if err = d.Set("ha_members", flattenSystemCsfTrustedListHaMembers2edl(o["ha-members"], d, "ha_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-members"], "SystemCsfTrustedList-HaMembers"); ok {
			if err = d.Set("ha_members", vv); err != nil {
				return fmt.Errorf("Error reading ha_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_members: %v", err)
		}
	}

	if err = d.Set("index", flattenSystemCsfTrustedListIndex2edl(o["index"], d, "index")); err != nil {
		if vv, ok := fortiAPIPatch(o["index"], "SystemCsfTrustedList-Index"); ok {
			if err = d.Set("index", vv); err != nil {
				return fmt.Errorf("Error reading index: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemCsfTrustedListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemCsfTrustedList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemCsfTrustedListSerial2edl(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemCsfTrustedList-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	return nil
}

func flattenSystemCsfTrustedListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCsfTrustedListAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListAuthorizationType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListCertificate2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListDownstreamAuthorization2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListHaMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListIndex2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListSerial2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCsfTrustedList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemCsfTrustedListAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("authorization_type"); ok || d.HasChange("authorization_type") {
		t, err := expandSystemCsfTrustedListAuthorizationType2edl(d, v, "authorization_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization-type"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandSystemCsfTrustedListCertificate2edl(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("downstream_authorization"); ok || d.HasChange("downstream_authorization") {
		t, err := expandSystemCsfTrustedListDownstreamAuthorization2edl(d, v, "downstream_authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downstream-authorization"] = t
		}
	}

	if v, ok := d.GetOk("ha_members"); ok || d.HasChange("ha_members") {
		t, err := expandSystemCsfTrustedListHaMembers2edl(d, v, "ha_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-members"] = t
		}
	}

	if v, ok := d.GetOk("index"); ok || d.HasChange("index") {
		t, err := expandSystemCsfTrustedListIndex2edl(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemCsfTrustedListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemCsfTrustedListSerial2edl(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	return &obj, nil
}
