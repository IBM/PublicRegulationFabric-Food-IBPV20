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
	"encoding/json"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ============================================================================================================================
// Get Product - get a product product from ledger
// ============================================================================================================================
func get_product(stub shim.ChaincodeStubInterface, id string) (Product, error) {
	var product Product
	productAsBytes, err := stub.GetState(id)                  //getState retreives a key/value from the ledger
	if err != nil {                                          //this seems to always succeed, even if key didn't exist
		return product, errors.New("Failed to find product - " + id)
	}
	json.Unmarshal(productAsBytes, &product)                   //un stringify it aka JSON.parse()

	if product.Id != id {                                     //test if product is actually here or just nil
		return product, errors.New("Product does not exist - " + id)
	}

	return product, nil
}

// ============================================================================================================================
// Get Owner - get the owner product from ledger
// ============================================================================================================================
func get_user(stub shim.ChaincodeStubInterface, id string) (User, error) {
	var user User
	userAsBytes, err := stub.GetState(id)                     //getState retreives a key/value from the ledger
	if err != nil {                                            //this seems to always succeed, even if key didn't exist
		return user, errors.New("Failed to get User - " + id)
	}
	json.Unmarshal(userAsBytes, &user)                       //un stringify it aka JSON.parse()

	// if len(user.id) == 0 {                              //test if owner is actually here or just nil
	// 	return user, errors.New("User does not exist - " + id )
	// }

	return user, nil
}

func get_regulator(stub shim.ChaincodeStubInterface, id string) (Regulator, error) {
	var regulator Regulator
	regulatorAsBytes, err := stub.GetState(id)                     //getState retreives a key/value from the ledger
	if err != nil {                                            //this seems to always succeed, even if key didn't exist
		return regulator, errors.New("Failed to get Regulator - " + id)
	}
	json.Unmarshal(regulatorAsBytes, &regulator)                       //un stringify it aka JSON.parse()

	if len(regulator.Id) == 0 {                              //test if owner is actually here or just nil
		return regulator, errors.New("Regulator does not exist - " + id)
	}

	return regulator, nil
}

// ========================================================
// Input Sanitation - dumb input checking, look for empty strings
// ========================================================
func sanitize_arguments(strs []string) error{
	for i, val:= range strs {
		if len(val) <= 0 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be a non-empty string")
		}
		if len(val) > 32 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be <= 32 characters")
		}
	}
	return nil
}
