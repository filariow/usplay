// package: reportgrpc
// file: reportgrpc/report.proto

var reportgrpc_report_pb = require("../reportgrpc/report_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ReportSvc = (function () {
  function ReportSvc() {}
  ReportSvc.serviceName = "reportgrpc.ReportSvc";
  return ReportSvc;
}());

ReportSvc.Create = {
  methodName: "Create",
  service: ReportSvc,
  requestStream: false,
  responseStream: false,
  requestType: reportgrpc_report_pb.CreateReportRequest,
  responseType: reportgrpc_report_pb.CreateReportReply
};

ReportSvc.Read = {
  methodName: "Read",
  service: ReportSvc,
  requestStream: false,
  responseStream: false,
  requestType: reportgrpc_report_pb.ReadReportRequest,
  responseType: reportgrpc_report_pb.ReadReportReply
};

ReportSvc.Delete = {
  methodName: "Delete",
  service: ReportSvc,
  requestStream: false,
  responseStream: false,
  requestType: reportgrpc_report_pb.DeleteReportRequest,
  responseType: reportgrpc_report_pb.DeleteReportReply
};

ReportSvc.List = {
  methodName: "List",
  service: ReportSvc,
  requestStream: false,
  responseStream: false,
  requestType: reportgrpc_report_pb.ListReportsRequest,
  responseType: reportgrpc_report_pb.ListReportsReply
};

exports.ReportSvc = ReportSvc;

function ReportSvcClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ReportSvcClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ReportSvc.Create, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ReportSvcClient.prototype.read = function read(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ReportSvc.Read, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ReportSvcClient.prototype.delete = function pb_delete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ReportSvc.Delete, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ReportSvcClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ReportSvc.List, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.ReportSvcClient = ReportSvcClient;

