package media_animation

import ()

var SQL_SelectPreRainAnimation = `
SELECT md.media_path, 
       md.filename, 
       md.media_datetime 
FROM   (SELECT filename, 
               Max(media_datetime) AS media_datetime 
        FROM   media_animation 
        WHERE  filename IN ( 'ani_d03_large.mp4', 'ani_d02_large.mp4', 
                             'ani_d01_large.mp4' ) 
               AND media_type_id IN ( '80', '81', '82' ) 
               AND deleted_at = To_timestamp(0) 
        GROUP  BY filename 
        ORDER  BY filename DESC)t 
       INNER JOIN media_animation md
               ON md.media_datetime = t.media_datetime 
                  AND md.filename = t.filename 
                  AND deleted_at = To_timestamp(0) 
GROUP  BY md.media_path, 
          md.filename, 
          md.media_datetime 
ORDER BY md.filename
`
var SQL_SelectPreWaveAnimationMP4 = `
SELECT media_path, a.filename, a.media_datetime
FROM
  (SELECT MAX(media_datetime) AS media_datetime,
          filename
   FROM media_animation
   WHERE filename IN ('wave_168hr.mp4')
   GROUP BY filename) b
JOIN media_animation a ON b.media_datetime = a.media_datetime
AND media_type_id = 83  AND a.filename IN ('wave_168hr.mp4')
`

var SQL_SelectPreWaveAnimation = `
SELECT media_path, a.filename, a.media_datetime
FROM
  (SELECT MAX(media_datetime) AS media_datetime,
          filename
   FROM media_animation
   WHERE filename IN ('wave_168hr.gif')
   GROUP BY filename) b
JOIN media_animation a ON b.media_datetime = a.media_datetime
AND media_type_id = 83
`

var SQL_SelectPreWindAnimation = `
SELECT media_path, a.filename, a.media_datetime
FROM
  (SELECT MAX(media_datetime) AS media_datetime,
          filename
   FROM media_animation
		WHERE filename IN ('ani_d02_large.mp4')
   GROUP BY filename) b
JOIN media_animation a ON b.media_datetime = a.media_datetime
AND media_type_id = 183
`
