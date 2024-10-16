package tencentcloudess

type Admin struct {
	// 超管名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"name"`

	// 超管手机号，打码显示
	// 示例值：138****1569
	//
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"mobile"`
}

type Agent struct {
	// 被代理机构在电子签平台的机构编号，集团代理下场景必传
	ProxyOrganizationId *string `json:"proxyOrganizationId"`
}

type ApproverComponentLimitType struct {
	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	RecipientId *string `json:"recipientId"`

	// 签署方经办人控件类型是个人印章签署控件（SIGN_SIGNATURE） 时，可选的签名方式，可多选
	//
	// 签名方式：
	// <ul>
	// <li>HANDWRITE-手写签名</li>
	// <li>ESIGN-个人印章类型</li>
	// <li>OCR_ESIGN-AI智能识别手写签名</li>
	// <li>SYSTEM_ESIGN-系统签名</li>
	// </ul>
	Values []*string `json:"values"`
}

type ApproverInfo struct {
	// 在指定签署方时，可选择企业B端或个人C端等不同的参与者类型，可选类型如下:
	// **0**：企业
	// **1**：个人
	// **3**：企业静默签署
	// 注：`类型为3（企业静默签署）时，此接口会默认完成该签署方的签署。静默签署仅进行盖章操作，不能自动签名。`
	// **7**: 个人自动签署，适用于个人自动签场景。
	// 注: `个人自动签场景为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverType *int64 `json:"approverType"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	ApproverName *string `json:"approverName"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此合同签署方。
	ApproverMobile *string `json:"approverMobile"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	// 如果签署方是企业签署方(approverType = 0 或者 approverType = 3)， 则企业名称必填。
	OrganizationName *string `json:"organizationName"`

	// 合同中的签署控件列表，列表中可支持下列多种签署控件,控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 个人签名/印章</li>
	// <li> 企业印章</li>
	// <li> 骑缝章等签署控件</li></ul>
	SignComponents []*Component `json:"signComponents"`

	// 签署方经办人的证件类型，支持以下类型，样式可以参考<a href="https://qian.tencent.com/developers/partner/id_card_support/" target="_blank">常见个人证件类型介绍</a>
	// <ul><li>ID_CARD 中国大陆居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE 其他证件</li></ul>
	//
	//
	//
	//
	// 注:
	// 1. <b>其他证件类型为白名单功能</b>，使用前请联系对接的客户经理沟通。
	// 2. 港澳居民来往内地通行证 和  港澳台居民居住证 类型的签署人<b>至少要过一次大陆的海关</b>才能使用。
	ApproverIdCardType *string `json:"approverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给香港居民，“M”字头签发给澳门居民；第2位至第11位为数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"approverIdCardNumber"`

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li>  **sms**  :  (默认)短信</li>
	// <li>   **none**   : 不通知</li></ul>
	//
	// 注意：
	// `如果使用的是通过文件发起合同（CreateFlowByFiles），NotifyType必须 是 sms 才会发送短信`
	NotifyType *string `json:"notifyType"`

	// 收据场景设置签署人角色类型, 可以设置如下<b>类型</b>:
	// <ul><li> **1**  :收款人</li>
	// <li>   **2**   :开具人</li>
	// <li>   **3** :见证人</li></ul>
	// 注: `收据场景为白名单功能，使用前请联系对接的客户经理沟通。`
	ApproverRole *int64 `json:"approverRole"`

	// 可以自定义签署人角色名：收款人、开具人、见证人等，长度不能超过20，只能由中文、字母、数字和下划线组成。
	//
	// 注: `如果是用模板发起, 优先使用此处上传的, 如果不传则用模板的配置的`
	ApproverRoleName *string `json:"approverRoleName"`

	// 签署方在签署合同之前，需要强制阅读合同的时长，可指定为3秒至300秒之间的任意值。
	//
	// 若未指定阅读时间，则会按照合同页数大小计算阅读时间，计算规则如下：
	// <ul><li>合同页数少于等于2页，阅读时间为3秒；</li>
	// <li>合同页数为3到5页，阅读时间为5秒；</li>
	// <li>合同页数大于等于6页，阅读时间为10秒。</li></ul>
	PreReadTime *int64 `json:"preReadTime"`

	// 签署人userId，仅支持本企业的员工userid， 可在控制台组织管理处获得
	//
	// 注：
	// 如果传进来的<font color="red">UserId已经实名， 则忽略ApproverName，ApproverIdCardType，ApproverIdCardNumber，ApproverMobile这四个入参</font>（会用此UserId实名的身份证和登录的手机号覆盖）
	UserId *string `json:"userId"`

	// 在企微场景下使用，需设置参数为**WEWORKAPP**，以表明合同来源于企微。
	ApproverSource *string `json:"approverSource"`

	// 在企业微信场景下，表明该合同流程为或签，其最大长度为64位字符串。
	// 所有参与或签的人员均需具备该标识。
	// 注意，在合同中，不同的或签参与人必须保证其CustomApproverTag唯一。
	// 如果或签签署人为本方企业微信参与人，则需要指定ApproverSource参数为WEWORKAPP。
	CustomApproverTag *string `json:"customApproverTag"`

	// 可以控制签署方在签署合同时能否进行某些操作，例如拒签、转交他人等。
	// 详细操作可以参考开发者中心的ApproverOption结构体。
	ApproverOption *ApproverOption `json:"approverOption"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注:
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段可传多个校验方式</li></ul>
	ApproverVerifyTypes []*int64 `json:"approverVerifyTypes"`

	// 您可以指定签署方签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证，需进行人脸识别成功后才能签署合同；</li>
	// <li>**2**：签署密码，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署；</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息（名字、手机号、证件号）校验一致才能成功进行合同签署。（如果是港澳台客户，建议不要选择这个）</li>
	// <li>**5**：设备指纹识别，需要对比手机机主预留的指纹信息，校验一致才能成功进行合同签署。（iOS系统暂不支持该校验方式）</li>
	// <li>**6**：设备面容识别，需要对比手机机主预留的人脸信息，校验一致才能成功进行合同签署。（Android系统暂不支持该校验方式）</li></ul>
	//
	//
	// 默认为1(人脸认证 ),2(签署密码),3(运营商三要素),5(设备指纹识别),6(设备面容识别)
	//
	// 注：
	// 1. 用<font color='red'>模板创建合同场景</font>, 签署人的认证方式需要在配置模板的时候指定, <font color='red'>在创建合同重新指定无效</font>
	// 2. 运营商三要素认证方式对手机号运营商及前缀有限制,可以参考[运营商支持列表类](https://qian.tencent.com/developers/company/mobile_support)得到具体的支持说明
	// 3. 校验方式不允许只包含<font color='red'>设备指纹识别</font>和<font color='red'>设备面容识别</font>，至少需要再增加一种其他校验方式。
	// 4. <font color='red'>设备指纹识别</font>和<font color='red'>设备面容识别</font>只支持小程序使用，其他端暂不支持。
	ApproverSignTypes []*int64 `json:"approverSignTypes"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li>**false**：（默认）不需要审批，直接签署。</li>
	// <li>**true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li>如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li>如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	//
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	ApproverNeedSignReview *bool `json:"approverNeedSignReview"`

	// [用PDF文件创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowByFiles)时,如果设置了外层参数SignBeanTag=1(允许签署过程中添加签署控件),则可通过此参数明确规定合同所使用的签署控件类型（骑缝章、普通章法人章等）和具体的印章（印章ID或者印章类型）或签名方式。
	//
	// 注：`限制印章控件或骑缝章控件情况下,仅本企业签署方可以指定具体印章（通过传递ComponentValue,支持多个），他方企业或个人只支持限制控件类型。`
	AddSignComponentsLimits []*ComponentLimit `json:"addSignComponentsLimits"`

	// 签署须知：支持传入富文本，最长字数：500个中文字符
	SignInstructionContent *string `json:"signInstructionContent"`

	// 签署人的签署截止时间，格式为Unix标准时间戳（秒）
	//
	// 注: `若不设置此参数，则默认使用合同的截止时间，此参数暂不支持合同组子合同`
	Deadline *int64 `json:"deadline"`

	// 签署人在合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>数据表格等填写控件</li></ul>
	//
	// 具体使用说明可参考[为签署方指定填写控件](https://qian.tencent.cn/developers/company/createFlowByFiles/#指定签署方填写控件)
	//
	// 注：`此参数仅在通过文件发起合同或者合同组时生效`
	Components []*Component `json:"components"`
}

type ApproverItem struct {
	// 签署方唯一编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"signId"`

	// 签署方角色编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"recipientId"`

	// 签署方角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverRoleName *string `json:"approverRoleName"`
}

type ApproverOption struct {
	// 签署方是否可以拒签
	//
	// <ul><li> **false** : ( 默认)可以拒签</li>
	// <li> **true** :不可以拒签</li></ul>
	NoRefuse *bool `json:"noRefuse"`

	// 签署方是否可以转他人处理
	//
	// <ul><li> **false** : ( 默认)可以转他人处理</li>
	// <li> **true** :不可以转他人处理</li></ul>
	NoTransfer *bool `json:"noTransfer"`

	// 允许编辑签署人信息（嵌入式使用） 默认true-可以编辑 false-不可以编辑
	CanEditApprover *bool `json:"canEditApprover"`

	// 签署人信息补充类型，默认无需补充。
	//
	// <ul><li> **1** : ( 动态签署人（可发起合同后再补充签署人信息）注：`企业自动签不支持动态补充`</li></ul>
	//
	// 注：
	// `使用动态签署人能力前，需登陆腾讯电子签控制台打开服务开关`
	FillType *int64 `json:"fillType"`

	// 签署人阅读合同限制参数
	//  <br/>取值：
	// <ul>
	// <li> LimitReadTimeAndBottom，阅读合同必须限制阅读时长并且必须阅读到底</li>
	// <li> LimitReadTime，阅读合同仅限制阅读时长</li>
	// <li> LimitBottom，阅读合同仅限制必须阅读到底</li>
	// <li> NoReadTimeAndBottom，阅读合同不限制阅读时长且不限制阅读到底（白名单功能，请联系客户经理开白使用）</li>
	// </ul>
	FlowReadLimit *string `json:"flowReadLimit"`
}

type ApproverRestriction struct {
	// 指定签署人名字
	Name *string `json:"name"`

	// 指定签署人手机号，11位数字
	Mobile *string `json:"mobile"`

	// 指定签署人证件类型，ID_CARD-身份证
	IdCardType *string `json:"idCardType"`

	// 指定签署人证件号码，字母大写
	IdCardNumber *string `json:"idCardNumber"`
}

type ArchiveDynamicApproverData struct {
	// 签署方唯一编号，一个全局唯一的标识符，不同的流程不会出现冲突。
	//
	// 可以使用签署方的唯一编号来生成签署链接（也可以通过RecipientId来生成签署链接）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"signId"`

	// 签署方角色编号，签署方角色编号是用于区分同一个流程中不同签署方的唯一标识。不同的流程会出现同样的签署方角色编号。
	//
	// 填写控件和签署控件都与特定的角色编号关联。
	//
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"recipientId"`
}

type AuthInfoDetail struct {
	// 扩展服务类型，和入参一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"type"`

	// 扩展服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"name"`

	// 授权员工列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthUserList []*HasAuthUser `json:"hasAuthUserList"`

	// 授权企业列表（企业自动签时，该字段有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthOrganizationList []*HasAuthOrganization `json:"hasAuthOrganizationList"`

	// 授权员工列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthUserTotal *int64 `json:"authUserTotal"`

	// 授权企业列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthOrganizationTotal *int64 `json:"authOrganizationTotal"`
}

type AuthRecord struct {
	// 经办人姓名。
	OperatorName *string `json:"operatorName"`

	// 经办人手机号。
	OperatorMobile *string `json:"operatorMobile"`

	// 认证授权方式：
	// <ul><li> **0**：未选择授权方式（默认值）</li>
	// <li> **1**：上传授权书</li>
	// <li> **2**：法人授权</li>
	// <li> **3**：法人认证</li></ul>
	AuthType *int64 `json:"authType"`

	// 企业认证授权书审核状态：
	// <ul><li> **0**：未提交授权书（默认值）</li>
	// <li> **1**：审核通过</li>
	// <li> **2**：审核驳回</li>
	// <li> **3**：审核中</li>
	// <li> **4**：AI识别中</li>
	// <li> **5**：客户确认AI信息</li></ul>
	AuditStatus *int64 `json:"auditStatus"`
}

type AuthorizedUser struct {
	// 电子签系统中的用户id
	UserId *string `json:"userId"`
}

type AutoSignConfig struct {
	// 自动签开通个人用户信息, 包括名字,身份证等
	UserInfo *UserThreeFactor `json:"userInfo"`

	// 是否回调证书信息:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**:需要</li></ul>
	CertInfoCallback *bool `json:"certInfoCallback"`

	// 是否支持用户自定义签名印章:
	// <ul><li>**false**: 不能自己定义(默认)</li>
	// <li>**true**: 可以自己定义</li></ul>
	UserDefineSeal *bool `json:"userDefineSeal"`

	// 回调中是否需要自动签将要使用的印章(签名) 图片的 base64:
	// <ul><li>**false**: 不需要(默认)</li>
	// <li>**true**: 需要</li></ul>
	SealImgCallback *bool `json:"sealImgCallback"`

	// 开通时候的身份验证方式, 取值为：
	// <ul><li>**WEIXINAPP** : 微信人脸识别</li>
	// <li>**INSIGHT** : 慧眼人脸认别</li>
	// <li>**TELECOM** : 运营商三要素验证</li></ul>
	// 注：
	// <ul><li>如果是小程序开通链接，支持传 WEIXINAPP / TELECOM。为空默认 WEIXINAPP</li>
	// <li>如果是 H5 开通链接，支持传 INSIGHT / TELECOM。为空默认 INSIGHT </li></ul>
	VerifyChannels []*string `json:"verifyChannels"`

	// 设置用户开通自动签时是否绑定个人自动签账号许可。
	//
	// <ul><li>**0**: (默认) 使用个人自动签账号许可进行开通，个人自动签账号许可有效期1年，注: `不可解绑释放更换他人`</li>
	// <li>**1**: 不绑定自动签账号许可开通，后续使用合同份额进行合同发起</li></ul>
	LicenseType *int64 `json:"licenseType"`

	// 开通成功后前端页面跳转的url，此字段的用法场景请联系客户经理确认。
	//
	// 注：`仅支持H5开通场景`, `跳转链接仅支持 https:// , qianapp:// 开头`
	//
	// 跳转场景：
	// <ul><li>**贵方H5 -> 腾讯电子签H5 -> 贵方H5** : JumpUrl格式: https://YOUR_CUSTOM_URL/xxxx，只需满足 https:// 开头的正确且合规的网址即可。</li>
	// <li>**贵方原生App -> 腾讯电子签H5 -> 贵方原生App** : JumpUrl格式: qianapp://YOUR_CUSTOM_URL，只需满足 qianapp:// 开头的URL即可。`APP实现方，需要拦截Webview地址跳转，发现url是qianapp:// 开头时跳转到原生页面。`APP拦截地址跳转可参考：<a href='https://stackoverflow.com/questions/41693263/android-webview-err-unknown-url-scheme'>Android</a>，<a href='https://razorpay.com/docs/payments/payment-gateway/web-integration/standard/webview/upi-intent-ios/'>IOS</a> </li></ul>
	//
	// 成功结果返回：
	// 若贵方需要在跳转回时通过链接query参数提示开通成功，JumpUrl中的query应携带如下参数：`appendResult=qian`。这样腾讯电子签H5会在跳转回的url后面会添加query参数提示贵方签署成功，例如： qianapp://YOUR_CUSTOM_URL?action=sign&result=success&from=tencent_ess
	JumpUrl *string `json:"jumpUrl"`
}

