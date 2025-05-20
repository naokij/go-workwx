# 文档接口

## Models

### `reqCreateDoc` 新建文档请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`SpaceID`|`spaceid,omitempty`|`string`|空间spaceid。若指定spaceid，则fatherid也要同时指定
`FatherID`|`fatherid`|`string`|	父目录fileid, 在根目录时为空间spaceid
`DocType`|`doc_type`|`DocType`|文档类型, 3:文档 4:表格 10:智能表格
`DocName`|`doc_name`|`string`|文档名字（注意：文件名最多填255个字符, 超过255个字符会被截断）
`AdminUsers`|`admin_users`|`[]string`|文档管理员userid

```go
// DocType 文档类型
type DocType uint32

const (
	DocTypeDoc        DocType = 3  // 文档
	DocTypeSheet      DocType = 4  // 表格
	DocTypeSmartSheet DocType = 10 // 智能表格
)
```

### `respCreateDoc` 新建文档返回

Name|JSON|Type|Doc
:---|:---|:---|:--
`respCommon`||-|通用响应
`URL`|`url`|`string`|新建文档的访问链接
`DocID`|`docid`|`string`|新建文档的docid

### `reqRenameDoc` 重命名文档请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`|文档docid（docid、formid只能填其中一个），仅可修改应用自己创建的文档
`FormID`|`formid`|`string`|收集表id（docid、formid只能填其中一个），仅可修改应用自己创建的收集表
`NewName`|`new_name`|`string`|重命名后的文档名（注意：文档名最多填255个字符，英文算1个，汉字算2个，超过255个字符会被截断）

### `reqDelDoc` 删除文档请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid,omitempty`|`string`|文档docid（docid、formid只能填其中一个），仅可删除应用自己创建的文档
`FormID`|`formid,omitempty`|`string`|收集表id（docid、formid只能填其中一个），仅可删除应用自己创建的收集表

### `reqGetDocBaseInfo` 获取文档基础信息请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`|文档docid

### `respGetDocBaseInfo` 获取文档基础信息返回

Name|JSON|Type|Doc
:---|:---|:---|:--
`respCommon`||-|通用响应
`DocBaseInfo`|`doc_base_info`|`DocBaseInfo`|文档基础信息


### `DocBaseInfo` 文档基础信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`|文档docid
`DocName`|`doc_name`|`string`|文档名字
`CreateTime`|`create_time`|`uint64`|文档创建时间
`ModifyTime`|`modify_time`|`uint64`|文档最后修改时间
`DocType`|`doc_type`|`DocType`|文档类型

### `reqDocShare` 分享文档请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid,omitempty`|`string`|文档id（docid、formid只能填其中一个）
`FormID`|`formid,omitempty`|`string`|收集表id（docid、formid只能填其中一个）

### `respDocShare` 分享文档返回

Name|JSON|Type|Doc
:---|:---|:---|:--
`respCommon`||-|通用响应
`ShareURL`|`share_url`|`string`|文档分享链接

### `reqGetSmartsheet` 查询子表请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`|文档的docid
`SheetID`|`sheet_id`|`string`|指定子表ID查询
`NeedAllTypeSheet`|`need_all_type_sheet`|`bool`|获取所有类型子表

### `respGetSmartsheet` 查询子表返回

Name|JSON|Type|Doc
:---|:---|:---|:--
`respCommon`||-|通用响应
`SheetList`|`sheet_list`|`[]SheetInfo`|智能表信息列表

### `SheetInfo` 子表信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`SheetID`|`sheet_id`|`string`|子表id
`Title`|`title`|`string`|子表名称
`IsVisible`|`is_visible`|`bool`|子表是否可见
`SheetType`|`type`|`string`|子表类型。"dashboard" 仪表盘。"external" 说明页，"smartsheet" 智能表

### `reqListSmartsheetViews` 获取智能表格视图列表请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`|文档的docid
`SheetID`|`sheet_id`|`string`|Smartsheet子表ID
`ViewIDs`|`view_ids`|`[]string`|需要查询的视图ID数组
`Offset`|`offset`|`uint32`|偏移量，初始值为0
`Limit`|`limit`|`uint32`|分页大小，不填或0时，如果总数大于1000，一次性返回1000个视图，否则返回全部视图；最大值为1000

### `respListSmartsheetViews` 获取智能表格视图列表响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`respCommon`||-|通用响应
`Total`|`total`|`uint32`|符合筛选条件的视图总数
`HasMore`|`has_more`|`bool`|是否还有更多项
`Next`|`next`|`uint32`|下次下一个搜索结果的偏移量
`Views`|`views`|`[]View`|视图数据

