package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	BatchID          string  `json:" batchID "`        //Batch ID of the data accorded
	IsActive         bool    `json:"isActive"`         //True when agreed by both parties
	Amount           int     `json:"amount"`           //Units that are accorded to be sold
	Price            int     `json:"price"`            //Price for each asset sold
	Payment          float64 `json:"payment"`          //How much will seller receive for each asset claimed
	Percentage       string  `json:"percentage"`       //Percentage of the 'Price' that will result in payment
	Percentage_Bonus string  `json:"percentage_bonus"` //Percentage of the 'Price' that will result in payment if Total_Devices > Amount
	AssetType        string  `json:"assetType"`        //What type of asset will be designed with this data
	TotalDevices     int     `json:"TotalDevices"`     //How many devices were designed with this data
}

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"color"`
	Size           int    `json:"size"`
	Owner          string `json:"owner"`
	AppraisedValue int    `json:"appraisedValue"`
}

type AriesAgent struct {
	AgentID     string `json:"agentID"`
	AgentType   string `json:"agentType"`
	AgentStatus string `json:"status"`
}

//agentType = [CONSORTIUM, OEM_GW, OEM_SD, CONSUMER, AI_SERVICE_PROVIDER, DATA_PURCHASER]
//agentStatus = [IN_GOOD_STANDING]

type Device struct {
	DeviceID       string `json:" deviceID "`
	ControllerID   string `json:" controllerID "`
	DeviceModelID  string `json:" DeviceModelID "`
	DeviceType     string `json:" DeviceType "`
	DeviceStatus   string `json:" status "`
	DTID           string `json:" DTID "`
	SellInvitation string `json:" sell_invitation "`
}

//deviceStatus = [AVAILABLE, CLAIMED, TWINNED, IN_TRANSIT, DECOMMISSIONED]

type DataBatch struct {
	BatchID      string `json:" batchID "`
	BatchURL     string `json:" batchURL "`
	Hash         string `json:" hash "`
	ControllerID string `json:" ControllerID "` // ?? Filipe doubt
	timestamp    string `json:" ControllerID "`
}

type DeviceModel struct {
	DeviceModelID string   `json:"deviceModelID"`
	Description   string   `json:"description"`
	Features      []string `json:"features"`
	Images        []string `json:"images"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	/*assets := []Asset{
	  {ID: "asset1", Color: "blue", Size: 5, Owner: "Tomoko", AppraisedValue: 300},
	  {ID: "asset2", Color: "red", Size: 5, Owner: "Brad", AppraisedValue: 400},
	  {ID: "asset3", Color: "green", Size: 10, Owner: "Jin Soo", AppraisedValue: 500},
	  {ID: "asset4", Color: "yellow", Size: 10, Owner: "Max", AppraisedValue: 600},
	  {ID: "asset5", Color: "black", Size: 15, Owner: "Adriana", AppraisedValue: 700},
	  {ID: "asset6", Color: "white", Size: 15, Owner: "Michel", AppraisedValue: 800},
	}*/

	ariesAgents := []AriesAgent{
		{AgentID: "agent1",
			AgentType:   "CONSORTIUM",
			AgentStatus: "IN_GOOD_STANDING"},
	}

	deviceModels := []DeviceModel{
		{DeviceModelID: "devicemodel1",
			Description: "iWatch Device Model",
			Features:    []string{"feature1"},
			Images:      []string{"image1"}},
	}

	devices := []Device{
		{DeviceID: "device1",
			ControllerID:   "controller1",
			DeviceModelID:  "devicemodel1",
			DeviceType:     "deviceType",
			DeviceStatus:   "deviceStatus",
			DTID:           "digitalTwin1",
			SellInvitation: "url1"},
	}

	dataBatchs := []DataBatch{
		{BatchID: "batch1",
			BatchURL: "batch url 1",
			Hash:     "hash1"},
	}

	for _, ariesagent := range ariesAgents {
		assetJSON, err := json.Marshal(ariesagent)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(ariesagent.AgentID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, devicemodel := range deviceModels {
		assetJSON, err := json.Marshal(devicemodel)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(devicemodel.DeviceModelID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, device := range devices {
		assetJSON, err := json.Marshal(device)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(device.DeviceID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	for _, databatch := range dataBatchs {
		assetJSON, err := json.Marshal(databatch)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(databatch.BatchID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}
	/*
	  for _, asset := range assets {
	    assetJSON, err := json.Marshal(asset)
	    if err != nil { return err }
	    err = ctx.GetStub().PutState(asset.ID, assetJSON)
	    if err != nil { return fmt.Errorf("failed to put to world state. %v", err) }
	  }*/

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
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

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
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

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Owner = newOwner
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}

		if asset.ID == "" {
			continue
		} else {
			assets = append(assets, &asset)
		}

	}

	return assets, nil
}

// ARIES AGENT FUNCTIONS
func (s *SmartContract) AgentExists(ctx contractapi.TransactionContextInterface, agentid string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(agentid)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return assetJSON != nil, nil
}

func (s *SmartContract) AgentRegistration(ctx contractapi.TransactionContextInterface, agentid string, agentype string, agentstatus string) error {
	exists, err := s.AgentExists(ctx, agentid)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the aries agent %s already exists", agentid)
	}

	ariesagent := AriesAgent{
		AgentID:     agentid,
		AgentType:   agentype,
		AgentStatus: agentstatus,
	}
	assetJSON, err := json.Marshal(ariesagent)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(agentid, assetJSON)
}

func (s *SmartContract) GetAllAriesAgents(ctx contractapi.TransactionContextInterface) ([]*AriesAgent, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var ariesagents []*AriesAgent
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var ariesagent AriesAgent
		err = json.Unmarshal(queryResponse.Value, &ariesagent)
		if err != nil {
			return nil, err
		}

		if ariesagent.AgentID == "" {
			continue
		} else {
			ariesagents = append(ariesagents, &ariesagent)
		}
	}

	return ariesagents, nil
}

// DEVICE MODEL FUNCTIONS
func (s *SmartContract) GetAllDeviceModels(ctx contractapi.TransactionContextInterface) ([]*DeviceModel, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var devicemodels []*DeviceModel
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var devicemodel DeviceModel
		err = json.Unmarshal(queryResponse.Value, &devicemodel)
		if err != nil {
			return nil, err
		}

		if devicemodel.DeviceModelID == "" {
			continue
		} else {
			devicemodels = append(devicemodels, &devicemodel)
		}
	}

	return devicemodels, nil
}

// DEVICE FUNCTIONS
func (s *SmartContract) DeviceExists(ctx contractapi.TransactionContextInterface, deviceid string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(deviceid)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return assetJSON != nil, nil
}

func (s *SmartContract) DeviceGenesisRegistration(ctx contractapi.TransactionContextInterface, deviceid string,
	controllerid string,
	devicemodelid string,
	devicetype string,
	devicestatus string,
	dtid string,
	sellinvitation string) error {
	exists, err := s.DeviceExists(ctx, deviceid)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the device %s already exists", deviceid)
	}

	device := Device{
		DeviceID:       deviceid,
		ControllerID:   controllerid,
		DeviceModelID:  devicemodelid,
		DeviceType:     devicetype,
		DeviceStatus:   devicestatus,
		DTID:           dtid,
		SellInvitation: sellinvitation,
	}
	assetJSON, err := json.Marshal(device)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(deviceid, assetJSON)
}

