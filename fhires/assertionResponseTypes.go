package fhires

// The type of response code to use for assertion.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/assert-response-code-types
type AssertionResponseTypes string

const (
	// Response code is 100.
	AssertionResponseTypesContinue AssertionResponseTypes = "continue"
	// Response code is 101.
	AssertionResponseTypesSwitchingProtocols AssertionResponseTypes = "switchingProtocols"
	// Response code is 200.
	AssertionResponseTypesOkay AssertionResponseTypes = "okay"
	// Response code is 201.
	AssertionResponseTypesCreated AssertionResponseTypes = "created"
	// Response code is 202.
	AssertionResponseTypesAccepted AssertionResponseTypes = "accepted"
	// Response code is 203.
	AssertionResponseTypesNonAuthoritativeInformation AssertionResponseTypes = "nonAuthoritativeInformation"
	// Response code is 204.
	AssertionResponseTypesNoContent AssertionResponseTypes = "noContent"
	// Response code is 205.
	AssertionResponseTypesResetContent AssertionResponseTypes = "resetContent"
	// Response code is 206.
	AssertionResponseTypesPartialContent AssertionResponseTypes = "partialContent"
	// Response code is 300.
	AssertionResponseTypesMultipleChoices AssertionResponseTypes = "multipleChoices"
	// Response code is 301.
	AssertionResponseTypesMovedPermanently AssertionResponseTypes = "movedPermanently"
	// Response code is 302.
	AssertionResponseTypesFound AssertionResponseTypes = "found"
	// Response code is 303.
	AssertionResponseTypesSeeOther AssertionResponseTypes = "seeOther"
	// Response code is 304.
	AssertionResponseTypesNotModified AssertionResponseTypes = "notModified"
	// Response code is 305.
	AssertionResponseTypesUseProxy AssertionResponseTypes = "useProxy"
	// Response code is 307.
	AssertionResponseTypesTemporaryRedirect AssertionResponseTypes = "temporaryRedirect"
	// Response code is 308.
	AssertionResponseTypesPermanentRedirect AssertionResponseTypes = "permanentRedirect"
	// Response code is 400.
	AssertionResponseTypesBadRequest AssertionResponseTypes = "badRequest"
	// Response code is 401.
	AssertionResponseTypesUnauthorized AssertionResponseTypes = "unauthorized"
	// Response code is 402.
	AssertionResponseTypesPaymentRequired AssertionResponseTypes = "paymentRequired"
	// Response code is 403.
	AssertionResponseTypesForbidden AssertionResponseTypes = "forbidden"
	// Response code is 404.
	AssertionResponseTypesNotFound AssertionResponseTypes = "notFound"
	// Response code is 405.
	AssertionResponseTypesMethodNotAllowed AssertionResponseTypes = "methodNotAllowed"
	// Response code is 406.
	AssertionResponseTypesNotAcceptable AssertionResponseTypes = "notAcceptable"
	// Response code is 407.
	AssertionResponseTypesProxyAuthenticationRequired AssertionResponseTypes = "proxyAuthenticationRequired"
	// Response code is 408.
	AssertionResponseTypesRequestTimeout AssertionResponseTypes = "requestTimeout"
	// Response code is 409.
	AssertionResponseTypesConflict AssertionResponseTypes = "conflict"
	// Response code is 410.
	AssertionResponseTypesGone AssertionResponseTypes = "gone"
	// Response code is 411.
	AssertionResponseTypesLengthRequired AssertionResponseTypes = "lengthRequired"
	// Response code is 412.
	AssertionResponseTypesPreconditionFailed AssertionResponseTypes = "preconditionFailed"
	// Response code is 413.
	AssertionResponseTypesContentTooLarge AssertionResponseTypes = "contentTooLarge"
	// Response code is 414.
	AssertionResponseTypesUriTooLong AssertionResponseTypes = "uriTooLong"
	// Response code is 415.
	AssertionResponseTypesUnsupportedMediaType AssertionResponseTypes = "unsupportedMediaType"
	// Response code is 416.
	AssertionResponseTypesRangeNotSatisfiable AssertionResponseTypes = "rangeNotSatisfiable"
	// Response code is 417.
	AssertionResponseTypesExpectationFailed AssertionResponseTypes = "expectationFailed"
	// Response code is 421.
	AssertionResponseTypesMisdirectedRequest AssertionResponseTypes = "misdirectedRequest"
	// Response code is 422.
	AssertionResponseTypesUnprocessableContent AssertionResponseTypes = "unprocessableContent"
	// Response code is 426.
	AssertionResponseTypesUpgradeRequired AssertionResponseTypes = "upgradeRequired"
	// Response code is 500.
	AssertionResponseTypesInternalServerError AssertionResponseTypes = "internalServerError"
	// Response code is 501.
	AssertionResponseTypesNotImplemented AssertionResponseTypes = "notImplemented"
	// Response code is 502.
	AssertionResponseTypesBadGateway AssertionResponseTypes = "badGateway"
	// Response code is 503.
	AssertionResponseTypesServiceUnavailable AssertionResponseTypes = "serviceUnavailable"
	// Response code is 504.
	AssertionResponseTypesGatewayTimeout AssertionResponseTypes = "gatewayTimeout"
	// Response code is 505.
	AssertionResponseTypesHttpVersionNotSupported AssertionResponseTypes = "httpVersionNotSupported"
)

