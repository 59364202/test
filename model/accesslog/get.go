// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package accesslog is a model for api.accesslog table. This table store accesslog information.
package accesslog

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/server/model/setting"
	result "haii.or.th/api/thaiwater30/util/result"
	selectOption "haii.or.th/api/thaiwater30/util/selectoption"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"

	"haii.or.th/api/util/rest"
)

// Get History Accesslog return accesslog details by service name and date range
//  Parameters:
//		dateStart
//					Start Date access log.
//		dateEnd
//					End Date access log.
//		ServiceName
//					Service name.
//  Return:
//		Access log in date range.
func GetHistory(dateStart, dateEnd string, ServiceName, User string) (*result.Result, error) {

	// connect database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql for get Access log table
	q := sqlGetHistory
	p := []interface{}{}
	var c int
	sqlService := ""
	sqlUser := ""
	// create condition between datetime and service name
	if dateStart != "" && dateEnd != "" && ServiceName != "" && User != "" {
		p = []interface{}{dateStart, dateEnd}
		sqlService, p, c = selectMultiService(3, ServiceName, p)
		sqlUser, p = selectMultiUser(c, User, p)
		q = q + " WHERE a.created_at >= $1 AND a.created_at <= $2 AND " + sqlService + sqlUser
	} else if dateStart != "" && dateEnd != "" && User != "" {
		p = []interface{}{dateStart, dateEnd}
		sqlUser, p = selectMultiUser(3, User, p)
		q = q + " WHERE a.created_at >= $1 AND a.created_at <= $2" + sqlUser

	} else if dateStart != "" && dateEnd != "" && ServiceName != "" {
		// case select multi service name
		p = []interface{}{dateStart, dateEnd}
		sqlService, p, c = selectMultiService(3, ServiceName, p)
		q = q + " WHERE a.created_at >= $1 AND a.created_at <= $2 AND " + sqlService
	} else if dateStart != "" && dateEnd != "" && ServiceName == "" {
		// case no input service
		q = q + " WHERE a.created_at >= $1 AND a.created_at <= $2"
		p = []interface{}{dateStart, dateEnd}
	} else if dateStart != "" && dateEnd == "" && ServiceName == "" {
		// case no input end date, no service
		q = q + " WHERE a.created_at >= $1 "
		p = []interface{}{dateStart}
	} else if dateStart == "" && dateEnd != "" && ServiceName == "" {
		// case no input start date, no service
		q = q + " WHERE a.created_at <= $1 "
		p = []interface{}{dateEnd}
	} else if dateStart == "" && dateEnd == "" && ServiceName != "" {
		// case no input start and end date
		p = []interface{}{}
		sqlService, p, c = selectMultiService(1, ServiceName, p)
		q = q + " WHERE " + sqlService
	} else if dateStart != "" && dateEnd == "" && ServiceName != "" {
		// case no input end date
		p = []interface{}{dateStart}
		sqlService, p, c = selectMultiService(2, ServiceName, p)
		q = q + " WHERE a.created_at >= $1 AND " + sqlService

	} else if dateStart == "" && dateEnd != "" && ServiceName != "" {
		// case no start date
		p = []interface{}{dateEnd}
		sqlService, p, c = selectMultiService(2, ServiceName, p)
		q = q + " WHERE a.created_at <= $1 AND " + sqlService
	}

	// process query access table
	rows, err := db.Query(q, p...)
	// check error
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	//declare value for get data from table
	var (
		id                sql.NullInt64
		access_time       time.Time
		agent_user        sql.NullString
		user              sql.NullString
		server_agent_user sql.NullString
		service           sql.NullString
		service_method    sql.NullString
		host              sql.NullString
		request_url       sql.NullString
		access_duration   sql.NullInt64
		reply_code        sql.NullInt64
		client_ip         sql.NullString
	)

	data := make([]*ResultAccessLog, 0)
	// loop for get data from table to struct for return
	for rows.Next() {
		accessLog := new(ResultAccessLog)
		// scan data from database
		rows.Scan(&id, &access_time, &agent_user, &user, &server_agent_user, &service, &service_method, &host, &request_url, &access_duration, &reply_code, &client_ip)
		accessLog.ID = id.Int64
		accessLog.AccessTime = access_time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		accessLog.AgentUser = agent_user.String
		accessLog.User = user.String
		accessLog.ServerAgentUser = server_agent_user.String
		accessLog.Service = service.String
		accessLog.ServiceMethod = service_method.String
		accessLog.Host = host.String
		accessLog.RequestURL = request_url.String
		accessLog.AccessDuration = access_duration.Int64
		accessLog.ReplyCode = reply_code.Int64
		accessLog.ClientIP = client_ip.String
		// add data row to array result
		data = append(data, accessLog)
	}

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	// return data access log
	return result.Result1(data), nil
}