type BillUsageDetail struct {
	// 合同流程ID，为32位字符串。
	// 可登录腾讯电子签控制台，在 "合同"->"合同中心" 中查看某个合同的FlowId(在页面中展示为合同ID)。
	FlowId *string `json:"flowId"`

	// 合同经办人名称
	// 如果有多个经办人用分号隔开。
	OperatorName *string `json:"operatorName"`

	// 发起方组织机构名称
	CreateOrganizationName *string `json:"createOrganizationName"`

	// 合同流程的名称。
	FlowName *string `json:"flowName"`

	// 当前合同状态,如下是状态码对应的状态。
	// <ul>
	// <li>**0**: 还没有发起</li>
	// <li>**1**: 等待签署</li>
	// <li>**2**: 部分签署 </li>
	// <li>**3**: 拒签</li>
	// <li>**4**: 已签署 </li>
	// <li>**5**: 已过期 </li>
	// <li>**6**: 已撤销 </li>
	// <li>**7**: 还没有预发起</li>
	// <li>**8**: 等待填写</li>
	// <li>**9**: 部分填写 </li>
	// <li>**10**: 拒填</li>
	// <li>**11**: 已解除</li>
	// </ul>
	Status *int64 `json:"status"`

	// 查询的套餐类型
	// 对应关系如下:
	// <ul>
	// <li>**CloudEnterprise**: 企业版合同</li>
	// <li>**SingleSignature**: 单方签章</li>
	// <li>**CloudProve**: 签署报告</li>
	// <li>**CloudOnlineSign**: 腾讯会议在线签约</li>
	// <li>**ChannelWeCard**: 微工卡</li>
	// <li>**SignFlow**: 合同套餐</li>
	// <li>**SignFace**: 签署意愿（人脸识别）</li>
	// <li>**SignPassword**: 签署意愿（密码）</li>
	// <li>**SignSMS**: 签署意愿（短信）</li>
	// <li>**PersonalEssAuth**: 签署人实名（腾讯电子签认证）</li>
	// <li>**PersonalThirdAuth**: 签署人实名（信任第三方认证）</li>
	// <li>**OrgEssAuth**: 签署企业实名</li>
	// <li>**FlowNotify**: 短信通知</li>
	// <li>**AuthService**: 企业工商信息查询</li>
	// </ul>
	QuotaType *string `json:"quotaType"`

	// 合同使用量
	// 注: `如果消耗类型是撤销返还，此值为负值代表返还的合同数量`
	UseCount *int64 `json:"useCount"`

	// 消耗的时间戳，格式为Unix标准时间戳（秒）。
	CostTime *int64 `json:"costTime"`

	// 消耗的套餐名称
	QuotaName *string `json:"quotaName"`

	// 消耗类型
	// **1**.扣费
	// **2**.撤销返还
	CostType *int64 `json:"costType"`

	// 备注
	Remark *string `json:"remark"`
}

type CcInfo struct {
	// 被抄送方手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	// 请确认手机号所有方为此业务通知方。
	Mobile *string `json:"mobile"`

	// 被抄送方姓名。
	// 抄送方的姓名将用于身份认证，请确保填写的姓名为抄送方的真实姓名，而非昵称等代名。
	Name *string `json:"name"`

	// 被抄送方类型, 可设置以下类型:
	// <ul><li> **0** :个人抄送方</li>
	// <li> **1** :企业员工抄送方</li></ul>
	CcType *int64 `json:"ccType"`

	// 被抄送方权限, 可设置如下权限:
	// <ul><li> **0** :可查看合同内容</li>
	// <li> **1** :可查看合同内容也可下载原文</li></ul>
	CcPermission *int64 `json:"ccPermission"`

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li> **sms** :  (默认)短信</li>
	// <li> **none** : 不通知</li></ul>
	NotifyType *string `json:"notifyType"`
}

