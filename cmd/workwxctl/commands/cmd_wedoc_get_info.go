package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdWedocGetInfo(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)

	if docID == "" {
		return fmt.Errorf("必须指定doc-id")
	}

	app := cfg.MakeWorkwxApp()

	info, err := app.GetWedocBaseInfo(docID)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("文档信息:\n")
	fmt.Printf("  DocID: %s\n", info.DocID)
	fmt.Printf("  DocName: %s\n", info.DocName)
	fmt.Printf("  CreateTime: %d\n", info.CreateTime)
	fmt.Printf("  ModifyTime: %d\n", info.ModifyTime)
	fmt.Printf("  DocType: %+v\n", info.DocType)

	return nil
}
