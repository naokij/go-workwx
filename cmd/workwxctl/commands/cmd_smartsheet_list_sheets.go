package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdSmartsheetListSheets(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)
	sheetID := c.String(flagSheetID)
	needAllTypeSheet := c.Bool("need-all-type-sheet")

	if docID == "" {
		return fmt.Errorf("必须指定doc-id")
	}

	app := cfg.MakeWorkwxApp()

	sheetList, err := app.ListSmartsheetSheets(docID, sheetID, needAllTypeSheet)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("智能表格子表列表 (共 %d 个):\n", len(sheetList))
	for i, sheet := range sheetList {
		fmt.Printf("  [%d] ID: %s\n", i+1, sheet.SheetID)
		fmt.Printf("      标题: %s\n", sheet.Title)
		fmt.Printf("      类型: %s\n", sheet.SheetType)
		fmt.Printf("      是否可见: %v\n", sheet.IsVisible)
	}

	return nil
}
