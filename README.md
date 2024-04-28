# 假设表格为如下形式

|收件人地址|标题|正文|附件|
|:---:|:---:|:---:|:---:|
|a@gmail.com|title|test|1.jpg|
|b@gmail.com|title|test|1.jpg|

# 成功后输出

```bash
Email Info:&{Email:zhangyiming748@gmail.com FileName:1 AttachURL:58a418be85d47.jpg}
2024/04/28 21:27:18 INFO 发送邮件
Email Info:&{Email:zhangyiming748@outlook.com FileName:2 AttachURL:58a418be85d47.jpg}
2024/04/28 21:27:20 INFO 发送邮件
Email Info:&{Email:zhangyiming748@protonmail.com FileName:3 AttachURL:58a418be85d47.jpg}
2024/04/28 21:27:22 INFO 发送邮件
Email Info:&{Email:578779391@qq.com FileName:4 AttachURL:58a418be85d47.jpg}
2024/04/28 21:27:23 INFO 发送邮件
```