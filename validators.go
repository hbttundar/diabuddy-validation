package validation

// Validator dynamically validates fields against rules.
type Validator struct {
	fields map[string][]Rule
	errors []ValidationError
}

// NewValidator creates a new Validator instance.
func NewValidator() *Validator {
	return &Validator{
		fields: make(map[string][]Rule),
		errors: []ValidationError{},
	}
}

// ForField defines rules for a specific field.
func (v *Validator) ForField(field string) *FieldValidator {
	return &FieldValidator{
		validator: v,
		field:     field,
	}
}

// Validate runs all registered rules on all fields.
func (v *Validator) Validate() {
	for field, rules := range v.fields {
		for _, rule := range rules {
			if err := rule.Validate(field); err != nil {
				v.errors = append(v.errors, ValidationError{
					Field:   field,
					Message: err.Error(),
				})
			}
		}
	}
}

// Errors returns all validation errors.
func (v *Validator) Errors() []ValidationError {
	return v.errors
}

// HasErrors checks if any validation errors exist.
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}