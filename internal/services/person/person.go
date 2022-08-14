package person

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jhabshoosh/etzer-api/internal/db"
	"github.com/jhabshoosh/etzer-api/internal/graph/model"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type PersonService struct {
	Neo4jClient	db.Neo4JDB
}

func (ps *PersonService) FetchPersons(ctx context.Context) ([]*model.Person, error) {
	query := `
		match (p:Person) return p.id, p.name
	`

	session, err := ps.Neo4jClient.Connection.Session(neo4j.AccessModeRead)
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
		db.ParseCypherQueryResult(result.Record(), "p", &person)

		persons = append(persons, &person)
	}

	return persons, err
}

func (ps *PersonService) CreatePerson(ctx context.Context, input model.CreatePersonInput) (*model.Person, error) {
	session, err := ps.Neo4jClient.Connection.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}
	defer session.Close()	

	query := `CREATE (p: Person {name: $name, id: $id}) RETURN p.name, p.id`
	parameters := map[string]interface{}{
		"name": input.Name,
		"id": 	uuid.NewString(),
	}
	
	result, err := session.Run(query, parameters)
	if err != nil {
		log.Panic("Cannot add person", err)
	}
	
	log.Println("CYPHER_QUERY", query)

	result.Next()
	if err = result.Err(); err != nil {
		log.Panic("Cannot iterate through person results", err)
	}

	person := model.Person{}
	db.ParseCypherQueryResult(result.Record(), "p", &person)

	return &person, err


}

func getParentRelationship(parentType string) string {

	if strings.ToLower(parentType) == "mother" {
		return "MOTHER_OF"
	}
	return "FATHER_OF"
}

func (ps *PersonService) UpdateParent(ctx context.Context, input model.UpdateParentInput) (string, error) {
	session, err := ps.Neo4jClient.Connection.Session(neo4j.AccessModeWrite)
	if err != nil {
		return "", err
	}
	defer session.Close()

	relationshipType := getParentRelationship(input.ParentType)

	query := fmt.Sprintf(`
		MATCH
			(p:Person),
			(c:Person)
 		WHERE p.id = $parentId AND c.id = $childId
  		CREATE (p)-[r:%s]->(c)
  		RETURN type(r)
	`, relationshipType)
	parameters := map[string]interface{}{
		"childId": input.ChildID,
		"parentId": input.ParentID,
	}
	result, err := session.Run(query, parameters)
	if err != nil {
		log.Panic("Cannot add parent relationship", err)
	}
	
	log.Println("CYPHER_QUERY", query)

	result.Next()
	if err = result.Err(); err != nil {
		log.Panic("Cannot iterate through relationship results", err)
	}
	return input.ChildID, nil
}

func (ps *PersonService) GetParent(ctx context.Context, child *model.Person, parentType string) (*model.Person, error) {
	relationshipType := getParentRelationship(parentType)
	query := fmt.Sprintf(`
		MATCH (c:Person {id: $childId})<-[:%s]-(p:Person)
	RETURN p
	`, relationshipType)
	parameters := map[string]interface{}{
		"childId": child.ID,
	}

	session, err := ps.Neo4jClient.Connection.Session(neo4j.AccessModeRead)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	result, err := session.Run(query, parameters)
	if err != nil {
		log.Panic("Cannot find parent", err)
	}

	log.Println("CYPHER_QUERY", query, parameters)

	person := model.Person{}
	result.Next()
	db.ParseCypherQueryResult(result.Record(), "p", &person)

	return &person, err
}

// func (ps *PersonService) FindPersonByID(ctx context.Context, id string) (*model.Person, error) {
// 	query := `
// 		match (p:Person) where p.id = $id return p.id, p.name
// 	`
// 	session, err := r.Connection.Session(neo4j.AccessModeWrite)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.Close()

// 	args := map[string]interface{}{
// 		"id": id,
// 	}

// 	result, err := session.Run(query, args)

// 	if err != nil {
// 		log.Panic("Cannot find person by id", id, err)
// 	}

// 	log.Println("CYPHER_QUERY", query,  args)

// 	person := models.Person{}
	
// 	for result.Next() {
// 		ParseCypherQueryResult(result.Record(), "m", &person)
// 	}

// 	return &person, err
// }
