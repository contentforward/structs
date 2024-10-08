package structs

import (
	"fmt"
)

type Manager struct {
	structure            any
	ruleFuncs            map[string]RuleFunc
	validationMessageTag string
	tags                 []string
}

var DefaultTags = []string{"arg", "short", "env", "json", "yaml"}

func NewManager(structure any, rules map[string]RuleFunc, tags ...string) *Manager {
	return &Manager{
		structure:            structure,
		validationMessageTag: "json",
		ruleFuncs:            rules,
		tags:                 tags,
	}
}

func NewManagerWithValidationTag(structure any, validationTag string, tags ...string) *Manager {
	return &Manager{
		structure:            structure,
		validationMessageTag: validationTag,
		tags:                 tags,
	}
}

func (m *Manager) Validate(inputs map[string]any) (map[string][]string, error) {
	structFields, err := GetStructFields(m.structure)
	if err != nil {
		return nil, fmt.Errorf("error getting struct fields for validation: %w", err)
	}

	errors, err := ValidateStructFields(m.ruleFuncs, structFields, inputs, m.validationMessageTag, m.tags...)
	if err != nil {
		return nil, fmt.Errorf("error validating struct with inputs: %w", err)
	}

	return errors, nil
}

func (m *Manager) SetFields(inputs map[string]any) error {
	err := SetStructFields(m.structure, m.tags, inputs)
	if err != nil {
		return fmt.Errorf("error setting struct fields: %w", err)
	}

	return nil
}
