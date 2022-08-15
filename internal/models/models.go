package models

import "github.com/mindstand/gogm/v2"

type 	Person struct {
	gogm.BaseUUIDNode

	UUID     string    `json:"uuid" gogm:"name=uuid"`
	Name     string    `json:"name" gogm:"name=name"`
	Parents	 []*Person `json:"parents" gogm:"direction=incoming;relationship=parent_of"`
	Children []*Person `json:"children" gogm:"direction=outgoing;relationship=parent_of"`
}