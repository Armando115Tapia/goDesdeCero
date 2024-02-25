package transactionRuleSamples

import (
	"strings"
)

func ValidFieldNameForTrxRules(field string) bool {
	invalidFieldNames := []string{"Name", "Action", "ProcessorName", "Code", "Message", "PartnerValidator", "Country", "Kind", "CreatedAt", "UpdatedAt", "Status", "GeneralAction"}

	return linq.From(invalidFieldNames).AllT(func(invalidField string) bool { return !strings.EqualFold(invalidField, field) })
}
