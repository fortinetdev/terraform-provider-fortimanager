// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: NTP server.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNtpNtpserver() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNtpNtpserverCreate,
		Read:   resourceSystemNtpNtpserverRead,
		Update: resourceSystemNtpNtpserverUpdate,
		Delete: resourceSystemNtpNtpserverDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"key_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maxpoll": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minpoll": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ntpv3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemNtpNtpserverCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemNtpNtpserver(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemNtpNtpserver resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	v, err := c.CreateSystemNtpNtpserver(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemNtpNtpserver resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceSystemNtpNtpserverRead(d, m)
		} else {
			return fmt.Errorf("Error creating SystemNtpNtpserver resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemNtpNtpserverRead(d, m)
}

func resourceSystemNtpNtpserverUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemNtpNtpserver(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtpNtpserver resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemNtpNtpserver(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtpNtpserver resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemNtpNtpserverRead(d, m)
}

func resourceSystemNtpNtpserverDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemNtpNtpserver(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNtpNtpserver resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNtpNtpserverRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemNtpNtpserver(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtpNtpserver resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNtpNtpserver(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtpNtpserver resource from API: %v", err)
	}
	return nil
}

func flattenSystemNtpNtpserverAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverKeyId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverMaxpoll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverMinpoll2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverNtpv32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNtpNtpserverServer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNtpNtpserver(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("authentication", flattenSystemNtpNtpserverAuthentication2edl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "SystemNtpNtpserver-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemNtpNtpserverId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemNtpNtpserver-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("key_id", flattenSystemNtpNtpserverKeyId2edl(o["key-id"], d, "key_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-id"], "SystemNtpNtpserver-KeyId"); ok {
			if err = d.Set("key_id", vv); err != nil {
				return fmt.Errorf("Error reading key_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_id: %v", err)
		}
	}

	if err = d.Set("maxpoll", flattenSystemNtpNtpserverMaxpoll2edl(o["maxpoll"], d, "maxpoll")); err != nil {
		if vv, ok := fortiAPIPatch(o["maxpoll"], "SystemNtpNtpserver-Maxpoll"); ok {
			if err = d.Set("maxpoll", vv); err != nil {
				return fmt.Errorf("Error reading maxpoll: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maxpoll: %v", err)
		}
	}

	if err = d.Set("minpoll", flattenSystemNtpNtpserverMinpoll2edl(o["minpoll"], d, "minpoll")); err != nil {
		if vv, ok := fortiAPIPatch(o["minpoll"], "SystemNtpNtpserver-Minpoll"); ok {
			if err = d.Set("minpoll", vv); err != nil {
				return fmt.Errorf("Error reading minpoll: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minpoll: %v", err)
		}
	}

	if err = d.Set("ntpv3", flattenSystemNtpNtpserverNtpv32edl(o["ntpv3"], d, "ntpv3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ntpv3"], "SystemNtpNtpserver-Ntpv3"); ok {
			if err = d.Set("ntpv3", vv); err != nil {
				return fmt.Errorf("Error reading ntpv3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ntpv3: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemNtpNtpserverServer2edl(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemNtpNtpserver-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	return nil
}

func flattenSystemNtpNtpserverFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNtpNtpserverAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNtpNtpserverKeyId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverMaxpoll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverMinpoll2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverNtpv32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverServer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNtpNtpserver(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandSystemNtpNtpserverAuthentication2edl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemNtpNtpserverId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandSystemNtpNtpserverKey2edl(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("key_id"); ok || d.HasChange("key_id") {
		t, err := expandSystemNtpNtpserverKeyId2edl(d, v, "key_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-id"] = t
		}
	}

	if v, ok := d.GetOk("maxpoll"); ok || d.HasChange("maxpoll") {
		t, err := expandSystemNtpNtpserverMaxpoll2edl(d, v, "maxpoll")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maxpoll"] = t
		}
	}

	if v, ok := d.GetOk("minpoll"); ok || d.HasChange("minpoll") {
		t, err := expandSystemNtpNtpserverMinpoll2edl(d, v, "minpoll")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minpoll"] = t
		}
	}

	if v, ok := d.GetOk("ntpv3"); ok || d.HasChange("ntpv3") {
		t, err := expandSystemNtpNtpserverNtpv32edl(d, v, "ntpv3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntpv3"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemNtpNtpserverServer2edl(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	return &obj, nil
}
