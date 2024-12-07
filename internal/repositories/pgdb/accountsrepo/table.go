package accountsrepo

const TABLE = "accounts"
const TABLE_JOINED = "accounts_users_map"
const TABLE_ACCOUNT_INSTANCE = "flow_instances_accounts"
const (
	ID                          = "id"
	PRIMARY_USER_ID             = "primary_user_id"
	GST_NO                      = "gst_no"
	IS_ACTIVE                   = "is_active"
	DEFAULT_WORKFLOW_PRE_ORDER  = "default_workflow_pre_order"
	DEFAULT_WORKFLOW_POST_ORDER = "default_workflow_post_order"
	CREATED_AT                  = "created_at"
	MODIFIED_AT                 = "modified_at"
)
const (
	USER_ID    = "user_id"
	ACCOUNT_ID = "account_id"
)

const (
	INSTANCE_ID = "instance_id"
)
