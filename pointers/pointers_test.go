package pointers

import (
	"testing"
)

func TestPointe(t *testing.T){
	

	


	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}	
		wallet.Deposit(Bitcoin(20))
		assetBalance(t, wallet, Bitcoin(20))


	})
	t.Run("withdrawl", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}	
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assetBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
	startingBalance := Bitcoin(20)
	wallet := Wallet{startingBalance}
	err := wallet.Withdraw(Bitcoin(100))
	assertError(t, err, ErrorMessage)
	assetBalance(t, wallet, startingBalance)
})	

}
func assetBalance (t testing.TB, wallet Wallet, want Bitcoin)  {
		t.Helper()	

		got := wallet.Balance()
		if got  != want {
		t.Errorf("got %s want %s", got, want)
		}	
	}

func assertError (t testing.TB, got error, want string)  {
		t.Helper()	
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("want %q got %q", want, got.Error() )
		}
}


func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}