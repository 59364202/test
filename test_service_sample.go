package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"haii.or.th/api/server"
	"haii.or.th/api/server/service"
	"io/ioutil"
	"os"
)

/*------------
   Sample test main
func main() {
	srv := server.New("TW30")

	//_, err := srv.GetServiceDispatcher()
	//if err != nil {
	//	os.Exit(1)
	//}
	//model_setting.SetSystemDefaultInt(realm.SettingAllowCreateUser, 1)

	TestServices(srv)
	//srv.Start()
}
*/

func TestServices(srv *server.Server) {
	srv.SetTestMode(true)

	// Testing for "GET /api/v1/dataimport/1/setting"
	ctx, err := srv.NewTestServiceRequestContext("POST", "/api/v1/dataimport/1/23/importdata")
	ctx.SetUserInfo(9, 0)
	path := "I:/TW30DATA/BMA/2016/08/28/012215"
	filename := "cv-20160828012215-Canal.csv"

	fn, err := os.Open(path + "/" + filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	content, err := ioutil.ReadAll(fn)
	fn.Close()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	ctx.AddUploadFile("canal_data", filename, content)

	//ctx.AddUploadFile("canal_data", filename, content)
	TestServiceRunDumpResult(ctx)

}

func TestServiceRunDumpResult(ctx *service.TestServiceRequestContext) {
	resp, body := ctx.TestService()

	// Get result
	fmt.Printf("Response Code :%d\n", resp)

	var pj bytes.Buffer
	if json.Indent(&pj, body, "", "    ") == nil {
		fmt.Printf("Response Body : \n")
		fmt.Println(pj.String())
	} else {
		fmt.Printf("Response Body : \n")
		fmt.Println(string(body))
	}
}
