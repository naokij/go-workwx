package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/xen0n/go-workwx/v2"
)

func cmdSmartsheetAddRecords(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)
	sheetID := c.String(flagSheetID)
	keyType := workwx.CellValueKeyType(c.String(flagKeyType))
	recordsFile := c.String("records-file")

	if docID == "" {
		return fmt.Errorf("必须指定doc-id")
	}

	if sheetID == "" {
		return fmt.Errorf("必须指定sheet-id")
	}

	if recordsFile == "" {
		return fmt.Errorf("必须指定records-file")
	}

	// 如果keyType为空，默认使用字段标题作为key
	if keyType == "" {
		keyType = workwx.CellValueKeyTypeFieldTitle
	}

	// 读取记录文件
	file, err := os.Open(recordsFile)
	if err != nil {
		return fmt.Errorf("无法打开记录文件: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("无法读取文件内容: %v", err)
	}

	// 解析记录JSON
	var records []map[string]interface{}
	if err := json.Unmarshal(content, &records); err != nil {
		return fmt.Errorf("无法解析记录JSON: %v", err)
	}

	if len(records) == 0 {
		return fmt.Errorf("没有找到记录数据")
	}

	app := cfg.MakeWorkwxApp()

	addedRecords, err := app.AddSmartsheetRecords(docID, sheetID, keyType, records)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("成功添加 %d 条记录:\n", len(addedRecords))
	for i, record := range addedRecords {
		fmt.Printf("  [%d] 记录ID: %s\n", i+1, record.RecordID)
		fmt.Printf("      创建时间: %s\n", record.CreateTime)
		fmt.Printf("      更新时间: %s\n", record.UpdateTime)
		fmt.Printf("      创建者: %s\n", record.CreatorName)
		fmt.Printf("      最后编辑者: %s\n", record.UpdaterName)
	}

	return nil
}
