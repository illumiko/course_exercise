package main

type InMemoryStore struct {
	scores map[string]int
}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMemoryStore) RecordWin(name string) {
	i.scores[name]++

}

func (i *InMemoryStore) GetLeague() League {
	league := League{}
	for name, wins := range i.scores {
		league = append(league, Player{name, wins})
	}
	return league
}

func newInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		scores: map[string]int{"John": 2},
	}
}
