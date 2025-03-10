package utils

import (
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type (
	YamlTemplateMetadata struct {
		Default  string `yaml:"default"`
		Required bool   `yaml:"required"`
	}
	YamlTemplate struct {
		Content  string                          `json:"content"`
		Metadata map[string]YamlTemplateMetadata `json:"metadata"`
	}
)

var (
	yamlTemplateSeparatorPattern = regexp.MustCompile(`(?m)^[-]{3,}$`)
)

// ParseYamlTemplate parses a YAML template string into a YamlTemplate struct.
func ParseYamlTemplate(input string) (*YamlTemplate, error) {
	template := &YamlTemplate{
		Content:  input,
		Metadata: map[string]YamlTemplateMetadata{},
	}

	parts := yamlTemplateSeparatorPattern.Split(input, 2)
	if len(parts) > 1 {
		template.Content = strings.TrimSpace(parts[1])

		unmarshalError := yaml.Unmarshal([]byte(parts[0]), &template.Metadata)
		if unmarshalError != nil {
			return template, unmarshalError
		}
	}

	return template, nil
}

// RenderYamlTemplate renders a YAML template with the provided variables.
func RenderYamlTemplate(input string, variables map[string]string) (*YamlTemplate, error) {
	template, parseError := ParseYamlTemplate(input)
	if parseError != nil {
		return template, parseError
	}

	for key, metadata := range template.Metadata {
		value := metadata.Default
		overrideValue, isSet := variables[key]
		if isSet {
			value = overrideValue
		}

		pattern, compileError := regexp.Compile(`{{\s+?` + key + `\s+?}}`)
		if compileError != nil {
			continue
		}

		template.Content = pattern.ReplaceAllString(template.Content, value)
	}

	return template, parseError
}
