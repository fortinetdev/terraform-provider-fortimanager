// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: File filter entries.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWebfilterProfileFileFilterEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebfilterProfileFileFilterEntriesCreate,
		Read:   resourceObjectWebfilterProfileFileFilterEntriesRead,
		Update: resourceObjectWebfilterProfileFileFilterEntriesUpdate,
		Delete: resourceObjectWebfilterProfileFileFilterEntriesDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password_protected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWebfilterProfileFileFilterEntriesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterProfileFileFilterEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfileFileFilterEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWebfilterProfileFileFilterEntries(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebfilterProfileFileFilterEntries resource: %v", err)
	}

	d.SetId(getStringKey(d, "filter"))

	return resourceObjectWebfilterProfileFileFilterEntriesRead(d, m)
}

func resourceObjectWebfilterProfileFileFilterEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWebfilterProfileFileFilterEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileFileFilterEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWebfilterProfileFileFilterEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebfilterProfileFileFilterEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "filter"))

	return resourceObjectWebfilterProfileFileFilterEntriesRead(d, m)
}

func resourceObjectWebfilterProfileFileFilterEntriesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWebfilterProfileFileFilterEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebfilterProfileFileFilterEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebfilterProfileFileFilterEntriesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWebfilterProfileFileFilterEntries(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileFileFilterEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebfilterProfileFileFilterEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebfilterProfileFileFilterEntries resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebfilterProfileFileFilterEntriesAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesComment3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesEncryption3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesFileType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWebfilterProfileFileFilterEntriesFilter3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesPasswordProtected3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebfilterProfileFileFilterEntriesProtocol3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectWebfilterProfileFileFilterEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectWebfilterProfileFileFilterEntriesAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectWebfilterProfileFileFilterEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectWebfilterProfileFileFilterEntriesComment3rdl(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWebfilterProfileFileFilterEntries-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("direction", flattenObjectWebfilterProfileFileFilterEntriesDirection3rdl(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "ObjectWebfilterProfileFileFilterEntries-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("encryption", flattenObjectWebfilterProfileFileFilterEntriesEncryption3rdl(o["encryption"], d, "encryption")); err != nil {
		if vv, ok := fortiAPIPatch(o["encryption"], "ObjectWebfilterProfileFileFilterEntries-Encryption"); ok {
			if err = d.Set("encryption", vv); err != nil {
				return fmt.Errorf("Error reading encryption: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("file_type", flattenObjectWebfilterProfileFileFilterEntriesFileType3rdl(o["file-type"], d, "file_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-type"], "ObjectWebfilterProfileFileFilterEntries-FileType"); ok {
			if err = d.Set("file_type", vv); err != nil {
				return fmt.Errorf("Error reading file_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_type: %v", err)
		}
	}

	if err = d.Set("filter", flattenObjectWebfilterProfileFileFilterEntriesFilter3rdl(o["filter"], d, "filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["filter"], "ObjectWebfilterProfileFileFilterEntries-Filter"); ok {
			if err = d.Set("filter", vv); err != nil {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("password_protected", flattenObjectWebfilterProfileFileFilterEntriesPasswordProtected3rdl(o["password-protected"], d, "password_protected")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-protected"], "ObjectWebfilterProfileFileFilterEntries-PasswordProtected"); ok {
			if err = d.Set("password_protected", vv); err != nil {
				return fmt.Errorf("Error reading password_protected: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_protected: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectWebfilterProfileFileFilterEntriesProtocol3rdl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectWebfilterProfileFileFilterEntries-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	return nil
}

func flattenObjectWebfilterProfileFileFilterEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebfilterProfileFileFilterEntriesAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesComment3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesEncryption3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesFileType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebfilterProfileFileFilterEntriesFilter3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesPasswordProtected3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebfilterProfileFileFilterEntriesProtocol3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectWebfilterProfileFileFilterEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectWebfilterProfileFileFilterEntriesAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWebfilterProfileFileFilterEntriesComment3rdl(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandObjectWebfilterProfileFileFilterEntriesDirection3rdl(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok || d.HasChange("encryption") {
		t, err := expandObjectWebfilterProfileFileFilterEntriesEncryption3rdl(d, v, "encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("file_type"); ok || d.HasChange("file_type") {
		t, err := expandObjectWebfilterProfileFileFilterEntriesFileType3rdl(d, v, "file_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandObjectWebfilterProfileFileFilterEntriesFilter3rdl(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("password_protected"); ok || d.HasChange("password_protected") {
		t, err := expandObjectWebfilterProfileFileFilterEntriesPasswordProtected3rdl(d, v, "password_protected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-protected"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectWebfilterProfileFileFilterEntriesProtocol3rdl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	return &obj, nil
}
