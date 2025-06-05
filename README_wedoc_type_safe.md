# 企业微信智能表格字段类型安全访问

本文档介绍了企业微信智能表格API中新增的类型安全字段访问功能。这些功能旨在解决使用`map[string]interface{}`访问字段值时可能出现的类型安全问题。

## 功能特点

1. 提供了一系列安全获取各种类型值的工具函数
2. 为`Record`结构体添加了类型安全的字段访问方法
3. 保持了原有API的兼容性，同时提供更安全的类型访问方式

## 安全工具函数

以下是一些核心的安全类型获取函数：

```go
// 安全地从interface{}获取字符串值
func SafeGetString(value interface{}) (string, bool)

// 安全地从interface{}获取float64值
func SafeGetFloat64(value interface{}) (float64, bool)

// 安全地从interface{}获取布尔值
func SafeGetBool(value interface{}) (bool, bool)

// 安全地从interface{}获取字符串切片
func SafeGetStringSlice(value interface{}) ([]string, bool)

// 安全地从interface{}获取map
func SafeGetMap(value interface{}) (map[string]interface{}, bool)

// 安全地从interface{}获取map切片
func SafeGetMapSlice(value interface{}) ([]map[string]interface{}, bool)
```

## Record结构体的类型安全方法

`Record`结构体新增了以下类型安全的访问方法：

```go
// 获取文本类型字段的值
func (r *Record) GetTextValue(fieldKey string) (string, error)

// 获取数字类型字段的值
func (r *Record) GetNumberValue(fieldKey string) (float64, error)

// 获取复选框类型字段的值
func (r *Record) GetCheckboxValue(fieldKey string) (bool, error)

// 获取日期类型字段的值
func (r *Record) GetDateValue(fieldKey string) (dateType string, value interface{}, err error)

// 获取日期类型字段的字符串值（适用于DATE和DATETIME类型）
func (r *Record) GetDateString(fieldKey string) (string, error)

// 获取日期范围类型字段的起止日期
func (r *Record) GetDateRangeStrings(fieldKey string) (start, end string, err error)

// 获取图片类型字段的图片ID列表
func (r *Record) GetImageIDs(fieldKey string) ([]string, error)

// 获取文件类型字段的文件信息列表
func (r *Record) GetAttachmentFiles(fieldKey string) ([]CellAttachmentFile, error)

// 获取成员类型字段的用户ID列表
func (r *Record) GetUserIDs(fieldKey string) ([]string, error)

// 获取超链接类型字段的链接信息
func (r *Record) GetURLValue(fieldKey string) (url, text string, err error)

// 获取单选类型字段的选项ID
func (r *Record) GetSingleSelectValue(fieldKey string) (string, error)

// 获取多选类型字段的选项ID列表
func (r *Record) GetSelectValues(fieldKey string) ([]string, error)

// 获取关联类型字段的记录ID列表
func (r *Record) GetReferenceIDs(fieldKey string) ([]string, error)

// 获取地理位置类型字段的位置信息
func (r *Record) GetLocationValue(fieldKey string) (address string, latitude, longitude float64, err error)

// 获取货币类型字段的数值
func (r *Record) GetCurrencyValue(fieldKey string) (float64, error)

// 获取群类型字段的群ID列表
func (r *Record) GetWwGroupIDs(fieldKey string) ([]string, error)

// 获取百分数类型字段的值
func (r *Record) GetPercentageValue(fieldKey string) (float64, error)
```

## 使用示例

以下是使用类型安全方法访问字段的示例：

```go
// 获取记录列表
_, _, _, records, err := app.ListSmartsheetRecords(
    docID,
    sheetID,
    "",  // 视图ID，可选
    nil, // 记录ID列表，可选
    workwx.CellValueKeyTypeFieldTitle, // 使用字段标题作为key
    nil, // 字段标题列表，可选
    nil, // 字段ID列表，可选
    nil, // 排序，可选
    0,   // 偏移量
    10,  // 每页记录数
)

if err != nil {
    fmt.Printf("获取记录失败: %v\n", err)
    return
}

// 使用类型安全的方法处理记录
for _, record := range records {
    // 获取文本字段
    if textValue, err := record.GetTextValue("标题"); err == nil {
        fmt.Printf("标题: %s\n", textValue)
    }
    
    // 获取数字字段
    if numberValue, err := record.GetNumberValue("金额"); err == nil {
        fmt.Printf("金额: %.2f\n", numberValue)
    }
    
    // 获取复选框字段
    if isCompleted, err := record.GetCheckboxValue("是否完成"); err == nil {
        fmt.Printf("是否完成: %v\n", isCompleted)
    }
    
    // 获取日期字段
    if dateStr, err := record.GetDateString("截止日期"); err == nil {
        fmt.Printf("截止日期: %s\n", dateStr)
    }
    
    // 获取日期范围字段
    if start, end, err := record.GetDateRangeStrings("日期范围"); err == nil {
        fmt.Printf("日期范围: %s 至 %s\n", start, end)
    }
}
```

## 添加记录示例

添加记录时使用各种类型的创建函数：

```go
newRecords, err := app.AddSmartsheetRecords(
    docID,
    sheetID,
    workwx.CellValueKeyTypeFieldTitle,
    []map[string]interface{}{
        {
            "标题":    workwx.NewTextCellValue("新任务"),
            "金额":    workwx.NewNumberCellValue(1000.50),
            "是否完成":  workwx.NewCheckboxCellValue(false),
            "截止日期":  workwx.NewDateCellValue("2023-06-30"),
            "时间范围":  workwx.NewDateRangeCellValue("2023-06-01", "2023-06-30"),
            "相关链接":  workwx.NewURLCellValue("https://example.com", "参考文档"),
            "负责人":   workwx.NewUserCellValue("zhangsan", "lisi"),
            "地点":    workwx.NewLocationCellValue("北京市海淀区", 39.9, 116.3),
            "进度":    workwx.NewPercentageCellValue(0.35), // 35%
            "状态":    workwx.NewSingleSelectCellValue("option1"),
            "标签":    workwx.NewSelectCellValue("tag1", "tag2"),
        },
    },
)
```

## 优势

1. 类型安全 - 避免运行时类型断言错误
2. 代码可读性高 - 明确字段的类型和预期值
3. IDE支持更好 - 提供自动完成和类型提示
4. 错误处理更明确 - 提供具体的错误信息而不是运行时panic 