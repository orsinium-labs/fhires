package fhires

import (
	"bytes"
	jsontext "encoding/json/jsontext"
	"encoding/json/v2"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"reflect"
)

type IResource interface {
	json.Marshaler
	GetBaseResource() BaseResource
	ResourceType() string
}

// Union type of all resources.
//
// Used as a type for fields (like [BundleEntryResponse.Outcome])
// that may contain any resource type.
type AnyResource struct {
	Value IResource
}

var (
	_ json.MarshalerTo = AnyResource{}
	_ json.Unmarshaler = &AnyResource{}
	_ xml.Marshaler    = AnyResource{}
	_ xml.Unmarshaler  = &AnyResource{}
)

// MarshalXML implements [xml.Marshaler].
func (r AnyResource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Value == nil {
		return nil
	}
	outName := start.Name
	resName := xml.Name{Local: r.Value.ResourceType()}
	err := e.EncodeToken(xml.StartElement{Name: outName})
	if err != nil {
		return err
	}
	inner := start.Copy()
	inner.Name = resName
	err = e.EncodeElement(r.Value, inner)
	if err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: outName})
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *AnyResource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	rawToken, err := d.Token()
	if err != nil {
		return err
	}
	token, isStart := rawToken.(xml.StartElement)
	if !isStart {
		return fmt.Errorf("expected XML opening tag, got %T", token)
	}
	if r.Value != nil {
		return errors.New("expected a single resource element")
	}
	r.Value = GetResource(token.Name.Local)
	if r.Value == nil {
		return fmt.Errorf("unsupported resourceType: %s", token.Name.Local)
	}
	ptr := reflect.New(reflect.TypeOf(r.Value))
	err = d.DecodeElement(ptr.Interface(), &token)
	if err != nil {
		return err
	}
	r.Value = ptr.Elem().Interface().(IResource)

	rawToken, err = d.Token()
	if err != nil {
		return err
	}
	_, isEnd := rawToken.(xml.EndElement)
	if !isEnd {
		return fmt.Errorf("expected XML closing tag, got %T", token)
	}
	return nil
}

// MarshalJSONTo implements [json.MarshalerTo].
func (r AnyResource) MarshalJSONTo(e *jsontext.Encoder) error {
	return json.MarshalEncode(e, r.Value)
}

// UnmarshalJSON implements [json.Unmarshaler].
func (r *AnyResource) UnmarshalJSON(b []byte) error {
	var res struct {
		ResourceType string `json:"resourceType"`
	}
	err := json.Unmarshal(b, &res, json.RejectUnknownMembers(false))
	if err != nil {
		return fmt.Errorf("read resourceType: %w", err)
	}
	if res.ResourceType == "" {
		return errors.New("not a valid resource")
	}
	val := GetResource(res.ResourceType)
	if val == nil {
		return fmt.Errorf("unsupported resourceType: %s", res.ResourceType)
	}
	err = json.Unmarshal(b, &val)
	if err != nil {
		return fmt.Errorf("read resource of the provided resourceType: %w", err)
	}
	r.Value = val
	return nil
}

func marshalJSON(t string, r any) ([]byte, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	b = append([]byte(`{"resourceType":"`+t+`",`), b[1:]...)
	return b, nil
}

func marshalXML(e *xml.Encoder, start xml.StartElement, t string, r any) error {
	var b bytes.Buffer
	start.Name.Local = t
	err := xml.NewEncoder(&b).EncodeElement(r, start)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	var pending *xml.StartElement
	var value []byte

	d := xml.NewDecoder(&b)
	for {
		token, err := d.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("get next token: %w", err)
		}

		// If there is inner text in the tag, store it
		// as `value` attribute of the tag instead.
		charData, isCharData := token.(xml.CharData)
		if isCharData {
			if pending != nil {
				value = append(value, charData...)
			}
			continue
		}

		// Flush the pending start tag.
		if pending != nil {
			if len(value) > 0 {
				pending.Attr = append(pending.Attr, xml.Attr{
					Name:  xml.Name{Local: "value"},
					Value: string(value),
				})
			}
			err = e.EncodeToken(*pending)
			pending = nil
			value = value[:0]
			if err != nil {
				return fmt.Errorf("encode start element: %w", err)
			}
		}

		startElement, isStart := token.(xml.StartElement)
		if isStart {
			pending = new(startElement.Copy())
			continue
		}

		err = e.EncodeToken(token)
		if err != nil {
			return fmt.Errorf("encode element: %w", err)
		}
	}
}

