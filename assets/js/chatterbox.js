/**
 * Created by ekeu on 25/01/16.
 */

$(function(){

    var socket = null;
    var msgBox = $('#client_message');
    var messages = $("#messages");

    function sendMessage(){
        if (!msgBox.val()) return false;
        if (!socket) {
            alert("Error: There is no socket connection!");
            return false;
        }
        socket.send(msgBox.val());
        msgBox.val("");
        return false;
    }

    $("#client_submit").click(sendMessage());

    if (!window["WebSocket"]){
        alert("Error: Your browser doesn't appear to support websockets.");
    } else {
        socket = new WebSocket("ws://localhost:8080/room");

        socket.onclose = function(){
            alert("Connection has been closed.");
        };

        socket.onmessage = function(e){
            messages.append($("<li class='messageboard'>").text(e.data));
        };
    }
});
