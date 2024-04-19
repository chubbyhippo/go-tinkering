package set

type Set map[interface{}]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(value interface{}) {
	s[value] = struct{}{}
}

func (s Set) Delete(value interface{}) {
	delete(s, value)
}

func (s Set) Contains(value interface{}) bool {
	_, ok := s[value]
	return ok
}
