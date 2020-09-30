package main

import (
	"database/sql"
	"fmt"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"
	"strconv"
)

func init() {
	runner.RegisterCommand("dummy", &DummyCmd{})
}

type DummyCmd struct {
}

func (cmd *DummyCmd) Run(args []string) error {
	// This should go to server log
	log.Logf("tw30bgjob dummy:Start background process ...")

	db, err := pqx.Open()
	if err != nil {
		return errors.Repack(err)
	}

	q := "select inet_server_addr(),inet_server_port(),version(),current_user"
	var ip, version, user sql.NullString
	var port sql.NullInt64

	r := db.QueryRow(q)
	if err = r.Scan(&ip, &port, &version, &user); err != nil {
		return errors.Repack(err)
	}

	fmt.Printf("tw30bgjob dummy: connection to %s:%d using user %s\n",
		ip.String, port.Int64, user.String)
	fmt.Printf("tw30bgjob dummy: PostgreSQL version is: %s\n", version.String)
	fmt.Printf("tw30bgjob dummy: 1st argument to command is: %s\n", args[0])
	fmt.Printf("tw30bgjob dummy: 2nd argument to command is: %s\n", args[1])

	s := args[10]

	//fmt.Printf("P: %s,%s -> %s\n", args[0], args[1], args[10])
	fmt.Printf("Hello %s\n", s)

	return nil
}

func (cmd *DummyCmd) GetUsage() string {
	return `test_p1 test_p2
	Example command for tw30bgjob
		test_p1 string parameter
		test_p2 integer parameter
`
}
func (cmd *DummyCmd) ValidateArgs(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing command parameter")
	}

	if _, err := strconv.ParseInt(args[1], 10, 64); err != nil {
		return fmt.Errorf("second parameter must be an interger")
	}

	return nil
}
