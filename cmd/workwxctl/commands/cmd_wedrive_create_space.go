package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/xen0n/go-workwx/v2"
)

func cmdWedriveCreateSpace(c *cli.Context) error {
	cfg := mustGetConfig(c)

	spaceName := c.String(flagSpaceName)
	if spaceName == "" {
		return fmt.Errorf("空间名称不能为空")
	}

	// 获取空间类型，目前只支持普通空间(0)
	spaceSubType := workwx.SpaceSubTypeNormal

	// 获取授权成员信息
	userIDs := c.StringSlice(flagAuthUserids)
	deptIDs := c.UintSlice(flagAuthDeptids)
	authType := c.Uint(flagAuthType)

	// 验证授权类型
	if authType != uint(workwx.AuthTypeDownloadOnly) &&
		authType != uint(workwx.AuthTypePreview) &&
		authType != uint(workwx.AuthTypeAdmin) {
		return fmt.Errorf("不支持的授权类型: %d，支持的类型: 1(仅下载), 4(可预览), 7(应用空间管理员)", authType)
	}

	// 构建授权成员信息
	authInfo := []workwx.AuthInfo{}

	// 添加用户授权
	for _, userid := range userIDs {
		authInfo = append(authInfo, workwx.AuthInfo{
			Type:   1, // 个人
			Userid: userid,
			Auth:   workwx.AuthType(authType),
		})
	}

	// 添加部门授权
	for _, deptid := range deptIDs {
		authInfo = append(authInfo, workwx.AuthInfo{
			Type:         2, // 部门
			Departmentid: uint32(deptid),
			Auth:         workwx.AuthType(authType),
		})
	}

	// 如果没有授权成员，给出提示
	if len(authInfo) == 0 {
		fmt.Println("警告: 未指定任何授权成员，空间将只有创建者有权限访问")
	}

	// 创建请求结构
	req := workwx.SpaceCreateReq{
		SpaceName:    spaceName,
		AuthInfo:     authInfo,
		SpaceSubType: spaceSubType,
	}

	app := cfg.MakeWorkwxApp()

	// 调用API创建空间
	spaceID, err := app.CreateSpace(req)
	if err != nil {
		fmt.Printf("创建空间失败: %+v\n", err)
		return err
	}

	fmt.Printf("空间创建成功:\nSpace ID: %s\n", spaceID)
	return nil
}
