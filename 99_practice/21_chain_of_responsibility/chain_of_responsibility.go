package p_chain_of_responsibility

import "fmt"

/*
概念的な例
Chain of Responsibility パターンを、 病院アプリを例に見ていきましょう
病院には以下のような複数の場所や役割があります：

- 受付
- 医者
- 薬局
- 支払い
病人が到着すると、まず受付に行き、 医者に会い、 薬局に行き、そして支払い窓口などへ行きます。
患者は、部局の連鎖の中を回され、各部局はその機能を終えると患者を次の部局へ送ります。

このパターンは、 同一リクエストを処理する複数の候補がある時、適用可能です。
リクエストを処理できる複数のオブジェクトがあり、 クライアントに受け手の選択をさせたくない場合もこのパターンが役に立ちます。
もう一つの役に立つケースとしては、 クライアントを受け手から分離したい場合です。
クライアントは、 連鎖の最初の要素だけ知っておく必要があります。
*/

/*
department.go: ハンドラー・インターフェース
*/
type Department interface {
	execute(*Patient)
	setNext(Department)
}

/*
reception.go: 具象ハンドラー
*/
type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

/*
doctor.go: 具象ハンドラー
*/
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

/*
medical.go: 具象ハンドラー
*/
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

/*
cashier.go: 具象ハンドラー
*/
type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

/*
patient.go
*/
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

/*
main.go: クライアント・コード
*/
func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
