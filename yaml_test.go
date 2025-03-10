package utils

import (
	"reflect"
	"testing"
)

type yamlTemplateTestCase struct {
	input     string
	output    string
	variables map[string]string
	template  *YamlTemplate
}

var (
	validParseYamlTemplateCases = []yamlTemplateTestCase{
		{
			input:  "PLACEHOLDER:\n    default: placeholder\n    required: true\n---\nname: {{ PLACEHOLDER }}",
			output: "name: placeholder",
			variables: map[string]string{
				"PLACEHOLDER": "placeholder",
			},
			template: &YamlTemplate{
				Content: "name: {{ PLACEHOLDER }}",
				Metadata: map[string]YamlTemplateMetadata{
					"PLACEHOLDER": {
						Default:  "placeholder",
						Required: true,
					},
				},
			},
		},
		{
			input:  "PLACEHOLDER:\n    default: placeholder\n    required: true\n---\nname: {{ PLACEHOLDER }}",
			output: "name: value",
			variables: map[string]string{
				"PLACEHOLDER": "value",
			},
			template: &YamlTemplate{
				Content: "name: {{ PLACEHOLDER }}",
				Metadata: map[string]YamlTemplateMetadata{
					"PLACEHOLDER": {
						Default:  "placeholder",
						Required: true,
					},
				},
			},
		},
		{
			input:     "name: {{ PLACEHOLDER }}",
			output:    "name: {{ PLACEHOLDER }}",
			variables: map[string]string{},
			template: &YamlTemplate{
				Content:  "name: {{ PLACEHOLDER }}",
				Metadata: map[string]YamlTemplateMetadata{},
			},
		},
	}
	invalidParseYamlTemplateCases = []yamlTemplateTestCase{
		{
			input: "PLACEHOLDER:\n default: placeholder\n  required: true\n---\nname: {{ PLACEHOLDER }}",
		},
		{
			input: "PLACEHOLDER:\n    default: placeholder\n    required: string\n---\nname: {{ PLACEHOLDER }}",
		},
	}
)

func TestParseYamlTemplate(test *testing.T) {
	for _, testCase := range validParseYamlTemplateCases {
		yamlTemplate, parseError := ParseYamlTemplate(testCase.input)
		if parseError != nil {
			test.Fatalf("failed to parse yaml template: %v", parseError)
		}

		if !reflect.DeepEqual(yamlTemplate, testCase.template) {
			test.Errorf("expected %v, got %v", testCase.template, yamlTemplate)
		}
	}

	for _, testCase := range invalidParseYamlTemplateCases {
		yamlTemplate, parseError := ParseYamlTemplate(testCase.input)
		if parseError == nil {
			test.Errorf("expected error, got %v", yamlTemplate)
		}
	}
}

func TestRenderYamlTemplate(test *testing.T) {
	for _, testCase := range validParseYamlTemplateCases {
		yamlTemplate, renderError := RenderYamlTemplate(testCase.input, testCase.variables)
		if renderError != nil {
			test.Fatalf("failed to render yaml template: %v", renderError)
		}

		if yamlTemplate.Content != testCase.output {
			test.Errorf("expected %v, got %v", testCase.output, yamlTemplate.Content)
		}
	}
}
