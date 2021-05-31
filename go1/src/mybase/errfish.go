package mybase

const (
	//My Error Code
	NEED_CLOSE_AGENT           = -1 + iota   //需要关闭客户端连接,
	SUCCESS                                  //返回成功
	ERROR_NOTHING_TODO         = 9997 + iota //服务器处理了，但是没有任何结果。
	ERROR_SERVER_FAIL                        // 10000, //服务器发生错误
	INVALID_UUID                             // 10001, //领取失败，同一个uuid已领取
	ERR_REPEAT                               // 10002, //该用户今天已领取奖励
	NOT_ENOUGH                               // 10003, //领取失败，今日奖励已领完，明日请早
	INVALID_ROOMID                           // 10004, //创建对战房失败，不是指定的房间id（～//4）
	ERR_ALLOC_DESK                           // 10005, //创建对战房失败，桌子分配
	ROOM_IS_NULL                             // 10006, //创建对战房失败，给定的房间ID对应的房间不存在
	DESK_IS_NULL                             // 10007, //给定的桌子ID对应的桌子不存在
	INVALID_PWD                              // 10008, //进入桌子密码错误
	INVALID_UID                              // 10009, //设置对战房信息，不是此房间的拥有者，不能设置房间信息。
	ERROR_COIN_NOENOUGH                      // 10010, //金币不足
	ERROR_SAME_ACCOUNT                       // 10011, //该用户名已被人使用
	MISSION_GETED                            // 10012, //奖励已领取。
	INVALID_UNFINISH                         // 10013, //任务未完成
	DESK_IS_FULL                             // 10014, //桌子已满员
	ERROR_SERVER_PAUSE                       // 10015, //停止服务（游戏维护中）
	INVALID_TOKEN                            // 10016, //登录错误，错误的token
	INVALID_ID                               // 10017, //身份证号错误
	INVALID_UNAME                            // 10018, //错误的用户名
	DB_UPDATE_ERROR                          // 10019, //设置昵称，密码，头像错误（数据库更新失败）
	PWD_FMT_ERROR                            // 10020, //设置密码错误（密码格式错误）
	OLDPWD_ERROR                             // 10021, //原密码错误
	SETNICK_ERROR                            // 10022, //设置昵称错误
	SETAVATAR_ERROR                          // 10023, //设置头像错误
	MERGE_DRAGONBALL_FAILED                  // 10024, //合成龙珠失败
	MERGE_NOTENOUGH_DRAGONBALL               // 10025, //合成龙珠错误，指定的数量错误，没有足够的可用的龙珠可供合成
	NEED_REFRESH                             // 10026, //客户端需要静默重新登录一下。。
	INVALID_DRAGONBALL_COUNT                 // 10027, //合成龙珠错误，指定的合成数量错误。
	USED_NOTENOUGH_DRAGONBALL                // 10028, //使用龙珠错误，没有足够的龙珠可供使用。
	FISH_ALREADY_DEAD                        // 10029, //鱼已经被其他玩家打死
	DRAGONBALL_TYPE_ERROR                    // 10030, //龙珠类型错误
	FISH_CATEGORY_ERROR                      // 10031, //龙珠击杀鱼的类型错误，请选择黄金鱼。
	RECEIVER_USER_OFFLINE                    // 10032, //赠送龙珠错误，接收用户不在线。
	SEND_NOTENOUGH_DRAGONBALL                // 10033, //赠送龙珠错误，没有足够的龙珠可供使用。
	SEND_SWITCH_OFF                          // 10034, //赠送龙珠错误，赠送功能未开启。
	SEND_VIP_TOO_LOW                         // 10035, //赠送龙珠错误，用户VIP等级太低。
	SEND_LEVEL_TOO_LOW                       // 10036, //赠送龙珠错误，用户经验等级太低。
	SEND_NOTENOUGH_COIN                      // 10037, //赠送龙珠错误，赠送手续费不够。
	CANNOT_SEND_SELF                         // 10038, //赠送龙珠错误，不能赠送给自己。
	GET_COLLAPSE_GOLD                        // 10039, //获取破产金错误
	NULL_COLLAPSE_GOLD                       // 10040, //今日领取次数已用完
	CANNOT_COLLAPSE_GOLD                     // 10041, //不能领取破产金
	NULL_PHONE_NUMBER                        // 10042, //未绑定手机
	//v1.1.0
	USED_NOTENOUGH_DRAGONCARD // 10043, //没有可使用的龙珠卡
	ERROR_CDKEY_OUTDATE       // 10044, //CDKEY码已过期
	CDKEY_USED                // 10045, //您已使用过该兑换码
	CDKEY_USER_EXCHANGED      // 10046, //CDKey该用户已领取
	INVALID_MAIL              // 10047, //邮件信息已失效
	//活动错误码提示
	CANNOT_GET         // 10048, //不能领取，用户等级不足或积分不足
	OUT_OF_REWARD      // 10049, //奖励已领完
	ALREADY_GET        // 10050, //已领取
	INVALID_ACTIVITY   // 10051, //无效的活动id
	INVALID_WEEKLYCARD // 10052, //无效的周卡
	//双倍金币，双倍爆率
	INVALID_VIP_LEVEL    // 10053, //错误的VIP等级，不能使用双倍金币，双倍爆率
	DOUBLE_GOLD_ENABLED  // 10054, //不能使用双倍爆率，正在使用双倍金币
	DOUBLE_SPEED_ENABLED // 10055, //不能使用双倍金币，正在使用双倍射速
	//抽奖活动
	INVALID_IP               // 10056, //无效的ip地址
	ERR_ENCRYPTION           // 10057, //加密错误
	CDKEY_IVALID             // 10058, //CDKEY码无效
	BUY_REPEATED             // 10059, //重复购买
	CANNOT_SPLIT             // 10060, //不能拆分
	INVALID_COUNT            // 10061, //给定的数量无效
	DRAGONBALL_NOTENOUGHT    // 10062, //拆分龙珠失败,可拆分龙珠不足
	DRAGONBALL_MAXCOUNT      // 10063, //拆分龙珠失败，已达到最大值
	NOTENOUGH_FLOOR_COIN     // 10064, //保底金币不足
	SEND_DRAGONBALL_MAXCOUNT // 10065, //赠送龙珠错误，接收者的龙珠大于最大值。
	INVALID_CREATEDESK_LEVEL // 10066, //等级X级以上才能创建房间
	PEARL_FULL               // 10067, //珍珠数量数量达到上限
	DB1_FULL                 // 10068, //一星龙珠数量达到上限
	DB2_FULL                 // 10069, //二星龙珠数量达到上限
	DB3_FULL                 // 10070, //三星龙珠数量达到上限
	DC_FULL                  // 10071, //龙珠卡数量达到上限
	LB_FULL                  // 10072, //喇叭数量达到上限
	INVALID_MAIL_REWARD      // 10073, //奖励信息不存在
	USER_OFFLINE             // 10074, //玩家不在线
	FUNCTION_NOT_COMPLETE    // 10075, //该功能未完成
	NO_LOG                   // 10076, //没有记录
	FREEZE_USER              // 10077, //登录错误，用户被冻结
	NOTICE_GET_FAILED        // 10078, //获取通知失败
	USER_NOT_EXIST           // 10079, //用户不存在
	INVALID_GAME_STATE       // 10080, //PK房不存在该游戏状态
	INVALID_PAGE             // 10081, //给定的页码错误
	NO_RETURN_CONTENT        // 10082, //没有返回数据
	DRAGONBALL_CANNOT_USE    // 10083, //本房间不能使用龙珠
	DRAGONCARD_CANNOT_USE    // 10084, //本房间不能使用龙珠卡
	MAX_COIN_LIMIT           // 10085, //超过当前场次最大金币限制
	FORBID_BROADCAST         // 10086, //禁止发送喇叭
	PHONE_NUMBER_OVER        // 10087, //该手机号已超过注册上限
	TRANSFER_COIN_OVER       // 10088, //超过今日赠送上限
	//领取vip奖励
	NOT_VIP           // 10089, //领取失败，不是vip
	ERROR_ALREADY_GET // 10090, //领取失败，已领取或者无奖励可领
	// 1.1.6
	BUY_FORT_EXISTS  // 10091, //指定的炮台已经存在
	BUY_INVALID_FORT // 10092, //炮台配置信息不存在
	ERROR_SET_FORT   // 10093, //设置炮台错误，没有可用的炮台
	// 1.1.8 v1.1.2
	UNLOCKED_FORTVALUE        // 10094, //炮台倍数已经解锁
	INVALID_FORTVALUE         // 10095, //错误的炮台倍数值
	ERROR_DIAMOND_NOENOUGH    // 10096, //钻石余额不足
	INVALID_FORT_VALUE        // 10097, //炮台需要升级才能使用响应功能
	ERROR_STAT_INVALID        // 10098, //状态错误,请稍后再重试
	TASK_INCOMPLETE           // 10099, //任务未完成，不能领取奖励
	INVALID_ALMS_TIME         // 10100, //领取救济金失败，倒计时中
	INVALID_ALMS_STATE        // 10101, //领取救济金失败，申请后倒计时结束才能领取
	EXPIRED_CARD              // 10102, //没有购买周/月卡，或者周/月卡已过期
	ERROR_FROZE_ROOM          // 10103, //不在房间中，不能使用冻结功能
	ERROR_FROZE_TIME          // 10104, //切换场景中，不能使用冻结功能
	ERROR_ALREADY_FROZE       // 10105, //正在使用冻结功能
	INVALID_PROPS_TYPE        // 10106, //错误的道具类型
	PROPS_IS_NULL             // 10107, //道具信息不存在
	INVALID_ROOM              // 10108, //该功能不能在当前房间使用
	ERROR_CODE_10109          // 10109, //邀请码无效
	ERROR_CODE_10110          // 10110, //必须是最近注册的用户才能填邀请码
	ERROR_CODE_DADABASE_ERROR // 10111, //数据服务器异常
	ERROR_CODE_10112          // 10112, //邀请人注册时间应早于被邀请人
	ERROR_CODE_10113          // 10113, //您已填写邀请人了
	ERROR_PARAM_INVALID       // 10114, //请求参数异常
	ERROR_NOTENOUGH_HBQ       // 10115, //红包券不足
	ERROR_OPT_TODYA_MAX       // 10116, //您今天不能再这样做了 今日次数达到上限
	//res new
	ERROR_NOTENOUGH_FUZI   // 10117, //福字不足
	ERROR_REPEAT_ZHOUKA    // 10118, //您已启用一张周卡了
	ERROR_ACCOUNT_PREFIX   // 10119, //用户名前缀不能以GW_或WX_开头.
	ERROR_APP_TOO_OLD      // 10120, //您的APP太老了，请前往官网更新
	ERROR_PROP_NOTENOUGH   // 10121, //您的该道具数量不足
	ERROR_VIP_NOENOUGH     // 10122, //操作失败,您的VIP等级不足
	ERROR_INYT_ALREADY     // 10123, //您已加入一个鱼塘
	ERROR_SERVER_BUSY      // 10124, //网络开小差了，请稍后重试
	ERROR_IN_SAVEMODE      // 10125, //处于保护期，不能加入新鱼塘
	ERROR_NO_RIGHT         // 10126, //您没有该操作的权限
	ERROR_TIME_NOARRIVE    // 10127, //操作失败，时间尚未达到
	ERROR_OPT_SAME_UID     // 10128, //您今天不能再对Ta这样做了
	ERROR_TIME_CLEAN       // 10129, //操作失败,服务器正在整理数据,请稍后再来
	ERROR_IN_PROTECTED     // 10130, //操作失败,对方处于保护期
	ERROR_APPLY_REPEATED   // 10131, //您已经提交申请,请耐心等待处理
	ERROR_MSG_OUTOFDATE    // 10132, //操作失败,请求已过时
	ERROR_INYT_NO          // 10133, //您尚未加入鱼塘
	ERROR_YTCOIN_NOENOUGH  // 10134, //鱼塘资金不足
	ERROR_CDKEY_NOUSE      // 10135, //操作失败,您无法使用该兑换码
	ERROR_YT_FULL          // 10136, //目前鱼塘已经人满了
	ERROR_NO_DCPORZT       // 10137, //您没有特殊炮台
	ERROR_ACCOUNT_LESS6    // 10138 "用户名少于6位",
	ERROR_SAME_AC_PWD      // 10139 "请确保账号、昵称、密码都不一样",
	ERROR_NAME_FORBIDDEN   // 10140 "用户昵称不合法",
	ERROR_NEED_WX_AUTH     // 10141 "微信授权失效,请重新授权",
	ERROR_NO_ACCOUNT       // 10142 "帐号不存在或密码错误",
	ERROR_SAME_PHONE       // 10143: "该手机号已绑定您当前账号,无需找回",
	ERROR_PHONE_NOACCOUNT  // 10144: "该手机号尚未绑定账号",
	ERROR_NO_WXUNIONID     // 10145: "该账号无需找回,您使用账号登录即可",
	ERROR_ONLY_WXACCOUNT   // 10146: "只支持微信登录用户找回老账号",
	ERROR_ROOM_LIMIT       // 10147, --当前状态不满足当前房间的进入要求
	ERROR_INPUT_TGY        // 10148, -- "您输入的推广员无效或者还在审核中"
	ERROR_BULLET_NOENOUGH  //= 10149, -- "炮弹数量不足"
	ERROR_LASTGAME_IS_OVER //= 10150, -- "您的上轮比赛已经结束了"
	ERROR_SAY_POLITE       //= 10151, --"请文明发言"
	ERROR_ID_LIMIT         //= 10152, -- "该身份证已绑定账号达上限",
	ERROR_SMBX_TIMEOUT     //= 10153, --"您的神秘宝藏已经过期了",
	ERROR_REPEAT_GIFT      //= 10154, --"您已购买此类型礼包",

	CODE_TIP = 30000 //, //通用提示信息，需要在返回的消息中添加msg
	// 鱼乐错误码
	INVALID_USER        = 30001 //, //找不到该玩家
	NOTENOUGH_GAMECOIN  = 30002 //, //游戏币不足
	INVALID_BETPOS      = 30003 //, //无效的押注位置
	INVALID_BET         = 30004 //, //无效的押注
	OVER_MAXBET_PER     = 30005 //, //超过本期最大押注值
	OVER_MAXBET_DAY     = 30006 //, //超过当天最大押注值
	NOT_BETSTATUS       = 30007 //, //不在押注时间
	OVER_CHANGECOIN_DAY = 30008 //, //超过当天金币变化值
	NO_HISTORY_INFO     = 30009 //, //历史数据为空
	NO_RANK_INFO        = 30010 //, //排行榜数据
	PARAMETER_ERROR     = 30011 //, //缺少参数
	CANNOT_ENTERGAME    = 30012 //, //不能进入游戏
	GET_COIN_ERR        = 30013 //, //获取金币错误
	// slot游戏错误码
	SLOT_INVALID_USER        = 31001 //, //lobby中的__uid2usr不存在用户信息
	SLOT_INVALID_TOKEN       = 31002 //, //token错误，登录失败
	SLOT_BET_NOT_ENOUGH_COIN = 31003 //, //bet错误，金币不足
	SLOT_STOP_SERVICE        = 31004 // //SLOT停止服务（tcpproxy中发送）

	ERR_USRINFO_NULL    = 40001 //找不到该用户的扫码信息
	ERR_RETURNCODE_NULL = 40002 //返回的code值为空
	ERR_INVALID_USER    = 40003 //
	ERR_NEED_REGMOBILE  = 40004 //需要先在移动版注册
)

