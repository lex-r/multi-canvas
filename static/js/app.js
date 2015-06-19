var ws = null;
var url = "ws://127.0.0.1:8090/ws";

var builder;

var canvas, context, ball, monitor;

window.onload = function() {
    url += window.location.search
    ws = new WebSocket(url);
    ws.binaryType = "arraybuffer";

    ws.onopen = function() {
        console.log("connected to " + url);
        init();
    };

    ws.onclose = function(e) {
        console.log("connection closed (" + e.code + ")");
    };

    ws.onmessage = function(e) {
        var data = e.data;

        clientRequest = builder.build("messages.ClientRequest");
        msg = clientRequest.decode64(e.data, 'utf8');
        var method = msg.method;

        if (typeof Service[method] == "function") {
            Service[method](msg);
        }
    };
};

function init() {
    var ProtoBuf = dcodeIO.ProtoBuf;
    builder = ProtoBuf.newBuilder();
    ProtoBuf.loadProtoFile("/proto/server_request.proto", builder);
    ProtoBuf.loadProtoFile("/proto/client_request.proto", builder);
    var ServerRequest = builder.build("messages.ServerRequest");
    var ServerRequestRegister = builder.build("messages.ServerRequestRegister");
    var canvas = document.getElementById("world");
    var register = new ServerRequestRegister(canvas.offsetWidth, canvas.offsetHeight);
    var request = new ServerRequest('register');
    request.requestRegister = register;
    var buf = request.toArrayBuffer();
    initCanvas();
    ws.send(buf);
};

function initCanvas() {
    canvas        = document.getElementById("world");
    context       = this.canvas.getContext("2d");
    canvas.width  = this.canvas.offsetWidth;
    canvas.height = this.canvas.offsetHeight;

    requestAnimationFrame(mainLoop);
}

function mainLoop() {
    requestAnimationFrame(mainLoop);
    if (ball == undefined) {
        return;
    }

    context.beginPath();
    context.fillStyle = "#ffffff";
    context.fillRect(0, 0, canvas.width, canvas.height);
    context.fill();

    if (monitor != undefined) {
        var offsetX = monitor.x;
        var offsetY = monitor.y;

        if ((ball.x >= offsetX && ball.x < canvas.width + offsetX)
            && (ball.y >= offsetY && ball.y < canvas.height + offsetY))
        {
            context.rect(ball.x - offsetX,ball.y-offsetY,10,10);
            context.stroke();
        }
    }
}