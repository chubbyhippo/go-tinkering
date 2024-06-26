package set

import "testing"

func TestSet(t *testing.T) {
	t.Run("should add item to the set", func(t *testing.T) {
		s := NewSet()
		s.Add("Go")
		got := s.Contains("Go")

		if got != true {
			t.Errorf("got %t is not true", got)
		}
	})
	t.Run("should add item to the set", func(t *testing.T) {
		s := NewSet()
		s.Add("Go")
		s.Delete("Go")
		got := s.Contains("Go")

		if got != false {
			t.Errorf("got %t is not false", got)
		}
	})

}
