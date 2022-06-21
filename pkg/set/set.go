package set

type Set map[string]bool

func New(ss... string) Set {
	res := Set(make(map[string]bool, len(ss)))
	for _, s := range ss {
		res.Insert(s)
	}
	return res
}

func (s Set) Contains(n string) bool {
	return s[n]
}

func (s Set) Insert(n string) {
	s[n] = true
}

func (s Set) Remove(n string) {
	delete(s, n)
}

func (s Set) Append(o Set) {
	for e, _ := range o {
		s.Insert(e)
	}
}

func Intersection(s1, s2 Set) Set {
	res := New()
	for e, _ := range s1 {
		if s2.Contains(e) {
			res.Insert(e)
		}
	}
	return res
}
