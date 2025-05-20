package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/xen0n/go-workwx/v2"
)

func cmdSmartsheetListFields(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)
	sheetID := c.String(flagSheetID)
	viewID := c.String(flagViewID)
	fieldIDs := c.StringSlice(flagFieldIDs)
	fieldTitles := c.StringSlice(flagFieldTitles)
	offset := c.Int(flagOffset)
	limit := c.Int(flagLimit)

	if docID == "" {
		return fmt.Errorf("必须指定doc-id")
	}

	if sheetID == "" {
		return fmt.Errorf("必须指定sheet-id")
	}

	app := cfg.MakeWorkwxApp()

	total, fields, err := app.ListSmartsheetFields(docID, sheetID, viewID, fieldIDs, fieldTitles, offset, limit)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("字段列表 (总计 %d 个, 本次返回 %d 个):\n", total, len(fields))

	for i, field := range fields {
		fmt.Printf("  [%d] ID: %s\n", i+1, field.FieldID)
		fmt.Printf("      标题: %s\n", field.FieldTitle)
		fmt.Printf("      类型: %s\n", field.FieldType)

		// 根据字段类型展示特定属性
		switch field.FieldType {
		case workwx.FieldTypeNumber:
			if field.PropertyNumber != nil {
				fmt.Printf("      小数位数: %d\n", field.PropertyNumber.DecimalPlaces)
				fmt.Printf("      使用千分位分隔符: %v\n", field.PropertyNumber.UseSeparate)
			}
		case workwx.FieldTypeSelect:
			if field.PropertySelect != nil {
				fmt.Printf("      允许快速添加选项: %v\n", field.PropertySelect.IsQuickAdd)
				if len(field.PropertySelect.Options) > 0 {
					fmt.Printf("      选项数量: %d\n", len(field.PropertySelect.Options))
					for j, opt := range field.PropertySelect.Options {
						if j < 3 { // 只展示前3个选项
							fmt.Printf("        - %s (%s)\n", opt.Text, opt.ID)
						}
					}
					if len(field.PropertySelect.Options) > 3 {
						fmt.Printf("        ... 等%d个选项\n", len(field.PropertySelect.Options)-3)
					}
				}
			}
		case workwx.FieldTypeSingleSelect:
			if field.PropertySingleSelect != nil {
				fmt.Printf("      允许快速添加选项: %v\n", field.PropertySingleSelect.IsQuickAdd)
				if len(field.PropertySingleSelect.Options) > 0 {
					fmt.Printf("      选项数量: %d\n", len(field.PropertySingleSelect.Options))
				}
			}
		case workwx.FieldTypeUser:
			if field.PropertyUser != nil {
				fmt.Printf("      允许添加多个人员: %v\n", field.PropertyUser.IsMultiple)
				fmt.Printf("      添加人员时通知用户: %v\n", field.PropertyUser.IsNotified)
			}
		case workwx.FieldTypeDateTime:
			if field.PropertyDateTime != nil {
				fmt.Printf("      日期格式: %s\n", field.PropertyDateTime.Format)
				fmt.Printf("      自动填充: %v\n", field.PropertyDateTime.AutoFill)
			}
		}
	}

	return nil
}
