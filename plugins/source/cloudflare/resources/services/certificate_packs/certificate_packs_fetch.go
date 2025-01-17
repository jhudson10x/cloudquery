package certificate_packs

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCertificatePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	packs, err := svc.ClientApi.ListCertificatePacks(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- packs
	return nil
}
