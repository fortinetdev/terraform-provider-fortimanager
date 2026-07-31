package fortimanager

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-fortimanager/sdk/sdkcore"
)

func resourceObjectCasbProfileSort() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileSortCreateUpdate,
		Read:   resourceObjectCasbProfileSortRead,
		Update: resourceObjectCasbProfileSortCreateUpdate,
		Delete: schema.Noop,

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
			"sortby": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					availableOptions := []string{"name"}
					var validValue bool
					for _, ele := range availableOptions {
						if ele == v {
							validValue = true
							break
						}
					}
					if !validValue {
						errs = append(errs, fmt.Errorf("%q must be one of the option of [\"name\"], got: \"%v\"", key, v))
					}
					return
				},
			},
			"sortdirection": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					availableOptions := []string{"ascending", "descending", "manual"}
					var validValue bool
					for _, ele := range availableOptions {
						if ele == v {
							validValue = true
							break
						}
					}
					if !validValue {
						errs = append(errs, fmt.Errorf("%q must be one of the option of [\"ascending\", \"descending\", \"manual\"], got: \"%v\"", key, v))
					}
					return
				},
			},
			"manual_order": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceObjectCasbProfileSortCreateUpdate(d *schema.ResourceData, m interface{}) (err error) {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiManager connection did not initialize successfully!")
	}

	c.Retries = 1

	sortby := d.Get("sortby").(string)
	sortdirection := d.Get("sortdirection").(string)
	manual_order_d := d.Get("manual_order").([]interface{})
	manual_order := make([]interface{}, len(manual_order_d))
	for cIndex, cValue := range manual_order_d {
		manual_order[cIndex] = fmt.Sprint(cValue)
	}

	if sortby != "name" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" && sortdirection != "manual" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv
	wsParams["adom"] = adomv

	var input_model forticlient.SortInputModel
	input_model.SortBy = sortby
	input_model.SortDirection = sortdirection
	input_model.ManualOrder = manual_order
	input_model.URLParams = paradict
	input_model.WSParams = wsParams
	err = c.CreateUpdateObjectCasbProfileSort(&input_model)
	if err != nil {
		return fmt.Errorf("Error sorting ObjectCasbProfile: %s", err)
	}

	d.SetId(sortby + sortdirection)

	return resourceObjectCasbProfileSortRead(d, m)
}

func resourceObjectCasbProfileSortRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiManager connection did not initialize successfully!")
	}

	c.Retries = 1

	sortby := d.Get("sortby").(string)
	sortdirection := d.Get("sortdirection").(string)
	manual_order_d := d.Get("manual_order").([]interface{})
	manual_order := make([]interface{}, len(manual_order_d))
	for cIndex, cValue := range manual_order_d {
		manual_order[cIndex] = fmt.Sprint(cValue)
	}

	if sortby != "name" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" && sortdirection != "manual" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv
	wsParams["adom"] = adomv

	var input_model forticlient.SortInputModel
	input_model.SortBy = sortby
	input_model.SortDirection = sortdirection
	input_model.ManualOrder = manual_order
	input_model.URLParams = paradict
	input_model.WSParams = wsParams

	sorted, o, err := c.ReadObjectCasbProfileSort(&input_model)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfile sort status: %s %s", err, mkey)
	}

	if sorted == false {
		d.Set("status", "unsorted")
	} else {
		d.Set("status", "")
	}

	if fr, ok := d.GetOk("force_recreate"); !ok || fr == "True" {
		d.Set("force_recreate", "False")
	}

	if o != nil {
		if err := d.Set("state_list", o); err != nil {
			log.Printf("[WARN] Error reading ObjectCasbProfile List for (%s): %s", d.Id(), err)
		}
	} else {
		d.Set("state_list", nil)
	}
	d.Set("manual_order", manual_order)

	return nil
}