type Component struct {
	// **如果是Component填写控件类型，则可选的字段为**：
	//
	// <ul><li> <b>TEXT</b> : 普通文本控件，输入文本字符串；</li>
	// <li> <b>MULTI_LINE_TEXT</b> : 多行文本控件，输入文本字符串；</li>
	// <li> <b>CHECK_BOX</b> : 勾选框控件，若选中填写ComponentValue 填写 true或者 false 字符串；</li>
	// <li> <b>FILL_IMAGE</b> : 图片控件，ComponentValue 填写图片的资源 ID；</li>
	// <li> <b>DYNAMIC_TABLE</b> : 动态表格控件；</li>
	// <li> <b>ATTACHMENT</b> : 附件控件,ComponentValue 填写附件图片的资源 ID列表，以逗号分隔；</li>
	// <li> <b>SELECTOR</b> : 选择器控件，ComponentValue填写选择的字符串内容；</li>
	// <li> <b>DATE</b> : 日期控件；默认是格式化为xxxx年xx月xx日字符串；</li>
	// <li> <b>WATERMARK</b> : 水印控件；只能分配给发起方，必须设置ComponentExtra；</li>
	// <li> <b>DISTRICT</b> : 省市区行政区控件，ComponentValue填写省市区行政区字符串内容；</li></ul>
	//
	// **如果是SignComponent签署控件类型，
	// 需要根据签署人的类型可选的字段为**
	// * 企业方
	// <ul><li> <b>SIGN_SEAL</b> : 签署印章控件；</li>
	// <li> <b>SIGN_DATE</b> : 签署日期控件；</li>
	// <li> <b>SIGN_SIGNATURE</b> : 用户签名控件；</li>
	// <li> <b>SIGN_PAGING_SEAL</b> : 骑缝章；若文件发起，需要对应填充ComponentPosY、ComponentWidth、ComponentHeight</li>
	// <li> <b>SIGN_OPINION</b> : 签署意见控件，用户需要根据配置的签署意见内容，完成对意见内容的确认；</li>
	// <li> <b>SIGN_LEGAL_PERSON_SEAL</b> : 企业法定代表人控件。</li></ul>
	//
	// * 个人方
	// <ul><li> <b>SIGN_DATE</b> : 签署日期控件；</li>
	// <li> <b>SIGN_SIGNATURE</b> : 用户签名控件；</li></ul>
	//
	// 注：` 表单域的控件不能作为印章和签名控件`
	ComponentType *string `json:"componentType"`

	// **在绝对定位方式和关键字定位方式下**，指定控件的高度， 控件高度是指控件在PDF文件中的高度，单位为pt（点）。
	ComponentHeight *float64 `json:"componentHeight"`

	// **在绝对定位方式和关键字定位方式下**，指定控件宽度，控件宽度是指控件在PDF文件中的宽度，单位为pt（点）。
	ComponentWidth *float64 `json:"componentWidth"`

	// **在绝对定位方式方式下**，指定控件所在PDF文件上的页码
	// **在使用文件发起的情况下**，绝对定位方式的填写控件和签署控件支持使用负数来指定控件在PDF文件上的页码，使用负数时，页码从最后一页开始。例如：ComponentPage设置为-1，即代表在PDF文件的最后一页，以此类推。
	//
	// 注：
	// 1. 页码编号是从<font color="red">1</font>开始编号的。
	// 2.  <font color="red">页面编号不能超过PDF文件的页码总数</font>。如果指定的页码超过了PDF文件的页码总数，在填写和签署时会出现错误，导致无法正常进行操作。
	ComponentPage *int64 `json:"componentPage"`

	// **在绝对定位方式和关键字定位方式下**，可以指定控件横向位置的位置，单位为pt（点）。
	ComponentPosX *float64 `json:"componentPosX"`

	// **在绝对定位方式和关键字定位方式下**，可以指定控件纵向位置的位置，单位为pt（点）。
	ComponentPosY *float64 `json:"componentPosY"`

	// <font color="red">【暂未使用】</font>控件所属文件的序号（取值为：0-N）。 目前单文件的情况下，值一直为0
	FileIndex *int64 `json:"fileIndex"`

	// 控件生成的方式：
	// <ul><li> <b>NORMAL</b> : 绝对定位控件</li>
	// <li> <b>FIELD</b> : 表单域</li>
	// <li> <b>KEYWORD</b> : 关键字（设置关键字时，请确保PDF原始文件内是关键字以文字形式保存在PDF文件中，不支持对图片内文字进行关键字查找）</li></ul>
	GenerateMode *string `json:"generateMode"`

	// 控件唯一ID。
	//
	// **在绝对定位方式方式下**，ComponentId为控件的ID，长度不能超过30，只能由中文、字母、数字和下划线组成，可以在后续的操作中使用该名称来引用控件。
	//
	// **在关键字定位方式下**，ComponentId不仅为控件的ID，也是关键字整词。此方式下可以通过"^"来决定是否使用关键字整词匹配能力。
	//
	// 例：
	//
	// - 如传入的关键字<font color="red">"^甲方签署^"</font >，则会在PDF文件中有且仅有"甲方签署"关键字的地方（<font color="red">前后不能有其他字符</font >）进行对应操作。
	// - 如传入的关键字为<font color="red">"甲方签署</font >"，则PDF文件中每个出现关键字的位置（<font color="red">前后可以有其他字符</font >）都会执行相应操作。
	//
	//
	// 注：`控件ID可以在一个PDF中不可重复`
	//
	// <a href="https://qcloudimg.tencent-cloud.cn/raw/93178569d07b4d7dbbe0967ae679e35c.png" target="_blank">点击查看ComponentId在模板编辑页面的位置</a>
	ComponentId *string `json:"componentId"`

	// **在绝对定位方式方式下**，ComponentName为控件名，长度不能超过20，只能由中文、字母、数字和下划线组成，可以在后续的操作中使用该名称来引用控件。
	//
	// **在表单域定位方式下**，ComponentName不仅为控件名，也是表单域名称。
	//
	// 注：`控件名可以在一个PDF中可以重复`
	//
	// <a href="https://qcloudimg.tencent-cloud.cn/raw/93178569d07b4d7dbbe0967ae679e35c.png" target="_blank">点击查看ComponentName在模板页面的位置</a>
	ComponentName *string `json:"componentName"`

	// 如果是<b>填写控件</b>，ComponentRequired表示在填写页面此控件是否必填
	// <ul><li>false（默认）：可以不填写</li>
	// <li>true ：必须填写此填写控件</li></ul>
	// 如果是<b>签署控件</b>，签批控件中签署意见等可以不填写， 其他签署控件不受此字段影响
	ComponentRequired *bool `json:"componentRequired"`

	// **在通过接口拉取控件信息场景下**，为出参参数，此控件归属的参与方的角色ID角色（即RecipientId），**发起合同时候不要填写此字段留空即可**
	ComponentRecipientId *string `json:"componentRecipientId"`

	// **在所有的定位方式下**，控件的扩展参数，为<font color="red">JSON格式</font>，不同类型的控件会有部分非通用参数。
	//
	// <font color="red">ComponentType为TEXT、MULTI_LINE_TEXT时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体</li>
	// <li> <b>FontSize</b>： 范围12 :72</li>
	// <li> <b>FontAlign</b>： Left/Right/Center，左对齐/居中/右对齐</li>
	// <li> <b>FontColor</b>：字符串类型，格式为RGB颜色数字</li></ul>
	// <b>参数样例</b>：`{"FontColor":"255,0,0","FontSize":12}`
	//
	// <font color="red">ComponentType为DATE时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体</li>
	// <li> <b>FontSize</b>： 范围12 :72</li></ul>
	// <b>参数样例</b>：`{"FontColor":"255,0,0","FontSize":12}`
	//
	// <font color="red">ComponentType为WATERMARK时</font>，支持以下参数：
	// <ul><li> <b>Font</b>：目前只支持黑体、宋体</li>
	// <li> <b>FontSize</b>： 范围6 :24</li>
	// <li> <b>Opacity</b>： 透明度，范围0 :1</li>
	// <li> <b>Density</b>： 水印样式，1-宽松，2-标准（默认值），3-密集，</li>
	// <li> <b>SubType</b>： 水印类型：CUSTOM_WATERMARK-自定义内容，PERSON_INFO_WATERMARK-访问者信息</li></ul>
	// <b>参数样例</b>：`"{\"Font\":\"黑体\",\"FontSize\":20,\"Opacity\":0.1,\"Density\":2,\"SubType\":\"PERSON_INFO_WATERMARK\"}"`
	//
	// <font color="red">ComponentType为FILL_IMAGE时</font>，支持以下参数：
	// <ul><li> <b>NotMakeImageCenter</b>：bool。是否设置图片居中。false：居中（默认）。 true : 不居中</li>
	// <li> <b>FillMethod</b> : int. 填充方式。0-铺满（默认）；1-等比例缩放</li></ul>
	//
	// <font color="red">ComponentType为SIGN_SIGNATURE类型时</font>，可以通过**ComponentTypeLimit**参数控制签名方式
	// <ul><li> <b>HANDWRITE</b> :  需要实时手写的手写签名</li>
	// <li> <b>HANDWRITTEN_ESIGN</b> : 长效手写签名， 是使用保存到个人中心的印章列表的手写签名(并且包含HANDWRITE)</li>
	// <li> <b>OCR_ESIGN</b> : AI智能识别手写签名</li>
	// <li> <b>ESIGN</b> : 个人印章类型</li>
	// <li> <b>SYSTEM_ESIGN</b> : 系统签名（该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署）</li>
	// <li> <b>IMG_ESIGN</b> : 图片印章(该类型支持用户在签署将上传的PNG格式的图片作为签名)</li></ul>
	// <b>参考样例</b>：`{"ComponentTypeLimit": ["SYSTEM_ESIGN"]}`
	// 印章的对应关系参考下图
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/ee0498856c060c065628a0c5ba780d6b.jpg)<br><br>
	//
	// <font color="red">ComponentType为SIGN_SEAL 或者 SIGN_PAGING_SEAL类型时</font>，可以通过**ComponentTypeLimit**参数控制签署方签署时要使用的印章类型，支持指定以下印章类型
	// <ul><li> <b>OFFICIAL</b> :  企业公章</li>
	// <li> <b>CONTRACT</b> : 合同专用章</li>
	// <li> <b>FINANCE</b> : 财务专用章</li>
	// <li> <b>PERSONNEL</b> : 人事专用章</li></ul>
	// <b>参考样例</b>：`{\"ComponentTypeLimit\":[\"PERSONNEL\",\"FINANCE\"]}` 表示改印章签署区,客户需使用人事专用章或财务专用章盖章签署。<br><br>
	//
	// <font color="red">ComponentType为SIGN_DATE时</font>，支持以下参数：
	// <ul><li> <b>Font</b> :字符串类型目前只支持"黑体"、"宋体"，如果不填默认为"黑体"</li>
	// <li> <b>FontSize</b> : 数字类型，范围6-72，默认值为12</li>
	// <li> <b>FontAlign</b> : 字符串类型，可取Left/Right/Center，对应左对齐/居中/右对齐</li>
	// <li> <b>Format</b> : 字符串类型，日期格式，必须是以下五种之一 “yyyy m d”，”yyyy年m月d日”，”yyyy/m/d”，”yyyy-m-d”，”yyyy.m.d”。</li>
	// <li> <b>Gaps</b> : 字符串类型，仅在Format为“yyyy m d”时起作用，格式为用逗号分开的两个整数，例如”2,2”，两个数字分别是日期格式的前后两个空隙中的空格个数</li></ul>
	// 如果extra参数为空，默认为”yyyy年m月d日”格式的居中日期
	// 特别地，如果extra中Format字段为空或无法被识别，则extra参数会被当作默认值处理（Font，FontSize，Gaps和FontAlign都不会起效）
	// <b>参数样例</b>： ` "{"Format":"yyyy m d","FontSize":12,"Gaps":"2,2", "FontAlign":"Right"}"`
	//
	// <font color="red">ComponentType为SIGN_SEAL类型时</font>，支持以下参数：
	// <ul><li> <b>PageRanges</b> :PageRange的数组，通过PageRanges属性设置该印章在PDF所有页面上盖章（适用于标书在所有页面盖章的情况）</li></ul>
	// <b>参数样例</b>：` "{"PageRanges":[{"BeginPage":1,"EndPage":-1}]}"`
	//
	//
	// <font color="red">关键字模式下支持关键字找不到的情况下不进行报错的设置</font>
	// <ul><li> <b>IgnoreKeywordError</b> :1-关键字查找不到时不进行报错</li></ul>
	// 场景说明：如果使用关键字进行定位，但是指定的PDF文件中又没有设置的关键字时，发起合同会进行关键字是否存在的校验，如果关键字不存在，会进行报错返回。如果不希望进行报错，可以设置"IgnoreKeywordError"来忽略错误。请注意，如果关键字签署控件对应的签署方在整个PDF文件中一个签署控件都没有，还是会触发报错逻辑。
	// <b>参数样例</b>：` "{"IgnoreKeywordError":1}"`
	ComponentExtra *string `json:"componentExtra"`

	// **在通过接口拉取控件信息场景下**，为出参参数，此控件是否通过表单域定位方式生成，默认false-不是，**发起合同时候不要填写此字段留空即可**
	IsFormType *bool `json:"isFormType"`

	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// <ul><li> <b>TEXT</b> : 文本内容</li>
	// <li> <b>MULTI_LINE_TEXT</b> : 文本内容，可以用  \n 来控制换行位置 </li>
	// <li> <b>CHECK_BOX</b> : true/false</li>
	// <li> <b>FILL_IMAGE、ATTACHMENT</b> : 附件的FileId，需要通过UploadFiles接口上传获取</li>
	// <li> <b>SELECTOR</b> : 选项值</li>
	// <li> <b>DYNAMIC_TABLE</b>  - 传入json格式的表格内容，详见说明：[数据表格](https://qian.tencent.com/developers/company/dynamic_table)</li>
	// <li> <b>DATE</b> : 格式化为：xxxx年xx月xx日（例如2024年05年28日）</li>
	// <li> <b>SIGN_SEAL</b> : 印章ID，于控制台查询获取， [点击查看在控制台上位置](https://qcloudimg.tencent-cloud.cn/raw/f7b0f2ea4a534aada4b893dbf9671eae.png)</li>
	// <li> <b>SIGN_PAGING_SEAL</b> : 可以指定印章ID，于控制台查询获取， [点击查看在控制台上位置](https://qcloudimg.tencent-cloud.cn/raw/f7b0f2ea4a534aada4b893dbf9671eae.png)</li></ul>
	//
	//
	// <b>控件值约束说明</b>：
	// <table> <thead> <tr> <th>特殊控件</th> <th>填写约束</th> </tr> </thead> <tbody> <tr> <td>企业全称控件</td> <td>企业名称中文字符中文括号</td> </tr> <tr> <td>统一社会信用代码控件</td> <td>企业注册的统一社会信用代码</td> </tr> <tr> <td>法人名称控件</td> <td>最大50个字符，2到25个汉字或者1到50个字母</td> </tr> <tr> <td>签署意见控件</td> <td>签署意见最大长度为50字符</td> </tr> <tr> <td>签署人手机号控件</td> <td>国内手机号 13,14,15,16,17,18,19号段长度11位</td> </tr> <tr> <td>签署人身份证控件</td> <td>合法的身份证号码检查</td> </tr> <tr> <td>控件名称</td> <td>控件名称最大长度为20字符，不支持表情</td> </tr> <tr> <td>单行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>多行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>勾选框控件</td> <td>选择填字符串true，不选填字符串false</td> </tr> <tr> <td>选择器控件</td> <td>同单行文本控件约束，填写选择值中的字符串</td> </tr> <tr> <td>数字控件</td> <td>请输入有效的数字(可带小数点)</td> </tr> <tr> <td>日期控件</td> <td>格式：yyyy年mm月dd日</td> </tr> <tr> <td>附件控件</td> <td>JPG或PNG图片，上传数量限制，1到6个，最大6个附件，填写上传的资源ID</td> </tr> <tr> <td>图片控件</td> <td>JPG或PNG图片，填写上传的图片资源ID</td> </tr> <tr> <td>邮箱控件</td> <td>有效的邮箱地址, w3c标准</td> </tr> <tr> <td>地址控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>省市区控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>性别控件</td> <td>选择值中的字符串</td> </tr> <tr> <td>学历控件</td> <td>选择值中的字符串</td> </tr> <tr> <td>水印控件</td> <td>水印控件设置为CUSTOM_WATERMARK类型时的水印内容</td> </tr> </tbody> </table>
	// 注：   `部分特殊控件需要在控制台配置模板形式创建`
	ComponentValue *string `json:"componentValue"`

	// **如果控件是关键字定位方式**，可以对关键字定位出来的区域进行横坐标方向的调整，单位为pt（点）。例如，如果关键字定位出来的区域偏左或偏右，可以通过调整横坐标方向的参数来使控件位置更加准确。
	// 注意： `向左调整设置为负数， 向右调整设置成正数`
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetX *float64 `json:"offsetX"`

	// **如果控件是关键字定位方式**，可以对关键字定位出来的区域进行纵坐标方向的调整，单位为pt（点）。例如，如果关键字定位出来的区域偏上或偏下，可以通过调整纵坐标方向的参数来使控件位置更加准确。
	// 注意： `向上调整设置为负数， 向下调整设置成正数`
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetY *float64 `json:"offsetY"`

	// **如果控件是关键字定位方式**，指定关键字排序规则时，可以选择Positive或Reverse两种排序方式。
	// <ul><li> <b>Positive</b> :表示正序，即根据关键字在PDF文件内的顺序进行排列</li>
	// <li> <b>Reverse</b> :表示倒序，即根据关键字在PDF文件内的反序进行排列</li></ul>
	//
	// 在指定KeywordIndexes时，如果使用Positive排序方式，0代表在PDF内查找内容时，查找到的第一个关键字；如果使用Reverse排序方式，0代表在PDF内查找内容时，查找到的最后一个关键字。
	KeywordOrder *string `json:"keywordOrder"`

	// **如果控件是关键字定位方式**，在KeywordPage中指定关键字页码时，将只会在该页码中查找关键字，非该页码的关键字将不会查询出来。如果不设置查找所有页面中的关键字。
	KeywordPage *int64 `json:"keywordPage"`

	// **如果控件是关键字定位方式**，关键字生成的区域的对齐方式， 可以设置下面的值
	// <ul><li> <b>Middle</b> :居中</li>
	// <li> <b>Below</b> :正下方</li>
	// <li> <b>Right</b> :正右方</li>
	// <li> <b>LowerRight</b> :右下角</li>
	// <li> <b>UpperRight</b> :右上角。</li></ul>
	// 示例：如果设置Middle的关键字盖章，则印章的中心会和关键字的中心重合，如果设置Below，则印章在关键字的正下方
	RelativeLocation *string `json:"relativeLocation"`

	// **如果控件是关键字定位方式**，关键字索引是指在PDF文件中存在多个相同的关键字时，通过索引指定使用哪一个关键字作为最后的结果。可以通过指定多个索引来同时使用多个关键字。例如，[0,2]表示使用PDF文件内第1个和第3个关键字位置作为最后的结果。
	//
	// 注意：关键字索引是从0开始计数的
	KeywordIndexes []*int64 `json:"keywordIndexes"`

	// **web嵌入发起合同场景下**， 是否锁定填写和签署控件值不允许嵌入页面进行编辑
	// <ul><li>false（默认）：不锁定控件值，允许在页面编辑控件值</li>
	// <li>true：锁定控件值，在页面编辑控件值</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockComponentValue *bool `json:"lockComponentValue"`

	// **web嵌入发起合同场景下**，是否禁止移动和删除填写和签署控件
	// <ul><li> <b>false（默认）</b> :不禁止移动和删除控件</li>
	// <li> <b>true</b> : 可以移动和删除控件</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForbidMoveAndDelete *bool `json:"forbidMoveAndDelete"`

	// <font color="red">【暂未使用】</font>日期签署控件的字号，默认为 12
	ComponentDateFontSize *int64 `json:"componentDateFontSize"`

	// <font color="red">【暂未使用】</font>第三方应用集成平台模板控件 ID 标识
	ChannelComponentId *string `json:"channelComponentId"`

	// <font color="red">【暂未使用】</font>第三方应用集成中子客企业控件来源。
	// <ul><li> <b>0</b> :平台指定；</li>
	// <li> <b>1</b> :用户自定义</li></ul>
	ChannelComponentSource *uint64 `json:"channelComponentSource"`
}

type ComponentLimit struct {
	// 控件类型，支持以下类型
	// <ul><li>SIGN_SEAL : 印章控件</li>
	// <li>SIGN_PAGING_SEAL : 骑缝章控件</li>
	// <li>SIGN_LEGAL_PERSON_SEAL : 企业法定代表人控件</li>
	// <li>SIGN_SIGNATURE : 用户签名控件</li></ul>
	ComponentType *string `json:"componentType"`

	// 签署控件类型的值(可选)，用与限制签署时印章或者签名的选择范围
	//
	// 1.当ComponentType 是 SIGN_SEAL 或者 SIGN_PAGING_SEAL 时可传入企业印章Id（支持多个）或者以下印章类型
	//
	// <ul><li> <b>OFFICIAL</b> :  企业公章</li>
	// <li> <b>CONTRACT</b> : 合同专用章</li>
	// <li> <b>FINANCE</b> : 财务专用章</li>
	// <li> <b>PERSONNEL</b> : 人事专用章</li></ul>
	//
	// **注：`限制印章控件或骑缝章控件情况下,仅本企业签署方可以指定具体印章（通过传递ComponentValue,支持多个),他方企业签署人只能限制类型.若同时指定了印章类型和印章Id,以印章Id为主,印章类型会被忽略`**
	//
	//
	// 2.当ComponentType 是 SIGN_SIGNATURE 时可传入以下类型（支持多个）
	//
	// <ul><li>HANDWRITE : 需要实时手写的手写签名</li>
	// <li>HANDWRITTEN_ESIGN : 长效手写签名， 是使用保存到个人中心的印章列表的手写签名(并且包含HANDWRITE)</li>
	// <li>OCR_ESIGN : OCR印章（智慧手写签名）</li>
	// <li>ESIGN : 个人印章</li>
	// <li>SYSTEM_ESIGN : 系统印章</li></ul>
	//
	// 3.当ComponentType 是 SIGN_LEGAL_PERSON_SEAL 时无需传递此参数。
	ComponentValue []*string `json:"componentValue"`
}

