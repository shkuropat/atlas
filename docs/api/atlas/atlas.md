# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [service_control_plane.proto](#service_control_plane.proto)
    - [ControlPlane](#atlas.ControlPlane)
  
- [service_reports_plane.proto](#service_reports_plane.proto)
    - [ReportsPlane](#atlas.ReportsPlane)
  
- [type_address.proto](#type_address.proto)
    - [Address](#atlas.Address)
  
    - [AddressType](#atlas.AddressType)
  
- [type_command.proto](#type_command.proto)
    - [Command](#atlas.Command)
  
    - [CommandType](#atlas.CommandType)
  
- [type_data_chunk.proto](#type_data_chunk.proto)
    - [DataChunk](#atlas.DataChunk)
  
    - [DataChunkType](#atlas.DataChunkType)
  
- [type_diff_task.proto](#type_diff_task.proto)
    - [DiffTask](#atlas.DiffTask)
  
- [type_digest.proto](#type_digest.proto)
    - [Digest](#atlas.Digest)
  
    - [DigestType](#atlas.DigestType)
  
- [type_domain.proto](#type_domain.proto)
    - [Domain](#atlas.Domain)
  
- [type_entity.proto](#type_entity.proto)
    - [Entity](#atlas.Entity)
  
- [type_kafka_address.proto](#type_kafka_address.proto)
    - [KafkaAddress](#atlas.KafkaAddress)
  
- [type_kafka_edpoint.proto](#type_kafka_edpoint.proto)
    - [KafkaEndpoint](#atlas.KafkaEndpoint)
  
- [type_metadata.proto](#type_metadata.proto)
    - [Metadata](#atlas.Metadata)
  
    - [MetadataType](#atlas.MetadataType)
  
- [type_metric.proto](#type_metric.proto)
    - [Metric](#atlas.Metric)
    - [Metrics](#atlas.Metrics)
  
    - [MetricType](#atlas.MetricType)
    - [MetricsType](#atlas.MetricsType)
  
- [type_report.proto](#type_report.proto)
    - [Report](#atlas.Report)
  
    - [ReportType](#atlas.ReportType)
  
- [type_report_reply.proto](#type_report_reply.proto)
    - [ReportReply](#atlas.ReportReply)
  
    - [ReportReplyType](#atlas.ReportReplyType)
  
- [type_report_request.proto](#type_report_request.proto)
    - [ReportRequest](#atlas.ReportRequest)
  
    - [ReportRequestType](#atlas.ReportRequestType)
  
- [type_s3_address.proto](#type_s3_address.proto)
    - [S3Address](#atlas.S3Address)
  
- [type_status_reply.proto](#type_status_reply.proto)
    - [StatusReply](#atlas.StatusReply)
  
    - [StatusType](#atlas.StatusType)
  
- [type_status_request_multi.proto](#type_status_request_multi.proto)
    - [StatusRequestMulti](#atlas.StatusRequestMulti)
  
    - [StatusRequestMode](#atlas.StatusRequestMode)
  
- [type_status_request.proto](#type_status_request.proto)
    - [StatusRequest](#atlas.StatusRequest)
  
- [type_user_id.proto](#type_user_id.proto)
    - [UserID](#atlas.UserID)
  
- [type_uuid.proto](#type_uuid.proto)
    - [UUID](#atlas.UUID)
  
- [Scalar Value Types](#scalar-value-types)



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
| Commands | [Command](#atlas.Command) stream | [Command](#atlas.Command) stream | Bi-directional Commands stream. Commands are sent from service to client and from client to server |
| DataChunks | [DataChunk](#atlas.DataChunk) stream | [DataChunk](#atlas.DataChunk) stream | Bi-directional Data stream. Some commands may be followed by data load. Be it logs, dumps, etc. |
| Metrics | [Metric](#atlas.Metric) stream | [Metric](#atlas.Metric) | Uni-directional Metrics stream from client to server. |
| EntityStatus | [StatusRequest](#atlas.StatusRequest) | [StatusReply](#atlas.StatusReply) | EntityStatus checks status of the entity on the server. |
| EntityStatusMultiple | [StatusRequestMulti](#atlas.StatusRequestMulti) | [StatusReply](#atlas.StatusReply) | EntityStatusMulti checks status of the multiple entities on server. |

 



<a name="service_reports_plane.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service_reports_plane.proto


 

 

 


<a name="atlas.ReportsPlane"></a>

### ReportsPlane


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Reports | [ReportRequest](#atlas.ReportRequest) stream | [ReportReply](#atlas.ReportReply) stream |  |

 



<a name="type_address.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_address.proto
Address is an abstraction over domain-specific addresses.
Represents all types of addresses in the system.


<a name="atlas.Address"></a>

### Address
Address describes general address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [AddressType](#atlas.AddressType) |  | Type of the address. |
| s3 | [S3Address](#atlas.S3Address) |  | S3 address option |
| kafka | [KafkaAddress](#atlas.KafkaAddress) |  | Kafka address option |





 


<a name="atlas.AddressType"></a>

### AddressType
AddressType represents all types of domain-specific addresses in the system

| Name | Number | Description |
| ---- | ------ | ----------- |
| ADDRESS_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| ADDRESS_S3 | 100 | S3 and MinIO address |
| ADDRESS_KAFKA | 200 | Kafka address |


 

 

 



<a name="type_command.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_command.proto
Command is a command used for commands exchange by Control Plane.


<a name="atlas.Command"></a>

### Command
Command represents a command used by Control Plane to exchange commands between server and client(s).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [Metadata](#atlas.Metadata) |  | Header of the command |
| bytes | [bytes](#bytes) |  | Optional. Any arbitrary sequence of bytes no longer than 2^32 |
| subjects | [Metadata](#atlas.Metadata) | repeated | Optional. Multiple command&#39;s subjects. |
| commands | [Command](#atlas.Command) | repeated | Optional. Recursive chain of commands |





 


<a name="atlas.CommandType"></a>

### CommandType
CommandType represents all commands in the system

| Name | Number | Description |
| ---- | ------ | ----------- |
| COMMAND_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| COMMAND_UNSPECIFIED | 100 | Unspecified |
| COMMAND_ECHO_REQUEST | 200 | Echo request expects echo reply as an answer |
| COMMAND_ECHO_REPLY | 300 | Echo reply is an answer to echo request |
| COMMAND_CONFIG_REQUEST | 400 | Request for configuration from the other party |
| COMMAND_CONFIG | 500 | Configuration |
| COMMAND_METRICS_SCHEDULE | 600 | Metrics schedule sends schedule by which metrics should be sent. |
| COMMAND_METRICS_REQUEST | 700 | Metrics request is an explicit request for metrics to be sent |
| COMMAND_METRICS | 800 | One-time metrics |
| COMMAND_DATA_SCHEDULE | 900 | Schedule to send data |
| COMMAND_DATA_REQUEST | 1000 | Explicit data request |
| COMMAND_DATA | 1100 | Data are coming |
| COMMAND_ADDRESS | 1200 | Address is coming |
| COMMAND_EXTRACT | 1300 |  |
| COMMAND_EXTRACT_EXECUTABLES | 1400 |  |


 

 

 



<a name="type_data_chunk.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_data_chunk.proto



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





 


<a name="atlas.DataChunkType"></a>

### DataChunkType
DataChunkType represents all data chunk types

| Name | Number | Description |
| ---- | ------ | ----------- |
| DATA_CHUNK_TYPE_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| DATA_CHUNK_TYPE_UNSPECIFIED | 100 | Unspecified means data chunk type is unknown |
| DATA_CHUNK_TYPE_DATA | 200 | Data chunk |


 

 

 



<a name="type_diff_task.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_diff_task.proto
DiffTask represents request for diff between two objects.


<a name="atlas.DiffTask"></a>

### DiffTask
DiffTask represents request for diff between two objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| a | [Address](#atlas.Address) |  | &#34;A&#34; party for diff |
| b | [Address](#atlas.Address) |  | &#34;B&#34; party for diff |
| meta | [Address](#atlas.Address) |  | TODO |





 

 

 

 



<a name="type_digest.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_digest.proto
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


 

 

 



<a name="type_domain.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_domain.proto
Domain represents abstract domain.


<a name="atlas.Domain"></a>

### Domain
Domain represents abstract domain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Domain name |





 

 

 

 



<a name="type_entity.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_entity.proto
Domain represents abstract domain.


<a name="atlas.Entity"></a>

### Entity
Entity represents abstract entity.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Entity name |





 

 

 

 



<a name="type_kafka_address.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_kafka_address.proto
KafkaAddress represents Kafka address within Kafka endpoint (cluster access endpoint).
Full Kafka address would be KafkaEndpoint &#43; KafkaAddress


<a name="atlas.KafkaAddress"></a>

### KafkaAddress
KafkaAddress represents Kafka address within Kafka endpoint (cluster).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| topic | [string](#string) |  | Topic within Kafka |
| partition | [int32](#int32) |  | Partition within Kafka topic |





 

 

 

 



<a name="type_kafka_edpoint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_kafka_edpoint.proto
KafkaEndpoint represents Kafka endpoint (cluster access endpoint).
Full Kafka address would be KafkaEndpoint &#43; KafkaAddress


<a name="atlas.KafkaEndpoint"></a>

### KafkaEndpoint
KafkaEndpoint represents Kafka cluster access endpoint


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| brokers | [string](#string) | repeated | Multiple brokers of Kafka&#39;s cluster |





 

 

 

 



<a name="type_metadata.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_metadata.proto
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
| user_id | [UserID](#atlas.UserID) |  | Optional. User identifier. Used to specify related user (owner, sender, etc) |
| id | [UUID](#atlas.UUID) |  | Optional. Unique identifier of the object. Used to reference this object from outside. |
| reference_id | [UUID](#atlas.UUID) |  | Optional. Unique identifier of the reference object, if necessary. Used to reference to another object. |
| ts | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Optional. Timestamp. |
| encoding | [string](#string) |  | Optional. Encoding represents encoding algo/type of the object. TODO may be we should switch to enum here |
| compression | [string](#string) |  | Optional. Compression represents compression algo/type of the object. TODO may be we should switch to enum here |
| filename | [string](#string) |  | Optional. Filename represents filename of the object. |
| url | [string](#string) |  | Optional. URL represents URL of the object. |
| address | [Address](#atlas.Address) |  | Optional. address represents external address of the object. |
| domain | [Domain](#atlas.Domain) |  | Optional. Domain represents domain of the object. |
| digest | [Digest](#atlas.Digest) |  | Optional. Digest represents digest of the object. |
| description | [string](#string) |  | Optional. Description represents string human-readable description of the object. |
| len | [int64](#int64) |  | Optional. Len represents length of the object. |
| offset | [int64](#int64) |  | Optional. Offset represents offset of the object within the stream. |
| last | [bool](#bool) |  | Optional. Last identifies last object within the stream. |





 


<a name="atlas.MetadataType"></a>

### MetadataType
MetadataType presents type of all metadata in the system

| Name | Number | Description |
| ---- | ------ | ----------- |
| METADATA_RESERVED | 0 | Due to first enum value has to be zero in proto3 |


 

 

 



<a name="type_metric.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_metric.proto
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





 


<a name="atlas.MetricType"></a>

### MetricType
MetricType represents all metric types in the system.

| Name | Number | Description |
| ---- | ------ | ----------- |
| METRIC_TYPE_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| METRIC_TYPE_UNSPECIFIED | 100 | Unspecified means we do not know its type |
| METRIC_TYPE_CPU | 200 | CPU usage metric |
| METRIC_TYPE_RAM | 300 | RAM usage metric |



<a name="atlas.MetricsType"></a>

### MetricsType
MetricsType represents areas of metrics accumulation - such as resource utilization, etc

| Name | Number | Description |
| ---- | ------ | ----------- |
| METRICS_TYPE_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| METRICS_TYPE_UNSPECIFIED | 100 | Unspecified means we do not know its type |
| METRICS_TYPE_RESOURCE_UTILIZATION | 200 | Resource utilization, such as CPU, RAM, etc |


 

 

 



<a name="type_report.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_report.proto



<a name="atlas.Report"></a>

### Report



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [Metadata](#atlas.Metadata) |  | Payload metadata of the whole data |
| bytes | [bytes](#bytes) |  | Bytes is the purpose of the whole data chunk type May contain any arbitrary sequence of bytes no longer than 2^32 |





 


<a name="atlas.ReportType"></a>

### ReportType


| Name | Number | Description |
| ---- | ------ | ----------- |
| REPORT_TYPE_RESERVED | 0 |  |
| REPORT_TYPE_UNSPECIFIED | 100 |  |


 

 

 



<a name="type_report_reply.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_report_reply.proto



<a name="atlas.ReportReply"></a>

### ReportReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ReportReplyType](#atlas.ReportReplyType) |  |  |
| ts | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp is optional |
| description | [string](#string) |  | Description is optional |
| reports | [Report](#atlas.Report) | repeated |  |





 


<a name="atlas.ReportReplyType"></a>

### ReportReplyType


| Name | Number | Description |
| ---- | ------ | ----------- |
| REPORT_REPLY_TYPE_RESERVED | 0 |  |
| REPORT_REPLY_TYPE_UNSPECIFIED | 100 |  |


 

 

 



<a name="type_report_request.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_report_request.proto



<a name="atlas.ReportRequest"></a>

### ReportRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ReportRequestType](#atlas.ReportRequestType) |  |  |
| ts | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp is optional |
| description | [string](#string) |  | Description is optional |





 


<a name="atlas.ReportRequestType"></a>

### ReportRequestType


| Name | Number | Description |
| ---- | ------ | ----------- |
| REPORT_REQUEST_TYPE_RESERVED | 0 |  |
| REPORT_REQUEST_TYPE_UNSPECIFIED | 100 |  |


 

 

 



<a name="type_s3_address.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_s3_address.proto
S3Address represents S3 and MinIO address


<a name="atlas.S3Address"></a>

### S3Address
S3Address represents S3 and MinIO address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket | [string](#string) |  | Bucket name |
| object | [string](#string) |  | Object name |





 

 

 

 



<a name="type_status_reply.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_status_reply.proto
StatusReply represents status of object(s)


<a name="atlas.StatusReply"></a>

### StatusReply
StatusReply represents status of object(s)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [StatusType](#atlas.StatusType) |  | Status of the object. |





 


<a name="atlas.StatusType"></a>

### StatusType
StatusType represents all types of statuses

| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_RESERVED | 0 | Due to first enum value has to be zero in proto3 |
| STATUS_UNSPECIFIED | 100 | Unspecified means we do not know its type |
| STATUS_NOT_FOUND | 200 | Object not found |
| STATUS_FOUND_PARTIALLY | 300 | Not all objects requested were found |
| STATUS_FOUND | 400 | Object found |
| STATUS_FOUND_ALL | 500 | All objects found |


 

 

 



<a name="type_status_request_multi.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_status_request_multi.proto



<a name="atlas.StatusRequestMulti"></a>

### StatusRequestMulti



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [Domain](#atlas.Domain) |  |  |
| mode | [StatusRequestMode](#atlas.StatusRequestMode) |  |  |
| entities | [StatusRequest](#atlas.StatusRequest) | repeated |  |





 


<a name="atlas.StatusRequestMode"></a>

### StatusRequestMode


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESERVED | 0 |  |
| UNSPECIFIED | 100 |  |
| ALL | 200 |  |
| FIRST_FOUND | 300 |  |
| FIRST_NOT_FOUND | 400 |  |


 

 

 



<a name="type_status_request.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_status_request.proto
StatusRequest represents status request of the object(s)


<a name="atlas.StatusRequest"></a>

### StatusRequest
StatusRequest represents status request of the object(s)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#atlas.Entity) |  | Entity which status is requested |
| digest | [Digest](#atlas.Digest) |  | Digest of entity to get status about |





 

 

 

 



<a name="type_user_id.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_user_id.proto
UsedID represents unique identifier of the user.
May contain any arbitrary sequence of bytes.


<a name="atlas.UserID"></a>

### UserID
Unique identifier of the user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  | Any arbitrary sequence of bytes no longer than 2^32 |





 

 

 

 



<a name="type_uuid.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## type_uuid.proto
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

