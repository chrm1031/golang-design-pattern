package p_flyweight

import "fmt"

/*
概念的な例

*/

/*
dressFactory.go: フライウェイト・ファクトリー
*/
var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

type DressFactory struct {
	dressMap map[string]Dress
}

const (
	TerroristDressType         = "tDress"
	CounterTerrroristDressType = "ctDress"
)

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

/*
dress.go: フライウェイト・インターフェース
*/
type Dress interface {
	getColor() string
}

/*
terroristDress.go: 具象フライウェイト・オブジェクト
*/
type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() Dress {
	return &TerroristDress{color: "red"}
}

/*
counterTerroristDress.go: 具象フライウェイト・オブジェクト
*/
type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

/*
player.go: コンテキスト
*/
type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType string, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		dress:      dress,
		playerType: playerType,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

/*
game.go: クライアント・コード
*/
type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", TerroristDressType)
	c.terrorists = append(c.terrorists, player)
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

/*
main.go: クライアント・コード
*/
func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
