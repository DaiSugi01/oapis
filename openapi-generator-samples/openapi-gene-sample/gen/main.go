/*
 * Swagger Petstore
 *
 * A sample API that uses a petstore as an example to demonstrate features in the OpenAPI 3.0 specification
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	PetsApiService := openapi.NewPetsApiService()
	PetsApiController := openapi.NewPetsApiController(PetsApiService)

	UsersApiService := openapi.NewUsersApiService()
	UsersApiController := openapi.NewUsersApiController(UsersApiService)

	router := openapi.NewRouter(PetsApiController, UsersApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}