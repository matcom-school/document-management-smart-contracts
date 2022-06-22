package chaincodes

import (
	"encoding/json"
	"fmt"
	"strconv"

	ca "github.com/hyperledger/fabric-contract-api-go/contractapi"
	cc "github.com/matcom-school/document-management-smart-contracts/chaincodes"
)

// File describes basic details of what makes up a simple asset
type Files struct {
	CreatedAt string   `json:"createdAt"`
	ID        uint32   `json:"id"`
	Name      string   `json:"name"`
	Owner     string   `json:"owner"`
	Customers []string `json:"customers"`
	Url       string   `json:"url"`
	Size      int      `json:"size"`
	Type      string   `json:"type"`
}

const idCount = 0

func getIdByUrl(url string) uint32 {
	str := strconv.Itoa(idCount)
	return cc.StringHash(url + str)
}

// InitLedger adds a base set of assets to the ledger
func (s *cc.SmartContract) InitLedger(ctx ca.TransactionContextInterface) error {
	assets := []Files{
		{
			Size: 5, Owner: "tomoko@gmail.com",
			CreatedAt: "2022-03-10", Customers: make([]string, 0),
			Type: "mp3", Name: "Im not", Url: "http//:localhost:9000"},
	}

	for _, asset := range assets {
		asset.ID = getIdByUrl((asset.Url))
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *cc.SmartContract) CreateFile(ctx ca.TransactionContextInterface,
	url string, name, string, createdAt string, size int, owner string, type_ string) error {

	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