type CreateFlowOption struct {
	// 是否允许修改发起合同时确认弹窗的合同信息（合同名称、合同类型、签署截止时间），若不允许编辑，则表单字段将被禁止输入。
	// <br/>true：允许编辑<br/>false：不允许编辑（默认值）<br/>
	CanEditFlow *bool `json:"canEditFlow"`

	// 是否允许编辑模板控件
	// <br/>true:允许编辑模板控件信息
	// <br/>false:不允许编辑模板控件信息（默认值）
	// <br/>
	CanEditFormField *bool `json:"canEditFormField"`

	// 发起页面隐藏合同名称展示
	// <br/>true:发起页面隐藏合同名称展示
	// <br/>false:发起页面不隐藏合同名称展示（默认值）
	// <br/>
	HideShowFlowName *bool `json:"hideShowFlowName"`

	// 发起页面隐藏合同类型展示
	// <br/>true:发起页面隐藏合同类型展示
	// <br/>false:发起页面不隐藏合同类型展示（默认值）
	// <br/>
	HideShowFlowType *bool `json:"hideShowFlowType"`

	// 发起页面隐藏合同截止日期展示
	// <br/>true:发起页面隐藏合同截止日期展示
	// <br/>false:发起页面不隐藏合同截止日期展示（默认值）
	// <br/>
	HideShowDeadline *bool `json:"hideShowDeadline"`

	// 发起页面允许跳过添加签署人环节
	// <br/>true:发起页面允许跳过添加签署人环节
	// <br/>false:发起页面不允许跳过添加签署人环节（默认值）
	// <br/>
	CanSkipAddApprover *bool `json:"canSkipAddApprover"`

	// 文件发起页面跳过文件上传步骤
	// <br/>true:文件发起页面跳过文件上传步骤
	// <br/>false:文件发起页面不跳过文件上传步骤（默认值）
	// <br/>
	SkipUploadFile *bool `json:"skipUploadFile"`

	// 禁止编辑填写控件
	// <br/>true:禁止编辑填写控件
	// <br/>false:允许编辑填写控件（默认值）
	// <br/>
	ForbidEditFillComponent *bool `json:"forbidEditFillComponent"`

	// 定制化发起合同弹窗的描述信息，描述信息最长500字符
	CustomCreateFlowDescription *string `json:"customCreateFlowDescription"`

	//   禁止添加签署方，若为true则在发起流程的可嵌入页面隐藏“添加签署人按钮”
	ForbidAddApprover *bool `json:"forbidAddApprover"`

	//   禁止设置设置签署流程属性 (顺序、合同签署认证方式等)，若为true则在发起流程的可嵌入页面隐藏签署流程设置面板
	ForbidEditFlowProperties *bool `json:"forbidEditFlowProperties"`

	// 在发起流程的可嵌入页面要隐藏的控件列表，和 ShowComponentTypes 参数 只能二选一使用，具体的控件类型如下
	// <ul><li>SIGN_SIGNATURE : 个人签名/印章</li>
	// <li>SIGN_SEAL : 企业印章</li>
	// <li>SIGN_PAGING_SEAL : 骑缝章</li>
	// <li>SIGN_LEGAL_PERSON_SEAL : 法定代表人章</li>
	// <li>SIGN_APPROVE : 签批</li>
	// <li>SIGN_OPINION : 签署意见</li>
	// <li>BUSI-FULL-NAME  : 企业全称</li>
	// <li>BUSI-CREDIT-CODE : 统一社会信用代码</li>
	// <li>BUSI-LEGAL-NAME : 法人/经营者姓名</li>
	// <li>PERSONAL-NAME : 签署人姓名</li>
	// <li>PERSONAL-MOBILE : 签署人手机号</li>
	// <li>PERSONAL-IDCARD-TYPE : 签署人证件类型</li>
	// <li>PERSONAL-IDCARD : 签署人证件号</li>
	// <li>TEXT : 单行文本</li>
	// <li>MULTI_LINE_TEXT : 多行文本</li>
	// <li>CHECK_BOX : 勾选框</li>
	// <li>SELECTOR : 选择器</li>
	// <li>DIGIT : 数字</li>
	// <li>DATE : 日期</li>
	// <li>FILL_IMAGE : 图片</li>
	// <li>ATTACHMENT : 附件</li>
	// <li>EMAIL : 邮箱</li>
	// <li>LOCATION : 地址</li>
	// <li>EDUCATION : 学历</li>
	// <li>GENDER : 性别</li>
	// <li>DISTRICT : 省市区</li></ul>
	HideComponentTypes []*string `json:"hideComponentTypes"`

	// 在发起流程的可嵌入页面要显示的控件列表，和 HideComponentTypes 参数 只能二选一使用，具体的控件类型如下
	// <ul><li>SIGN_SIGNATURE : 个人签名/印章</li>
	// <li>SIGN_SEAL : 企业印章</li>
	// <li>SIGN_PAGING_SEAL : 骑缝章</li>
	// <li>SIGN_LEGAL_PERSON_SEAL : 法定代表人章</li>
	// <li>SIGN_APPROVE : 签批</li>
	// <li>SIGN_OPINION : 签署意见</li>
	// <li>BUSI-FULL-NAME  : 企业全称</li>
	// <li>BUSI-CREDIT-CODE : 统一社会信用代码</li>
	// <li>BUSI-LEGAL-NAME : 法人/经营者姓名</li>
	// <li>PERSONAL-NAME : 签署人姓名</li>
	// <li>PERSONAL-MOBILE : 签署人手机号</li>
	// <li>PERSONAL-IDCARD-TYPE : 签署人证件类型</li>
	// <li>PERSONAL-IDCARD : 签署人证件号</li>
	// <li>TEXT : 单行文本</li>
	// <li>MULTI_LINE_TEXT : 多行文本</li>
	// <li>CHECK_BOX : 勾选框</li>
	// <li>SELECTOR : 选择器</li>
	// <li>DIGIT : 数字</li>
	// <li>DATE : 日期</li>
	// <li>FILL_IMAGE : 图片</li>
	// <li>ATTACHMENT : 附件</li>
	// <li>EMAIL : 邮箱</li>
	// <li>LOCATION : 地址</li>
	// <li>EDUCATION : 学历</li>
	// <li>GENDER : 性别</li>
	// <li>DISTRICT : 省市区</li></ul>
	ShowComponentTypes []*string `json:"showComponentTypes"`

	// 发起流程的可嵌入页面结果页配置
	ResultPageConfig []*CreateResultPageConfig `json:"resultPageConfig"`
}

type CreateResultPageConfig struct {
	// <ul>
	//   <li>0 : 发起审批成功页面（通过接口<a href="https://qian.tencent.com/developers/companyApis/embedPages/CreatePrepareFlow/" target="_blank">创建发起流程web页面</a>发起时设置了NeedCreateReview参数为true）</li>
	// </ul>
	Type *int64 `json:"type"`

	// 结果页标题，不超过50字
	Title *string `json:"title"`

	// 结果页描述，不超过200字
	Description *string `json:"description"`
}

type Department struct {
	// 部门ID。
	DepartmentId *string `json:"departmentId"`

	// 部门名称。
	DepartmentName *string `json:"departmentName"`
}

type EmbedUrlOption struct {
	// 合同详情预览，允许展示控件信息
	// <ul>
	// <li><b>true</b>：允许在合同详情页展示控件</li>
	// <li><b>false</b>：（默认）不允许在合同详情页展示控件</li>
	// </ul>
	ShowFlowDetailComponent *bool `json:"showFlowDetailComponent"`

	// 模板预览，允许展示模板控件信息
	// <ul><li> <b>true</b> :允许在模板预览页展示控件</li>
	// <li> <b>false</b> :（默认）不允许在模板预览页展示控件</li></ul>
	ShowTemplateComponent *bool `json:"showTemplateComponent"`
}

type ExtendAuthInfo struct {
	// 扩展服务的类型，可能是以下值：
	// <ul><li>OPEN_SERVER_SIGN：企业自动签署</li>
	// <li>BATCH_SIGN：批量签署</li>
	// <li>OVERSEA_SIGN：企业与港澳台居民签署合同</li>
	// <li>AGE_LIMIT_EXPANSION：拓宽签署方年龄限制</li>
	// <li>MOBILE_CHECK_APPROVER：个人签署方仅校验手机号</li>
	// <li>HIDE_OPERATOR_DISPLAY：隐藏合同经办人姓名</li>
	// <li>ORGANIZATION_OCR_FALLBACK：正楷临摹签名失败后更换其他签名类型</li>
	// <li>ORGANIZATION_FLOW_NOTIFY_TYPE：短信通知签署方</li>
	// <li>HIDE_ONE_KEY_SIGN：个人签署方手动签字</li>
	// <li>PAGING_SEAL：骑缝章</li>
	// <li>ORGANIZATION_FLOW_PASSWD_NOTIFY：签署密码开通引导</li></ul>
	Type *string `json:"type"`

	// 扩展服务的名称
	Name *string `json:"name"`

	// 扩展服务的开通状态：
	// <ul>
	// <li>ENABLE : 已开通</li>
	// <li>DISABLE : 未开通</li>
	// </ul>
	Status *string `json:"status"`

	// 操作扩展服务的操作人UserId，员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserId *string `json:"operatorUserId"`

	// 扩展服务的操作时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateOn *int64 `json:"operateOn"`

	// 该扩展服务若可以授权，此参数对应授权人员的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAuthUserList []*HasAuthUser `json:"hasAuthUserList"`
}

type ExtendScene struct {
	// 印章来源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenerateType *string `json:"generateType"`

	// 印章来源类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenerateTypeDesc *string `json:"generateTypeDesc"`

	// 印章来源logo
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenerateTypeLogo *string `json:"generateTypeLogo"`
}

type FailedCreateRoleData struct {
	// 用户userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"userId"`

	// 角色id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleIds []*string `json:"roleIds"`
}

type FailedCreateStaffData struct {
	// 员工名
	DisplayName *string `json:"displayName"`

	// 员工手机号
	Mobile *string `json:"mobile"`

	// 传入的企微账号id
	WeworkOpenId *string `json:"weworkOpenId"`

	// 失败原因
	Reason *string `json:"reason"`
}

type FailedDeleteStaffData struct {
	// 员工在电子签的userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"userId"`

	// 员工在第三方平台的openId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"openId"`

	// 失败原因
	Reason *string `json:"reason"`
}

type FailedUpdateStaffData struct {
	// 用户传入的名称
	DisplayName *string `json:"displayName"`

	// 用户传入的手机号，明文展示
	Mobile *string `json:"mobile"`

	// 失败原因
	Reason *string `json:"reason"`

	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	UserId *string `json:"userId"`

	// 员工在第三方平台的openId
	OpenId *string `json:"openId"`
}

type FileInfo struct {
	// 文件ID
	FileId *string `json:"fileId"`

	// 文件名
	FileName *string `json:"fileName"`

	// 文件大小，单位为Byte
	FileSize *int64 `json:"fileSize"`

	// 文件上传时间，格式为Unix标准时间戳（秒）
	CreatedOn *int64 `json:"createdOn"`
}

type FileUrl struct {
	// 下载文件的URL，有效期为输入的UrlTtl，默认5分钟
	Url *string `json:"url"`

	// 下载文件的附加信息。如果是pdf文件，会返回pdf文件每页的有效高宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	Option *string `json:"option"`
}

type FillApproverInfo struct {
	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	// 模板发起合同时，该参数为必填项。
	// 文件发起合同是，该参数无需传值。
	// 如果开发者后序用合同模板发起合同，建议保存此值，在用合同模板发起合同中需此值绑定对应的签署经办人 。
	RecipientId *string `json:"recipientId"`

	// 签署人来源
	// WEWORKAPP: 企业微信
	// <br/>仅【企微或签】时指定WEWORKAPP
	ApproverSource *string `json:"approverSource"`

	// 企业微信UserId
	// <br/>当ApproverSource为WEWORKAPP的企微或签场景下，必须指企业自有应用获取企业微信的UserId
	CustomUserId *string `json:"customUserId"`

	// 补充企业签署人员工姓名
	ApproverName *string `json:"approverName"`

	// 补充企业签署人员工手机号
	ApproverMobile *string `json:"approverMobile"`

	// 补充企业动态签署人时，需要指定对应企业名称
	OrganizationName *string `json:"organizationName"`

	// 签署方经办人的证件类型，支持以下类型
	// <ul><li>ID_CARD 中国大陆居民身份证</li>
	// <li>HONGKONG_AND_MACAO 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN 港澳台居民居住证(格式同居民身份证)</li>
	// <li>OTHER_CARD_TYPE 其他证件</li></ul>
	//
	// 注: `1.其他证件类型为白名单功能，使用前请联系对接的客户经理沟通。`
	// `2.补充个人签署方时，若该用户已在电子签完成实名则可通过指定姓名和证件类型、证件号码完成补充。`
	ApproverIdCardType *string `json:"approverIdCardType"`

	// 签署方经办人的证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给香港居民，“M”字头签发给澳门居民；第2位至第11位为数字。。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	//
	// 注：`补充个人签署方时，若该用户已在电子签完成实名则可通过指定姓名和证件类型、证件号码完成补充。`
	ApproverIdCardNumber *string `json:"approverIdCardNumber"`

	// 合同流程ID，补充合同组子合同动态签署人时必传。
	FlowId *string `json:"flowId"`
}

type FilledComponent struct {
	// 控件Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentId *string `json:"componentId"`

	// 控件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"componentName"`

	// 控件填写状态；0-未填写；1-已填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentFillStatus *string `json:"componentFillStatus"`

	// 控件填写内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentValue *string `json:"componentValue"`

	// 控件所属参与方Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentRecipientId *string `json:"componentRecipientId"`

	// 图片填充控件下载链接，如果是图片填充控件时，这里返回图片的下载链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"imageUrl"`
}

type FlowApproverDetail struct {
	// 签署时的相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMessage *string `json:"approveMessage"`

	// 签署方姓名
	ApproveName *string `json:"approveName"`

	// 签署方的签署状态
	// 0：还没有发起
	// 1：流程中 没有开始处理
	// 2：待签署
	// 3：已签署
	// 4：已拒绝
	// 5：已过期
	// 6：已撤销
	// 7：还没有预发起
	// 8：待填写
	// 9：因为各种原因而终止
	// 10：填写完成
	// 15：已解除
	// 19：转他人处理
	ApproveStatus *int64 `json:"approveStatus"`

	// 模板配置中的参与方ID,与控件绑定
	ReceiptId *string `json:"receiptId"`

	// 客户自定义的用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomUserId *string `json:"customUserId"`

	// 签署人手机号
	Mobile *string `json:"mobile"`

	// 签署顺序，如果是有序签署，签署顺序从小到大
	SignOrder *int64 `json:"signOrder"`

	// 签署人签署时间，时间戳，单位秒
	ApproveTime *int64 `json:"approveTime"`

	// 签署方类型，ORGANIZATION-企业员工，PERSON-个人，ENTERPRISESERVER-企业静默签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveType *string `json:"approveType"`

	// 签署方侧用户来源，如WEWORKAPP-企业微信等
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverSource *string `json:"approverSource"`

	// 客户自定义签署方标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomApproverTag *string `json:"customApproverTag"`

	// 签署方企业Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"organizationId"`

	// 签署方企业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"organizationName"`

	// 签署参与人在本流程中的编号ID（每个流程不同），可用此ID来定位签署参与人在本流程的签署节点，也可用于后续创建签署链接等操作。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"signId"`

	// 自定义签署人角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverRoleName *string `json:"approverRoleName"`
}

