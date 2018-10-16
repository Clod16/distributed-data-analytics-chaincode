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
		return t.getAnalyticsInstances(stub, isEnabled, args)
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
    } else if function == "deleteEdgeGatewaysByEgid" {
		return t.deleteEdgeGatewaysByEgid(stub, isEnabled, args)
	 }
	return shim.Error("Invalid invoke function name")
}






func main() {
	twc := new(DistributedDataAnalyticsChaincode)
	twc.testMode = true
	err := shim.Start(twc)
	if err != nil {
		logger.Error("Error starting Distributed-Data-Analytics chaincode: ", err)
	}
}
