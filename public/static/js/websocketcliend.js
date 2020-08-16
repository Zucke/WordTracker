var webSocket = new WebSocket("ws://"+window.location.host+"/ws");

function CreateWebSocket(){
    try {  
        webSocket = new WebSocket("ws://"+window.location.host+"/ws");
        webSocket.onopen = function(openEvent){
            console.log("websocket OPEN"+ JSON.stringify(openEvent, null, 4));
            document.getElementById("start-button").disabled = true;
        };
        webSocket.onmessage = function(messageEvent){
            document.getElementById("URLs").innerHTML += "<a href="+messageEvent.data+">"+messageEvent.data+"</a>";
        };
        
        webSocket.onclose = function(openEvent){
            console.log("websocket CLOSE"+ JSON.stringify(openEvent, null, 4));
            document.getElementById("start-button").disabled = false;
            }; 
    
    }
    catch (error) {
        console.error(error) 
    }

}

function checkWhiteList(){
        document.getElementById("whitelist").disabled = !document.getElementById("domains-filter").checked;
}

function sendMessage(){
    CreateWebSocket();
    setTimeout(function(){
    let url_wl = document.getElementById("whitelist").value;
    if (url_wl){
        url_wl = url_wl.split(",");
    }

    let info = {
        "to_find":document.getElementById("word").value,
        "currend_url":document.getElementById("url").value,
        "use_urls_white_list":document.getElementById("domains-filter").checked,
        "urls_white_list":url_wl,
        "use_regex":document.getElementById("regex").checked,
    }
    
    webSocket.send(JSON.stringify(info));
},1000);
}
