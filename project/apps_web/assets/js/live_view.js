let socket = new WebSocket("ws://localhost:8080/live/websocket");
let Handlebars = require("handlebars/dist/handlebars");
// import Handlebars from "handlebars/dist/handlebars";
// import $ from "expose-loader?exposes=$,jQuery!jquery";
// import Handlebars from "handlebars/lib";
/*
// import EasyMDE from "easymde/dist/easymde.min.js";
// socket.onopen = function (event) {
//     // socket.send("Here's some text that the server is urgently awaiting!");
// };*/
let template = "";
let data = "";

sent = function () {
    let data = `{"person":[
        {
            "Firstname": "Yehuda",
                "Lastname": "Katz"
        },
        {
            "Firstname": "Yehuda",
                "Lastname": "Katz"
        }]}
    `
    let d = new Date();
    let n = d.getTime();
    socket.send(data);
}

document.addEventListener('DOMContentLoaded', () => {
    let baseNodeList = document.querySelectorAll('[eta-update]');
    if (baseNodeList.length > 0) {
        let baseElem = baseNodeList[0];
        let id = baseElem.getAttribute("id");
        // console.log("DOCUMENT LOAD : "+id);
        template = document.getElementById("template_"+id).innerHTML;
        // console.log(template);
    }
});
socket.onmessage = function (event) {
    data = event.data
    // console.log(data);
    etaUpdate(data)
}

function etaUpdate(data) {
    let baseNodeList = document.querySelectorAll('[eta-update]');

    if (baseNodeList.length > 0) {
        let baseElem = baseNodeList[0];
        let typeView = baseElem.getAttribute("eta-update")
        let compiled = Handlebars.compile(template);
        let obj = JSON.parse(data);
        let html = compiled(obj);
        switch (typeView) {
            case "replace":
                etaUpdateReplace(baseElem, html);
                break;
            case "append":
                etaUpdateAppend(baseElem, html)
                break;
            case "prepend":
                etaUpdatePrepend(baseElem, html)
                break;
            default:
                // console.log("ignore");
                break;
        }
    }
}

function etaUpdateReplace(elem, data){
    elem.innerHTML = data;

}
function etaUpdateAppend(elem, data){
    elem.innerHTML += data
}
function etaUpdatePrepend(elem, data){
    elem.innerHTML = data + elem.innerHTML
}
