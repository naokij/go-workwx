# 文档接口

## Models

### `reqCreateDoc` 新建文档请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`SpaceID`|`spaceid`|`string`|空间spaceid。若指定spaceid，则fatherid也要同时指定
`FatherID`|`fatherid`|`string`|	父目录fileid, 在根目录时为空间spaceid
`DocType`|`doc_type`|`DocType`|文档类型, 3:文档 4:表格 10:智能表格
`DocName`|`doc_name`|`string`|文档名字（注意：文件名最多填255个字符, 超过255个字符会被截断）
`AdminUsers`|`admin_users`|`[]string`|文档管理员userid

### `DocType` 文档类型
`DocType` 是无符号32位整数类型 (`uint32`)，可能的值包括：
- `3` - 文档
- `4` - 表格
- `10` - 智能表格

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
`DocID`|`docid`|`string`|文档docid（docid、formid只能填其中一个），仅可删除应用自己创建的文档
`FormID`|`formid`|`string`|收集表id（docid、formid只能填其中一个），仅可删除应用自己创建的收集表

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
`DocID`|`docid`|`string`|文档id（docid、formid只能填其中一个）
`FormID`|`formid`|`string`|收集表id（docid、formid只能填其中一个）

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
```

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
