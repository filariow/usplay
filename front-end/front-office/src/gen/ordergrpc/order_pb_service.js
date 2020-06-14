// package: ordergrpc
// file: ordergrpc/order.proto

var ordergrpc_order_pb = require("../ordergrpc/order_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var OrderSvc = (function () {
  function OrderSvc() {}
  OrderSvc.serviceName = "ordergrpc.OrderSvc";
  return OrderSvc;
}());

OrderSvc.Create = {
  methodName: "Create",
  service: OrderSvc,
  requestStream: false,
  responseStream: false,
  requestType: ordergrpc_order_pb.CreateOrderRequest,
  responseType: ordergrpc_order_pb.CreateOrderReply
};

OrderSvc.Read = {
  methodName: "Read",
  service: OrderSvc,
  requestStream: false,
  responseStream: false,
  requestType: ordergrpc_order_pb.ReadOrderRequest,
  responseType: ordergrpc_order_pb.ReadOrderReply
};

OrderSvc.Exist = {
  methodName: "Exist",
  service: OrderSvc,
  requestStream: false,
  responseStream: false,
  requestType: ordergrpc_order_pb.ExistOrderRequest,
  responseType: ordergrpc_order_pb.ExistOrderReply
};

OrderSvc.Delete = {
  methodName: "Delete",
  service: OrderSvc,
  requestStream: false,
  responseStream: false,
  requestType: ordergrpc_order_pb.DeleteOrderRequest,
  responseType: ordergrpc_order_pb.DeleteOrderReply
};

OrderSvc.Update = {
  methodName: "Update",
  service: OrderSvc,
  requestStream: false,
  responseStream: false,
  requestType: ordergrpc_order_pb.UpdateOrderRequest,
  responseType: ordergrpc_order_pb.UpdateOrderReply
};

OrderSvc.List = {
  methodName: "List",
  service: OrderSvc,
  requestStream: false,
  responseStream: false,
  requestType: ordergrpc_order_pb.ListOrdersRequest,
  responseType: ordergrpc_order_pb.ListOrdersReply
};

exports.OrderSvc = OrderSvc;

function OrderSvcClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

OrderSvcClient.prototype.create = function create(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderSvc.Create, {
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

OrderSvcClient.prototype.read = function read(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderSvc.Read, {
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

OrderSvcClient.prototype.exist = function exist(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderSvc.Exist, {
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

OrderSvcClient.prototype.delete = function pb_delete(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderSvc.Delete, {
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

OrderSvcClient.prototype.update = function update(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderSvc.Update, {
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

OrderSvcClient.prototype.list = function list(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(OrderSvc.List, {
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

exports.OrderSvcClient = OrderSvcClient;

