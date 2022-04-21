package objek

type Person struct {
	Name  string
	saldo int64
}

func Constructor(name string, saldo int64) Person {
	return Person{
		Name:  name,
		saldo: saldo,
	}
}

func (p Person) KasihTauSaldo() int64 {
	return p.saldo
}
