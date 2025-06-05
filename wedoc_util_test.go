package workwx

import (
	"testing"
)

func TestSafeGetString(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantValue string
		wantOk    bool
	}{
		{"字符串值", "测试字符串", "测试字符串", true},
		{"nil值", nil, "", false},
		{"非字符串值", 123, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := SafeGetString(tt.value)
			if gotValue != tt.wantValue || gotOk != tt.wantOk {
				t.Errorf("SafeGetString() = (%v, %v), want (%v, %v)", gotValue, gotOk, tt.wantValue, tt.wantOk)
			}
		})
	}
}

func TestSafeGetFloat64(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantValue float64
		wantOk    bool
	}{
		{"float64值", 123.45, 123.45, true},
		{"int值", 123, 123.0, true},
		{"nil值", nil, 0, false},
		{"非数字值", "123", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := SafeGetFloat64(tt.value)
			if gotValue != tt.wantValue || gotOk != tt.wantOk {
				t.Errorf("SafeGetFloat64() = (%v, %v), want (%v, %v)", gotValue, gotOk, tt.wantValue, tt.wantOk)
			}
		})
	}
}

func TestSafeGetBool(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantValue bool
		wantOk    bool
	}{
		{"true值", true, true, true},
		{"false值", false, false, true},
		{"nil值", nil, false, false},
		{"非布尔值", "true", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := SafeGetBool(tt.value)
			if gotValue != tt.wantValue || gotOk != tt.wantOk {
				t.Errorf("SafeGetBool() = (%v, %v), want (%v, %v)", gotValue, gotOk, tt.wantValue, tt.wantOk)
			}
		})
	}
}

