// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreatePersonInput struct {
	Name string `json:"name"`
}

type GetPersonInput struct {
	UUID string `json:"uuid"`
}

type UpdateParentsInput struct {
	Child  string  `json:"child"`
	Father *string `json:"father"`
	Mother *string `json:"mother"`
}