CREATE TABLE api_journal
(
  /* d is an ordering/grouping key */
  d DateTime,

  /**** Call section ****/

  /* endpoint_id specifies api endpoint being called (a.k.a. API function) */
  endpoint_id UInt16,

  /* source_id specifies source which is call is being journalled */
  source_id Nullable(UUID),

  /* call_id specifies ID of this particular call */
  call_id UUID,

  /* action_id specifies action */
  action_id UInt8,

  /* duration is a nanoseconds duration since start, if applicable */
  duration Int64,

  /**** Object section ****/

  /* type specifies object type, if any */
  type UInt8,

  /* size specifies object size, if any */
  size UInt64,

  /* address specifies object address, if any */
  address String,

  /* name specifies object name, if any */
  name String,

  /* data specifies object data, if any */
  data String,

  /**** Error section ****/

  /* error specifies error, if any */
  error String

)
ENGINE = MergeTree
ORDER BY d
;
