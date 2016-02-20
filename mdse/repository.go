package mdse

import ()

type Repositor interface {
	Add(mdse *Mdse)
	GetAll() []*Mdse
}

type InMemRepository struct {
	list []*Mdse
}

func (repo *InMemRepository) Add(mdse *Mdse) {
	repo.list = append(repo.list, mdse)
}

// return copy of memory.
func (repo *InMemRepository) GetAll() []*Mdse {
	result := make([]*Mdse, 0, len(repo.list))
	for _, mdse := range repo.list {
		result = append(result, mdse)
	}
	return result
}

var repository = &InMemRepository{
	list: make([]*Mdse, 0, 32),
}

func GetRepository(repoType string) *InMemRepository {
	if repository == nil {
		repository = &InMemRepository{}
	}
	return repository
}