func (code AssertionResponseTypes) String() string {
	return code.Code()
}

func (code AssertionResponseTypes) Code() string {
	return string(code)
}

func (code AssertionResponseTypes) Display() string {
	switch code {
	case AssertionResponseTypesContinue:
		return "Continue"
	case AssertionResponseTypesSwitchingProtocols:
		return "Switching Protocols"
	case AssertionResponseTypesOkay:
		return "OK"
	case AssertionResponseTypesCreated:
		return "Created"
	case AssertionResponseTypesAccepted:
		return "Accepted"
	case AssertionResponseTypesNonAuthoritativeInformation:
		return "Non-Authoritative Information"
	case AssertionResponseTypesNoContent:
		return "No Content"
	case AssertionResponseTypesResetContent:
		return "Reset Content"
	case AssertionResponseTypesPartialContent:
		return "Partial Content"
	case AssertionResponseTypesMultipleChoices:
		return "Multiple Choices"
	case AssertionResponseTypesMovedPermanently:
		return "Moved Permanently"
	case AssertionResponseTypesFound:
		return "Found"
	case AssertionResponseTypesSeeOther:
		return "See Other"
	case AssertionResponseTypesNotModified:
		return "Not Modified"
	case AssertionResponseTypesUseProxy:
		return "Use Proxy"
	case AssertionResponseTypesTemporaryRedirect:
		return "Temporary Redirect"
	case AssertionResponseTypesPermanentRedirect:
		return "Permanent Redirect"
	case AssertionResponseTypesBadRequest:
		return "Bad Request"
	case AssertionResponseTypesUnauthorized:
		return "Unauthorized"
	case AssertionResponseTypesPaymentRequired:
		return "Payment Required"
	case AssertionResponseTypesForbidden:
		return "Forbidden"
	case AssertionResponseTypesNotFound:
		return "Not Found"
	case AssertionResponseTypesMethodNotAllowed:
		return "Method Not Allowed"
	case AssertionResponseTypesNotAcceptable:
		return "Not Acceptable"
	case AssertionResponseTypesProxyAuthenticationRequired:
		return "Proxy Authentication Required"
	case AssertionResponseTypesRequestTimeout:
		return "Request Timeout"
	case AssertionResponseTypesConflict:
		return "Conflict"
	case AssertionResponseTypesGone:
		return "Gone"
	case AssertionResponseTypesLengthRequired:
		return "Length Required"
	case AssertionResponseTypesPreconditionFailed:
		return "Precondition Failed"
	case AssertionResponseTypesContentTooLarge:
		return "Content Too Large"
	case AssertionResponseTypesUriTooLong:
		return "URI Too Long"
	case AssertionResponseTypesUnsupportedMediaType:
		return "Unsupported Media Type"
	case AssertionResponseTypesRangeNotSatisfiable:
		return "Range Not Satisfiable"
	case AssertionResponseTypesExpectationFailed:
		return "Expectation Failed"
	case AssertionResponseTypesMisdirectedRequest:
		return "Misdirected Request"
	case AssertionResponseTypesUnprocessableContent:
		return "Unprocessable Content"
	case AssertionResponseTypesUpgradeRequired:
		return "Upgrade Required"
	case AssertionResponseTypesInternalServerError:
		return "Internal Server Error"
	case AssertionResponseTypesNotImplemented:
		return "Not Implemented"
	case AssertionResponseTypesBadGateway:
		return "Bad Gateway"
	case AssertionResponseTypesServiceUnavailable:
		return "Service Unavailable"
	case AssertionResponseTypesGatewayTimeout:
		return "Gateway Timeout"
	case AssertionResponseTypesHttpVersionNotSupported:
		return "HTTP Version Not Supported"
	}
	return "<unknown>"
}

