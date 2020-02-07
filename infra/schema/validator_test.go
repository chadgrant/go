package schema

import (
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		schema string
		json   string
		errors []string
	}{{
		"http://schemas.vevo.com/models/content/3.0/artist.json",
		`{ "name": "Metallica", "urlSafeName" : "metallica", "created" : "2020-02-04T16:50:15.1093474Z" }`,
		[]string{},
	}, {
		"product.json",
		`{ "id":"741983bf-c6da-4310-b1f3-843b7b7181ec", "name":"test", "category":"test", "price": 1.99 }`,
		[]string{},
	}, {
		"http://schemas.sentex.io/product.json",
		`{ "id":"741983bf-c6da-4310-b1f3-843b7b7181ec", "name":"test", "category":"test", "price": 1.99 }`,
		[]string{},
	}, {
		"product.json",
		`{ "id":"10-b1f3-843b7b7181ec", "name":"test", "category":"test", "price": 1.99 }`,
		[]string{`id: Does not match pattern '^[A-Za-z0-9\-]{36}$'. value was: 10-b1f3-843b7b7181ec`},
	}, {
		"product.json",
		`{ }`,
		[]string{"id is required", "category is required", "name is required", "price is required"},
	}, {
		"product.json",
		`{ "id":"741983bf-c6da-4310-b1f3-843b7b7181ec", "name":"test", "category":"test", "price": "notanumber" }`,
		[]string{"price: Invalid type. Expected: number, given: string"},
	}}

	root := "./test-schemas"
	reg := NewRegistry()
	err := reg.AddDirectory(root)
	if err != nil {
		t.Fatalf("loading schemas from disk : %s %v", root, err)
	}

	val, err := NewValidator(reg)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {

		var file string
		if strings.HasPrefix(test.schema, "http") {
			file = test.schema
		} else {
			abs, err := filepath.Abs(path.Join(root, test.schema))
			if err != nil {
				t.Errorf("creating abs path %v", err)
			}
			file = "file://" + abs
		}

		err, validationErors := val.Validate(file, []byte(test.json))
		if err != nil {
			t.Errorf("validating basic json : %v", err)
		}

		errs := make(map[string]struct{})
		for _, e := range validationErors {
			errs[e.Error()] = struct{}{}
		}

		if len(validationErors) != len(test.errors) {
			t.Errorf("expected %d errors, got %d", len(test.errors), len(validationErors))
			for _, e := range validationErors {
				t.Error(e)
			}
		}

		for _, v := range test.errors {
			if _, ok := errs[v]; !ok {
				t.Errorf("expected but did not find %s", v)
				for _, e := range validationErors {
					t.Error(e)
				}
			}
		}
	}
}
