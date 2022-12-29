package common

const (
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
	//bucket api
	CREATE_BUCKET = "/api/v2/bucket/create"
	BUCKET_LIST   = "/api/v2/bucket/get_bucket_list"
	DELETE_BUCKET = "/api/v2/bucket/delete?bucket_uid="
	FILE_INFO     = "/api/v2/oss_file/get_file_info"
	DELETE_FILE   = "/api/v2/oss_file/delete"
	CREATE_FOLDER = "/api/v2/oss_file/create_folder"
	CHECK_UPLOAD  = "/api/v2/oss_file/check"
	UPLOAD_CHUNK  = "/api/v2/oss_file/upload"
	MERGE_FILE    = "/api/v2/oss_file/merge"
	FILE_LIST     = "/api/v2/oss_file/get_file_list"
)
