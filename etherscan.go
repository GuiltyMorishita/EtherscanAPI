package etherscan

// Etherscan - the Etherscan API toolkit
type Etherscan struct {
	apiKey string
}

// NewEtherscan - create an api object with the correct key
func NewEtherscan(apiKey string) *Etherscan {
	return &Etherscan{apiKey}
}
