package p_builder

import "fmt"

/*
概念的な例
異なる種類の家（iglooとnormalHouse）が、iglooBuilderとnormalBuilderによって構築されている
どちらの種類の家も、同じ構築ステップを使用する
ディレクター構造体は、構築過程を組織化するのに役に立つ
*/

/*
Builderインターフェース
*/
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuiler()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

/*
normalBuilder: 具象ビルダー
*/
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuiler() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

/*
iglooBuilder: 具象ビルダー
*/
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 2
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

/*
house: 製品
*/
type House struct {
	windowType string
	doorType   string
	floor      int
}

/*
director: ディレクター
*/
type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

/*
main: クライアントコード
*/
func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