func unmarshalXML(d *xml.Decoder, start xml.StartElement, t string, r any) error {
	var b bytes.Buffer
	encoder := xml.NewEncoder(&b)
	if start.Name.Local != t {
		return fmt.Errorf("unexpected XML tag: expected %s, got %s", t, start.Name.Local)
	}
	err := encoder.EncodeToken(start)
	if err != nil {
		return fmt.Errorf("encode start token: %w", err)
	}
	for {
		token, err := d.Token()
		if err != nil && err != io.EOF {
			return fmt.Errorf("get next token: %w", err)
		}
		if token == nil {
			break
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			fhirValue := findValue(startElement.Attr)
			if fhirValue != "" {
				err := encoder.EncodeElement(fhirValue, xml.StartElement{Name: startElement.Name})
				if err != nil {
					return fmt.Errorf("encode value element: %w", err)
				}
				err = d.Skip()
				if err != nil {
					return fmt.Errorf(": %w", err)
				}
			} else {
				err := encoder.EncodeToken(startElement)
				if err != nil {
					return fmt.Errorf("encode non-value element: %w", err)
				}
			}
		default:
			err := encoder.EncodeToken(token)
			if err != nil {
				return fmt.Errorf("encode element: %w", err)
			}
		}
	}
	err = encoder.Flush()
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	return xml.NewDecoder(&b).Decode(r)
}

func findValue(attr []xml.Attr) string {
	for _, a := range attr {
		if a.Name.Local == "value" {
			return a.Value
		}
	}
	return ""
}

