package workwx

// CreateWedoc 创建企业微信文档
// spaceID: 空间spaceid，若指定spaceid，则fatherid也要同时指定
// fatherID: 父目录fileid, 在根目录时为空间spaceid
// docType: 文档类型, 3:文档(DocTypeDoc) 4:表格(DocTypeSheet) 10:智能表格(DocTypeSmartsheet)
// docName: 文档名字（注意：文件名最多填255个字符, 超过255个字符会被截断）
// adminUsers: 文档管理员userid列表
func (c *WorkwxApp) CreateWedoc(spaceID, fatherID string, docType DocType, docName string, adminUsers []string) (url, docID string, err error) {
	req := reqCreateDoc{
		SpaceID:    spaceID,
		FatherID:   fatherID,
		DocType:    docType,
		DocName:    docName,
		AdminUsers: adminUsers,
	}

	resp, err := c.execWedocCreatDoc(req)
	if err != nil {
		return "", "", err
	}

	return resp.URL, resp.DocID, nil
}

// RenameWedoc 重命名企业微信文档
// docID: 文档docid，仅可修改应用自己创建的文档，docID和formID只能填其中一个
// formID: 收集表id，仅可修改应用自己创建的收集表，docID和formID只能填其中一个
// newName: 重命名后的文档名（注意：文档名最多填255个字符，英文算1个，汉字算2个，超过255个字符会被截断）
func (c *WorkwxApp) RenameWedoc(docID, formID, newName string) error {
	req := reqRenameDoc{
		DocID:   docID,
		FormID:  formID,
		NewName: newName,
	}

	_, err := c.execWedocRenameDoc(req)
	return err
}

// DeleteWedoc 删除企业微信文档
// docID: 文档docid，仅可删除应用自己创建的文档，docID和formID只能填其中一个
// formID: 收集表id，仅可删除应用自己创建的收集表，docID和formID只能填其中一个
func (c *WorkwxApp) DeleteWedoc(docID, formID string) error {
	req := reqDelDoc{
		DocID:  docID,
		FormID: formID,
	}

	_, err := c.execWedocDelDoc(req)
	return err
}

// GetWedocBaseInfo 获取企业微信文档基础信息
// docID: 文档docid
func (c *WorkwxApp) GetWedocBaseInfo(docID string) (DocBaseInfo, error) {
	req := reqGetDocBaseInfo{
		DocID: docID,
	}

	resp, err := c.execWedocGetDocBaseInfo(req)
	if err != nil {
		return DocBaseInfo{}, err
	}

	return resp.DocBaseInfo, nil
}

// ShareWedoc 分享企业微信文档
// docID: 文档id，docID和formID只能填其中一个
// formID: 收集表id，docID和formID只能填其中一个
func (c *WorkwxApp) ShareWedoc(docID, formID string) (string, error) {
	req := reqDocShare{
		DocID:  docID,
		FormID: formID,
	}

	resp, err := c.execWedocDocShare(req)
	if err != nil {
		return "", err
	}

	return resp.ShareURL, nil
}

// ListSmartsheetSheets 获取文档中的子表列表
// docID: 文档的docid
// sheetID: 指定子表ID查询（可选）
// needAllTypeSheet: 是否获取所有类型子表
func (c *WorkwxApp) ListSmartsheetSheets(docID string, sheetID string, needAllTypeSheet bool) ([]SheetInfo, error) {
	req := reqGetSmartsheet{
		DocID:            docID,
		SheetID:          sheetID,
		NeedAllTypeSheet: needAllTypeSheet,
	}

	resp, err := c.execWedocSmartsheetGetSheet(req)
	if err != nil {
		return nil, err
	}

	return resp.SheetList, nil
}

// ListSmartsheetViews 获取智能表格视图列表
// docID: 文档的docid
// sheetID: Smartsheet子表ID
// viewIDs: 可选，需要查询的视图ID数组
// offset: 可选，偏移量，初始值为0
// limit: 可选，分页大小，不填或0时，如果总数大于1000，一次性返回1000个视图，否则返回全部视图；最大值为1000
func (c *WorkwxApp) ListSmartsheetViews(docID, sheetID string, viewIDs []string, offset, limit uint32) (total uint32, hasMore bool, next uint32, views []View, err error) {
	req := reqListSmartsheetViews{
		DocID:   docID,
		SheetID: sheetID,
		ViewIDs: viewIDs,
		Offset:  offset,
		Limit:   limit,
	}

	resp, err := c.execWedocSmartsheetGetViews(req)
	if err != nil {
		return 0, false, 0, nil, err
	}

	return resp.Total, resp.HasMore, resp.Next, resp.Views, nil
}
