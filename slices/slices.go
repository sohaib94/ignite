package slices

func UnorderedEquals[T comparable](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	s2Map := make(map[T]int)
	for i := 0; i < len(s2); i +=1 {
		if val, ok := s2Map[s2[i]]; ok {
			s2Map[s2[i]] = val + 1
		} else {
			s2Map[s2[i]] = 1
		}
	}

	for i := range s1 {
		if val, ok := s2Map[s1[i]]; (!ok) || (ok && val < 1) {
			return false
		} else {
			s2Map[s1[i]] = val - 1
		}
	}

	return true
}
