package wechatgo

// UserFormInfoFlag 微信卡券会员卡格式化的选项类型
type UserFormInfoFlag string

const (
	UserFormInfoFlagMobile          UserFormInfoFlag = "USER_FORM_INFO_FLAG_MOBILE"            // 手机号
	UserFormInfoFlagSex             UserFormInfoFlag = "USER_FORM_INFO_FLAG_SEX"               // 性别
	UserFormInfoFlagName            UserFormInfoFlag = "USER_FORM_INFO_FLAG_NAME"              // 姓名
	UserFormInfoFlagBirthday        UserFormInfoFlag = "USER_FORM_INFO_FLAG_BIRTHDAY"          // 生日
	UserFormInfoFlagIDCard          UserFormInfoFlag = "USER_FORM_INFO_FLAG_IDCARD"            // 身份证
	UserFormInfoFlagEmail           UserFormInfoFlag = "USER_FORM_INFO_FLAG_EMAIL"             // 邮箱
	UserFormInfoFlagLocation        UserFormInfoFlag = "USER_FORM_INFO_FLAG_LOCATION"          // 详细地址
	UserFormInfoFlagEducationBackgr UserFormInfoFlag = "USER_FORM_INFO_FLAG_EDUCATION_BACKGRO" // 教育背景
	UserFormInfoFlagIndustry        UserFormInfoFlag = "USER_FORM_INFO_FLAG_INDUSTRY"          // 行业
	UserFormInfoFlagIncome          UserFormInfoFlag = "USER_FORM_INFO_FLAG_INCOME"            // 收入
	UserFormInfoFlagHabit           UserFormInfoFlag = "USER_FORM_INFO_FLAG_HABIT"             // 兴趣爱好
)

// ReimburseStatus 发票报销状态
type ReimburseStatus string

const (
	ReimburseStatusInit    ReimburseStatus = "INVOICE_REIMBURSE_INIT"    // 初始状态，未锁定，可提交报销
	ReimburseStatusLock    ReimburseStatus = "INVOICE_REIMBURSE_LOCK"    // 已锁定，无法重复提交报销
	ReimburseStatusClosure ReimburseStatus = "INVOICE_REIMBURSE_CLOSURE" // 已核销，从用户卡包中移除
)

// WeChatErrorCode 微信接口返回码
// 全局返回码请参考 https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Global_Return_Code.html
type WeChatErrorCode int

