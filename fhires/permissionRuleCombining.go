package fhires

// Codes identifying rule combining algorithm.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/permission-rule-combining
type PermissionRuleCombining string

const (
	// The deny overrides combining algorithm is intended for those cases where a deny decision should have priority over a permit decision.
	PermissionRuleCombiningDenyOverrides PermissionRuleCombining = "deny-overrides"
	// The permit overrides combining algorithm is intended for those cases where a permit decision should have priority over a deny decision.
	PermissionRuleCombiningPermitOverrides PermissionRuleCombining = "permit-overrides"
	// The behavior of this algorithm is identical to that of the “Deny-overrides” rule-combining algorithm with one exception.  The order in which the collection of rules is evaluated SHALL match the order as listed in the permission.
	PermissionRuleCombiningOrderedDenyOverrides PermissionRuleCombining = "ordered-deny-overrides"
	// The behavior of this algorithm is identical to that of the “Permit-overrides” rule-combining algorithm with one exception.  The order in which the collection of rules is evaluated SHALL match the order as listed in the permission.
	PermissionRuleCombiningOrderedPermitOverrides PermissionRuleCombining = "ordered-permit-overrides"
	// The “Deny-unless-permit” combining algorithm is intended for those cases where a permit decision should have priority over a deny decision, and an “Indeterminate” or “NotApplicable” must never be the result. It is particularly useful at the top level in a policy structure to ensure that a PDP will always return a definite “Permit” or “Deny” result.
	PermissionRuleCombiningDenyUnlessPermit PermissionRuleCombining = "deny-unless-permit"
	// The “Permit-unless-deny” combining algorithm is intended for those cases where a deny decision should have priority over a permit decision, and an “Indeterminate” or “NotApplicable” must never be the result. It is particularly useful at the top level in a policy structure to ensure that a PDP will always return a definite “Permit” or “Deny” result. This algorithm has the following behavior.
	PermissionRuleCombiningPermitUnlessDeny PermissionRuleCombining = "permit-unless-deny"
)

func (code PermissionRuleCombining) String() string {
	return code.Code()
}

func (code PermissionRuleCombining) Code() string {
	return string(code)
}

func (code PermissionRuleCombining) Display() string {
	switch code {
	case PermissionRuleCombiningDenyOverrides:
		return "Deny-overrides"
	case PermissionRuleCombiningPermitOverrides:
		return "Permit-overrides"
	case PermissionRuleCombiningOrderedDenyOverrides:
		return "Ordered-deny-overrides"
	case PermissionRuleCombiningOrderedPermitOverrides:
		return "Ordered-permit-overrides"
	case PermissionRuleCombiningDenyUnlessPermit:
		return "Deny-unless-permit"
	case PermissionRuleCombiningPermitUnlessDeny:
		return "Permit-unless-deny"
	}
	return "<unknown>"
}

func (code PermissionRuleCombining) Definition() string {
	switch code {
	case PermissionRuleCombiningDenyOverrides:
		return "The deny overrides combining algorithm is intended for those cases where a deny decision should have priority over a permit decision."
	case PermissionRuleCombiningPermitOverrides:
		return "The permit overrides combining algorithm is intended for those cases where a permit decision should have priority over a deny decision."
	case PermissionRuleCombiningOrderedDenyOverrides:
		return "The behavior of this algorithm is identical to that of the “Deny-overrides” rule-combining algorithm with one exception.  The order in which the collection of rules is evaluated SHALL match the order as listed in the permission."
	case PermissionRuleCombiningOrderedPermitOverrides:
		return "The behavior of this algorithm is identical to that of the “Permit-overrides” rule-combining algorithm with one exception.  The order in which the collection of rules is evaluated SHALL match the order as listed in the permission."
	case PermissionRuleCombiningDenyUnlessPermit:
		return "The “Deny-unless-permit” combining algorithm is intended for those cases where a permit decision should have priority over a deny decision, and an “Indeterminate” or “NotApplicable” must never be the result. It is particularly useful at the top level in a policy structure to ensure that a PDP will always return a definite “Permit” or “Deny” result."
	case PermissionRuleCombiningPermitUnlessDeny:
		return "The “Permit-unless-deny” combining algorithm is intended for those cases where a deny decision should have priority over a permit decision, and an “Indeterminate” or “NotApplicable” must never be the result. It is particularly useful at the top level in a policy structure to ensure that a PDP will always return a definite “Permit” or “Deny” result. This algorithm has the following behavior."
	}
	return "<unknown>"
}
