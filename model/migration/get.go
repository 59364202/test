package migrate

import (
	"database/sql"
	//	"fmt"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/pqx"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"haii.or.th/api/server/model/backgroundjob"
)

func GetUrlImageNHC(q string) ([]string, error) {

	//	db, err := sql.Open("postgres", "postgres://nhc:nhcdbadmin@master1.nhc.in.th:5432/nhc?sslmode=disable")
	db, err := OpenNhc()
	if err != nil {
		return nil, err
	}
	//	defer db.Close()
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	img := make([]string, 0)
	for rows.Next() {
		var media_path sql.NullString
		rows.Scan(&media_path)
		if media_path.String != "" {
			img = append(img, media_path.String)
		}
	}
	return img, nil
}

func CountResponseCode200(img []string) (int64, int64) {

	var count200, countNot200 int64

	for _, url := range img {
		resp, err := http.Get(url)
		if err != nil {
			countNot200++
		} else if resp.StatusCode != 200 {
			countNot200++
			resp.Body.Close()
		} else {
			count200++
			resp.Body.Close()
			//			fmt.Println(url)
		}
		//		fmt.Println(url)
	}

	return count200, countNot200
}

func RegenDataImg() (interface{}, error) {
	//	log.WrapPanicGO(func() { UpdateImg() })
	_, err := backgroundjob.RunBackgroundJob(Cmd, Cmd_RegenDataImg)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func RegenDataImg1() (interface{}, error) {
	log.WrapPanicGO(func() { UpdateCVRow() })
	//	if err != nil {
	//		return nil, err
	//	}
	return nil, nil
}

func RegenDataImg2() (interface{}, error) {
	log.WrapPanicGO(func() { countFileTW30() })
	//	if err != nil {
	//		return nil, err
	//	}
	return nil, nil
}

func RegenDataImg3() (interface{}, error) {
	log.WrapPanicGO(func() { CalRowDiff() })
	//	if err != nil {
	//		return nil, err
	//	}
	return nil, nil
}

func CalRowDiff() error {
	q := "SELECT (tw_cv_ss+tw_cv_err)-nhc_row,media_type_id FROM migrate_log.summary_img "
	p := []interface{}{}
	qupdate := "UPDATE migrate_log.summary_img SET rec_diff=$2,last_update_media_type=NOW() WHERE media_type_id=$1"
	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	rows, err := db.Query(q, p...)
	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			rowdiff    int64
			media_type int64
		)
		rows.Scan(&rowdiff, &media_type)
		p1 := []interface{}{media_type, rowdiff}
		err := updateImgTableByMediaType(qupdate, p1)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateImg() error {
	q := "UPDATE migrate_log.summary_img SET nhc_row=$2,nhc_file_count=$3, last_update_media_type=NOW() WHERE media_type_id=$1"
	msql := getImgSql()
	for k, v := range msql {
		UpdateTimeImg()
		a, err := GetUrlImageNHC(v)
		//		fmt.Println(err)
		if err != nil {
			continue
		}
		b, _ := CountResponseCode200(a)

		p := []interface{}{k, len(a), b}
		err = updateImgTableByMediaType(q, p)
		if err != nil {
			return err
		}

	}
	UpdateCVRow()
	countFileTW30()
	CalRowDiff()
	//	UpdateTimeImg1()

	return nil
}
func UpdateTimeImg1() error {
	q := "UPDATE migrate_log.summary_img SET last_import_date='2017-01-01 00:00:00' WHERE media_type_id=$1"
	p := []interface{}{1}
	err := updateImgTableByMediaType(q, p)

	return err
}

func UpdateTimeImg() error {
	q := "UPDATE migrate_log.summary_img SET last_import_date=NOW() WHERE media_type_id=$1"
	p := []interface{}{1}
	err := updateImgTableByMediaType(q, p)
	return err
}

var prefixPath = "/data/thaiwater/thaiwaterdata/data/"

func countFileTW30() (interface{}, error) {
	q := "UPDATE migrate_log.summary_img SET tw30_file_count=$2 ,last_update_media_type=NOW() WHERE media_type_id=$1"
	a := getPath()
	for k, _ := range a {
		b, err := countFile(a[k])
		if err != nil {
			return nil, err
		}
		p := []interface{}{k, b}
		err = updateImgTableByMediaType(q, p)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

var countf int64

func countFile(path string) (int64, error) {
	countf = 0
	filepath.Walk(prefixPath+path, printFile)

	return countf, nil
}

var rex = regexp.MustCompile("thumb-*")

func printFile(path string, info os.FileInfo, err error) error {

	if err != nil {
		return nil
	}
	if !info.IsDir() {
		if !rex.MatchString(path) {
			countf++
		}
	}
	return nil
}

func updateImgTableByMediaType(q string, p []interface{}) error {

	db, err := pqx.Open()
	if err != nil {
		return errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	stmt, err := db.Prepare(q)
	if err != nil {
		return pqx.GetRESTError(err)
	}

	res, err := stmt.Exec(p...)
	if err != nil {
		return pqx.GetRESTError(err)
	}
	_, err = res.RowsAffected()

	if err != nil {
		return pqx.GetRESTError(err)
	}

	return nil
}

func UpdateCVRow() (interface{}, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := "SELECT sum(convert_success_row), sum(convert_error_row) FROM api.dataimport_dataset_log where deleted_at is null and dataimport_dataset_id IN "
	p := []interface{}{}
	qupdate := "UPDATE migrate_log.summary_img SET tw_cv_ss=$2, tw_cv_err=$3 WHERE media_type_id=$1"
	mds := getImgDataset()
	for k, v := range mds {
		p = []interface{}{}
		condition := "("
		for j, o := range v {
			if j >= 1 {
				condition += ",$" + strconv.Itoa(j+1)
			} else {
				condition += "$" + strconv.Itoa(j+1)
			}
			p = append(p, o)
		}
		condition += ")"

		rows, err := db.Query(q+condition, p...)
		if err != nil {
			return nil, pqx.GetRESTError(err)
		}

		for rows.Next() {
			var (
				cvss  int64
				cverr int64
			)
			rows.Scan(&cvss, &cverr)
			p1 := []interface{}{k, cvss, cverr}
			err := updateImgTableByMediaType(qupdate, p1)
			if err != nil {
				return nil, err
			}
		}
	}

	return nil, nil
}

func getPath() map[int64]string {
	a := map[int64]string{}
	a[1] = "product/image/countour_image/humidity/haii/media"
	a[2] = "product/image/countour_image/temperature/haii"
	a[3] = "product/image/countour_image/pressure/haii"
	a[4] = "product/image/rain_accumulate/thailand/haii"
	a[5] = "product/image/rain_accumulate/asia/haii"
	a[6] = "product/image/rain_accumulate/southeast_asia/haii"
	a[7] = "product/image/wind/asia/haii"
	a[8] = "product/image/wind/southeast_asia/haii"
	a[9] = "product/image/wind/thailand/haii"
	a[10] = "product/image/upper_wind_600m/asia"
	a[11] = "product/image/upper_wind_1500m/asia/haii"
	a[14] = "product/image/sst_w/global/haii"
	a[15] = "product/image/pressure_600m/asia/haii/media"
	a[16] = "product/image/pressure_1500m/asia/haii/media"
	a[17] = "product/image/precipitation/thailand/haii"
	a[18] = "product/image/precipitation/asia/haii"
	a[19] = "product/image/precipitation/southeast_asia/haii"
	a[23] = "product/image/wave_height/south_east_asia/hd"
	a[24] = "product/image/wave_height/south_china_sea/hd"
	a[25] = "product/image/wave_height/indian_ocean_northern/hd"
	a[26] = "product/image/sea_surface_elevation/global/hd"
	a[27] = "product/image/weather_map/thailand"
	a[28] = "product/image/upper_wind_850hpa/thailand/tmd"
	a[29] = "product/image/upper_wind_925hpa/thailand/tmd"
	a[30] = "product/image/radar"
	a[32] = "product/image/precip_forecast/global/ham"
	a[33] = "product/image/radar/coastal/gistda"
	a[140] = "product/image/ssh_w/global/haii"
	a[141] = "product/image/himawari-8/thailand"
	a[142] = "product/image/ssh_event/global/haii"
	a[149] = "product/image/modis/precipitaion/usda"
	a[150] = "product/image/modis/precipitaion/usda"
	a[151] = "product/image/modis/precipitaion/usda"
	a[152] = "product/image/modis/precipitaion/usda"
	a[157] = "product/image/satellite_rain/10km/haii"

	return a
}

func getImgSql() map[int64]string {
	a := map[int64]string{}
	a[1] = NhcMediaType1
	a[2] = NhcMediaType2
	a[3] = NhcMediaType3
	a[4] = NhcMediaType4
	a[5] = NhcMediaType5
	a[6] = NhcMediaType6
	a[7] = NhcMediaType7
	a[8] = NhcMediaType8
	a[9] = NhcMediaType9
	a[10] = NhcMediaType10
	a[11] = NhcMediaType11
	a[14] = NhcMediaType14
	a[15] = NhcMediaType15
	a[16] = NhcMediaType16
	a[17] = NhcMediaType17
	a[18] = NhcMediaType18
	a[19] = NhcMediaType19
	a[23] = NhcMediaType23
	a[24] = NhcMediaType24
	a[25] = NhcMediaType25
	a[26] = NhcMediaType26
	a[27] = NhcMediaType27
	a[28] = NhcMediaType28
	a[29] = NhcMediaType29
	a[30] = NhcMediaType30
	a[32] = NhcMediaType32
	a[33] = NhcMediaType33
	a[140] = NhcMediaType140
	a[141] = NhcMediaType141
	a[142] = NhcMediaType142
	a[149] = NhcMediaType149
	a[150] = NhcMediaType150
	a[151] = NhcMediaType151
	a[152] = NhcMediaType152
	a[157] = NhcMediaType157
	return a
}

func getImgDataset() map[int64][]interface{} {
	a := map[int64][]interface{}{}
	a[1] = []interface{}{344}
	a[2] = []interface{}{352}
	a[3] = []interface{}{353, 480}
	a[4] = []interface{}{354, 463}
	a[5] = []interface{}{355, 464}
	a[6] = []interface{}{356, 465}
	a[7] = []interface{}{361, 436}
	a[8] = []interface{}{362, 437}
	a[9] = []interface{}{363, 438}
	a[10] = []interface{}{364, 474, 478, 479, 483, 484, 485, 486}
	a[11] = []interface{}{365, 475}
	a[14] = []interface{}{374}
	a[15] = []interface{}{369, 476}
	a[16] = []interface{}{370, 477}
	a[17] = []interface{}{373, 456, 461, 462}
	a[18] = []interface{}{371, 454, 457, 458}
	a[19] = []interface{}{372, 455, 459, 460}
	a[23] = []interface{}{382, 383}
	a[24] = []interface{}{385, 386}
	a[25] = []interface{}{387, 428}
	a[26] = []interface{}{429, 430}
	a[27] = []interface{}{376, 384, 432, 433, 434, 435}
	a[28] = []interface{}{466, 467, 468, 469}
	a[29] = []interface{}{471, 473}
	a[30] = []interface{}{349, 439, 440, 441, 442, 443, 444, 445, 446, 447, 448, 449, 450, 451, 452, 453}
	a[32] = []interface{}{481}
	a[33] = []interface{}{482}
	a[140] = []interface{}{366}
	a[141] = []interface{}{368, 348, 351}
	a[142] = []interface{}{367}
	a[149] = []interface{}{357}
	a[150] = []interface{}{358}
	a[151] = []interface{}{359}
	a[152] = []interface{}{360}
	a[157] = []interface{}{375}

	return a
}