func (code AssertionResponseTypes) Definition() string {
	switch code {
	case AssertionResponseTypesContinue:
		return "Response code is 100."
	case AssertionResponseTypesSwitchingProtocols:
		return "Response code is 101."
	case AssertionResponseTypesOkay:
		return "Response code is 200."
	case AssertionResponseTypesCreated:
		return "Response code is 201."
	case AssertionResponseTypesAccepted:
		return "Response code is 202."
	case AssertionResponseTypesNonAuthoritativeInformation:
		return "Response code is 203."
	case AssertionResponseTypesNoContent:
		return "Response code is 204."
	case AssertionResponseTypesResetContent:
		return "Response code is 205."
	case AssertionResponseTypesPartialContent:
		return "Response code is 206."
	case AssertionResponseTypesMultipleChoices:
		return "Response code is 300."
	case AssertionResponseTypesMovedPermanently:
		return "Response code is 301."
	case AssertionResponseTypesFound:
		return "Response code is 302."
	case AssertionResponseTypesSeeOther:
		return "Response code is 303."
	case AssertionResponseTypesNotModified:
		return "Response code is 304."
	case AssertionResponseTypesUseProxy:
		return "Response code is 305."
	case AssertionResponseTypesTemporaryRedirect:
		return "Response code is 307."
	case AssertionResponseTypesPermanentRedirect:
		return "Response code is 308."
	case AssertionResponseTypesBadRequest:
		return "Response code is 400."
	case AssertionResponseTypesUnauthorized:
		return "Response code is 401."
	case AssertionResponseTypesPaymentRequired:
		return "Response code is 402."
	case AssertionResponseTypesForbidden:
		return "Response code is 403."
	case AssertionResponseTypesNotFound:
		return "Response code is 404."
	case AssertionResponseTypesMethodNotAllowed:
		return "Response code is 405."
	case AssertionResponseTypesNotAcceptable:
		return "Response code is 406."
	case AssertionResponseTypesProxyAuthenticationRequired:
		return "Response code is 407."
	case AssertionResponseTypesRequestTimeout:
		return "Response code is 408."
	case AssertionResponseTypesConflict:
		return "Response code is 409."
	case AssertionResponseTypesGone:
		return "Response code is 410."
	case AssertionResponseTypesLengthRequired:
		return "Response code is 411."
	case AssertionResponseTypesPreconditionFailed:
		return "Response code is 412."
	case AssertionResponseTypesContentTooLarge:
		return "Response code is 413."
	case AssertionResponseTypesUriTooLong:
		return "Response code is 414."
	case AssertionResponseTypesUnsupportedMediaType:
		return "Response code is 415."
	case AssertionResponseTypesRangeNotSatisfiable:
		return "Response code is 416."
	case AssertionResponseTypesExpectationFailed:
		return "Response code is 417."
	case AssertionResponseTypesMisdirectedRequest:
		return "Response code is 421."
	case AssertionResponseTypesUnprocessableContent:
		return "Response code is 422."
	case AssertionResponseTypesUpgradeRequired:
		return "Response code is 426."
	case AssertionResponseTypesInternalServerError:
		return "Response code is 500."
	case AssertionResponseTypesNotImplemented:
		return "Response code is 501."
	case AssertionResponseTypesBadGateway:
		return "Response code is 502."
	case AssertionResponseTypesServiceUnavailable:
		return "Response code is 503."
	case AssertionResponseTypesGatewayTimeout:
		return "Response code is 504."
	case AssertionResponseTypesHttpVersionNotSupported:
		return "Response code is 505."
	}
	return "<unknown>"
}
