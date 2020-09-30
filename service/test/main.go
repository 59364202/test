// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

// package for service
//รวม registter service ต่าง ๆ ใน folder นี้
package test

import (
	"haii.or.th/api/util/service"
)

const (
	ModuleName     = "test"
	ServiceName    = ModuleName + "/"
	ServiceVersion = service.APIVersion1

	ServiceCim       = ServiceName + "cim"       // test/cim
	ServiceManorot   = ServiceName + "manorot"   // test/manorot
	ServiceJuiss     = ServiceName + "juiss"     // test/manorot
	ServiceFern      = ServiceName + "test_fern" // test/test_fern
	ServiceWerawan   = ServiceName + "werawan"
	ServiceThitiporn = ServiceName + "thitiporn" // test/thitiporn
	ServicePeerapong = ServiceName + "peerapong" // test/peerapong
)

func RegisterService(dpt service.Dispatcher) {

	cim := &Cim{}
	dpt.Register(ServiceVersion, service.MethodGET, ServiceCim, cim.handlerGetCim)

	// register service
	manorot := &Manorot{}
	dpt.Register(ServiceVersion, service.MethodGET, ServiceManorot, manorot.handlerGetManorot)

	// register service
	juiss := &Juiss{}
	dpt.Register(ServiceVersion, service.MethodGET, ServiceJuiss, juiss.handlerGetJuiss)

	werawan := &Werawan{}
	dpt.Register(ServiceVersion, service.MethodGET, ServiceWerawan, werawan.handlerGetWerawan)

	// register service
	fern := &Fern{}
	dpt.Register(ServiceVersion, service.MethodGET, ServiceFern, fern.handlerGetFern)
	
	// register service
	peerapong := &Peerapong{}
	dpt.Register(ServiceVersion, service.MethodGET, ServicePeerapong, peerapong.handlerGetPeerapong)
	
	// register service
	//	api2.thaiwater.net:9200/api/v1/test/thitiporn
	//	reference service struct in file thitiporn.go
	thitiporn := &ThitipornServiceProvinceStruc{}
	dpt.Register(ServiceVersion, service.MethodGET, ServiceThitiporn, thitiporn.handlerGetProvinceThitiporn)

}