### `View` 视图信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`ViewID`|`view_id`|`string`|视图ID
`ViewTitle`|`view_title`|`string`|视图标题
`ViewType`|`view_type`|`ViewType`|视图类型
`Property`|`property`|`*ViewProperty`|视图属性

```go
// ViewType 视图类型
type ViewType string

const (
	ViewTypeUnknown ViewType = "VEW_UNKNOWN"       // 未知类型视图
	ViewTypeGrid    ViewType = "VIEW_TYPE_GRID"    // 网格视图
	ViewTypeKanban  ViewType = "VIEW_TYPE_KANBAN"  // 看板视图
	ViewTypeGallery ViewType = "VIEW_TYPE_GALLERY" // 画册视图
	ViewTypeGantt   ViewType = "VIEW_TYPE_GANTT"   // 甘特视图
)
```

### `ViewProperty` 视图属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`AutoSort`|`auto_sort`|`bool`|记录变更后自动重新排序
`SortSpec`|`sort_spec`|`*SortSpec`|排序设置
`GroupSpec`|`group_spec`|`*GroupSpec`|分组设置
`FilterSpec`|`filter_spec`|`*FilterSpec`|过滤设置
`IsFieldStatEnabled`|`is_field_stat_enabled`|`bool`|是否使用数据统计
`FieldVisibility`|`field_visibility`|`map[string]bool`|字段可见性
`FrozenFieldCount`|`frozen_field_count`|`int32`|冻结列数量，从首列开始

### `SortSpec` 排序设置

Name|JSON|Type|Doc
:---|:---|:---|:--
`SortInfos`|`sort_infos`|`[]SortInfo`|参与排序的字段列表

### `SortInfo` 排序信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`|字段id
`Desc`|`desc`|`bool`|是否降序

### `GroupSpec` 分组设置

Name|JSON|Type|Doc
:---|:---|:---|:--
`Groups`|`groups`|`[]GroupInfo`|参与分组的字段列表

### `GroupInfo` 分组信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`|字段id
`Desc`|`desc`|`bool`|是否降序

### `FilterSpec` 过滤设置

Name|JSON|Type|Doc
:---|:---|:---|:--
`Conjunction`|`conjunction`|`string`|多个conditions之间是以and("CONJUNCTION_AND")还是or("CONJUNCTION_OR")进行组合
`Conditions`|`conditions`|`[]Condition`|判断条件

### `Condition` 过滤条件

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`|字段ID
`FieldType`|`field_type`|`string`|字段类型
`Operator`|`operator`|`string`|判断类型
`StringValue`|`string_value`|`*StringValue`|文本值
`NumberValue`|`number_value`|`*NumberValue`|数字值
`BoolValue`|`bool_value`|`*BoolValue`|布尔值
`UserValue`|`user_value`|`*UserValue`|用户值
`DateTimeValue`|`date_time_value`|`*FilterDataTimeValue`|日期时间值

### `StringValue` 字符串值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Value`|`value`|`[]string`|字符串值数组

### `NumberValue` 数字值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Value`|`value`|`float64`|数字值

### `BoolValue` 布尔值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Value`|`value`|`bool`|布尔值

### `UserValue` 用户值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Value`|`value`|`[]string`|用户ID数组

### `FilterDataTimeValue` 日期时间值

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`|日期类型
`Value`|`value`|`[]string`|具体日期值，type为具体日期或具体日期范围时必填


### `reqListSmartsheetFields` 获取智能表格字段请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`|文档的docid
`SheetID`|`sheet_id`|`string`|表格ID
`ViewID`|`view_id,omitempty`|`string`|视图ID
`FieldIDs`|`field_ids,omitempty`|`[]string`|由字段ID组成的数组
`FieldTitles`|`field_titles,omitempty`|`[]string`|由字段标题组成的数组
`Offset`|`offset,omitempty`|`int`|偏移量，初始值为0
`Limit`|`limit,omitempty`|`int`|分页大小，不填或0时，如果总数大于1000，一次性返回1000个字段，当总数小于1000时，返回全部字段；最大值为1000

### `respListSmartsheetFields` 获取智能表格字段响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`respCommon`||-|通用响应
`Total`|`total`|`int`|字段总数
`Fields`|`fields`|`[]Field`|字段详情

