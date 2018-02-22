package database

// Schema for the configuration GraphQL service.
var Schema = `
	schema {
		query: Query
	}
	type Query {
		configuration(id: ID!): Configuration
	}
	type Configuration {
		userId: ID!
		country: String!
	}
`
