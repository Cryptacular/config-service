package schema

import (
	"github.com/graph-gophers/graphql-go"
	"time"
)

type configuration struct {
	ID          graphql.ID
	Country     string
	Environment string
}

var configurations = []*configuration{
	{
		ID:          "1",
		Country:     "NZ",
		Environment: "com",
	},
}

var configurationData = make(map[graphql.ID]*configuration)

func init() {
	for _, c := range configurations {
		configurationData[c.ID] = c
	}
}

type abTest struct {
	ID      graphql.ID
	UserID  graphql.ID
	Name    string
	Segment string
	Expiry  string
}

var abTests = []*abTest{
	{
		ID:      "1",
		UserID:  "1",
		Name:    "NZ",
		Segment: "com",
		Expiry:  time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC).String(),
	},
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

func (r *configurationResolver) ID() graphql.ID {
	return r.c.ID
}

func (r *configurationResolver) Country() string {
	return r.c.Country
}

func (r *configurationResolver) Environment() string {
	return r.c.Environment
}

// AbTests resolver.
func (r *Resolver) AbTests(args struct{ ID graphql.ID }) []*abTestsResolver {
	result := []*abTestsResolver{}

	for i := range abTests {
		a := abTests[i]
		if a.UserID == args.ID {
			result = append(result, &abTestsResolver{a})
		}
	}

	return result
}

type abTestsResolver struct {
	a *abTest
}

func (r *abTestsResolver) ID() graphql.ID {
	return r.a.ID
}

func (r *abTestsResolver) UserID() graphql.ID {
	return r.a.UserID
}

func (r *abTestsResolver) Name() string {
	return r.a.Name
}

func (r *abTestsResolver) Segment() string {
	return r.a.Segment
}

func (r *abTestsResolver) Expiry() string {
	return r.a.Expiry
}
