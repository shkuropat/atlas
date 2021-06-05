CREATE DATABASE atlas;

CREATE TABLE atlas.api_journal
(
  /* d is an ordering/grouping key */
  d DateTime,

  /**********************/
  /**** Call section ****/
  /**********************/

  /* EndpointID [MANDATORY] specifies ID of the endpoint (API call handler/Task processor/etc) which produces the entry */
  endpoint_id Int32,

  /* EndpointInstanceID [OPTIONAL] specifies ID of the particular endpoint instance (ex.: process) which produces the entry */
  endpoint_instance_id String,

  /* SourceID [OPTIONAL] specifies ID of the source (possibly external) of the entry */
  source_id String,

  /* ContextID [OPTIONAL] specifies ID of the execution/rpc context associated with the entry */
  context_id String,

  /* TaskID [OPTIONAL] specifies ID of the task associated with the entry */
  task_id String,

  /* Type [MANDATORY] specifies type of the entry - what this entry is about. */
  type_id Int32,

  /* duration is a nanoseconds duration since start, if applicable */
  duration Int64,

  /************************/
  /**** Object section ****/
  /************************/

  /* type specifies object type, if any */
  type Int32,

  /* size specifies object size, if any */
  size UInt64,

  /* address specifies object address, if any */
  address String,

  /* domain specifies object domain, if any */
  domain String,

  /* name specifies object name, if any */
  name String,

  /* digest specifies object digestm if any*/
  digest String,

  /* data specifies object data, if any */
  data String,

  /************************/
  /**** Result section ****/
  /************************/

  /* Result [OPTIONAL] specifies result of the operation, specified by Type */
  result String,
  /* Result [OPTIONAL] specifies status of the operation, specified by Type */
  status String,
  /* Error [OPTIONAL] tells about error encountered during the operation, specified by Type */
  error String
)
ENGINE = MergeTree
ORDER BY d
;