type FlowApproverUrlInfo struct {
	// 签署短链接。
	// 注意:
	// 1. 该链接有效期为<b>30分钟</b>，同时需要注意保密，不要外泄给无关用户。
	// 2. 该链接不支持小程序嵌入，仅支持<b>移动端浏览器</b>打开。
	// 3. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	SignUrl *string `json:"signUrl"`

	// 签署人类型。
	// - **1**: 个人
	ApproverType *int64 `json:"approverType"`

	// 签署人姓名。
	ApproverName *string `json:"approverName"`

	// 签署人手机号。
	ApproverMobile *string `json:"approverMobile"`

	// 签署长链接。
	// 注意:
	// 1. 该链接有效期为**30分钟**，同时需要注意保密，不要外泄给无关用户。
	// 2. 该链接不支持小程序嵌入，仅支持**移动端浏览器**打开。
	// 3. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）
	LongUrl *string `json:"longUrl"`
}

type FlowBatchApproverInfo struct {
	// 合同流程ID。
	FlowId *string `json:"flowId"`

	// 签署节点ID，用于生成动态签署人链接完成领取。注：`生成动态签署人补充链接时必传。`
	RecipientId *string `json:"recipientId"`
}

type FlowBatchUrlInfo struct {
	// 批量签署合同和签署方的信息，用于补充动态签署人。
	FlowBatchApproverInfos []*FlowBatchApproverInfo `json:"flowBatchApproverInfos"`
}

type FlowBrief struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"flowId"`

	// 合同流程的名称。
	FlowName *string `json:"flowName"`

	// 合同流程描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"flowDescription"`

	// 合同流程的类别分类（如销售合同/入职合同等）。
	FlowType *string `json:"flowType"`

	// 合同流程当前的签署状态, 会存在下列的状态值
	// <ul><li> **0** : 未开启流程(合同中不存在填写环节)</li>
	// <li> **1** : 待签署</li>
	// <li> **2** : 部分签署</li>
	// <li> **3** : 已拒签</li>
	// <li> **4** : 已签署</li>
	// <li> **5** : 已过期</li>
	// <li> **6** : 已撤销</li>
	// <li> **7** : 未开启流程(合同中存在填写环节)</li>
	// <li> **8** : 等待填写</li>
	// <li> **9** : 部分填写</li>
	// <li> **10** : 已拒填</li>
	// <li> **21** : 已解除</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowStatus *int64 `json:"flowStatus"`

	// 合同流程创建时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedOn *int64 `json:"createdOn"`

	// 当合同流程状态为已拒签（即 FlowStatus=3）或已撤销（即 FlowStatus=6）时，此字段 FlowMessage 为拒签或撤销原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"flowMessage"`

	//  合同流程发起方的员工编号, 即员工在腾讯电子签平台的唯一身份标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"creator"`

	// 合同流程的签署截止时间，格式为Unix标准时间戳（秒）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deadline *int64 `json:"deadline"`
}

type FlowCreateApprover struct {
	// 在指定签署方时，可以选择企业B端或个人C端等不同的参与者类型，可选类型如下：
	//
	// <ul><li> <b>0</b> :企业B端。</li>
	// <li> <b>1</b> :个人C端。</li>
	// <li> <b>3</b> :企业B端静默（自动）签署，无需签署人参与，自动签署可以参考<a href="https://qian.tencent.com/developers/company/autosign_guide" target="_blank" rel="noopener noreferrer">自动签署使用说明</a>文档。</li>
	// <li> <b>7</b> :个人C端自动签署，适用于个人自动签场景。注: <b>个人自动签场景为白名单功能，使用前请联系对接的客户经理沟通。</b> </li></ul>
	ApproverType *int64 `json:"approverType"`

	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	//
	// 注: `当approverType=0(企业签署方) 或 approverType=3(企业静默签署)时，必须指定`
	//
	OrganizationName *string `json:"organizationName"`

	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	//
	// 在未指定签署人电子签UserId情况下，为必填参数
	ApproverName *string `json:"approverName"`

	// 签署方经办人手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。 此手机号用于通知和用户的实名认证等环境，请确认手机号所有方为此合同签署方。
	//
	// 注：`在未指定签署人电子签UserId情况下，为必填参数`
	ApproverMobile *string `json:"approverMobile"`

	// 证件类型，支持以下类型
	// <ul><li><b>ID_CARD</b>: 居民身份证 (默认值)</li>
	// <li><b>HONGKONG_AND_MACAO</b> : 港澳居民来往内地通行证</li>
	// <li><b>HONGKONG_MACAO_AND_TAIWAN</b> : 港澳台居民居住证(格式同居民身份证)</li></ul>
	ApproverIdCardType *string `json:"approverIdCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给香港居民，“M”字头签发给澳门居民；第2位至第11位为数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	ApproverIdCardNumber *string `json:"approverIdCardNumber"`

	// 签署方经办人在模板中配置的参与方ID，与控件绑定，是控件的归属方，ID为32位字符串。
	//
	// <b>模板发起合同时，该参数为必填项，可以通过[查询模板信息接口](https://qian.tencent.com/developers/companyApis/templatesAndFiles/DescribeFlowTemplates)获得。</b>
	// <b>文件发起合同时，该参数无需传值。</b>
	//
	// 如果开发者后续用合同模板发起合同，建议保存此值，在用合同模板发起合同中需此值绑定对应的签署经办人 。
	RecipientId *string `json:"recipientId"`

	// 通知签署方经办人的方式,  有以下途径:
	// <ul><li>  **sms**  :  (默认)短信</li>
	// <li>   **none**   : 不通知</li></ul>
	//
	// 注: `既是发起方又是签署方时，不给此签署方发送短信`
	NotifyType *string `json:"notifyType"`

	// 合同强制需要阅读全文，无需传此参数
	IsFullText *bool `json:"isFullText"`

	// 签署方在签署合同之前，需要强制阅读合同的时长，可指定为3秒至300秒之间的任意值。
	//
	// 若未指定阅读时间，则会按照合同页数大小计算阅读时间，计算规则如下：
	// <ul>
	// <li>合同页数少于等于2页，阅读时间为3秒；</li>
	// <li>合同页数为3到5页，阅读时间为5秒；</li>
	// <li>合同页数大于等于6页，阅读时间为10秒。</li>
	// </ul>
	PreReadTime *uint64 `json:"preReadTime"`

	// 签署人userId，仅支持本企业的员工userid， 可在控制台组织管理处获得
	//
	// 注：
	// 如果传进来的<font color="red">UserId已经实名， 则忽略ApproverName，ApproverIdCardType，ApproverIdCardNumber，ApproverMobile这四个入参</font>（会用此UserId实名的身份证和登录的手机号覆盖）
	UserId *string `json:"userId"`

	// 在企微场景下使用，需设置参数为**WEWORKAPP**，以表明合同来源于企微。
	ApproverSource *string `json:"approverSource"`

	// 在企业微信场景下，表明该合同流程为或签，其最大长度为64位字符串。
	// 所有参与或签的人员均需具备该标识。
	// 注意，在合同中，不同的或签参与人必须保证其CustomApproverTag唯一。
	// 如果或签签署人为本方企业微信参与人，则需要指定ApproverSource参数为WEWORKAPP。
	CustomApproverTag *string `json:"customApproverTag"`

	// 签署人个性化能力值，如是否可以转发他人处理、是否可以拒签、是否为动态补充签署人等功能开关。
	ApproverOption *ApproverOption `json:"approverOption"`

	// 签署人的签署ID
	//
	// <ul>
	// <li>在CreateFlow、CreatePrepareFlow等发起流程时不需要传入此参数，电子签后台系统会自动生成。</li>
	// <li>在CreateFlowSignUrl、CreateBatchQuickSignUrl等生成签署链接时，可以通过查询详情接口获取签署人的SignId，然后可以将此值传入，为该签署人创建签署链接。这样可以避免重复传输姓名、手机号、证件号等其他信息。</li>
	// </ul>
	SignId *string `json:"signId"`

	// 发起方企业的签署人进行签署操作前，是否需要企业内部走审批流程，取值如下：
	// <ul><li>**false**：（默认）不需要审批，直接签署。</li>
	// <li>**true**：需要走审批流程。当到对应参与人签署时，会阻塞其签署操作，等待企业内部审批完成。</li></ul>
	// 企业可以通过CreateFlowSignReview审批接口通知腾讯电子签平台企业内部审批结果
	// <ul><li>如果企业通知腾讯电子签平台审核通过，签署方可继续签署动作。</li>
	// <li>如果企业通知腾讯电子签平台审核未通过，平台将继续阻塞签署方的签署动作，直到企业通知平台审核通过。</li></ul>
	//
	// 注：`此功能可用于与企业内部的审批流程进行关联，支持手动、静默签署合同`
	ApproverNeedSignReview *bool `json:"approverNeedSignReview"`

	// 签署人签署控件， 此参数仅针对文件发起（CreateFlowByFiles）生效
	//
	// 合同中的签署控件列表，列表中可支持下列多种签署控件,控件的详细定义参考开发者中心的Component结构体
	// <ul><li> 个人签名/印章</li>
	// <li> 企业印章</li>
	// <li> 骑缝章等签署控件</li></ul>
	//
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	SignComponents []*Component `json:"signComponents"`

	// 签署人填写控件 此参数仅针对文件发起（CreateFlowByFiles）生效
	//
	// 合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>动态表格等填写控件</li></ul>
	//
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	Components []*Component `json:"components"`

	// 当签署方控件类型为 <b>SIGN_SIGNATURE</b> 时，可以指定签署方签名方式。如果不指定，签署人可以使用所有的签名类型，可指定的签名类型包括：
	//
	// <ul><li> <b>HANDWRITE</b> :需要实时手写的手写签名。</li>
	// <li> <b>HANDWRITTEN_ESIGN</b> :长效手写签名， 是使用保存到个人中心的印章列表的手写签名。(并且包含HANDWRITE)</li>
	// <li> <b>OCR_ESIGN</b> :AI智能识别手写签名。</li>
	// <li> <b>ESIGN</b> :个人印章类型。</li>
	// <li> <b>IMG_ESIGN</b>  : 图片印章。该类型支持用户在签署将上传的PNG格式的图片作为签名。</li>
	// <li> <b>SYSTEM_ESIGN</b> :系统签名。该类型可以在用户签署时根据用户姓名一键生成一个签名来进行签署。</li></ul>
	//
	// 各种签名的样式可以参考下图：
	// ![image](https://qcloudimg.tencent-cloud.cn/raw/ee0498856c060c065628a0c5ba780d6b.jpg)
	ComponentLimitType []*string `json:"componentLimitType"`

	// 指定个人签署方查看合同的校验方式,可以传值如下:
	// <ul><li>  **1**   : （默认）人脸识别,人脸识别后才能合同内容</li>
	// <li>  **2**  : 手机号验证, 用户手机号和参与方手机号(ApproverMobile)相同即可查看合同内容（当手写签名方式为OCR_ESIGN时，该校验方式无效，因为这种签名方式依赖实名认证）
	// </li></ul>
	// 注:
	// <ul><li>如果合同流程设置ApproverVerifyType查看合同的校验方式,    则忽略此签署人的查看合同的校验方式</li>
	// <li>此字段可传多个校验方式</li></ul>
	//
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	//
	// .
	ApproverVerifyTypes []*int64 `json:"approverVerifyTypes"`

	// 您可以指定签署方签署合同的认证校验方式，可传递以下值：
	// <ul><li>**1**：人脸认证，需进行人脸识别成功后才能签署合同；</li>
	// <li>**2**：签署密码，需输入与用户在腾讯电子签设置的密码一致才能校验成功进行合同签署；</li>
	// <li>**3**：运营商三要素，需到运营商处比对手机号实名信息（名字、手机号、证件号）校验一致才能成功进行合同签署。（如果是港澳台客户，建议不要选择这个）</li>
	// <li>**5**：设备指纹识别，需要对比手机机主预留的指纹信息，校验一致才能成功进行合同签署。（iOS系统暂不支持该校验方式）</li>
	// <li>**6**：设备面容识别，需要对比手机机主预留的人脸信息，校验一致才能成功进行合同签署。（Android系统暂不支持该校验方式）</li></ul>
	//
	// 注：
	// <ul><li>默认情况下，认证校验方式为人脸认证和签署密码两种形式；</li>
	// <li>您可以传递多种值，表示可用多种认证校验方式。</li>
	// <li>校验方式不允许只包含设备指纹识别和设备面容识别，至少需要再增加一种其他校验方式。</li>
	// <li>设备指纹识别和设备面容识别只支持小程序使用，其他端暂不支持。</li></ul>
	//
	// 注:
	// `此参数仅针对文件发起设置生效,模板发起合同签署流程, 请以模板配置为主`
	ApproverSignTypes []*uint64 `json:"approverSignTypes"`

	// 生成H5签署链接时，您可以指定签署方签署合同的认证校验方式的选择模式，可传递一下值：
	// <ul><li>**0**：签署方自行选择，签署方可以从预先指定的认证方式中自由选择；</li>
	// <li>**1**：自动按顺序首位推荐，签署方无需选择，系统会优先推荐使用第一种认证方式。</li></ul>
	// 注：
	// `不指定该值时，默认为签署方自行选择。`
	SignTypeSelector *uint64 `json:"signTypeSelector"`

	// 签署人的签署截止时间，格式为Unix标准时间戳（秒）, 超过此时间未签署的合同变成已过期状态，不能在继续签署
	//
	// 注: `若不设置此参数，则默认使用合同的截止时间，此参数暂不支持合同组子合同`
	Deadline *int64 `json:"deadline"`

	// <b>只有在生成H5签署链接的情形下</b>（ 如调用<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowSignUrl" target="_blank">获取H5签署链接</a>、<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateBatchQuickSignUrl" target="_blank">获取H5批量签署链接</a>等接口），该配置才会生效。
	//
	// 您可以指定H5签署视频核身的意图配置，选择问答模式或点头模式的语音文本。
	//
	// 注意：
	// 1. 视频认证为<b>白名单功能，使用前请联系对接的客户经理沟通</b>。
	// 2. 使用视频认证时，<b>生成H5签署链接必须将签署认证方式指定为人脸</b>（即ApproverSignTypes设置成人脸签署）。
	// 3. 签署完成后，可以通过<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeSignFaceVideo" target="_blank">查询签署认证人脸视频</a>获取到当时的视频。
	Intention *Intention `json:"intention"`
}

type FlowDetailInfo struct {
	// 合同流程ID，为32位字符串。
	FlowId *string `json:"flowId"`

	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	FlowName *string `json:"flowName"`

	// 合同流程的类别分类（如销售合同/入职合同等）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowType *string `json:"flowType"`

	// 合同流程当前的签署状态, 会存在下列的状态值 <ul><li> **0** : 未开启流程(合同中不存在填写环节)</li> <li> **1** : 待签署</li> <li> **2** : 部分签署</li> <li> **3** : 已拒签</li> <li> **4** : 已签署</li> <li> **5** : 已过期</li> <li> **6** : 已撤销</li> <li> **7** : 未开启流程(合同中存在填写环节)</li> <li> **8** : 等待填写</li> <li> **9** : 部分填写</li> <li> **10** : 已拒填</li> <li> **21** : 已解除</li></ul>
	FlowStatus *int64 `json:"flowStatus"`

	// 当合同流程状态为已拒签（即 FlowStatus=3）或已撤销（即 FlowStatus=6）时，此字段 FlowMessage 为拒签或撤销原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMessage *string `json:"flowMessage"`

	// 合同流程描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDescription *string `json:"flowDescription"`

	// 合同流程的创建时间戳，格式为Unix标准时间戳（秒）。
	CreatedOn *int64 `json:"createdOn"`

	// 合同流程的签署方数组
	FlowApproverInfos []*FlowApproverDetail `json:"flowApproverInfos"`

	// 合同流程的关注方信息数组
	CcInfos []*FlowApproverDetail `json:"ccInfos"`

	// 合同流程发起方的员工编号, 即员工在腾讯电子签平台的唯一身份标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"reator"`
}

