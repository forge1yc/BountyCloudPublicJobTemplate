package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"time"
)

// 进行消息消费，以及读取，只是模拟程序，并不真实的消费，不过可以带回日志，我申请一个测试的邮箱账号，来接收这里发送的信息
// v2 只是测试基本的运行功能
// v3 测试了计费功能
// v4 与v2相比，只是将日志改为普通的控制台输出，而不是文件
// v5 与v3相比，更改了输出形式，使用go自带的输出
func main() {
	// 简单设置 log 参数
	//log.SetFlags(log.Lshortfile | log.LstdFlags)
	//log.SetOutput(os.Stdout)
	fmt.Sprintf("%+v\n","来自 boutycloud email")

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "bountycloud@163.com"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{"bountyemail@163.com"}

	// 设置主题
	em.Subject = "来自BountyCloud的测试邮件"

	fmt.Sprintf("sending email to %s", em.To)

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("hello , 尊敬的BountyCloud测试用户您好，这封邮件来自某个Porvider机器运行 bounty_test_email 镜像的自动发送邮件，收到该邮件代表链路畅通， 10秒钟之后请返回 https://wwww.bountycloud.com 控制台查看job结果")

	//设置服务器相关的配置
	fmt.Println("send successfully ... ")
	fmt.Println("wait 10s to end, to test billing function, send a email every 10 minutes, please check email")

	err := em.Send("smtp.163.com:25", smtp.PlainAuth("", "bountycloud@163.com", "TPYQKLWCCIHDYCSP", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stderr,"this is error test\n")
	time.Sleep(time.Second * 10)

}
