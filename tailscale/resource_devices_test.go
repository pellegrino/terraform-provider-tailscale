package tailscale_test

import (
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/davidsbond/tailscale-client-go/tailscale"
)

const testResourceDevice = `
	resource "tailscale_device" "test_device" {
		name = "testdevice.example.com"
	}`

func TestProvider_TestAssignResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		PreCheck: func() {
			testServer.ResponseCode = http.StatusOK
			testServer.ResponseBody = map[string][]tailscale.Device{}
		},
		ProviderFactories: testProviderFactories(t),
		Steps: []resource.TestStep{
			testResourceCreated("tailscale_device.test_device", testResourceDevice),
		},
	})
}
