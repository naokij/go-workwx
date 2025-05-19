package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdWedocDelete(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)
	formID := c.String(flagFormID)

	if docID == "" && formID == "" {
		return fmt.Errorf("必须指定doc-id或form-id其中之一")
	}

	app := cfg.MakeWorkwxApp()

	err := app.DeleteWedoc(docID, formID)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Println("文档删除成功")
	return nil
}
