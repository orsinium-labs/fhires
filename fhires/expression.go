package fhires

// Expression Type: A expression that is evaluated in a specified context and returns a value.
//
// The context of use of the expression must specify the context in which the expression is evaluated, and how the result of the expression is used.
//
// Status: draft
//
// http://hl7.org/fhir/StructureDefinition/Expression
type Expression struct {
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
	// A brief, natural language description of the condition that effectively communicates the intended semantics.
	Description *string `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// A short name assigned to the expression to allow for multiple reuse of the expression in the context where it is defined.
	Name *string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The media type of the language for the expression.
	Language *string `bson:"language,omitempty" json:"language,omitempty" xml:"language,omitempty"`
	// An expression in the specified language that returns a value.
	//
	// Comment: If Expression.expression and Expression.reference are both present, the Expression.expression might just be a name pointing something within the referenced content.
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty" xml:"expression,omitempty"`
	// A URI that defines where the expression is found.
	//
	// Comment: If both a reference and an expression is found, the reference SHALL point to the same expression.
	Reference *string `bson:"reference,omitempty" json:"reference,omitempty" xml:"reference,omitempty"`
}
