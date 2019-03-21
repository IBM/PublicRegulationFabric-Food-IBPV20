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
  // "encoding/csv"
	"fmt"
	// "strconv"
	"strings"
	// "reflect"
	// "math"
	// "time"
	// "encoding/gob"
	// "bytes"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// write() - genric write variable into ledger
//
// Shows Off PutState() - writting a key/value into the ledger
//
// Inputs - Array of strings
//    0   ,    1
//   key  ,  value
//  "abc" , "test"
// ============================================================================================================================
func write(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, value string
	var err error
	fmt.Println("starting write")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2. key of the variable and value to set")
	}

	// input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	key = args[0]                                   //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value))         //write the variable into the ledger
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end write")
	return shim.Success(nil)
}

// ============================================================================================================================
// delete_marble() - remove a marble from state and from marble index
//
// Shows Off DelState() - "removing"" a key/value from the ledger
//
// Inputs - Array of strings
//      0      ,         1
//     id      ,  authed_by_company
// "m999999999", "united marbles"
// ============================================================================================================================
func delete(stub shim.ChaincodeStubInterface, args []string) (pb.Response) {
	fmt.Println("starting delete")

	id := args[0]

	err := stub.DelState(id)                                                 //remove the key from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	fmt.Println("- end delete")
	return shim.Success(nil)
}

// ============================================================================================================================
// Init Product - create a new product, store into chaincode state
//
// Shows off building a key's JSON value manually
//
// Inputs - Array of strings

//      0      ,    1  ,     				2  ,      																			3          			 ,    4     ,     5					, 	6
//     id      ,  loan amount ,  borrower_info , 															, state								 , interest ,  balance due	, grade
// "m999999999", "545,000"    ,   object																				deliquent/in payment , 	3.0    , 		520,000			,  BBB
															// credit/income verification/debt to income,

// ============================================================================================================================
func init_product(stub shim.ChaincodeStubInterface, args []string) (pb.Response) {
	var err error
	fmt.Println("starting init_product")
	var product Product
	product.Id = args[0]
	product.Quantity =  args[1]
	product.CountryId = args[2]
	// check if product already exists
	// TODO, uncomment
	// _, err = get_product(stub, product.Id)
	// if err == nil {
	// 	fmt.Println("This product already exists - " + product.Id)
	// 	return shim.Error("This product already exists - " + product.Id)
	// }
	//store product
	productAsBytes, _ := json.Marshal(product)                         //convert to array of bytes
	fmt.Println("writing product to state")
	fmt.Println(string(productAsBytes))
	err = stub.PutState(product.Id, productAsBytes)                    //store owner by its Id
	if err != nil {
		fmt.Println("Could not store product")
		return shim.Error(err.Error())
	}
	fmt.Println("- end init_product")

	return shim.Success(nil)
}


// update_product


// generate_securities







// pool_product (product)
// underwriting info gives grade
// grade determines pool
// pool determines return, but also likliehood of failure/delinquency


// ============================================================================================================================
// Init Owner - create a new owner aka end user, store into chaincode state
//
// Shows off building key's value from GoLang Structure
//
// Inputs - Array of Strings
//           0     ,     1   ,   2
//      owner id   , username, company
// "o9999999999999",     bob", "united marbles"
// ============================================================================================================================
func init_user(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("starting init_user")

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

  // TODO, should probably allow JSON object as third arg
	// user.ObjectType = "product_user"
  id := args[0]
  userType := args[1]

  var user User
	user.Id = id
  user.Type = userType

  // userType := "supplier"
  // switch user.Type := {

  // TODO, remove userType arg
  // if strings.Contains(id, "supplier") {
  //
  // } else if strings.Contains(id, "importer") {
  //
  // } else if strings.Contains(id, "retailer") {
  //
  // }
  switch userType {
    case "supplier":
      var supplier Supplier
      supplier.User = user
      supplier.countryId = args[2]
      supplier.orgId = args[3]
      supplierAsBytes, _ := json.Marshal(supplier)                         //convert to array of bytes
    	err = stub.PutState(id, supplierAsBytes)                    //store owner by its Id
    	if err != nil {
    		fmt.Println("Could not store supplier")
    		return shim.Error(err.Error())
    	}
    case "importer":
      var importer Importer
      importer.User = user
      importerAsBytes, _ := json.Marshal(importer)                         //convert to array of bytes
      err = stub.PutState(id, importerAsBytes)                    //store owner by its Id
    	if err != nil {
    		fmt.Println("Could not store importer")
    		return shim.Error(err.Error())
    	}
    case "retailer":
      var retailer Retailer
      retailer.User = user
      retailer.Products = nil //[]Product
      retailerAsBytes, _ := json.Marshal(retailer)                         //convert to array of bytes
      err = stub.PutState(id, retailerAsBytes)                    //store owner by its Id
    	if err != nil {
    		fmt.Println("Could not store retailer")
    		return shim.Error(err.Error())
    	}
  }
	fmt.Println("- end init_user")
	return shim.Success(nil)
}



