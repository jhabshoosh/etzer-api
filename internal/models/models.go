package models

import (
	"fmt"
	"reflect"

	"github.com/mindstand/gogm/v2"
)

type Person struct {
	gogm.BaseUUIDNode

	UUID     string      `json:"uuid" gogm:"name=uuid"`
	Name     string      `json:"name" gogm:"name=name"`
	Parents  []*ParentOf `json:"parents" gogm:"direction=incoming;relationship=parent_of"`
	Children []*ParentOf `json:"children" gogm:"direction=outgoing;relationship=parent_of"`
}

type ParentType string

const (
	Father ParentType = "FATHER"
	Mother ParentType = "MOTHER"
)

// ParentOf implements Edge
type ParentOf struct {
	// provides required node fields
	gogm.BaseUUIDNode

	Parent     *Person
	Child      *Person
	ParentType ParentType `gogm:"name=parent_type"`
}

func (e *ParentOf) GetStartNode() interface{} {
	return e.Parent
}

func (e *ParentOf) GetStartNodeType() reflect.Type {
	return reflect.TypeOf(&Person{})
}

func (e *ParentOf) SetStartNode(v interface{}) error {
	val, ok := v.(*Person)
	if !ok {
		return fmt.Errorf("unable to cast [%T] to *VertexA", v)
	}

	e.Parent = val
	return nil
}

func (e *ParentOf) GetEndNode() interface{} {
	return e.Child
}

func (e *ParentOf) GetEndNodeType() reflect.Type {
	return reflect.TypeOf(&Person{})
}

func (e *ParentOf) SetEndNode(v interface{}) error {
	val, ok := v.(*Person)
	if !ok {
		return fmt.Errorf("unable to cast [%T] to *VertexB", v)
	}

	e.Child = val
	return nil
}
