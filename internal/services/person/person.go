package person

import (
	"context"

	"github.com/jhabshoosh/etzer-api/internal/graph/model"
	"github.com/jhabshoosh/etzer-api/internal/models"
	"github.com/mindstand/gogm/v2"
)

type PersonService struct {
	Ogm		gogm.Gogm
}

// func (ps *PersonService) FetchPersons(ctx context.Context) ([]*models.Person, error) {
	// return nil, nil
	// sess, err := ps.Neo4JOGM.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	// if err != nil {
	// 	panic(err)
	// }

	// sess.LoadAll()


	// query := `
	// 	match (p:Person) return p.id, p.name
	// `

	// session, err := ps.Neo4jClient.Connection.Session(neo4j.AccessModeRead)
	// if err != nil {
	// 	return nil, err
	// }
	// defer session.Close()

	// result, err := session.Run(query, nil)
	// if err != nil {
	// 	log.Panic("Cannot find persons", err)
	// }

	// log.Println("CYPHER_QUERY", query)

	// var persons []*models.Person

	// for result.Next() {
	// 	person := models.Person{}
	// 	db.ParseCypherQueryResult(result.Record(), "p", &person)

	// 	persons = append(persons, &person)
	// }

	// return persons, err
// }

func (ps *PersonService) CreatePerson(ctx context.Context, input model.CreatePersonInput) (*models.Person, error) {

	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	newPerson := &models.Person{
		Name: 	input.Name,
	}

	err = sess.Save(context.Background(), newPerson)
	if err != nil {
		panic(err)
	}

	var readin models.Person
	err = sess.Load(context.Background(), &readin, newPerson.UUID)
	if err != nil {
		panic(err)
	}

	return &readin, err
}

func (ps *PersonService) GetPerson(ctx context.Context, input model.GetPersonInput) (*models.Person, error) {
	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	var readin models.Person
	err = sess.Load(context.Background(), &readin, input.UUID)
	if err != nil {
		panic(err)
	}

	return &readin, err
}

func (ps *PersonService) Parents(ctx context.Context, obj *models.Person) ([]*models.Person, error) {
	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	var readin models.Person
	err = sess.Load(context.Background(), &readin, obj.UUID)
	if err != nil {
		panic(err)
	}

	return readin.Parents, err
}


func (ps *PersonService) Children(ctx context.Context, obj *models.Person) ([]*models.Person, error) {
	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	var readin models.Person
	err = sess.Load(context.Background(), &readin, obj.UUID)
	if err != nil {
		panic(err)
	}

	return readin.Children, err
}

func (ps *PersonService) UpdateParents(ctx context.Context, input model.UpdateParentsInput) (string, error) {
	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	var childPerson models.Person
	err = sess.Load(context.Background(), &childPerson, input.Child)
	if err != nil {
		panic(err)
	}

	for _, p := range input.Parents {
		var parentPerson models.Person
		err = sess.Load(context.Background(), &parentPerson, p)
		if err != nil {
			panic(err)
		}
		childPerson.LinkToPersonOnFieldParents(&parentPerson)
	}

	err = sess.Save(context.Background(), &childPerson)
	if err != nil {
		panic(err)
	}
	
	return input.Child, err
}

// func (ps *PersonService) GetParent(ctx context.Context, child *models.Person, parentType string) (*models.Person, error) {
// 	relationshipType := getParentRelationship(parentType)
// 	query := fmt.Sprintf(`
// 		MATCH (c:Person {id: $childId})<-[:%s]-(p:Person)
// 	RETURN p
// 	`, relationshipType)
// 	parameters := map[string]interface{}{
// 		"childId": child.ID,
// 	}

// 	session, err := ps.Neo4jClient.Connection.Session(neo4j.AccessModeRead)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.Close()	

// 	result, err := session.Run(query, parameters)
// 	if err != nil {
// 		log.Panic("Cannot find parent", err)
// 	}

// 	log.Println("CYPHER_QUERY", query, parameters)

// 	person := model.Person{}
// 	result.Next()
// 	db.ParseCypherQueryResult(result.Record(), "p", &person)

// 	return &person, err
// }

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
