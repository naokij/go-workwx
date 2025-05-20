package workwx

import (
	"encoding/json"
	"testing"
)

func TestListSmartsheetFields(t *testing.T) {
	// 这个测试仅验证结构体定义和解析
	jsonStr := `{
		"errcode": 0,
		"errmsg": "ok",
		"total": 2,
		"fields": [
			{
				"field_id": "field_1",
				"field_title": "测试字段1",
				"field_type": "text",
				"property_number": {
					"decimal_places": 2,
					"use_separate": true
				}
			},
			{
				"field_id": "field_2",
				"field_title": "测试字段2",
				"field_type": "select",
				"property_select": {
					"is_quick_add": true,
					"options": [
						{
							"id": "opt_1",
							"text": "选项1",
							"style": 1
						},
						{
							"id": "opt_2",
							"text": "选项2",
							"style": 2
						}
					]
				}
			}
		]
	}`

	var resp respListSmartsheetFields
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal respListSmartsheetFields: %v", err)
	}

	if resp.Total != 2 {
		t.Errorf("Expected 2 fields, got %d", resp.Total)
	}

	if len(resp.Fields) != 2 {
		t.Errorf("Expected 2 fields in the array, got %d", len(resp.Fields))
	}

	if resp.Fields[0].FieldID != "field_1" || resp.Fields[0].FieldTitle != "测试字段1" {
		t.Errorf("First field data is incorrect")
	}

	if resp.Fields[0].PropertyNumber == nil {
		t.Errorf("Expected PropertyNumber to be set for first field")
	} else {
		if resp.Fields[0].PropertyNumber.DecimalPlaces != 2 {
			t.Errorf("Expected DecimalPlaces to be 2, got %d", resp.Fields[0].PropertyNumber.DecimalPlaces)
		}
		if !resp.Fields[0].PropertyNumber.UseSeparate {
			t.Errorf("Expected UseSeparate to be true")
		}
	}

	if resp.Fields[1].PropertySelect == nil {
		t.Errorf("Expected PropertySelect to be set for second field")
	} else {
		if !resp.Fields[1].PropertySelect.IsQuickAdd {
			t.Errorf("Expected IsQuickAdd to be true")
		}
		if len(resp.Fields[1].PropertySelect.Options) != 2 {
			t.Errorf("Expected 2 options, got %d", len(resp.Fields[1].PropertySelect.Options))
		} else {
			if resp.Fields[1].PropertySelect.Options[0].ID != "opt_1" {
				t.Errorf("Expected first option ID to be opt_1, got %s", resp.Fields[1].PropertySelect.Options[0].ID)
			}
		}
	}
}
