package main

import (
	"encoding/gob"
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
	/*} else if function == "deleteDataSources" {
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
		return t.deleteEdgeGateways(stub, isEnabled, args) */
	}
	return shim.Error("Invalid invoke function name")
}


func (t *DistributedDataAnalyticsChaincode) createDataSources(stub shim.ChaincodeStubInterface, isEnabled bool, args []string) pb.Response {

	var data DataSource
	var dataEvent DataSource
	var dataID, dataEGID string
	var payload string
	var err error
	logger.Info(" createDataSources()\n")

	if len(args) == 2 {
		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(args[0])
		bs := buf.Bytes()
		err = json.Unmarshal(bs, &data)
		if err != nil{
			return shim.Error(" json.Unmarshal() ERROR: " +err.Error())
		}else{
			if (len(data.Id) == 0){
				xidAnalytics := xid.New()
				dataID = xidAnalytics.String()
				dataEGID = args[1]
				payload = args[0]
			} else {
				dataID = data.Id
				dataEGID = args[1]
				payload = data.Payload
		}}
	} else {
		dataID = args[0]
		dataEGID = args[2]
		payload = args[1]
	}

	DDAKey, err1 := getAnalyticsKey(stub, dataID, dataEGID)
	if err1!= nil{
		return shim.Error("CreateCompositeKey() ERROR: " +err1.Error())
	}

	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(payload)
	bufByte := buf.Bytes()
	err = stub.PutState(DDAKey, bufByte)
	if err != nil{
		return shim.Error("PutState() ERROR: " +err.Error())
	}

	analyticsEvent.Id = dataID
	analyticsEvent.Egid = dataEGID
	analyticsEvent.Payload = payload
	jsonAnalytics, err2 := json.Marshal(&analyticsEvent)
	if err2 != nil{
		logger.Error("Error starting Distributed-Data-Analytics chaincode: ", err)
		return shim.Error("json.Marshal() ERROR: " +err2.Error())
	}

	stringEvent := "createAnalyticsInstances Event --- payload:" +string(jsonAnalytics)
	err = setAnalytcsEvent(stub, stringEvent)
	if err != nil{
		return shim.Error(" setEvent() ERROR: " +err.Error())
	}

	return shim.Success([]byte(jsonAnalytics))


	

}
/*
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
} */

func (t *DistributedDataAnalyticsChaincode) getAnalyticsInstances(stub shim.ChaincodeStubInterface, isEnabled bool) pb.Response {

	logger.Info(" getAnalyticsInstances()\n")
	var emptyArgs  []string
	var bufByte []byte
	//var analytic *AnalitycsInstances
	var analyticsArrayString  []string
	analyticsResponse, err := stub.GetStateByPartialCompositeKey("FE_Analytics_Instances",emptyArgs)
	if err != nil{
		logger.Error(" GetStateByPartialCompositeKey() ERROR:\n")
		return shim.Error(err.Error())
	}
	for analyticsResponse.HasNext(){
		analyticsArray, err1 := analyticsResponse.Next()
		if err1 != nil {
			return shim.Error(err1.Error())
		}
		payloadByte := analyticsArray.Value
		payload := BytesToString(payloadByte)
		analyticsArrayString = append(analyticsArrayString, payload)
		buf :=&bytes.Buffer{}
		gob.NewEncoder(buf).Encode(analyticsArrayString)
		bufByte = buf.Bytes()
	}
	stringEvent := " getAnalyticsInstances Event :" +analyticsArrayString
	err2 := setAnalytcsEvent(stub, stringEvent )
	if err2 != nil{
		return shim.Error(" setEvent() ERROR: " +err2.Error())
	}	
	return shim.Success(bufByte)

	
}
func (t *DistributedDataAnalyticsChaincode) getAnalyticsInstancesByEgid(stub shim.ChaincodeStubInterface, isEnabled bool, args []string)pb.Response  {
	logger.Info(" getAnalyticsInstancesByEgid()\n")
	var analyticsArrayString  []string
	var bufByte []byte

	if len(args) != 1{
		logger.Error(" getAnalyticsInstancesByEgid() ERROR: wrong argument\n")
		return shim.Error("getAnalyticsInstancesByEgid() ERROR: wrong argument" )
	}
	analyticsResponse, err := stub.GetStateByPartialCompositeKey("FE_Analytics_Instances" , args)
	if err != nil{
		logger.Error(" GetStateByPartialCompositeKey() ERROR:\n")
		return shim.Error(err.Error())
	}

	for analyticsResponse.HasNext(){
		analyticsArray, err1 := analyticsResponse.Next()
		if err1 != nil {
			return shim.Error(err1.Error())
		}
		payloadByte := analyticsArray.Value
		payload := BytesToString(payloadByte)
		analyticsArrayString = append(analyticsArrayString, payload)
		buf :=&bytes.Buffer{}
		gob.NewEncoder(buf).Encode(analyticsArrayString)
		bufByte = buf.Bytes()
	}
	stringEvent := "getAnalyticsInstancesByEgid Event :" +analyticsArrayString
	err2 := setAnalytcsEvent(stub, stringEvent )
	if err2 != nil{
		return shim.Error(" setEvent() ERROR: " +err2.Error())
	}	
	return shim.Success(bufByte)

	
}


