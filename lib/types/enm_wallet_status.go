package types

type WalletStatus int

const (
	WalletActive WalletStatus = iota + 1
	WalletExpired
)

func (ws WalletStatus) String() string {
	switch ws {
	case WalletActive:
		return "active"
	case WalletExpired:
		return "expired"
	default:
		return "undefined"
	}
}

func (ws WalletStatus) IsValid() bool {
	switch ws {
	case WalletActive, WalletExpired:
		return true
	}
	return false
}
