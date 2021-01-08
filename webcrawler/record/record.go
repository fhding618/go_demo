package record

import (
	"fmt"
	"os"
)

// 文件名暂不允许修改
const record_file = "webcrawler.log"

// 记录Level
type RecordLevel byte

const (
	RECORD_LEVEL_INFO RecordLevel = 0
	RECORD_LEVEL_WARN RecordLevel = 1
	RECORD_LEVEL_ERR  RecordLevel = 2
)

var recordLevelName = map[RecordLevel]string{
	RECORD_LEVEL_INFO: "INFO",
	RECORD_LEVEL_WARN: "WARN",
	RECORD_LEVEL_ERR:  "ERROR",
}

// 记录类型
type RecordTypz byte

const (
	CONSOLE_RECORD_TYPZ RecordTypz = 0
	FILE_RECORD_TYPZ    RecordTypz = 1
)

var recordTypzName = map[RecordTypz]string{
	CONSOLE_RECORD_TYPZ: "console_record",
	FILE_RECORD_TYPZ:    "file_record",
}

// 日志记录函数类型
// 参数level代表日志级别。级别设定：0-普通 1-警告 2-错误
type RecordFun func(level RecordLevel, content string)

// 数据日志记录
type Record interface {
	// 返回记录方式
	Typz() string
	// 记录内存
	Record(level RecordLevel, content string)
}

func NewRecord(typz RecordTypz) Record {
	return &fileRecord{
		typz: typz,
	}
}

type fileRecord struct {
	typz RecordTypz
	fd   *os.File
}

func (re *fileRecord) Typz() string {
	return recordTypzName[re.typz]
}

func (re *fileRecord) Record(level RecordLevel, content string) {
	if content == "" {
		return
	}
	fd, err := re.getFd()
	if err != nil {
		panic(fmt.Sprintf("Open record file failed: %v", err))

	}
	fd.WriteString(fmt.Sprintf("%s: %s\n", recordLevelName[level], content))
	return
}

func (re *fileRecord) getFd() (*os.File, error) {
	if re.fd == nil {
		f, err := os.OpenFile(record_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		re.fd = f
	}
	return re.fd, nil
}

type consoleRecord struct {
	typz RecordTypz
}

func (con *consoleRecord) Typz() string {
	return recordTypzName[con.typz]
}

func (con *consoleRecord) Record(level byte, content string) {
	return
}
