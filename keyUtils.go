

package main

import (
	"encoding/gob"
	"bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


func setAnalytcsEvent(stub shim.ChaincodeStubInterface, input string) (error){
	strs := []string{}
	strs = append(strs, input)
	buf := &bytes.Buffer{}
    gob.NewEncoder(buf).Encode(strs)
    bs := buf.Bytes()
	err:= stub.SetEvent("FE_Analytics_Instances", bs)
	return err
	
}
func getAnalyticsKey(stub shim.ChaincodeStubInterface, id string, egid string) (string, error) {
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

