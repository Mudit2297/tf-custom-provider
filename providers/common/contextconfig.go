package common

import (
	"context"

	"cpf/providers/client"

	"cpf"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ConfigContex = cpf.CustomConfigureContextFunc{
	ConfigureContextFunc: providerConfigure,
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var newClient client.Client
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	gitToken := d.Get("github_token").(string)
	if gitToken != "" {
		return newClient.GitHubNewClient(gitToken), diags
	} else if schema.EnvDefaultFunc("GH_TOKEN", nil) == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create git client",
			Detail:   "No credentials provided for authentication. Please add your credentials and try again.",
		})
		return nil, diags
	}

	// username := d.Get("username").(string)
	// password := d.Get("password").(string)
	// if (username != "") && (password != "") {
	// 	c, err := newClient.HashiCupsClient(nil, &username, &password)
	// 	if err != nil {
	// 		return nil, diag.FromErr(err)
	// 	}

	// 	return c, diags
	// }

	// c, err := newClient.HashiCupsClient(nil, nil, nil)
	// if err != nil {
	// 	return nil, diag.FromErr(err)
	// }

	return nil, diags
}