func (s *SmartContract) GetAllDevices(ctx contractapi.TransactionContextInterface) ([]*Device, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var devices []*Device
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var device Device
		err = json.Unmarshal(queryResponse.Value, &device)
		if err != nil {
			return nil, err
		}
		if device.DeviceID == "" {
			continue
		} else {
			devices = append(devices, &device)
		}
	}

	return devices, nil
}

func (s *SmartContract) ReadDevice(ctx contractapi.TransactionContextInterface, deviceid string) (*Device, error) {
	assetJSON, err := ctx.GetStub().GetState(deviceid)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the device %s does not exist", deviceid)
	}

	var device Device
	err = json.Unmarshal(assetJSON, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (s *SmartContract) BuyDevice(ctx contractapi.TransactionContextInterface, deviceid string, status string) error {
	device, err := s.ReadDevice(ctx, deviceid)
	if err != nil {
		return fmt.Errorf("failed to buy device: %v", err)
	}
	device.DeviceStatus = status
	assetJSON, err := json.Marshal(device)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(deviceid, assetJSON)
}

func (s *SmartContract) ClaimDevice(ctx contractapi.TransactionContextInterface, deviceid string,
	status string,
	controllerid string) float64 {
	device, _ := s.ReadDevice(ctx, deviceid)

	device.DeviceStatus = status
	device.ControllerID = controllerid
	assetJSON, _ := json.Marshal(device)

	ctx.GetStub().PutState(deviceid, assetJSON)

	payload_device := Device{
		DeviceID:       device.DeviceID,
		ControllerID:   device.ControllerID,
		DeviceModelID:  device.DeviceModelID,
		DeviceType:     device.DeviceType,
		DeviceStatus:   device.DeviceStatus,
		DTID:           device.DTID,
		SellInvitation: device.SellInvitation,
	}

	deviceAsBytes, _ := json.Marshal(payload_device)
	// if err != nil {
	// 	return err
	// }

	fmt.Printf("Payload: %s\n", deviceAsBytes)

	// //Emit an event
	// ctx.GetStub().SetEvent("ClaimedDevice", []byte(deviceAsBytes))
	// if err != nil {
	// 	return fmt.Errorf("failed to set event: %v", err)
	// }

	// After the device has been successfully claimed, we can check the promissory notes:
	agreementCounterAsBytes, _ := ctx.GetStub().GetState("agreementCounter")
	agreementCounter, _ := strconv.Atoi(string(agreementCounterAsBytes))

	result := 0.0
	for i := 1; i <= agreementCounter; i++ {
		agreementAsBytes, _ := ctx.GetStub().GetState("Agreement" + strconv.Itoa(i))
		var agreement Agreement
		json.Unmarshal(agreementAsBytes, &agreement)
		if agreement.AssetType == device.DeviceType {
			agreement.TotalDevices = agreement.TotalDevices + 1 // assuming totalDevices has been added to your Agreement struct

			percent := ParseStringToInt(agreement.Percentage)
			percent_bonus := ParseStringToInt(agreement.Percentage_Bonus)

			agreement.Payment = calculatePayment(percent,
				percent_bonus,
				agreement.Amount,
				agreement.TotalDevices,
				agreement.Price)

			agreementAsBytes, _ = json.Marshal(agreement)
			ctx.GetStub().PutState("Agreement"+strconv.Itoa(i), agreementAsBytes)
		}
	}

	//Emit an event
	ctx.GetStub().SetEvent("ClaimedDevice", []byte(deviceAsBytes))

	return result
}

func calculatePayment(percentage int, percentageBonus int, amount int, total int, pricePerUnit int) float64 {
	var payment float64
	if total > amount {
		payment = float64(pricePerUnit) * (float64(percentageBonus) / 100.0)
	} else {
		payment = float64(pricePerUnit) * (float64(percentage) / 100.0)
	}
	return payment
}

func ParseStringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return value
}

