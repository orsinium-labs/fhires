package fhires

// Codes to identify how UDI data was entered.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/udi-entry-type
type UDIEntryType string

const (
	// a barcodescanner captured the data from the device label.
	UDIEntryTypeBarcode UDIEntryType = "barcode"
	// An RFID chip reader captured the data from the device label.
	UDIEntryTypeRfid UDIEntryType = "rfid"
	// The data was read from the label by a person and manually entered. (e.g.  via a keyboard).
	UDIEntryTypeManual UDIEntryType = "manual"
	// The data originated from a patient's implant card and was read by an operator.
	UDIEntryTypeCard UDIEntryType = "card"
	// The data originated from a patient source and was not directly scanned or read from a label or card.
	UDIEntryTypeSelfReported UDIEntryType = "self-reported"
	// The UDI information was received electronically from the device through a communication protocol, such as the IEEE 11073 20601 version 4 exchange protocol over Bluetooth or USB.
	UDIEntryTypeElectronicTransmission UDIEntryType = "electronic-transmission"
	// The method of data capture has not been determined.
	UDIEntryTypeUnknown UDIEntryType = "unknown"
)

func (code UDIEntryType) String() string {
	return code.Code()
}

func (code UDIEntryType) Code() string {
	return string(code)
}

func (code UDIEntryType) Display() string {
	switch code {
	case UDIEntryTypeBarcode:
		return "Barcode"
	case UDIEntryTypeRfid:
		return "RFID"
	case UDIEntryTypeManual:
		return "Manual"
	case UDIEntryTypeCard:
		return "Card"
	case UDIEntryTypeSelfReported:
		return "Self Reported"
	case UDIEntryTypeElectronicTransmission:
		return "Electronic Transmission"
	case UDIEntryTypeUnknown:
		return "Unknown"
	}
	return "<unknown>"
}

func (code UDIEntryType) Definition() string {
	switch code {
	case UDIEntryTypeBarcode:
		return "a barcodescanner captured the data from the device label."
	case UDIEntryTypeRfid:
		return "An RFID chip reader captured the data from the device label."
	case UDIEntryTypeManual:
		return "The data was read from the label by a person and manually entered. (e.g.  via a keyboard)."
	case UDIEntryTypeCard:
		return "The data originated from a patient's implant card and was read by an operator."
	case UDIEntryTypeSelfReported:
		return "The data originated from a patient source and was not directly scanned or read from a label or card."
	case UDIEntryTypeElectronicTransmission:
		return "The UDI information was received electronically from the device through a communication protocol, such as the IEEE 11073 20601 version 4 exchange protocol over Bluetooth or USB."
	case UDIEntryTypeUnknown:
		return "The method of data capture has not been determined."
	}
	return "<unknown>"
}
