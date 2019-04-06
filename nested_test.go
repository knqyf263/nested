package nested

import (
	"reflect"
	"strings"
	"testing"
)

var setTests = []struct {
	name  string
	key   []string
	value interface{}
	out   Nested
}{
	{"first", []string{"a"}, 1, Nested{"a": 1}},
	{"first2", []string{"b"}, "test", Nested{"a": 1, "b": "test"}},
	{"second", []string{"a", "c"}, 1, Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}},
	{"second2", []string{"a", "d"}, true, Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}},
	{"third", []string{"e", "f", "g"}, []int{1, 2}, Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}},
	{"third2", []string{"e", "f", "h"}, map[string]int{"i": 100}, Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}, "h": map[string]int{"i": 100}}}}},
}

func TestSet(t *testing.T) {
	nested := Nested{}
	for _, tt := range setTests {
		t.Run(tt.name, func(t *testing.T) {
			nested.Set(tt.key, tt.value)
			if !reflect.DeepEqual(nested, tt.out) {
				t.Errorf("got %v, want %v", nested, tt.out)
			}
		})
	}
}

var setByStringTests = []struct {
	name  string
	key   string
	sep   string
	value interface{}
	out   Nested
}{
	{"first", "a", ".", 1, Nested{"a": 1}},
	{"first2", "b", ".", "test", Nested{"a": 1, "b": "test"}},
	{"second", "a.c", ".", 1, Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}},
	{"second2", "a.d", ".", true, Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}},
	{"second3", "a/d", "/", false, Nested{"a": map[string]interface{}{"c": 1, "d": false}, "b": "test"}},
	{"second4", "/a/d", "/", 2, Nested{"a": map[string]interface{}{"c": 1, "d": 2}, "b": "test"}},
	{"third", "e/f/g", "/", []int{1, 2}, Nested{"a": map[string]interface{}{"c": 1, "d": 2}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}},
	{"third2", "e.f.h", ".", map[string]int{"i": 100}, Nested{"a": map[string]interface{}{"c": 1, "d": 2}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}, "h": map[string]int{"i": 100}}}}},
}

func TestSetByString(t *testing.T) {
	nested := Nested{}
	for _, tt := range setByStringTests {
		t.Run(tt.name, func(t *testing.T) {
			nested.SetByString(tt.key, tt.sep, tt.value)
			if !reflect.DeepEqual(nested, tt.out) {
				t.Errorf("got %v, want %v", nested, tt.out)
			}
		})
	}
}

var getIntTests = []struct {
	name string
	in   Nested
	key  []string
	out  int
	err  error
}{
	{"first", Nested{"a": 1}, []string{"a"}, 1, nil},
	{"second", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, []string{"a", "c"}, 1, nil},
}

func TestGetInt(t *testing.T) {
	for _, tt := range getIntTests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.in.GetInt(tt.key)
			if err != tt.err {
				t.Errorf("err: got %v, want %v", err, tt.err)
			}
			if actual != tt.out {
				t.Errorf("got %v, want %v", actual, tt.out)
			}
		})
	}
}

var getStringTests = []struct {
	name string
	in   Nested
	key  []string
	out  string
	err  error
}{
	{"first", Nested{"a": "test"}, []string{"a"}, "test", nil},
	{"second", Nested{"a": map[string]interface{}{"c": "test2"}, "b": "test"}, []string{"a", "c"}, "test2", nil},
}

func TestGetString(t *testing.T) {
	for _, tt := range getStringTests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.in.GetString(tt.key)
			if err != tt.err {
				t.Errorf("err: got %v, want %v", err, tt.err)
			}
			if actual != tt.out {
				t.Errorf("got %v, want %v", actual, tt.out)
			}
		})
	}
}

var getBoolTests = []struct {
	name string
	in   Nested
	key  []string
	out  bool
	err  error
}{
	{"first", Nested{"a": true}, []string{"a"}, true, nil},
	{"second", Nested{"a": map[string]interface{}{"c": false}, "b": "test"}, []string{"a", "c"}, false, nil},
}

func TestGetBool(t *testing.T) {
	for _, tt := range getBoolTests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.in.GetBool(tt.key)
			if err != tt.err {
				t.Errorf("err: got %v, want %v", err, tt.err)
			}
			if actual != tt.out {
				t.Errorf("got %v, want %v", actual, tt.out)
			}
		})
	}
}

