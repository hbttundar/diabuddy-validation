package validation

// FieldValidator manages rules for a specific field.
type FieldValidator struct {
	validator *Validator
	field     string
}

// AddRule adds a validation rule to the field.
func (fv *FieldValidator) AddRule(rule Rule) *FieldValidator {
	fv.validator.fields[fv.field] = append(fv.validator.fields[fv.field], rule)
	return fv
}
