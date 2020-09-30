package dam_daily_sum_by_region

// ====== Data by region ======
type DamSumByRegionInput struct {
	Region_id   string 	`json:"region_id"`	// รหัสภาคกรมอุตุ  เช่น 1
	Year 		[]int64 `json:"year"` 		// ปีสำหรับดึงข้อมูล เช่น [2015,2016]
}

type DamSumByRegionBound struct {
	Upper		float64 	`json:"upper"`		// ค่าเก็บกักสูงสุด เช่น 1234.50
	Lower		float64 	`json:"lower"`		// ค่าเก็บกักสูงสุด เช่น 1234.50
	Normal		float64 	`json:"normal"`		// ค่าเก็บกักสูงสุด เช่น 1234.50
}

type DamSumByRegionByYear struct {
	// Region_id   string 						`json:"region_id"`	// รหัสภาคกรมอุตุ  เช่น 1
	Year		int64 						`json:"year"`		// ปีข้อมูล เช่น 2019
	Data        []*DamSumByRegionGraphData 	`json:"data"`      	// ข้อมูลกราฟ	
}

type DamSumByRegionGraphData struct {
	DamDate         string          `json:"dam_date"`           	 // example:`2006-01-02` วันที่เก็บข้อมูล
	DamStorage      interface{}     `json:"total_dam_storage"`       // example:`140` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
	DamInflow       interface{}     `json:"total_dam_inflow"`        // example:`10` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	DamReleased     interface{}     `json:"total_dam_released"`      // example:`11` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
	DamUsesWater    interface{}     `json:"total_dam_uses_water"`    // ปริมาณน้ำที่ใช้ได้ (ล้าน ลบ.ม.)
	DamInflowAcc    interface{}     `json:"total_dam_inflow_acc"`   // example:`10` ปริมาณน้ำไหลเข้าอ่างสะสม (ล้าน ลบ.ม)
	DamReleasedAcc  interface{}     `json:"total_dam_released_acc"` // example:`11` ปริมาณการระบายผ่านสะสม (ล้าน ลบ.ม.)
}

type DamSumByRegionByYearOutput struct {
	GraphData   []*DamSumByRegionByYear		`json:"graph_data"`	// ข้อมูลกราฟ
	Bound		*DamSumByRegionBound 		`json:"bound"`		// ข้อมูล metadata
}
// ====== END Data by region ======

// ====== Compare data same date each year ======
type DamCompareSumByRegionInput struct {
	Region_id   string 	`json:"region_id"`	// รหัสภาคกรมอุตุ  เช่น 1
	Day     	string  `json:"day"`   		// วันที่เข้อมูล 1-31
	Month		string	`json:"month"`   	// เดือนข้อมูล 1-12
}

type DamCompareSumByRegionOutput struct {
	DamYear			int64 			`json:"year"`					// ปีข้อมูล เช่น 2019
	DamDate     	string  		`json:"date"`         			// example:`2006-01-02` วันที่เก็บข้อมูล
	DamStorage      interface{}     `json:"total_dam_storage"`      // example:`140` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.)
	DamInflow       interface{}     `json:"total_dam_inflow"`       // example:`10` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม)
	DamReleased     interface{}     `json:"total_dam_released"`     // example:`11` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.)
	DamUsesWater    interface{}     `json:"total_dam_uses_water"`   // ปริมาณน้ำที่ใช้ได้ (ล้าน ลบ.ม.)
	DamInflowAcc    interface{}     `json:"total_dam_inflow_acc"`   // example:`10` ปริมาณน้ำไหลเข้าอ่างสะสม (ล้าน ลบ.ม)
	DamReleasedAcc  interface{}     `json:"total_dam_released_acc"` // example:`11` ปริมาณการระบายผ่านสะสม (ล้าน ลบ.ม.)
}
// ====== END Compare data same date each year ======