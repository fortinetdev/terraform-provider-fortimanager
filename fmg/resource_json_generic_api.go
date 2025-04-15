package fortimanager

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceJsonGenericAPI() *schema.Resource {
	return &schema.Resource{
		Create: createGeneric,
		Update: updateGeneric,
		Delete: deleteGeneric,
		Read:   schema.Noop,

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
			"json_content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createGeneric(d *schema.ResourceData, m interface{}) error {
	data := d.Get("json_content").(string)

	c := m.(*FortiClient).Client
	c.Retries = 1
	cfg := m.(*FortiClient).Cfg

	adomv, err := adomChecking(cfg, d)
	wsParams := make(map[string]string)
	wsParams["adom"] = adomv
	res, err := c.JsonGenericAPI(data, wsParams)

	if err != nil {
		return fmt.Errorf("Error createGeneric: %v", err)
	}

	d.Set("response", res)

	d.SetId("JSONRPC-Requst-" + uuid.New().String())

	return nil
}

func updateGeneric(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	data := d.Get("json_content").(string)
	cfg := m.(*FortiClient).Cfg
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := adomChecking(cfg, d)
	wsParams := make(map[string]string)
	wsParams["adom"] = adomv
	res, err := c.JsonGenericAPI(data, wsParams)

	if err != nil {
		return fmt.Errorf("Error updateGeneric: %v", err)
	}

	d.Set("response", res)

	d.SetId(mkey)

	return nil
}

func deleteGeneric(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}
