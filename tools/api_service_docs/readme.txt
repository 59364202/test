api doc ของระบบให้บริการข้อมูล

รันที่เครื่องตัวเอง
รันที่ api\thaiwater30\tools\api_service_docs\main.go

set run configuration
-config=src/haii.or.th/test/db.conf  // connnection config
-out=src/haii.or.th/api/thaiwater30/tools/api_service_docs/_test.go // output file

เปิดไฟล์ main run 
ได้ไฟล์ _test.go เอาเนื้่อไฟล์ไปใส่ใน

thaiwater30/service/api_service/apidocs.go

upload to git