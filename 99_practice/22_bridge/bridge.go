package p_bridge

import (
	"fmt"
)

/*
概念的な例
ここに 2 種類のコンピューターがあるとします。 Mac と Windows です。
そして 2 種類のプリンターがあるとします。 Epson と HP です。
両方のコンピューターとプリンターは、 任意の組み合わせで動作する必要があります。
クライアントは、 プリンターをコンピューターに接続する詳細について気を使いたくありません。

新しいプリンターを導入するとして、 コードが指数関数的に増大するのは避けたいものです。
2 × 2 の組み合わせのために 4 個の構造体を作成する代わりに、 二つの階層を作成します：

抽象化階層： コンピュータに対応します
実装階層： プリンターに対応します

この二つの階層は、 ブリッジを介して通信します。
ここでは、 抽象化 （コンピューター） は、 実装 （プリンター） への参照を保持します。
抽象化階層と実装階層は、 互いに悪影響を与えることなく、 独立して開発できます。
*/

/*
computer.go: 抽象化
*/
type Computer interface {
	Print()
	SetPrinter(Printer)
}

/*
mac.go: 特化した抽象化
*/
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

/*
windows.go: 特化した抽象化
*/
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

/*
printer.go: 実装
*/
type Printer interface {
	PrintFile()
}

/*
epson.go: 具象実装
*/
type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

/*
hp.go: 具象実装
*/
type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

/*
main.go: クライアント・コード
*/
func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
