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
