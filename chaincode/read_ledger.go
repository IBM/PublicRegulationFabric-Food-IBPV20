/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
  // "reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// Read - read a generic variable from ledger
//
// Shows Off GetState() - reading a key/value from the ledger
//
// Inputs - Array of strings
//  0
//  key
//  "abc"
//
// Returns - string
// ============================================================================================================================
func read(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, jsonResp string
	var err error
	fmt.Println("starting read")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}

	// input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)           //get the var from ledger
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("Object keys")

	fmt.Println(string(valAsbytes))
	// fmt.Printf("%+v\n", security.Investor)


	fmt.Println("- end read")
	return shim.Success(valAsbytes)                  //send it onward
}

// ============================================================================================================================
// Get everything we need (owners + marbles + companies)
//
// Inputs - none
//
// Returns:
// {
//	"owners": [{
//			"id": "o99999999",
//			"company": "United Marbles"
//			"username": "alice"
//	}],
//	"marbles": [{
//		"id": "m1490898165086",
//		"color": "white",
//		"docType" :"marble",
//		"owner": {
//			"company": "United Marbles"
//			"username": "alice"
//		},
//		"size" : 35
//	}]
// }
// ============================================================================================================================
func read_everything(stub shim.ChaincodeStubInterface) pb.Response {
	type Everything struct {
		// Users   []User   `json:"users"`
		Products  			[]Product     `json:"products"`
		Retailers				  []Retailer				 `json:"retailers"`
		Importers		[]Importer		 `json:"importers"`
		Suppliers     []Supplier		 `json:"suppliers"`
		Regulators     []Regulator		 `json:"regulators"`
		ProductListingContracts []ProductListingContract `json:listingcontracts`
	}
	var everything Everything

	// ---- Get All Marbles ---- //
	productsIterator, err := stub.GetStateByRange("product0", "product9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer productsIterator.Close()

	for productsIterator.HasNext() {
		aKeyValue, err := productsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryValAsBytes := aKeyValue.Value
		var product Product
		json.Unmarshal(queryValAsBytes, &product)                  //un stringify it aka JSON.parse()
		everything.Products = append(everything.Products, product)   //add this marble to the list
	}

	retailersIterator, err := stub.GetStateByRange("retailer0", "retailer9999999999999999999")
	if err != nil {
	  return shim.Error(err.Error())
	}
	defer retailersIterator.Close()

	for retailersIterator.HasNext() {
	  aKeyValue, err := retailersIterator.Next()
	  if err != nil {
	    return shim.Error(err.Error())
	  }
	  queryValAsBytes := aKeyValue.Value
	  var retailer Retailer
	  json.Unmarshal(queryValAsBytes, &retailer)                  //un stringify it aka JSON.parse()
	  everything.Retailers = append(everything.Retailers, retailer)   //add this marble to the list
	}

	importersIterator, err := stub.GetStateByRange("importer0", "importer9999999999999999999")
	if err != nil {
	  return shim.Error(err.Error())
	}
	defer importersIterator.Close()

	for importersIterator.HasNext() {
	  aKeyValue, err := importersIterator.Next()
	  if err != nil {
	    return shim.Error(err.Error())
	  }
	  queryValAsBytes := aKeyValue.Value
	  var importer Importer
	  json.Unmarshal(queryValAsBytes, &importer)                  //un stringify it aka JSON.parse()
	  everything.Importers = append(everything.Importers, importer)   //add this marble to the list
	}

	suppliersIterator, err := stub.GetStateByRange("supplier0", "supplier9999999999999999999")
	if err != nil {
	  return shim.Error(err.Error())
	}
	defer suppliersIterator.Close()

	for suppliersIterator.HasNext() {
	  aKeyValue, err := suppliersIterator.Next()
	  if err != nil {
	    return shim.Error(err.Error())
	  }
	  queryValAsBytes := aKeyValue.Value
	  var supplier Supplier
	  json.Unmarshal(queryValAsBytes, &supplier)                  //un stringify it aka JSON.parse()
	  everything.Suppliers = append(everything.Suppliers, supplier)   //add this marble to the list
	}

	regulatorsIterator, err := stub.GetStateByRange("regulator0", "regulator9999999999999999999")
	if err != nil {
	  return shim.Error(err.Error())
	}
	defer regulatorsIterator.Close()

	for regulatorsIterator.HasNext() {
	  aKeyValue, err := regulatorsIterator.Next()
	  if err != nil {
	    return shim.Error(err.Error())
	  }
	  queryValAsBytes := aKeyValue.Value
	  var regulator Regulator
	  json.Unmarshal(queryValAsBytes, &regulator)                  //un stringify it aka JSON.parse()
	  everything.Regulators = append(everything.Regulators, regulator)   //add this marble to the list
	}

	productlistingcontractsIterator, err := stub.GetStateByRange("productlistingcontract0", "productlistingcontract9999999999999999999")
	if err != nil {
	  return shim.Error(err.Error())
	}
	defer productlistingcontractsIterator.Close()

	for productlistingcontractsIterator.HasNext() {
	  aKeyValue, err := productlistingcontractsIterator.Next()
	  if err != nil {
	    return shim.Error(err.Error())
	  }
	  queryValAsBytes := aKeyValue.Value
	  var productlistingcontract ProductListingContract
	  json.Unmarshal(queryValAsBytes, &productlistingcontract)                  //un stringify it aka JSON.parse()
	  everything.ProductListingContracts = append(everything.ProductListingContracts, productlistingcontract)   //add this marble to the list
	}

	fmt.Println("result", everything)

	//change to array of bytes
	everythingAsBytes, _ := json.Marshal(everything)              //convert to array of bytes
	return shim.Success(everythingAsBytes)
}

// ============================================================================================================================
// Get history of product
//
// Shows Off GetHistoryForKey() - reading complete history of a key/value
//
// Inputs - Array of strings
//  0
//  id
//  "m01490985296352SjAyM"
// ============================================================================================================================
func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	type AuditHistory struct {
		TxId    string   `json:"txId"`
		Value   Product   `json:"value"`
	}
	var history []AuditHistory;
	var product Product

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productId := args[0]
	fmt.Printf("- start getHistoryForProduct: %s\n", productId)

	// Get History
	resultsIterator, err := stub.GetHistoryForKey(productId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		historyData, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var tx AuditHistory
		tx.TxId = historyData.TxId                     //copy transaction id over
		json.Unmarshal(historyData.Value, &product)     //un stringify it aka JSON.parse()
		if historyData.Value == nil {                  //marble has been deleted
			var emptyProduct Product
			tx.Value = emptyProduct                 //copy nil marble
		} else {
			json.Unmarshal(historyData.Value, &product) //un stringify it aka JSON.parse()
			tx.Value = product                      //copy marble over
		}
		history = append(history, tx)              //add this tx to the list
	}
	fmt.Printf("- getHistoryForMarble returning:\n%s", history)

	//change to array of bytes
	historyAsBytes, _ := json.Marshal(history)     //convert to array of bytes
	return shim.Success(historyAsBytes)
}

// ============================================================================================================================
// Get history of product - performs a range query based on the start and end keys provided.
//
// Shows Off GetStateByRange() - reading a multiple key/values from the ledger
//
// Inputs - Array of strings
//       0     ,    1
//   startKey  ,  endKey
//  "marbles1" , "marbles5"
// ============================================================================================================================
func getMarblesByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryResultKey := aKeyValue.Key
		queryResultValue := aKeyValue.Value

		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResultKey)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResultValue))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getMarblesByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
