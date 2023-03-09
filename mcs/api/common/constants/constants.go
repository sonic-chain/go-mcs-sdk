package constants

const (
	HTTP_STATUS_SUCCESS = "success"
	HTTP_STATUS_ERROR   = "error"

	PAYMENT_CHAIN_NAME_POLYGON_MUMBAI  = "polygon.mumbai"
	PAYMENT_CHAIN_NAME_POLYGON_MAINNET = "polygon.mainnet"
	PAYMENT_CHAIN_NAME_BSC_TESTNET     = "bsc.testnet"

	// common
	API_URL_MCS_POLYGON_MAINNET = "https://api.multichain.storage"
	API_URL_MCS_POLYGON_MUMBAI  = "http://127.0.0.1:8889" // "https://calibration-mcs-api.filswan.com" //
	API_URL_MCS_BSC_TESTNET     = "https://calibration-mcs-bsc.filswan.com"
	API_URL_FIL_PRICE_API       = "https://api.filswan.com/stats/storage"

	// mcs api
	API_URL_MCS_GET_PARAMS = "/api/v1/common/system/params"

	API_URL_USER_REGISTER        = "/api/v1/user/register"
	API_URL_USER_LOGIN           = "/api/v1/user/login_by_metamask_signature"
	API_URL_USER_GENERATE_APIKEY = "/api/v1/user/generate_api_key"
	API_URL_USER_LOGIN_BY_APIKEY = "/api/v1/user/login_by_api_key"
	API_URL_USER_CHECK_LOGIN     = "/api/v1/user/check_login"

	API_URL_BILLING_HISTORY          = "/api/v1/billing"
	API_URL_BILLING_FILECOIN_PRICE   = "/api/v1/billing/price/filecoin"
	API_URL_BILLING_GET_PAYMENT_INFO = "/api/v1/billing/deal/lockpayment/info"

	API_URL_STORAGE_UPLOAD_FILE            = "/api/v1/storage/ipfs/upload"
	API_URL_STORAGE_GET_DEALS              = "/api/v1/storage/tasks/deals"
	API_URL_STORAGE_GET_DEAL_DETAIL        = "/api/v1/storage/deal/detail"
	API_URL_STORAGE_GET_DEAL_LOG           = "/api/v1/storage/deal/log"
	API_URL_STORAGE_GET_SOURCE_FILE_UPLOAD = "/api/v1/storage/source_file_upload"
	API_URL_STORAGE_UNPIN_SOURCE_FILE      = "/api/v1/storage/unpin_source_file"

	API_URL_STORAGE_WRITE_NFT_COLLECTION = "/api/v1/storage/mint/nft_collection"
	API_URL_STORAGE_GET_NFT_COLLECTIONS  = "/api/v1/storage/mint/nft_collections"
	API_URL_STORAGE_RECORD_MINT_INFO     = "/api/v1/storage/mint/info"
	API_URL_STORAGE_GET_MINT_INFO        = "/api/v1/storage/mint/info"

	API_URL_DAO_GET_DEALS_2_PRE_SIGN  = "/api/v1/dao/deals_to_pre_sign/x"
	API_URL_DAO_GET_DEALS_2_SIGN      = "/api/v1/dao/deals_to_sign/x"
	API_URL_DAO_GET_DEALS_2_SIGN_HASH = "/api/v1/dao/deals_to_sign_hash/x"

	// bucket api
	API_URL_BUCKET_CREATE_BUCKET          = "/api/v2/bucket/create"
	API_URL_BUCKET_GET_BUCKET_LIST        = "/api/v2/bucket/get_bucket_list"
	API_URL_BUCKET_DELETE_BUCKET          = "/api/v2/bucket/delete"
	API_URL_BUCKET_RENAME_BUCKET          = "/api/v2/bucket/rename"
	API_URL_BUCKET_GET_TOTAL_STORAGE_SIZE = "/api/v2/bucket/get_address_storage_total"

	API_URL_BUCKET_FILE_GET_FILE_INFO                = "/api/v2/oss_file/get_file_info"
	API_URL_BUCKET_FILE_DELETE_FILE                  = "/api/v2/oss_file/delete"
	API_URL_BUCKET_FILE_CREATE_FOLDER                = "/api/v2/oss_file/create_folder"
	API_URL_BUCKET_FILE_GET_FILE_INFO_BY_OBJECT_NAME = "/api/v2/oss_file/get_file_by_object_name"
	API_URL_BUCKET_FILE_CHECK_UPLOAD                 = "/api/v2/oss_file/check"
	API_URL_BUCKET_FILE_UPLOAD_CHUNK                 = "/api/v2/oss_file/upload"
	API_URL_BUCKET_FILE_MERGE_FILE                   = "/api/v2/oss_file/merge"
	API_URL_BUCKET_FILE_GET_FILE_LIST                = "/api/v2/oss_file/get_file_list"

	BYTES_1KB = 1024
	BYTES_1MB = BYTES_1KB * BYTES_1KB
	BYTES_1GB = BYTES_1MB * BYTES_1KB

	FILE_CHUNK_SIZE_MAX = 1 * BYTES_1KB

	DURATION_DAYS_DEFAULT = 525
	SECOND_PER_DAY        = 24 * 60 * 60
	DAY_PER_YEAR          = 365

	COPY_NUMBER_LIMIT = 5

	SOURCE_FILE_TYPE_NORMAL = 0
	SOURCE_FILE_TYPE_MINT   = 1
)
