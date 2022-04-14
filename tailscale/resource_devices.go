package tailscale

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidsbond/tailscale-client-go/tailscale"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const NON_EXISTING_DEVICE = "-1"

func resourceDevice() *schema.Resource {
	return &schema.Resource{
		Description:   "The device resource allows control tailscale devices in your Tailscale tailnet. See https://tailscale.com/kb/1017/install/ for more information.",
		ReadContext:   resourceDeviceRead,
		CreateContext: resourceDeviceCreate,
		UpdateContext: resourceDeviceUpdate,
		DeleteContext: resourceDeviceDelete,
		Schema: map[string]*schema.Schema{
			"device_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID for the device",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the device",
			},
			"user": {
				Type:        schema.TypeString,
				Description: "The user associated with the device",
				Computed:    true,
			},
			"addresses": {
				Type:        schema.TypeList,
				Description: "The list of device's IPs",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// resourceDeviceUpdate there isn't support for updating, so it will only read the resource
func resourceDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChange("name") {
		d.Set("addresses", nil)
		d.SetId("")
		d.Set("user", "")
	}
	d.SetId("")
	return resourceDeviceRead(ctx, d, m)

}

// resourceDeviceCreate will perform a read and store the value as an ID.
// it doesn't create a device as tailscale doesn't support it
func resourceDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tailscale.Client)
	deviceName := d.Get("name").(string)

	// will find the corresponding device and register its id
	device, err := findDevice(ctx, client, deviceName)
	if err != nil {
		d.SetId(NON_EXISTING_DEVICE)
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Device isn't registered yet",
			Detail:   "This device hasn't been registered yet. It must be created by calling `tailscale up` from the running instance.",
		}}
	}

	//if err != nil {
	//	d.Set("name", deviceName)
	//	d.SetId("")
	//	return diag.Diagnostics{
	//		diag.Diagnostic{
	//			Severity: diag.Warning,
	//			Summary:  "Device isn't registered yet",
	//			Detail:   "This device hasn't been registered yet. It must be created by calling `tailscale up` from the running instance.",
	//		},
	//	}
	//}
	// stores the retrieved id as ID
	d.SetId(device.ID)
	d.Set("name", device.Name)
	d.Set("user", device.User)
	d.Set("addresses", device.Addresses)
	return nil
}

// findDevice gets all the devices for a user account and locates the device based on its name
// will return all devices at every invocation. Might add some rudimentary caching if needed
func findDevice(ctx context.Context, client *tailscale.Client, deviceName string) (*tailscale.Device, error) {
	devices, err := client.Devices(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch devices")
	}

	var selected *tailscale.Device
	for _, device := range devices {
		if device.Name != deviceName {
			continue
		}

		selected = &device
		break
	}

	if selected == nil {
		return nil, errors.New(fmt.Sprintf("could not find device with name %s", deviceName))
	}

	return selected, nil
}

// resourceDeviceRead finds a device given its name and stores its ID in terraform state.
func resourceDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tailscale.Client)
	deviceName := d.Get("name").(string)

	// will find the corresponding device and register its id
	device, err := findDevice(ctx, client, deviceName)
	if err != nil {
		d.Set("name", deviceName)
		d.SetId("")
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Device isn't registered yet",
				Detail:   "This device hasn't been registered yet. It must be created by calling `tailscale up` from the running instance.",
			},
		}
	}
	// stores the retrieved id as ID
	d.SetId(device.ID)
	d.Set("name", device.Name)
	d.Set("user", device.User)
	d.Set("addresses", device.Addresses)
	return nil
}

// resourceDeviceDelete deletes the device from the user account based on it's ID
func resourceDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*tailscale.Client)
	id := d.Id()

	err := client.DeleteDevice(ctx, id)

	if tailscale.IsNotFound(err) {
		return nil
	}

	if err != nil {
		return diagnosticsError(err, "Failed to delete device")
	}

	return nil
}
