// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/jhabshoosh/etzer-api/internal/models"
)

type CreateChildInput struct {
	ChildName  string `json:"childName"`
	ParentID   string `json:"parentId"`
	ParentType string `json:"parentType"`
}

type CreateParentInput struct {
	ChildID    string `json:"childId"`
	ParentName string `json:"parentName"`
	ParentType string `json:"parentType"`
}

type CreatePersonInput struct {
	Name string `json:"name"`
}

type GetFamilyResponse struct {
	Persons       []*models.Person `json:"persons"`
	Relationships []*Relationship  `json:"relationships"`
}

type GetPersonInput struct {
	UUID string `json:"uuid"`
}

type Relationship struct {
	Parent     string `json:"parent"`
	Child      string `json:"child"`
	ParentType string `json:"parentType"`
}

type UpdateParentsInput struct {
	Child  string  `json:"child"`
	Father *string `json:"father"`
	Mother *string `json:"mother"`
}
