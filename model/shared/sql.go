package shared

import ()

var sqlGetTableName = ` SELECT table_name
						FROM information_schema.tables
						WHERE table_schema='public' `

var sqlGetMetadataTable = ` SELECT c.table_name, obj_description(pg.oid) AS table_desc, string_agg(c.column_name, ',') AS column_list
							 FROM information_schema.columns c
							 LEFT JOIN pg_class pg ON c.table_name = pg.relname
							 WHERE c.table_schema = 'public'
							   AND EXISTS (SELECT tb.table_name
							               FROM information_schema.tables tb
							               WHERE tb.table_schema = 'public' AND left(tb.table_name,2) = 'm_'
							                 AND tb.table_name = c.table_name)
							   AND (c.column_name = 'geocode_id' OR c.column_name = 'subbasin_id' OR c.column_name = substring(c.table_name from '^m_(\w*)$') || '_lat' OR c.column_name = substring(c.table_name from '^m_(\w*)$') || '_long')
							AND relkind = 'r'
							GROUP BY c.table_name, obj_description(pg.oid)
							ORDER BY c.table_name  `
