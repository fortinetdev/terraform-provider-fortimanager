package fortimanager

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceJsonGenericAPI() *schema.Resource {
	return &schema.Resource{
		Create: createGeneric,
		Update: updateGeneric,
		Delete: deleteGeneric,
		Read:   schema.Noop,

		Schema: map[string]*schema.Schema{
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

	res, err := c.JsonGenericAPI(data)

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

	c := m.(*FortiClient).Client
	c.Retries = 1

	res, err := c.JsonGenericAPI(data)

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
