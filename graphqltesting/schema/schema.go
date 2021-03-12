// Package schema define graphql schema in GoLang
package schema

// import "github.com/graphql-go/graphql"

// Product Struct to represent products offered by a Company
type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	SalePrice float64 `json:"sale_price"`
	CostPrice float64 `json:"cost_price"`
	Stock     int     `json:"stock"`
}

// productType definition of graphql todoType in GoLang
// var productType = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "Todo",
// 	Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.Int,
// 		},
// 		"name": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"sale_price": &graphql.Field{
// 			Type: graphql.Float,
// 		},
// 		"cost_price": &graphql.Field{
// 			Type: graphql.Float,
// 		},
// 		"stock": &graphql.Field{
// 			Type: graphql.Int,
// 		},
// 	},
// })

// // ProductSchema Schema Definition with rootQuery and rootMutation
// var ProductSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
// 	Query:    rootQuery,
// 	Mutation: rootMutation,
// })
