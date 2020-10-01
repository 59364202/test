// Copyright (c) 2016-2018 CIM Systems (Thailand) Co., Ltd. (cim.co.th)
// All Rights Reserved
//
// This product is protected by copyright and distributed under
// licenses restricting copying, distribution and decompilation.

// Package datafile provides data file name validate and manipulation function.
//
package datafile

import (
	"io"
	"log"
	"os"
	"strings"

	"haii.or.th/api/server/model"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/filepathx"
)

var DefaultFileMode = os.FileMode(0640)
var DefaultFolderMode = os.FileMode(0750)

// GetAbsoluatePath returns an absolute path name of the given file name
// and path prefix after snaitized the given file name.
//	Parameter:
//		ps
//			A list of path name element
//  Return:
//		An aboulute path name of the given input or empty string if there
//		is an error.
//
func GetAbsoluatePath(ps ...string) (string, error) {
	fs := filepathx.JoinPath(ps...)
	for _, s := range strings.Split(fs, "/") {
		if filepathx.SanitizeFileName(s) != s {
			return "", errors.Newf("invalid characher in path name '%s'", fs)
		}
	}

	return filepathx.JoinPath(model.GetStoragePath(), fs), nil
}

// DecryptPath decrypts a path name using system crypter.
//	Parameter:
//		path
//			An encrypted path name
//  Return:
//		A decrypted path name or empty string if there is an error.
//
func DecryptPath(path string) (string, error) {
	p, err := model.GetCipher().DecryptText(path)
	return p, errors.Repack(err)
}

// EncryptPath encrypts a path name using system crypter.
//	Parameter:
//		path
//			A relative path name in encrypted format
//  Return:
//		An encrypted path name or empty string if there is an error
//
func EncryptPath(p string) (string, error) {
	if p == "" {
		return "", nil
	}

	p, err := model.GetCipher().EncryptText(p)
	return p, errors.Repack(err)
}

// HasPrefix determines whether the given path as the specified prefix
//  Parameter:
//		path
//			A path name
//		prefix
//			A path prefix
//  Return
//		true if the path name has a prefix
//
func HasPrefix(path, prefix string) bool {
	path = strings.TrimLeft(path, "/")
	prefix = strings.TrimLeft(prefix, "/")

	if !strings.HasPrefix(path, prefix) {
		return false
	}

	return len(path) == len(prefix) || path[len(prefix)] == '/'
}

// CopyFile copies a content from given source to the destination file name
//  Parameter:
//		prefix
//			A path name prefix
//		filename
//			A file name
//		src
//			A source content
//  Return
//		A relative name of destination file name or empty string if
//		there is an error
//
func CopyFile(prefix, filename string, src io.Reader) (string, error) {
	relname := filepathx.JoinPath(prefix, filename)
	absname, err := GetAbsoluatePath(relname)
	if err != nil {
		return "", err
	}

	os.MkdirAll(filepathx.Dir(absname), DefaultFolderMode)
	log.Println(absname)
	f, err := os.OpenFile(absname, os.O_WRONLY|os.O_CREATE, DefaultFileMode)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(f, src); err != nil {
		return "", err
	}
	return relname, nil
}

// CreatePath creates a new folder.
//  Parameter:
//		path
//			A relative path name
//  Return
//		A nil if there is an error.
//
func CreatePath(path string) error {
	abspath, err := GetAbsoluatePath(path)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepathx.Dir(abspath), DefaultFolderMode); err != nil {
		return err
	}

	return nil
}

func RemoveFileCheckPrefix(filename, pathprefix string) error {
	if pathprefix != "" {
		if !HasPrefix(filename, pathprefix) {
			return errors.New("Not a coal awards file")
		}
	}

	ap, err := GetAbsoluatePath(filename)
	if err != nil {
		return errors.Repack(err)
	}

	return errors.Repack(os.Remove(ap))
}

func Rename(o, n string) error {
	o, err := GetAbsoluatePath(o)
	if err != nil {
		return errors.Repack(err)
	}
	n, err = GetAbsoluatePath(n)
	if err != nil {
		return errors.Repack(err)
	}

	return errors.Repack(os.Rename(o, n))
}
