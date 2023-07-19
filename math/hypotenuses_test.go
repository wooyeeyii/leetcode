package math

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Cal(t *testing.T) {
	cs1 := cal(750.0, 550.0, 250.0, 450.0, 100.0)
	fmt.Println("1号图")
	fmt.Printf("最长一根: %.2f， 每根减少 %.2f\n", cs1[0], cs1[1])

	cs2 := cal(750.0, 750.0, 350.0, 450.0, 100.0)
	fmt.Println("2号图")
	fmt.Printf("最长一根: %.2f， 每根减少 %.2f\n", cs2[0], cs2[1])
}

type con struct {
	A    string
	F    func() int
	Next *con
}

func Test_con(t *testing.T) {
	c1 := &con{
		A: "A",
		F: func() int {
			return 1
		},
	}
	c2 := &con{
		A: "B",
		F: func() int {
			return 2
		},
		Next: c1,
	}

	c1r := reflect.ValueOf(c1)
	y := c1r.Elem().FieldByName("A")
	fmt.Println(y)
	y.Set(reflect.ValueOf("A1"))
	fmt.Println(c1r.Elem().FieldByName("A"))
	fmt.Println(c1r.Elem().FieldByName("F"))
	fmt.Println(c1r.Elem().FieldByName("Next"))
	c2r := reflect.ValueOf(c2)
	fmt.Println(c2r.Elem().FieldByName("A"))
	fmt.Println(c2r.Elem().FieldByName("F"))
	fmt.Println(c2r.Elem().FieldByName("Next"))
	c2rNextr := reflect.ValueOf(c2r.Elem().FieldByName("Next").Interface().(*con))
	fmt.Println(c2rNextr.Elem().FieldByName("A"))
}

func Test_r(t *testing.T) {
	type T struct {
		N int
	}
	var n = T{42}
	// N at start
	fmt.Println(n.N)
	// pointer to struct - addressable
	ps := reflect.ValueOf(&n)
	// struct
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		// exported field
		f := s.FieldByName("N")
		if f.IsValid() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if f.CanSet() {
				// change value of N
				if f.Kind() == reflect.Int {
					x := int64(7)
					if !f.OverflowInt(x) {
						f.SetInt(x)
					}
				}
			}
		}
	}
	// N at end
	fmt.Println(n.N)
}
