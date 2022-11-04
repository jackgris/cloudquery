// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func EnterpriseAccountMembers() *schema.Table {
	return &schema.Table{
		Name:        "heroku_enterprise_account_members",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#enterprise-account-member-attributes`,
		Resolver:    fetchEnterpriseAccountMembers,
		Columns: []schema.Column{
			{
				Name:     "enterprise_account",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnterpriseAccount"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "identity_provider",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProvider"),
			},
			{
				Name:     "permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Permissions"),
			},
			{
				Name:     "two_factor_authentication",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TwoFactorAuthentication"),
			},
			{
				Name:     "user",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("User"),
			},
		},
	}
}

func fetchEnterpriseAccountMembers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.EnterpriseAccount, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.EnterpriseAccountList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		items = append(items, v...)
	}

	for _, it := range items {
		nextRange = &heroku.ListRange{
			Field: "id",
			Max:   1000,
		}
		// Roundtripper middleware in client/pagination.go
		// sets the nextRange value after each request
		for nextRange.Max != 0 {
			ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
			v, err := c.Heroku.EnterpriseAccountMemberList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
