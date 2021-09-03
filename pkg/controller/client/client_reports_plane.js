
const {ObjectRequest} = require('../../api/atlas/object_request_pb');
const {ObjectsRequest} = require('../../api/atlas/objects_request_pb');
const {Address} = require('../../api/atlas/address_pb');
const {Domain} = require('../../api/atlas/domain_pb');
const grpc = {};
grpc.web = require('grpc-web');

// gets status of specified task
export async function GetTaskStatus(ReportsPlaneClient, taskUUID) {
    // One object request
    var taskAddress = new Address();
    taskAddress.setUuid(taskUUID);
    var objectRequest = new ObjectRequest();
    objectRequest.setAddress(taskAddress);
    // Multi-object request
    var requestDomain = new Domain();
    requestDomain.setName("task");
    var resultDomain = new Domain();
    resultDomain.setName("status");

    var request = new ObjectsRequest();
    request.setRequestDomain(requestDomain);
    request.setResultDomain(resultDomain);
    request.addRequests(objectRequest);

    /*
    let call = reportsPlaneClient.objectsReport(
        request,
        {"custom-header-1": "value1"},
        callbackGetTaskStatus,
    );
    console.log("set on function GetTaskStatus");
    call.on('status', function(status) {
        console.log('on status')
        console.log(status)
        //if (status.metadata) {
        //    console.log("Received metadata");
        //    console.log(status.metadata);
        //}
    });
    call.on('data', function(data) {
        console.log('on data')
    });
    var completed = false;
    call.on('end', function() {
        console.log('on end')
        completed = true;
    });
    call.on('error', function(error) {
        console.log('on error')
    });

     */
    try {
        const response = await ReportsPlaneClient.objectsReport(
            request,
            {"custom-header-1": "value1"},
        );
        return response
    } catch(error){
        console.log(error);
        return null
    }
}

// gets report of specified task
export async function GetTaskReport(ReportsPlaneClient, taskUUID) {
    // One object request
    var taskAddress = new Address();
    taskAddress.setUuid(taskUUID);
    var objectRequest = new ObjectRequest();
    objectRequest.setAddress(taskAddress);
    // Multi-object request
    var requestDomain = new Domain();
    requestDomain.setName("task");
    var resultDomain = new Domain();
    resultDomain.setName("report");

    var request = new ObjectsRequest();
    request.setRequestDomain(requestDomain);
    request.setResultDomain(resultDomain);
    request.addRequests(objectRequest);

    try {
        const response = await ReportsPlaneClient.objectsReport(
            request,
            {"custom-header-1": "value1"},
        );
        return response
    } catch(error){
        console.log(error);
        return null
    }
}
