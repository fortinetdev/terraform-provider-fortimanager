// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Replacement messages.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempSystemReplacemsgWebproxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempSystemReplacemsgWebproxyUpdate,
		Read:   resourceSystempSystemReplacemsgWebproxyRead,
		Update: resourceSystempSystemReplacemsgWebproxyUpdate,
		Delete: resourceSystempSystemReplacemsgWebproxyDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"buffer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"msg_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystempSystemReplacemsgWebproxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempSystemReplacemsgWebproxy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemReplacemsgWebproxy resource while getting object: %v", err)
	}

	_, err = c.UpdateSystempSystemReplacemsgWebproxy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystempSystemReplacemsgWebproxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempSystemReplacemsgWebproxy")

	return resourceSystempSystemReplacemsgWebproxyRead(d, m)
}

func resourceSystempSystemReplacemsgWebproxyDelete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	err = c.DeleteSystempSystemReplacemsgWebproxy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystempSystemReplacemsgWebproxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempSystemReplacemsgWebproxyRead(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	paradict["devprof"] = devprof

	o, err := c.ReadSystempSystemReplacemsgWebproxy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemReplacemsgWebproxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempSystemReplacemsgWebproxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempSystemReplacemsgWebproxy resource from API: %v", err)
	}
	return nil
}

func flattenSystempSystemReplacemsgWebproxyBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemReplacemsgWebproxyFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemReplacemsgWebproxyHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempSystemReplacemsgWebproxyMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempSystemReplacemsgWebproxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("buffer", flattenSystempSystemReplacemsgWebproxyBuffer(o["buffer"], d, "buffer")); err != nil {
		if vv, ok := fortiAPIPatch(o["buffer"], "SystempSystemReplacemsgWebproxy-Buffer"); ok {
			if err = d.Set("buffer", vv); err != nil {
				return fmt.Errorf("Error reading buffer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("format", flattenSystempSystemReplacemsgWebproxyFormat(o["format"], d, "format")); err != nil {
		if vv, ok := fortiAPIPatch(o["format"], "SystempSystemReplacemsgWebproxy-Format"); ok {
			if err = d.Set("format", vv); err != nil {
				return fmt.Errorf("Error reading format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("header", flattenSystempSystemReplacemsgWebproxyHeader(o["header"], d, "header")); err != nil {
		if vv, ok := fortiAPIPatch(o["header"], "SystempSystemReplacemsgWebproxy-Header"); ok {
			if err = d.Set("header", vv); err != nil {
				return fmt.Errorf("Error reading header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("msg_type", flattenSystempSystemReplacemsgWebproxyMsgType(o["msg-type"], d, "msg_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["msg-type"], "SystempSystemReplacemsgWebproxy-MsgType"); ok {
			if err = d.Set("msg_type", vv); err != nil {
				return fmt.Errorf("Error reading msg_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	return nil
}

func flattenSystempSystemReplacemsgWebproxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempSystemReplacemsgWebproxyBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemReplacemsgWebproxyFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemReplacemsgWebproxyHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempSystemReplacemsgWebproxyMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempSystemReplacemsgWebproxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("buffer"); ok || d.HasChange("buffer") {
		t, err := expandSystempSystemReplacemsgWebproxyBuffer(d, v, "buffer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["buffer"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok || d.HasChange("format") {
		t, err := expandSystempSystemReplacemsgWebproxyFormat(d, v, "format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok || d.HasChange("header") {
		t, err := expandSystempSystemReplacemsgWebproxyHeader(d, v, "header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("msg_type"); ok || d.HasChange("msg_type") {
		t, err := expandSystempSystemReplacemsgWebproxyMsgType(d, v, "msg_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg-type"] = t
		}
	}

	return &obj, nil
}
