package p_adapter

import "fmt"

/*
概念的な例
オブジェクトの何らかの機能 （ライトニング・ポート） の存在を期待するクライアント・コードがあります。
しかし、 ここで adaptee と呼ばれるオブジェクト （Windows ラップトップ） は、
同じ機能を異なるインターフェース （USB ポート） を通して提供します。

ここで Adapter パターンが登場します。 以下を行う adapter という struct 型を作成します：

- クライアントが期待しているのと同じインターフェース （ライトニング・ポート） に従う。

- クライアントからのリクエストを adaptee が期待する形式に翻訳する。 adapter は、 ライトニング・コネクターを受け入れ、 その信号を USB 形式に変換し、 Windows ラップトップの USB ポートに渡す。
*/
/*
クライアントコード
*/
type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

/*
クライアントインターフェース
*/
type Computer interface {
	InsertIntoLightningPort()
}

/*
サービス
*/
type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

/*
不明なサービス
*/
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

/*
アダプター
*/
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

/*
main
*/
func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
