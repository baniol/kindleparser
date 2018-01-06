package kindleparser

import (
	"testing"
)

func Test_splitIntoRecords(t *testing.T) {

	records := ParseClippngs("sample_clippings.txt")

	var expected string

	expected = "GPRS Tutoial"
	if records[0].Title != expected {
		t.Errorf("Expected %s, but got %s", expected, records[0].Title)
	}

	expected = "K. K. Panigrahi"
	if records[0].Author != expected {
		t.Errorf("Expected %s, but got %s", expected, records[0].Author)
	}

	expected = "A variable declaration provides assurance to the compiler that there is one variable existing with the given type and name"
	if records[0].Content != expected {
		t.Errorf("Expected %s, but got %s", expected, records[0].Content)
	}
}
