package workwx

// SpaceCreateReq 微盘创建空间请求参数
type SpaceCreateReq struct {
	// SpaceName 空间标题
	SpaceName string
	// AuthInfo 空间其他成员信息
	AuthInfo []AuthInfo
	// SpaceSubType 区分创建空间类型, 0:普通（目前只支持0）
	SpaceSubType SpaceSubType
}

// CreateSpace 创建微盘空间
//
// 该接口用于创建微盘空间，企业需要使用"微盘"应用才能创建。
//
// See: https://developer.work.weixin.qq.com/document/path/96845
func (c *WorkwxApp) CreateSpace(req SpaceCreateReq) (string, error) {
	// 转换为内部请求结构
	internalReq := reqWedriveSpaceCreate{
		SpaceName:    req.SpaceName,
		SpaceSubType: uint32(req.SpaceSubType),
		AuthInfo:     req.AuthInfo,
	}

	// 调用底层 API
	resp, err := c.wedriveCreateSpace(internalReq)
	if err != nil {
		return "", err
	}

	return resp.SpaceID, nil
}