### `Field` 字段信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldID`|`field_id`|`string`|字段ID
`FieldTitle`|`field_title`|`string`|字段标题
`FieldType`|`field_type`|`FieldType`|字段类型
`PropertyNumber`|`property_number,omitempty`|`*NumberFieldProperty`|数字类型的字段属性
`PropertyCheckbox`|`property_checkbox,omitempty`|`*CheckboxFieldProperty`|复选框类型的字段属性
`PropertyDateTime`|`property_date_time,omitempty`|`*DateTimeFieldProperty`|日期类型的字段属性
`PropertyAttachment`|`property_attachment,omitempty`|`*AttachmentFieldProperty`|文件类型的字段属性
`PropertyUser`|`property_user,omitempty`|`*UserFieldProperty`|人员类型的字段属性
`PropertyURL`|`property_url,omitempty`|`*UrlFieldProperty`|超链接类型的字段属性
`PropertySelect`|`property_select,omitempty`|`*SelectFieldProperty`|多选类型的字段属性
`PropertyCreatedTime`|`property_created_time,omitempty`|`*CreatedTimeFieldProperty`|创建时间类型的字段属性
`PropertyModifiedTime`|`property_modified_time,omitempty`|`*ModifiedTimeFieldProperty`|最后编辑时间类型的字段属性
`PropertyProgress`|`property_progress,omitempty`|`*ProgressFieldProperty`|进度类型的字段属性
`PropertySingleSelect`|`property_single_select,omitempty`|`*SingleSelectFieldProperty`|单选类型的字段属性
`PropertyReference`|`property_reference,omitempty`|`*ReferenceFieldProperty`|引用类型的字段属性
`PropertyLocation`|`property_location,omitempty`|`*LocationFieldProperty`|地理位置类型的字段属性
`PropertyAutoNumber`|`property_auto_number,omitempty`|`*AutoNumberFieldProperty`|自动编号类型的字段属性
`PropertyCurrency`|`property_currency,omitempty`|`*CurrencyFieldProperty`|货币类型的字段属性
`PropertyWwGroup`|`property_ww_group,omitempty`|`*WwGroupFieldProperty`|群类型的字段属性
`PropertyPercentage`|`property_percentage,omitempty`|`*PercentageFieldProperty`|百分数类型的字段属性


```go
// FieldType 字段类型
type FieldType string

const (
	FieldTypeText         FieldType = "FIELD_TYPE_TEXT"         // 文本
	FieldTypeNumber       FieldType = "FIELD_TYPE_NUMBER"       // 数字
	FieldTypeCheckbox     FieldType = "FIELD_TYPE_CHECKBOX"     // 复选框
	FieldTypeDateTime     FieldType = "FIELD_TYPE_DATE_TIME"    // 日期
	FieldTypeImage        FieldType = "FIELD_TYPE_IMAGE"        // 图片
	FieldTypeAttachment   FieldType = "FIELD_TYPE_ATTACHMENT"   // 文件
	FieldTypeUser         FieldType = "FIELD_TYPE_USER"         // 成员
	FieldTypeURL          FieldType = "FIELD_TYPE_URL"          // 超链接
	FieldTypeSelect       FieldType = "FIELD_TYPE_SELECT"       // 多选
	FieldTypeCreatedUser  FieldType = "FIELD_TYPE_CREATED_USER" // 创建人
	FieldTypeModifiedUser FieldType = "FIELD_TYPE_MODIFIED_USER" // 最后编辑人
	FieldTypeCreatedTime  FieldType = "FIELD_TYPE_CREATED_TIME" // 创建时间
	FieldTypeModifiedTime FieldType = "FIELD_TYPE_MODIFIED_TIME" // 最后编辑时间
	FieldTypeProgress     FieldType = "FIELD_TYPE_PROGRESS"     // 进度
	FieldTypePhoneNumber  FieldType = "FIELD_TYPE_PHONE_NUMBER" // 电话
	FieldTypeEmail        FieldType = "FIELD_TYPE_EMAIL"        // 邮件
	FieldTypeSingleSelect FieldType = "FIELD_TYPE_SINGLE_SELECT" // 单选
	FieldTypeReference    FieldType = "FIELD_TYPE_REFERENCE"    // 关联
	FieldTypeLocation     FieldType = "FIELD_TYPE_LOCATION"     // 地理位置
	FieldTypeFormula      FieldType = "FIELD_TYPE_FORMULA"      // 公式
	FieldTypeCurrency     FieldType = "FIELD_TYPE_CURRENCY"     // 货币
	FieldTypeWwGroup      FieldType = "FIELD_TYPE_WWGROUP"      // 群
	FieldTypeAutoNumber   FieldType = "FIELD_TYPE_AUTONUMBER"   // 自动编号
	FieldTypePercentage   FieldType = "FIELD_TYPE_PERCENTAGE"   // 百分数
)
```

