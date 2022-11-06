// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/{{.Service}}"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func {{.Table.Resolver}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input {{.Service}}.Describe{{.StructName}}sInput
	c := meta.(*client.Client)
	svc := c.Services().{{.Service | ToCamel}}
	for {
		response, err := svc.Describe{{.StructName}}s(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.{{.StructName}}s
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}