package setting

import ()

var sqlUpdateSystemSetting = ` UPDATE api.system_setting
									  SET value = $4
									    , updated_by = $5
									    , updated_at = NOW()
									  WHERE user_id = $1
										AND is_public = $3 
									    AND name = $2 
									    AND deleted_at IS NULL 
									    AND deleted_by IS NULL `

var sqlInsertSystemSetting = ` INSERT INTO api.system_setting (
										      user_id, name, is_public, value, description, created_by, created_at, updated_by, updated_at)
                                      VALUES ($1, $2, $3, $4, $5, $6, NOW(), $6, NOW()) RETURNING id `

var sqlDeleteSystemSetting = ` DELETE FROM api.system_setting WHERE name = $1 AND deleted_at IS NULL AND deleted_by IS NULL `
