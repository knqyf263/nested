package nested_test

import (
	"fmt"
	"strings"

	"github.com/knqyf263/nested"
)

func Example() {
	n := nested.Nested{}

	n.Set([]string{"a", "b"}, 1)
	n.SetByString("a.c.d", ".", "test")
	n.SetByString("/e/f", "/", true)

	var result interface{}
	result, _ = n.Get([]string{"a", "c", "d"})
	fmt.Println(result)

	result, _ = n.GetByString("e/f", "/")
	fmt.Println(result)

	var b int
	b, _ = n.GetInt([]string{"a", "b"})
	fmt.Println(b)

	// Output:
	// test
	// true
	// 1
}

func ExampleNested_Set() {
	n := nested.Nested{}
	n.Set([]string{"a", "b", "c"}, 1)
	fmt.Printf("%v", n)
	// Output: map[a:map[b:map[c:1]]]
}

func ExampleNested_SetByString() {
	n := nested.Nested{}
	n.SetByString("a.b.c", ".", "test")
	fmt.Printf("%v", n)
	// Output: map[a:map[b:map[c:test]]]
}

func ExampleNested_Get() {
	n := nested.Nested{"a": map[string]interface{}{"b": 1}}
	result, _ := n.Get([]string{"a", "b"})
	fmt.Printf("%v", result)
	// Output: 1
}

func ExampleNested_GetByString() {
	n := nested.Nested{"a": map[string]interface{}{"b": map[string]interface{}{"c": 2}}}
	result, _ := n.GetByString("a.b.c", ".")
	fmt.Printf("%v", result)
	// Output: 2
}

func ExampleNested_Delete() {
	n := nested.Nested{"a": map[string]interface{}{"b": 1, "c": 2}}
	n.Delete([]string{"a", "b"})
	fmt.Printf("%v", n)
	// Output: map[a:map[c:2]]
}

func ExampleNested_DeleteByString() {
	n := nested.Nested{"a": map[string]interface{}{"b": 1, "c": 2}}
	n.DeleteByString("a/b", "/")
	fmt.Printf("%v", n)
	// Output: map[a:map[c:2]]
}

func ExampleNested_Walk() {
	n := nested.Nested{}
	n.Set([]string{"a", "b", "c"}, 1)
	n.Set([]string{"d", "e"}, 2)

	walkFn := func(keys []string, value interface{}) error {
		fmt.Println(keys, value)
		return nil
	}
	n.Walk(walkFn)

	// [a] map[b:map[c:1]]
	// [a b] map[c:1]
	// [a b c] 1
	// [d] map[e:2]
	// [d e] 2
}

func main() {
	n := nested.Nested{}
	n.Set([]string{"a", "b", "c"}, 1)
	n.Set([]string{"d", "e"}, 2)
	n.SetByString("f.g.h", ".", false)

	walkFn := func(keys []string, value interface{}) error {
		key := strings.Join(keys, ".")
		// Skip all keys under "f"
		if key == "f" {
			return nested.SkipKey
		}
		fmt.Println(key, value)
		return nil
	}
	n.Walk(walkFn)

	// Output:
	// a map[b:map[c:1]]
	// a.b map[c:1]
	// a.b.c 1
	// d map[e:2]
	// d.e 2
	// f map[g:map[h:1]]
}