func (t *DistributedDataAnalyticsChaincode) getAnalyticsInstancesById(stub shim.ChaincodeStubInterface, isEnabled bool, args []string) pb.Response  {

	logger.Info(" getAnalyticsInstancesById()\n")
	var analyticsArrayString  []string
	var bufByte []byte

	//var analyticsArrayString []string
	if len(args) != 1{
		logger.Error(" getAnalyticsInstancesById() ERROR: wronh argument\n")
		return shim.Error("getAnalyticsInstancesById() ERROR: wrong argument" )
	}
	analyticsResponse, err := stub.GetStateByPartialCompositeKey("FE_Analytics_Instances" , args)
	if err != nil{
		logger.Error(" GetStateByPartialCompositeKey() ERROR:\n")
		return shim.Error(err.Error())
	}
	for analyticsResponse.HasNext(){
		analyticsArray, err1 := analyticsResponse.Next()
		if err1 != nil {
			return shim.Error(err1.Error())
		}
		payloadByte := analyticsArray.Value
		payload := BytesToString(payloadByte)
		analyticsArrayString = append(analyticsArrayString, payload)
		buf :=&bytes.Buffer{}
		gob.NewEncoder(buf).Encode(analyticsArrayString)
		bufByte = buf.Bytes()
	}

	stringEvent := "getAnalyticsInstancesById Event :" +analyticsArrayString
	err2 := setAnalytcsEvent(stub, stringEvent )
	if err2 != nil{
		return shim.Error(" setEvent() ERROR: " +err2.Error())
	}	
	return shim.Success(bufByte)

	
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

}
func (t *DistributedDataAnalyticsChaincode) delateAnalyticsInstances(stub shim.ChaincodeStubInterface, isEnabled bool, args []string) pb.Response {

	logger.Info(" delateAnalyticsInstances()\n")

	if len(args) != 2{
		return shim.Error("delateAnalyticsInstances() ERROR: wrong argument" )
	}

	DDAKey , err := getAnalyticsKey(stub, args[0], args[1])
	if err != nil{
		return shim.Error("CreateCompositeKey() ERROR: " +err.Error())
	}

	err1 := stub.DelState(DDAKey)
	if err1 != nil{
		return shim.Error("DelState() ERROR" +err1.Error()) 
	}
	stringEvent :=  "Analytics Instances deleted with key: " +DDAKey
	err2 := setAnalytcsEvent(stub, stringEvent )
	if err2 != nil{
		return shim.Error(" setEvent() ERROR: " +err2.Error())
	}	
	return shim.Success(nil)

}

