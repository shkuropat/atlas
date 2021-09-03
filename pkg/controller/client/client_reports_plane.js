
const {ReportsPlaneClient} = require('../../api/atlas/service_reports_plane_grpc_web_pb');
const {ObjectRequest} = require('../../api/atlas/object_request_pb');
const {ObjectsRequest} = require('../../api/atlas/objects_request_pb');
const {Address} = require('../../api/atlas/address_pb');
const {Domain} = require('../../api/atlas/domain_pb');
const {UUID} = require('../../api/atlas/uuid_pb');
const {Status} = require('../../api/atlas/status_pb');
const {Report} = require('../../api/atlas/report_pb');
const { UuidToString, UuidFromString } = require('./uuid')
const { ReportToString } = require('./report')
const grpc = {};
grpc.web = require('grpc-web');

const reportsPlaneClient = new ReportsPlaneClient('http://localhost:8080', null, null);

var globalObjectStatus = [];

// callbackGetTaskStatus
function callbackGetTaskStatus(err, response) {
    console.log('callbackGetTaskStatus');
    if (err) {
        console.log('Error code: '+err.code+' "'+err.message+'"');
        return
    }
    console.log('Call completed');
    const statuses = response.getObjectStatusesList();
    for (let i = 0; i < statuses.length; i++) {
        const status = statuses[i];
        const code = status.getStatus().getCode();
        const uuid = UuidToString(status.getAddress().getUuid())
        console.log(i + ' : ' +  'object status: ' + code + " : " + uuid);
    }
}

// gets status of specified task
async function GetTaskStatus(ReportsPlaneClient, taskUUID) {
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

    //let call = await reportsPlaneClient.objectsReport(
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
}

// callbackGetTaskReport
function callbackGetTaskReport(err, response) {
    console.log('callbackGetTaskReport');
    if (err) {
        console.log('Error code: '+err.code+' "'+err.message+'"');
        return;
    }
    console.log('Call completed');
    const reports = response.getReportsList();
    for (let i = 0; i < reports.length; i++) {
        console.log('get report object: ' + i);
        const report = reports[i];
        console.log('report: ');
        const children = report.getChildrenList()
        for (let j = 0; j < children.length; j++) {
            console.log('get child report object: ' + j);
            let childReport = children[j];
            // Extract report text from child report
            console.log('child report: ' + ReportToString(childReport));
        }
    }
}

// gets report of specified task
async function GetTaskReport(ReportsPlaneClient, taskUUID) {
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

    //let call = await reportsPlaneClient.objectsReport(
    await reportsPlaneClient.objectsReport(
        request,
        {"custom-header-1": "value1"},
        callbackGetTaskReport,
    );
    console.log("set on function GetTaskReport");
//    call.on('status', function(status) {
//        if (status.metadata) {
//            console.log("Received metadata");
//            console.log(status.metadata);
//        }
//    });
}

async function main(taskUUID) {
    await GetTaskStatus(reportsPlaneClient, taskUUID);
//    await GetTaskReport(reportsPlaneClient, taskUUID);
}

//taskUUID = UuidFromString("40fa8055-4fa9-438c-9bf1-bb5639173bc4");
taskUUID = null;
main(taskUUID);
