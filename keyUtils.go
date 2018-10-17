

package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

func setDDAEvent(stub shim.ChaincodeStubInterface, json string) (error){
	err:= stub.SetEvent("FE_Analytics_Instances", json)
	return err
	
}
func getDDAKey(stub shim.ChaincodeStubInterface, id string, egid string) (string, error) {
	ddaKey, err := stub.CreateCompositeKey("FE_Analytics_Instances:", []string{id, egid})
	if err != nil {
		return "", err
	} else {
		return ddaKey, nil
	}

}



func BytesToString(data []byte) string {
	return string(data[:])
}

