package wkhtmltopdf

import (
	"reflect"
	"testing"
)

func TestNewDocument(t *testing.T) {

	doc := NewDocument()
	if !reflect.DeepEqual(doc, &Document{pages: []string{}}) {
		t.Errorf("NewDocument not produced correctly. Expected: %v, Got: %v",
			&Document{pages: []string{}}, doc)
	}
}

func TestAddPages(t *testing.T) {

	testcases := []struct {
		Pages []string
		All   []string
	}{
		{[]string{}, []string{}},
		{[]string{"test.html"}, []string{"test.html"}},
		{[]string{}, []string{"test.html"}},
		{[]string{"test2.html"}, []string{"test.html", "test2.html"}},
		{[]string{"test3.html", "test4.html"}, []string{"test.html", "test2.html", "test3.html", "test4.html"}},
	}

	doc := NewDocument()
	for _, tc := range testcases {
		doc.AddPages(tc.Pages...)
		if !reflect.DeepEqual(doc.pages, tc.All) {
			t.Errorf("Wrong page list. Expected: %v, Got: %v", tc.All, doc.pages)
		}
	}
}