func init_product_listing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("starting init_product_listing")

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}
	// user.ObjectType = "product_user"
  product_listing_id := args[0]
  supplier_id := args[1]
  // product_ids := args[2] // Expecting JSON array of product ids. TODO, csv would probably be easier?

  productListing := ProductListingContract{}
  productListing.Id = product_listing_id
  productListing.Status = "INITIALREQUEST"
  productListing.Owner = supplier_id
  productListing.Supplier = supplier_id
  productListing.OwnerType = "Supplier"

  numProducts := len(args) - 2
  var products = make([]string, numProducts)
  // all array elements after first 2 are parsed as product ids
  for i := 2; i < len(args); i++ {
    products[i - 2] = args[i]
  }
  productListing.Products = products //csv.NewReader(product_ids) //json.Unmarshal(product_ids)

  // TODO? update product location to same as supplier
  // supplierAsBytes, err := stub.GetState(supplier_id)
  // supplier := Supplier{}
	// err = json.Unmarshal(supplierAsBytes, &supplier)           //un stringify it aka JSON.parse()
	// if err != nil {
	// 	return shim.Error("Error loading supplier")
	// }
  productListingAsBytes, _ := json.Marshal(productListing)                         //convert to array of bytes
  err = stub.PutState(product_listing_id, productListingAsBytes)                    //store owner by its Id
	if err != nil {
		fmt.Println("Could not store product listing")
		return shim.Error(err.Error())
	}
	fmt.Println("- end init_product_listing")
	return shim.Success(nil)
}

func init_regulator(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("starting init_regulator")

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}
  // regulator_id := args[0]
  regulator := Regulator{}
  regulator.Id = args[0]
  regulator.countryId = args[1]
  regulatorAsBytes, _ := json.Marshal(regulator)                         //convert to array of bytes
  err = stub.PutState(regulator.Id, regulatorAsBytes)                    //store owner by its Id
	if err != nil {
		fmt.Println("Could not store regulator")
		return shim.Error(err.Error())
	}
	fmt.Println("- end init_regulator")
	return shim.Success(nil)
}

