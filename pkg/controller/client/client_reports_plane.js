
const {ReportsPlaneClient} = require('../../api/atlas/service_reports_plane_grpc_web_pb');
const {ObjectRequest} = require('../../api/atlas/object_request_pb');
const {ObjectsRequest} = require('../../api/atlas/objects_request_pb');
const {Address} = require('../../api/atlas/address_pb');
const {Domain} = require('../../api/atlas/domain_pb');
const {UUID} = require('../../api/atlas/uuid_pb');
const {Status} = require('../../api/atlas/status_pb');
const {Report} = require('../../api/atlas/report_pb');
const grpc = {};
grpc.web = require('grpc-web');

var reportsPlaneClient = new ReportsPlaneClient('http://localhost:8080', null, null);

// gets status of specified task
function GetTaskStatus(ReportsPlaneClient, taskUUID) {
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

    var result = false;

    var call = reportsPlaneClient.objectsReport(
        request,
        {"custom-header-1": "value1"},
        function(err, response) {
            if (err) {
                console.log('Error code: '+err.code+' "'+err.message+'"');
            } else {
                console.log('Call completed');
                var statuses = response.getStatusesList();
                for (var i = 0; i < statuses.length; i++) {
                    console.log('get status object: ' + i);
                    var status = statuses[i];
                    console.log('status: ' + status.getStatus());
                    if (status.getStatus() == 200) {
                        console.log('success');
                        result = true;
                    } else {
                        console.log('status error');
                    }
                }
            }
        });

    call.on('status', function(status) {
        if (status.metadata) {
            console.log("Received metadata");
            console.log(status.metadata);
        }
    });

    return result;
}

// gets report of specified task
function GetTaskReport(ReportsPlaneClient, taskUUID) {
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

    var result = false;

    var call = reportsPlaneClient.objectsReport(
        request,
        {"custom-header-1": "value1"},
        function(err, response) {
            if (err) {
                console.log('Error code: '+err.code+' "'+err.message+'"');
            } else {
                console.log('Call completed');
                var reports = response.getReportsList();
                for (var i = 0; i < reports.length; i++) {
                    console.log('get report object: ' + i);
                    var report = reports[i];
                    console.log('report: ');
                    var children = report.getChildrenList()
                    for (var j = 0; j < children.length; j++) {
                        console.log('get child report object: ' + j);
                        var childReport = children[j];
                        // Extract report text from child report
                        var uint8array = childReport.getBytes_asU8()
                        var reportText = new TextDecoder("utf-8").decode(uint8array);
                        console.log('child report: ' + reportText);
                    }
                }
            }
        });

    call.on('status', function(status) {
        if (status.metadata) {
            console.log("Received metadata");
            console.log(status.metadata);
        }
    });

    return result;
}

var strTaskUUID = "89ec7e42-8290-45b2-b20b-3f106d821390";
var bytesTaskUUID = new TextEncoder("utf-8").encode(strTaskUUID);

var taskUUID = new UUID();
taskUUID.setData(bytesTaskUUID);

if (GetTaskStatus(reportsPlaneClient, null)) {
//if (GetTaskStatus(reportsPlaneClient, taskUUID)) {
    console.log("GetTaskStatus returned true");
} else {
    console.log("GetTaskStatus returned false");
}

/*
if (GetTaskReport(reportsPlaneClient, taskUUID)) {
    console.log("GetTaskReport returned true");
} else {
    console.log("GetTaskReport returned false");
}

*/