package db

import (
	"context"
	"log"

	"github.com/jhabshoosh/etzer-api/internal/graph/model"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func (r *Neo4JDB) FetchPersons(ctx context.Context) ([]*model.Person, error) {
	query := `
		match (p:Person) return p.id, p.name, m.released, m.tagline
	`


	session, err := r.Connection.Session(neo4j.AccessModeWrite)

	if err != nil {
		return nil, err
	}

	defer session.Close()

	result, err := session.Run(query, nil)
	if err != nil {
		log.Panic("Cannot find persons", err)
	}

	log.Println("CYPHER_QUERY", query)

	var persons []*model.Person

	for result.Next() {
		person := model.Person{}
		ParseCypherQueryResult(result.Record(), "p", &person)

		persons = append(persons, &person)
	}

	return persons, err
}

func (r *Neo4JDB) FindPersonsByID(ctx context.Context, id string) (*model.Person, error) {
	query := `
		match (p:Person) where p.id = $id return p.id, p.name
	`
	session, err := r.Connection.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	args := map[string]interface{}{
		"id": id,
	}

	result, err := session.Run(query, args)

	if err != nil {
		log.Panic("Cannot find person by id", id, err)
	}

	log.Println("CYPHER_QUERY", query,  args)

	person := model.Person{}
	
	for result.Next() {
		ParseCypherQueryResult(result.Record(), "m", &person)
	}

	return &person, err
}