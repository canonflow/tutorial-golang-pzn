package _11_sync_rwmutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* ===== INTRO =====
- RWMutex (Read Write Mutex)
- Kadang ada kasus dimana kita ingin melakukan LOCKING tidak hanya pada proses PENGUBAHAN DATA, tapi juga MEMBACA DATA
- Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan REBUTAN antara proses MEMBACA dan MENGUBAH
- Di Golang telah disediakan stuct sync.RWMutex untuk menangani hal ini, dimana Mutex jenis ini memiliki 2 LOCK:
  - LOCK untuk Read
  - LOCK untuk Write
*/

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance(amount int) int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}
func TestSyncRWMutex(t *testing.T) {

	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance(1))
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance:", account.Balance) // 10000
	/*
		...
		9995
		9996
		9997
		9930
		9998
		9999
		10000
		Final Balance: 10000
		--- PASS: TestSyncRWMutex (5.00s)
		PASS
	*/
}