### `NumberFieldProperty` 数字类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`DecimalPlaces`|`decimal_places`|`int`|表示小数点的位数，即数字精度
`UseSeparate`|`use_separate`|`bool`|是否使用千位符，设置此属性后数字字段将以英文逗号分隔千分位，如1,000

### `CheckboxFieldProperty` 复选框类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Checked`|`checked`|`bool`|新增时是否默认勾选

### `DateTimeFieldProperty` 日期类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Format`|`format`|`string`|设置日期格式
`AutoFill`|`auto_fill`|`bool`|新建记录时，是否自动填充时间

### `AttachmentFieldProperty` 文件类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`DisplayMode`|`display_mode`|`string`|展示样式

### `UserFieldProperty` 成员类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`IsMultiple`|`is_multiple`|`bool`|允许添加多个人员
`IsNotified`|`is_notified`|`bool`|添加人员时通知用户，关闭后不通知

### `UrlFieldProperty` 超链接类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`|超链接展示样式

### `SelectFieldProperty` 多选类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`IsQuickAdd`|`is_quick_add`|`bool`|是否允许填写时新增选项
`Options`|`options`|`[]Option`|多选选项的格式设置

### `CreatedTimeFieldProperty` 创建时间类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Format`|`format`|`string`|设置日期格式

### `ModifiedTimeFieldProperty` 最后编辑时间类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Format`|`format`|`string`|设置日期格式

### `ProgressFieldProperty` 进度类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`DecimalPlaces`|`decimal_places`|`int`|小数位数

### `SingleSelectFieldProperty` 单选类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`IsQuickAdd`|`is_quick_add`|`bool`|是否允许填写时新增选项
`Options`|`options`|`[]Option`|单选选项的格式设置

### `ReferenceFieldProperty` 关联字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`SubID`|`sub_id`|`string`|关联的子表id，为空时，表示关联本子表
`FiledID`|`filed_id`|`string`|关联的字段id
`IsMultiple`|`is_multiple`|`bool`|是否允许多选
`ViewID`|`view_id`|`string`|视图id

### `LocationFieldProperty` 地理位置字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`InputType`|`input_type`|`string`|输入类型

### `AutoNumberFieldProperty` 自动编号字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`|输入类型
`Rules`|`rules`|`[]NumberRule`|自定义规则
`ReformatExistingRecord`|`reformat_existing_record`|`bool`|是否应用于已有编号

### `CurrencyFieldProperty` 货币类型字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`CurrencyType`|`currency_type`|`string`|输入类型
`DecimalPlaces`|`decimal_places`|`int`|表示小数点的位数，即数字精度
`UseSeparate`|`use_separate`|`bool`|是否使用千位符

### `WwGroupFieldProperty` 群类型的字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`AllowMultiple`|`allow_multiple`|`bool`|是否允许多个群聊

### `PercentageFieldProperty` 百分数类型的字段属性

Name|JSON|Type|Doc
:---|:---|:---|:--
`DecimalPlaces`|`decimal_places`|`int`|表示小数点的位数，即数字精度
`UseSeparate`|`use_separate`|`bool`|是否使用千位符

### `Option` 选项参数

Name|JSON|Type|Doc
:---|:---|:---|:--
`ID`|`id`|`string`|选项ID
`Text`|`text`|`string`|要填写的选项内容
`Style`|`style`|`int`|选项颜色

### `NumberRule` 自动编号规则

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`string`|规则类型
`Value`|`value`|`string`|存放创建时间格式或固定字符，自增数字位数

### `reqSmartsheetGetRecords` 查询记录请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`DocID`|`docid`|`string`|文档的docid
`SheetID`|`sheet_id`|`string`|Smartsheet 子表ID
`ViewID`|`view_id,omitempty`|`string`|视图 ID
`RecordIDs`|`record_ids,omitempty`|`[]string`|由记录 ID 组成的 JSON 数组
`KeyType`|`key_type,omitempty`|`CellValueKeyType`|返回记录中单元格的key类型
`FieldTitles`|`field_titles,omitempty`|`[]string`|返回指定列，由字段标题组成的 JSON 数组 ，key_type 为 CELL_VALUE_KEY_TYPE_FIELD_TITLE 时有效
`FieldIDs`|`field_ids,omitempty`|`[]string`|返回指定列，由字段 ID 组成的 JSON 数组 ，key_type 为 CELL_VALUE_KEY_TYPE_FIELD_ID 时有效
`Sort`|`sort,omitempty`|`[]Sort`|对返回记录进行排序
`Offset`|`offset,omitempty`|`uint32`|偏移量，初始值为 0
`Limit`|`limit,omitempty`|`uint32`|分页大小，当不填写该参数或将该参数设置为 0 时，如果总数大于 1000，一次性返回 1000 行记录，当总数小于 1000 时，返回全部记录；limit 最大值为 1000