type FlowGroupApproverInfo struct {
	// 合同流程ID。
	FlowId *string `json:"flowId"`

	// 签署节点ID，用于生成动态签署人链接完成领取。注：`生成动态签署人补充链接时必传。`
	RecipientId *string `json:"fecipientId"`
}

type FlowGroupApprovers struct {
	// 合同流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"flowId"`

	// 签署方信息，包含合同ID和角色ID用于定位RecipientId。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Approvers []*ApproverItem `json:"approvers"`
}

type FlowGroupInfo struct {
	// 合同流程的名称（可自定义此名称），长度不能超过200，只能由中文、字母、数字和下划线组成。
	// 该名称还将用于合同签署完成后的下载文件名。
	FlowName *string `json:"flowName"`

	// 签署流程参与者信息，最大限制50方
	// 注意 approver中的顺序需要和模板中的顺序保持一致， 否则会导致模板中配置的信息无效。
	Approvers []*ApproverInfo `json:"approvers"`

	// 文件资源ID，通过多文件上传[UploadFiles](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)接口获得，为32位字符串。
	// 建议开发者保存此资源ID，后续创建合同或创建合同流程需此资源ID。
	FileIds []*string `json:"fileIds"`

	// 合同模板ID，为32位字符串。
	// 建议开发者保存此模板ID，后续用此模板发起合同流程需要此参数。
	// 可登录腾讯电子签控制台，在 "模板"->"模板中心"->"列表展示设置"选中模板 ID 中查看某个模板的TemplateId(在页面中展示为模板ID)。
	TemplateId *string `json:"templateId"`

	// 签署流程的类型(如销售合同/入职合同等)，最大长度200个字符
	// 示例值：劳务合同
	FlowType *string `json:"flowType"`

	// 签署流程描述,最大长度1000个字符
	FlowDescription *string `json:"flowDescription"`

	// 签署流程的签署截止时间。
	//
	// 值为unix时间戳,精确到秒,不传默认为当前时间一年后
	// 示例值：1604912664
	Deadline *int64 `json:"deadline"`

	// 调用方自定义的个性化字段(可自定义此字段的值)，并以base64方式编码，支持的最大数据大小为 20480长度。
	// 在合同状态变更的回调信息等场景中，该字段的信息将原封不动地透传给贵方。
	// 回调的相关说明可参考开发者中心的<a href="https://qian.tencent.com/developers/company/callback_types_v2" target="_blank">回调通知</a>模块。
	UserData *string `json:"userData"`

	// 发送类型：
	// true：无序签
	// false：有序签
	// 注：默认为false（有序签），请和模板中的配置保持一致
	// 示例值：true
	Unordered *bool `json:"unordered"`

	// 模板或者合同中的填写控件列表，列表中可支持下列多种填写控件，控件的详细定义参考开发者中心的Component结构体
	// <ul><li>单行文本控件</li>
	// <li>多行文本控件</li>
	// <li>勾选框控件</li>
	// <li>数字控件</li>
	// <li>图片控件</li>
	// <li>动态表格等填写控件</li></ul>
	Components []*Component `json:"components"`

	// 发起方企业的签署人进行签署操作是否需要企业内部审批。使用此功能需要发起方企业有参与签署。
	// 若设置为true，审核结果需通过接口 [CreateFlowSignReview](https://qian.tencent.com/developers/companyApis/operateFlows/CreateFlowSignReview) 通知电子签，审核通过后，发起方企业签署人方可进行签署操作，否则会阻塞其签署操作。
	//
	// 注：企业可以通过此功能与企业内部的审批流程进行关联，支持手动、静默签署合同。
	// 示例值：true
	NeedSignReview *bool `json:"needSignReview"`

	// 个人自动签场景。发起自动签署时，需设置对应自动签署场景，目前仅支持场景：处方单-E_PRESCRIPTION_AUTO_SIGN
	// 示例值：E_PRESCRIPTION_AUTO_SIGN
	AutoSignScene *string `json:"autoSignScene"`

	// 在短信通知、填写、签署流程中，若标题、按钮、合同详情等地方存在“合同”字样时，可根据此配置指定文案，可选文案如下：  <ul><li> <b>0</b> :合同（默认值）</li> <li> <b>1</b> :文件</li> <li> <b>2</b> :协议</li></ul>效果如下:![FlowDisplayType](https://qcloudimg.tencent-cloud.cn/raw/e4a2c4d638717cc901d3dbd5137c9bbc.png)
	FlowDisplayType *int64 `json:"flowDisplayType"`
}

type FlowGroupOptions struct {
	// 签署人校验方式,支持以下类型
	// <ul><li>VerifyCheck : 人脸识别 (默认值)</li>
	// <li>MobileCheck : 手机号验证</li></ul>
	// 参数说明：此参数仅在合同组文件发起有效，可选人脸识别或手机号验证两种方式，若选择后者，未实名个人签署方在签署合同时，无需经过实名认证和意愿确认两次人脸识别，该能力仅适用于个人签署方。
	ApproverVerifyType *string `json:"approverVerifyType"`

	// 发起合同（流程）组本方企业经办人通知方式
	// 签署通知类型，支持以下类型
	// <ul><li>sms : 短信 (默认值)</li><li>none : 不通知</li></ul>
	SelfOrganizationApproverNotifyType *string `json:"selfOrganizationApproverNotifyType"`

	// 发起合同（流程）组他方经办人通知方式
	// 签署通知类型，支持以下类型
	// <ul><li>sms : 短信 (默认值)</li><li>none : 不通知</li></ul>
	OtherApproverNotifyType *string `json:"otherApproverNotifyType"`
}

type FlowGroupUrlInfo struct {
	// 合同组子合同和签署方的信息，用于补充动态签署人。
	FlowGroupApproverInfos []*FlowGroupApproverInfo `json:"flowGroupApproverInfos"`
}

type FormField struct {
	// 控件填充vaule，ComponentType和传入值类型对应关系：
	// <ul><li> <b>TEXT</b> : 文本内容</li>
	// <li> <b>MULTI_LINE_TEXT</b> : 文本内容， 可以用  \n 来控制换行位置</li>
	// <li> <b>CHECK_BOX</b> : true/false</li>
	// <li> <b>FILL_IMAGE、ATTACHMENT</b> : 附件的FileId，需要通过UploadFiles接口上传获取</li>
	// <li> <b>SELECTOR</b> : 选项值</li>
	// <li> <b>DYNAMIC_TABLE</b>  - 传入json格式的表格内容，详见说明：[数据表格](https://qian.tencent.com/developers/company/dynamic_table)</li>
	// <li> <b>DATE</b> : 格式化：xxxx年xx月xx日（例如：2024年05月28日）</li>
	// </ul>
	//
	//
	// <b>控件值约束说明</b>：
	// <table> <thead> <tr> <th>特殊控件</th> <th>填写约束</th> </tr> </thead> <tbody> <tr> <td>企业全称控件</td> <td>企业名称中文字符中文括号</td> </tr> <tr> <td>统一社会信用代码控件</td> <td>企业注册的统一社会信用代码</td> </tr> <tr> <td>法人名称控件</td> <td>最大50个字符，2到25个汉字或者1到50个字母</td> </tr> <tr> <td>签署意见控件</td> <td>签署意见最大长度为50字符</td> </tr> <tr> <td>签署人手机号控件</td> <td>国内手机号 13,14,15,16,17,18,19号段长度11位</td> </tr> <tr> <td>签署人身份证控件</td> <td>合法的身份证号码检查</td> </tr> <tr> <td>控件名称</td> <td>控件名称最大长度为20字符，不支持表情</td> </tr> <tr> <td>单行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>多行文本控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>勾选框控件</td> <td>选择填字符串true，不选填字符串false</td> </tr> <tr> <td>选择器控件</td> <td>同单行文本控件约束，填写选择值中的字符串</td> </tr> <tr> <td>数字控件</td> <td>请输入有效的数字(可带小数点)</td> </tr> <tr> <td>日期控件</td> <td>格式：yyyy年mm月dd日</td> </tr> <tr> <td>附件控件</td> <td>JPG或PNG图片，上传数量限制，1到6个，最大6个附件，填写上传的资源ID</td> </tr> <tr> <td>图片控件</td> <td>JPG或PNG图片，填写上传的图片资源ID</td> </tr> <tr> <td>邮箱控件</td> <td>有效的邮箱地址, w3c标准</td> </tr> <tr> <td>地址控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>省市区控件</td> <td>只允许输入中文，英文，数字，中英文标点符号，不支持表情</td> </tr> <tr> <td>性别控件</td> <td>选择值中的字符串</td> </tr> <tr> <td>学历控件</td> <td>选择值中的字符串</td> </tr> </tbody> </table>
	ComponentValue *string `json:"componentValue"`

	// 控件id，和ComponentName选择一项传入即可
	//
	// <a href="https://dyn.ess.tencent.cn/guide/apivideo/component_name.mp4" target="_blank">点击查看在模板中找到控件ID的方式</a>
	ComponentId *string `json:"componentId"`

	// 控件名字，最大长度不超过30字符，和ComponentId选择一项传入即可
	//
	// <a href="https://dyn.ess.tencent.cn/guide/apivideo/component_name.mp4" target="_blank">点击查看在模板中找到控件名字的方式</a>
	ComponentName *string `json:"componentName"`
}

type GroupOrganization struct {
	// 成员企业名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"name"`

	// 成员企业别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"alias"`

	// 成员企业id，为 32 位字符串，可在电子签PC 控制台，企业设置->企业电子签账号 获取
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"organizationId"`

	// 记录更新时间， unix时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"updateTime"`

	// 成员企业加入集团的当前状态
	// <ul><li> **1**：待授权</li>
	// <li> **2**：已授权待激活</li>
	// <li> **3**：拒绝授权</li>
	// <li> **4**：已解除</li>
	// <li> **5**：已加入</li>
	// </ul>
	//
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"status"`

	// 是否为集团主企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMainOrganization *bool `json:"isMainOrganization"`

	// 企业社会信用代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardNumber *string `json:"idCardNumber"`

	// 企业超管信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminInfo *Admin `json:"adminInfo"`

	// 企业许可证Id，此字段暂时不需要关注
	// 注意：此字段可能返回 null，表示取不到有效值。
	License *string `json:"license"`

	// 企业许可证过期时间，unix时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseExpireTime *uint64 `json:"licenseExpireTime"`

	// 成员企业加入集团时间，unix时间戳，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinTime *uint64 `json:"joinTime"`

	// 是否使用自建审批流引擎（即不是企微审批流引擎）
	// <ul><li> **true**：是</li>
	// <li> **false**：否</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowEngineEnable *bool `json:"flowEngineEnable"`
}

type HasAuthOrganization struct {
	// 授权企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationId *string `json:"organizationId"`

	// 授权企业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"organizationName"`

	// 被授权企业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedOrganizationId *string `json:"authorizedOrganizationId"`

	// 被授权企业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedOrganizationName *string `json:"authorizedOrganizationName"`

	// 授权模板id（仅当授权方式为模板授权时有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"templateId"`

	// 授权模板名称（仅当授权方式为模板授权时有值）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"templateName"`

	// 授权时间，格式为时间戳，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizeTime *int64 `json:"authorizeTime"`
}

type HasAuthUser struct {
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"userId"`

	// 当前员工的归属情况，可能值是：
	// MainOrg：在集团企业的场景下，返回此值代表是归属主企业
	// CurrentOrg：在普通企业场景下返回此值；或者在集团企业的场景下，返回此值代表归属子企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongTo *string `json:"belongTo"`

	// 集团主企业id，当前企业为集团子企业时，该字段有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainOrganizationId *string `json:"mainOrganizationId"`
}

type IntegrateRole struct {
	// 角色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"roleId"`

	// 角色名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"roleName"`

	// 角色状态，1-启用，2-禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleStatus *uint64 `json:"roleStatus"`

	// 是否是集团角色，true-是，false-否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGroupRole *bool `json:"isGroupRole"`

	// 管辖的子企业列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubOrgIdList []*string `json:"subOrgIdList"`

	// 权限树
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionGroups []*PermissionGroup `json:"permissionGroups"`
}