// make condition for select servicename return sql get service id
//  Parameters:
//		count
//					number of parameters in sql.
//		p
//					value of parameters
//		ServiceName
//					Service name.
//  Return:
//		sql with condition for get data
func selectMultiService(count int, service string, p []interface{}) (string, []interface{}, int) {
	arrService := strings.Split(service, ",")
	// arrUser := strings.Split(user, ",")
	var c int
	q := "("

	for i, v := range arrService {
		//		log.Println(v)
		vInt, _ := strconv.Atoi(v)
		p = append(p, vInt)
		if i == 0 {
			q = q + "s.id = $" + strconv.Itoa(count)
		} else {
			q = q + " OR s.id = $" + strconv.Itoa(count)
		}
		count = count + 1
		c = count
	}
	q = q + ")"

	return q, p, c
}

func selectMultiUser(count int, user string, p []interface{}) (string, []interface{}) {
	arrUser := strings.Split(user, ",")
	q := "AND ("

	for i, v := range arrUser {
		//		log.Println(v)
		vInt, _ := strconv.Atoi(v)
		p = append(p, vInt)
		if i == 0 {
			q = q + "u2.id = $" + strconv.Itoa(count)
		} else {
			q = q + " OR u2.id = $" + strconv.Itoa(count)
		}
		count = count + 1
	}
	q = q + ")"

	return q, p
}

// get service name from api.service return service name and limit date range
//  Parameters:
//		None
//  Return:
//		Service name
func GetServiceName() (*result.Result, error) {

	// connect database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql for get Service name from api.service
	q := sqlGetServiceName
	p := []interface{}{}
	// process query access table
	rows, err := db.Query(q, p...)
	// check error
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	var data = selectOption.NewSelect()
	// loop for get data from table to struct for return
	for rows.Next() {
		//declare value for get data from table
		var (
			id     sql.NullInt64
			text   sql.NullString
			method sql.NullString
			module sql.NullString
		)
		rows.Scan(&id, &text, &method, &module)
		// add data one row to array data
		data.Add(id.Int64, module.String+"/"+text.String+"("+method.String+")")
	}
	//declare value for result
	type ServiceName struct {
		DateRange     string                 `json:"date_range"`
		ServiceOption []*selectOption.Option `json:"service_name"`
	}

	dataOption := &ServiceName{}
	dataOption.ServiceOption = data.Option
	// get date range
	dataOption.DateRange = setting.GetSystemSetting("setting.Default.DateRange")
	// return service
	return result.Result1(dataOption), nil
}

// get accesslog service id 107 by order detail id
func GetOrderDetailLog(orderDetailID int64, dateString, dateEnd string) ([]*ResultAccessLog, error) {
	if orderDetailID <= 0 {
		return nil, rest.NewError(422, "invalid order_detail_id", nil)
	}
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	param := []interface{}{orderDetailID}
	strWhere := ""
	if dateString != "" {
		param = append(param, dateString)
		strWhere += " AND a.created_at >= $" + strconv.Itoa(len(param))
	}
	if dateEnd != "" {
		param = append(param, dateString+" 23:59:59")
		strWhere += " AND a.created_at <= $" + strconv.Itoa(len(param))
	}

	if strWhere != "" {
		SQL_GetOrderDetailLog = strings.Replace(SQL_GetOrderDetailLog, "--WHERE", strWhere, 1)
	}

	rows, err := db.Query(SQL_GetOrderDetailLog, param...)
	if err != nil {
		return nil, err
	}

	var (
		id                sql.NullInt64
		access_time       time.Time
		agent_user        sql.NullString
		user              sql.NullString
		server_agent_user sql.NullString
		service           sql.NullString
		service_method    sql.NullString
		host              sql.NullString
		request_url       sql.NullString
		access_duration   sql.NullInt64
		reply_code        sql.NullInt64
		client_ip         sql.NullString
		reply_resaon      sql.NullString
	)

	data := make([]*ResultAccessLog, 0)
	// loop for get data from table to struct for return
	for rows.Next() {
		accessLog := new(ResultAccessLog)
		// scan data from database
		rows.Scan(&id, &access_time, &agent_user, &user, &server_agent_user, &service, &service_method, &host, &request_url, &access_duration, &reply_code, &client_ip, &reply_resaon)
		accessLog.ID = id.Int64
		accessLog.AccessTime = access_time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		accessLog.AgentUser = agent_user.String
		accessLog.User = user.String
		accessLog.ServerAgentUser = server_agent_user.String
		accessLog.Service = service.String
		accessLog.ServiceMethod = service_method.String
		accessLog.Host = host.String
		accessLog.RequestURL = request_url.String
		accessLog.AccessDuration = access_duration.Int64
		accessLog.ReplyCode = reply_code.Int64
		accessLog.ClientIP = client_ip.String
		accessLog.ReplyReason = reply_resaon.String
		// add data row to array result
		data = append(data, accessLog)
	}

	return data, nil
}

func GetAgentName() (*result.Result, error) {

	// connect database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// sql for get Service name from api.service
	q := sqlGetAgentName
	p := []interface{}{}
	// process query access table
	rows, err := db.Query(q, p...)
	// check error
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*AgentName, 0)
	for rows.Next() {
		d := &AgentName{}
		var (
			id        sql.NullInt64
			agentName sql.NullString
		)
		rows.Scan(&id, &agentName)

		d.ID = id.Int64
		d.AgentName = agentName.String

		data = append(data, d)
	}

	// log.Println("333333333")
	return result.Result1(data), nil

}

