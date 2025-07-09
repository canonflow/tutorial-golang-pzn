package _12_deadlock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* ===== INTRO =====
- Hati - hati saat membuat aplikasi yang PARALLEL atau CONCCURENT, masalah yang sering dihadapi adalah DEADLOCK
- Deadlock adalah keadaan dimana sebuah proses goroutine SALING MENUNGGU LOCK sehingga TIDAK ADA satupun goroutine
  yang BISA JALAN
*/

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Charge(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User 1:", user1.Name)
	user1.Charge(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2:", user2.Name)
	user2.Charge(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := &UserBalance{
		Name:    "Nathan",
		Balance: 10000,
	}

	user2 := &UserBalance{
		Name:    "Garyza",
		Balance: 20000,
	}

	go Transfer(user1, user2, 1000)
	go Transfer(user2, user1, 2000)

	time.Sleep(5 * time.Second)
	fmt.Println("User:", user1.Name, "Balance:", user1.Balance)
	fmt.Println("User:", user2.Name, "Balance:", user2.Balance)
	/*
		=== RUN   TestDeadlock
		Lock User 1: Garyza
		Lock User 1: Nathan
		User: Nathan Balance: 9000
		User: Garyza Balance: 18000
		--- PASS: TestDeadlock (5.00s)
		PASS
	*/
	/*
		Data yang ditransfer sebelumnya hilang, karena blm sempet kita transfer karena
		2 2 nya saling menunggu
		Bisa dilihat bahwa saldo akhir Nathan hanya 9000 (hasil pemotongan 1000) dan
		Garyza hanya 18000 (hasil pemotongan 2000)
	*/
}
