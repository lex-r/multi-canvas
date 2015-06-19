function ServiceClass() {

}

ServiceClass.prototype.pong = function (clientRequest) {
    var requestPong = clientRequest.requestPong;
    console.log('requestPong', requestPong);
    console.log("pong test", requestPong.text);
};

ServiceClass.prototype.createWorld = function(clientRequest) {
    var world = clientRequest.requestCreateWorld;

    ball = {
        x: world.ball.x,
        y: world.ball.y,
        dirX: world.ball.dirX,
        dirY: world.ball.dirY
    }
}

ServiceClass.prototype.monitor = function(clientRequest) {
    monitor = {
        x: clientRequest.requestMonitor.x,
        y: clientRequest.requestMonitor.y
    };

    console.log("Monitor", monitor)
}

var Service = new ServiceClass();
