package repository

import (
	"testing"
	"time"

	"github.com/energywork/pseudo-paysystem/balance/internal/repository"
	"github.com/energywork/pseudo-paysystem/balance/internal/repository/mocks"
)

func TestBalanceCreate(t *testing.T) {
	/*cfg := config.New("balance_repo_test")
	if cfg == nil {
		t.Fatal("Config isn't initialized")
	}
	set := setup.GetSetupForTest(t, cfg, false)*/
	// -----------------------------------------------------------------------------------------------------------------

	repo := mocks.NewRepository(t)

	b := &repository.Balance{
		LimitID:        1,
		WalletID:       "123",
		Value:          0,
		Hold:           0,
		Identification: 1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	repo.On("Create", b).Return(nil)

	err := repo.Create(b)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Balance created!")
}
