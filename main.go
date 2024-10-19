package main

import (
	"pdf2png-node/depot"
	"pdf2png-node/mqcode"
)

func main() {
	depot.MkdirTemp()
	// Cleanup temporary working directory after finished
	defer depot.CleanTemp()

	// Configure unoconvert options
	//unoAddr := c.String("unoserver-addr")
	//host, port, _ := net.SplitHostPort(unoAddr)
	//unoconvert.SetInterface(host)
	//unoconvert.SetPort(port)
	//unoconvert.SetExecutable(c.String("unoconvert-bin"))
	//unoconvert.SetContextTimeout(c.Duration("unoconvert-timeout"))
	var producer = mqcode.CreateProducer()
	// 启动rocketmq消费者开始等待需要转换的文件进入，一次最大允许一个
	var consumers = mqcode.CreateConsumer(producer)
	consumers.StartConsumer()
}
