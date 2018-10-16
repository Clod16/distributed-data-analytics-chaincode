##### ANALYTICS INSTANCES

**POST /analytics-instances?edgeGatewayReferenceID=:egid**

Creates an analytics instance either at the edge or at the ledger tier.

**PUT /analytics-instances/:id/specification?edgeGatewayReferenceID=:egid**

Changes the specification of the analytics instance with the specified ID that
exists either at the edge or at the ledger tier.

**DELETE /analytics-instances/:id?edgeGatewayReferenceID=:egid**

Deletes the analytics instance with the specified ID that exists either at the
edge or at the ledger tier.

**GET /analytics-instances/:id?edgeGatewayReferenceID=:egid**

Gets the analytics instance with the specified ID that exists either at the
edge or at the ledger tier.

**GET /analytics-instances/:id/specification?edgeGatewayReferenceID=:egid**

Gets the specification of the analytics instance with the specified ID that
exists either at the edge or at the ledger tier.

**GET /analytics-instances/:id/state?edgeGatewayReferenceID=:egid**

Gets the state of the analytics instance with the specified ID that exists
either at the edge or at the ledger tier.

**POST /analytics-instances/discover/discover?edgeGatewayReferenceID=:egid**

Discovers analytics instances that match the specified criteria either at the
edge or at the ledger tier.

**POST /analytics-instances/:id/start?edgeGatewayReferenceID=:egid**

Starts the analytics instance with the specified ID that exists either at the
edge or at the ledger tier.

**POST /analytics-instances/:id/stop?edgeGatewayReferenceID=:egid**

Stops the analytics instance with the specified ID that exists either at the
edge or at the ledger tier.

---

##### DATA SOURCES

**POST /data-sources?edgeGatewayReferenceID=:egid**

Registers a data source either at the edge or at the ledger tier.

**DELETE /data-sources/:id?edgeGatewayReferenceID=:egid**

Unregisters the data source with the specified ID that exists either at the
edge or at the ledger tier.

**POST /data-sources/discover?edgeGatewayReferenceID=:egid**

Discovers data sources that match the specified criteria either at the edge or
at the ledger tier.

**GET /data-sources/:id/data?edgeGatewayReferenceID=:egid**

Gets data from the data source with the specified ID that exists either at the
edge or at the ledger tier.

---

##### EDGE GATEWAYS

**POST /edge-gateways**

Creates an edge gateway.

**PUT /edge-gateways/:id**

Changes the edge gateway with the specified ID.

**GET /edge-gateways/:id**

Gets the edge gateway with the specified ID.

**DELETE /edge-gateways/:id**

Deletes the edge gateway with the specified ID.

**POST /edge-gateways/discover**

Discovers edge gateways that match the specified criteria.

---
