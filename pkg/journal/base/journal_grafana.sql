-- All API endpoint calls count
SELECT
  $timeSeries as t,
  COUNT()
FROM $table
WHERE $timeFilter
  AND action_id=1
GROUP BY t
ORDER BY t

-- All API endpoint errors count
SELECT
  $timeSeries as t,
  COUNT()
FROM $table
WHERE $timeFilter
  AND action_id=7
GROUP BY t
ORDER BY t

-- DataChunks endpoint calls count
SELECT
  $timeSeries as t,
  COUNT()
FROM $table
WHERE $timeFilter
  AND endpoint_id=1
  AND action_id=1
GROUP BY t
ORDER BY t

-- DataChunks endpoint errors count
SELECT
  $timeSeries as t,
  COUNT()
FROM $table
WHERE $timeFilter
  AND endpoint_id=1
  AND action_id=7
GROUP BY t
ORDER BY t

-- Bootguard endpoint calls count
SELECT
  $timeSeries as t,
  COUNT()
FROM $table
WHERE $timeFilter
  AND endpoint_id=10
  AND action_id=1
GROUP BY t
ORDER BY t

-- Bootguard endpoint errors count
SELECT
  $timeSeries as t,
  COUNT()
FROM $table
WHERE $timeFilter
  AND endpoint_id=10
  AND action_id=7
GROUP BY t
ORDER BY t

-- Reports endpoint calls count
SELECT
  $timeSeries as t,
  COUNT()
FROM $table
WHERE $timeFilter
  AND endpoint_id=2
  AND action_id=1
GROUP BY t
ORDER BY t

-- Reports endpoint errors count
SELECT
  $timeSeries as t,
  COUNT()
FROM $table
WHERE $timeFilter
  AND endpoint_id=2
  AND action_id=7
GROUP BY t
ORDER BY t

-- All API endpoint MAX duration
SELECT
  $timeSeries as t,
  MAX(duration)
FROM $table
WHERE $timeFilter
GROUP BY t, context_id
ORDER BY t

-- DataChunks endpoint MAX duration
SELECT
  $timeSeries as t,
  MAX(duration)
FROM $table
WHERE $timeFilter
  AND endpoint_id=1
GROUP BY t, context_id
ORDER BY t

-- Reports endpoint MAX duration
SELECT
  $timeSeries as t,
  MAX(duration)
FROM $table
WHERE $timeFilter
  AND endpoint_id=2
GROUP BY t, context_id
ORDER BY t

-- Bootguard Results Table - ALL
SELECT
  $timeSeries as t,
  JSONExtractBool(data, 'BootguardInfoParsed') AS BootguardInfoParsed,
  JSONExtractBool(data, 'KMFound') AS KMFound,
  JSONExtractBool(data, 'KMSignatureVerificationOk') AS KMSignatureVerificationOk,
  JSONExtractBool(data, 'BPMFound') AS BPMFound,
  JSONExtractBool(data, 'BPMKeyHashEqKMKeyHash') AS BPMKeyHashEqKMKeyHash,
  JSONExtractBool(data, 'BPMSignatureVerificationOk') AS BPMSignatureVerificationOk,
  JSONExtractBool(data, 'BGProtectedRangesFound') AS BGProtectedRangesFound,
  JSONExtractBool(data, 'BGProtectedRangesHashVerificationOk') AS BGProtectedRangesHashVerificationOk,
  JSONExtractBool(data, 'VendorProtectedRangesFound') AS VendorProtectedRangesFound,
  JSONExtractBool(data, 'VendorProtectedRangesHashVerificationOk') AS VendorProtectedRangesHashVerificationOk,
  JSONExtractRaw(data, 'BGErrors') AS BGErrors
FROM $table
WHERE $timeFilter
  AND LENGTH(data) > 0

-- Sources count
SELECT
  $timeSeries as t,
  COUNT(DISTINCT(source_id))
FROM $table
WHERE $timeFilter
  AND LENGTH(source_id) > 0
GROUP BY t
ORDER BY t

-- TOP 5 active sources
SELECT
  COUNT() as t,
  source_id as source
FROM $table
WHERE $timeFilter
  AND LENGTH(source_id) > 0
GROUP BY source_id
ORDER BY t DESC
LIMIT 5

