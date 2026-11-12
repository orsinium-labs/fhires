package fhires

import (
	jsontext "encoding/json/jsontext"
	"encoding/json/v2"
	"encoding/xml"
	"errors"
	"fmt"
)

type IResource interface {
	json.Marshaler
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
	_ json.MarshalerTo     = AnyResource{}
	_ json.UnmarshalerFrom = &AnyResource{}
	_ xml.Marshaler        = AnyResource{}
	_ xml.Unmarshaler      = &AnyResource{}
)

// MarshalXML implements [xml.Marshaler].
func (r AnyResource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(r.Value)
}

// UnmarshalXML implements [xml.Unmarshaler].
func (r *AnyResource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var res struct {
		ResourceType string `xml:"resourceType"`
	}
	err := d.Decode(&res)
	if err != nil {
		return err
	}
	if res.ResourceType == "" {
		return errors.New("not a valid resource")
	}
	val := GetResource(res.ResourceType)
	if val == nil {
		return fmt.Errorf("unsupported resourceType: %s", res.ResourceType)
	}
	err = d.Decode(&val)
	if err != nil {
		return err
	}
	r.Value = val
	return nil
}

// MarshalJSONTo implements [json.MarshalerTo].
func (r AnyResource) MarshalJSONTo(e *jsontext.Encoder) error {
	return json.MarshalEncode(e, r.Value)
}

// UnmarshalJSONFrom implements [json.UnmarshalerFrom].
func (r *AnyResource) UnmarshalJSONFrom(d *jsontext.Decoder) error {
	var res struct {
		ResourceType string `json:"resourceType"`
	}
	err := json.UnmarshalDecode(d, &res, json.RejectUnknownMembers(false))
	if err != nil {
		return err
	}
	if res.ResourceType == "" {
		return errors.New("not a valid resource")
	}
	val := GetResource(res.ResourceType)
	if val == nil {
		return fmt.Errorf("unsupported resourceType: %s", res.ResourceType)
	}
	err = json.UnmarshalDecode(d, &val)
	if err != nil {
		return err
	}
	r.Value = val
	return nil
}

// Get resource instance for the given resource name.
//
// For example, returns [Encounter] for "Encounter".
// Returns [nil] if there is no match.
func GetResource(resourceType string) IResource {
	return registry[resourceType]
}