const (
	_              = iota
	ErrorCode20001 = 20000 + iota // "支付请求返回超时!!!",
	ErrorCode20002                // "支付请求返回内容无法解析!!!",
	ErrorCode20003                // "支付错误，重复的交易ID!!!",
	ErrorCode20004                // "支付宝支付错误,请联系客服",
	ErrorCode20005                // "微信支付错误,请联系客服",
	ErrorCode20006                // "用户token错误!!!",
	ErrorCode20007                // "商品id错误!!!",
	ErrorCode20008                // "支付失败!!!",
	ErrorCode20009                // "支付失败,错误的BundleId",
	ErrorCode20010                // "支付错误，错误的渠道商",
	ErrorCode20011                // "支付错误，错误的订单ID!!!",
	ErrorCode20012                // "支付错误，错误的商品ID!!!",
	ErrorCode20013                // "支付错误，已经达到单日充值最大限额!!!",
	ErrorCode20014                // "支付错误，订单已取消!!!",
	ErrorCode20015                // 您不能再购买该商品了！
	//新加错误码龙珠卡新手场对战房不能购买
	ErrorCode20016 // "应用程序商店无法读取您提供的JSON对象!!!",
	ErrorCode20017 // "接收数据属性中的数据格式错误或丢失!!!",
	ErrorCode20018 // "无法验证收据!!!",
	ErrorCode20019 // "您提供的密钥与您的帐户上的密钥不匹配!!!" // 仅返回用于iOS自动更新订阅的iOS6交易收据。
	ErrorCode20020 // "收据服务器当前不可用!!!",
	ErrorCode20021 // "此收据有效，但订阅已过期!!!" // 当此状态代码返回到服务器时，接收数据也会作为响应的一部分进行解码和返回。仅返回用于自动续订订阅的IOS6样式事务收据。
	ErrorCode20022 // "此收据来自正式环境，但它被发送到测试环境进行验证。把它发送到正式环境!!!",
)

