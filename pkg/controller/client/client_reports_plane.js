
const { ReportsPlaneClient } = require('../../api/atlas/service_reports_plane_grpc_web_pb')
const { ObjectRequest } = require('../../api/atlas/object_request_pb')
const { ObjectsRequest } = require('../../api/atlas/objects_request_pb')
const { Address } = require('../../api/atlas/address_pb')
const { Domain } = require('../../api/atlas/domain_pb')
const { UUID } = require('../../api/atlas/uuid_pb')
const { Status } = require('../../api/atlas/status_pb')
const { Report } = require('../../api/atlas/report_pb')
const { callbackTask, callbackTaskFiles, callbackTaskReport, callbackTaskStatus } = require('./callbacks')
const { UuidFromString } = require('./uuid')
const grpc = {}
grpc.web = require('grpc-web')

const reportsPlaneClient = new ReportsPlaneClient('http://localhost:8080/', null, null)

    // Gets status of specified task
    function getTaskStatus(taskUUID) {
        // One object request
        const taskAddress = new Address()
        taskAddress.setUuid(taskUUID)
        const objectRequest = new ObjectRequest()
        objectRequest.setAddress(taskAddress)

        // Multi-object request
        const requestDomain = new Domain()
        requestDomain.setName("task")
        const resultDomain = new Domain()
        resultDomain.setName("status")

        const request = new ObjectsRequest()
        request.setRequestDomain(requestDomain)
        request.setResultDomain(resultDomain)
        request.addRequests(objectRequest)

        let result = false

        const call = reportsPlaneClient.objectsReport(
            request, 
            { "custom-header-1": "value1" }, 
            (err, response) => { callbackTaskStatus(err, response) }
        )

        call.on('status', function(status) { if (status.metadata) console.log(status.metadata) })

        return result
    }

    // Gets report of specified task
    function getTaskReport(taskUUID) {
        // One object request
        const taskAddress = new Address()
        taskAddress.setUuid(taskUUID)
        const objectRequest = new ObjectRequest()
        objectRequest.setAddress(taskAddress)

        // Multi-object request
        const requestDomain = new Domain()
        requestDomain.setName("task")
        const resultDomain = new Domain()
        resultDomain.setName("report")

        const request = new ObjectsRequest()
        request.setRequestDomain(requestDomain)
        request.setResultDomain(resultDomain)
        request.addRequests(objectRequest)

        let result = false

        const call = reportsPlaneClient.objectsReport(
            request,
            { "custom-header-1": "value1" },
            (err, response) => { callbackTaskReport(err, response) }
        )

        call.on('status', function(status) {
            if (status.metadata) {
                console.log("Received metadata")
                console.log(status.metadata)
            }
        })

        return result
    }

    // GetTask requests task
    function getTask(taskUUID) {
        // One object request
        const taskAddress = new Address()
        taskAddress.setUuid(taskUUID)
        const objectRequest = new ObjectRequest()
        objectRequest.setAddress(taskAddress)
  
        // Multi-object request
        const requestDomain = new Domain()
        requestDomain.setName("task")
        const resultDomain = new Domain()
        resultDomain.setName("task")
  
        const request = new ObjectsRequest()
        request.setRequestDomain(requestDomain)
        request.setResultDomain(resultDomain)
        request.addRequests(objectRequest)
  
        let result = false

        const call = reportsPlaneClient.objectsReport(
            request,
            { "custom-header-1": "value1" },
            (err, response) => { callbackTask(err, response) }
        )

    }

    // GetTaskFiles requests file(es) of the task
    function getTaskFiles(taskUUID) {
        // One object request
        const taskAddress = new Address()
        taskAddress.setUuid(taskUUID)
        const objectRequest = new ObjectRequest()
        objectRequest.setAddress(taskAddress)

        // Multi-object request
        const requestDomain = new Domain()
        requestDomain.setName("task")
        const resultDomain = new Domain()
        resultDomain.setName("file")

        const request = new ObjectsRequest()
        request.setRequestDomain(requestDomain)
        request.setResultDomain(resultDomain)
        request.addRequests(objectRequest)

        let result = false

        const call = reportsPlaneClient.objectsReport(
            request,
            { "custom-header-1": "value1" },
            (err, response) => { callbackTaskFiles(err, response) }
        )
        
    }


console.log('starting at:' + Date.now())
getTaskStatus(UuidFromString("40fa8055-4fa9-438c-9bf1-bb5639173bc4"))
//reporter.getTaskReport(taskUUID)
//reporter.getTask(taskUUID)
//reporter.getTaskFiles(taskUUID)
// Promise returned from getTaskFiles is ignored