type IntegrationDepartment struct {
	// 部门ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptId *string `json:"deptId"`

	// 部门名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptName *string `json:"deptName"`

	// 父部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentDeptId *string `json:"parentDeptId"`

	// 客户系统部门ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeptOpenId *string `json:"deptOpenId"`

	// 序列号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *uint64 `json:"orderNo"`
}

type Intention struct {
	// 视频认证类型，支持以下类型
	// <ul><li>1 : 问答模式</li>
	// <li>2 : 点头模式</li></ul>
	//
	// 注: `视频认证为白名单功能，使用前请联系对接的客户经理沟通。`
	IntentionType *int64 `json:"intentionType"`

	// 意愿核身语音问答模式（即语音播报+语音回答）使用的文案，包括：系统语音播报的文本、需要核验的标准文本。当前仅支持1轮问答。
	//
	// 注：`选择问答模式时，此字段可不传，不传则使用默认语音文本：请问，您是否同意签署本协议？可语音回复“同意”或“不同意”。`
	IntentionQuestions []*IntentionQuestion `json:"intentionQuestions"`

	// 意愿核身（点头确认模式）使用的文案，若未使用意愿核身（点头确认模式），则该字段无需传入。当前仅支持一个提示文本。
	//
	// 注：`选择点头模式时，此字段可不传，不传则使用默认语音文本：请问，您是否同意签署本协议？可点头同意。`
	IntentionActions []*IntentionAction `json:"intentionActions"`
}

type IntentionAction struct {
	// 点头确认模式下，系统语音播报使用的问题文本，问题最大长度为150个字符。
	Text *string `json:"text"`
}

type IntentionQuestion struct {
	// 当选择语音问答模式时，系统自动播报的问题文本，最大长度为150个字符。
	Question *string `json:"question"`

	//  当选择语音问答模式时，用于判断用户回答是否通过的标准答案列表，传入后可自动判断用户回答文本是否在标准文本列表中。
	Answers []*string `json:"answers"`
}

type OccupiedSeal struct {
	// 电子印章编号
	SealId *string `json:"sealId"`

	// 电子印章名称
	SealName *string `json:"sealName"`

	// 电子印章授权时间戳，单位秒
	CreateOn *int64 `json:"createOn"`

	// 电子印章授权人的UserId
	Creator *string `json:"creator"`

	// 电子印章策略Id
	SealPolicyId *string `json:"sealPolicyId"`

	// 印章状态，有以下六种：CHECKING（审核中）SUCCESS（已启用）FAIL（审核拒绝）CHECKING-SADM（待超管审核）DISABLE（已停用）STOPPED（已终止）
	SealStatus *string `json:"sealStatus"`

	// 审核失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"failReason"`

	// 印章图片url，5分钟内有效
	Url *string `json:"url"`

	// 印章类型,OFFICIAL-企业公章, CONTRACT-合同专用章,ORGANIZATIONSEAL-企业印章(本地上传印章类型),LEGAL_PERSON_SEAL-法人印章
	SealType *string `json:"sealType"`

	// 用印申请是否为永久授权，true-是，false-否
	IsAllTime *bool `json:"isAllTime"`

	// 授权人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizedUsers []*AuthorizedUser `json:"authorizedUsers"`

	// 印章扩展数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendScene *ExtendScene `json:"extendScene"`
}

type OrgBillSummary struct {
	// 套餐总数
	Total *int64 `json:"total"`

	// 套餐使用数
	Used *int64 `json:"used"`

	// 套餐剩余数
	Available *int64 `json:"available"`

	// 套餐类型
	// 对应关系如下:
	// <ul>
	// <li>**CloudEnterprise**: 企业版合同</li>
	// <li>**SingleSignature**: 单方签章</li>
	// <li>**CloudProve**: 签署报告</li>
	// <li>**CloudOnlineSign**: 腾讯会议在线签约</li>
	// <li>**ChannelWeCard**: 微工卡</li>
	// <li>**SignFlow**: 合同套餐</li>
	// <li>**SignFace**: 签署意愿（人脸识别）</li>
	// <li>**SignPassword**: 签署意愿（密码）</li>
	// <li>**SignSMS**: 签署意愿（短信）</li>
	// <li>**PersonalEssAuth**: 签署人实名（腾讯电子签认证）</li>
	// <li>**PersonalThirdAuth**: 签署人实名（信任第三方认证）</li>
	// <li>**OrgEssAuth**: 签署企业实名</li>
	// <li>**FlowNotify**: 短信通知</li>
	// <li>**AuthService**: 企业工商信息查询</li>
	// </ul>
	QuotaType *string `json:"auotaType"`
}

type OrganizationInfo struct {
}

type Permission struct {
	// 权限名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"name"`

	// 权限key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"key"`

	// 权限类型 1前端，2后端
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"type"`

	// 是否隐藏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hide *int64 `json:"hide"`

	// 数据权限标签 1:表示根节点，2:表示叶子结点
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataLabel *int64 `json:"dataLabel"`

	// 数据权限独有，1:关联其他模块鉴权，2:表示关联自己模块鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *int64 `json:"dataType"`

	// 数据权限独有，表示数据范围，1：全公司，2:部门及下级部门，3:自己
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRange *int64 `json:"dataRange"`

	// 关联权限, 表示这个功能权限要受哪个数据权限管控
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTo *string `json:"dataTo"`

	// 父级权限key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentKey *string `json:"parentKey"`

	// 是否选中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsChecked *bool `json:"isChecked"`

	// 子权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*Permission `json:"children"`
}

type PermissionGroup struct {
	// 权限组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"groupName"`

	// 权限组key
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupKey *string `json:"groupKey"`

	// 是否隐藏分组，0否1是
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hide *int64 `json:"hide"`

	// 权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*Permission `json:"permissions"`
}

type Recipient struct {
	// 签署参与者ID，唯一标识
	RecipientId *string `json:"recipientId"`

	// 参与者类型。
	// 默认为空。
	// ENTERPRISE-企业；
	// INDIVIDUAL-个人；
	// PROMOTER-发起方
	RecipientType *string `json:"recipientType"`

	// 描述信息
	Description *string `json:"description"`

	// 角色名称
	RoleName *string `json:"roleName"`

	// 是否需要验证，
	// 默认为false-不需要验证
	RequireValidation *bool `json:"requireValidation"`

	// 是否需要签署，
	// 默认为true-需要签署
	RequireSign *bool `json:"requireSign"`

	// 此参与方添加的顺序，从0～N
	RoutingOrder *int64 `json:"routingOrder"`

	// 是否需要发送，
	// 默认为true-需要发送
	RequireDelivery *bool `json:"requireDelivery"`

	// 邮箱地址
	Email *string `json:"email"`

	// 电话号码
	Mobile *string `json:"mobile"`

	// 关联的用户ID，电子签系统的用户ID
	UserId *string `json:"userId"`

	// 发送方式，默认为EMAIL。
	// EMAIL-邮件；
	// MOBILE-手机短信；
	// WECHAT-微信通知
	DeliveryMethod *string `json:"deliveryMethod"`

	// 参与方的一些附属信息，json格式
	RecipientExtra *string `json:"recipientExtra"`

	// 签署人查看合同校验方式, 支持的类型如下:
	// <ul><li> 1 :实名认证查看</li>
	// <li> 2 :手机号校验查看</li></ul>
	ApproverVerifyTypes []*int64 `json:"approverVerifyTypes"`

	// 签署人进行合同签署时的认证方式，支持的类型如下:
	// <ul><li> 1 :人脸认证</li>
	// <li> 2 :签署密码</li>
	// <li> 3 :运营商三要素认证</li>
	// <li> 4 :UKey认证</li>
	// <li> 5 :设备指纹识别</li>
	// <li> 6 :设备面容识别</li></ul>
	ApproverSignTypes []*int64 `json:"approverSignTypes"`

	// 签署方是否可以转他人处理
	//
	// <ul><li> **false** : ( 默认)可以转他人处理</li>
	// <li> **true** :不可以转他人处理</li></ul>
	NoTransfer *bool `json:"noTransfer"`
}

type RecipientComponentInfo struct {
	// 签署方经办人在合同流程中的参与方ID，与控件绑定，是控件的归属方
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientId *string `json:"recipientId"`

	// 参与方填写状态
	// <ul>
	// <li>**空值** : 此参与方没有填写控件</li>
	// <li>**0**:  未填写, 表示此参与方还没有填写合同的填写控件</li>
	// <li>**1**:  已填写, 表示此参与方已经填写所有的填写控件</li></ul>
	//
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecipientFillStatus *string `json:"recipientFillStatus"`

	// 是否为发起方
	// <ul><li>true-发起方</li>
	// <li>false-参与方</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPromoter *bool `json:"isPromoter"`

	// 改参与方填写控件信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*FilledComponent `json:"components"`
}

type RegisterInfo struct {
	// 法人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalName *string `json:"legalName"`

	// 社会统一信用代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnifiedSocialCreditCode *string `json:"unifiedSocialCreditCode"`
}

type RegistrationOrganizationInfo struct {
	// 组织机构名称。
	// 请确认该名称与企业营业执照中注册的名称一致。
	// 如果名称中包含英文括号()，请使用中文括号（）代替。
	OrganizationName *string `json:"organizationName"`

	// 组织机构企业统一社会信用代码。
	// 请确认该企业统一社会信用代码与企业营业执照中注册的统一社会信用代码一致。
	UniformSocialCreditCode *string `json:"uniformSocialCreditCode"`

	// 组织机构法人的姓名。
	// 请确认该企业统一社会信用代码与企业营业执照中注册的法人姓名一致。
	LegalName *string `json:"legalName"`

	// 组织机构企业注册地址。
	// 请确认该企业注册地址与企业营业执照中注册的地址一致。
	Address *string `json:"address"`

	// 组织机构超管姓名。
	// 在注册流程中，必须是超管本人进行操作。
	// 如果法人做为超管管理组织机构,超管姓名就是法人姓名
	// 如果入参中传递超管授权书PowerOfAttorneys，则此参数为必填参数。
	AdminName *string `json:"adminName"`

	// 组织机构超管手机号。
	// 在注册流程中，这个手机号必须跟操作人在电子签注册的个人手机号一致。
	// 如果入参中传递超管授权书PowerOfAttorneys，则此参数为必填参数
	AdminMobile *string `json:"adminMobile"`

	// 可选的此企业允许的授权方式, 可以设置的方式有:
	// 1：上传授权书
	// 2：法人授权超管
	// 5：授权书+对公打款
	//
	//
	// 注:
	// `1. 当前仅支持一种认证方式`
	// `2. 如果当前的企业类型是政府/事业单位, 则只支持上传授权书+对公打款`
	// `3. 如果当前操作人是法人,则是法人认证`
	AuthorizationTypes []*uint64 `json:"authorizationTypes"`

	// 认证人身份证号，如果入参中传递超管授权书PowerOfAttorneys，则此参数为必填参数
	AdminIdCardNumber *string `json:"adminIdCardNumber"`

	// 认证人证件类型
	// 支持以下类型
	// <ul><li>ID_CARD : 居民身份证  (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	AdminIdCardType *string `json:"adminIdCardType"`

	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	BusinessLicense *string `json:"businessLicense"`

	// 授权书(PNG或JPG或PDF) base64格式, 大小不超过8M 。
	// p.s. 如果上传授权书 ，需遵循以下条件
	// 1. 超管的信息（超管姓名，超管身份证，超管手机号）必须为必填参数。
	// 2. 超管的个人身份必须在电子签已经实名。
	// 2. 认证方式AuthorizationTypes必须只能是上传授权书方式
	PowerOfAttorneys []*string `json:"powerOfAttorneys"`
}

type ReleasedApprover struct {
	// 签署人姓名，最大长度50个字。
	Name *string `json:"name"`

	// 签署人手机号。
	Mobile *string `json:"mobile"`

	// 要更换的原合同参与人RecipientId编号。(可通过接口<a href="https://qian.tencent.com/developers/companyApis/queryFlows/DescribeFlowInfo/">DescribeFlowInfo</a>查询签署人的RecipientId编号)<br/>
	RelievedApproverReceiptId *string `json:"relievedApproverReceiptId"`

	// 指定签署人类型，目前仅支持
	// <ul><li> **ORGANIZATION**：企业（默认值）</li>
	// <li> **ENTERPRISESERVER**：企业静默签</li></ul>
	ApproverType *string `json:"approverType"`

	// 签署控件类型，支持自定义企业签署方的签署控件类型
	// <ul><li> **SIGN_SEAL**：默认为印章控件类型（默认值）</li>
	// <li> **SIGN_SIGNATURE**：手写签名控件类型</li></ul>
	ApproverSignComponentType *string `json:"approverSignComponentType"`

	// 参与方在合同中的角色是按照创建合同的时候来排序的，解除协议默认会将第一个参与人叫`甲方`,第二个叫`乙方`,  第三个叫`丙方`，以此类推。
	//
	// 如果需改动此参与人的角色名字，可用此字段指定，由汉字,英文字符,数字组成，最大20个字。
	ApproverSignRole *string `json:"approverSignRole"`

	// 印章Id，签署控件类型为印章时，用于指定本企业签署方在解除协议中使用那个印章进行签署
	ApproverSignSealId *string `json:"approverSignSealId"`
}

type RelieveInfo struct {
	// 解除理由，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	Reason *string `json:"reason"`

	// 解除后仍然有效的条款，保留条款，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	RemainInForceItem *string `json:"remainInForceItem"`

	// 原合同事项处理-费用结算，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	OriginalExpenseSettlement *string `json:"originalExpenseSettlement"`

	// 原合同事项处理-其他事项，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	OriginalOtherSettlement *string `json:"originalOtherSettlement"`

	// 其他约定，长度不能超过200，只能由中文、字母、数字、中文标点和英文标点组成(不支持表情)。
	OtherDeals *string `json:"otherDeals"`
}

type RemindFlowRecords struct {
	// 合同流程是否可以催办：
	// true - 可以，false - 不可以。
	// 若无法催办，将返回RemindMessage以解释原因。
	CanRemind *bool `json:"canRemind"`

	// 合同流程ID，为32位字符串。
	FlowId *string `json:"flowId"`

	// 在合同流程无法催办的情况下，系统将返回RemindMessage以阐述原因。
	RemindMessage *string `json:"remindMessage"`
}

