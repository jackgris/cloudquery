package storage

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "azure_storage_accounts",
		Resolver:    fetchAccounts,
		Description: "https://learn.microsoft.com/en-us/rest/api/storagerp/storage-accounts/list?tabs=HTTP#storageaccount",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_storage_accounts", client.Namespacemicrosoft_storage),
		Transform:   transformers.TransformWithStruct(&armstorage.Account{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},

		Relations: []*schema.Table{
			tables(),
			containers(),
			blob_services(),
		},
	}
}

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armstorage.NewAccountsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
