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

// ListSmartsheetFields 获取智能表格字段列表
// docID: 文档的docid
// sheetID: 表格ID
// viewID: 可选，视图ID
// fieldIDs: 可选，由字段ID组成的数组
// fieldTitles: 可选，由字段标题组成的数组
// offset: 可选，偏移量，初始值为0
// limit: 可选，分页大小，不填或0时，如果总数大于1000，一次性返回1000个字段，当总数小于1000时，返回全部字段；最大值为1000
func (c *WorkwxApp) ListSmartsheetFields(docID, sheetID string, viewID string, fieldIDs, fieldTitles []string, offset, limit int) (total int, fields []Field, err error) {
	req := reqListSmartsheetFields{
		DocID:       docID,
		SheetID:     sheetID,
		ViewID:      viewID,
		FieldIDs:    fieldIDs,
		FieldTitles: fieldTitles,
		Offset:      offset,
		Limit:       limit,
	}

	resp, err := c.execWedocSmartsheetGetFields(req)
	if err != nil {
		return 0, nil, err
	}

	return resp.Total, resp.Fields, nil
}

// ListSmartsheetRecords 获取智能表格记录列表
// docID: 文档的docid
// sheetID: Smartsheet子表ID
// viewID: 可选，视图ID
// recordIDs: 可选，由记录ID组成的数组
// keyType: 可选，返回记录中单元格的key类型，默认为CellValueKeyTypeFieldTitle
// fieldTitles: 可选，返回指定列，由字段标题组成的数组，keyType为CellValueKeyTypeFieldTitle时有效
// fieldIDs: 可选，返回指定列，由字段ID组成的数组，keyType为CellValueKeyTypeFieldID时有效
// sort: 可选，对返回记录进行排序
// offset: 可选，偏移量，初始值为0
// limit: 可选，分页大小，不填或0时，如果总数大于1000，一次性返回1000行记录，当总数小于1000时，返回全部记录；最大值为1000
func (c *WorkwxApp) ListSmartsheetRecords(docID, sheetID string, viewID string, recordIDs []string, keyType CellValueKeyType, fieldTitles, fieldIDs []string, sort []Sort, offset, limit uint32) (total uint32, hasMore bool, next uint32, records []Record, err error) {
	req := reqSmartsheetGetRecords{
		DocID:       docID,
		SheetID:     sheetID,
		ViewID:      viewID,
		RecordIDs:   recordIDs,
		KeyType:     keyType,
		FieldTitles: fieldTitles,
		FieldIDs:    fieldIDs,
		Sort:        sort,
		Offset:      offset,
		Limit:       limit,
	}

	resp, err := c.execWedocSmartsheetGetRecords(req)
	if err != nil {
		return 0, false, 0, nil, err
	}

	return resp.Total, resp.HasMore, resp.Next, resp.Records, nil
}

// AddSmartsheetRecords 添加智能表格记录
// docID: 文档的docid
// sheetID: Smartsheet子表ID
// keyType: 可选，返回记录中单元格的key类型，默认为CellValueKeyTypeFieldTitle
// records: 需要添加的记录内容，map中的key为字段标题或字段ID，value类型根据字段类型不同而异
// 参考文档: https://developer.work.weixin.qq.com/document/path/100224
func (c *WorkwxApp) AddSmartsheetRecords(docID, sheetID string, keyType CellValueKeyType, records []map[string]interface{}) ([]Record, error) {
	addRecords := make([]AddRecord, 0, len(records))
	for _, r := range records {
		addRecords = append(addRecords, AddRecord{
			Values: r,
		})
	}

	req := reqSmartsheetAddRecords{
		DocID:   docID,
		SheetID: sheetID,
		KeyType: keyType,
		Records: addRecords,
	}

	resp, err := c.execWedocSmartsheetAddRecords(req)
	if err != nil {
		return nil, err
	}

	return resp.Records, nil
}

// NewTextCellValue 创建文本类型的单元格值
func NewTextCellValue(text string) string {
	return text
}

// NewNumberCellValue 创建数字类型的单元格值
func NewNumberCellValue(number float64) float64 {
	return number
}

// NewCheckboxCellValue 创建复选框类型的单元格值
func NewCheckboxCellValue(checked bool) bool {
	return checked
}

// NewDateCellValue 创建日期类型的单元格值
func NewDateCellValue(dateStr string) map[string]interface{} {
	return map[string]interface{}{
		"type":  "DATE",
		"value": dateStr,
	}
}

// NewDateTimeCellValue 创建日期时间类型的单元格值
func NewDateTimeCellValue(dateTimeStr string) map[string]interface{} {
	return map[string]interface{}{
		"type":  "DATETIME",
		"value": dateTimeStr,
	}
}

// NewDateRangeCellValue 创建日期范围类型的单元格值
func NewDateRangeCellValue(startDate, endDate string) map[string]interface{} {
	return map[string]interface{}{
		"type":  "DATE_RANGE",
		"value": []string{startDate, endDate},
	}
}

// NewImageCellValue 创建图片类型的单元格值
func NewImageCellValue(imageIDs ...string) []string {
	return imageIDs
}

// NewAttachmentFile 创建文件信息
func NewAttachmentFile(fileID, fileName string) map[string]string {
	return map[string]string{
		"file_id":   fileID,
		"file_name": fileName,
	}
}

// NewAttachmentCellValue 创建文件类型的单元格值
func NewAttachmentCellValue(files ...map[string]string) []map[string]string {
	return files
}

// NewUserCellValue 创建成员类型的单元格值
func NewUserCellValue(userIDs ...string) []string {
	return userIDs
}

// NewURLCellValue 创建超链接类型的单元格值
func NewURLCellValue(url, text string) map[string]string {
	return map[string]string{
		"url":  url,
		"text": text,
	}
}

// NewSingleSelectCellValue 创建单选类型的单元格值
func NewSingleSelectCellValue(optionID string) string {
	return optionID
}

// NewSelectCellValue 创建多选类型的单元格值
func NewSelectCellValue(optionIDs ...string) []string {
	return optionIDs
}

// NewReferenceCellValue 创建关联类型的单元格值
func NewReferenceCellValue(recordIDs ...string) []string {
	return recordIDs
}

// NewLocationCellValue 创建地理位置类型的单元格值
func NewLocationCellValue(address string, latitude, longitude float64) map[string]interface{} {
	return map[string]interface{}{
		"address":   address,
		"latitude":  latitude,
		"longitude": longitude,
	}
}

// NewCurrencyCellValue 创建货币类型的单元格值
func NewCurrencyCellValue(amount float64) float64 {
	return amount
}

// NewWwGroupCellValue 创建群类型的单元格值
func NewWwGroupCellValue(groupIDs ...string) []string {
	return groupIDs
}

// NewPercentageCellValue 创建百分数类型的单元格值
func NewPercentageCellValue(percentage float64) float64 {
	return percentage
}
