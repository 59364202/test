package main

import (
	"fmt"
	"haii.or.th/api/util/log"

	model "haii.or.th/api/thaiwater30/model/migration"
)

func init() {
	runner.RegisterCommand("regen", &RegenCmd{})
}

type RegenCmd struct {
}

func (cmd *RegenCmd) Run(args []string) error {
	log.Logf("tw30bgjob regen:%s background process ...", args[0])
	var err error

	switch args[0] {
	case model.Cmd_RegenData:
		err = model.UpdateData()
	case model.Cmd_RegenMasterData:
		err = model.UpdateMasterData()
	case model.Cmd_RegenDataImg:
		err = model.UpdateImg()
	default:
		log.Logf("tw30bgjob regen:%s no func", args[0])
		return nil
	}

	if err != nil {
		log.Logf("tw30bgjob regen error : %s", err.Error())
		return err
	}
	log.Logf("tw30bgjob regen:%s finish", args[0])

	return nil
}

func (cmd *RegenCmd) GetUsage() string {
	return `func_name
	regen command for tw30bgjob
		func_name string parameter exmaple:RegenDataImg, RegenData, RegenMasterData
`
}
func (cmd *RegenCmd) ValidateArgs(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing command parameter")
	}

	return nil
}
