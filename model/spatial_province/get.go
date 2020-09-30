package spatial_province

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

type Struct_Province struct {
	Province_id   int64  `json:"province_id"`   // example:`10` รหัสจังหวัด
	Region_id     int64  `json:"region_id"`     // example:`2` รหัสภาค
	Province_name string `json:"province_name"` // example:`กรุงเทพมหานคร` ชื่อจังหวัด
}

func GerProv(lat, long string) (*Struct_Province, error) {
	if lat == "" || long == "" {
		return nil, errors.New("invalid lat long")
	}

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	q := `
	SELECT sp.prov_code, g.area_code, g.province_name->>'th' 
	FROM spatial_province sp
	INNER JOIN lt_geocode g ON g.province_code = sp.prov_code
	WHERE ST_Contains(sp.geom, ST_GeomFromText($1, 4326))
	AND (g.amphoe_code = '  ' AND g.tumbon_code = '  ' OR g.amphoe_name->>'th' = '' AND g.tumbon_name->>'th' = '')
	`
	p := "POINT(" + long + " " + lat + " )"

	prov := &Struct_Province{}
	err = db.QueryRow(q, p).Scan(&prov.Province_id, &prov.Region_id, &prov.Province_name)
	if err != nil {
		return nil, errors.Repack(err)
	}

	return prov, nil
}