const (
	// 系统错误
	SystemError WeChatErrorCode = -1000

	// 系统繁忙，此时请开发者稍候再试
	SystemBusy WeChatErrorCode = -1

	// 请求成功
	Success WeChatErrorCode = 0

	// AppSecret 错误，或是 Access Token 无效
	// 请开发者认真比对AppSecret的正确性，或查看是否正在为恰当的公众号调用接口
	InvalidCredential WeChatErrorCode = 40001

	// 错误的凭证类型
	InvalidCredentialType WeChatErrorCode = 40002

	// 错误的 OpenID
	// 请开发者确认 OpenID 是否已关注公众号，或是否是其他公众号的 OpenID
	InvalidOpenID WeChatErrorCode = 40003

	// 不支持��媒体文件类型
	InvalidMediaType WeChatErrorCode = 40004

	// 不支持的文件类型
	InvalidFileType WeChatErrorCode = 40005

	// 不支持的文件大小
	InvalidFileSize WeChatErrorCode = 40006

	// 错误的 MediaID
	InvalidMediaID WeChatErrorCode = 40007

	// 错误的消息类型
	InvalidMessageType WeChatErrorCode = 40008

	// 不支持的图片大小，图片格式不对有时也会报这个错
	InvalidImageSize WeChatErrorCode = 40009

	// 不支持的语音文件大小
	InvalidVoiceSize WeChatErrorCode = 40010

	// 不支持的视频文件大小
	InvalidVideoSize WeChatErrorCode = 40011

	// 不支持的缩略图大小
	InvalidThumbSize WeChatErrorCode = 40012

	// 错误的 AppID，目前 AppID 格式都是 /^wx\d{16}$/
	InvalidAppID WeChatErrorCode = 40013

	// 不合法的 Access Token
	// 请开发者认真比对 Access Token 的有效性（如是否过期），或查看是否正在为恰当的公众号调用接口
	InvalidAccessToken WeChatErrorCode = 40014

	// 错误的按钮类型
	InvalidButtonType WeChatErrorCode = 40015

	// 不支持的主菜单按钮个数，微信自定义菜单按钮个数应该在 1~3 个之间
	InvalidButtonSize WeChatErrorCode = 40016

	// 不支持的子菜单按钮个数，微信自定义子菜单按钮个数应该在 1~5 个之间
	InvalidSubButtonSize WeChatErrorCode = 40017

	// 不支持的按钮名字长度
	InvalidButtonNameSize WeChatErrorCode = 40018

	// 不支持的按钮 key 长度
	InvalidButtonKeySize WeChatErrorCode = 40019

	// 不支持的按钮 url 长度
	InvalidButtonURLSize WeChatErrorCode = 40020

	// 不合法的菜单版本号
	InvalidMenuVersion WeChatErrorCode = 40021

	// 不合法的子菜单级数
	InvalidSubButtonLevel WeChatErrorCode = 40022

	// 不合法的子菜单按钮个数
	InvalidSubButtonCount WeChatErrorCode = 40023

	// 不合法的子菜单按钮类型
	InvalidSubButtonType WeChatErrorCode = 40024

	// 不合法的子菜单按钮名字长度
	InvalidSubButtonNameSize WeChatErrorCode = 40025

	// 不合法的子菜单按钮 key 长度
	InvalidSubButtonKeySize WeChatErrorCode = 40026

	// 不合法的子菜单按钮 url 长度
	InvalidSubButtonURLSize WeChatErrorCode = 40027

	// 不合法的自定义菜单使用用户
	InvalidMenuUser WeChatErrorCode = 40028

	// 错误的 OAuth Code
	InvalidOAuthCode WeChatErrorCode = 40029

	// 错误的 Refresh Token
	InvalidRefreshToken WeChatErrorCode = 40030

	// 错误的 OpenID 列表
	InvalidOpenIDList WeChatErrorCode = 40031

	// 错误的 OpenID 列表长度，列表内最多10000个 OpenID
	InvalidOpenIDListSize WeChatErrorCode = 40032

	// 不支持的请求字符，不能包含 \uxxxx 格式的字符
	InvalidRequestCharset WeChatErrorCode = 40033

	// 不合法的参数
	InvalidParameter WeChatErrorCode = 40035

	// 错误的模板消息 ID，Template ID 失效了，请重新刷新一次 Template ID
	InvalidTemplate WeChatErrorCode = 40037

	// 不合法的请求格式
	InvalidRequestFormat WeChatErrorCode = 40038

	// 不合法的 url 长度
	InvalidURLSize WeChatErrorCode = 40039

	// 无效的 url
	InvalidURLDomain WeChatErrorCode = 40048

	// 不合法的分组 ID
	InvalidGroupID WeChatErrorCode = 40050

	// 不合法的分组名字，40117 也是这个错误
	InvalidGroupName WeChatErrorCode = 40051

	// 不支持的操作，可能是该公众号已经申请完了十万个二维码
	InvalidActionInfo WeChatErrorCode = 40053

	// 自定义菜单的按钮里，网址有误
	InvalidButtonDomain WeChatErrorCode = 40054

	// 自定义子菜单的按钮里，网址有误
	InvalidSubButtonDomain WeChatErrorCode = 40055

	// 删除单篇图文时，指定的 article_idx 不合法
	InvalidDeleteArticleID WeChatErrorCode = 40060

	// 错误的行业号，有一些模板消息只会在特定的行业下申请
	InvalidIndustryID WeChatErrorCode = 40102

	// 不支持的 MediaID 长度
	InvalidMediaIDSize WeChatErrorCode = 40118

	// button 类型错误
	InvalidUseButtonType WeChatErrorCode = 40119

	// 子 button 类型错误
	InvalidUseSubButtonType WeChatErrorCode = 40120

	// 不支持的 MediaID 类型
	InvalidMediaIDType WeChatErrorCode = 40121

	// 无效的 AppSecret
	InvalidAppSecret WeChatErrorCode = 40125

	// 微信号不合法
	InvalidWeChatID WeChatErrorCode = 40132

	// 不支持的图片格式
	InvalidImageFormat WeChatErrorCode = 40137

	// 请勿添加其他公众号的主页链接
	ContainOtherHomePageURL WeChatErrorCode = 40155

	// OAuth Code 已使用
	CodeBeenUsed WeChatErrorCode = 40163

	// 缺少 Access Token 参数
	MissingAccessToken WeChatErrorCode = 41001

	// 缺少 AppID 参数
	MissingAppID WeChatErrorCode = 41002

	// 缺少 Refresh Token 参数
	MissingRefreshToken WeChatErrorCode = 41003

	// 缺少 AppSecret 参数
	MissingAppSecret WeChatErrorCode = 41004

	// 缺少多媒体文件数据
	MissingMediaData WeChatErrorCode = 41005

	// 缺少 MediaID 参数
	MissingMediaID WeChatErrorCode = 41006

	// 缺少子菜单数据
	MissingSubButtons WeChatErrorCode = 41007

	// 缺少 OAuth Code
	MissingOAuthCode WeChatErrorCode = 41008

	// 缺少 OpenID
	MissingOpenID WeChatErrorCode = 41009

	// page 路径不正确，需要保证在现网版本小程序中存在，与 app.json 保持一致
	InvalidPage WeChatErrorCode = 41030

	// Access Token 已失效，请检查 Access Token 的有效期，重新刷新 Access Token
	ExpiredAccessToken WeChatErrorCode = 42001

	// Refresh Token 已失效
	ExpiredRefreshToken WeChatErrorCode = 42002

	// OAuth Code 已失效
	ExpiredOAuthCode WeChatErrorCode = 42003

	// 授权已失效，用户修改微信密码，Access Token, Refresh Token 均已失效，需要重新授权
	ExpiredAuthorization WeChatErrorCode = 42007

	// 需要 Get 请求
	RequireGet WeChatErrorCode = 43001

	// 需要 Post 请求
	RequirePost WeChatErrorCode = 43002

	// 需要 Https 请求
	RequireHTTPS WeChatErrorCode = 43003

	// 用户没有关注公众号
	RequireSubscribe WeChatErrorCode = 43004

	// 需要好友关系
	RequireFriend WeChatErrorCode = 43005

	// 用户被拉黑，需要公众号把该用户从黑名单里移除
	RequireUnblockUser WeChatErrorCode = 43019

	// 超过了更换行业的限制，一个月最多换一次
	OutOfChangeIndustryLimit WeChatErrorCode = 43100

	// 用户拒绝接受消息，如果用户之前曾经订阅过，则表示用户取消了订阅关系
	UserRefuseToAcceptTheMessage WeChatErrorCode = 43101

	// 多媒体文件为空
	EmptyMediaData WeChatErrorCode = 44001

	// POST 的数据包为空
	EmptyPostData WeChatErrorCode = 44002

	// 图文消息内容为空
	EmptyNewsData WeChatErrorCode = 44003

	// 文本消息内容为空
	EmptyContent WeChatErrorCode = 44004

	// 多媒体文件大小超过限制，最大允许 1MB
	OutOfMediaSizeLimit WeChatErrorCode = 45001

	// 消息内容超过限制
	OutOfContentSizeLimit WeChatErrorCode = 45002

	// 标题长度超过限制，最长允许 64 字符长度
	OutOfTitleSizeLimit WeChatErrorCode = 45003

	// 描述字段超过限制
	OutOfDescriptionSizeLimit WeChatErrorCode = 45004

	// 链接字段超过限制
	OutOfURLSizeLimit WeChatErrorCode = 45005

	// 图片链接字段超过限制
	OutOfPicURLSizeLimit WeChatErrorCode = 45006

	// 语音播放时间超过限制，最长允许 60 秒
	OutOfVoiceTimeLimit WeChatErrorCode = 45007

	// 图文消息数量超过限制，最多 10 条图文消息
	OutOfArticleSizeLimit WeChatErrorCode = 45008

	// 接口调用频率超过限制
	OutOfAPIFreqLimit WeChatErrorCode = 45009

	// 创建菜单个数超过限制
	OutOfMenuSizeLimit WeChatErrorCode = 45010

	// API 调用太频繁，请稍候再试
	APIMinuteQuotaReachLimit WeChatErrorCode = 45011

	// 回复时间超过限制
	// 接受推送后，5 秒内未被动响应。或者是用户与公众号 48 小时无互动后，调用客服接口主动推送消息。
	OutOfResponseTimeLimit WeChatErrorCode = 45015

	// 系统分组，不允许修改
	SystemGroupCannotChange WeChatErrorCode = 45016

	// 分组名字过长
	OutOfGroupNameSizeLimit WeChatErrorCode = 45017

	// 分组数量超过上限
	OutOfGroupSizeLimit WeChatErrorCode = 45018

	// 模板消息数量超过限制
	OutOfTemplateSizeLimit WeChatErrorCode = 45026

	// 模板消息与行业信息冲突
	TemplateConflictWithIndustry WeChatErrorCode = 45027

	// 客服接口下行条数超过上限
	OutOfResponseCountLimit WeChatErrorCode = 45047

	// 创建菜单包含未关联的小程序
	NoPermissionToUseWeappInMenu WeChatErrorCode = 45064

	// 相同 clientmsgid 已存在群发记录，返回数据中带有已存在的群发任务的 msgid
	ClientMsgIDExist WeChatErrorCode = 45065

	// 相同 clientmsgid 重试速度过快，请间隔1分钟重试
	OutOfClientMsgIDAPIFreqLimit WeChatErrorCode = 45066

	// clientmsgid 长度超过限制
	ClientMsgIDSizeOutOfLimit WeChatErrorCode = 45067

	// 不支持的图文消息内容，请确认 content 里没有超链接标签
	InvalidContent WeChatErrorCode = 45166

	// 不存在媒体数据
	MediaDataNoExist WeChatErrorCode = 46001

	// 不存在的菜单版本
	MenuVersionNotExist WeChatErrorCode = 46002

	// 不存在的菜单数据
	MenuNoExist WeChatErrorCode = 46003

	// 不存在的用户
	UserNotExist WeChatErrorCode = 46004

	// 解析 JSON/XML 内容错误
	DataFormatError WeChatErrorCode = 47001

	// 模板参数不准确，可能为空或者不满足规则，errmsg 会提示具体是哪个字段出错
	InvalidTemplateArgument WeChatErrorCode = 47003

	// API 功能未授权，请确认公众号已获得该接口，可以在公众平台官网-开发者中心页中查看接口权限
	UnauthorizedAPI WeChatErrorCode = 48001

	// 用户拒收公众号消息 (在公众号选项中，关闭了"接收消息")
	UserBlockMessage WeChatErrorCode = 48002

	// 公众号管理员没有同意微信群发协议，请登录公众号后台点一下同意
	UserNotAgreeProtocol WeChatErrorCode = 48003

	// API 接口被封禁，请登录公众号后台查看详情
	APIBanned WeChatErrorCode = 48004

	// API 禁止删除被自动回复和自定义菜单引用的素材
	APIDeleteProhibited WeChatErrorCode = 48005

	// API 清零次数失败，因为清零次数达到上限
	OutOfResetLimit WeChatErrorCode = 48006

	// 没有该类型消息的发送权限
	NoPermissionForThisMsgType WeChatErrorCode = 48008

	// 用户未授权该 API
	UserUnauthorized WeChatErrorCode = 50001

	// 用户受限，可能是违规后接口被封禁
	UserLimited WeChatErrorCode = 50002

	// 用户未关注的公众号
	UnsubscribeOfficialAccount WeChatErrorCode = 50005

	// 发布功能被封禁
	PublishLimited WeChatErrorCode = 53500

	// 频繁请求发布
	OutOfPublishLimit WeChatErrorCode = 53501

	// Publish ID 无效
	InvalidPublishID WeChatErrorCode = 53502

	// Article ID 无效
	InvalidArticleID WeChatErrorCode = 53600

	// 公众号未授权给开放平台
	UnauthorizedComponent WeChatErrorCode = 61003

	// 公众号未授权该 API 给开放平台
	UnauthorizedComponentAPI WeChatErrorCode = 61007

	// 错误的开放平台 Refresh Token
	InvalidComponentRefreshToken WeChatErrorCode = 61023

	// 参数错误
	ErrorParameter WeChatErrorCode = 61451

	// 无效客服账号
	InvalidKFAccount WeChatErrorCode = 61452

	// 客服帐号已存在
	KFAccountExisted WeChatErrorCode = 61453

	// 客服帐号名长度超过限制 ( 仅允许 10 个英文字符，不包括 @ 及 @ 后的公众号的微信号 )
	InvalidKFAccountLength WeChatErrorCode = 61454

	// 客服帐号名包含非法字符 ( 仅允许英文 + 数字 )
	IllegalCharterInKFAccount WeChatErrorCode = 61455

	// 客服帐号个数超过限制 (10 个客服账号 )
	KFAccountExceeded WeChatErrorCode = 61456

	// 无效头像文件类型
	InvalidAvatarFileType WeChatErrorCode = 61457

	// 日期格式错误
	DateFormatError WeChatErrorCode = 61500

	// 部分参数为空
	MissingParameter WeChatErrorCode = 63001

	// 无效的 JS SDK 签名
	InvalidJSSDKSignature WeChatErrorCode = 63002

	// 不存在此 menuid 对应的个性化菜单
	InvalidMenuID WeChatErrorCode = 65301

	// 没有默认菜单，不能创建个性化菜单
	ThereIsNoSelfMenu WeChatErrorCode = 65303

	// MatchRule 信息为空
	MatchRuleEmpty WeChatErrorCode = 65304

	// 个性化菜单数量受限
	MenuCountLimit WeChatErrorCode = 65305

	// 不支持个性化菜单的帐号
	InvalidAccountForMenu WeChatErrorCode = 65306

	// 个性化菜单信息为空
	EmptyMenu WeChatErrorCode = 65307

	// 包含没有响应类型的 button
	ButtonMissingResponse WeChatErrorCode = 65308

	// 个性化菜单开关处于关闭状态
	DisableMenu WeChatErrorCode = 65309

	// 填写了省份或城市信息，国家信息不能为空
	MissingCountry WeChatErrorCode = 65310

	// 填写了城市信息，省份信息不能为空
	MissingProvince WeChatErrorCode = 65311

	// 不合法的国家信息
	InvalidCountry WeChatErrorCode = 65312

	// 不合法的省份信息
	InvalidProvince WeChatErrorCode = 65313

	// 不合法的城市信息
	InvalidCityInfo WeChatErrorCode = 65314

	// 该公众号的菜单设置了过多的域名外跳（最多跳转到 3 个域名的链接）
	DomainCountReachLimit WeChatErrorCode = 65316

	// 不合法的 URL
	InvalidURL WeChatErrorCode = 65317

	// 无效的签名
	InvalidSignature WeChatErrorCode = 87009

	// 内容可能潜在风险
	RiskyContent WeChatErrorCode = 87014

	// POST 数据参数不合法
	InvalidPostData WeChatErrorCode = 9001001

	// 远端服务不可用
	RemoteServiceUnavailable WeChatErrorCode = 9001002

	// Ticket 不合法
	InvalidTicket WeChatErrorCode = 9001003

	// 获取摇周边用户信息失败
	GetUserFailed WeChatErrorCode = 9001004

	// 获取商户信息失败
	GetMerchantFailed WeChatErrorCode = 9001005

	// 获取 OpenID 失败
	GetOpenIDFailed WeChatErrorCode = 9001006

	// 上传文件缺失
	UploadFileMissing WeChatErrorCode = 9001007

	// 上传素材的文件类型不合法
	UploadFileTypeError WeChatErrorCode = 9001008

	// 上传素材的文件尺寸不合法
	UploadFileSizeError WeChatErrorCode = 9001009

	// 上传失败
	UploadFailed WeChatErrorCode = 9001010

	// 帐号不合法
	InvalidAccount WeChatErrorCode = 9001020

	// 已有设备激活率低于 50% ，不能新增设备
	ActiveDeviceLessThanHalf WeChatErrorCode = 9001021

	// 设备申请数不合法，必须为大于 0 的数字
	InvalidDeviceCount WeChatErrorCode = 9001022

	// 已存在审核中的设备 ID 申请
	DeviceIDExisted WeChatErrorCode = 9001023

	// 一次查询设备 ID 数量不能超过 50
	OutOfDeviceQueryLimit WeChatErrorCode = 9001024

	// 设备 ID 不合法
	InvalidDeviceID WeChatErrorCode = 9001025

	// 页面 ID 不合法
	InvalidPageID WeChatErrorCode = 9001026

	// 页面参数不合法
	InvalidPageParameter WeChatErrorCode = 9001027

	// 一次删除页面 ID 数量不能超过 10
	OutOfPageDeleteLimit WeChatErrorCode = 9001028

	// 页面已应用在设备中，请先解除应用关系再删除
	PageAppliedInDevice WeChatErrorCode = 9001029

	// 一次查询页面 ID 数量不能超过 50
	OutOfPageQueryLimit WeChatErrorCode = 9001030

	// 时间区间不合法
	InvalidTimeRange WeChatErrorCode = 9001031

	// 保存设备与页面的绑定关系参数错误
	BindDevicePageParameterError WeChatErrorCode = 9001032

	// 门店 ID 不合法
	InvalidLocationID WeChatErrorCode = 9001033

	// 设备备注信息过长
	OutOfDeviceDescLengthLimit WeChatErrorCode = 9001034

	// 设备申请参数不合法
	InvalidDeviceParameter WeChatErrorCode = 9001035

	// 查询起始值 begin 不合法
	InvalidBegin WeChatErrorCode = 9001036
)
