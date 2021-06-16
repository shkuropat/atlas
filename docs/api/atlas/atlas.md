# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [address_list.proto](#address_list.proto)
    - [AddressList](#atlas.AddressList)
  
- [address_map.proto](#address_map.proto)
    - [AddressMap](#atlas.AddressMap)
    - [AddressMap.MapEntry](#atlas.AddressMap.MapEntry)
  
- [address.proto](#address.proto)
    - [Address](#atlas.Address)
  
- [compression.proto](#compression.proto)
    - [Compression](#atlas.Compression)
  
- [data_chunk_properties.proto](#data_chunk_properties.proto)
    - [DataChunkProperties](#atlas.DataChunkProperties)
  
- [data_chunk.proto](#data_chunk.proto)
    - [DataChunk](#atlas.DataChunk)
  
- [diff_task.proto](#diff_task.proto)
    - [DiffTask](#atlas.DiffTask)
  
- [digest.proto](#digest.proto)
    - [Digest](#atlas.Digest)
  
    - [DigestType](#atlas.DigestType)
  
- [dirname.proto](#dirname.proto)
    - [Dirname](#atlas.Dirname)
  
- [domain.proto](#domain.proto)
    - [Domain](#atlas.Domain)
  
- [encoding.proto](#encoding.proto)
    - [Encoding](#atlas.Encoding)
  
- [filename.proto](#filename.proto)
    - [Filename](#atlas.Filename)
  
- [kafka_address.proto](#kafka_address.proto)
    - [KafkaAddress](#atlas.KafkaAddress)
  
- [kafka_endpoint.proto](#kafka_endpoint.proto)
    - [KafkaEndpoint](#atlas.KafkaEndpoint)
  
- [metadata.proto](#metadata.proto)
    - [Metadata](#atlas.Metadata)
  
- [metric.proto](#metric.proto)
    - [Metric](#atlas.Metric)
  
    - [MetricType](#atlas.MetricType)
  
- [metrics.proto](#metrics.proto)
    - [Metrics](#atlas.Metrics)
  
    - [MetricsType](#atlas.MetricsType)
  
- [object_request.proto](#object_request.proto)
    - [ObjectRequest](#atlas.ObjectRequest)
  
- [objects_list.proto](#objects_list.proto)
    - [ObjectsList](#atlas.ObjectsList)
  
- [objects_request.proto](#objects_request.proto)
    - [ObjectsRequest](#atlas.ObjectsRequest)
  
- [presentation_options.proto](#presentation_options.proto)
    - [PresentationOptions](#atlas.PresentationOptions)
  
- [report.proto](#report.proto)
    - [Report](#atlas.Report)
  
- [request_mode.proto](#request_mode.proto)
    - [RequestMode](#atlas.RequestMode)
  
- [s3_address.proto](#s3_address.proto)
    - [S3Address](#atlas.S3Address)
  
- [service_control_plane.proto](#service_control_plane.proto)
    - [ControlPlane](#atlas.ControlPlane)
  
- [service_reports_plane.proto](#service_reports_plane.proto)
    - [ReportsPlane](#atlas.ReportsPlane)
  
- [status.proto](#status.proto)
    - [Status](#atlas.Status)
  
- [task.proto](#task.proto)
    - [Task](#atlas.Task)
  
- [url.proto](#url.proto)
    - [URL](#atlas.URL)
  
- [user_id.proto](#user_id.proto)
    - [UserID](#atlas.UserID)
  
- [uuid.proto](#uuid.proto)
    - [UUID](#atlas.UUID)
  
- [Scalar Value Types](#scalar-value-types)



<a name="address_list.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## address_list.proto
Address is an abstraction over domain-specific addresses.
Represents all types of addresses in the system.


<a name="atlas.AddressList"></a>

### AddressList
AddressList describes list of general address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#atlas.Domain) |  | Domain where address is specified |
| addresses | [Address](#atlas.Address) | repeated |  |





 

 

 

 



<a name="address_map.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## address_map.proto
Address is an abstraction over domain-specific addresses.
Represents all types of addresses in the system.


<a name="atlas.AddressMap"></a>

### AddressMap
AddressMap describes map of address lists


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#atlas.Domain) |  | Domain where address is specified |
| map | [AddressMap.MapEntry](#atlas.AddressMap.MapEntry) | repeated |  |






<a name="atlas.AddressMap.MapEntry"></a>

### AddressMap.MapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [AddressList](#atlas.AddressList) |  |  |





 

 

 

 



<a name="address.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## address.proto
Address is an abstraction over domain-specific addresses.
Represents all types of addresses in the system.


<a name="atlas.Address"></a>

### Address
Address describes general address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| explicit_domain | [Domain](#atlas.Domain) |  | Domain where address is specified |
| s3 | [S3Address](#atlas.S3Address) |  | S3 address option |
| kafka | [KafkaAddress](#atlas.KafkaAddress) |  | Kafka address option |
| digest | [Digest](#atlas.Digest) |  | Digest-based address option |
| uuid | [UUID](#atlas.UUID) |  | UUID-based address option |
| user_id | [UserID](#atlas.UserID) |  | UserID-based address. Used to specify any related user (owner, sender, etc) |
| dirname | [Dirname](#atlas.Dirname) |  | Dirname/path-based address |
| filename | [Filename](#atlas.Filename) |  | Filename/filepath-based address |
| url | [URL](#atlas.URL) |  | URL address |
| domain | [Domain](#atlas.Domain) |  | Domain address |
| custom_string | [string](#string) |  | Custom string |





 

 

 

 



<a name="compression.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## compression.proto



<a name="atlas.Compression"></a>

### Compression
Compression describes compression of the object


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [int32](#int32) |  |  |





 

 

 

 



<a name="data_chunk_properties.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## data_chunk_properties.proto
Metadata defines all possible metadata for objects.
For example, for stream of dataChunks, is used to represent:
  - chunk header,
  - data descriptions,
  - data encoding descriptions,
  - etc.
Can also be used as metadata description for any other objects.
Since protobuf has ability not to send &#39;optional&#39; fields, metadata can have full set of fields, describing
all possible metadata options and have only few of them which are really used sent to the correspondent.


<a name="atlas.DataChunkProperties"></a>

### DataChunkProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| digest | [Digest](#atlas.Digest) |  | Optional. Digest represents digest of the object. |
| len | [int64](#int64) |  | Optional. Len represents length of the object. |
| offset | [int64](#int64) |  | Optional. Offset represents offset of the object within the stream. |
| last | [bool](#bool) |  | Optional. Last identifies last object within the stream. |





 

 

 

 



<a name="data_chunk.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## data_chunk.proto



<a name="atlas.DataChunk"></a>

### DataChunk
DataChunk is a chunk of data transferred as a single piece.
Can be part of bigger data, transferred by smaller chunks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Metadata](#atlas.Metadata) |  | Header describes this chunk |
| data | [bytes](#bytes) |  | Data is the purpose of the whole data chunk type. May contain any arbitrary sequence of bytes no longer than 2^32. |
| transport_metadata | [Metadata](#atlas.Metadata) |  | Optional. Transport metadata describes transport-level. |
| payload_metadata | [Metadata](#atlas.Metadata) |  | Optional. Payload metadata of the whole data |





 

 

 

 



<a name="diff_task.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## diff_task.proto
DiffTask represents request for diff between two objects.


<a name="atlas.DiffTask"></a>

### DiffTask
DiffTask represents request for diff between two objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| a | [Address](#atlas.Address) |  | &#34;A&#34; party for diff |
| b | [Address](#atlas.Address) |  | &#34;B&#34; party for diff |
| meta | [Address](#atlas.Address) |  | TODO |





 

 

 

 



<a name="digest.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## digest.proto
Digest represents abstract digest of multiple types.


<a name="atlas.Digest"></a>

### Digest
Digest represents abstract digest.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [DigestType](#atlas.DigestType) |  | Type of the digest. MD5 or SHA256 or something else |
| data | [bytes](#bytes) |  | Data. Any arbitrary sequence of bytes no longer than 2^32 |





 


<a name="atlas.DigestType"></a>

### DigestType
DigestType represents all types of digests in the system.

| Name | Number | Description |
| ---- | ------ | ----------- |
| DIGEST_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| DIGEST_UNSPECIFIED | 100 | Unspecified means we do not know its type |
| DIGEST_MD5 | 200 | MD5 digest |
| DIGEST_SHA256 | 300 | SHA256 digest |


 

 

 



<a name="dirname.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dirname.proto
Domain represents abstract domain.


<a name="atlas.Dirname"></a>

### Dirname
Dirname represents abstract directory name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dirname | [string](#string) |  | Dirname |





 

 

 

 



<a name="domain.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## domain.proto
Domain represents abstract domain.


<a name="atlas.Domain"></a>

### Domain
Domain represents abstract domain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Domain name |





 

 

 

 



<a name="encoding.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## encoding.proto



<a name="atlas.Encoding"></a>

### Encoding
Encoding describes encoding of the object


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method | [string](#string) |  |  |





 

 

 

 



<a name="filename.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## filename.proto
Domain represents abstract domain.


<a name="atlas.Filename"></a>

### Filename
Filename represents abstract filename.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filename | [string](#string) |  | Filename |





 

 

 

 



<a name="kafka_address.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## kafka_address.proto
KafkaAddress represents Kafka address within Kafka endpoint (cluster access endpoint).
Full Kafka address would be KafkaEndpoint &#43; KafkaAddress


<a name="atlas.KafkaAddress"></a>

### KafkaAddress
KafkaAddress represents Kafka address within Kafka endpoint (cluster).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| topic | [string](#string) |  | Topic within Kafka |
| partition | [int32](#int32) |  | Partition within Kafka topic |





 

 

 

 



<a name="kafka_endpoint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## kafka_endpoint.proto
KafkaEndpoint represents Kafka endpoint (cluster access endpoint).
Full Kafka address would be KafkaEndpoint &#43; KafkaAddress


<a name="atlas.KafkaEndpoint"></a>

### KafkaEndpoint
KafkaEndpoint represents Kafka cluster access endpoint


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| brokers | [string](#string) | repeated | Multiple brokers of Kafka&#39;s cluster |





 

 

 

 



<a name="metadata.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## metadata.proto
Metadata defines all possible metadata for objects.
For example, for stream of dataChunks, is used to represent:
  - chunk header,
  - data descriptions,
  - data encoding descriptions,
  - etc.
Can also be used as metadata description for any other objects.
Since protobuf has ability not to send &#39;optional&#39; fields, metadata can have full set of fields, describing
all possible metadata options and have only few of them which are really used sent to the correspondent.


<a name="atlas.Metadata"></a>

### Metadata
Metadata describes metadata of the object


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [int32](#int32) |  | Optional. Type of the object. Object has to have either type or name, one of the is mandatory. Object can be typed/identified either by type or by name. |
| name | [string](#string) |  | Optional. Name of the object. Object has to have either type or name, one of the is mandatory. Object can be typed/identified either by type or by name. |
| version | [int32](#int32) |  | Optional. Version of the object. |
| description | [string](#string) |  | Optional. Description represents string human-readable description of the object. |
| status | [int32](#int32) |  | Optional. Status represents status code of the object. |
| mode | [int32](#int32) |  | Optional. Mode represents mode code of the object. |
| ts | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Optional. Timestamp of the object. |
| addresses | [AddressMap](#atlas.AddressMap) |  | Optional. Addresses represents map of address of the object. Internal and external or whatever may be needed. |
| presentation_options | [PresentationOptions](#atlas.PresentationOptions) |  | Optional. Presentation options of the object. |
| data_chunk_properties | [DataChunkProperties](#atlas.DataChunkProperties) |  | Optional. Data chunk -level properties ob the object. |





 

 

 

 



<a name="metric.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## metric.proto
Metrics come from client to server in order to track client&#39;s activities.


<a name="atlas.Metric"></a>

### Metric
Metric is one metric tuple
(metric identifier [optional], timestamp [optional], description [optional], metric value)
Metric can be classified either by type or by name.
Identifier can be omitted, in case either both sides understand what they are talking about or
identifier is provided in parent&#39;s (outer) data type, such as Metric


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [MetricType](#atlas.MetricType) |  | Optional. Type of the metric |
| name | [string](#string) |  | Optional. Name of the metric |
| ts | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Optional. Timestamp. |
| description | [string](#string) |  | Optional. Description represents human-readable description on what it is. |
| string_value | [string](#string) |  | OneOf. string value of the metric |
| double_value | [double](#double) |  | OneOf. double value of the metric |
| int32_value | [int32](#int32) |  | OneOf. int32 value of the metric |
| uint32_value | [uint32](#uint32) |  | OneOf. uint32 value of the metric |
| int64_value | [int64](#int64) |  | OneOf. int64 value of the metric |
| uint64_value | [uint64](#uint64) |  | OneOf. uint64 value of the metric |
| bytes_value | [bytes](#bytes) |  | OneOf. bytes value of the metric |





 


<a name="atlas.MetricType"></a>

### MetricType
MetricType represents all metric types in the system.

| Name | Number | Description |
| ---- | ------ | ----------- |
| METRIC_TYPE_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| METRIC_TYPE_UNSPECIFIED | 100 | Unspecified means we do not know its type |
| METRIC_TYPE_CPU | 200 | CPU usage metric |
| METRIC_TYPE_RAM | 300 | RAM usage metric |


 

 

 



<a name="metrics.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## metrics.proto
Metrics come from client to server in order to track client&#39;s activities.


<a name="atlas.Metrics"></a>

### Metrics
Metrics is a set of Metric tuples.
Metric can be classified either by type or by name.
MetricType can be specified once for the whole set of metrics, instead of specifying in each one of them.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Metadata](#atlas.Metadata) |  | Header represents metadata of the set of metrics |
| type | [int32](#int32) |  | Optional. Type of metrics set. Object can be typed/identified either by type or by name. |
| name | [string](#string) |  | Optional. Name of metrics set. Object can be typed/identified either by type or by name. |
| metrics | [Metric](#atlas.Metric) | repeated | Metrics is the purpose of the whole Metrics data type, is expected to be present at all time. |





 


<a name="atlas.MetricsType"></a>

### MetricsType
MetricsType represents areas of metrics accumulation - such as resource utilization, etc

| Name | Number | Description |
| ---- | ------ | ----------- |
| METRICS_TYPE_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| METRICS_TYPE_UNSPECIFIED | 100 | Unspecified means we do not know its type |
| METRICS_TYPE_RESOURCE_UTILIZATION | 200 | Resource utilization, such as CPU, RAM, etc |


 

 

 



<a name="object_request.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## object_request.proto
StatusRequest represents status request of the object(s)


<a name="atlas.ObjectRequest"></a>

### ObjectRequest
ObjectRequest represents request for the object(s)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Metadata](#atlas.Metadata) |  |  |





 

 

 

 



<a name="objects_list.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## objects_list.proto



<a name="atlas.ObjectsList"></a>

### ObjectsList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Metadata](#atlas.Metadata) |  |  |
| reports | [Report](#atlas.Report) | repeated |  |
| tasks | [Task](#atlas.Task) | repeated |  |
| statuses | [Status](#atlas.Status) | repeated |  |





 

 

 

 



<a name="objects_request.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## objects_request.proto



<a name="atlas.ObjectsRequest"></a>

### ObjectsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#atlas.Domain) |  |  |
| request_mode | [RequestMode](#atlas.RequestMode) |  |  |
| requests | [ObjectRequest](#atlas.ObjectRequest) | repeated |  |





 

 

 

 



<a name="presentation_options.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## presentation_options.proto



<a name="atlas.PresentationOptions"></a>

### PresentationOptions
PresentationOptions describes presentation options of the object


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| encoding | [Encoding](#atlas.Encoding) |  | Optional. Encoding represents encoding of the object. |
| compression | [Compression](#atlas.Compression) |  | Optional. Compression represents compression of the object. |
| digest | [Digest](#atlas.Digest) |  | Optional. Digest represents digest of the object. |





 

 

 

 



<a name="report.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## report.proto



<a name="atlas.Report"></a>

### Report



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Metadata](#atlas.Metadata) |  |  |
| bytes | [bytes](#bytes) |  | Bytes is the purpose of the whole report type May contain any arbitrary sequence of bytes no longer than 2^32 |
| children | [Report](#atlas.Report) | repeated | Report may contain nested reports - parts of combined report |





 

 

 

 



<a name="request_mode.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## request_mode.proto


 


<a name="atlas.RequestMode"></a>

### RequestMode


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESERVED | 0 |  |
| UNSPECIFIED | 100 |  |
| ALL | 200 |  |
| ANY | 300 |  |


 

 

 



<a name="s3_address.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## s3_address.proto
S3Address represents S3 and MinIO address


<a name="atlas.S3Address"></a>

### S3Address
S3Address represents S3 and MinIO address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket | [string](#string) |  | Bucket name |
| object | [string](#string) |  | Object name |





 

 

 

 



<a name="service_control_plane.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service_control_plane.proto
Control Plane.

Control Plane represents communication between server and client(s).
Communication is logically structured into the following areas:

  - Commands. Commands flow into both directions, thus server can send commands to clients and clients can send
    commands to server. Commands is usually a long-live connection, in order for clients and server
    to receive commands with minimal delay. However, nothing prevents clients to connect to server
    from time to time as necessary and exchange commands.

  - DataChunks. DataChunks flow into both directions, thus server can send data stream to clients and clients
    can send data stream to Server. Some commands can be followed by bulk data steamed from client and/or server.

  - Metrics. Metrics are sent from client to server in order for server to keep track of client&#39;s activities.

  - Status. Status requests are sent from client to server in order to check status of the entity(es). Status calls
    are usually used for &#34;check status and send updates&#34; logic.

 

 

 


<a name="atlas.ControlPlane"></a>

### ControlPlane


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Tasks | [Task](#atlas.Task) stream | [Task](#atlas.Task) stream | Bi-directional Commands stream. Commands are sent from service to client and from client to server |
| DataChunks | [DataChunk](#atlas.DataChunk) stream | [DataChunk](#atlas.DataChunk) stream | Bi-directional Data stream. Some commands may be followed by data load. Be it logs, dumps, etc. |
| UploadObject | [DataChunk](#atlas.DataChunk) stream | [Status](#atlas.Status) |  |
| UploadObjects | [DataChunk](#atlas.DataChunk) stream | [ObjectsList](#atlas.ObjectsList) |  |
| Metrics | [Metric](#atlas.Metric) stream | [Metric](#atlas.Metric) | Uni-directional Metrics stream from client to server. |

 



<a name="service_reports_plane.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service_reports_plane.proto


 

 

 


<a name="atlas.ReportsPlane"></a>

### ReportsPlane


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ObjectsReport | [ObjectsRequest](#atlas.ObjectsRequest) | [ObjectsList](#atlas.ObjectsList) |  |

 



<a name="status.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## status.proto
StatusRequest represents status request of the object(s)


<a name="atlas.Status"></a>

### Status
Status represents status of the object(s)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Metadata](#atlas.Metadata) |  |  |





 

 

 

 



<a name="task.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## task.proto



<a name="atlas.Task"></a>

### Task



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Metadata](#atlas.Metadata) |  | Header of the task |
| bytes | [bytes](#bytes) |  | Optional. Any arbitrary sequence of bytes no longer than 2^32 |
| subjects | [Metadata](#atlas.Metadata) | repeated | Optional. Multiple task&#39;s subjects. |
| parents | [Task](#atlas.Task) | repeated | Optional. Recursive chain of tasks |
| children | [Task](#atlas.Task) | repeated | Optional. Recursive chain of tasks |





 

 

 

 



<a name="url.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## url.proto
Domain represents abstract domain.


<a name="atlas.URL"></a>

### URL
URL represents abstract url.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | URL |





 

 

 

 



<a name="user_id.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user_id.proto
UsedID represents unique identifier of the user.
May contain any arbitrary sequence of bytes.


<a name="atlas.UserID"></a>

### UserID
Unique identifier of the user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  | Any arbitrary sequence of bytes no longer than 2^32 |





 

 

 

 



<a name="uuid.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## uuid.proto
UUID represents unique identifier.
May contain any arbitrary sequence of bytes.


<a name="atlas.UUID"></a>

### UUID
UUID represents unique identifier.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  | Any arbitrary sequence of bytes no longer than 2^32 |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

