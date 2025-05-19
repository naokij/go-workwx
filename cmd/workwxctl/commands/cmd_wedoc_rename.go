package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdWedocRename(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)
	formID := c.String(flagFormID)
	newName := c.String(flagNewName)

	if docID == "" && formID == "" {
		return fmt.Errorf("必须指定doc-id或form-id其中之一")
	}

	if newName == "" {
		return fmt.Errorf("必须指定新文档名称 (new-name)")
	}

	app := cfg.MakeWorkwxApp()

	err := app.RenameWedoc(docID, formID, newName)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Println("文档重命名成功")
	return nil
}
