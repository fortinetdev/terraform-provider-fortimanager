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

func resourceSystemLocallogFortianalyzer2Setting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogFortianalyzer2SettingUpdate,
		Read:   resourceSystemLocallogFortianalyzer2SettingRead,
		Update: resourceSystemLocallogFortianalyzer2SettingUpdate,
		Delete: resourceSystemLocallogFortianalyzer2SettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"peer_cert_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
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

func resourceSystemLocallogFortianalyzer2SettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogFortianalyzer2Setting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogFortianalyzer2Setting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogFortianalyzer2Setting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogFortianalyzer2Setting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogFortianalyzer2Setting")

	return resourceSystemLocallogFortianalyzer2SettingRead(d, m)
}

func resourceSystemLocallogFortianalyzer2SettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogFortianalyzer2Setting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogFortianalyzer2Setting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogFortianalyzer2SettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogFortianalyzer2Setting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogFortianalyzer2Setting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogFortianalyzer2Setting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogFortianalyzer2Setting resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogFortianalyzer2SettingPeerCertCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer2SettingReliable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer2SettingSecureConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer2SettingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer2SettingSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer2SettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer2SettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemLocallogFortianalyzer2Setting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("peer_cert_cn", flattenSystemLocallogFortianalyzer2SettingPeerCertCn(o["peer-cert-cn"], d, "peer_cert_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-cert-cn"], "SystemLocallogFortianalyzer2Setting-PeerCertCn"); ok {
			if err = d.Set("peer_cert_cn", vv); err != nil {
				return fmt.Errorf("Error reading peer_cert_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_cert_cn: %v", err)
		}
	}

	if err = d.Set("reliable", flattenSystemLocallogFortianalyzer2SettingReliable(o["reliable"], d, "reliable")); err != nil {
		if vv, ok := fortiAPIPatch(o["reliable"], "SystemLocallogFortianalyzer2Setting-Reliable"); ok {
			if err = d.Set("reliable", vv); err != nil {
				return fmt.Errorf("Error reading reliable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	if err = d.Set("secure_connection", flattenSystemLocallogFortianalyzer2SettingSecureConnection(o["secure-connection"], d, "secure_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-connection"], "SystemLocallogFortianalyzer2Setting-SecureConnection"); ok {
			if err = d.Set("secure_connection", vv); err != nil {
				return fmt.Errorf("Error reading secure_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_connection: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemLocallogFortianalyzer2SettingServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemLocallogFortianalyzer2Setting-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("severity", flattenSystemLocallogFortianalyzer2SettingSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "SystemLocallogFortianalyzer2Setting-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLocallogFortianalyzer2SettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLocallogFortianalyzer2Setting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenSystemLocallogFortianalyzer2SettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-time"], "SystemLocallogFortianalyzer2Setting-UploadTime"); ok {
			if err = d.Set("upload_time", vv); err != nil {
				return fmt.Errorf("Error reading upload_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogFortianalyzer2SettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogFortianalyzer2SettingPeerCertCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer2SettingReliable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer2SettingSecureConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer2SettingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer2SettingSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer2SettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer2SettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemLocallogFortianalyzer2Setting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("peer_cert_cn"); ok || d.HasChange("peer_cert_cn") {
		t, err := expandSystemLocallogFortianalyzer2SettingPeerCertCn(d, v, "peer_cert_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-cert-cn"] = t
		}
	}

	if v, ok := d.GetOk("reliable"); ok || d.HasChange("reliable") {
		t, err := expandSystemLocallogFortianalyzer2SettingReliable(d, v, "reliable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reliable"] = t
		}
	}

	if v, ok := d.GetOk("secure_connection"); ok || d.HasChange("secure_connection") {
		t, err := expandSystemLocallogFortianalyzer2SettingSecureConnection(d, v, "secure_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-connection"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemLocallogFortianalyzer2SettingServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandSystemLocallogFortianalyzer2SettingSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLocallogFortianalyzer2SettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok || d.HasChange("upload_time") {
		t, err := expandSystemLocallogFortianalyzer2SettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	return &obj, nil
}
