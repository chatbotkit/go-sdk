package agent

import (
	"reflect"
	"strings"
	"testing"

	"github.com/chatbotkit/go-sdk/types"
)

func TestValidateRemoteConversationOptions(t *testing.T) {
	err := validateRemoteConversationOptions("conv_123", []Message{{Type: "user", Text: "hello"}}, "gpt-4o", "", "dataset_123", "")
	if err == nil {
		t.Fatal("expected validation error for unsupported remote options")
	}

	message := err.Error()
	for _, part := range []string{"messages", "model", "datasetID"} {
		if !strings.Contains(message, part) {
			t.Fatalf("expected validation error to mention %s, got %q", part, message)
		}
	}
}

func TestValidateRemoteConversationOptionsAllowsStatefulInputs(t *testing.T) {
	err := validateRemoteConversationOptions("conv_123", nil, "", "", "", "")
	if err != nil {
		t.Fatalf("expected remote options to be valid, got %v", err)
	}
}

func TestBuildToolParameters(t *testing.T) {
	properties, required := buildToolParameters(FunctionParameters{
		"properties": map[string]any{
			"location": map[string]any{"type": "string"},
			"count":    map[string]any{"type": "number"},
		},
		"required": []any{"location"},
	})

	if len(properties) != 2 {
		t.Fatalf("expected 2 properties, got %d", len(properties))
	}
	if len(required) != 1 || required[0] != "location" {
		t.Fatalf("unexpected required values: %v", required)
	}
}

func TestConvertMessageExtensions(t *testing.T) {
	backstory := "You are a helpful assistant"
	datasetName := "kb"
	skillsetName := "ops"
	secretID := "secret_123"
	spaceID := "space_123"

	extensions := &types.ConversationCompleteRequestExtensions{
		Backstory: &backstory,
		Datasets: []types.CompleteDataset{{
			Name: &datasetName,
			Records: []types.CompleteRecord{{
				Text: "A record",
				Meta: map[string]interface{}{"source": "test"},
			}},
		}},
		Features: []types.CompleteFeature{{
			Name:    "skills",
			Options: map[string]interface{}{"enabled": true},
		}},
		Skillsets: []types.CompleteSkillset{{
			Name: &skillsetName,
			Abilities: []types.CompleteAbility{{
				Name:           "lookup",
				Description:    "Look something up",
				Instruction:    "Search the knowledge base",
				Meta:           map[string]interface{}{"kind": "search"},
				LinkedSecretID: &secretID,
				LinkedSpaceID:  &spaceID,
			}},
		}},
	}

	converted := convertMessageExtensions(extensions)
	if converted == nil {
		t.Fatal("expected converted extensions")
	}
	if converted.Backstory == nil || *converted.Backstory != backstory {
		t.Fatalf("unexpected backstory: %#v", converted.Backstory)
	}
	if len(converted.Datasets) != 1 || len(converted.Datasets[0].Records) != 1 {
		t.Fatalf("unexpected dataset conversion: %#v", converted.Datasets)
	}
	if len(converted.Features) != 1 || converted.Features[0].Name != "skills" {
		t.Fatalf("unexpected feature conversion: %#v", converted.Features)
	}
	if len(converted.Skillsets) != 1 || len(converted.Skillsets[0].Abilities) != 1 {
		t.Fatalf("unexpected skillset conversion: %#v", converted.Skillsets)
	}
	ability := converted.Skillsets[0].Abilities[0]
	if ability.LinkedSecretID == nil || *ability.LinkedSecretID != secretID {
		t.Fatalf("unexpected ability linkedSecretId: %#v", ability)
	}
	if ability.LinkedSpaceID == nil || *ability.LinkedSpaceID != spaceID {
		t.Fatalf("unexpected ability linkedSpaceId: %#v", ability)
	}
}

// fillStructNonZero sets every exported field of the struct pointed to by v to
// a non-zero value derived from the field name, so a round-trip can detect any
// field that is silently dropped.
func fillStructNonZero(t *testing.T, v reflect.Value) {
	t.Helper()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		name := v.Type().Field(i).Name
		if !field.CanSet() {
			continue
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(name)
		case reflect.Ptr:
			elem := reflect.New(field.Type().Elem())
			if elem.Elem().Kind() == reflect.String {
				elem.Elem().SetString(name)
			} else if elem.Elem().Kind() == reflect.Struct {
				fillStructNonZero(t, elem.Elem())
			}
			field.Set(elem)
		case reflect.Map:
			m := reflect.MakeMap(field.Type())
			m.SetMapIndex(reflect.ValueOf(name), reflect.ValueOf(name).Convert(field.Type().Elem()))
			field.Set(m)
		case reflect.Slice:
			s := reflect.MakeSlice(field.Type(), 1, 1)
			if s.Index(0).Kind() == reflect.Struct {
				fillStructNonZero(t, s.Index(0))
			}
			field.Set(s)
		case reflect.Bool:
			field.SetBool(true)
		case reflect.Int, reflect.Int64, reflect.Float64:
			field.SetInt(1)
		default:
			t.Fatalf("fillStructNonZero: unsupported kind %s for field %s", field.Kind(), name)
		}
	}
}

// TestConvertMessageExtensionsAbilityParity is a parity guard: every field on
// the inline ability request type must survive convertMessageExtensions. If the
// generated types gain a new ability field that the conversion does not copy,
// this test fails.
func TestConvertMessageExtensionsAbilityParity(t *testing.T) {
	inType := reflect.TypeOf(types.CompleteAbility{})
	outType := reflect.TypeOf(types.MessageCompleteAbility{})

	for i := 0; i < inType.NumField(); i++ {
		if _, ok := outType.FieldByName(inType.Field(i).Name); !ok {
			t.Fatalf("output ability type is missing input field %s", inType.Field(i).Name)
		}
	}
	for i := 0; i < outType.NumField(); i++ {
		if _, ok := inType.FieldByName(outType.Field(i).Name); !ok {
			t.Fatalf("input ability type is missing output field %s", outType.Field(i).Name)
		}
	}

	var input types.CompleteAbility
	fillStructNonZero(t, reflect.ValueOf(&input).Elem())

	skillsetName := "ops"
	converted := convertMessageExtensions(&types.ConversationCompleteRequestExtensions{
		Skillsets: []types.CompleteSkillset{{
			Name:      &skillsetName,
			Abilities: []types.CompleteAbility{input},
		}},
	})
	if converted == nil || len(converted.Skillsets) != 1 || len(converted.Skillsets[0].Abilities) != 1 {
		t.Fatalf("unexpected skillset conversion: %#v", converted)
	}

	out := reflect.ValueOf(converted.Skillsets[0].Abilities[0])
	for i := 0; i < out.NumField(); i++ {
		name := out.Type().Field(i).Name
		if out.Field(i).IsZero() {
			t.Fatalf("convertMessageExtensions dropped ability field %s", name)
		}
		in := reflect.ValueOf(input).FieldByName(name)
		if !reflect.DeepEqual(in.Interface(), out.Field(i).Interface()) {
			t.Fatalf("ability field %s changed during conversion: %#v != %#v", name, in.Interface(), out.Field(i).Interface())
		}
	}
}
