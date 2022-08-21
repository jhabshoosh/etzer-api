package person

import (
	"context"

	"github.com/jhabshoosh/etzer-api/internal/graph/model"
	"github.com/jhabshoosh/etzer-api/internal/models"
	"github.com/mindstand/gogm/v2"
)

type PersonService struct {
	Ogm gogm.Gogm
}

func (ps *PersonService) CreatePerson(ctx context.Context, input model.CreatePersonInput) (*models.Person, error) {

	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	newPerson := &models.Person{
		Name: input.Name,
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

	return getParentsFromParentOf(readin.Parents), err
}

func getParentsFromParentOf(parentOf []*models.ParentOf) []*models.Person {
	var parents []*models.Person
	for _, p := range parentOf {
		parents = append(parents, p.Parent)
	}
	return parents
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

	return getChildrenFromParentOf(readin.Children), err
}

func getChildrenFromParentOf(parentOf []*models.ParentOf) []*models.Person {
	var children []*models.Person
	for _, p := range parentOf {
		children = append(children, p.Child)
	}
	return children
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

	if input.Father != nil {
		var fatherPerson models.Person
		err = sess.Load(context.Background(), &fatherPerson, input.Father)
		if err != nil {
			panic(err)
		}
		parentOf := models.ParentOf{
			Parent:     &fatherPerson,
			Child:      &childPerson,
			ParentType: models.Father,
		}
		childPerson.LinkToPersonOnFieldParents(&fatherPerson, &parentOf)
	}

	if input.Mother != nil {
		var motherPerson models.Person
		err = sess.Load(context.Background(), &motherPerson, input.Mother)
		if err != nil {
			panic(err)
		}
		parentOf := models.ParentOf{
			Parent:     &motherPerson,
			Child:      &childPerson,
			ParentType: models.Mother,
		}
		childPerson.LinkToPersonOnFieldParents(&motherPerson, &parentOf)
	}

	err = sess.Save(context.Background(), &childPerson)
	if err != nil {
		panic(err)
	}

	return input.Child, err
}

func (ps *PersonService) GetRootAncestor(ctx context.Context) (*models.Person, error) {
	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	var readin models.Person
	err = sess.Query(context.Background(), "MATCH (p:Person)WHERE NOT (p)<-[:parent_of]-(:Person) RETURN p", nil, &readin)

	return &readin, err

}

func (ps *PersonService) GetFamily(ctx context.Context) (*model.GetFamilyResponse, error) {
	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeRead})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	var readin []*models.Person
	err = sess.Query(context.Background(), "MATCH (p:Person) RETURN p", nil, &readin)
	if err != nil {
		panic(err)
	}

	err = sess.LoadAll(context.Background(), &readin)
	if err != nil {
		panic(err)
	}

	response := &model.GetFamilyResponse{
		Persons:       readin,
		Relationships: make([]*model.Relationship, 0),
	}

	for _, p := range readin {
		for _, c := range p.Children {
			newRelationship := model.Relationship{
				Parent:     c.Parent.UUID,
				Child:      c.Child.UUID,
				ParentType: string(c.ParentType),
			}
			response.Relationships = append(response.Relationships, &newRelationship)
		}
	}

	return response, nil
}

func (ps *PersonService) CreateChild(ctx context.Context, input *model.CreateChildInput) (string, error) {
	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	newPersonId, err := ps.createNewPerson(sess, ctx, input.ChildName)
	if err != nil {
		panic(err)
	}

	var readin models.Person
	err = sess.Load(context.Background(), &readin, newPersonId)
	if err != nil {
		panic(err)
	}

	if input.ParentType == "FATHER" {
		updateParentsInput := &model.UpdateParentsInput{
			Child:  newPersonId,
			Father: &input.ParentID,
		}
		ps.UpdateParents(ctx, *updateParentsInput)
	} else {
		updateParentsInput := &model.UpdateParentsInput{
			Child:  newPersonId,
			Mother: &input.ParentID,
		}
		ps.UpdateParents(ctx, *updateParentsInput)

	}

	return newPersonId, err
}

func (ps *PersonService) CreateParent(ctx context.Context, input *model.CreateParentInput) (string, error) {
	sess, err := ps.Ogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	newPersonId, err := ps.createNewPerson(sess, ctx, input.ParentName)
	if err != nil {
		panic(err)
	}

	var readin models.Person
	err = sess.Load(context.Background(), &readin, newPersonId)
	if err != nil {
		panic(err)
	}

	if input.ParentType == "FATHER" {
		updateParentsInput := &model.UpdateParentsInput{
			Child:  input.ChildID,
			Father: &newPersonId,
		}
		ps.UpdateParents(ctx, *updateParentsInput)
	} else {
		updateParentsInput := &model.UpdateParentsInput{
			Child:  input.ChildID,
			Mother: &newPersonId,
		}
		ps.UpdateParents(ctx, *updateParentsInput)

	}

	return newPersonId, err
}

func (ps *PersonService) createNewPerson(sess gogm.SessionV2, ctx context.Context, name string) (string, error) {
	newPerson := &models.Person{
		Name: name,
	}

	err := sess.Save(context.Background(), newPerson)
	if err != nil {
		panic(err)
	}

	return newPerson.UUID, err
}
