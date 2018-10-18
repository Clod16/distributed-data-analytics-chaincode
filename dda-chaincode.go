package main

import (
	"bytes"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/rs/xid"
)

var logger = shim.NewLogger("dda-chaincode-log")

//var logger = shim.NewLogger("dcot-chaincode")

// DistributedDataAnalyticsChaincode implementation
type DistributedDataAnalyticsChaincode struct {
	testMode bool
}

func (t *DistributedDataAnalyticsChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	logger.Info("Chaincode Interface - Init()\n")
	logger.SetLevel(shim.LogDebug)
	_, args := stub.GetFunctionAndParameters()
	//var err error

	// Upgrade Mode 1: leave ledger state as it was
	if len(args) == 0 {
		//logger.Info("Args correctly!!!")
		return shim.Success(nil)
	}

	return shim.Success(nil)
}

func (t *DistributedDataAnalyticsChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	var creatorOrg, creatorCertIssuer string
	//var attrValue string
	var err error
	var isEnabled bool
	var callerRole string

	logger.Debug("Chaincode Interface - Invoke()\n")

	if !t.testMode {
		creatorOrg, creatorCertIssuer, err = getTxCreatorInfo(stub)
		if err != nil {
			logger.Error("Error extracting creator identity info: \n", err.Error())
			return shim.Error(err.Error())
		}
		logger.Info("DistributedDataAnalyticsChaincode Invoke by '', ''\n", creatorOrg, creatorCertIssuer)
		callerRole, _, err = getTxCreatorInfo(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

		isEnabled, _, err = isInvokerOperator(stub, callerRole)
		if err != nil {
			logger.Error("Error getting attribute info: \n", err.Error())
			return shim.Error(err.Error())
		}
	}

	function, args := stub.GetFunctionAndParameters()

	if function == "createAnalyticsInstances" {
		return t.createAnalyticsInstances(stub, isEnabled, args)
	} else if function == "updateAnalyticsInstances" {
		return t.updateAnalyticsInstances(stub, isEnabled, args)
	} else if function == "delateAnalyticsInstances" {
		return t.delateAnalyticsInstances(stub, isEnabled, args)
	} else if function == "getAnalyticsInstancesById" {
		return t.getAnalyticsInstancesById(stub, isEnabled, args)
	} else if function == "getAnalyticsInstancesByEgid" {
		return t.getAnalyticsInstancesByEgid(stub, isEnabled, args)
	} else if function == "getAnalyticsInstances" {
		return t.getAnalyticsInstances(stub, isEnabled)
	} else if function == "createDataSources" {
		return t.createDataSources(stub, isEnabled, args)
	} else if function == "deleteDataSources" {
		return t.deleteDataSources(stub, isEnabled, args)
	} else if function == "getDataSources" {
		return t.getDataSources(stub, isEnabled)
	} else if function == "getDataSourcesbyId" {
		return t.getDataSourcesbyId(stub, isEnabled, args)
	} else if function == "createEdgeGateways" {
		return t.createEdgeGateways(stub, isEnabled, args)
	} else if function == "updateEdgeGateways" {
		return t.updateEdgeGateways(stub, isEnabled, args)
	} else if function == "getEdgeGateways" {
		return t.getEdgeGateways(stub, isEnabled)
	} else if function == "getEdgeGatewaysByEgid" {
		return t.getEdgeGatewaysByEgid(stub, isEnabled, args)
	} else if function == "deleteEdgeGateways" {
		return t.deleteEdgeGateways(stub, isEnabled, args)
	}
	return shim.Error("Invalid invoke function name")
}


func (t *DistributedDataAnalyticsChaincode) createDataSources(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	



}
func (t *DistributedDataAnalyticsChaincode) deleteDataSources(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}
func (t *DistributedDataAnalyticsChaincode) getDataSources(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}
func (t *DistributedDataAnalyticsChaincode) getDataSources(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}
func (t *DistributedDataAnalyticsChaincode) getDataSourcesbyId(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}
func (t *DistributedDataAnalyticsChaincode) createEdgeGateways(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}
func (t *DistributedDataAnalyticsChaincode) updateEdgeGateways(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}
func (t *DistributedDataAnalyticsChaincode) getEdgeGateways(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}
func (t *DistributedDataAnalyticsChaincode) getEdgeGatewaysByEgid(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}
func (t *DistributedDataAnalyticsChaincode) deleteEdgeGateways(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	var err error
	if(args[0]){
		return shim.Error(err.Error())
	}else{
		return shim.Success(nil)
	}
}

func (t *DistributedDataAnalyticsChaincode) getAnalyticsInstances(shim.ChaincodeStubInterface, isEnabled bool) {

	analyticsArrayBytes, err := stub.GetStateByPartialCompositeKey("FE_Analytics_Instances" , DDAPartialKey)
	if err != nil{
		logger.error(" GetStateByPartialCompositeKey() ERROR:\n")
		return shim.Error(err.Error())
	}

	return shim.Success(analyticsArrayBytes)

	
}
func (t *DistributedDataAnalyticsChaincode) getAnalyticsInstancesByEgid(shim.ChaincodeStubInterface, isEnabled bool, args []string) {
	logger.info(" getAnalyticsInstancesByEgid()\n")

	if len(args) != 1{
		logger.error(" getAnalyticsInstancesByEgid() ERROR: wrong argument\n")
		return shim.Error("getAnalyticsInstancesByEgid() ERROR: wrong argument" )
	}
	DDAPartialKey := args[0]
	analyticsArrayBytes, err := stub.GetStateByPartialCompositeKey("FE_Analytics_Instances" , DDAPartialKey)
	if err != nil{
		logger.error(" GetStateByPartialCompositeKey() ERROR:\n")
		return shim.Error(err.Error())
	}

	return shim.Success(analyticsArrayBytes)

}


func (t *DistributedDataAnalyticsChaincode) getAnalyticsInstancesById(shim.ChaincodeStubInterface, isEnabled bool, args []string) {

	logger.info(" getAnalyticsInstancesById()\n")
	//var analyticsArrayString []string
	if len(args) != 1{
		logger.error(" getAnalyticsInstancesById() ERROR: wronh argument\n")
		return shim.Error("getAnalyticsInstancesById() ERROR: wrong argument" )
	}
	DDAPartialKey := args[0]
	analyticsArrayBytes, err := stub.GetStateByPartialCompositeKey("FE_Analytics_Instances" , DDAPartialKey)
	if err != nil{
		logger.error(" GetStateByPartialCompositeKey() ERROR:\n")
		return shim.Error(err.Error())
	}
l
	/* for( i=1, i<len(analyticsArrayBytes), i++){

		payloadByte := analyticsArrayBytes[i]
		payload := BytesToString(payloadByte)
		analytics := new(AnalitycsInstances)
		analytics.Id = args[0]
		analytics.Payload = payload
		analyticsJson,err := json.Marshal(analytics)
		if err != nil{
			return shim.Error("json.Marshal() ERROR: ", err.Error())
		}
		analyticsArrayString = append(analyticsArray, analyticsJson)		
	}
	logger.Info("Query Response:\n", analyticsArrayString) */
	return shim.Success(analyticsArrayBytes)

}
func (t *DistributedDataAnalyticsChaincode) delateAnalyticsInstances(shim.ChaincodeStubInterface, isEnabled bool, args []string) {

	logger.info(" delateAnalyticsInstances()\n")

	if len(args) != 2{
		return shim.Error("delateAnalyticsInstances() ERROR: wrong argument" )
	}

	DDAKey , err := getAnalyticsKey(stub, args[0], args[1])
	if err != nil{
		return shim.Error("CreateCompositeKey() ERROR: ", err.Error())
	}

	err := stub.DelState(DDAKey)
	if err != nil{
		return shim.Error("DelState() ERROR", err.Error()) 
	}
	stringEvent :=  "Analytics Instances deleted with key: " +DDAKey
	err := setDDAEvent(stub, stringEvent )
	if err != nil{
		return shim.Error(" setEvent() ERROR: ", err.Error())
	}
	return shim.Success(nil)

}


func (t *DistributedDataAnalyticsChaincode) updateAnalyticsInstancesstub(shim.ChaincodeStubInterface, isEnabled bool, args []string) {

	var analytics AnalitycsInstances
	logger.info(" updateAnalyticsInstancesstub()\n")

	if len(args) != 3{
		return shim.Error("updateAnalyticsInstancesstub() ERROR: wrong argument")
	}

	DDAKey, err := getAnalyticsKey(stub, args[0], args[2]){
		if err != nil{
			return shim.Error("CreateCompositeKey() ERROR: ", err.Error())
		}
	}

	DDABytes, err := stub.GetState(DDAKey)
	if err!= nil{
		return shim.Error("GetState() ERROR: ", err.Error())
	}
	
	err := json.Unmarshal([]byte(DDABytes), &analytics)
	if err != nil{
		return shim.Error("json.Unmarshal() ERROR: ", err.Error())
	}

	newPayload = args[1]
	analytics.Payload = newPayload
	err :=stub.PutState(DDAKey, newPayload)
	if err != nil{
		return shim.Error("PutState() ERROR: ", err.Error())
	}
	jsonDDA, err := json.Marshal(&analytics)
	if err != nil{
		return shim.Error("json.Marshal() ERROR: ", err.Error())
	}
	err := setDDAEvent(stub, jsonDDA)
	if err != nil{
		return shim.Error(" setEvent() ERROR: ", err.Error())
	}

	return shim.Success(nil)
}


func (t *DistributedDataAnalyticsChaincode) createAnalyticsInstances(stub shim.ChaincodeStubInterface, isEnabled bool, args []string) {

	var analytics AnalitycsInstances
	var analyticsEvent AnalitycsInstances
	var analytcsID, analytcsEGID string
	var payload string
	logger.info(" createAnalyticsInstances()\n")

	if len(args) == 2 {
		err := json.Unmarshal([]byte(args[0], &analytics))
		if (len(analytics.Id) == 0){
			analytcsID = xid.New()
			analytcsEGID = args[1]
			payload = args[0]
		} else {
			analyticsID = analytics.Id
			analytcsEGID = args[1]
			payload = analytics.Payload
		}
	} else {
		analytcsID = args[0]
		analytcsEGID = args[2]
		payload = args[1]
	}

	DDAKey, err := getAnalyticsKey(stub, analytcsID, analytcsEGID)
	if err!= nil{
		return shim.Error("CreateCompositeKey() ERROR: ",err.Error())
	}


	err :=stub.PutState(DDAKey, payload)
	if err != nil{
		return shim.Error("PutState() ERROR: " ,err.Error())
	}

	analyticsEvent.Id = analytcsID
	analyticsEvent.Egid = analytcsEGID
	analyticsEvent.Payload = payload
	jsonAnalytics, err := json.Marshal(&analyticsEvent)
	if err != nil{
		return shim.Error("json.Marshal() ERROR: ", err.Error())
	}
	err := setDDAEvent(stub, jsonAnalytics)
	if err != nil{
		return shim.Error(" setEvent() ERROR: ", err.Error())
	}

	return shim.Success([]byte(jsonAnalytics))

}


func main() {
	twc := new(DistributedDataAnalyticsChaincode)
	twc.testMode = true
	err := shim.Start(twc)
	if err != nil {
		logger.Error("Error starting Distributed-Data-Analytics chaincode: ", err)
	}
}
