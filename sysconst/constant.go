package sysconst

// ThirdPartyPlatfrom
const (
	THIRD_PARTY_PLATFROM_APOLLO = "/ThirdPartyPlatfrom/Apollo" // http 第三方平台過來的資料 (阿波羅)
	THIRD_PARTY_PLATFROM_DT     = "/ThirdPartyPlatfrom/Dt"     // http 第三方平台過來的資料 (阿波羅)
	THIRD_PARTY_PLATFROM_DIOS   = "/ThirdPartyPlatfrom/Dios"   // http 第三方平台過來的資料 (阿波羅)
)

// 錯誤代碼
type Base int

const (
	ERROR_CODE_SUCCESS          Base = iota // 0沒有錯誤
	ERROR_CODE_NO_FIND_CMD                  // 1找不到Cmd
	ERROR_CODE_NO_FIND_ACCOUNT              // 2找不到帳號
	ERROR_CODE_NO_LOGIN                     // 3帳號未登入
	ERROR_CODE_CLIENT_TOO_MATCH             // 4服務器上限滿額
	ERROR_CODE_NO_FIND_TABLE                // 5找不到桌子
	ERROR_CODE_NO_FIND_SEAT                 // 6找不到位子
	ERROR_CODE_NO_USE                       // 7資源未使用

	ERROR_CODE_FULL_PLAYER              // 8人滿了
	ERROR_CODE_CARRY_BALANCE_NOT_ENOUGH // 9想帶進來的錢不夠
	ERROR_CODE_RE_JOIN_TABLE            // 10重複入桌
	ERROR_CODE_ERROR_DATA               // 11資料錯誤 (一般物件數值錯誤)
	ERROR_CODE_ERROR_PARAMETER          // 12參數錯誤 (函數傳入值錯誤)
	ERROR_CODE_ERROR_OPEN_TABLE         // 13開桌失敗
	ERROR_CODE_ERROR_USER_ID            // 14錯誤的User_ID

	ERROR_CODE_BALANCE_UPDATE_FAIL      // 15更新錢錯誤
	ERROR_CODE_TABLE_BALANCE_NOT_ENOUGH // 16桌內的錢不夠
	ERROR_CODE_BALANCE_CHECK_FAIL       // 17檢查錢錯誤
	ERROR_CODE_DATA_UPDATE_FAIL         // 18更新資料錯誤
	ERROR_CODE_ERROR_JSON_MARSHAL       // 19Json解析錯誤
	ERROR_CODE_ERROR_GAME_MODE          // 20錯誤的GameMode
	ERROR_CODE_ERROR_JOIN_TABLE         // 21加入桌失敗
	ERROR_CODE_RE_OPEN_TABLE            // 22重複開桌
	ERROR_CODE_ERROR_ACCOUNT_CREATE     // 23建立帳號失敗

	ERROR_CODE_RE_ACCOUNT_CREATE          // 24同一個連線,重複建立帳號 (請先登出或是...)
	ERROR_CODE_RE_LOGIN                   // 25同一個連線,重複登入兩次
	ERROR_CODE_TABLE_BET_IS_ZERO          // 26押注金額不能為零
	ERROR_CODE_ERROR_POINT_IS_NIL         // 27指標錯誤
	ERROR_CODE_ERROR_TABLE_UPDATE         // 28更新桌子錯誤
	ERROR_CODE_ERROR_SEAT_UPDATE          // 29更新位子錯誤
	ERROR_CODE_DUPLICATE_ACCOUNT          // 30帳號重複
	ERROR_CODE_ERROR_AUTH_PLATFORM        // 31平台驗證失敗
	ERROR_CODE_ERROR_INSTER_TRANSFER_LOG  // 32交易log寫入失敗
	ERROR_CODE_ERROR_UPDATE_TRANSFER_LOG  // 33交易log更新失敗
	ERROR_CODE_LOBBY_BALANCE_NOT_ENOUGH   // 34大廳的錢不夠(轉出)
	ERROR_CODE_ERROR_RE_PLATFORM_TRANSFER // 35平台重複轉帳交易
	ERROR_CODE_NO_FIND_GAME               // 36找不到此遊戲
	ERROR_CODE_NO_FIND_BULLET             // 37找不到子彈
	ERROR_CODE_ERROR_INTERFACE_OBJ        // 38界面物件轉換失敗

	ERROR_CODE_ERROR_PROB_SERVICE // 39服務器異常
	ERROR_CODE_ERROR_PROB         // 40機率系統失敗

	ERROR_CODE_NO_FIND_FEATURE_BULLET  // 41找不到特殊子彈
	ERROR_CODE_NO_FIND_BETCLUSTER      // 42找不到注單
	ERROR_CODE_ERROR_INTERFACE_GAMEOBJ // 43遊戲物件轉換失敗
	ERROR_CODE_NO_FIND_ODDS            // 44找不到賠率表

	ERROR_CODE_ERROR_HTTP_REQ_GET  // 45http get 錯誤
	ERROR_CODE_ERROR_HTTP_REQ_POST // 46http post 錯誤
	ERROR_CODE_ERROR_MINIGAME_BET  // 47小遊戲押注錯誤

	ERROR_CODE_ERROR_THIRD_PARTY_AUTH // 48(网络异常)第三方平台驗證錯誤

	ERROR_CODE_ERROR_OVERLIMIT_BET // 49押注超過上限

	ERROR_CODE_ERROR_MEMBER_LOCK // 50會員凍結
	ERROR_CODE_NO_FIND_PLATFORM  // 51找不到平台

	ERROR_CODE_ERROR_OPEN_TATBL_MAX // 52(開桌數已達滿額)(目前不支援同帳號多開)
	ERROR_CODE_ERROR_TIMELIMIT_BET  // 53押注超過時間

	ERROR_CODE_PERMISSION_DENIED      // 54權限不足
	ERROR_CODE_GM_CLOSE_TABLE         // 55GM關閉桌子中
	ERROR_CODE_ERROR_TABLE_DELETE     // 56刪除桌子錯誤
	ERROR_CODE_MISMATCH_BET_ROUNDCODE // 57押注局號不合
	ERROR_CODE_ERROR_OVER_BETCOUNT    // 58投注次數已達上限
	ERROR_CODE_ERROR_DECODE           // 59加解密失敗

	ERROR_CODE_ERROR_RECONNECT_AGAIN // 60與伺服器失去連線,請重新啟動客戶端 (超過5分鐘用戶沒封包動作, 就踢掉玩家)

	ERROR_CODE_DATA_INSERT_FAIL // 61寫入資料錯誤
	ERROR_CODE_DATA_GET_FAIL    // 62讀取資料錯誤

	ERROR_CODE_GAME_CLOSE // 63遊戲關閉維護中
	ERROR_CODE_RE_LOGIN2  // 64同一個帳號,重複登入兩次

	ERROR_CODE_ERROR_AUTH_PLATFORM_IP      // 65平台IP驗證失敗
	ERROR_CODE_ERROR_AUTH_PLATFORM_ACCOUNT // 66平台帳密驗證失敗

	ERROR_CODE_ERROR_ACCOUNT_RECREATE // 67重複建立帳號
	ERROR_CODE_NO_FIND_LOBBY          // 68找不到大廳
	ERROR_CODE_ERROR_DEMOO_CMD        // 69錯誤的測試CMD
	ERROR_CODE_ERROR_SEND_MESSAGE     // 70無法傳送封包
	ERROR_CODE_TRY_ENTER_GAME         // 71本系统不支持用户多开，请重新操作

	ERROR_CODE_ERROR_GAME_CODE           // 72錯誤的GameCode
	ERROR_CODE_ERROR_LINE_BET            // 73錯誤的押注區間
	ERROR_CODE_INGAME_NOT_TRANSFER_MONEY // 74遊戲中禁止點數轉出

	ERROR_CODE_PLATFORM_MONEY_IN_TRANSFER // 75平台會員額度處理中

	ERROR_CODE_ERROR_GAME_STATUS // 76錯誤的遊戲狀態(e.g. 不該slot_bonus_spin的時候送此command)
	ERROR_CODE_WEBAPI_TIMEOUT    // 77 webapi req time out

	ERROR_CODE_NO_ENTER_GAME // 78人不在遊戲內

	ERROR_CODE_ERROR_TOKEN              // 79token錯誤
	ERROR_CODE_ERROR_BETID              // 80BetID錯誤
	ERROR_CODE_ERROR_BETLINE_MULTIPLIER // 81 total bet和line bet並非bet line倍

	ERROR_CODE_DT_TOKEN_OUTDATE // 82 DT token 過期

	ERROR_CODE_MAX
)

const (
	HTTP_CMD_BET_CLUSTER_GET = "http_bet_cluster_get" // http 取得注單資訊
	HTTP_CMD_BET_DETAIL_GET  = "http_bet_detail_get"  // http 取得細單資訊
)

// 定義遊戲模式
const (
	GAME_MODE_NULL     Base = iota // 0 沒有定義
	GAME_MODE_FISH                 // 1 魚機
	GAME_MODE_SLOT                 // 2 Slot
	GAME_MODE_POKER                // 3 撲克牌
	GAME_MODE_MAHJONG              // 4 麻將
	GAME_MODE_MINIGAME             // 5 小遊戲
	GAME_MODE_MAX                  // x 最大值
)
