package utils

import (
	"bytes"
	"crypto/md5"
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"time"
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// JSONTime format json time field by myself
type JSONDate struct {
	time.Time
}

// MarshalJSON on JSONDate format Time field with %Y-%m-%d %H:%M:%S
func (t JSONDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func MkDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		//文件夹不存在，创建
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// MD5 加密字符串
func GetMD5(plainText string) string {
	h := md5.New()
	h.Write([]byte(plainText))
	return hex.EncodeToString(h.Sum(nil))
}

//计算文件的md5,适用于本地大文件计算
func GetMd5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(md5hash.Sum(nil)), nil
}

//计算文件的md5，不能用于计算大文件流否则内存占用很大
//@return io.Reader @params file的副本
func GetMd52(file io.Reader) (io.Reader, string, error) {
	var b bytes.Buffer
	md5hash := md5.New()
	if _, err := io.Copy(&b, io.TeeReader(file, md5hash)); err != nil {
		return nil, "", err
	}
	return &b, hex.EncodeToString(md5hash.Sum(nil)), nil
}

func GetStackInfo() string {
	//return string(debug.Stack())
	line, funcName := 0, "???"
	pc, _, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcName = runtime.FuncForPC(pc).Name() // main.(*MyStruct).foo
		//funcName = filepath.Ext(funcName)            // .foo
		//funcName = strings.TrimPrefix(funcName, ".") // foo
		//funcName = filepath.Base(funcName)           // /full/path/basename.go => basename.go
	}
	return "funcName: " + funcName + "  line: " + strconv.FormatInt(int64(line), 10)
}
