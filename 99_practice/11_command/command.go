package p_command

import "fmt"

/*
概念的な例
テレビを例にとってみてみましょう。 テレビの電源を入れるには以下のどちらかを押します

- リモコンのオンのボタン
- 実際のテレビのオンのボタン

「オンにする」 コマンド・オブジェクトの実装はテレビを受け手として開始します
このコマンドの実行メソッドが呼ばれるとそれは次に TV.on 関数を呼び出します
次に、 インボーカーを定義します
実際のところインボーカーは二つあります
リモコンとテレビ自身です。 両方とも 「オンにする」 コマンド・オブジェクトを埋め込んでいます。

同じリクエストを複数のインボーカーでラップしていることに注目してください。他のコマンドでも同様です。
別々のコマンド・オブジェクトを作成する利点は、UI ロジックをビジネス・ロジックから切り離すことです。
インボーカーごとに異なるハンドラーを開発する必要はありません。コマンド・オブジェクトには実行に必要なすべての情報が含まれています。そのため、 遅延実行にも使用できます。
*/

/*
button.go: インボーカー
*/
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

/*
command.go: コマンド・インターフェース
*/
type Command interface {
	execute()
}

/*
onCommand.go: 具象コマンド
*/
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

/*
offCommand.go: 具象コマンド
*/
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

/*
device.go: 受け手インターフェース
*/
type Device interface {
	on()
	off()
}

/*
tv.go: 具象受け手
*/
type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

/*
main.go: クライアント・コード
*/
func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
