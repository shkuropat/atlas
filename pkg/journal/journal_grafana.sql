-- All API endpoint calls count
SELECT
  $timeSeries as t,
  count()
FROM $table
WHERE $timeFilter
  AND action_id=1
GROUP BY t
ORDER BY t

-- All API endpoint errors count
SELECT
  $timeSeries as t,
  count()
FROM $table
WHERE $timeFilter
  AND action_id=7
GROUP BY t
ORDER BY t

-- DataChunks endpoint calls count
SELECT
  $timeSeries as t,
  count()
FROM $table
WHERE $timeFilter
  AND endpoint_id=1
  AND action_id=1
GROUP BY t
ORDER BY t

-- DataChunks endpoint errors count
SELECT
  $timeSeries as t,
  count()
FROM $table
WHERE $timeFilter
  AND endpoint_id=1
  AND action_id=7
GROUP BY t
ORDER BY t

-- Reports endpoint calls count
SELECT
  $timeSeries as t,
  count()
FROM $table
WHERE $timeFilter
  AND endpoint_id=2
  AND action_id=1
GROUP BY t
ORDER BY t

-- Reports endpoint errors count
SELECT
  $timeSeries as t,
  count()
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
GROUP BY t, call_id
ORDER BY t

-- DataChunks endpoint MAX duration
SELECT
  $timeSeries as t,
  MAX(duration)
FROM $table
WHERE $timeFilter
  AND endpoint_id=1
GROUP BY t, call_id
ORDER BY t

-- Reports endpoint MAX duration
SELECT
  $timeSeries as t,
  MAX(duration)
FROM $table
WHERE $timeFilter
  AND endpoint_id=2
GROUP BY t, call_id
ORDER BY t