// Get resource instance for the given resource name.
//
// For example, returns [Encounter] for "Encounter".
// Returns [nil] if there is no match.
func GetResource(resourceType string) IResource {
	switch resourceType {
	case "Account":
		return Account{}
	case "ActivityDefinition":
		return ActivityDefinition{}
	case "ActorDefinition":
		return ActorDefinition{}
	case "AdministrableProductDefinition":
		return AdministrableProductDefinition{}
	case "AdverseEvent":
		return AdverseEvent{}
	case "AllergyIntolerance":
		return AllergyIntolerance{}
	case "Appointment":
		return Appointment{}
	case "AppointmentResponse":
		return AppointmentResponse{}
	case "ArtifactAssessment":
		return ArtifactAssessment{}
	case "AuditEvent":
		return AuditEvent{}
	case "Basic":
		return Basic{}
	case "Binary":
		return Binary{}
	case "BiologicallyDerivedProductDispense":
		return BiologicallyDerivedProductDispense{}
	case "BiologicallyDerivedProduct":
		return BiologicallyDerivedProduct{}
	case "BodyStructure":
		return BodyStructure{}
	case "Bundle":
		return Bundle{}
	case "CanonicalResource":
		return CanonicalResource{}
	case "CapabilityStatement":
		return CapabilityStatement{}
	case "CarePlan":
		return CarePlan{}
	case "CareTeam":
		return CareTeam{}
	case "ChargeItemDefinition":
		return ChargeItemDefinition{}
	case "ChargeItem":
		return ChargeItem{}
	case "Citation":
		return Citation{}
	case "Claim":
		return Claim{}
	case "ClaimResponse":
		return ClaimResponse{}
	case "ClinicalImpression":
		return ClinicalImpression{}
	case "ClinicalUseDefinition":
		return ClinicalUseDefinition{}
	case "CodeSystem":
		return CodeSystem{}
	case "Communication":
		return Communication{}
	case "CommunicationRequest":
		return CommunicationRequest{}
	case "CompartmentDefinition":
		return CompartmentDefinition{}
	case "Composition":
		return Composition{}
	case "ConceptMap":
		return ConceptMap{}
	case "ConditionDefinition":
		return ConditionDefinition{}
	case "Condition":
		return Condition{}
	case "Consent":
		return Consent{}
	case "Contract":
		return Contract{}
	case "CoverageEligibilityRequest":
		return CoverageEligibilityRequest{}
	case "CoverageEligibilityResponse":
		return CoverageEligibilityResponse{}
	case "Coverage":
		return Coverage{}
	case "DetectedIssue":
		return DetectedIssue{}
	case "DeviceAssociation":
		return DeviceAssociation{}
	case "DeviceDefinition":
		return DeviceDefinition{}
	case "DeviceDispense":
		return DeviceDispense{}
	case "Device":
		return Device{}
	case "DeviceMetric":
		return DeviceMetric{}
	case "DeviceRequest":
		return DeviceRequest{}
	case "DeviceUsage":
		return DeviceUsage{}
	case "DiagnosticReport":
		return DiagnosticReport{}
	case "DocumentReference":
		return DocumentReference{}
	case "DomainResource":
		return DomainResource{}
	case "Encounter":
		return Encounter{}
	case "EncounterHistory":
		return EncounterHistory{}
	case "Endpoint":
		return Endpoint{}
	case "EnrollmentRequest":
		return EnrollmentRequest{}
	case "EnrollmentResponse":
		return EnrollmentResponse{}
	case "EpisodeOfCare":
		return EpisodeOfCare{}
	case "EventDefinition":
		return EventDefinition{}
	case "Evidence":
		return Evidence{}
	case "EvidenceReport":
		return EvidenceReport{}
	case "EvidenceVariable":
		return EvidenceVariable{}
	case "ExampleScenario":
		return ExampleScenario{}
	case "ExplanationOfBenefit":
		return ExplanationOfBenefit{}
	case "FamilyMemberHistory":
		return FamilyMemberHistory{}
	case "Flag":
		return Flag{}
	case "FormularyItem":
		return FormularyItem{}
	case "GenomicStudy":
		return GenomicStudy{}
	case "Goal":
		return Goal{}
	case "GraphDefinition":
		return GraphDefinition{}
	case "Group":
		return Group{}
	case "GuidanceResponse":
		return GuidanceResponse{}
	case "HealthcareService":
		return HealthcareService{}
	case "ImagingSelection":
		return ImagingSelection{}
	case "ImagingStudy":
		return ImagingStudy{}
	case "ImmunizationEvaluation":
		return ImmunizationEvaluation{}
	case "Immunization":
		return Immunization{}
	case "ImmunizationRecommendation":
		return ImmunizationRecommendation{}
	case "ImplementationGuide":
		return ImplementationGuide{}
	case "Ingredient":
		return Ingredient{}
	case "InsurancePlan":
		return InsurancePlan{}
	case "InventoryItem":
		return InventoryItem{}
	case "InventoryReport":
		return InventoryReport{}
	case "Invoice":
		return Invoice{}
	case "Library":
		return Library{}
	case "Linkage":
		return Linkage{}
	case "List":
		return List{}
	case "Location":
		return Location{}
	case "ManufacturedItemDefinition":
		return ManufacturedItemDefinition{}
	case "Measure":
		return Measure{}
	case "MeasureReport":
		return MeasureReport{}
	case "MedicationAdministration":
		return MedicationAdministration{}
	case "MedicationDispense":
		return MedicationDispense{}
	case "Medication":
		return Medication{}
	case "MedicationKnowledge":
		return MedicationKnowledge{}
	case "MedicationRequest":
		return MedicationRequest{}
	case "MedicationStatement":
		return MedicationStatement{}
	case "MedicinalProductDefinition":
		return MedicinalProductDefinition{}
	case "MessageDefinition":
		return MessageDefinition{}
	case "MessageHeader":
		return MessageHeader{}
	case "MetadataResource":
		return MetadataResource{}
	case "MolecularSequence":
		return MolecularSequence{}
	case "NamingSystem":
		return NamingSystem{}
	case "NutritionIntake":
		return NutritionIntake{}
	case "NutritionOrder":
		return NutritionOrder{}
	case "NutritionProduct":
		return NutritionProduct{}
	case "ObservationDefinition":
		return ObservationDefinition{}
	case "Observation":
		return Observation{}
	case "OperationDefinition":
		return OperationDefinition{}
	case "OperationOutcome":
		return OperationOutcome{}
	case "OrganizationAffiliation":
		return OrganizationAffiliation{}
	case "Organization":
		return Organization{}
	case "PackagedProductDefinition":
		return PackagedProductDefinition{}
	case "Parameters":
		return Parameters{}
	case "Patient":
		return Patient{}
	case "PaymentNotice":
		return PaymentNotice{}
	case "PaymentReconciliation":
		return PaymentReconciliation{}
	case "Permission":
		return Permission{}
	case "Person":
		return Person{}
	case "PlanDefinition":
		return PlanDefinition{}
	case "Practitioner":
		return Practitioner{}
	case "PractitionerRole":
		return PractitionerRole{}
	case "Procedure":
		return Procedure{}
	case "Provenance":
		return Provenance{}
	case "Questionnaire":
		return Questionnaire{}
	case "QuestionnaireResponse":
		return QuestionnaireResponse{}
	case "RegulatedAuthorization":
		return RegulatedAuthorization{}
	case "RelatedPerson":
		return RelatedPerson{}
	case "RequestOrchestration":
		return RequestOrchestration{}
	case "Requirements":
		return Requirements{}
	case "ResearchStudy":
		return ResearchStudy{}
	case "ResearchSubject":
		return ResearchSubject{}
	case "Resource":
		return Resource{}
	case "RiskAssessment":
		return RiskAssessment{}
	case "Schedule":
		return Schedule{}
	case "SearchParameter":
		return SearchParameter{}
	case "ServiceRequest":
		return ServiceRequest{}
	case "Slot":
		return Slot{}
	case "SpecimenDefinition":
		return SpecimenDefinition{}
	case "Specimen":
		return Specimen{}
	case "StructureDefinition":
		return StructureDefinition{}
	case "StructureMap":
		return StructureMap{}
	case "Subscription":
		return Subscription{}
	case "SubscriptionStatus":
		return SubscriptionStatus{}
	case "SubscriptionTopic":
		return SubscriptionTopic{}
	case "SubstanceDefinition":
		return SubstanceDefinition{}
	case "Substance":
		return Substance{}
	case "SubstanceNucleicAcid":
		return SubstanceNucleicAcid{}
	case "SubstancePolymer":
		return SubstancePolymer{}
	case "SubstanceProtein":
		return SubstanceProtein{}
	case "SubstanceReferenceInformation":
		return SubstanceReferenceInformation{}
	case "SubstanceSourceMaterial":
		return SubstanceSourceMaterial{}
	case "SupplyDelivery":
		return SupplyDelivery{}
	case "SupplyRequest":
		return SupplyRequest{}
	case "Task":
		return Task{}
	case "TerminologyCapabilities":
		return TerminologyCapabilities{}
	case "TestPlan":
		return TestPlan{}
	case "TestReport":
		return TestReport{}
	case "TestScript":
		return TestScript{}
	case "Transport":
		return Transport{}
	case "ValueSet":
		return ValueSet{}
	case "VerificationResult":
		return VerificationResult{}
	case "VisionPrescription":
		return VisionPrescription{}
	default:
		return nil
	}
}
