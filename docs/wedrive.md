# 微盘接口

## Models

### `reqWedriveSpaceCreate` 新建空间请求

Name|JSON|Type|Doc
:---|:---|:---|:--
`SpaceName`|`space_name`|`string`|空间标题
`AuthInfo`|`auth_info`|`[]AuthInfo`|空间其他成员信息
`SpaceSubType`|`space_sub_type`|`uint32`|区分创建空间类型, 0:普通（目前只支持0）

```go
//SpaceSubType 空间类型
type SpaceSubType uint32
//SpaceSubTypeNormal 普通
const SpaceSubTypeNormal SpaceSubType = 0
```

### `AuthInfo` 空间成员信息

Name|JSON|Type|Doc
:---|:---|:---|:--
`Type`|`type`|`uint32`|成员类型 1:个人 2:部门
`Userid`|`userid,omitempty`|`string`|成员userid
`Departmentid`|`departmentid,omitempty`|`uint32`|部门departmentid
`Auth`|`auth`|`AuthType`|成员权限 1:仅下载 4:可预览 7:应用空间管理员

```go
// AuthType 表示微盘空间成员权限类型
type AuthType uint32

// 微盘成员权限类型常量
const (
	// AuthTypeDownloadOnly 仅下载权限
	AuthTypeDownloadOnly AuthType = 1
	// AuthTypePreview 可预览权限
	AuthTypePreview AuthType = 4 
	// AuthTypeAdmin 应用空间管理员权限
	AuthTypeAdmin AuthType = 7
) 
```

### `respWedriveSpaceCreate` 新建空间返回

Name|JSON|Type|Doc
:---|:---|:---|:--
`respCommon`||-|通用响应
`SpaceID`|`spaceid`|`string`|空间id

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`wedriveCreateSpace`|`reqWedriveSpaceCreate`|`respWedriveSpaceCreate`|Y|`POST /cgi-bin/wedrive/space_create`|`https://developer.work.weixin.qq.com/document/path/96845`