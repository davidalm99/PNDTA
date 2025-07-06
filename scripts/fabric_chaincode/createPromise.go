package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Agreement struct
type Agreement struct {
	Buyer            string  `json:"buyer"`            //Who is getting access to the data
	Buyer_addr       string  `json:"eth_buyer_addr"`   //Ethereum wallet address of buyer
	Seller           string  `json:"seller"`           //Who is giving access to the data
	Seller_addr      string  `json:"eth_seller_addr"`  //Ethereum wallet address of seller
	ID               int     `json:"id"`               //Agreement's ID
	BatchID          string  `json:"batchID "`         //Batch ID of the data accorded
	IsActive         bool    `json:"isActive"`         //True when agreed by both parties
	Amount           int     `json:"amount"`           //Units that are accorded to be sold
	Price            int     `json:"price"`            //Price for each asset sold
	Payment          float64 `json:"payment"`          //How much will seller receive for each asset claimed
	Percentage       string  `json:"percentage"`       //Percentage of the 'Price' that will result in payment
	Percentage_Bonus string  `json:"percentage_bonus"` //Percentage of the 'Price' that will result in payment if Total_Devices > Amount
	AssetType        string  `json:"assetType"`        //What type of asset will be designed with this data
	TotalDevices     int     `json:"TotalDevices"`     //How many devices were designed with this data
}

// PromissoryNote contract - this contract will include methods to change 'Agreement' objcts in the ledger
type PromissoryNote struct {
	contractapi.Contract
}

func (s *PromissoryNote) InitLedger(ctx contractapi.TransactionContextInterface) error {
	counter := 0
	counterAsBytes := []byte(strconv.Itoa(counter)) //counter to string and  then converting to byte slice format
	ctx.GetStub().PutState("agreementCounter", counterAsBytes)
	return nil
}

func (s *PromissoryNote) CreatePromissoryNote(ctx contractapi.TransactionContextInterface, buyer string, b_addr string, seller string, s_addr string, batchID string, amount string, price int, assetType string, percent string, percent_bonus string) (int, error) {
	counterAsBytes, _ := ctx.GetStub().GetState("agreementCounter")
	counter, _ := strconv.Atoi(string(counterAsBytes))

	counter++

	_amount, err := ParseStringToInt(amount)

	if err != nil {
		return -1, err
	}

	_percent, err := ParseStringToInt(percent)

	if err != nil {
		return -1, err
	}

	_percent_bonus, err := ParseStringToInt(percent_bonus)

	if err != nil {
		return -1, err
	}

	total := 0

	_price := calculatePrice(_percent, _percent_bonus, _amount, total, price)

	agreement := Agreement{
		Buyer:            buyer,
		Buyer_addr:       b_addr,
		Seller:           seller,
		Seller_addr:      s_addr,
		ID:               counter,
		BatchID:          batchID,
		IsActive:         true,
		Amount:           _amount,
		Price:            price,
		Payment:          _price,
		Percentage:       percent,
		Percentage_Bonus: percent_bonus,
		AssetType:        assetType,
		TotalDevices:     total,
	}

	agreementAsBytes, err := json.Marshal(agreement)

	if err != nil {
		return -1, err
	}

	ctx.GetStub().PutState("Agreement"+strconv.Itoa(counter), agreementAsBytes) //Store in ledger as key value; Key: Agreement'x', Value: agreementAsBytes

	counterAsBytes = []byte(strconv.Itoa(counter))
	ctx.GetStub().PutState("agreementCounter", counterAsBytes)

	return counter, nil
}

func calculatePrice(percentage int, percentageBonus int, amount int, total int, pricePerUnit int) float64 {
	var price float64
	if total > amount {
		price = float64(pricePerUnit) * (float64(percentageBonus) / 100.0)
	} else {
		price = float64(pricePerUnit) * (float64(percentage) / 100.0)
	}
	return price
}

func ParseStringToInt(s string) (int, error) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("ParseStringToInt: %v is not a valid integer", s)
	}
	return value, nil
}

func (s *PromissoryNote) GetAllAgreements(ctx contractapi.TransactionContextInterface) ([]*Agreement, error) {
	startKey := "Agreement0" //Lexicographically smaller than any Agreement_ value (ASCII table)
	endKey := "AgreementZ"   //Lexicographically bigger than any Agreement_ value (ASCII table)

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var agreements []*Agreement
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var agreement Agreement
		err = json.Unmarshal(queryResponse.Value, &agreement)
		if err != nil {
			return nil, err
		}
		agreements = append(agreements, &agreement)
	}

	return agreements, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&PromissoryNote{})
	if err != nil {
		fmt.Printf("Error create product chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting product chaincode: %s", err.Error())
	}
}
