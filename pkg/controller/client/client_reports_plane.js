
const { ReportsPlaneClient } = require('../../api/atlas/service_reports_plane_grpc_web_pb')
const { ObjectRequest } = require('../../api/atlas/object_request_pb')
const { ObjectsRequest } = require('../../api/atlas/objects_request_pb')
const { Address } = require('../../api/atlas/address_pb')
const { Domain } = require('../../api/atlas/domain_pb')
const { UUID } = require('../../api/atlas/uuid_pb')
const { Status } = require('../../api/atlas/status_pb')
const { Report } = require('../../api/atlas/report_pb')
const { callbackTask, callbackTaskFiles, callbackTaskReport, callbackTaskStatus } = require('./callbacks')
const grpc = {}
grpc.web = require('grpc-web')

const reportsPlaneClient = new ReportsPlaneClient('http://localhost:8080/', null, null)
const strTaskUUID = "89ec7e42-8290-45b2-b20b-3f106d821390"

class ClientReportsPlane {
    constructor(uuid) {
        this.globalObjectStatus = []
        this.bytesTaskUUID = null
        this.taskUUID = new UUID()
        this.passUUID(uuid)
    }

    // Gets status of specified task
    async getTaskStatus(ReportsPlaneClient, taskUUID) {
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

        const call = await reportsPlaneClient.objectsReport(
            request, 
            { "custom-header-1": "value1" }, 
            (err, response) => { callbackTaskStatus(err, response) }
        )

        call.on('status', function(status) { if (status.metadata) console.log(status.metadata) })

        return result
    }

    // Gets report of specified task
    async getTaskReport(ReportsPlaneClient, taskUUID) {
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

        const call = await reportsPlaneClient.objectsReport(
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
    async getTask(ReportsPlaneClient, taskUUID) {
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

        const call = await reportsPlaneClient.objectsReport(
            request,
            { "custom-header-1": "value1" },
            (err, response) => { callbackTask(err, response) }
        )

    }

    // GetTaskFiles requests file(es) of the task
    async getTaskFiles(ReportsPlaneClient, taskUUID) {
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

        const call = await reportsPlaneClient.objectsReport(
            request,
            { "custom-header-1": "value1" },
            (err, response) => { callbackTaskFiles(err, response) }
        )
        
    }

    // pass uuid into taskUUID data
    passUUID(uuid) {
        this.bytesTaskUUID = new TextEncoder("utf-8").encode(uuid) //strTaskUUID
        this.taskUUID.setData(this.bytesTaskUUID)
    }

}

const clientReportsPlane = new ClientReportsPlane(strTaskUUID)
clientReportsPlane.getTaskStatus(reportsPlaneClient, null)
clientReportsPlane.getTaskReport(reportsPlaneClient, null)
clientReportsPlane.getTask(reportsPlaneClient, null)
clientReportsPlane.getTaskFiles(reportsPlaneClient, null)