package validation

import "github.com/hbttundar/diabuddy-errors"

// Rule interface defines the behavior of a validation rule.
type Rule interface {
	Validate(value any) errors.ApiErrors
	ValidationErrorMessage(baseMessage string) string
	Message() string
}

// RuleRegistry stores dynamically added rules.
type RuleRegistry struct {
	rules map[string]Rule
}

// NewRuleRegistry creates a new rule registry.
func NewRuleRegistry() *RuleRegistry {
	return &RuleRegistry{rules: make(map[string]Rule)}
}

// RegisterRule adds a new rule to the registry.
func (r *RuleRegistry) RegisterRule(name string, rule Rule) {
	r.rules[name] = rule
}

// GetRule retrieves a rule by its name.
func (r *RuleRegistry) GetRule(name string) (Rule, bool) {
	rule, exists := r.rules[name]
	return rule, exists
}
