package item

import ()

type Repositor interface {
	Add(item *Item)
	GetAll() []*Item
}

type InMemRepository struct {
	list []*Item
}

func (repo *InMemRepository) Add(item *Item) {
	repo.list = append(repo.list, item)
}

// return copy of memory.
func (repo *InMemRepository) GetAll() []*Item {
	result := make([]*Item, 0, len(repo.list))
	for _, item := range repo.list {
		result = append(result, item)
	}
	return result
}

var repository = &InMemRepository{
	list: make([]*Item, 0, 32),
}

func GetRepository(repoType string) *InMemRepository {
	if repository == nil {
		repository = &InMemRepository{}
	}
	return repository
}
