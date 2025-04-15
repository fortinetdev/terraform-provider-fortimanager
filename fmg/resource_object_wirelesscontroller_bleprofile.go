// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Bluetooth Low Energy profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerBleProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerBleProfileCreate,
		Read:   resourceObjectWirelessControllerBleProfileRead,
		Update: resourceObjectWirelessControllerBleProfileUpdate,
		Delete: resourceObjectWirelessControllerBleProfileDelete,

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
			"advertising": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"beacon_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ble_scanning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eddystone_instance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eddystone_namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eddystone_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eddystone_url_encode_hex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ibeacon_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"major_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minor_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"scan_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"txpower": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerBleProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerBleProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerBleProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerBleProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerBleProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerBleProfileRead(d, m)
}

func resourceObjectWirelessControllerBleProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectWirelessControllerBleProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerBleProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerBleProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerBleProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerBleProfileRead(d, m)
}

func resourceObjectWirelessControllerBleProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectWirelessControllerBleProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerBleProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerBleProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerBleProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerBleProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerBleProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerBleProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerBleProfileAdvertising(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerBleProfileBeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileBleScanning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileEddystoneInstance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileEddystoneNamespace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileEddystoneUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileEddystoneUrlEncodeHex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileIbeaconUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileMajorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileMinorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileScanInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileScanPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileScanThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileScanTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileScanType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileScanWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBleProfileTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerBleProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("advertising", flattenObjectWirelessControllerBleProfileAdvertising(o["advertising"], d, "advertising")); err != nil {
		if vv, ok := fortiAPIPatch(o["advertising"], "ObjectWirelessControllerBleProfile-Advertising"); ok {
			if err = d.Set("advertising", vv); err != nil {
				return fmt.Errorf("Error reading advertising: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading advertising: %v", err)
		}
	}

	if err = d.Set("beacon_interval", flattenObjectWirelessControllerBleProfileBeaconInterval(o["beacon-interval"], d, "beacon_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["beacon-interval"], "ObjectWirelessControllerBleProfile-BeaconInterval"); ok {
			if err = d.Set("beacon_interval", vv); err != nil {
				return fmt.Errorf("Error reading beacon_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beacon_interval: %v", err)
		}
	}

	if err = d.Set("ble_scanning", flattenObjectWirelessControllerBleProfileBleScanning(o["ble-scanning"], d, "ble_scanning")); err != nil {
		if vv, ok := fortiAPIPatch(o["ble-scanning"], "ObjectWirelessControllerBleProfile-BleScanning"); ok {
			if err = d.Set("ble_scanning", vv); err != nil {
				return fmt.Errorf("Error reading ble_scanning: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ble_scanning: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectWirelessControllerBleProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerBleProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("eddystone_instance", flattenObjectWirelessControllerBleProfileEddystoneInstance(o["eddystone-instance"], d, "eddystone_instance")); err != nil {
		if vv, ok := fortiAPIPatch(o["eddystone-instance"], "ObjectWirelessControllerBleProfile-EddystoneInstance"); ok {
			if err = d.Set("eddystone_instance", vv); err != nil {
				return fmt.Errorf("Error reading eddystone_instance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eddystone_instance: %v", err)
		}
	}

	if err = d.Set("eddystone_namespace", flattenObjectWirelessControllerBleProfileEddystoneNamespace(o["eddystone-namespace"], d, "eddystone_namespace")); err != nil {
		if vv, ok := fortiAPIPatch(o["eddystone-namespace"], "ObjectWirelessControllerBleProfile-EddystoneNamespace"); ok {
			if err = d.Set("eddystone_namespace", vv); err != nil {
				return fmt.Errorf("Error reading eddystone_namespace: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eddystone_namespace: %v", err)
		}
	}

	if err = d.Set("eddystone_url", flattenObjectWirelessControllerBleProfileEddystoneUrl(o["eddystone-url"], d, "eddystone_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["eddystone-url"], "ObjectWirelessControllerBleProfile-EddystoneUrl"); ok {
			if err = d.Set("eddystone_url", vv); err != nil {
				return fmt.Errorf("Error reading eddystone_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eddystone_url: %v", err)
		}
	}

	if err = d.Set("eddystone_url_encode_hex", flattenObjectWirelessControllerBleProfileEddystoneUrlEncodeHex(o["eddystone-url-encode-hex"], d, "eddystone_url_encode_hex")); err != nil {
		if vv, ok := fortiAPIPatch(o["eddystone-url-encode-hex"], "ObjectWirelessControllerBleProfile-EddystoneUrlEncodeHex"); ok {
			if err = d.Set("eddystone_url_encode_hex", vv); err != nil {
				return fmt.Errorf("Error reading eddystone_url_encode_hex: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eddystone_url_encode_hex: %v", err)
		}
	}

	if err = d.Set("ibeacon_uuid", flattenObjectWirelessControllerBleProfileIbeaconUuid(o["ibeacon-uuid"], d, "ibeacon_uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ibeacon-uuid"], "ObjectWirelessControllerBleProfile-IbeaconUuid"); ok {
			if err = d.Set("ibeacon_uuid", vv); err != nil {
				return fmt.Errorf("Error reading ibeacon_uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ibeacon_uuid: %v", err)
		}
	}

	if err = d.Set("major_id", flattenObjectWirelessControllerBleProfileMajorId(o["major-id"], d, "major_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["major-id"], "ObjectWirelessControllerBleProfile-MajorId"); ok {
			if err = d.Set("major_id", vv); err != nil {
				return fmt.Errorf("Error reading major_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading major_id: %v", err)
		}
	}

	if err = d.Set("minor_id", flattenObjectWirelessControllerBleProfileMinorId(o["minor-id"], d, "minor_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["minor-id"], "ObjectWirelessControllerBleProfile-MinorId"); ok {
			if err = d.Set("minor_id", vv); err != nil {
				return fmt.Errorf("Error reading minor_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minor_id: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerBleProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerBleProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("scan_interval", flattenObjectWirelessControllerBleProfileScanInterval(o["scan-interval"], d, "scan_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-interval"], "ObjectWirelessControllerBleProfile-ScanInterval"); ok {
			if err = d.Set("scan_interval", vv); err != nil {
				return fmt.Errorf("Error reading scan_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_interval: %v", err)
		}
	}

	if err = d.Set("scan_period", flattenObjectWirelessControllerBleProfileScanPeriod(o["scan-period"], d, "scan_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-period"], "ObjectWirelessControllerBleProfile-ScanPeriod"); ok {
			if err = d.Set("scan_period", vv); err != nil {
				return fmt.Errorf("Error reading scan_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_period: %v", err)
		}
	}

	if err = d.Set("scan_threshold", flattenObjectWirelessControllerBleProfileScanThreshold(o["scan-threshold"], d, "scan_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-threshold"], "ObjectWirelessControllerBleProfile-ScanThreshold"); ok {
			if err = d.Set("scan_threshold", vv); err != nil {
				return fmt.Errorf("Error reading scan_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_threshold: %v", err)
		}
	}

	if err = d.Set("scan_time", flattenObjectWirelessControllerBleProfileScanTime(o["scan-time"], d, "scan_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-time"], "ObjectWirelessControllerBleProfile-ScanTime"); ok {
			if err = d.Set("scan_time", vv); err != nil {
				return fmt.Errorf("Error reading scan_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_time: %v", err)
		}
	}

	if err = d.Set("scan_type", flattenObjectWirelessControllerBleProfileScanType(o["scan-type"], d, "scan_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-type"], "ObjectWirelessControllerBleProfile-ScanType"); ok {
			if err = d.Set("scan_type", vv); err != nil {
				return fmt.Errorf("Error reading scan_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_type: %v", err)
		}
	}

	if err = d.Set("scan_window", flattenObjectWirelessControllerBleProfileScanWindow(o["scan-window"], d, "scan_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-window"], "ObjectWirelessControllerBleProfile-ScanWindow"); ok {
			if err = d.Set("scan_window", vv); err != nil {
				return fmt.Errorf("Error reading scan_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_window: %v", err)
		}
	}

	if err = d.Set("txpower", flattenObjectWirelessControllerBleProfileTxpower(o["txpower"], d, "txpower")); err != nil {
		if vv, ok := fortiAPIPatch(o["txpower"], "ObjectWirelessControllerBleProfile-Txpower"); ok {
			if err = d.Set("txpower", vv); err != nil {
				return fmt.Errorf("Error reading txpower: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading txpower: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerBleProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerBleProfileAdvertising(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerBleProfileBeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileBleScanning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileEddystoneInstance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileEddystoneNamespace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileEddystoneUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileEddystoneUrlEncodeHex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileIbeaconUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileMajorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileMinorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileScanInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileScanPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileScanThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileScanTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileScanType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileScanWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBleProfileTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerBleProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("advertising"); ok || d.HasChange("advertising") {
		t, err := expandObjectWirelessControllerBleProfileAdvertising(d, v, "advertising")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertising"] = t
		}
	}

	if v, ok := d.GetOk("beacon_interval"); ok || d.HasChange("beacon_interval") {
		t, err := expandObjectWirelessControllerBleProfileBeaconInterval(d, v, "beacon_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-interval"] = t
		}
	}

	if v, ok := d.GetOk("ble_scanning"); ok || d.HasChange("ble_scanning") {
		t, err := expandObjectWirelessControllerBleProfileBleScanning(d, v, "ble_scanning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-scanning"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandObjectWirelessControllerBleProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_instance"); ok || d.HasChange("eddystone_instance") {
		t, err := expandObjectWirelessControllerBleProfileEddystoneInstance(d, v, "eddystone_instance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-instance"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_namespace"); ok || d.HasChange("eddystone_namespace") {
		t, err := expandObjectWirelessControllerBleProfileEddystoneNamespace(d, v, "eddystone_namespace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-namespace"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_url"); ok || d.HasChange("eddystone_url") {
		t, err := expandObjectWirelessControllerBleProfileEddystoneUrl(d, v, "eddystone_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-url"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_url_encode_hex"); ok || d.HasChange("eddystone_url_encode_hex") {
		t, err := expandObjectWirelessControllerBleProfileEddystoneUrlEncodeHex(d, v, "eddystone_url_encode_hex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-url-encode-hex"] = t
		}
	}

	if v, ok := d.GetOk("ibeacon_uuid"); ok || d.HasChange("ibeacon_uuid") {
		t, err := expandObjectWirelessControllerBleProfileIbeaconUuid(d, v, "ibeacon_uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibeacon-uuid"] = t
		}
	}

	if v, ok := d.GetOk("major_id"); ok || d.HasChange("major_id") {
		t, err := expandObjectWirelessControllerBleProfileMajorId(d, v, "major_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["major-id"] = t
		}
	}

	if v, ok := d.GetOk("minor_id"); ok || d.HasChange("minor_id") {
		t, err := expandObjectWirelessControllerBleProfileMinorId(d, v, "minor_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minor-id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerBleProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("scan_interval"); ok || d.HasChange("scan_interval") {
		t, err := expandObjectWirelessControllerBleProfileScanInterval(d, v, "scan_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-interval"] = t
		}
	}

	if v, ok := d.GetOk("scan_period"); ok || d.HasChange("scan_period") {
		t, err := expandObjectWirelessControllerBleProfileScanPeriod(d, v, "scan_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-period"] = t
		}
	}

	if v, ok := d.GetOk("scan_threshold"); ok || d.HasChange("scan_threshold") {
		t, err := expandObjectWirelessControllerBleProfileScanThreshold(d, v, "scan_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-threshold"] = t
		}
	}

	if v, ok := d.GetOk("scan_time"); ok || d.HasChange("scan_time") {
		t, err := expandObjectWirelessControllerBleProfileScanTime(d, v, "scan_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-time"] = t
		}
	}

	if v, ok := d.GetOk("scan_type"); ok || d.HasChange("scan_type") {
		t, err := expandObjectWirelessControllerBleProfileScanType(d, v, "scan_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-type"] = t
		}
	}

	if v, ok := d.GetOk("scan_window"); ok || d.HasChange("scan_window") {
		t, err := expandObjectWirelessControllerBleProfileScanWindow(d, v, "scan_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-window"] = t
		}
	}

	if v, ok := d.GetOk("txpower"); ok || d.HasChange("txpower") {
		t, err := expandObjectWirelessControllerBleProfileTxpower(d, v, "txpower")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["txpower"] = t
		}
	}

	return &obj, nil
}
