package main

import (
	"fmt"
	"readandsend/email"
	"readandsend/excel"
	"time"
)

func main() {
	filename := "example.xlsx" // 替换为实际的Excel文件路径
	emailInfos, err := excel.ReadExcelSave(filename)
	if err != nil {
		fmt.Println("Error reading Excel file:", err)
		return
	}
	for _, emailInfo := range emailInfos {
		fmt.Printf("Email Info:%+v\n", emailInfo)
		email.Send(email.Username, emailInfo.Email, emailInfo.Subject, emailInfo.Body, emailInfo.AttachURL)
		time.Sleep(1 * time.Second)
	}

}
