package models

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
