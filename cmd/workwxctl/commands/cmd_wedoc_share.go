package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdWedocShare(c *cli.Context) error {
	cfg := mustGetConfig(c)

	docID := c.String(flagDocID)
	formID := c.String(flagFormID)

	if docID == "" && formID == "" {
		return fmt.Errorf("必须指定doc-id或form-id其中之一")
	}

	app := cfg.MakeWorkwxApp()

	shareURL, err := app.ShareWedoc(docID, formID)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("文档分享链接: %s\n", shareURL)
	return nil
}
