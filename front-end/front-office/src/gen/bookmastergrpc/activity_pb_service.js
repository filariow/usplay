// package: bookmastergrpc
// file: bookmastergrpc/activity.proto

var bookmastergrpc_activity_pb = require("../bookmastergrpc/activity_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ActivitySvc = (function () {
  function ActivitySvc() {}
  ActivitySvc.serviceName = "bookmastergrpc.ActivitySvc";
  return ActivitySvc;
}());

ActivitySvc.Create = {
  methodName: "Create",
  service: ActivitySvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activity_pb.CreateActivityRequest,
  responseType: bookmastergrpc_activity_pb.CreateActivityReply
};

ActivitySvc.Read = {
  methodName: "Read",
  service: ActivitySvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activity_pb.ReadActivityRequest,
  responseType: bookmastergrpc_activity_pb.ReadActivityReply
};

ActivitySvc.Delete = {
  methodName: "Delete",
  service: ActivitySvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activity_pb.DeleteActivityRequest,
  responseType: bookmastergrpc_activity_pb.DeleteActivityReply
};

ActivitySvc.Update = {
  methodName: "Update",
  service: ActivitySvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activity_pb.UpdateActivityRequest,
  responseType: bookmastergrpc_activity_pb.UpdateActivityReply
};

ActivitySvc.List = {
  methodName: "List",
  service: ActivitySvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activity_pb.ListActivitiesRequest,
  responseType: bookmastergrpc_activity_pb.ListActivitiesReply
};

ActivitySvc.ListInInterval = {
  methodName: "ListInInterval",
  service: ActivitySvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activity_pb.ListInIntervalActivitiesRequest,
  responseType: bookmastergrpc_activity_pb.ListActivitiesReply
};

exports.ActivitySvc = ActivitySvc;

function ActivitySvcClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ActivitySvcClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivitySvc.Create, {
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

ActivitySvcClient.prototype.read = function read(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivitySvc.Read, {
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

ActivitySvcClient.prototype.delete = function pb_delete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivitySvc.Delete, {
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

ActivitySvcClient.prototype.update = function update(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivitySvc.Update, {
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

ActivitySvcClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivitySvc.List, {
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

ActivitySvcClient.prototype.listInInterval = function listInInterval(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivitySvc.ListInInterval, {
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

exports.ActivitySvcClient = ActivitySvcClient;

