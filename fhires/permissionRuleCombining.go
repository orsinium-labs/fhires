package fhires

import (
	"encoding/json/v2"
	"fmt"
	"strings"
)

// Codes identifying rule combining algorithm.
//
// Status: draft
//
// http://hl7.org/fhir/ValueSet/permission-rule-combining
type PermissionRuleCombining int

const (
	PermissionRuleCombiningDenyOverrides PermissionRuleCombining = iota
	// The permit overrides combining algorithm is intended for those cases where a permit decision should have priority over a deny decision.
	PermissionRuleCombiningPermitOverrides
	// The behavior of this algorithm is identical to that of the “Deny-overrides” rule-combining algorithm with one exception.  The order in which the collection of rules is evaluated SHALL match the order as listed in the permission.
	PermissionRuleCombiningOrderedDenyOverrides
	// The behavior of this algorithm is identical to that of the “Permit-overrides” rule-combining algorithm with one exception.  The order in which the collection of rules is evaluated SHALL match the order as listed in the permission.
	PermissionRuleCombiningOrderedPermitOverrides
	// The “Deny-unless-permit” combining algorithm is intended for those cases where a permit decision should have priority over a deny decision, and an “Indeterminate” or “NotApplicable” must never be the result. It is particularly useful at the top level in a policy structure to ensure that a PDP will always return a definite “Permit” or “Deny” result.
	PermissionRuleCombiningDenyUnlessPermit
	// The “Permit-unless-deny” combining algorithm is intended for those cases where a deny decision should have priority over a permit decision, and an “Indeterminate” or “NotApplicable” must never be the result. It is particularly useful at the top level in a policy structure to ensure that a PDP will always return a definite “Permit” or “Deny” result. This algorithm has the following behavior.
	PermissionRuleCombiningPermitUnlessDeny
)

func (code PermissionRuleCombining) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}

func (code *PermissionRuleCombining) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "deny-overrides":
		*code = PermissionRuleCombiningDenyOverrides
	case "permit-overrides":
		*code = PermissionRuleCombiningPermitOverrides
	case "ordered-deny-overrides":
		*code = PermissionRuleCombiningOrderedDenyOverrides
	case "ordered-permit-overrides":
		*code = PermissionRuleCombiningOrderedPermitOverrides
	case "deny-unless-permit":
		*code = PermissionRuleCombiningDenyUnlessPermit
	case "permit-unless-deny":
		*code = PermissionRuleCombiningPermitUnlessDeny
	default:
		return fmt.Errorf("unknown PermissionRuleCombining code `%s`", s)
	}
	return nil
}

func (code PermissionRuleCombining) String() string {
	return code.Code()
}

func (code PermissionRuleCombining) Code() string {
	switch code {
	case PermissionRuleCombiningDenyOverrides:
		return "deny-overrides"
	case PermissionRuleCombiningPermitOverrides:
		return "permit-overrides"
	case PermissionRuleCombiningOrderedDenyOverrides:
		return "ordered-deny-overrides"
	case PermissionRuleCombiningOrderedPermitOverrides:
		return "ordered-permit-overrides"
	case PermissionRuleCombiningDenyUnlessPermit:
		return "deny-unless-permit"
	case PermissionRuleCombiningPermitUnlessDeny:
		return "permit-unless-deny"
	}
	return "<unknown>"
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