func transfer_product_listing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	fmt.Println("-starting transfer_product_listing")
  //input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}
  product_listing_id := args[0]
  new_owner_id := args[1]
  // user_type := args[1]
  // user_id := args[1]
  // retailer_id := args[2]

  productListingAsBytes, err := stub.GetState(product_listing_id)
  productListing := ProductListingContract{}
	err = json.Unmarshal(productListingAsBytes, &productListing)           //un stringify it aka JSON.parse()
	if err != nil {
		fmt.Println(string(productListingAsBytes))
		return shim.Error(err.Error())
	}
  productListing.Owner = new_owner_id
  if (strings.ToLower(productListing.OwnerType) == "supplier") {
    productListing.OwnerType = "Importer"
    productListing.Status = "EXEMPTCHECKREQ"
    productListing.Owner = new_owner_id

  } else if ( strings.ToLower(productListing.OwnerType) == "importer" ) {
    productListing.OwnerType = "Retailer"
    if ( productListing.Status == "EXEMPTCHECKREQ" ){
      return shim.Error("Products in listing need to be checked by regulator.")
    } else if ( productListing.Status == "HAZARDANALYSISCHECKREQ" ){
      return shim.Error("Products cannot be transferred as they've been flagged by regulator.")
    }
    retailerAsBytes, err := stub.GetState(new_owner_id)
    retailer := Retailer{}
  	err = json.Unmarshal(retailerAsBytes, &retailer)           //un stringify it aka JSON.parse()
  	if err != nil {
  		fmt.Println(string(retailerAsBytes))
  		return shim.Error(err.Error())
  	}
    // _, products := json.Marshal(productListing.Products)
    for _, product := range productListing.Products {
      retailer.Products = append(retailer.Products, product)
    }
    retailerAsBytes, _ = json.Marshal(retailer)           //convert to array of bytes
    err = stub.PutState(new_owner_id, retailerAsBytes)     //rewrite the marble with id as key
    if err != nil {
      return shim.Error(err.Error())
    }
  } else {
      return shim.Error("Invalid user type provided.")
  }
  productListingAsBytes, _ = json.Marshal(productListing)                         //convert to array of bytes
  err = stub.PutState(product_listing_id, productListingAsBytes)                    //store owner by its Id
	if err != nil {
		fmt.Println("Could not store product listing")
		return shim.Error(err.Error())
	}
  fmt.Println("- end transfer_product_listing")
	return shim.Success(nil)
}

func update_exempted_list(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  fmt.Println("- start update_exempted_list")
  // ["regulator1", "org", "org1", "org2"..]
  // add list of exempted orgs or products to regulator
  regulator_id := args[0]
  exempted_type := args[1]
  // all remaining args are list of ids


  regulatorAsBytes, err := stub.GetState(regulator_id)
  regulator := Regulator{}
	err = json.Unmarshal(regulatorAsBytes, &regulator)           //un stringify it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

  numIds := len(args) - 2
  ids := make([]string, numIds)
  // all array elements after first 2 are parsed as product ids
  for i := 2; i < len(args); i++ {
    ids[i - 2] = args[i]
  }

  // TODO, should probably append list
  if (exempted_type == "org") {
    regulator.ExemptedOrgIds = ids
  } else if (exempted_type == "product") {
    regulator.ExemptedProductIds = ids
  }

  regulatorAsBytes, _ = json.Marshal(regulator)                         //convert to array of bytes
  err = stub.PutState(regulator_id, regulatorAsBytes)                    //store owner by its Id
	if err != nil {
		fmt.Println("Could not store regulator")
		return shim.Error(err.Error())
	}
  fmt.Println("- end update_exempted_list")
	return shim.Success(nil)

}

func check_products(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  product_listing_id := args[0]
  regulator_id := args[1]

  productListingAsBytes, err := stub.GetState(product_listing_id)
  productListing := ProductListingContract{}
	err = json.Unmarshal(productListingAsBytes, &productListing)           //un stringify it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

  supplierAsBytes, err := stub.GetState(productListing.Supplier)
  supplier := Supplier{}
	err = json.Unmarshal(supplierAsBytes, &supplier)           //un stringify it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

  regulatorAsBytes, err := stub.GetState(regulator_id)
  regulator := Regulator{}
	err = json.Unmarshal(regulatorAsBytes, &regulator)           //un stringify it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

  if (productListing.Status != "EXEMPTCHECKREQ" && productListing.Status != "HAZARDANALYSISCHECKREQ"){
    return shim.Error("Invalid state, listing cannot be checked");
  }

  check := true
  if (productListing.Status=="EXEMPTCHECKREQ"){
    // check if the supplier org is exempted by regulator

    for _, orgId := range regulator.ExemptedOrgIds {
      if ( supplier.orgId == orgId ) {
        check = false
      }
    }
  }

  if (check) {
    productListing.Status="CHECKCOMPLETED"
  } else {
    productListing.Status="HAZARDANALYSISCHECKREQ"
  }

  productListingAsBytes, _ = json.Marshal(productListing)           //convert to array of bytes
  err = stub.PutState(product_listing_id, productListingAsBytes)     //rewrite the marble with id as key
  if err != nil {
    return shim.Error(err.Error())
  }
  fmt.Println("- end check_products")
	return shim.Success(nil)
}
