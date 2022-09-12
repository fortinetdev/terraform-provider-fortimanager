// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for locallog to fortianalyzer.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLocallogFortianalyzerSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogFortianalyzerSettingUpdate,
		Read:   resourceSystemLocallogFortianalyzerSettingRead,
		Update: resourceSystemLocallogFortianalyzerSettingUpdate,
		Delete: resourceSystemLocallogFortianalyzerSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"peer_cert_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reliable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocallogFortianalyzerSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogFortianalyzerSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogFortianalyzerSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogFortianalyzerSetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogFortianalyzerSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogFortianalyzerSetting")

	return resourceSystemLocallogFortianalyzerSettingRead(d, m)
}

func resourceSystemLocallogFortianalyzerSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogFortianalyzerSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogFortianalyzerSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogFortianalyzerSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogFortianalyzerSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogFortianalyzerSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogFortianalyzerSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogFortianalyzerSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogFortianalyzerSettingPeerCertCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerSettingReliable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerSettingSecureConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerSettingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerSettingSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemLocallogFortianalyzerSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("peer_cert_cn", flattenSystemLocallogFortianalyzerSettingPeerCertCn(o["peer-cert-cn"], d, "peer_cert_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-cert-cn"], "SystemLocallogFortianalyzerSetting-PeerCertCn"); ok {
			if err = d.Set("peer_cert_cn", vv); err != nil {
				return fmt.Errorf("Error reading peer_cert_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_cert_cn: %v", err)
		}
	}

	if err = d.Set("reliable", flattenSystemLocallogFortianalyzerSettingReliable(o["reliable"], d, "reliable")); err != nil {
		if vv, ok := fortiAPIPatch(o["reliable"], "SystemLocallogFortianalyzerSetting-Reliable"); ok {
			if err = d.Set("reliable", vv); err != nil {
				return fmt.Errorf("Error reading reliable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	if err = d.Set("secure_connection", flattenSystemLocallogFortianalyzerSettingSecureConnection(o["secure-connection"], d, "secure_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-connection"], "SystemLocallogFortianalyzerSetting-SecureConnection"); ok {
			if err = d.Set("secure_connection", vv); err != nil {
				return fmt.Errorf("Error reading secure_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_connection: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemLocallogFortianalyzerSettingServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemLocallogFortianalyzerSetting-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("severity", flattenSystemLocallogFortianalyzerSettingSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "SystemLocallogFortianalyzerSetting-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLocallogFortianalyzerSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLocallogFortianalyzerSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenSystemLocallogFortianalyzerSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-time"], "SystemLocallogFortianalyzerSetting-UploadTime"); ok {
			if err = d.Set("upload_time", vv); err != nil {
				return fmt.Errorf("Error reading upload_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogFortianalyzerSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogFortianalyzerSettingPeerCertCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerSettingReliable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerSettingSecureConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerSettingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerSettingSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemLocallogFortianalyzerSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("peer_cert_cn"); ok || d.HasChange("peer_cert_cn") {
		t, err := expandSystemLocallogFortianalyzerSettingPeerCertCn(d, v, "peer_cert_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-cert-cn"] = t
		}
	}

	if v, ok := d.GetOk("reliable"); ok || d.HasChange("reliable") {
		t, err := expandSystemLocallogFortianalyzerSettingReliable(d, v, "reliable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reliable"] = t
		}
	}

	if v, ok := d.GetOk("secure_connection"); ok || d.HasChange("secure_connection") {
		t, err := expandSystemLocallogFortianalyzerSettingSecureConnection(d, v, "secure_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-connection"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemLocallogFortianalyzerSettingServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandSystemLocallogFortianalyzerSettingSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLocallogFortianalyzerSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok || d.HasChange("upload_time") {
		t, err := expandSystemLocallogFortianalyzerSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	return &obj, nil
}
