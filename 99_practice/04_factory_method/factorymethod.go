package p_factorymethod

import "fmt"

/*
概念的な例
この例では、 ファクトリー構造体を使用して、 多種の兵器を製造します。
*/

/*
iGun プロダクトインターフェース
銃が持つべき全メソッド（機能）を定義する
*/
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

/*
具象プロダクト
*/
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

type gunType string

const (
	_ak47   = "ak47"
	_musket = "musket"
)

func getGun(t gunType) (IGun, error) {
	if t == _ak47 {
		return newAk47(), nil
	}
	if t == _musket {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

/*
クライアントコード
*/
func main() {
	ak47, _ := getGun(_ak47)
	musket, _ := getGun(_musket)

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
