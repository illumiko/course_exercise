package main

type League []Player

func (l League) Find(name string) *Player {
	for i, v := range l {
		if v.Name == name {
			return &l[i]
		}

	}
	return nil
}
