// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Alert emails.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemMail() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMailCreate,
		Read:   resourceSystemMailRead,
		Update: resourceSystemMailUpdate,
		Delete: resourceSystemMailDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_passwd": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"secure_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemMailCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemMail(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemMail resource while getting object: %v", err)
	}

	_, err = c.CreateSystemMail(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemMail resource: %v", err)
	}

	d.SetId(getStringKey(d, "fosid"))

	return resourceSystemMailRead(d, m)
}

func resourceSystemMailUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemMail(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemMail resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMail(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemMail resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "fosid"))

	return resourceSystemMailRead(d, m)
}

func resourceSystemMailDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemMail(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMail resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMailRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemMail(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemMail resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMail(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemMail resource from API: %v", err)
	}
	return nil
}

func flattenSystemMailAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMailAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMailFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMailFromPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemMailId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMailLocalCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMailPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemMailPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMailSecureOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMailServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMailUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemMail(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth", flattenSystemMailAuth(o["auth"], d, "auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth"], "SystemMail-Auth"); ok {
			if err = d.Set("auth", vv); err != nil {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenSystemMailAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "SystemMail-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("from", flattenSystemMailFrom(o["from"], d, "from")); err != nil {
		if vv, ok := fortiAPIPatch(o["from"], "SystemMail-From"); ok {
			if err = d.Set("from", vv); err != nil {
				return fmt.Errorf("Error reading from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading from: %v", err)
		}
	}

	if err = d.Set("from_passwd", flattenSystemMailFromPasswd(o["from_passwd"], d, "from_passwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["from_passwd"], "SystemMail-FromPasswd"); ok {
			if err = d.Set("from_passwd", vv); err != nil {
				return fmt.Errorf("Error reading from_passwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading from_passwd: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemMailId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemMail-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenSystemMailLocalCert(o["local-cert"], d, "local_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-cert"], "SystemMail-LocalCert"); ok {
			if err = d.Set("local_cert", vv); err != nil {
				return fmt.Errorf("Error reading local_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemMailPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemMail-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secure_option", flattenSystemMailSecureOption(o["secure-option"], d, "secure_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-option"], "SystemMail-SecureOption"); ok {
			if err = d.Set("secure_option", vv); err != nil {
				return fmt.Errorf("Error reading secure_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_option: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemMailServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemMail-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemMailUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "SystemMail-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenSystemMailFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemMailAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMailAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMailFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMailFromPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemMailId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMailLocalCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMailPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemMailPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMailSecureOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMailServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMailUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMail(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandSystemMailAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandSystemMailAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("from"); ok || d.HasChange("from") {
		t, err := expandSystemMailFrom(d, v, "from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["from"] = t
		}
	}

	if v, ok := d.GetOk("from_passwd"); ok || d.HasChange("from_passwd") {
		t, err := expandSystemMailFromPasswd(d, v, "from_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["from_passwd"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemMailId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("local_cert"); ok || d.HasChange("local_cert") {
		t, err := expandSystemMailLocalCert(d, v, "local_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok || d.HasChange("passwd") {
		t, err := expandSystemMailPasswd(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemMailPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secure_option"); ok || d.HasChange("secure_option") {
		t, err := expandSystemMailSecureOption(d, v, "secure_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-option"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemMailServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandSystemMailUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
