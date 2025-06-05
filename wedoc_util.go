package workwx

import (
	"fmt"
)

// SafeGetString 安全地从interface{}获取字符串值
func SafeGetString(value interface{}) (string, bool) {
	if value == nil {
		return "", false
	}

	str, ok := value.(string)
	return str, ok
}

// SafeGetFloat64 安全地从interface{}获取float64值
func SafeGetFloat64(value interface{}) (float64, bool) {
	if value == nil {
		return 0, false
	}

	switch v := value.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case int32:
		return float64(v), true
	}

	return 0, false
}

// SafeGetBool 安全地从interface{}获取布尔值
func SafeGetBool(value interface{}) (bool, bool) {
	if value == nil {
		return false, false
	}

	b, ok := value.(bool)
	return b, ok
}

// SafeGetStringSlice 安全地从interface{}获取字符串切片
func SafeGetStringSlice(value interface{}) ([]string, bool) {
	if value == nil {
		return nil, false
	}

	// 检查是否直接是字符串切片
	strSlice, ok := value.([]string)
	if ok {
		return strSlice, true
	}

	// 检查是否是[]interface{}，可以转换为[]string
	interfaceSlice, ok := value.([]interface{})
	if !ok {
		return nil, false
	}

	strSlice = make([]string, 0, len(interfaceSlice))
	for _, v := range interfaceSlice {
		str, ok := v.(string)
		if !ok {
			return nil, false
		}
		strSlice = append(strSlice, str)
	}

	return strSlice, true
}

// SafeGetMap 安全地从interface{}获取map
func SafeGetMap(value interface{}) (map[string]interface{}, bool) {
	if value == nil {
		return nil, false
	}

	m, ok := value.(map[string]interface{})
	return m, ok
}

// SafeGetMapSlice 安全地从interface{}获取map切片
func SafeGetMapSlice(value interface{}) ([]map[string]interface{}, bool) {
	if value == nil {
		return nil, false
	}

	// 检查是否直接是map切片
	mapSlice, ok := value.([]map[string]interface{})
	if ok {
		return mapSlice, true
	}

	// 检查是否是[]interface{}，可以转换为[]map[string]interface{}
	interfaceSlice, ok := value.([]interface{})
	if !ok {
		return nil, false
	}

	mapSlice = make([]map[string]interface{}, 0, len(interfaceSlice))
	for _, v := range interfaceSlice {
		m, ok := v.(map[string]interface{})
		if !ok {
			return nil, false
		}
		mapSlice = append(mapSlice, m)
	}

	return mapSlice, true
}

// SafeGetDateValue 安全地从interface{}获取日期值信息
func SafeGetDateValue(value interface{}) (dateType string, dateValue interface{}, ok bool) {
	m, ok := SafeGetMap(value)
	if !ok {
		return "", nil, false
	}

	typeStr, ok := SafeGetString(m["type"])
	if !ok {
		return "", nil, false
	}

	dateValue = m["value"]
	return typeStr, dateValue, true
}

// GetTextValue 从Record获取文本类型字段的值
func (r *Record) GetTextValue(fieldKey string) (string, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return "", fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	strValue, ok := SafeGetString(value)
	if !ok {
		return "", fmt.Errorf("字段 %s 不是文本类型", fieldKey)
	}

	return strValue, nil
}

// GetNumberValue 从Record获取数字类型字段的值
func (r *Record) GetNumberValue(fieldKey string) (float64, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return 0, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	numValue, ok := SafeGetFloat64(value)
	if !ok {
		return 0, fmt.Errorf("字段 %s 不是数字类型", fieldKey)
	}

	return numValue, nil
}

// GetCheckboxValue 从Record获取复选框类型字段的值
func (r *Record) GetCheckboxValue(fieldKey string) (bool, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return false, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	boolValue, ok := SafeGetBool(value)
	if !ok {
		return false, fmt.Errorf("字段 %s 不是复选框类型", fieldKey)
	}

	return boolValue, nil
}