func (t *DistributedDataAnalyticsChaincode) updateAnalyticsInstances(stub shim.ChaincodeStubInterface, isEnabled bool, args []string) pb.Response {

	var analytics AnalitycsInstances
	logger.Info(" updateAnalyticsInstances()\n")

	if len(args) != 3{
		return shim.Error("updateAnalyticsInstances() ERROR: wrong argument")
	}

	DDAKey, err := getAnalyticsKey(stub, args[0], args[2])
	if err != nil{
			return shim.Error("CreateCompositeKey() ERROR: " +err.Error())
		}
	

	DDABytes, err1 := stub.GetState(DDAKey)
	if err1 != nil{
		return shim.Error("GetState() ERROR: " +err1.Error())
	}
	
	err2 := json.Unmarshal([]byte(DDABytes), &analytics)
	if err2 != nil{
		return shim.Error("json.Unmarshal() ERROR: " +err2.Error())
	}

	newPayload := args[1]
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(newPayload)
	bufByte := buf.Bytes()
	analytics.Payload = newPayload
	err4 := stub.PutState(DDAKey, bufByte)
	if err4 != nil{
		return shim.Error("PutState() ERROR: " +err4.Error())
	}
	jsonDDA, err3 := json.Marshal(&analytics)
	if err3 != nil{
		return shim.Error("json.Marshal() ERROR: " +err3.Error())
	}
	stringEvent := "updateAnalyticsInstances Event --- New payload :" +string(jsonDDA)
	err0:= setAnalytcsEvent(stub, stringEvent)
	if err0 != nil{
		return shim.Error(" setEvent() ERROR: " +err0.Error())
	}

	return shim.Success(nil)
}


func (t *DistributedDataAnalyticsChaincode) createAnalyticsInstances(stub shim.ChaincodeStubInterface, isEnabled bool, args []string) pb.Response {

	var analytics AnalitycsInstances
	var analyticsEvent AnalitycsInstances
	var analyticsID, analytcsEGID string
	var payload string
	var err error
	logger.Info(" createAnalyticsInstances()\n")

	if len(args) == 2 {
		buf := &bytes.Buffer{}
		gob.NewEncoder(buf).Encode(args[0])
		bs := buf.Bytes()
		err = json.Unmarshal(bs, &analytics)
		if err != nil{
			return shim.Error(" json.Unmarshal() ERROR: " +err.Error())
		}else{
			if (len(analytics.Id) == 0){
				xidAnalytics := xid.New()
				analyticsID = xidAnalytics.String()
				analytcsEGID = args[1]
				payload = args[0]
			} else {
				analyticsID = analytics.Id
				analytcsEGID = args[1]
				payload = analytics.Payload
		}}
	} else {
		analyticsID = args[0]
		analytcsEGID = args[2]
		payload = args[1]
	}

	DDAKey, err1 := getAnalyticsKey(stub, analyticsID, analytcsEGID)
	if err1!= nil{
		return shim.Error("CreateCompositeKey() ERROR: " +err1.Error())
	}

	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(payload)
	bufByte := buf.Bytes()
	err = stub.PutState(DDAKey, bufByte)
	if err != nil{
		return shim.Error("PutState() ERROR: " +err.Error())
	}

	analyticsEvent.Id = analyticsID
	analyticsEvent.Egid = analytcsEGID
	analyticsEvent.Payload = payload
	jsonAnalytics, err2 := json.Marshal(&analyticsEvent)
	if err2 != nil{
		logger.Error("Error starting Distributed-Data-Analytics chaincode: ", err)
		return shim.Error("json.Marshal() ERROR: " +err2.Error())
	}

	stringEvent := "createAnalyticsInstances Event --- payload:" +string(jsonAnalytics)
	err = setAnalytcsEvent(stub, stringEvent)
	if err != nil{
		return shim.Error(" setEvent() ERROR: " +err.Error())
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
