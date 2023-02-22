package constants

const (
	HTTP_STATUS_SUCCESS = "success"
	HTTP_STATUS_ERROR   = "error"

	PAYMENT_CHAIN_NAME_POLYGON_MUMBAI  = "polygon.mumbai"
	PAYMENT_CHAIN_NAME_POLYGON_MAINNET = "polygon.mainnet"
	PAYMENT_CHAIN_NAME_BSC_TESTNET     = "bsc.testnet"

	// common
	API_URL_MCS_POLYGON_MAINNET = "https://api.multichain.storage"
	API_URL_MCS_POLYGON_MUMBAI  = "https://calibration-mcs-api.filswan.com"
	API_URL_MCS_BSC_TESTNET     = "https://calibration-mcs-bsc.filswan.com"
	API_URL_FIL_PRICE_API       = "https://api.filswan.com/stats/storage"

	// mcs api
	MCS_PARAMS      = "/api/v1/common/system/params"
	PRICE_RATE      = "/api/v1/billing/price/filecoin"
	PAYMENT_INFO    = "/api/v1/billing/deal/lockpayment/info"
	TASKS_DEALS     = "/api/v1/storage/tasks/deals"
	MINT_INFO       = "/api/v1/storage/mint/info"
	UPLOAD_FILE     = "/api/v1/storage/ipfs/upload"
	DEAL_DETAIL     = "/api/v1/storage/deal/detail/"
	USER_REGISTER   = "/api/v1/user/register"
	USER_LOGIN      = "/api/v1/user/login_by_metamask_signature"
	GENERATE_APIKEY = "/api/v1/user/generate_api_key"
	APIKEY_LOGIN    = "/api/v1/user/login_by_api_key"

	// bucket api
	CREATE_BUCKET = "/api/v2/bucket/create"
	BUCKET_LIST   = "/api/v2/bucket/get_bucket_list"
	DELETE_BUCKET = "/api/v2/bucket/delete"
	FILE_INFO     = "/api/v2/oss_file/get_file_info"
	DELETE_FILE   = "/api/v2/oss_file/delete"
	CREATE_FOLDER = "/api/v2/oss_file/create_folder"
	CHECK_UPLOAD  = "/api/v2/oss_file/check"
	UPLOAD_CHUNK  = "/api/v2/oss_file/upload"
	MERGE_FILE    = "/api/v2/oss_file/merge"
	FILE_LIST     = "/api/v2/oss_file/get_file_list"

	// contract
	USDC_ABI          = "ERC20.json"
	SWAN_PAYMENT_ABI  = "SwanPayment.json"
	MINT_ABI          = "SwanNFT.json"
	CONTRACT_TIME_OUT = 300

	BYTES_1KB = 1024
	BYTES_1MB = BYTES_1KB * BYTES_1KB
	BYTES_1GB = BYTES_1MB * BYTES_1KB

	DURATION_DAYS_DEFAULT = 525
)
