package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/Canto-Network/Canto/v6/x/govshuttle/client/cli"
)

var (
	LendingMarketProposalHandler = govclient.NewProposalHandler(cli.NewLendingMarketProposalCmd)
	TreasuryProposalHandler      = govclient.NewProposalHandler(cli.NewTreasuryProposalCmd)
)
