
package types
type Money int

type Category string
type Status string
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

    

type Payment struct {
	ID int // 'card'
	Amount Money // номер вида '5058 xxxx xxxx 8888'
	Category Category // баланс в дирамах
    Status Status
}