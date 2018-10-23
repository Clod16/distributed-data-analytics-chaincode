# distributed-data-analytics-chaincode

<BR>
 
These are the APIs that we implement in the chaincode. These are the APIs that we implement in the chaincode.  

```java
    public interface AnalyticsLedgerClient {

    //String createAnalyticsInstances(AnalyticsInstances analyticsInstances);
    String createAnalyticsInstances(String id, String payload, String egid);
    void  updateAnalyticsInstances(String id, String payload, String egid);
    void delateAnalyticsInstances(String id, String egid);
    AnalyticsInstances getAnalyticsInstancesById(String id);
    AnalyticsInstances getAnalyticsInstancesByIdByEgid(String id, String egid);
    Collection<AnalyticsInstances> getAnalyticsInstances();

    //String createDataSources(DataSources dataSources) ;
    String createDataSources(String id, String payload, String egid) ;
    void updateDataSources(String id, String payload, String egid);
    void deleteDataSources(String id, String egid);
    Collection<DataSources> getDataSources() ;
    DataSources getDataSourcesbyId( String id) ;

    //String createEdgeGateways(EdgeGateways edgeGateways) ;
    String createEdgeGateways(String id, String payload, String egid);
    void updateEdgeGateways(String egid, String payload);
    Collection<EdgeGateways> getEdgeGateways() ;
    EdgeGateways getEdgeGatewaysByEgid(String egid);
    void deleteEdgeGateways(String egid);
}

```