```go
// CellValueKeyType 记录中key的类型
type CellValueKeyType string

const (
	CellValueKeyTypeFieldTitle CellValueKeyType = "CELL_VALUE_KEY_TYPE_FIELD_TITLE" // key用字段标题表示
	CellValueKeyTypeFieldID    CellValueKeyType = "CELL_VALUE_KEY_TYPE_FIELD_ID"    // key用字段 ID 表示
)
```

### `Sort` 排序参数

Name|JSON|Type|Doc
:---|:---|:---|:--
`FieldTitle`|`field_title`|`string`|需要排序的字段标题
`Desc`|`desc,omitempty`|`bool`|是否进行降序排序，默认值为 false

### `respSmartsheetGetRecords` 查询记录响应

Name|JSON|Type|Doc
:---|:---|:---|:--
`respCommon`||-|通用响应
`Total`|`total`|`uint32`|符合筛选条件的视图总数
`HasMore`|`has_more`|`bool`|是否还有更多项
`Next`|`next`|`uint32`|下次下一个搜索结果的偏移量
`Records`|`records`|`[]Record`|由查询记录的具体内容组成的 JSON 数组

### `Record` 记录信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`RecordID`|`record_id`|`string`|记录 ID
`CreateTime`|`create_time`|`string`|记录的创建时间
`UpdateTime`|`update_time`|`string`|记录的更新时间
`Values`|`values`|`map[string]interface{}`|记录的具体内容，key为字段标题或字段ID，value类型根据字段类型不同而异
`CreatorName`|`creator_name`|`string`|创建者名字
`UpdaterName`|`updater_name`|`string`|最后编辑者名字


## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execWedocCreatDoc`|`reqCreateDoc`|`respCreateDoc`|Y|`POST /cgi-bin/wedoc/create_doc`|[新建文档](https://developer.work.weixin.qq.com/document/path/97460)
`execWedocRenameDoc`|`reqRenameDoc`|`respCommon`|Y|`POST /cgi-bin/wedoc/rename_doc`|[重命名文档](https://developer.work.weixin.qq.com/document/path/97736)
`execWedocDelDoc`|`reqDelDoc`|`respCommon`|Y|`POST /cgi-bin/wedoc/del_doc`|[删除文档](https://developer.work.weixin.qq.com/document/path/97735)
`execWedocGetDocBaseInfo`|`reqGetDocBaseInfo`|`respGetDocBaseInfo`|Y|`POST /cgi-bin/wedoc/get_doc_base_info`|[获取文档基础信息](https://developer.work.weixin.qq.com/document/path/97734)
`execWedocDocShare`|`reqDocShare`|`respDocShare`|Y|`POST /cgi-bin/wedoc/doc_share`|[分享文档](https://developer.work.weixin.qq.com/document/path/97733)
`execWedocSmartsheetGetSheet`|`reqGetSmartsheet`|`respGetSmartsheet`|Y|`POST /cgi-bin/wedoc/smartsheet/get_sheet`|[查询子表](https://developer.work.weixin.qq.com/document/path/99911)
`execWedocSmartsheetGetViews`|`reqListSmartsheetViews`|`respListSmartsheetViews`|Y|`POST /cgi-bin/wedoc/smartsheet/get_views`|[查询视图](https://developer.work.weixin.qq.com/document/path/99913)
`execWedocSmartsheetGetFields`|`reqListSmartsheetFields`|`respListSmartsheetFields`|Y|`POST /cgi-bin/wedoc/smartsheet/get_fields`|[查询字段](https://developer.work.weixin.qq.com/document/path/100229)
`execWedocSmartsheetGetRecords`|`reqSmartsheetGetRecords`|`respSmartsheetGetRecords`|Y|`POST /cgi-bin/wedoc/smartsheet/get_records`|[查询记录](https://developer.work.weixin.qq.com/document/path/100230)


