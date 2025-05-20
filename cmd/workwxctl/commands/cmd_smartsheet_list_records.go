package commands

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/xen0n/go-workwx/v2"
)

func cmdSmartsheetListRecords(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)
	sheetID := c.String(flagSheetID)
	viewID := c.String(flagViewID)
	recordIDs := c.StringSlice(flagRecordIDs)
	keyType := workwx.CellValueKeyType(c.String(flagKeyType))
	fieldTitles := c.StringSlice(flagFieldTitles)
	fieldIDs := c.StringSlice(flagFieldIDs)
	offset := uint32(c.Uint(flagOffset))
	limit := uint32(c.Uint(flagLimit))

	if docID == "" {
		return fmt.Errorf("必须指定doc-id")
	}

	if sheetID == "" {
		return fmt.Errorf("必须指定sheet-id")
	}

	// 如果keyType为空，默认使用字段标题作为key
	if keyType == "" {
		keyType = workwx.CellValueKeyTypeFieldTitle
	}

	app := cfg.MakeWorkwxApp()

	// 对于排序参数，从命令行参数中解析
	sort := make([]workwx.Sort, 0)
	sortFields := c.StringSlice(flagSortFields)
	sortDescStr := c.StringSlice(flagSortDesc)

	// 将sortDesc字符串数组转为bool数组
	for i, field := range sortFields {
		desc := false
		// 如果有对应的desc参数且值为"true"，则将desc设为true
		if i < len(sortDescStr) && (sortDescStr[i] == "true" || sortDescStr[i] == "1") {
			desc = true
		}
		sort = append(sort, workwx.Sort{
			FieldTitle: field,
			Desc:       desc,
		})
	}

	total, hasMore, next, records, err := app.ListSmartsheetRecords(docID, sheetID, viewID, recordIDs, keyType, fieldTitles, fieldIDs, sort, offset, limit)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("记录列表 (总计 %d 条, 本次返回 %d 条):\n", total, len(records))
	if hasMore {
		fmt.Printf("还有更多数据，下一页偏移量: %d\n", next)
	}

	for i, record := range records {
		fmt.Printf("  [%d] 记录ID: %s\n", i+1, record.RecordID)
		fmt.Printf("      创建时间: %s\n", record.CreateTime)
		fmt.Printf("      更新时间: %s\n", record.UpdateTime)
		fmt.Printf("      创建者: %s\n", record.CreatorName)
		fmt.Printf("      最后编辑者: %s\n", record.UpdaterName)

		// 显示记录值
		if len(record.Values) > 0 {
			fmt.Printf("      字段值:\n")
			// 将values转为JSON字符串并格式化打印，便于查看
			valuesJSON, err := json.MarshalIndent(record.Values, "        ", "  ")
			if err != nil {
				fmt.Printf("        无法格式化显示字段值: %v\n", err)
			} else {
				fmt.Printf("        %s\n", string(valuesJSON))
			}
		}
	}

	return nil
}
