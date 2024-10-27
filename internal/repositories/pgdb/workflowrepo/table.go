package workflowrepo

const (
	UPDATED_AT = "updated_at"
	CREATED_AT = "created_at"
	UPDATED_BY = "updated_by"
	CREATED_BY = "created_by"
)
const TABLE_WORKFLOW = "workflows"
const TABLE_FLOW = "flows"
const TABLE_FLOW_PARAMS = "flow_params"
const TABLE_FLOW_INSTANCE = "flow_instances"
const TABLE_FLOW_INSTANCE_PARAMS = "flow_instance_params"
const (
	ID   = "id"
	NAME = "name"
	TYPE = "type"
)

const (
	FLOW_ID     = "id"
	WORKFLOW_ID = "workflow_id"
	DESCRIPTION = "description"
	FLOW_TYPE   = "type"
	TITLE       = "title"
	ORDER       = "order"
	ACTIVE      = "active"
	TAT         = "tat"
)
const (
	FLOW_PARAMS_ID  = "id"
	FLOW_ID_PARAM   = "flow_id"
	PARAM_NAME      = "name"
	PARAM_TYPE      = "type"
	PARAM_MANDATORY = "mandatory"
)

const (
	FLOW_INSTANCES_ID    = "id"
	WORKFLOW_ID_INSTANCE = "workflow_id"
	DESCRIPTION_INSTANCE = "description"
	TYPE_INSTANCE        = "type"
	TITLE_INSTANCE       = "title"
	ORDER_INSTANCE       = "order"
	ACTIVE_INSTANCE      = "active"
	TAT_INSTANCE         = "tat"
	INSTANCE_ID          = "instance_id"
	INSTANCE_TYPE        = "instance_type"
	STATUS               = "status"
	ASSIGNED_TO          = "assigned_to"
	EXPIRES_AT           = "expires_at"
)

const (
	FLOW_INSTANCE_PARAMS_ID  = "id"
	FLOW_INSTANCE_ID_PARAM   = "flow_instance_id"
	INSTANCE_PARAM_NAME      = "name"
	INSTANCE_PARAM_TYPE      = "type"
	INSTANCE_PARAM_MANDATORY = "mandatory"
	INSTANCE_PARAM_VALUE     = "value"
)
