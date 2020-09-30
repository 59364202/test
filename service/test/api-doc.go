package test

import ()

// @Document		v1.test
// @Version			1.0
// @Title			WebService API for training
// @Description    	WebService ในกลุ่มนี้ เป็น WebService ที่ใช้สำหรับ ให้บริการข้อมูล ที่ผู้ใช้ ร้องขอ ผ่าน ระบบ ให้บริการข้อมูลแก่บุคลภายนอก
// @
// @				เมื่อผู้ใช้ซึ่งเป็นบุคลภายนอก ร้องข้อข้อมูล ระบบ จะทำการจัดเตรียมข้อมูล จากนั้น ระบบ จะทำการสร้าง จดหมายอิเล็คโทรนิค เพื่อ
// @    			แจ้งผู้ใช้ถึง URL ที่ ผู้ใช้ สามารถใช้เพื่อเข้าถึงข้อมูลที่ร้องขอได้
// @
// @				โดย WebService ที่ให้บริการข้อมูลเหล่านี้ ทุก service จะมี parameter พิเศษ หนึ่งตัวชื่อ eid ซึ่งจะถูกสร้าง จากระบบ
// @				โดยจะเป็น ค่าตัวอักษรที่ไม่ซ้ำกัน (unique RFC7515 Unpadded 'base64url') กับการร้องข้อข้อมูลอื่นก่อนหน้านี้
// @				เพื่อป้องกันไม่ให้คนอื่นซึ่งไม่ใช่ผู้ร้องขอข้อมูลเข้าถึงข้อมูลชุดนี้ได้
// @
// @				สามารถดู eid ได้จาก จดหมายอิเล็คโทรนิค ที่อยู่ใน ช่องทางการรับข้อมูล
// @TermsOfService 	http://www.haii.or.th/tos
// @ContactEmail    api@haii.or.th
// @License      	http://www.haii.or.th/license HAII License
// @ExternalDoc		http://swagger.io/swagger-ui/ Find out more about Swagger-UI
