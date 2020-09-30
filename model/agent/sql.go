package agent

import ()

var sqlGetKeyAccessTable = "SELECT a.id,a.callback_url,l.name as agent_type ,a.secret_key, p.name as permission_realm, l.id as agent_type_id, p.id as permission_realm_id , u.account , u.full_name from api.agent a left join api.lt_agent_type l on a.agent_type_id = l.id left join api.permission_realm p on a.permission_realm_id = p.id left join api.user u on a.user_id = u.id WHERE a.deleted_at is null AND l.id !=1 ORDER BY a.id"

var sqlGetAgentType = "SELECT id,name from api.lt_agent_type WHERE deleted_at is null"

var sqlGetPermissionRealm = "SELECT id,name from api.permission_realm WHERE deleted_at is null"

var sqlGetAgentUserInfo = "SELECT u.account FROM api.agent a LEFT JOIN api.user u ON a.user_id=u.id WHERE a.id = $1"

var sqlPostUser = "INSERT INTO api.user(user_type_id,account,is_active) VALUES(1,$1,TRUE) ON CONFLICT (id) DO UPDATE SET id = excluded.id RETURNING id"

var sqlPostPermissionGroup = "INSERT INTO api.permission_group_user(permission_group_id, user_id) VALUES (10,$1) RETURNING id"

var sqlPostAgent = "INSERT INTO api.agent(user_id, agent_type_id, callback_url, secret_key, permission_realm_id, created_by, updated_by, created_at, updated_at) " +
	"VALUES ($6, $1, $2, $3, $4, $5, $5, NOW(), NOW()) " +
	"ON CONFLICT (user_id) DO UPDATE SET user_id = excluded.user_id RETURNING id"

var sqlPutUpdateKey = "UPDATE api.agent SET secret_key = $1, updated_by = $2, updated_at = NOW() WHERE id=$3"

var sqlPutAgent = "UPDATE api.agent SET agent_type_id=$1, callback_url=$2,  permission_realm_id=$3, updated_by = $4, updated_at = NOW() WHERE id=$5 RETURNING user_id"

var sqlPutLookupAgentType = "SELECT a.id, a.name from api.lt_agent_type a where a.id = $1 WHERE a.deleted_at is null"

var sqlPutLookupPermissionRealm = "SELECT a.id, pr.name from api.permission_realm pr where pr.id = $1 WHERE pr.deleted_at is null"

var sqlDeleteKey = "UPDATE api.agent SET secret_key = NULL, updated_by = $1, updated_at = NOW() WHERE id=$2"

var sqlDeleteLookupKeyAccess = "SELECT al.agent_user_id as access_log,dd.agent_user_id as dataimport,el.agent_user_id as eventlog FROM api.agent a left join api.access_log al on a.id = al.agent_user_id left join api.dataimport_dataset dd on a.id = dd.agent_user_id left join api.event_log el on a.id = el.agent_user_id where a.id = $1 and (al.agent_user_id is not null or dd.agent_user_id is not null or el.agent_user_id is not null) group by 1,2,3"

var sqlDeleteLookupUser = "SELECT user_id FROM api.agent where id=$1"

var sqlDeleteLookupAccessLog = "SELECT al.user_id as access_log,dd.user_id as permissiongroupuser,el.user_id as eventlog, ss.user_id as system_setting, pru.user_id as permissionrealmuser FROM api.user a left join api.access_log al on a.id = al.user_id left join api.permission_group_user dd on a.id = dd.user_id left join api.event_log el on a.id = el.user_id left join api.permission_realm_user pru on a.id = pru.user_id left join api.system_setting ss on a.id = ss.user_id WHERE a.id = $1 and (al.user_id is not null or dd.user_id is not null or el.user_id is not null or pru.user_id is not null or ss.user_id is not null) group by 1,2,3,4,5"

var sqlDeleteUser = "DELETE FROM api.user WHERE id=$1"

var sqlDeleteAgent = "DELETE FROM api.agent WHERE id=$1"
