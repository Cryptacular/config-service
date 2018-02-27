package schema

// Schema for the configuration GraphQL service.
var Schema = `
	schema {
		query: Query
	}
	type Query {
		configuration(id: ID!): Configuration
		abTests(id: ID!): [AbTest]!
	}
	type Configuration {
		id: ID!
		country: String!
		environment: String!
	}
	type AbTest {
		id: ID!
		userId: ID!
		name: String!
		segment: String!
		expiry: String!
	}
`
