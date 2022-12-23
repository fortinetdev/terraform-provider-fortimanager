// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure connector.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemConnectorUpdate,
		Read:   resourceSystemConnectorRead,
		Update: resourceSystemConnectorUpdate,
		Delete: resourceSystemConnectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"conn_refresh_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fsso_refresh_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fsso_sess_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"px_refresh_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"px_svr_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemConnectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemConnector(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemConnector resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemConnector(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemConnector")

	return resourceSystemConnectorRead(d, m)
}

func resourceSystemConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemConnector(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemConnectorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemConnector(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemConnector resource from API: %v", err)
	}
	return nil
}

func flattenSystemConnectorConnRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConnectorFssoRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConnectorFssoSessTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConnectorPxRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConnectorPxSvrTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("conn_refresh_interval", flattenSystemConnectorConnRefreshInterval(o["conn-refresh-interval"], d, "conn_refresh_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-refresh-interval"], "SystemConnector-ConnRefreshInterval"); ok {
			if err = d.Set("conn_refresh_interval", vv); err != nil {
				return fmt.Errorf("Error reading conn_refresh_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_refresh_interval: %v", err)
		}
	}

	if err = d.Set("fsso_refresh_interval", flattenSystemConnectorFssoRefreshInterval(o["fsso-refresh-interval"], d, "fsso_refresh_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-refresh-interval"], "SystemConnector-FssoRefreshInterval"); ok {
			if err = d.Set("fsso_refresh_interval", vv); err != nil {
				return fmt.Errorf("Error reading fsso_refresh_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_refresh_interval: %v", err)
		}
	}

	if err = d.Set("fsso_sess_timeout", flattenSystemConnectorFssoSessTimeout(o["fsso-sess-timeout"], d, "fsso_sess_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsso-sess-timeout"], "SystemConnector-FssoSessTimeout"); ok {
			if err = d.Set("fsso_sess_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fsso_sess_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsso_sess_timeout: %v", err)
		}
	}

	if err = d.Set("px_refresh_interval", flattenSystemConnectorPxRefreshInterval(o["px-refresh-interval"], d, "px_refresh_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["px-refresh-interval"], "SystemConnector-PxRefreshInterval"); ok {
			if err = d.Set("px_refresh_interval", vv); err != nil {
				return fmt.Errorf("Error reading px_refresh_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading px_refresh_interval: %v", err)
		}
	}

	if err = d.Set("px_svr_timeout", flattenSystemConnectorPxSvrTimeout(o["px-svr-timeout"], d, "px_svr_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["px-svr-timeout"], "SystemConnector-PxSvrTimeout"); ok {
			if err = d.Set("px_svr_timeout", vv); err != nil {
				return fmt.Errorf("Error reading px_svr_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading px_svr_timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemConnectorConnRefreshInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConnectorFssoRefreshInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConnectorFssoSessTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConnectorPxRefreshInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConnectorPxSvrTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemConnector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("conn_refresh_interval"); ok || d.HasChange("conn_refresh_interval") {
		t, err := expandSystemConnectorConnRefreshInterval(d, v, "conn_refresh_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-refresh-interval"] = t
		}
	}

	if v, ok := d.GetOk("fsso_refresh_interval"); ok || d.HasChange("fsso_refresh_interval") {
		t, err := expandSystemConnectorFssoRefreshInterval(d, v, "fsso_refresh_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-refresh-interval"] = t
		}
	}

	if v, ok := d.GetOk("fsso_sess_timeout"); ok || d.HasChange("fsso_sess_timeout") {
		t, err := expandSystemConnectorFssoSessTimeout(d, v, "fsso_sess_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-sess-timeout"] = t
		}
	}

	if v, ok := d.GetOk("px_refresh_interval"); ok || d.HasChange("px_refresh_interval") {
		t, err := expandSystemConnectorPxRefreshInterval(d, v, "px_refresh_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["px-refresh-interval"] = t
		}
	}

	if v, ok := d.GetOk("px_svr_timeout"); ok || d.HasChange("px_svr_timeout") {
		t, err := expandSystemConnectorPxSvrTimeout(d, v, "px_svr_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["px-svr-timeout"] = t
		}
	}

	return &obj, nil
}
