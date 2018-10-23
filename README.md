# distributed-data-analytics-chaincode


#### Possible API's
 
These are the APIs that we implement in the chaincode. 

```java
    public interface AnalyticsLedgerClient {

    //String createAnalyticsInstances(AnalyticsInstances analyticsInstances);
    String createAnalyticsInstances(String id, String payload, String egid);
    //void  updateAnalyticsInstances(AnalyticsInstances analyticsInstances);
    void  updateAnalyticsInstances(String id, String payload, String egid);
    void delateAnalyticsInstances(String id, String egid);
    AnalyticsInstances getAnalyticsInstancesById(String id);
    AnalyticsInstances getAnalyticsInstancesByIdByEgid(String id, String egid);
    Collection<AnalyticsInstances> getAnalyticsInstances();

    //String createDataSources(DataSources dataSources) ;
    String createDataSources(String id, String payload, String egid) ;
    //void updateDataSources(DataSources dataSources) ;
    void updateDataSources(String id, String payload, String egid);
    void deleteDataSources(String id, String egid);
    Collection<DataSources> getDataSources() ;
    DataSources getDataSourcesbyId( String id) ;
    DataSources getDataSourcesbyIdByEgid( String id, String egid) ;

    //String createEdgeGateways(EdgeGateways edgeGateways) ;
    String createEdgeGateways(String id, String payload, String egid);
    //void updateEdgeGateways(EdgeGateways edgeGateways);
    void updateEdgeGateways(String egid, String payload);
    Collection<EdgeGateways> getEdgeGateways() ;
    EdgeGateways getEdgeGatewaysByEgid(String egid);
    void deleteEdgeGateways(String egid);
    
}

```

#### Possible data structures


```java
public class AnalyticsInstances {
    private String id;
    private String payload;
    //private Collection<EdgeGateways> edgeGatewaysArrayList;
    private String egid;
    }   
public class DataSources {
    private String id;
    private String payload;
    private String egid;
    //private Collection<EdgeGateways> edgeGatewaysCollection;
    }
 public class EdgeGateways {
    private String egid;
    private String payload;
    }
```

```go
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
```
