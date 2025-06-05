package commands

import (
	"github.com/urfave/cli/v2"

	"github.com/xen0n/go-workwx/v2"
)

// InitApp defines the workwxctl CLI.
func InitApp() *cli.App {
	return &cli.App{
		Name:  "workwxctl",
		Usage: "企业微信命令行客户端 powered by go-workwx",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagCorpID,
				Usage:   "使用 `CORPID` 作为企业 ID",
				EnvVars: []string{"WORKWXCTL_CORPID"},
			},
			&cli.StringFlag{
				Name:    flagCorpSecret,
				Usage:   "使用 `SECRET` 作为应用凭证密钥",
				EnvVars: []string{"WORKWXCTL_CORPSECRET"},
			},
			&cli.Int64Flag{
				Name:    flagAgentID,
				Usage:   "使用 `AGENTID` 作为企业应用 ID",
				EnvVars: []string{"WORKWXCTL_AGENTID"},
			},
			&cli.StringFlag{
				Name:    flagWebhookKey,
				Usage:   "使用 `KEY` 作为群机器人 webhook key",
				EnvVars: []string{"WORKWXCTL_WEBHOOK_KEY"},
			},
			&cli.StringFlag{
				Name:    flagQyapiHostOverride,
				Usage:   "使用 `HOST` 覆盖默认企业微信 API 地址",
				EnvVars: []string{"WORKWXCTL_QYAPI_HOST_OVERRIDE"},
			},
			&cli.StringFlag{
				Name:    flagTLSKeyLogFile,
				Usage:   "将 HTTPS 会话所用密钥写入 `LOGFILE` 以便 Wireshark 等工具读取",
				EnvVars: []string{"WORKWXCTL_TLS_KEY_LOGFILE"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "user-get",
				Usage:  "读取成员",
				Action: cmdUserGet,
			},
			{
				Name:   "user-list-by-dept",
				Usage:  "获取部门成员详情",
				Action: cmdUserListByDept,
			},
			{
				Name:   "dept-list",
				Usage:  "获取部门列表",
				Action: cmdDeptList,
			},
			{
				Name:   "appchat-create",
				Usage:  "创建群聊会话",
				Action: cmdAppchatCreate,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagChatID,
						Usage: "欲建立的群聊 chatid，不能与已有的群重复，最长32个字符。只允许字符0-9及字母a-zA-Z。如果不填，系统会随机生成群id",
					},
					&cli.StringFlag{
						Name:  flagName,
						Usage: "群聊名。最多50个utf8字符，超过将截断",
					},
					&cli.StringFlag{
						Name:  flagOwner,
						Usage: "群主的 ID。如果不指定，系统会随机从群成员列表中选一人作为群主",
					},
					&cli.StringSliceFlag{
						Name:    flagUser,
						Aliases: []string{flagToUserShort},
						Usage:   "群成员 ID，可重复指定。至少2人，至多2000人",
					},
				},
			},
			{
				Name:   "appchat-get",
				Usage:  "获取群聊会话",
				Action: cmdAppchatGet,
			},
			{
				Name:   "send-message",
				Usage:  "发送消息",
				Action: cmdSendMessage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagMessageType,
						Usage: "发送消息的类型: text, image, voice, video, file, textcard, news, mpnews, markdown",
					},
					&cli.StringSliceFlag{
						Name:    flagToUser,
						Aliases: []string{flagToUserShort},
						Usage:   "收信用户 ID (可指定多次)",
					},
					&cli.StringSliceFlag{
						Name:    flagToParty,
						Aliases: []string{flagToPartyShort},
						Usage:   "收信部门 ID (可指定多次)",
					},
					&cli.StringSliceFlag{
						Name:    flagToTag,
						Aliases: []string{flagToTagShort},
						Usage:   "收信标签 ID (可指定多次)",
					},
					&cli.StringFlag{
						Name:    flagToChat,
						Aliases: []string{flagToChatShort},
						Usage:   "收信群聊 chatid (不可与其他收信人选项同时指定)",
					},
					&cli.BoolFlag{
						Name:  flagSafe,
						Usage: "作为保密消息发送",
					},

					// 发消息参数
					&cli.StringFlag{
						Name:  flagMediaID,
						Usage: "图片媒体文件id，可以调用上传临时素材接口获取",
					},
					&cli.StringFlag{
						Name:  flagThumbMediaID,
						Usage: "图文消息缩略图的media_id, 可以通过素材管理接口获得。",
					},
					&cli.StringFlag{
						Name:  flagAuthor,
						Usage: "图文消息的作者，不超过64个字节",
					},
					&cli.StringFlag{
						Name:  flagDescription,
						Usage: "描述，不超过512个字节，超过会自动截断",
					},
					&cli.StringFlag{
						Name:  flagTitle,
						Usage: "标题，不超过128个字节，超过会自动截断",
					},
					&cli.StringFlag{
						Name:  flagURL,
						Usage: "点击后跳转的链接。",
					},
					&cli.StringFlag{
						Name:  flagPicURL,
						Usage: "图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图1068*455，小图150*150。",
					},
					&cli.StringFlag{
						Name:  flagButtonText,
						Usage: "按钮文字。 默认为\"详情\"， 不超过4个文字，超过自动截断。",
					},
					&cli.StringFlag{
						Name:  flagSourceContentURL,
						Usage: "图文消息点击\"阅读原文\"之后的页面链接",
					},
					&cli.StringFlag{
						Name:  flagDigest,
						Usage: "图文消息的描述，不超过512个字节，超过会自动截断",
					},
				},
			},
			{
				Name:   "upload-temp-media",
				Usage:  "上传临时素材",
				Action: cmdUploadTempMedia,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagMediaType,
						Usage: "媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)",
					},
				},
			},
			{
				Name:   "webhook-send-message",
				Usage:  "使用群机器人接口发送消息",
				Action: cmdWebhookSendMessage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagMessageType,
						Usage: "发送消息的类型: text, markdown",
					},
					&cli.StringSliceFlag{
						Name:    flagMentionUser,
						Aliases: []string{flagToUserShort},
						Usage:   "需要被提醒的用户 ID (可指定多次), 特殊值 '" + workwx.MentionAll + "' 表示提醒所有人",
					},
					&cli.StringSliceFlag{
						Name:    flagMentionMobile,
						Aliases: []string{flagMentionMobileShort},
						Usage:   "需要被提醒的用户手机号 (可指定多次), 特殊值 '" + workwx.MentionAll + "' 表示提醒所有人",
					},
				},
			},
			// 企业微信文档管理命令
			{
				Name:   "wedoc-create",
				Usage:  "创建企业微信文档",
				Action: cmdWedocCreate,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagSpaceID,
						Usage: "空间spaceid，若指定spaceid，则father-id也要同时指定",
					},
					&cli.StringFlag{
						Name:  flagFatherID,
						Usage: "父目录fileid, 在根目录时为空间spaceid",
					},
					&cli.StringFlag{
						Name:  flagDocType,
						Usage: "文档类型: doc(文档), sheet(表格), smartsheet(智能表格)",
					},
					&cli.StringFlag{
						Name:  flagDocName,
						Usage: "文档名字（注意：文件名最多填255个字符, 超过255个字符会被截断）",
					},
					&cli.StringSliceFlag{
						Name:  flagAdminUsers,
						Usage: "文档管理员userid列表，可重复指定",
					},
				},
			},
			{
				Name:   "wedoc-rename",
				Usage:  "重命名企业微信文档",
				Action: cmdWedocRename,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档docid，仅可修改应用自己创建的文档，doc-id和form-id只能填其中一个",
					},
					&cli.StringFlag{
						Name:  flagFormID,
						Usage: "收集表id，仅可修改应用自己创建的收集表，doc-id和form-id只能填其中一个",
					},
					&cli.StringFlag{
						Name:  flagNewName,
						Usage: "重命名后的文档名（注意：文档名最多填255个字符，英文算1个，汉字算2个，超过255个字符会被截断）",
					},
				},
			},
			{
				Name:   "wedoc-delete",
				Usage:  "删除企业微信文档",
				Action: cmdWedocDelete,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档docid，仅可删除应用自己创建的文档，doc-id和form-id只能填其中一个",
					},
					&cli.StringFlag{
						Name:  flagFormID,
						Usage: "收集表id，仅可删除应用自己创建的收集表，doc-id和form-id只能填其中一个",
					},
				},
			},
			{
				Name:   "wedoc-get-info",
				Usage:  "获取企业微信文档基础信息",
				Action: cmdWedocGetInfo,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档docid",
					},
				},
			},
			{
				Name:   "wedoc-share",
				Usage:  "分享企业微信文档",
				Action: cmdWedocShare,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档id，doc-id和form-id只能填其中一个",
					},
					&cli.StringFlag{
						Name:  flagFormID,
						Usage: "收集表id，doc-id和form-id只能填其中一个",
					},
				},
			},
			{
				Name:   "smartsheet-list-sheets",
				Usage:  "获取文档中的子表列表",
				Action: cmdSmartsheetListSheets,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档的docid",
					},
					&cli.StringFlag{
						Name:  flagSheetID,
						Usage: "指定子表ID查询（可选）",
					},
					&cli.BoolFlag{
						Name:  "need-all-type-sheet",
						Usage: "是否获取所有类型子表",
					},
				},
			},
			{
				Name:   "smartsheet-list-views",
				Usage:  "获取智能表格视图列表",
				Action: cmdSmartsheetListViews,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档的docid",
					},
					&cli.StringFlag{
						Name:  flagSheetID,
						Usage: "Smartsheet子表ID",
					},
					&cli.StringSliceFlag{
						Name:  flagViewIDs,
						Usage: "需要查询的视图ID数组（可选）",
					},
					&cli.UintFlag{
						Name:  flagOffset,
						Usage: "偏移量，初始值为0",
						Value: 0,
					},
					&cli.UintFlag{
						Name:  flagLimit,
						Usage: "分页大小，不填或0时，如果总数大于1000，一次性返回1000个视图，否则返回全部视图；最大值为1000",
						Value: 0,
					},
				},
			},
			{
				Name:   "smartsheet-list-fields",
				Usage:  "获取智能表格字段列表",
				Action: cmdSmartsheetListFields,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档的docid",
					},
					&cli.StringFlag{
						Name:  flagSheetID,
						Usage: "表格ID",
					},
					&cli.StringFlag{
						Name:  flagViewID,
						Usage: "视图ID（可选）",
					},
					&cli.StringSliceFlag{
						Name:  flagFieldIDs,
						Usage: "由字段ID组成的数组（可选）",
					},
					&cli.StringSliceFlag{
						Name:  flagFieldTitles,
						Usage: "由字段标题组成的数组（可选）",
					},
					&cli.IntFlag{
						Name:  flagOffset,
						Usage: "偏移量，初始值为0",
						Value: 0,
					},
					&cli.IntFlag{
						Name:  flagLimit,
						Usage: "分页大小，不填或0时，如果总数大于1000，一次性返回1000个字段，当总数小于1000时，返回全部字段；最大值为1000",
						Value: 0,
					},
				},
			},
			{
				Name:   "smartsheet-list-records",
				Usage:  "获取智能表格记录列表",
				Action: cmdSmartsheetListRecords,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档的docid",
					},
					&cli.StringFlag{
						Name:  flagSheetID,
						Usage: "Smartsheet子表ID",
					},
					&cli.StringFlag{
						Name:  flagViewID,
						Usage: "视图ID（可选）",
					},
					&cli.StringSliceFlag{
						Name:  flagRecordIDs,
						Usage: "由记录ID组成的数组（可选）",
					},
					&cli.StringFlag{
						Name:  flagKeyType,
						Usage: "返回记录中单元格的key类型: CELL_VALUE_KEY_TYPE_FIELD_TITLE(字段标题), CELL_VALUE_KEY_TYPE_FIELD_ID(字段ID)，默认使用字段标题",
						Value: string(workwx.CellValueKeyTypeFieldTitle),
					},
					&cli.StringSliceFlag{
						Name:  flagFieldTitles,
						Usage: "返回指定列，由字段标题组成的数组（可选），key_type为CELL_VALUE_KEY_TYPE_FIELD_TITLE时有效",
					},
					&cli.StringSliceFlag{
						Name:  flagFieldIDs,
						Usage: "返回指定列，由字段ID组成的数组（可选），key_type为CELL_VALUE_KEY_TYPE_FIELD_ID时有效",
					},
					&cli.StringSliceFlag{
						Name:  flagSortFields,
						Usage: "排序字段标题列表（可选），可指定多个字段进行排序",
					},
					&cli.StringSliceFlag{
						Name:  flagSortDesc,
						Usage: "排序方式列表（可选），与sort-fields一一对应，true表示降序，false表示升序，默认为false",
					},
					&cli.UintFlag{
						Name:  flagOffset,
						Usage: "偏移量，初始值为0",
						Value: 0,
					},
					&cli.UintFlag{
						Name:  flagLimit,
						Usage: "分页大小，不填或0时，如果总数大于1000，一次性返回1000行记录，当总数小于1000时，返回全部记录；最大值为1000",
						Value: 0,
					},
				},
			},
			{
				Name:   "smartsheet-add-records",
				Usage:  "添加智能表格记录",
				Action: cmdSmartsheetAddRecords,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagDocID,
						Usage: "文档的docid",
					},
					&cli.StringFlag{
						Name:  flagSheetID,
						Usage: "Smartsheet子表ID",
					},
					&cli.StringFlag{
						Name:  flagKeyType,
						Usage: "记录中单元格的key类型: CELL_VALUE_KEY_TYPE_FIELD_TITLE(字段标题), CELL_VALUE_KEY_TYPE_FIELD_ID(字段ID)，默认使用字段标题",
						Value: string(workwx.CellValueKeyTypeFieldTitle),
					},
					&cli.StringFlag{
						Name:  "records-file",
						Usage: "包含记录数据的JSON文件路径，格式为数组，每个元素是一个记录对象，键为字段名或ID（取决于key_type）",
					},
				},
			},
			// 企业微信微盘管理命令
			{
				Name:   "wedrive-create-space",
				Usage:  "创建企业微信微盘空间",
				Action: cmdWedriveCreateSpace,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  flagSpaceName,
						Usage: "空间名称（必填）",
					},
					&cli.StringSliceFlag{
						Name:  flagAuthUserids,
						Usage: "授权用户ID列表，可重复指定",
					},
					&cli.UintSliceFlag{
						Name:  flagAuthDeptids,
						Usage: "授权部门ID列表，可重复指定",
					},
					&cli.UintFlag{
						Name:  flagAuthType,
						Usage: "授权类型: 1(仅下载), 4(可预览), 7(应用空间管理员)",
						Value: 4, // 默认可预览
					},
				},
			},
		},
	}
}
