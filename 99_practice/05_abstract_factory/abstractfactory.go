package p_abstractfactory

import "fmt"

/*
概念的な例
靴 1 足とシャツ 1 枚の二つの異なった製品の組み合わせであるスポーツ・キットを買う必要があるとします。
あなたは、 同じブランドのスポーツ・キットを購入したいと思っています。

もしこれをコードに転換したければ、 Abstract Factory が、 常に互いにマッチするプロダクトの組を作成するのに役立ちます。
*/

/*
Abstract Factory インターフェース
*/
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

type Brand int

const (
	adidas Brand = iota
	nike
)

func GetSportsFactory(brand Brand) (ISportsFactory, error) {
	if brand == adidas {
		return &Adidas{}, nil
	} else if brand == nike {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong branc type passed\n")
}

/*
具象ファクトリー
*/
type Adidas struct{}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

type Nike struct{}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}

/*
抽象プロダクト
*/
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getSize() int {
	return s.size
}

/*
具象プロダクト
*/
type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}

/*
抽象プロダクト
*/
type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}

/*
具象プロダクト
*/
type AdidasShirt struct {
	Shirt
}

type NikeShirt struct {
	Shirt
}

/*
クライアントコード
*/
func main() {
	adidasFactory, _ := GetSportsFactory(adidas)
	nikeFactory, _ := GetSportsFactory(nike)

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
