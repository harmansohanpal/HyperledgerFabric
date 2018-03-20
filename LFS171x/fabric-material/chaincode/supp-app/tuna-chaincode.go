package main

import (
	//"errors"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Supp struct {
	Emailaddress string `json:"emailaddress"`
	Name string `json:"name"`
	Role  string `json:"role"`
	Holder  string `json:"holder"`
	Balance  string `json:"balance"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}


func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordSupp" {
		return s.recordSupp(APIstub, args)
	} else if function == "queryAllSupp" {
		return s.queryAllSupp(APIstub)
	} 
	// else if function == "changeSuppHolder" {
	// 	return s.changeSuppHolder(APIstub, args)
	// }

	return shim.Error("Invalid Smart Contract function name.")
}



func (t *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	fmt.Printf("Init called, initializing chaincode")
	supp := []Supp{
		Supp{Emailaddress: "shikherwalia@gmail.com", Name: "Shikher", Role: "Operator", Holder: "A", Balance: "0"},
	    Supp{Emailaddress: "harman@gmail.com", Name: "Harman", Role: "Operator", Holder: "B", Balance: "0"},
	   	Supp{Emailaddress: "akhil@gmail.com", Name: "Akhil", Role: "Operator", Holder: "C", Balance: "0"},
		Supp{Emailaddress: "ram@gmail.com", Name: "Ram", Role: "Customer", Holder: "", Balance: "1000"},
		Supp{Emailaddress: "shyam@gmail.com", Name: "Shyam", Role: "Cutomer", Holder: "", Balance: "1000"},
       	Supp{Emailaddress: "sita@gmail.com", Name: "Sita", Role: "Customer", Holder: "", Balance: "1000"},

}

	i := 0
	for i < len(supp) {
		fmt.Println("i is ", i)
		suppAsBytes, _ := json.Marshal(supp[i])
		APIstub.PutState(strconv.Itoa(i+1), suppAsBytes)
		fmt.Println("Added", supp[i])
		i = i + 1
	}

	return shim.Success(nil)
}


func (s *SmartContract) recordSupp(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	fmt.Println(args[1])
	var supp = Supp{ Emailaddress: args[1], Name: args[2], Role: args[3], Holder: args[4], Balance: args[5] }

	suppAsBytes, _ := json.Marshal(supp)
	err := APIstub.PutState(args[0], suppAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record : %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryAllSupp(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllSupp:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


func (s *SmartContract) changeSuppHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	suppAsBytes, _ := APIstub.GetState(args[0])
	if suppAsBytes == nil {
		return shim.Error("Could not locate supp")
	}
	supp := Supp{}

	json.Unmarshal(suppAsBytes, &supp)
	supp.Holder = args[1]

	suppAsBytes, _ = json.Marshal(supp)
	err := APIstub.PutState(args[0], suppAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change holder: %s", args[0]))
	}

	return shim.Success(nil)
}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

