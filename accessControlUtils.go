
package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, string, error) {

	var err error
	var attrValue1, attrValue2 string
	var found bool
	//FIXME ROLE-UID!!
	const ROLE string = ""
	const UID string = ""

	attrValue1, found, err = cid.GetAttributeValue(stub, ROLE)
	if err != nil {
		fmt.Printf("Error getting Attribute Value: %s\n", err.Error())
		return "", "", err
	}
	if found == false {
		fmt.Printf("Error getting ROLE --> NOT FOUND!!!\n")
	//	err.Error()
	//	return "", "", err
	}

	attrValue2, found, err = cid.GetAttributeValue(stub, UID)
	if err != nil {
		fmt.Printf("Error getting Attribute Value UID: %s\n", err.Error())
		return "", "", err
	}
	if found == false {
		fmt.Printf("Error getting UID --> NOT FOUND!!!\n")
		return "", "", err
	}

	return attrValue1, attrValue2 , nil
}

func isInvokerOperator(stub shim.ChaincodeStubInterface, attrName string) (bool, string, error) {
	var found bool
	var attrValue string
	var err error

	attrValue, found, err = cid.GetAttributeValue(stub, attrName)
	if err != nil {
		fmt.Printf("Error getting Attribute Value: %s\n", err.Error())
		return false, "", err
	}
	return found, attrValue, nil
}