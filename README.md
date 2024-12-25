<p align="center"><img src="art/diabuddy-validation.webp" alt="Diabuddy Validation Package"></p>

# Diabuddy Validation Package

## Introduction

The **Diabuddy Validation package** provides a powerful and extensible library to centralize validation logic across all Diabuddy APIs. It ensures consistency, reusability, and simplicity for validating user input, dynamically applying validation rules inspired by frameworks like Laravel.

### Features
- Centralized validation for all Diabuddy APIs.
- Dynamic rule registration and application.
- Modular design to support custom validation rules.
- Unified error format leveraging the **Diabuddy Errors** package.

---

## Installation

With Go's module support, Go [build|run|test] automatically fetches the necessary dependencies when you add the import in your code:

```go
import "github.com/hbttundar/diabuddy-validation"
```

Alternatively, use `go get`:

```bash
go get -u github.com/hbttundar/diabuddy-validation
```

---

## Usage

### Basic Example

```go
package main

import (
	"fmt"
	"github.com/hbttundar/diabuddy-validation"
	"github.com/hbttundar/diabuddy-validation/rules"
)

func main() {
	// Initialize the validator
	validator := validation.NewValidator()

	// Add validation rules
	validator.ForField("email").AddRule(rules.EmailRule{})
	validator.ForField("password").AddRule(rules.PasswordRule{})

	// Input data to validate
	data := map[string]string{
		"email":    "invalid-email",
		"password": "short",
	}

	// Apply validation rules
	for field, value := range data {
		validator.ForField(field).AddRule(rules.EmailRule{}).AddRule(rules.PasswordRule{})
	}

	// Validate and handle errors
	validator.Validate()
	if validator.HasErrors() {
		for _, err := range validator.Errors() {
			fmt.Printf("Field: %s, Error: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("Validation passed!")
	}
}
```

---

### Dynamic Rule Registration

You can register your own custom rules dynamically to extend the library’s capabilities:

```go
package rules

import "github.com/hbttundar/diabuddy-errors"

type CustomRule struct{}

func (c CustomRule) Validate(value any) errors.ApiErrors {
	if value != "expected_value" {
		return errors.NewApiError(errors.BadRequestErrorType, "value does not match expected")
	}
	return nil
}

func (c CustomRule) Message() string {
	return "must match the expected value"
}
```

Then apply it:

```go
validator.ForField("customField").AddRule(CustomRule{})
```

---

## Key Rules

- **EmailRule**: Validates email format and domain.
- **PasswordRule**: Ensures password strength (e.g., length, special characters).
- **PhoneNumberRule**: Validates phone number format using `libphonenumbers`.
- **UuidRule**: Ensures UUIDs are valid and non-nil.

---

## Extending the Library

You can easily add new validation rules to fit your specific needs:

1. Create a new struct implementing the `Rule` interface:

   ```go
   type CustomRule struct{}
   
   func (c CustomRule) Validate(value any) errors.ApiErrors {
       // Add validation logic here
       return nil
   }
   
   func (c CustomRule) Message() string {
       return "custom validation message"
   }
   ```

2. Add the rule to your validator:

   ```go
   validator.ForField("customField").AddRule(CustomRule{})
   ```

---

## Dependencies

The **Diabuddy Validation** package relies on:
- [Diabuddy Errors](https://github.com/hbttundar/diabuddy-errors): For unified error formatting.
- [libphonenumbers](https://github.com/nyaruka/phonenumbers): For phone number validation.

---

## Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request with a clear description of your changes.

---

## License

This library is licensed under the MIT License. See the [MIT License](https://opensource.org/licenses/MIT) for details.

```

