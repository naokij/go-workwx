package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/xen0n/go-workwx/v2"
)

// DocType 整数常量定义
const (
	DocTypeDoc        = 3  // 文档
	DocTypeSheet      = 4  // 表格
	DocTypeSmartsheet = 10 // 智能表格
)

func cmdWedocCreate(c *cli.Context) error {
	cfg := mustGetConfig(c)

	spaceID := c.String(flagSpaceID)
	fatherID := c.String(flagFatherID)
	docName := c.String(flagDocName)
	adminUsers := c.StringSlice(flagAdminUsers)

	// 获取文档类型
	docTypeStr := c.String(flagDocType)
	var docTypeInt int
	switch docTypeStr {
	case "doc":
		docTypeInt = DocTypeDoc
	case "sheet":
		docTypeInt = DocTypeSheet
	case "smartsheet":
		docTypeInt = DocTypeSmartsheet
	default:
		return fmt.Errorf("不支持的文档类型: %s，支持的类型: doc, sheet, smartsheet", docTypeStr)
	}

	app := cfg.MakeWorkwxApp()

	_ = docTypeInt
	var docType workwx.DocType
	docType = workwx.DocType(docTypeInt)

	url, docID, err := app.CreateWedoc(spaceID, fatherID, docType, docName, adminUsers)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("文档创建成功:\nURL: %s\nDocID: %s\n", url, docID)
	return nil
}
