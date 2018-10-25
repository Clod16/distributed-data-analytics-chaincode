package main 

type AnalitycsInstances struct{

	Id                       string `json:"id"`
	Payload                  string `json:"payload"`
	Egid					 string `json:"egid"`	
	//EdgeGateway				 		`json:"edgeGateway"`
}

type  DataSource struct {

	Id                       string `json:"id"`
	Payload                  string `json:"payload"`
	Egid					 string `json:"egid"`			
	//EdgeGateway				 		`json:"edgeGateway"`
}


type EdgeGateway struct{

	Egid                       string `json:"egid"`
	Payload                  string `json:"payload"`
}