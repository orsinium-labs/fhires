package fhires

import (
	"encoding/json/v2"
	"encoding/xml"
	"testing"

	"github.com/orsinium-labs/fhires/fhires"
)

func eq[V comparable](t *testing.T, a, b V) {
	t.Helper()
	if a != b {
		t.Fatalf("got %#v, expected %#v", a, b)
	}
}

func noErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestJSONEncode(t *testing.T) {
	resIn := fhires.Encounter{
		ID:     new("123"),
		Status: fhires.EncounterStatusCompleted,
	}
	raw, err := json.Marshal(resIn)
	noErr(t, err)
	eq(t, string(raw), `{"resourceType":"Encounter","id":"123","status":"completed"}`)
}

func TestJSONRoundTrip(t *testing.T) {
	resIn := fhires.Encounter{
		ID:     new("123"),
		Status: fhires.EncounterStatusCompleted,
	}
	raw, err := json.Marshal(resIn)
	noErr(t, err)
	var resOut fhires.Encounter
	err = json.Unmarshal(raw, &resOut)
	noErr(t, err)
	eq(t, *resIn.ID, *resOut.ID)
	eq(t, resIn.Status, resOut.Status)
}

func TestXMLEncode(t *testing.T) {
	resIn := fhires.Encounter{
		ID:     new("123"),
		Status: fhires.EncounterStatusCompleted,
	}
	raw, err := xml.Marshal(resIn)
	noErr(t, err)
	eq(t, string(raw), `<Encounter ><id>123</id><status>completed</status></Encounter>`)
}

func TestXMLDecode(t *testing.T) {
	var res fhires.Encounter
	input := `<Encounter><id value="123"/><status value="completed"/></Encounter>`
	err := xml.Unmarshal([]byte(input), &res)
	noErr(t, err)
	eq(t, *res.ID, "123")
	eq(t, res.Status, fhires.EncounterStatusCompleted)
}

func TestXMLRoundTrip(t *testing.T) {
	resIn := fhires.Encounter{
		ID:     new("123"),
		Status: fhires.EncounterStatusCompleted,
	}
	raw, err := xml.Marshal(resIn)
	noErr(t, err)
	var resOut fhires.Encounter
	err = xml.Unmarshal(raw, &resOut)
	noErr(t, err)
	eq(t, *resIn.ID, *resOut.ID)
	eq(t, resIn.Status, resOut.Status)
}

func TestAnyResourceJSONRoundTrip(t *testing.T) {
	bunIn := fhires.Bundle{
		Entry: []fhires.BundleEntry{{
			Resource: fhires.AnyResource{
				Value: fhires.Encounter{
					ID:     new("123"),
					Status: fhires.EncounterStatusCompleted,
				},
			},
		}},
	}
	raw, err := json.Marshal(bunIn)
	noErr(t, err)

	var bunOut fhires.Bundle
	err = json.Unmarshal(raw, &bunOut)
	noErr(t, err)
	resIn := bunIn.Entry[0].Resource.Value.(fhires.Encounter)
	resOut := bunOut.Entry[0].Resource.Value.(fhires.Encounter)
	eq(t, *resIn.ID, *resOut.ID)
	eq(t, resIn.Status, resOut.Status)
}

func TestAnyResourceXMLRoundTrip(t *testing.T) {
	bunIn := fhires.Bundle{
		Entry: []fhires.BundleEntry{{
			Resource: fhires.AnyResource{
				Value: fhires.Encounter{
					ID:     new("123"),
					Status: fhires.EncounterStatusCompleted,
				},
			},
		}},
	}
	raw, err := xml.Marshal(bunIn)
	noErr(t, err)
	println(string(raw))

	var bunOut fhires.Bundle
	err = xml.Unmarshal(raw, &bunOut)
	noErr(t, err)
	resIn := bunIn.Entry[0].Resource.Value.(fhires.Encounter)
	resOut := bunOut.Entry[0].Resource.Value.(fhires.Encounter)
	// eq(t, *resIn.ID, *resOut.ID)
	eq(t, resIn.Status, resOut.Status)
}
