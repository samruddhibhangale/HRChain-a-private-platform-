package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	peer "github.com/hyperledger/fabric/protos/peer"
)

type hrchain struct {
}
type permissionToAccess struct {
	CandidateID string `json:"candidateID"`
	CompanyID   string `json:"companyID"`
}
type onbordingData struct {
	Company          string `json:"companyName"`
	AadharNumber     string `json:"aadharNumber"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	HighestEducation string `json:"highestEducation"`
	Experiance       string `json:"experiance"`
	SkillSet         string `json:"skillSet"`
	Profile          string `json:"profile"`
	JoiningDate      string `json:"joiningDate"`
}
type relievingData struct {
	Position        string `json:"position"`
	Experiance      string `json:"experiance"`
	RelievingDate   string `json:"relievingDate"`
	WorkingSkillSet string `json:"workingSkillSet"`
}

type candidateData struct {
	permissionToAccess `json:"permissionToAccess"`
	onbordingData      `json:"onbordingData"`
	relievingData      `json:"relievingData"`
}
type candidateBasicData struct {
	CandidateName string `json:"candidateName"`
	CompanyName   string `json:"companyName"`
}

func main() {
	fmt.Println("Function main")

	err := shim.Start(new(hrchain))

	if err != nil {
		fmt.Println("Error stating main function:")
	}

}

func (hrpoc *hrchain) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("------------Function inside Init")
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments, expecting none")
	}
	return shim.Success(nil)
}

func (hrpoc *hrchain) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("----------Function inside Invoke")
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is running     " + function)

	if function == "givePermission" {
		return hrpoc.givePermission(stub, args)
	} else if function == "pushOnbordingData" {
		return hrpoc.pushOnbordingData(stub, args)
	} else if function == "updateCanditateData" {
		return hrpoc.updateCanditateData(stub, args)
	} else if function == "getCandidateData" {
		return hrpoc.getCandidateData(stub, args)
	} else if function == "getCandidateBasicData" {
		return hrpoc.getCandidateBasicData(stub, args)
	} else if function == "pushCandidateBasicData" {
		return hrpoc.pushCandidateBasicData(stub, args)
	}

	fmt.Println("Invoke did not find func: " + function)
	return shim.Error("Received unknown function invocation")
}
func (hrpoc *hrchain) givePermission(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("inside givePermission function")

	if len(args) != 2 {
		return shim.Error("Incorrect no of args, expecting 2")
	}
	candidatedataAsBytes, err := stub.GetState(args[0])
	newCandidateData := &candidateData{}

	if len(candidatedataAsBytes) != 0 {
		fmt.Println("Updating the Information...")
		err = json.Unmarshal(candidatedataAsBytes, newCandidateData)

		candidateInfoJSON, err := json.Marshal(newCandidateData)
		fmt.Println("Fetched candidate data", candidateInfoJSON)
		if err != nil {
			shim.Error("Error while Marshalling")
		}
		newCandidateData.permissionToAccess.CandidateID = args[0]
		newCandidateData.permissionToAccess.CompanyID = args[1]

	} else {

		fmt.Println("start pushCandidateData")
		newCandidateData := &candidateData{

			permissionToAccess: permissionToAccess{
				CandidateID: args[0],
				CompanyID:   args[1],
			},
			onbordingData: onbordingData{
				Company:          "NA",
				AadharNumber:     "NA",
				FirstName:        "NA",
				LastName:         "NA",
				HighestEducation: "NA",
				Experiance:       "NA",
				SkillSet:         "NA",
				Profile:          "NA",
				JoiningDate:      "NA",
			},
			relievingData: relievingData{
				Position:        "NA",
				Experiance:      "NA",
				RelievingDate:   "NA",
				WorkingSkillSet: "NA",
			},
		}
		err = json.Unmarshal(candidatedataAsBytes, newCandidateData)
	}

	candidateInfoByte, err := json.Marshal(newCandidateData)
	fmt.Println("CandidateData", candidateInfoByte)

	err = stub.PutState(args[0], candidateInfoByte)
	if err != nil {
		return shim.Error("Error while PutState function")
	} else {
		result, _ := json.Marshal("Success")
		return shim.Success(result)
	}

}

func (hrpoc *hrchain) pushCandidateBasicDatar(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}
func (hrpoc *hrchain) pushOnbordingData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	fmt.Println("----- inside pushCandidateData function")

	if len(args) != 11 {
		return shim.Error("Incorrect no of args, expecting 10")
	}

	candidatedataAsBytes, err := stub.GetState(args[0])
	newCandidateData := candidateData{}

	if len(candidatedataAsBytes) != 0 {
		fmt.Println("Updating the Information...")
		err = json.Unmarshal(candidatedataAsBytes, &newCandidateData)

		if err != nil {
			shim.Error("Error while Marshilling")
		}
		newCandidateData.onbordingData.Company = args[2]
		newCandidateData.onbordingData.AadharNumber = args[3]
		newCandidateData.onbordingData.FirstName = args[4]
		newCandidateData.onbordingData.LastName = args[5]
		newCandidateData.onbordingData.HighestEducation = args[6]
		newCandidateData.onbordingData.Experiance = args[7]
		newCandidateData.onbordingData.SkillSet = args[8]
		newCandidateData.onbordingData.Profile = args[9]
		newCandidateData.onbordingData.JoiningDate = args[10]

		fmt.Println("Fetched candidate data", newCandidateData)

	} else {
		fmt.Println("Candidate does not exists. Please add candidate ID details first.")
		return shim.Error("Candidate Information does not exists")
	}
	if args[1] == newCandidateData.permissionToAccess.CompanyID {
		candidateInfoJSON, err := json.Marshal(newCandidateData)
		err = stub.PutState(args[0], candidateInfoJSON)
		if err != nil {
			return shim.Error("Error while pushingData")
		}
		fmt.Println("Candidates details are:-")
		fmt.Println(
			args[0], "\t",
			args[1], "\t",
			args[2], "\t",
			args[3], "\t",
			args[4], "\t")

	} else {
		fmt.Println("Not an valid User to do the operations")
		return shim.Error("Not an valid User to do the operations")
	}
	result, _ := json.Marshal("Success")
	return shim.Success(result)

}

func (hrpoc *hrchain) updateCanditateData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("Inside updateCandidateData function")
	var err error
	if len(args) != 6 {
		fmt.Println("Unexpected number of arguments, expecting 6")
	}
	candidatedataAsBytes, err := stub.GetState(args[0])
	newCandidateData := &candidateData{}
	fmt.Println(err)
	if len(candidatedataAsBytes) != 0 {
		fmt.Println("Updating the Information...")
		err = json.Unmarshal(candidatedataAsBytes, newCandidateData)

		candidateInfoJSON, err := json.Marshal(newCandidateData)
		fmt.Println("Fetched candidate data", candidateInfoJSON)
		if err != nil {
			shim.Error("Error while Marshilling")
		}
		newCandidateData.relievingData.Position = args[2]
		newCandidateData.relievingData.Experiance = args[3]
		newCandidateData.relievingData.RelievingDate = args[4]
		newCandidateData.relievingData.WorkingSkillSet = args[5]
	} else {
		fmt.Println("Candidate does not exists. Please add candidate ID details first.")
		return shim.Error("Candidate Information does not exists")
	}
	if args[1] == newCandidateData.permissionToAccess.CompanyID {
		candidateInfoJSON, err := json.Marshal(newCandidateData)
		fmt.Println("CandidateData", candidateInfoJSON)

		err = stub.PutState(args[0], candidateInfoJSON)
		if err != nil {
			return shim.Error("Error while PutState function")
		}
	} else {
		fmt.Println("Not an valid User to do the operations")
		return shim.Error("Not an valid User to do the operations")
	}
	result, _ := json.Marshal("Success")
	return shim.Success(result)

}

/*func GetMspID(stub ChaincodeStubInterface) string {
	msp, err := cid.GetMSPID(stub)
	if err != nil {
		fmt.Println("MSPID is-----------------")
	}
	fmt.Println("MSPID is-----------------")
	fmt.Println(msp)

	id, err := cid.GetID(stub)
	fmt.Println("MSPID is-----------------")
	fmt.Println(id)
}*/
/*func getTrueComapnyAccess(stub shim.ChaincodeStubInterface, args []string) bool {

	fetchedData := &permissionToAccess{}
	candidateDataAsByte, _ := stub.GetState(args[0])
	json.Unmarshal(candidateDataAsByte, fetchedData)
	candidateDataJSON, _ := json.Marshal(fetchedData)
	if fetchedData.CandidateID == args[0] && fetchedData.CompanyID == args[1] {
		return true
	}
	return false
}*/
func (hrpoc *hrchain) getCandidateBasicData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("----- inside getCandidateData function")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// candidateDataAsByte, _ := stub.GetState(args[0])
	// fetchedData := permissionToAccess{}
	// candidateData := candidateData{}
	// json.Unmarshal(candidateDataAsByte, &fetchedData)
	// json.Unmarshal(candidateDataAsByte, &candidateData)
	// //candidateDataJSON, _ := json.Marshal(fetchedData)
	// fmt.Println(fetchedData)
	// fmt.Println("=============")
	// fmt.Println(candidateData)

	// if args[0] == candidateData.permissionToAccess.CandidateID && args[1] == candidateData.permissionToAccess.CompanyID {

	// candidateDataAsByte1, _ := stub.GetState(args[0])
	// fetchedDeatilData := onbordingData{}
	// json.Unmarshal(candidateDataAsByte1, &fetchedDeatilData)

	// if candidateDataAsByte != nil {
	// 	fmt.Println(candidateDataAsByte1)
	// 	return shim.Success(candidateDataAsByte1)

	// }

	candidateDataAsByte1, _ := stub.GetState(args[0])
	fetchedData1 := onbordingData{}
	json.Unmarshal(candidateDataAsByte1, &fetchedData1)

	return shim.Success(fetchedData1)
	//}
	//return shim.Error("Record Not Found!")

}

func (hrpoc *hrchain) getCandidateData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("----- inside getCandidateData function")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")

	}
	// fetchedData := &onbordingData{}
	// candidateDataAsByte, _ := stub.GetState(args[0])
	// json.Unmarshal(candidateDataAsByte, fetchedData)
	// candidateDataJSON, err := json.Marshal(fetchedData)
	// if err != nil {
	// 	return shim.Error(" Please provide the valid Name")
	// } else if candidateDataJSON != nil {
	// 	fmt.Println(candidateDataJSON)

	// }
	// return shim.Success(candidateDataJSON)
	candidateDataAsByte, _ := stub.GetState(args[0])

	candidateData := candidateData{}
	json.Unmarshal(candidateDataAsByte, &candidateData)

	if args[0] == candidateData.permissionToAccess.CandidateID && args[1] == candidateData.permissionToAccess.CompanyID {
		historyQueryIterator, err := stub.GetHistoryForKey(args[0])

		// In case of error - return error
		if err != nil {
			return shim.Error("Error in fetching history !!!" + err.Error())
		}

		// Local variable to hold the history record
		var resultModification *queryresult.KeyModification
		counter := 0
		resultJSON := "["

		// Start a loop with check for more rows
		for historyQueryIterator.HasNext() {

			// Get the next record
			resultModification, err = historyQueryIterator.Next()

			if err != nil {
				return shim.Error("Error in reading history record!!!" + err.Error())
			}

			// Append the data to local variable
			data := "{\"txnid\": " + fmt.Sprintf(`"%s"`, resultModification.GetTxId())
			data += " , \"value\": " + string(resultModification.GetValue()) + "}  "
			if counter > 0 {
				data = ", " + data
			}
			resultJSON += data

			counter++
		}

		// Close the iterator
		historyQueryIterator.Close()

		// finalize the return string
		resultJSON += "]"
		resultJSON = "{ \"counter\": " + strconv.Itoa(counter) + ", \"txns\":" + resultJSON + "}"

		return shim.Success([]byte(resultJSON))
	}
	return shim.Error("Record Not Found!")
}
