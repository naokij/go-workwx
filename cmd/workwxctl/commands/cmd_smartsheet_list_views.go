package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdSmartsheetListViews(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)
	sheetID := c.String(flagSheetID)
	viewIDs := c.StringSlice(flagViewIDs)
	offset := uint32(c.Uint(flagOffset))
	limit := uint32(c.Uint(flagLimit))

	if docID == "" {
		return fmt.Errorf("必须指定doc-id")
	}

	if sheetID == "" {
		return fmt.Errorf("必须指定sheet-id")
	}

	app := cfg.MakeWorkwxApp()

	total, hasMore, next, views, err := app.ListSmartsheetViews(docID, sheetID, viewIDs, offset, limit)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("视图列表 (总计 %d 个, 本次返回 %d 个):\n", total, len(views))
	if hasMore {
		fmt.Printf("还有更多数据，下一页偏移量: %d\n", next)
	}

	for i, view := range views {
		fmt.Printf("  [%d] ID: %s\n", i+1, view.ViewID)
		fmt.Printf("      标题: %s\n", view.ViewTitle)
		fmt.Printf("      类型: %s\n", view.ViewType)
		if view.Property != nil {
			fmt.Printf("      自动排序: %v\n", view.Property.AutoSort)
			fmt.Printf("      冻结列数: %d\n", view.Property.FrozenFieldCount)
			fmt.Printf("      字段统计已启用: %v\n", view.Property.IsFieldStatEnabled)
		}
	}

	return nil
}