func TestSafeGetStringSlice(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantValue []string
		wantOk    bool
	}{
		{"字符串切片", []string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{"interface切片", []interface{}{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{"nil值", nil, nil, false},
		{"非切片值", "abc", nil, false},
		{"混合类型切片", []interface{}{"a", 123, "c"}, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := SafeGetStringSlice(tt.value)
			if gotOk != tt.wantOk {
				t.Errorf("SafeGetStringSlice() ok = %v, want %v", gotOk, tt.wantOk)
				return
			}
			if !gotOk {
				return // 不需要比较值
			}
			if len(gotValue) != len(tt.wantValue) {
				t.Errorf("SafeGetStringSlice() = %v, want %v", gotValue, tt.wantValue)
				return
			}
			for i, v := range gotValue {
				if v != tt.wantValue[i] {
					t.Errorf("SafeGetStringSlice() 索引 %d = %v, want %v", i, v, tt.wantValue[i])
				}
			}
		})
	}
}

func TestRecordGetters(t *testing.T) {
	// 创建一个包含各种类型字段的Record进行测试
	record := Record{
		RecordID:    "1",
		CreateTime:  "2023-05-30 12:00:00",
		UpdateTime:  "2023-05-30 14:00:00",
		CreatorName: "张三",
		UpdaterName: "李四",
		Values: map[string]interface{}{
			"文本字段":  "这是一段文本",
			"数字字段":  123.45,
			"复选框字段": true,
			"日期字段": map[string]interface{}{
				"type":  "DATE",
				"value": "2023-05-30",
			},
			"日期时间字段": map[string]interface{}{
				"type":  "DATETIME",
				"value": "2023-05-30 12:00:00",
			},
			"日期范围字段": map[string]interface{}{
				"type":  "DATE_RANGE",
				"value": []string{"2023-05-01", "2023-05-31"},
			},
			"图片字段": []string{"img1", "img2"},
			"文件字段": []map[string]interface{}{
				{
					"file_id":   "file1",
					"file_name": "文件1.docx",
				},
				{
					"file_id":   "file2",
					"file_name": "文件2.xlsx",
				},
			},
			"成员字段":   []string{"user1", "user2"},
			"超链接字段":  map[string]interface{}{"url": "https://example.com", "text": "示例链接"},
			"单选字段":   "option1",
			"多选字段":   []string{"option1", "option2"},
			"关联字段":   []string{"record1", "record2"},
			"地理位置字段": map[string]interface{}{"address": "北京市海淀区", "latitude": 39.9, "longitude": 116.3},
			"货币字段":   123.45,
			"群字段":    []string{"group1", "group2"},
			"百分比字段":  0.75,
		},
	}

	t.Run("GetTextValue", func(t *testing.T) {
		value, err := record.GetTextValue("文本字段")
		if err != nil {
			t.Errorf("GetTextValue() 错误: %v", err)
			return
		}
		if value != "这是一段文本" {
			t.Errorf("GetTextValue() = %v, want %v", value, "这是一段文本")
		}

		// 测试错误情况
		_, err = record.GetTextValue("不存在的字段")
		if err == nil {
			t.Errorf("GetTextValue() 应当返回错误")
		}
	})

	t.Run("GetNumberValue", func(t *testing.T) {
		value, err := record.GetNumberValue("数字字段")
		if err != nil {
			t.Errorf("GetNumberValue() 错误: %v", err)
			return
		}
		if value != 123.45 {
			t.Errorf("GetNumberValue() = %v, want %v", value, 123.45)
		}
	})

	t.Run("GetCheckboxValue", func(t *testing.T) {
		value, err := record.GetCheckboxValue("复选框字段")
		if err != nil {
			t.Errorf("GetCheckboxValue() 错误: %v", err)
			return
		}
		if value != true {
			t.Errorf("GetCheckboxValue() = %v, want %v", value, true)
		}
	})

	t.Run("GetDateString", func(t *testing.T) {
		value, err := record.GetDateString("日期字段")
		if err != nil {
			t.Errorf("GetDateString() 错误: %v", err)
			return
		}
		if value != "2023-05-30" {
			t.Errorf("GetDateString() = %v, want %v", value, "2023-05-30")
		}
	})

	t.Run("GetDateRangeStrings", func(t *testing.T) {
		start, end, err := record.GetDateRangeStrings("日期范围字段")
		if err != nil {
			t.Errorf("GetDateRangeStrings() 错误: %v", err)
			return
		}
		if start != "2023-05-01" || end != "2023-05-31" {
			t.Errorf("GetDateRangeStrings() = (%v, %v), want (%v, %v)", start, end, "2023-05-01", "2023-05-31")
		}
	})

	t.Run("GetImageIDs", func(t *testing.T) {
		ids, err := record.GetImageIDs("图片字段")
		if err != nil {
			t.Errorf("GetImageIDs() 错误: %v", err)
			return
		}
		if len(ids) != 2 || ids[0] != "img1" || ids[1] != "img2" {
			t.Errorf("GetImageIDs() = %v, want %v", ids, []string{"img1", "img2"})
		}
	})

	t.Run("GetURLValue", func(t *testing.T) {
		url, text, err := record.GetURLValue("超链接字段")
		if err != nil {
			t.Errorf("GetURLValue() 错误: %v", err)
			return
		}
		if url != "https://example.com" || text != "示例链接" {
			t.Errorf("GetURLValue() = (%v, %v), want (%v, %v)", url, text, "https://example.com", "示例链接")
		}
	})

	t.Run("GetCurrencyValue", func(t *testing.T) {
		value, err := record.GetCurrencyValue("货币字段")
		if err != nil {
			t.Errorf("GetCurrencyValue() 错误: %v", err)
			return
		}
		if value != 123.45 {
			t.Errorf("GetCurrencyValue() = %v, want %v", value, 123.45)
		}
	})

	t.Run("GetPercentageValue", func(t *testing.T) {
		value, err := record.GetPercentageValue("百分比字段")
		if err != nil {
			t.Errorf("GetPercentageValue() 错误: %v", err)
			return
		}
		if value != 0.75 {
			t.Errorf("GetPercentageValue() = %v, want %v", value, 0.75)
		}
	})
}