// GetDateValue 从Record获取日期类型字段的值
func (r *Record) GetDateValue(fieldKey string) (dateType string, value interface{}, err error) {
	fieldValue, exists := r.Values[fieldKey]
	if !exists {
		return "", nil, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	dateType, dateValue, ok := SafeGetDateValue(fieldValue)
	if !ok {
		return "", nil, fmt.Errorf("字段 %s 不是有效的日期类型", fieldKey)
	}

	return dateType, dateValue, nil
}

// GetDateString 从Record获取日期类型字段的字符串值（适用于DATE和DATETIME类型）
func (r *Record) GetDateString(fieldKey string) (string, error) {
	dateType, dateValue, err := r.GetDateValue(fieldKey)
	if err != nil {
		return "", err
	}

	if dateType != "DATE" && dateType != "DATETIME" {
		return "", fmt.Errorf("字段 %s 不是DATE或DATETIME类型", fieldKey)
	}

	dateStr, ok := SafeGetString(dateValue)
	if !ok {
		return "", fmt.Errorf("字段 %s 的日期值格式无效", fieldKey)
	}

	return dateStr, nil
}

// GetDateRangeStrings 从Record获取日期范围类型字段的起止日期
func (r *Record) GetDateRangeStrings(fieldKey string) (start, end string, err error) {
	dateType, dateValue, err := r.GetDateValue(fieldKey)
	if err != nil {
		return "", "", err
	}

	if dateType != "DATE_RANGE" {
		return "", "", fmt.Errorf("字段 %s 不是DATE_RANGE类型", fieldKey)
	}

	dateSlice, ok := SafeGetStringSlice(dateValue)
	if !ok || len(dateSlice) != 2 {
		return "", "", fmt.Errorf("字段 %s 的日期范围值格式无效", fieldKey)
	}

	return dateSlice[0], dateSlice[1], nil
}

// GetImageIDs 从Record获取图片类型字段的图片ID列表
func (r *Record) GetImageIDs(fieldKey string) ([]string, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return nil, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	imageIDs, ok := SafeGetStringSlice(value)
	if !ok {
		return nil, fmt.Errorf("字段 %s 不是图片类型", fieldKey)
	}

	return imageIDs, nil
}

// GetAttachmentFiles 从Record获取文件类型字段的文件信息列表
func (r *Record) GetAttachmentFiles(fieldKey string) ([]CellAttachmentFile, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return nil, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	// 处理文件附件格式
	files := []CellAttachmentFile{}

	// 尝试获取[]map[string]interface{}
	mapsSlice, ok := SafeGetMapSlice(value)
	if ok {
		for _, fileMap := range mapsSlice {
			fileID, _ := SafeGetString(fileMap["file_id"])
			fileName, _ := SafeGetString(fileMap["file_name"])
			files = append(files, CellAttachmentFile{
				FileID:   fileID,
				FileName: fileName,
			})
		}
		return files, nil
	}

	// 尝试获取[]interface{}，可能包含map[string]interface{}
	interfaceSlice, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("字段 %s 不是文件类型", fieldKey)
	}

	for _, item := range interfaceSlice {
		fileMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		fileID, _ := SafeGetString(fileMap["file_id"])
		fileName, _ := SafeGetString(fileMap["file_name"])
		files = append(files, CellAttachmentFile{
			FileID:   fileID,
			FileName: fileName,
		})
	}

	return files, nil
}

// GetUserIDs 从Record获取成员类型字段的用户ID列表
func (r *Record) GetUserIDs(fieldKey string) ([]string, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return nil, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	userIDs, ok := SafeGetStringSlice(value)
	if !ok {
		return nil, fmt.Errorf("字段 %s 不是成员类型", fieldKey)
	}

	return userIDs, nil
}

// GetURLValue 从Record获取超链接类型字段的链接信息
func (r *Record) GetURLValue(fieldKey string) (url, text string, err error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return "", "", fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	urlMap, ok := SafeGetMap(value)
	if !ok {
		return "", "", fmt.Errorf("字段 %s 不是超链接类型", fieldKey)
	}

	url, _ = SafeGetString(urlMap["url"])
	text, _ = SafeGetString(urlMap["text"])

	return url, text, nil
}

// GetSingleSelectValue 从Record获取单选类型字段的选项ID
func (r *Record) GetSingleSelectValue(fieldKey string) (string, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return "", fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	optionID, ok := SafeGetString(value)
	if !ok {
		return "", fmt.Errorf("字段 %s 不是单选类型", fieldKey)
	}

	return optionID, nil
}

// GetSelectValues 从Record获取多选类型字段的选项ID列表
func (r *Record) GetSelectValues(fieldKey string) ([]string, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return nil, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	optionIDs, ok := SafeGetStringSlice(value)
	if !ok {
		return nil, fmt.Errorf("字段 %s 不是多选类型", fieldKey)
	}

	return optionIDs, nil
}

// GetReferenceIDs 从Record获取关联类型字段的记录ID列表
func (r *Record) GetReferenceIDs(fieldKey string) ([]string, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return nil, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	recordIDs, ok := SafeGetStringSlice(value)
	if !ok {
		return nil, fmt.Errorf("字段 %s 不是关联类型", fieldKey)
	}

	return recordIDs, nil
}

// GetLocationValue 从Record获取地理位置类型字段的位置信息
func (r *Record) GetLocationValue(fieldKey string) (address string, latitude, longitude float64, err error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return "", 0, 0, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	locMap, ok := SafeGetMap(value)
	if !ok {
		return "", 0, 0, fmt.Errorf("字段 %s 不是地理位置类型", fieldKey)
	}

	address, _ = SafeGetString(locMap["address"])
	latitude, _ = SafeGetFloat64(locMap["latitude"])
	longitude, _ = SafeGetFloat64(locMap["longitude"])

	return address, latitude, longitude, nil
}

// GetCurrencyValue 从Record获取货币类型字段的数值
func (r *Record) GetCurrencyValue(fieldKey string) (float64, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return 0, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	amount, ok := SafeGetFloat64(value)
	if !ok {
		return 0, fmt.Errorf("字段 %s 不是货币类型", fieldKey)
	}

	return amount, nil
}

// GetWwGroupIDs 从Record获取群类型字段的群ID列表
func (r *Record) GetWwGroupIDs(fieldKey string) ([]string, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return nil, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	groupIDs, ok := SafeGetStringSlice(value)
	if !ok {
		return nil, fmt.Errorf("字段 %s 不是群类型", fieldKey)
	}

	return groupIDs, nil
}

// GetPercentageValue 从Record获取百分数类型字段的值
func (r *Record) GetPercentageValue(fieldKey string) (float64, error) {
	value, exists := r.Values[fieldKey]
	if !exists {
		return 0, fmt.Errorf("字段 %s 不存在", fieldKey)
	}

	percentage, ok := SafeGetFloat64(value)
	if !ok {
		return 0, fmt.Errorf("字段 %s 不是百分数类型", fieldKey)
	}

	return percentage, nil
}
