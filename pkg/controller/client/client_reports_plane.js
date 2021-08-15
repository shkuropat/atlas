
const {ReportsPlaneClient} = require('../../api/atlas/service_reports_plane_grpc_web_pb');
const {ObjectRequest} = require('../../api/atlas/object_request_pb');
const {ObjectsRequest} = require('../../api/atlas/objects_request_pb');
const {Metadata} = require('../../api/atlas/metadata_pb');
const grpc = {};
grpc.web = require('grpc-web');

var reportsPlaneClient = new ReportsPlaneClient('http://localhost:8080', null, null);

function Status(ReportsPlaneClient, meta) {
    /*
    request := atlas.NewObjectsRequest().Append(atlas.NewObjectRequest().SetHeader(meta))
    request.EnsureHeader().SetDomain(atlas.DomainTask).SetResultDomain(atlas.DomainStatus)
    list, err := ReportsPlaneClient.ObjectsReport(ctx, request)
    if len(list.GetStatuses()) > 0 {
        result.Recv.Status = list.GetStatuses()[0]
    }
    result.Error = err
     */

    // What object to get status about
    var objectRequest = new ObjectRequest()
    objectRequest.setHeader(meta)

    // Build request
    var request = new ObjectsRequest()
    var header = new Metadata()
    request.setHeader(header)
    //request.EnsureHeader().SetDomain(atlas.DomainTask).SetResultDomain(atlas.DomainStatus)

    request.addRequests(objectRequest)

    var call = reportsPlaneClient.objectsReport(
        objectRequest,
        {"custom-header-1": "value1"},
        function(err, response) {
            if (err) {
                console.log('Error code: '+err.code+' "'+err.message+'"');
            } else {
                console.log('Call completed');
            }
        });
    call.on('status', function(status) {
        if (status.metadata) {
            console.log("Received metadata");
            console.log(status.metadata);
        }
    });

}

Status(reportsPlaneClient, null);
