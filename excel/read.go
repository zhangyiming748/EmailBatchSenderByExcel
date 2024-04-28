package excel

import (
	"github.com/tealeg/xlsx"
)

type EmailInfo struct {
	Email     string // 收件人地址
	Subject   string // 标题
	Body      string // 正文
	AttachURL string // 附件
}

func ReadExcelSave(filename string) (emailInfos []*EmailInfo, err error) {
	//var emailInfos []*EmailInfo

	// 打开Excel文件
	file, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	// 获取第一个工作表
	sheet := file.Sheets[0]

	// 遍历每一行数据
	for _, row := range sheet.Rows {
		e := new(EmailInfo)
		e.Email = row.Cells[0].String()     // email地址
		e.Subject = row.Cells[1].String()   // 标题
		e.Body = row.Cells[2].String()      // 正文
		e.AttachURL = row.Cells[3].String() // 附件地址

		emailInfos = append(emailInfos, e)
	}

	return emailInfos, nil
}