// 奖励类型
const (
	UNKNOWN                 = -1 + iota //未知
	DAYAWARDS                           //每日奖励金币
	MISSION                             //每日任务
	FISHGAME                            //捕鱼游戏中获得的奖励
	PKBONUS                             //捕鱼游戏对战房间奖金
	PKREADY                             //捕鱼游戏对战房间准备金、手续费扣除
	PKBACK                              //捕鱼游戏对战房间准备金、手续费退回
	ADMIN                               //后台增加金币 废弃。。
	WECHAT                              //微信公众号充值
	WECHAT_APP                          //微信APP支付
	APPLE                               //苹果支付
	FIRSTPAY                            //首次充值iapppay中的5元时，赠送的金币值。
	NEWER7                              //用户七天奖励
	ALMS                                //救济金
	CATCHFISH_BY_BALL                   // 使用龙珠获得的金币奖励-----物品兑换。
	THIRD_PAY                           // 外接第三方充值
	ALI                                 // 支付宝充值
	LOTTERY                             //抽奖游戏
	CDKEY                               //CDKey兑换
	WEEKLYCARD                          //周卡
	CHONGJIDAJIANGSAI                   //冲级大奖赛活动
	LONGZHUDUOBAO                       //龙珠夺宝活动
	UPGRADEVIP                          //vip升级
	GAME_YULE                           //鱼乐游戏
	WEBPAY                              //官网充值（7,8,9,14,15 合集给福字专用）
	ROBOT_ADD                           //机器人添加
	AGENTPAY_AWARD                      // 代理充值奖励
	RANKREWARDS                         // 排行奖励
	YULE_FEE                            //服务费
	PROP_USE                            //道具使用
	VIP_DAILY_REWARD                    //vip每日奖励
	SLOT_SLOT_BET                       //押注
	SLOT_SLOT_REWARD                    //水浒传Slot和Bonus奖励
	SLOT_BIBEI_BET                      //水浒传比倍押注
	SLOT_BIBEI_REWARD                   //水浒传比倍奖励
	BUY_FORT                            //购买炮台消耗的金币
	USE_DOUBLE_GOLD                     //使用双倍金币消耗钻石记录
	USE_DOUBLE_SPEED                    //使用双倍射速消耗钻石记录
	GROWTH_TASK                         //成长任务奖励
	ONLINE_REWARD                       //在线奖励
	UNLOCK_FORTVALUE                    //解锁炮倍，减少钻石，增加金币
	VIP_REWARD                          //VIP奖励
	LEVEL_REWARD                        //LEVEL奖励
	USE_LOCK                            //使用锁定道具消耗钻石记录
	USE_FROZE                           //使用冻结道具消耗钻石记录
	USE_DRAGONCARD                      //使用龙珠卡消耗钻石记录
	WEEKLYCARD_DAILY_REWARD             //周卡每日奖励
	MONTHCARD_REWARD                    //月卡奖励
	MONTHCARD_DAILY_REWARD              //月卡每日奖励
	BUY_PROPS                           //商城中购买道具
	ACTIVITY_LEVEL                      //冲级活动奖励
	ACTIVITY_DIAMONDS                   //钻石消耗活动奖励
	ACTIVITY_SHARING                    //分享有礼活动奖励
	MAIL_REWARD                         //未知类型的邮件奖励
	SYSTEM                              //系统
	NAME_DIAMOND                        //改名字消耗钻石
	REWARD_INVITE                       //邀请新用户奖励
	MERGE                               //合成
	SEND                                //赠送
	RECV                                //接收
	SPLIT                               //拆分
	EXCHANGE                            //商店兑换
	HB_TIXIAN                           //微信红包提现
	JF_REWARD                           //新年集褔活动
	FXBOSS                              // 分享捕获boss活动
	USE_DIAMOND_FOR_CALL                //使用召唤黄金鱼消耗钻石记录
	LZHB                                //龙珠红包抽取奖励。
	LZDBHB                              //龙珠夺宝奖励
	HF_TIXIAN                           //话费提现，
	YT_CHECKIN                          //鱼塘签到，
	YT_YUHUO                            //鱼塘鱼货收取，
	YT_STEAL                            //鱼塘鱼货偷取，
	YT_DONATE                           //鱼塘资金捐献，
	YT_CREATE                           //鱼塘创建
	YT_MODIFY                           //鱼塘修改
	FISHGAME_YT                         //鱼塘捕鱼游戏奖励-如果用户加入了俱乐部，那么改成用这个
	BACK_LZXG                           //龙珠峡谷返还
	ACTIVITY_REWARD                     //活动奖励
	PAY_REWARD                          //充值返利/金币售卖返利
)

const (
	REG_TYPE_WX = int32(iota)
	REG_TYPE_HW
	REG_TYPE_YYB
	REG_TYPE_VIVO
	REG_TYPE_PHONE
)