var getTests = []struct {
	name  string
	in    Nested
	key   []string
	value interface{}
	err   error
}{
	{"first", Nested{"a": 1}, []string{"a"}, 1, nil},
	{"first2", Nested{"a": 1, "b": "test"}, []string{"b"}, "test", nil},
	{"second", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, []string{"a", "c"}, 1, nil},
	{"second2", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}, []string{"a", "d"}, true, nil},
	{"third", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, []string{"e", "f", "g"}, []int{1, 2}, nil},
	{"third2", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}, "h": map[string]int{"i": 100}}}}, []string{"e", "f", "h"}, map[string]int{"i": 100}, nil},
	{"error", Nested{"a": 1}, []string{"z"}, nil, ErrNoSuchKey},
	{"error2", Nested{"a": 1, "b": "test"}, []string{"a", "y"}, nil, ErrNoSuchKey},
	{"error3", Nested{"a": 1, "b": "test"}, []string{"x"}, nil, ErrNoSuchKey},
}

func TestGet(t *testing.T) {
	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.in.Get(tt.key)
			if err != tt.err {
				t.Errorf("err: got %v, want %v", err, tt.err)
			}
			if !reflect.DeepEqual(actual, tt.value) {
				t.Errorf("got %v, want %v", actual, tt.value)
			}
		})
	}
}

var getByStringTests = []struct {
	name  string
	in    Nested
	key   string
	sep   string
	value interface{}
	err   error
}{
	{"first", Nested{"a": 1}, "a", ".", 1, nil},
	{"first2", Nested{"a": 1, "b": "test"}, "b", ".", "test", nil},
	{"second_dot", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, "a.c", ".", 1, nil},
	{"second_slash", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, "a/c", "/", 1, nil},
	{"second_slash_prefix", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, "/a/c", "/", 1, nil},
	{"second2", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}, "a.d", ".", true, nil},
	{"third", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, "e.f.g", ".", []int{1, 2}, nil},
	{"third2", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}, "h": map[string]int{"i": 100}}}}, "e/f/h", "/", map[string]int{"i": 100}, nil},
	{"error", Nested{"a": 1}, "z", ",", nil, ErrNoSuchKey},
	{"error2", Nested{"a": 1, "b": "test"}, "a.y", ".", nil, ErrNoSuchKey},
	{"error3", Nested{"a": 1, "b": "test"}, "a.b", "/", nil, ErrNoSuchKey},
	{"error4", Nested{"a": 1, "b": "test"}, "x", ".", nil, ErrNoSuchKey},
}

func TestGetByString(t *testing.T) {
	for _, tt := range getByStringTests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.in.GetByString(tt.key, tt.sep)
			if err != tt.err {
				t.Errorf("err: got %v, want %v", err, tt.err)
			}
			if !reflect.DeepEqual(actual, tt.value) {
				t.Errorf("got %v, want %v", actual, tt.value)
			}
		})
	}
}

var deleteTests = []struct {
	name  string
	in    Nested
	key   []string
	value interface{}
	err   error
}{
	{"first", Nested{"a": 1}, []string{"a"}, Nested{}, nil},
	{"first2", Nested{"a": 1, "b": "test"}, []string{"b"}, Nested{"a": 1}, nil},
	{"second", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, []string{"a", "c"}, Nested{"a": map[string]interface{}{}, "b": "test"}, nil},
	{"second2", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, []string{"a"}, Nested{"b": "test"}, nil},
	{"second3", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}, []string{"a", "d"}, Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, nil},
	{"second4", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}, []string{"b"}, Nested{"a": map[string]interface{}{"c": 1, "d": true}}, nil},
	{"third", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, []string{"e", "f", "g"}, Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{}}}, nil},
	{"third2", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, []string{"e", "f"}, Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{}}, nil},
	{"third3", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, []string{"e"}, Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}, nil},
	{"error", Nested{"a": 1}, []string{"z"}, Nested{"a": 1}, ErrNoSuchKey},
	{"error2", Nested{"a": 1, "b": "test"}, []string{"a", "y"}, Nested{"a": 1, "b": "test"}, ErrNoSuchKey},
	{"error3", Nested{"a": 1, "b": "test"}, []string{"x"}, Nested{"a": 1, "b": "test"}, ErrNoSuchKey},
}

func TestDelete(t *testing.T) {
	for _, tt := range deleteTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.in
			err := actual.Delete(tt.key)
			if err != tt.err {
				t.Errorf("err: got %v, want %v", err, tt.err)
			}
			if !reflect.DeepEqual(actual, tt.value) {
				t.Errorf("got %v, want %v", actual, tt.value)
			}
		})
	}
}

