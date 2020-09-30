package spatial_thailand_boundary

import (
	"haii.or.th/api/util/errors"	// errors
	"haii.or.th/api/util/pqx"	// database connection
	"haii.or.th/api/util/rest"
	"strconv"
)

// Table : Spatial_Thailand_Boundary
type Struct_Spatial_Thailand_Boundary struct {
	Geocode string `json:"tambon_idn"` 		// example:240410` รหัส Geocode
	TambonName string `json:"tam_nam_t"` 	// example:ต.เกาะนางคำ` ชื่อตำบล
	AmphoeName string `json:"amphoe_t"` 	// example:อ.ปากพะยูน` ชื่ออำเภอ
	ProvinceName string `json:"prov_nam_t"` // example:จ.พัทลุง` ชื่อจังหวัด
}

type Struct_Geocode_Id struct {
	Id          	int64  `json:"id,omitempty"`        // example:`3` ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
	Geocode     	string `json:"geocode,omitempty"`      // example:`100101` รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	Province_name	string `json:"province_name,omitempty"`		// example:`กรุงเทพมหานคร` ชื่อจังหวัด
	Amphoe_name		string `json:"amphoe_name,omitempty"`			// example:`พระนคร` ชื่ออำเภอ
	Tumbon_name		string `json:"tumbon_name,omitempty"`			// example:`พระบรมมหาราชวัง` ชื่อตำบล
}

// Get geocode from lat, lon
func GetGeocode(lat, lon string) (*Struct_Geocode_Id, error) {
	lt_geocode := &Struct_Geocode_Id{}

	// validate parameters
	if len(lat) != 0 {
		_ , err := strconv.ParseFloat(lat, 64)
		if  err != nil  {
			return nil, rest.NewError(422, "Invalid latitude", nil)
		}
	} else {
		return nil, rest.NewError(422, "No latitude provided.", nil)
	}

	if len(lon) != 0 {
		_ , err := strconv.ParseFloat(lon, 64)
		if  err != nil  {
			return nil, rest.NewError(422, "Invalid longitude", nil)
		}
	} else {
		return nil, rest.NewError(422, "No longitude provided.", nil)
	}

	// connection
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	// main query
	q := `
	SELECT sp.tambon_idn, sp.tam_name_t, sp.amphoe_t, sp.prov_nam_t
	FROM spatial_thailand_boundary sp
	WHERE ST_Within (ST_PointFromText($1, 4326), geom)
	`

	// point parameter
	p := "POINT(" + lon + " " + lat + ")"

	boundary := &Struct_Spatial_Thailand_Boundary{}

	// query geocode
	err = db.QueryRow(q, p).Scan(&boundary.Geocode, &boundary.TambonName, &boundary.AmphoeName, &boundary.ProvinceName)

	if err != nil {
		return nil, errors.Repack(err)
	}

	// query geocode id from geocode
	if (len(boundary.Geocode) > 0) {
		q2 := `SELECT id, geocode, province_name->>'th' AS province_name, amphoe_name->>'th' AS amphoe_name, tumbon_name->>'th' AS tumbon_name
				FROM lt_geocode
				WHERE geocode = $1`
		p2 := boundary.Geocode

		err = db.QueryRow(q2, p2).Scan(&lt_geocode.Id, &lt_geocode.Geocode, &lt_geocode.Province_name, &lt_geocode.Amphoe_name, &lt_geocode.Tumbon_name)

		if err != nil {
			return nil, errors.Repack(err)
		}

		return lt_geocode, err
	} else {
		err := errors.New("Cannot find geocode from provided coodinate")

		return nil, err
	}
}