func (s *SmartContract) TwinDevice(ctx contractapi.TransactionContextInterface, deviceid string, status string, dtid string) error {
	device, err := s.ReadDevice(ctx, deviceid)
	if err != nil {
		return err
	}
	device.DeviceStatus = status
	device.DTID = dtid
	assetJSON, err := json.Marshal(device)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(deviceid, assetJSON)
}

func (s *SmartContract) UntwinDevice(ctx contractapi.TransactionContextInterface, deviceid string, status string, dtid string) error {
	device, err := s.ReadDevice(ctx, deviceid)
	if err != nil {
		return err
	}
	device.DeviceStatus = status
	device.DTID = dtid
	assetJSON, err := json.Marshal(device)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(deviceid, assetJSON)
}

func (s *SmartContract) SellDevice(ctx contractapi.TransactionContextInterface, deviceid string, status string) error {
	device, err := s.ReadDevice(ctx, deviceid)
	if err != nil {
		return err
	}
	device.DeviceStatus = status
	assetJSON, err := json.Marshal(device)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(deviceid, assetJSON)
}

//BATCH FUNCTIONS

func (s *SmartContract) BatchExists(ctx contractapi.TransactionContextInterface, batchid string, hash string) (bool, error) {
	assetJSON, err1 := ctx.GetStub().GetState(batchid)
	assetJSON, err2 := ctx.GetStub().GetState(hash)
	if err1 != nil || err2 != nil {
		return false, fmt.Errorf("failed to read from world state: %v and %v", err1, err2)
	}
	return assetJSON != nil, nil
}

// 1 batch = 1 hash
func (s *SmartContract) BatchRegistration(ctx contractapi.TransactionContextInterface, batchid string, batchurl string, hash string) error {
	exists, err := s.BatchExists(ctx, batchid, hash)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the device %s already exists", batchid)
	}

	batch := DataBatch{
		BatchID:  batchid,
		BatchURL: batchurl,
		Hash:     hash,
	}
	assetJSON, err := json.Marshal(batch)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(batchid, assetJSON)
}

func (s *SmartContract) GetAllDataBatchs(ctx contractapi.TransactionContextInterface) ([]*DataBatch, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var databatchs []*DataBatch
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var databatch DataBatch
		err = json.Unmarshal(queryResponse.Value, &databatch)
		if err != nil {
			return nil, err
		}

		if databatch.BatchID == "" {
			continue
		} else {
			databatchs = append(databatchs, &databatch)
		}
	}

	return databatchs, nil
}

func main() {
	assetChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}
	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
