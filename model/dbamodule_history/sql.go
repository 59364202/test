package dbamodule_history

import ()

var sqlSelectHistory = ` SELECT h.id
							  , h.created_at
							  , h.month
							  , h.dba_remark
							  , h.created_by
							  , u.full_name
						FROM public.dbamodule_history h
						LEFT JOIN api.user u ON h.created_by = u.id
						WHERE table_name = $1 AND year = $2
						ORDER BY h.created_at DESC `

var sqlInsertHistory = ` INSERT INTO public.dbamodule_history (table_name, year, month, dba_remark, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $5, NOW(), NOW()) RETURNING id; `
