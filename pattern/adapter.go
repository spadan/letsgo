// 适配器
package pattern

import "fmt"

// 一种客户端，可通过闪电接口与其他设备通信
type client struct{}

// 带闪电接口的设备
type computer interface {
	insertLightningConnector()
}

func (c *client) insertLightningConnector2Computer(com computer) {
	com.insertLightningConnector()
}

// mac电脑，自带闪电接口
type mac struct{}

func (m *mac) insertLightningConnector() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// win电脑，只有usb接口
type windows struct{}

func (w *windows) insertUsbConnector() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// 怎么让客户端通过闪电接口连接windows电脑
// 转接设备，它可以连接闪电接口，同时另一端连usb
type adapter struct {
	win *windows
}

func (a *adapter) insertLightningConnector() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	a.win.insertUsbConnector()
}
