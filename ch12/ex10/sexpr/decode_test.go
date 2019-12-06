package sexpr

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type wrapper struct {
		Hoge1 interface{}
		Hoge2 interface{}
		Hoge3 interface{}
		Hoge4 interface{}
		Hoge5 interface{}
	}
	type TestStruct struct {
		StringField  string
		HogeaceField interface{}
	}
	INTERFACES["sexpr.TestStruct"] = reflect.TypeOf(TestStruct{})
	for _, test := range []struct {
		input string
		want  wrapper
	}{
		{`nil`, wrapper{}},
		{`()`, wrapper{}},
		{`(
			 (Hoge1 ("[]int" (1 2 3)))
			 (Hoge2 ("[]string" ("a" "b" "c")))
			 (Hoge3 nil)
       (Hoge4 ("[]complex128" (#C(1.2 3.4) #C(5.6 7.8) #C(9.10 11.12))))
			 (Hoge5 nil)
			)`,
			wrapper{
				Hoge1: []int{1, 2, 3},
				Hoge2: []string{"a", "b", "c"},
				Hoge3: nil,
				Hoge4: []complex128{complex(1.2, 3.4), complex(5.6, 7.8), complex(9.10, 11.12)},
				Hoge5: nil,
			},
		},
		{`(
			 (Hoge1 ("[]int" (1 2 3)))
			 (Hoge2 ("[]string" ("a" "b" "c")))
			 (Hoge3 ("[]float64" (1.23 4.56 7.89)))
       (Hoge4 ("[]complex128" (#C(1.2 3.4) #C(5.6 7.8) #C(9.10 11.12))))
			 (Hoge5 ("sexpr.TestStruct" ((StringField "abc")
                                     (HogeaceField ("sexpr.TestStruct" ((StringField "def")
                                                                          (HogeaceField nil))))))
		  )
		)`,
			wrapper{
				Hoge1: []int{1, 2, 3},
				Hoge2: []string{"a", "b", "c"},
				Hoge3: []float64{1.23, 4.56, 7.89},
				Hoge4: []complex128{complex(1.2, 3.4), complex(5.6, 7.8), complex(9.10, 11.12)},
				Hoge5: TestStruct{StringField: "abc", HogeaceField: TestStruct{StringField: "def"}},
			}},
	} {
		got := wrapper{}
		err := Unmarshal([]byte(test.input), &got)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("\ngot: \n%v\nwant: \n%v\n", got, test.want)
		}
	}
}
