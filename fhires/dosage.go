package fhires

// Dosage Type: Indicates how the medication is/was taken or should be taken by the patient.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Dosage
type Dosage struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants.
	//
	// Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty" xml:"modifierExtension,omitempty"`
	// Indicates the order in which the dosage instructions should be applied or interpreted.
	Sequence *int `bson:"sequence,omitempty" json:"sequence,omitempty" xml:"sequence,omitempty"`
	// Free text dosage instructions e.g. SIG.
	Text *string `bson:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// Supplemental instructions to the patient on how to take the medication  (e.g.
	//
	// "with meals" or"take half to one hour before food") or warnings for the patient about the medication (e.g. "may cause drowsiness" or "avoid exposure of skin to direct sunlight or sunlamps").
	//
	// Comment: Information about administration or preparation of the medication (e.g. "infuse as rapidly as possibly via intraperitoneal port" or take "immediately following drug x") should be populated in dosage.text.
	AdditionalInstruction []CodeableConcept `bson:"additionalInstruction,omitempty" json:"additionalInstruction,omitempty" xml:"additionalInstruction,omitempty"`
	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *string `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty" xml:"patientInstruction,omitempty"`
	// When medication should be administered.
	//
	// Comment: This attribute might not always be populated while the Dosage.text is expected to be populated.  If both are populated, then the Dosage.text should reflect the content of the Dosage.timing.
	Timing *Timing `bson:"timing,omitempty" json:"timing,omitempty" xml:"timing,omitempty"`
	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option).
	//
	// Comment: Can express "as needed" without a reason by setting the Boolean = True.  In this case the CodeableConcept is not populated.
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty" xml:"asNeeded,omitempty"`
	// Indicates whether the Medication is only taken based on a precondition for taking the Medication (CodeableConcept).
	//
	// Comment: Can express "as needed" with a reason by including the CodeableConcept.  In this case the Boolean is assumed to be False, then the dose is given according to the schedule and is not "prn" or "as needed".
	AsNeededFor []CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty" xml:"asNeededFor,omitempty"`
	// Body site to administer to.
	//
	// Comment: If the use case requires attributes from the BodySite resource (e.g. to identify and track separately) then use the standard extension [http://hl7.org/fhir/StructureDefinition/bodySite](http://hl7.org/fhir/extensions/StructureDefinition-bodySite.html).  May be a summary code, or a reference to a very precise definition of the location, or both.
	Site *CodeableConcept `bson:"site,omitempty" json:"site,omitempty" xml:"site,omitempty"`
	// How drug should enter body.
	Route *CodeableConcept `bson:"route,omitempty" json:"route,omitempty" xml:"route,omitempty"`
	// Technique for administering medication.
	//
	// Comment: Terminologies used often pre-coordinate this term with the route and or form of administration.
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// Depending on the resource,this is the amount of medication administered, to  be administered or typical amount to be administered.
	DoseAndRate []DosageDoseAndRate `bson:"doseAndRate,omitempty" json:"doseAndRate,omitempty" xml:"doseAndRate,omitempty"`
	// Upper limit on medication per unit of time.
	//
	// Comment: This is intended for use as an adjunct to the dosage when there is an upper cap.  For example "2 tablets every 4 hours to a maximum of 8/day".
	MaxDosePerPeriod []Ratio `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty" xml:"maxDosePerPeriod,omitempty"`
	// Upper limit on medication per administration.
	//
	// Comment: This is intended for use as an adjunct to the dosage when there is an upper cap.  For example, a body surface area related dose with a maximum amount, such as 1.5 mg/m2 (maximum 2 mg) IV over 5 – 10 minutes would have doseQuantity of 1.5 mg/m2 and maxDosePerAdministration of 2 mg.
	MaxDosePerAdministration *Quantity `bson:"maxDosePerAdministration,omitempty" json:"maxDosePerAdministration,omitempty" xml:"maxDosePerAdministration,omitempty"`
	// Upper limit on medication per lifetime of the patient.
	MaxDosePerLifetime *Quantity `bson:"maxDosePerLifetime,omitempty" json:"maxDosePerLifetime,omitempty" xml:"maxDosePerLifetime,omitempty"`
}

type DosageDoseAndRate struct {
	// Unique id for the element within a resource (for internal references).
	//
	// This may be any string value that does not contain spaces.
	ID *string `bson:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// May be used to represent additional information that is not part of the basic definition of the element.
	//
	// To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	//
	// Comment: There can be no stigma associated with the use of extensions by any application, project, or standard - regardless of the institution or jurisdiction that uses or defines the extensions.  The use of extensions is what allows the FHIR specification to retain a core level of simplicity for everyone.
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty" xml:"extension,omitempty"`
	// The kind of dose or rate specified, for example, ordered or calculated.
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Amount of medication per dose.
	//
	// Comment: Note that this specifies the quantity of the specified medication, not the quantity for each active ingredient(s). Each ingredient amount can be communicated in the Medication resource. For example, if one wants to communicate that a tablet was 375 mg, where the dose was one tablet, you can use the Medication resource to document that the tablet was comprised of 375 mg of drug XYZ. Alternatively if the dose was 375 mg, then you may only need to use the Medication resource to indicate this was a tablet. If the example were an IV such as dopamine and you wanted to communicate that 400mg of dopamine was mixed in 500 ml of some IV solution, then this would all be communicated in the Medication resource. If the administration is not intended to be instantaneous (rate is present or timing has a duration), this can be specified to convey the total amount to be administered over the period of time as indicated by the schedule e.g. 500 ml in dose, with timing used to convey that this should be done over 4 hours.
	DoseRange    *Range    `bson:"doseRange,omitempty" json:"doseRange,omitempty" xml:"doseRange,omitempty"`
	DoseQuantity *Quantity `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty" xml:"doseQuantity,omitempty"`
	// Amount of medication per unit of time.
	//
	// Comment: It is possible to supply both a rate and a doseQuantity to provide full details about how the medication is to be administered and supplied. If the rate is intended to change over time, depending on local rules/regulations, each change should be captured as a new version of the MedicationRequest with an updated rate, or captured with a new MedicationRequest with the new rate.It is possible to specify a rate over time (for example, 100 ml/hour) using either the rateRatio and rateQuantity.  The rateQuantity approach requires systems to have the capability to parse UCUM grammar where ml/hour is included rather than a specific ratio where the time is specified as the denominator.  Where a rate such as 500ml over 2 hours is specified, the use of rateRatio may be more semantically correct than specifying using a rateQuantity of 250 mg/hour.
	RateRatio    *Ratio    `bson:"rateRatio,omitempty" json:"rateRatio,omitempty" xml:"rateRatio,omitempty"`
	RateRange    *Range    `bson:"rateRange,omitempty" json:"rateRange,omitempty" xml:"rateRange,omitempty"`
	RateQuantity *Quantity `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty" xml:"rateQuantity,omitempty"`
}
