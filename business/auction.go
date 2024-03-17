package business

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

type AuctionStatus struct {
	Ended      bool  `json:"ended"`
	HighestBid int64 `json:"highestBid"`
}

type EthConnection struct {
	client *ethclient.Client
}

func NewBlockchainConnection() (*EthConnection, error) {

	instanceUrl, found := os.LookupEnv("INSTANCE_URL")
	if !found {
		return nil, errors.New("Instance url is needed to connect to an ethereum node")
	}

	client, err := ethclient.Dial(instanceUrl)
	if err != nil {
		return nil, err
	}

	return &EthConnection{
		client: client,
	}, nil
}

func (ec *EthConnection) GetAuctionStatus() (AuctionStatus, error) {

	actionStatus := AuctionStatus{}

	contractAddr, found := os.LookupEnv("CONTRACT_DEPLOYMENT_ADDR")
	if !found {
		return AuctionStatus{}, errors.New("Instance url is needed to connect to an ethereum node")
	}

	addr := common.HexToAddress(contractAddr)

	acn, err := NewSimpleAuction(addr, ec.client)
	if err != nil {
		return AuctionStatus{}, err
	}

	endTime, err := acn.AuctionEndTime(nil)
	if err != nil {
		return AuctionStatus{}, err
	}

	header, err := ec.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return AuctionStatus{}, err
	}

	actionStatus.Ended = header.Time > endTime.Uint64()

	highestBid, err := acn.HighestBid(nil)
	if err != nil {
		return AuctionStatus{}, err
	}

	actionStatus.HighestBid = highestBid.Int64()

	return actionStatus, nil
}