type ReviewerInfo struct {
	// 姓名
	Name *string `json:"name"`

	// 手机号
	Mobile *string `json:"mobile"`
}

type SealInfo struct {
	// 印章ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealId *string `json:"sealId"`

	// 印章类型。LEGAL_PERSON_SEAL: 法定代表人章；
	// ORGANIZATIONSEAL：企业印章；
	// OFFICIAL：企业公章；
	// CONTRACT：合同专用章
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealType *string `json:"sealType"`

	// 印章名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SealName *string `json:"sealName"`
}

type SignQrCode struct {
	// 二维码ID，为32位字符串。
	QrCodeId *string `json:"qrCodeId"`

	// 二维码URL，可通过转换二维码的工具或代码组件将此URL转化为二维码，以便用户扫描进行流程签署。
	QrCodeUrl *string `json:"qrCodeUrl"`

	// 二维码的有截止时间，格式为Unix标准时间戳（秒）。
	// 一旦超过二维码的有效期限，该二维码将自动失效。
	ExpiredTime *int64 `json:"expiredTime"`
}

type SignUrl struct {
	// 跳转至电子签名小程序签署的链接地址。
	// 适用于客户端APP及小程序直接唤起电子签名小程序。
	AppSignUrl *string `json:"appSignUrl"`

	// 签署链接有效时间，格式类似"2022-08-05 15:55:01"
	EffectiveTime *string `json:"effectiveTime"`

	// 跳转至电子签名小程序签署的链接地址，格式类似于https://essurl.cn/xxx。
	// 打开此链接将会展示H5中间页面，随后唤起电子签名小程序以进行合同签署。
	HttpSignUrl *string `json:"httpSignUrl"`
}

type Staff struct {
	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 注：`创建和更新场景无需填写。`
	UserId *string `json:"userId"`

	// 显示的用户名/昵称。
	DisplayName *string `json:"displayName"`

	// 用户手机号码， 支持国内手机号11位数字(无需加+86前缀或其他字符)。
	Mobile *string `json:"mobile"`

	// 用户邮箱。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"email"`

	// 用户在第三方平台ID。
	// 注：`如需在此接口提醒员工实名，该参数不传。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"openId"`

	// 员工角色信息。
	// 注：`创建和更新场景无需填写。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*StaffRole `json:"roles"`

	// 员工部门信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"department"`

	// 员工是否实名。
	// 注：`创建和更新场景无需填写。`
	Verified *bool `json:"verified"`

	// 员工创建时间戳，单位秒。
	// 注：`创建和更新场景无需填写。`
	CreatedOn *int64 `json:"createdOn"`

	// 员工实名时间戳，单位秒。
	// 注：`创建和更新场景无需填写。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifiedOn *int64 `json:"verifiedOn"`

	// 员工是否离职：
	// <ul><li>**0**：未离职</li><li>**1**：离职</li></ul>
	// 注：`创建和更新场景无需填写。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuiteJob *int64 `json:"quiteJob"`

	// 员工离职交接人用户ID。
	// 注：`创建和更新场景无需填写。`
	ReceiveUserId *string `json:"receiveUserId"`

	// 员工离职交接人用户OpenId。
	// 注：`创建和更新场景无需填写。`
	ReceiveOpenId *string `json:"receiveOpenId"`

	// 企业微信用户账号ID。
	// 注：`仅企微类型的企业创建员工接口支持该字段。`
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeworkOpenId *string `json:"weworkOpenId"`
}

type StaffRole struct {
	// 角色ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleId *string `json:"roleId"`

	// 角色名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"roleName"`
}

type SubOrgBillSummary struct {
	// 子企业名称
	OrganizationName *string `json:"organizationName"`

	//
	Usage []*SubOrgBillUsage `json:"usage"`
}

type SubOrgBillUsage struct {
	// 套餐使用数
	Used *int64 `json:"used"`

	// 套餐类型
	// 对应关系如下:
	// <ul>
	// <li>**CloudEnterprise**: 企业版合同</li>
	// <li>**SingleSignature**: 单方签章</li>
	// <li>**CloudProve**: 签署报告</li>
	// <li>**CloudOnlineSign**: 腾讯会议在线签约</li>
	// <li>**ChannelWeCard**: 微工卡</li>
	// <li>**SignFlow**: 合同套餐</li>
	// <li>**SignFace**: 签署意愿（人脸识别）</li>
	// <li>**SignPassword**: 签署意愿（密码）</li>
	// <li>**SignSMS**: 签署意愿（短信）</li>
	// <li>**PersonalEssAuth**: 签署人实名（腾讯电子签认证）</li>
	// <li>**PersonalThirdAuth**: 签署人实名（信任第三方认证）</li>
	// <li>**OrgEssAuth**: 签署企业实名</li>
	// <li>**FlowNotify**: 短信通知</li>
	// <li>**AuthService**: 企业工商信息查询</li>
	// </ul>
	QuotaType *string `json:"quotaType"`
}

type SuccessCreateStaffData struct {
	// 员工名
	DisplayName *string `json:"displayName"`

	// 员工手机号
	Mobile *string `json:"mobile"`

	// 员工在电子签平台的id
	UserId *string `json:"userId"`

	// 提示，当创建已存在未实名用户时，该字段有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Note *string `json:"note"`

	// 传入的企微账号id
	WeworkOpenId *string `json:"weworkOpenId"`

	// 员工邀请返回链接 根据入参的 InvitationNotifyType 和 Endpoint 返回链接 <table><tbody><tr><td>链接类型</td><td>有效期</td><td>示例</td></tr><tr><td>HTTP_SHORT_URL（短链）</td><td>一天</td><td>https://test.essurl.cn/fvG7UBEd0F</td></tr><tr><td>HTTP（长链）</td><td>一天</td><td>https://res.ess.tencent.cn/cdn/h5-activity-dev/jump-mp.html?where=mini&from=MSG&to=USER_VERIFY&verifyToken=yDCVbUUckpwocmfpUySko7IS83LTV0u0&expireTime=1710840183</td></tr><tr><td>H5</td><td>30 天</td><td>https://quick.test.qian.tencent.cn/guide?Code=yDCVbUUckpwtvxqoUbTw4VBBjLbfAtW7&CodeType=QUICK&shortKey=yDCVbUY7lhqV7mZlCL2d</td></tr><tr><td>APP</td><td>一天</td><td>/pages/guide/index?to=USER_VERIFY&verifyToken=yDCVbUUckpwocm96UySko7ISvEIZH7Yz&expireTime=1710840455 </td></tr></tbody></table>
	Url *string `json:"url"`
}

type SuccessDeleteStaffData struct {
	// 员工名
	DisplayName *string `json:"displayName"`

	// 员工手机号
	Mobile *string `json:"mobile"`

	// 员工在电子签平台的id
	UserId *string `json:"userId"`
}

type SuccessUpdateStaffData struct {
	// 传入的用户名称
	DisplayName *string `json:"displayName"`

	// 传入的手机号，没有打码
	Mobile *string `json:"mobile"`

	// 员工在腾讯电子签平台的唯一身份标识，为32位字符串。
	// 可登录腾讯电子签控制台，在 "更多能力"->"组织管理" 中查看某位员工的UserId(在页面中展示为用户ID)。
	UserId *string `json:"userId"`

	// H5端员工实名链接
	//
	// 只有入参 InvitationNotifyType = H5的时候才会进行返回。
	Url *string `json:"url"`
}

type TemplateInfo struct {
	// 模板ID，模板的唯一标识
	TemplateId *string `json:"templateId"`

	// 模板的名字
	TemplateName *string `json:"templateName"`

	// 此模块需要签署的各个参与方的角色列表。RecipientId标识每个参与方角色对应的唯一标识符，用于确定此角色的信息。
	//
	// [点击查看在模板中配置的签署参与方角色列表的样子](https://qcloudimg.tencent-cloud.cn/raw/e082bbcc0d923f8cb723d98382410aa2.png)
	//
	Recipients []*Recipient `json:"recipients"`

	// 模板的填充控件列表
	//
	// [点击查看在模板中配置的填充控件的样子](https://qcloudimg.tencent-cloud.cn/raw/cb2f58529fca8d909258f9d45a56f7f4.png)
	Components []*Component `json:"components"`

	// 此模板中的签署控件列表
	//
	// [点击查看在模板中配置的签署控件的样子](https://qcloudimg.tencent-cloud.cn/raw/29bc6ed753a5a0fce4a3ab02e2c0d955.png)
	SignComponents []*Component `json:"signComponents"`

	// 模板描述信息
	Description *string `json:"description"`

	// 此模板的资源ID
	DocumentResourceIds []*string `json:"documentResourceIds"`

	// 生成模板的文件基础信息
	FileInfos []*FileInfo `json:"fileInfos"`

	// 此模板里边附件的资源ID
	AttachmentResourceIds []*string `json:"attachmentResourceIds"`

	// 签署人参与签署的顺序，可以分为以下两种方式：
	//
	// <b>无序</b>：不限定签署人的签署顺序，签署人可以在任何时间签署。此种方式值为 ：｛-1｝
	// <b>有序</b>：通过序列数字标识签署顺序，从0开始编码，数字越大签署顺序越靠后，签署人按照指定的顺序依次签署。此种方式值为： ｛0，1，2，3………｝
	SignOrder []*int64 `json:"signOrder"`

	// 此模板的状态可以分为以下几种：
	//
	// <b>-1</b>：不可用状态。
	// <b>0</b>：草稿态，即模板正在编辑或未发布状态。
	// <b>1</b>：正式态，只有正式态的模板才可以发起合同。
	Status *int64 `json:"status"`

	// 模板的创建者信息，用户的名字
	//
	// 注： `是创建者的名字，而非创建者的用户ID`
	Creator *string `json:"creator"`

	// 模板创建的时间戳，格式为Unix标准时间戳（秒）
	CreatedOn *int64 `json:"createdOn"`

	// 此模板创建方角色信息。
	//
	// [点击查看在模板中配置的创建方角色的样子](https://qcloudimg.tencent-cloud.cn/raw/e082bbcc0d923f8cb723d98382410aa2.png)
	Promoter *Recipient `json:"promoter"`

	// 模板类型可以分为以下两种：
	//
	// <b>1</b>：带有本企业自动签署的模板，即签署过程无需签署人手动操作，系统自动完成签署。
	// <b>3</b>：普通模板，即签署人需要手动进行签署操作。
	TemplateType *int64 `json:"templateType"`

	// 模板可用状态可以分为以下两种：
	//
	// <b>1</b>：（默认）启用状态，即模板可以正常使用。
	// <b>2</b>：停用状态，即模板暂时无法使用。
	//
	// 可到控制台启停模板
	Available *int64 `json:"available"`

	// 创建模板的企业ID，电子签的机构ID
	OrganizationId *string `json:"organizationId"`

	// 模板创建人用户ID
	CreatorId *string `json:"creatorId"`

	// 模板的H5预览链接,有效期5分钟。
	// 可以通过浏览器打开此链接预览模板，或者嵌入到iframe中预览模板。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewUrl *string `json:"previewUrl"`

	// 用户自定义合同类型。
	//
	// 返回配置模板的时候选择的合同类型。[点击查看配置的位置](https://qcloudimg.tencent-cloud.cn/raw/4a766f0540253bf2a05d50c58bd14990.png)
	//
	// 自定义合同类型配置的地方如链接图所示。[点击查看自定义合同类型管理的位置](https://qcloudimg.tencent-cloud.cn/raw/36582cea03ae6a2559894844942b5d5c.png)
	//
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserFlowType *UserFlowType `json:"userFlowType"`

	// 模板版本的编号，旨在标识其独特的版本信息，通常呈现为一串字符串，由日期和递增的数字组成
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateVersion *string `json:"templateVersion"`

	// 模板是否已发布可以分为以下两种状态：
	//
	// <b>true</b>：已发布状态，表示该模板已经发布并可以正常使用。
	// <b>false</b>：未发布状态，表示该模板还未发布，无法使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Published *bool `json:"published"`

	// <b>集体账号场景下</b>： 集团账号分享给子企业的模板的来源模板ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareTemplateId *string `json:"shareTemplateId"`

	// 此模板配置的预填印章列表（包括自动签署指定的印章）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateSeals []*SealInfo `json:"templateSeals"`
}

type UploadFile struct {
	// Base64编码后的文件内容
	FileBody *string `json:"fileBody"`

	// 文件名，最大长度不超过200字符
	FileName *string `json:"fileName"`
}

type UserFlowType struct {
	// 合同类型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserFlowTypeId *string `json:"userFlowTypeId"`

	// 合同类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"name"`

	// 合同类型说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"description"`
}

type UserInfo struct {
	// 用户在平台的编号
	UserId *string `json:"userId"`
}

type UserThreeFactor struct {
	// 签署方经办人的姓名。
	// 经办人的姓名将用于身份认证和电子签名，请确保填写的姓名为签署方的真实姓名，而非昵称等代名。
	Name *string `json:"name"`

	// 证件类型，支持以下类型
	// <ul><li>ID_CARD : 中国大陆居民身份证 (默认值)</li>
	// <li>HONGKONG_AND_MACAO : 港澳居民来往内地通行证</li>
	// <li>HONGKONG_MACAO_AND_TAIWAN : 港澳台居民居住证(格式同居民身份证)</li></ul>
	IdCardType *string `json:"idCardType"`

	// 证件号码，应符合以下规则
	// <ul><li>居民身份证号码应为18位字符串，由数字和大写字母X组成（如存在X，请大写）。</li>
	// <li>港澳居民来往内地通行证号码共11位。第1位为字母，“H”字头签发给香港居民，“M”字头签发给澳门居民；第2位至第11位为数字。</li>
	// <li>港澳台居民居住证号码编码规则与中国大陆身份证相同，应为18位字符串。</li></ul>
	IdCardNumber *string `json:"idCardNumber"`
}

type WebThemeConfig struct {
	// 是否显示页面底部电子签logo，取值如下：
	// <ul><li> **true**：页面底部显示电子签logo</li>
	// <li> **false**：页面底部不显示电子签logo（默认）</li></ul>
	DisplaySignBrandLogo *bool `json:"displaySignBrandLogo"`

	// 主题颜色：
	// 支持十六进制颜色值以及RGB格式颜色值，例如：#D54941，rgb(213, 73, 65)
	// <br/>
	WebEmbedThemeColor *string `json:"webEmbedThemeColor"`
}
