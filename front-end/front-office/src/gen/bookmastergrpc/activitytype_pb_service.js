// package: bookmastergrpc
// file: bookmastergrpc/activitytype.proto

var bookmastergrpc_activitytype_pb = require("../bookmastergrpc/activitytype_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ActivityTypeSvc = (function () {
  function ActivityTypeSvc() {}
  ActivityTypeSvc.serviceName = "bookmastergrpc.ActivityTypeSvc";
  return ActivityTypeSvc;
}());

ActivityTypeSvc.Create = {
  methodName: "Create",
  service: ActivityTypeSvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activitytype_pb.CreateActivityTypeRequest,
  responseType: bookmastergrpc_activitytype_pb.CreateActivityTypeReply
};

ActivityTypeSvc.Exist = {
  methodName: "Exist",
  service: ActivityTypeSvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activitytype_pb.ExistActivityTypeRequest,
  responseType: bookmastergrpc_activitytype_pb.ExistActivityTypeReply
};

ActivityTypeSvc.Read = {
  methodName: "Read",
  service: ActivityTypeSvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activitytype_pb.ReadActivityTypeRequest,
  responseType: bookmastergrpc_activitytype_pb.ReadActivityTypeReply
};

ActivityTypeSvc.Delete = {
  methodName: "Delete",
  service: ActivityTypeSvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activitytype_pb.DeleteActivityTypeRequest,
  responseType: bookmastergrpc_activitytype_pb.DeleteActivityTypeReply
};

ActivityTypeSvc.Update = {
  methodName: "Update",
  service: ActivityTypeSvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activitytype_pb.UpdateActivityTypeRequest,
  responseType: bookmastergrpc_activitytype_pb.UpdateActivityTypeReply
};

ActivityTypeSvc.List = {
  methodName: "List",
  service: ActivityTypeSvc,
  requestStream: false,
  responseStream: false,
  requestType: bookmastergrpc_activitytype_pb.ListActivityTypesRequest,
  responseType: bookmastergrpc_activitytype_pb.ListActivityTypesReply
};

exports.ActivityTypeSvc = ActivityTypeSvc;

function ActivityTypeSvcClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ActivityTypeSvcClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivityTypeSvc.Create, {
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

ActivityTypeSvcClient.prototype.exist = function exist(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivityTypeSvc.Exist, {
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

ActivityTypeSvcClient.prototype.read = function read(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivityTypeSvc.Read, {
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

ActivityTypeSvcClient.prototype.delete = function pb_delete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivityTypeSvc.Delete, {
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

ActivityTypeSvcClient.prototype.update = function update(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivityTypeSvc.Update, {
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

ActivityTypeSvcClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ActivityTypeSvc.List, {
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

exports.ActivityTypeSvcClient = ActivityTypeSvcClient;

