CREATE DATABASE atlas;

CREATE TABLE atlas.api_journal
(
  /* d is an ordering/grouping key */
  d DateTime,

  /**********************/
  /**** Call section ****/
  /**********************/

  /* endpoint_id specifies api endpoint being called (a.k.a. API function) */
  endpoint_id UInt32,

  /* source_id specifies source which is call is being journalled */
  source_id String,

  /* context_id specifies ID of the context */
  context_id String,

  /* action_id specifies action */
  action_id UInt32,

  /* duration is a nanoseconds duration since start, if applicable */
  duration Int64,

  /************************/
  /**** Object section ****/
  /************************/

  /* type specifies object type, if any */
  type UInt32,

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

  /***********************/
  /**** Error section ****/
  /***********************/

  /* error specifies error, if any */
  error String

)
ENGINE = MergeTree
ORDER BY d
;
