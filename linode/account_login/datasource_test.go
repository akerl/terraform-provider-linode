package account_login_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/linode/terraform-provider-linode/linode/acceptance"
	"github.com/linode/terraform-provider-linode/linode/account_login/tmpl"
)

func TestAccDataSourceLinodeAccountLogin_basic(t *testing.T) {
	t.Parallel()

	resourceName := "data.linode_account_login.foobar"

	client, err := acceptance.GetClientForSweepers()
	if err != nil {
		t.Fail()
		t.Log("Failed to get testing client.")
	}

	logins, err := client.ListLogins(context.TODO(), nil)
	login := logins[0]
	accountID := login.ID

	if err != nil {
		t.Fail()
		t.Log("Failed to get testing login.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: tmpl.DataBasic(t, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", strconv.Itoa(login.ID)),
					resource.TestCheckResourceAttr(resourceName, "ip", login.IP),
					resource.TestCheckResourceAttr(resourceName, "username", login.Username),
					resource.TestCheckResourceAttr(resourceName, "datetime", login.Datetime.Format(time.RFC3339)),
					resource.TestCheckResourceAttr(resourceName, "restricted", strconv.FormatBool(login.Restricted)),
				),
			},
		},
	})
}
