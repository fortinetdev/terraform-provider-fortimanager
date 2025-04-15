// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ICAP forwarded request headers.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectIcapProfileIcapHeaders() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectIcapProfileIcapHeadersCreate,
		Read:   resourceObjectIcapProfileIcapHeadersRead,
		Update: resourceObjectIcapProfileIcapHeadersUpdate,
		Delete: resourceObjectIcapProfileIcapHeadersDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"base64_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectIcapProfileIcapHeadersCreate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectIcapProfileIcapHeaders(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapProfileIcapHeaders resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	v, err := c.CreateObjectIcapProfileIcapHeaders(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectIcapProfileIcapHeaders resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceObjectIcapProfileIcapHeadersRead(d, m)
		} else {
			return fmt.Errorf("Error creating ObjectIcapProfileIcapHeaders resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIcapProfileIcapHeadersRead(d, m)
}

func resourceObjectIcapProfileIcapHeadersUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectIcapProfileIcapHeaders(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapProfileIcapHeaders resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectIcapProfileIcapHeaders(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectIcapProfileIcapHeaders resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectIcapProfileIcapHeadersRead(d, m)
}

func resourceObjectIcapProfileIcapHeadersDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectIcapProfileIcapHeaders(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectIcapProfileIcapHeaders resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectIcapProfileIcapHeadersRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectIcapProfileIcapHeaders(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapProfileIcapHeaders resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectIcapProfileIcapHeaders(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectIcapProfileIcapHeaders resource from API: %v", err)
	}
	return nil
}

func flattenObjectIcapProfileIcapHeadersBase64Encoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileIcapHeadersContent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileIcapHeadersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectIcapProfileIcapHeadersName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectIcapProfileIcapHeaders(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("base64_encoding", flattenObjectIcapProfileIcapHeadersBase64Encoding2edl(o["base64-encoding"], d, "base64_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["base64-encoding"], "ObjectIcapProfileIcapHeaders-Base64Encoding"); ok {
			if err = d.Set("base64_encoding", vv); err != nil {
				return fmt.Errorf("Error reading base64_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading base64_encoding: %v", err)
		}
	}

	if err = d.Set("content", flattenObjectIcapProfileIcapHeadersContent2edl(o["content"], d, "content")); err != nil {
		if vv, ok := fortiAPIPatch(o["content"], "ObjectIcapProfileIcapHeaders-Content"); ok {
			if err = d.Set("content", vv); err != nil {
				return fmt.Errorf("Error reading content: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectIcapProfileIcapHeadersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectIcapProfileIcapHeaders-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectIcapProfileIcapHeadersName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectIcapProfileIcapHeaders-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectIcapProfileIcapHeadersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectIcapProfileIcapHeadersBase64Encoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileIcapHeadersContent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileIcapHeadersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectIcapProfileIcapHeadersName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectIcapProfileIcapHeaders(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("base64_encoding"); ok || d.HasChange("base64_encoding") {
		t, err := expandObjectIcapProfileIcapHeadersBase64Encoding2edl(d, v, "base64_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base64-encoding"] = t
		}
	}

	if v, ok := d.GetOk("content"); ok || d.HasChange("content") {
		t, err := expandObjectIcapProfileIcapHeadersContent2edl(d, v, "content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectIcapProfileIcapHeadersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectIcapProfileIcapHeadersName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