var registry = map[string]IResource{
	"Account":                            Account{},
	"ActivityDefinition":                 ActivityDefinition{},
	"ActorDefinition":                    ActorDefinition{},
	"AdministrableProductDefinition":     AdministrableProductDefinition{},
	"AdverseEvent":                       AdverseEvent{},
	"AllergyIntolerance":                 AllergyIntolerance{},
	"Appointment":                        Appointment{},
	"AppointmentResponse":                AppointmentResponse{},
	"ArtifactAssessment":                 ArtifactAssessment{},
	"AuditEvent":                         AuditEvent{},
	"Basic":                              Basic{},
	"Binary":                             Binary{},
	"BiologicallyDerivedProductDispense": BiologicallyDerivedProductDispense{},
	"BiologicallyDerivedProduct":         BiologicallyDerivedProduct{},
	"BodyStructure":                      BodyStructure{},
	"Bundle":                             Bundle{},
	"CanonicalResource":                  CanonicalResource{},
	"CapabilityStatement":                CapabilityStatement{},
	"CarePlan":                           CarePlan{},
	"CareTeam":                           CareTeam{},
	"ChargeItemDefinition":               ChargeItemDefinition{},
	"ChargeItem":                         ChargeItem{},
	"Citation":                           Citation{},
	"Claim":                              Claim{},
	"ClaimResponse":                      ClaimResponse{},
	"ClinicalImpression":                 ClinicalImpression{},
	"ClinicalUseDefinition":              ClinicalUseDefinition{},
	"CodeSystem":                         CodeSystem{},
	"Communication":                      Communication{},
	"CommunicationRequest":               CommunicationRequest{},
	"CompartmentDefinition":              CompartmentDefinition{},
	"Composition":                        Composition{},
	"ConceptMap":                         ConceptMap{},
	"ConditionDefinition":                ConditionDefinition{},
	"Condition":                          Condition{},
	"Consent":                            Consent{},
	"Contract":                           Contract{},
	"CoverageEligibilityRequest":         CoverageEligibilityRequest{},
	"CoverageEligibilityResponse":        CoverageEligibilityResponse{},
	"Coverage":                           Coverage{},
	"DetectedIssue":                      DetectedIssue{},
	"DeviceAssociation":                  DeviceAssociation{},
	"DeviceDefinition":                   DeviceDefinition{},
	"DeviceDispense":                     DeviceDispense{},
	"Device":                             Device{},
	"DeviceMetric":                       DeviceMetric{},
	"DeviceRequest":                      DeviceRequest{},
	"DeviceUsage":                        DeviceUsage{},
	"DiagnosticReport":                   DiagnosticReport{},
	"DocumentReference":                  DocumentReference{},
	"DomainResource":                     DomainResource{},
	"Encounter":                          Encounter{},
	"EncounterHistory":                   EncounterHistory{},
	"Endpoint":                           Endpoint{},
	"EnrollmentRequest":                  EnrollmentRequest{},
	"EnrollmentResponse":                 EnrollmentResponse{},
	"EpisodeOfCare":                      EpisodeOfCare{},
	"EventDefinition":                    EventDefinition{},
	"Evidence":                           Evidence{},
	"EvidenceReport":                     EvidenceReport{},
	"EvidenceVariable":                   EvidenceVariable{},
	"ExampleScenario":                    ExampleScenario{},
	"ExplanationOfBenefit":               ExplanationOfBenefit{},
	"FamilyMemberHistory":                FamilyMemberHistory{},
	"Flag":                               Flag{},
	"FormularyItem":                      FormularyItem{},
	"GenomicStudy":                       GenomicStudy{},
	"Goal":                               Goal{},
	"GraphDefinition":                    GraphDefinition{},
	"Group":                              Group{},
	"GuidanceResponse":                   GuidanceResponse{},
	"HealthcareService":                  HealthcareService{},
	"ImagingSelection":                   ImagingSelection{},
	"ImagingStudy":                       ImagingStudy{},
	"ImmunizationEvaluation":             ImmunizationEvaluation{},
	"Immunization":                       Immunization{},
	"ImmunizationRecommendation":         ImmunizationRecommendation{},
	"ImplementationGuide":                ImplementationGuide{},
	"Ingredient":                         Ingredient{},
	"InsurancePlan":                      InsurancePlan{},
	"InventoryItem":                      InventoryItem{},
	"InventoryReport":                    InventoryReport{},
	"Invoice":                            Invoice{},
	"Library":                            Library{},
	"Linkage":                            Linkage{},
	"List":                               List{},
	"Location":                           Location{},
	"ManufacturedItemDefinition":         ManufacturedItemDefinition{},
	"Measure":                            Measure{},
	"MeasureReport":                      MeasureReport{},
	"MedicationAdministration":           MedicationAdministration{},
	"MedicationDispense":                 MedicationDispense{},
	"Medication":                         Medication{},
	"MedicationKnowledge":                MedicationKnowledge{},
	"MedicationRequest":                  MedicationRequest{},
	"MedicationStatement":                MedicationStatement{},
	"MedicinalProductDefinition":         MedicinalProductDefinition{},
	"MessageDefinition":                  MessageDefinition{},
	"MessageHeader":                      MessageHeader{},
	"MetadataResource":                   MetadataResource{},
	"MolecularSequence":                  MolecularSequence{},
	"NamingSystem":                       NamingSystem{},
	"NutritionIntake":                    NutritionIntake{},
	"NutritionOrder":                     NutritionOrder{},
	"NutritionProduct":                   NutritionProduct{},
	"ObservationDefinition":              ObservationDefinition{},
	"Observation":                        Observation{},
	"OperationDefinition":                OperationDefinition{},
	"OperationOutcome":                   OperationOutcome{},
	"OrganizationAffiliation":            OrganizationAffiliation{},
	"Organization":                       Organization{},
	"PackagedProductDefinition":          PackagedProductDefinition{},
	"Parameters":                         Parameters{},
	"Patient":                            Patient{},
	"PaymentNotice":                      PaymentNotice{},
	"PaymentReconciliation":              PaymentReconciliation{},
	"Permission":                         Permission{},
	"Person":                             Person{},
	"PlanDefinition":                     PlanDefinition{},
	"Practitioner":                       Practitioner{},
	"PractitionerRole":                   PractitionerRole{},
	"Procedure":                          Procedure{},
	"Provenance":                         Provenance{},
	"Questionnaire":                      Questionnaire{},
	"QuestionnaireResponse":              QuestionnaireResponse{},
	"RegulatedAuthorization":             RegulatedAuthorization{},
	"RelatedPerson":                      RelatedPerson{},
	"RequestOrchestration":               RequestOrchestration{},
	"Requirements":                       Requirements{},
	"ResearchStudy":                      ResearchStudy{},
	"ResearchSubject":                    ResearchSubject{},
	"Resource":                           BaseResource{},
	"RiskAssessment":                     RiskAssessment{},
	"Schedule":                           Schedule{},
	"SearchParameter":                    SearchParameter{},
	"ServiceRequest":                     ServiceRequest{},
	"Slot":                               Slot{},
	"SpecimenDefinition":                 SpecimenDefinition{},
	"Specimen":                           Specimen{},
	"StructureDefinition":                StructureDefinition{},
	"StructureMap":                       StructureMap{},
	"Subscription":                       Subscription{},
	"SubscriptionStatus":                 SubscriptionStatus{},
	"SubscriptionTopic":                  SubscriptionTopic{},
	"SubstanceDefinition":                SubstanceDefinition{},
	"Substance":                          Substance{},
	"SubstanceNucleicAcid":               SubstanceNucleicAcid{},
	"SubstancePolymer":                   SubstancePolymer{},
	"SubstanceProtein":                   SubstanceProtein{},
	"SubstanceReferenceInformation":      SubstanceReferenceInformation{},
	"SubstanceSourceMaterial":            SubstanceSourceMaterial{},
	"SupplyDelivery":                     SupplyDelivery{},
	"SupplyRequest":                      SupplyRequest{},
	"Task":                               Task{},
	"TerminologyCapabilities":            TerminologyCapabilities{},
	"TestPlan":                           TestPlan{},
	"TestReport":                         TestReport{},
	"TestScript":                         TestScript{},
	"Transport":                          Transport{},
	"ValueSet":                           ValueSet{},
	"VerificationResult":                 VerificationResult{},
	"VisionPrescription":                 VisionPrescription{},
}
