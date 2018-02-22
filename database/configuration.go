package database

import (
	"github.com/neelance/graphql-go"
)

type configuration struct {
	UserID  graphql.ID
	Country string
}

var configurations = []*configuration{
	{
		UserID:  "1",
		Country: "NZ",
	},
}

var configurationData = make(map[graphql.ID]*configuration)

func init() {
	for _, c := range configurations {
		configurationData[c.UserID] = c
	}
}

// Resolver provides the mechanism to resolve queries.
type Resolver struct{}

// Configuration resolver.
func (r *Resolver) Configuration(args struct{ ID graphql.ID }) *configurationResolver {
	if c := configurationData[args.ID]; c != nil {
		return &configurationResolver{c}
	}
	return nil
}

type configurationResolver struct {
	c *configuration
}

func (r *configurationResolver) UserID() graphql.ID {
	return r.c.UserID
}

func (r *configurationResolver) Country() string {
	return r.c.Country
}
