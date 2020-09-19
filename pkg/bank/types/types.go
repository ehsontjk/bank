
package types
type Money int

type Category string

	

    

type Payment struct {
	ID int // 'card'
	Amount Money // номер вида '5058 xxxx xxxx 8888'
	Category Category // баланс в дирамах

}