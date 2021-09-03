
const {ReportsPlanePromiseClient} = require('../../api/atlas/service_reports_plane_grpc_web_pb');
const {ObjectRequest} = require('../../api/atlas/object_request_pb');
const {ObjectsRequest} = require('../../api/atlas/objects_request_pb');
const {Address} = require('../../api/atlas/address_pb');
const {Domain} = require('../../api/atlas/domain_pb');
const {UUID} = require('../../api/atlas/uuid_pb');
const {Status} = require('../../api/atlas/status_pb');
const {Report} = require('../../api/atlas/report_pb');
const { UuidToString, UuidFromString } = require('./uuid');
const { ReportToString } = require('./report');
const { GetTaskStatus, GetTaskReport } = require('./client_reports_plane');
const grpc = {};
grpc.web = require('grpc-web');

// ProcessTaskStatus returns array of task UUIDs of completed tasks
function ProcessTaskStatus(response) {
    let tasks = [];
    if (response == null) {
        return tasks;
    }
    const statuses = response.getObjectStatusesList();
    for (let i = 0; i < statuses.length; i++) {
        const status = statuses[i];
        const code = status.getStatus().getCode();
        if (code == 200) {
            const uuid = status.getAddress().getUuid();
            const text_uuid = UuidToString(uuid);
            tasks.push(uuid);
            console.log(i + ' : ' +  'object status: ' + code + " : " + text_uuid);
        }
    }
    return tasks;
}


// ProcessTaskReport returns array of reports texts from provided response's reports list
function ProcessTaskReport(response) {
    let report_texts = [];
    if (response == null) {
        return report_texts;
    }
    const reports = response.getReportsList();
    for (let i = 0; i < reports.length; i++) {
        console.log('get report object: ' + i);
        const report = reports[i];
        console.log('report: ');
        const children = report.getChildrenList()
        for (let j = 0; j < children.length; j++) {
            console.log('get child report object: ' + j);
            const childReport = children[j];
            // Extract report text from child report
            const report_text = ReportToString(childReport);
            console.log('child report: ' + report_text);
            report_texts.push(report_text);
        }
    }
    return report_texts;
}

// main entry
async function main(taskUUID) {
    const reportsPlaneClient = new ReportsPlanePromiseClient('http://localhost:8080', null, null);
    const taskStatusResponse = await GetTaskStatus(reportsPlaneClient, taskUUID);
    const taskUUIDs = ProcessTaskStatus(taskStatusResponse);
    for (const uuid of taskUUIDs) {
        const taskReportResponse = await GetTaskReport(reportsPlaneClient, uuid);
        const texts = ProcessTaskReport(taskReportResponse);
        for (const text of texts) {
            console.log('report = ' + text);
        }
    }
}

//taskUUID = UuidFromString("40fa8055-4fa9-438c-9bf1-bb5639173bc4");
// null means ALL tasks
taskUUID = null;
main(taskUUID);
