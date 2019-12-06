package sexpr

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "",
		Year:     0,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "",
			"Pres. Merkin Muffley":       "",
			"Gen. Buck Turgidson":        "",
			"Brig. Gen. Jack D. Ripper":  "",
			`Maj. T.J. "King" Kong`:      "",
		},
		Oscars: []string{
			"",
			"",
			"",
			"",
		},
	}

	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = \n%s\n", data)
	fmt.Printf("%+v\n", data)

	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}