var deleteByStringTests = []struct {
	name  string
	in    Nested
	key   string
	sep   string
	value interface{}
	err   error
}{
	{"first", Nested{"a": 1}, "a", ".", Nested{}, nil},
	{"first2", Nested{"a": 1, "b": "test"}, "b", "/", Nested{"a": 1}, nil},
	{"second", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, "a.c", ".", Nested{"a": map[string]interface{}{}, "b": "test"}, nil},
	{"second2", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, "a", ".", Nested{"b": "test"}, nil},
	{"second3", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}, "a.d", ".", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, nil},
	{"third", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, "e.f.g", ".", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{}}}, nil},
	{"third2", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, "e/f", "/", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{}}, nil},
	{"error", Nested{"a": 1}, "z", ".", Nested{"a": 1}, ErrNoSuchKey},
	{"error2", Nested{"a": 1, "b": "test"}, "a.b", ",", Nested{"a": 1, "b": "test"}, ErrNoSuchKey},
	{"error3", Nested{"a": 1, "b": "test"}, "x", "/", Nested{"a": 1, "b": "test"}, ErrNoSuchKey},
}

func TestDeleteByString(t *testing.T) {
	for _, tt := range deleteByStringTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.in
			err := actual.DeleteByString(tt.key, tt.sep)
			if err != tt.err {
				t.Errorf("err: got %v, want %v", err, tt.err)
			}
			if !reflect.DeepEqual(actual, tt.value) {
				t.Errorf("got %v, want %v", actual, tt.value)
			}
		})
	}
}

var walkTests = []struct {
	name string
	in   Nested
	out  map[string]interface{}
}{
	{"first", Nested{"a": 1}, map[string]interface{}{"a": 1}},
	{"first2", Nested{"a": 1, "b": "test"}, map[string]interface{}{"a": 1, "b": "test"}},
	{"second", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, map[string]interface{}{"a": map[string]interface{}{"c": 1}, "a.c": 1, "b": "test"}},
	{"second2", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}, map[string]interface{}{"a": map[string]interface{}{"c": 1, "d": true}, "a.c": 1, "a.d": true, "b": "test"}},
	{"third", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, map[string]interface{}{"a": map[string]interface{}{"c": 1, "d": true}, "a.c": 1, "a.d": true, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}, "e.f": map[string]interface{}{"g": []int{1, 2}}, "e.f.g": []int{1, 2}}},
}

func TestWalk(t *testing.T) {
	for _, tt := range walkTests {
		walkResults := map[string]interface{}{}

		walkFn := func(keys []string, value interface{}) error {
			key := strings.Join(keys, ".")
			walkResults[key] = value
			return nil
		}

		t.Run(tt.name, func(t *testing.T) {
			tt.in.Walk(walkFn)
			if !reflect.DeepEqual(walkResults, tt.out) {
				t.Errorf("got %v, want %v", walkResults, tt.out)
			}
		})
	}
}

var walkSkipTests = []struct {
	name string
	in   Nested
	out  map[string]interface{}
}{
	{"first", Nested{"a": 1}, map[string]interface{}{"a": 1}},
	{"first2", Nested{"a": 1, "b": "test"}, map[string]interface{}{"a": 1, "b": "test"}},
	{"second", Nested{"a": map[string]interface{}{"c": 1}, "b": "test"}, map[string]interface{}{"a": map[string]interface{}{"c": 1}, "b": "test"}},
	{"second2", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}, map[string]interface{}{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test"}},
	{"third", Nested{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}}, map[string]interface{}{"a": map[string]interface{}{"c": 1, "d": true}, "b": "test", "e": map[string]interface{}{"f": map[string]interface{}{"g": []int{1, 2}}}, "e.f": map[string]interface{}{"g": []int{1, 2}}}},
}

func TestSkipWalk(t *testing.T) {
	for _, tt := range walkSkipTests {
		walkResults := map[string]interface{}{}

		walkFn := func(keys []string, value interface{}) error {
			key := strings.Join(keys, ".")
			walkResults[key] = value
			if key == "a" || key == "e.f" {
				return SkipKey

			}
			return nil
		}

		t.Run(tt.name, func(t *testing.T) {
			tt.in.Walk(walkFn)
			if !reflect.DeepEqual(walkResults, tt.out) {
				t.Errorf("got %v, want %v", walkResults, tt.out)
			}
		})
	}
}
