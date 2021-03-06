package foundation

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type FuncMap map[string]interface{}

func Create(args []string) {
	table := args[0]
	record := args[1:]
	createDir()
	createFile(table, record...)
	// ディレクトリの生成

}

func createDir() {
	os.MkdirAll("./model", os.ModeDir)
	os.MkdirAll("./db", os.ModeDir)
}

func createFile(table string, record ...string) {

	var attributes string
	tmp, err := template.ParseFiles("./template/foundation.go.tmpl")
	isPanic(err, "テンプレートが読み込めませんでした。")

	for _, v := range record {
		value, err := createAttributes(v)
		attributes += value
		isPanic(err, "createAttributesでエラーを起こしました。")
	}

	var builtins = FuncMap{
		"ModelName":  table,
		"Attributes": attributes,
	}
	// ファイルの書き込み
	file, err := os.Create("./model/" + table + ".go")
	isPanic(err, "モデルが作成できませんでした。")
	err = tmp.Execute(file, builtins)
	isPanic(err, "ファイルへ書き込めませんでした。")
	file.Close()
}

// craeteAttributes Table:Attribute:primary_key
func createAttributes(data string) (string, error) {
	work := strings.Split(data, ":")

	if work[1] == "datetime" || work[1] == "timestamp" {
		work[1] = "time.Time"
	}

	data = "\t" + work[0] + " " + work[1]
	if len(work) > 2 {
		data += setProps(work[2:]...)
	}

	data += "\n"
	return data, nil
}

// craeteAttributes Table:Attribute:primary_key
func setProps(data ...string) string {
	for i, s := range data {
		if s == "not_null" {
			data[i] = "not null"
		}
	}

	var work string
	work += " `gorm:"
	work += "\""
	work += strings.Join(data, ";")
	work += "\""
	work += "`"
	fmt.Println(work)
	return work
}
