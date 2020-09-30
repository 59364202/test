package file

import (
	"io"
	//	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

var (
	//	/data/thaiwater/thaiwaterdata
	UploadPath string = filepath.Join("/data", "thaiwater", "thaiwaterdata")
	//	UploadPath string = filepath.Join("D:/", "www_data", "haii-thaiwater30")
	DataserviceUploadLetterPath = "dataservice/"

	UploadDirPerm  uint32 = 0755
	UploadFilePerm uint32 = 0644
)

func SaveFile(file multipart.File, path, saveFileName string) error {
	fp := filepath.Join(UploadPath, path)

	if b, err := Exists(fp); err != nil || !b {
		err = os.MkdirAll(fp, os.FileMode(UploadDirPerm))
		if err != nil {
			return err
		}
	}

	fp = filepath.Join(fp, saveFileName)
	if _, err := Exists(fp); err != nil {
		os.Remove(fp)
	}
	dx, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, os.FileMode(UploadFilePerm))
	if err != nil {
		return err
	}
	defer dx.Close()

	file.Seek(0, 0)
	if _, err := io.Copy(dx, file); err != nil {
		return err
	}

	return nil
}

// exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
