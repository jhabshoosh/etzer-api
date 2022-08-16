package graph

import "github.com/jhabshoosh/etzer-api/internal/services/person"

type Resolver struct {
	PersonService person.PersonService
}
