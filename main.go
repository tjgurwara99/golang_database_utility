package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go/v4"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"golang.org/x/crypto/bcrypt"
)

// RootResolver structure for
type RootResolver struct{}

// Info method that acts like a field for graphql schema
func (r *RootResolver) Info() (string, error) {
	return "this is a thing", nil
}

// Link struct for returning []Link in Feed() resolver
type Link struct {
	ID          graphql.ID
	Description string
	URL         string
}

// Testing `links` data
var links = []Link{
	{
		ID:          "0",
		URL:         "www.howtographql.com",
		Description: "Fullstack tutorial for GraphQL",
	},
	{
		ID:          "1",
		Description: "TEsting 123",
		URL:         "localhost",
	},
}

// Feed Method that acts like a field and resolver for graphql schema
func (r *RootResolver) Feed() ([]Link, error) {
	return links, nil
}

var opts = []graphql.SchemaOpt{
	graphql.UseFieldResolvers(),
}

// Post Method that acts like a field and resolver for graphql schema
func (r *RootResolver) Post(args struct {
	URL         string
	Description string
}) (Link, error) {
	newLink := Link{
		ID: graphql.ID(fmt.Sprintf("%d", len(links))),
	}
	links = append(links, newLink)
	return newLink, nil
}

// User Struct for authentication
type User struct {
	ID        graphql.ID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Links     []Link
}

// test data for users
var users = []User{}

// AuthPayload for jwt token authentication
type AuthPayload struct {
	Token *string
	User  *User
}

// Signup form graphql resolver
func (r *RootResolver) Signup(args struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}) (*AuthPayload, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(args.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	newUser := User{
		ID:        graphql.ID(fmt.Sprintf("%d", len(users))),
		Email:     args.Email,
		FirstName: args.FirstName,
		LastName:  args.LastName,
		Password:  string(passwordHash),
	}
	users = append(users, newUser)

	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("supersecret"))

	// // generate token
	// token := jwt.NewWithClaim(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	ID: newUser.ID,
	// })
	// tokenString, err := token.SignedString([]byte("VerySecret"))
	if err != nil {
		return nil, err
	}
	return &AuthPayload{
		Token: &tokenString,
		User:  &newUser,
	}, nil
}

func parseSchema(path string, resolver interface{}) *graphql.Schema {
	byteString, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err) // handle this better
	}
	schemaString := string(byteString)
	parsedSchema, err := graphql.ParseSchema(
		schemaString,
		resolver,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	return parsedSchema
}

func main() {
	http.Handle("/graphql", &relay.Handler{
		Schema: parseSchema("./graphqltesting/schema.graphql", &RootResolver{}),
	})
	fmt.Printf("Serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
