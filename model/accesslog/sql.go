package accesslog

var sqlGetHistory = `SELECT  a.id,a.access_time,u.account as agent_user,u2.account as user,u3.account as server_agent_user,s.name as service,ltsm.name as service_method,a.host,a.request_url,a.access_duration,a.reply_code, a.client_ip 
					FROM api.access_log a left join api.user u on a.agent_user_id = u.id left join api.user u2 on a.user_id = u2.id left join api.user u3 on a.server_agent_user_id = u3.id left join api.service s on a.service_id = s.id left join api.lt_service_method ltsm on s.service_method_id = ltsm.id `

var sqlGetServiceName = "SELECT s.id, s.name, sm.name, smo.name FROM api.service s LEFT JOIN api.lt_service_method sm ON s.service_method_id=sm.id LEFT JOIN api.service_module smo ON s.service_module_id=smo.id ORDER BY smo.name,s.name,sm.name"

var sqlGetAgentName = "Select u.id , u.account from api.user u where user_type_id = 2 and deleted_at IS NULL Order by id asc"

var sqlGetServiceMethod = "Select id , name from api.lt_service_method Order by id asc"

//var SQL_GetOrderDetailLog = `SELECT access_time, client_ip, access_duration, reply_code, reply_reason FROM api.access_log WHERE request_params->>'id' = $1 AND service_id = 107
//ORDER BY access_time`

var SQL_GetOrderDetailLog = `SELECT  a.id,a.access_time,u.account as agent_user,u2.account as user,u3.account as server_agent_user,s.name as service,ltsm.name as service_method,a.host,a.request_url,a.access_duration,a.reply_code, a.client_ip, reply_reason 
					FROM api.access_log a left join api.user u on a.agent_user_id = u.id left join api.user u2 on a.user_id = u2.id left join api.user u3 on a.server_agent_user_id = u3.id left join api.service s on a.service_id = s.id left join api.lt_service_method ltsm on s.service_method_id = ltsm.id 
					WHERE request_params->>'id' = $1 AND service_id = 107 --WHERE ORDER BY access_time
					`
