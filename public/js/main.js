import Bacon from 'baconjs'
import $ from 'jquery'
import template from 'micro-template'
import _ from 'lodash'
import "js/app_materialize"

var tmpl = template.template
tmpl.variable = 'data';

var createWebSocket = function(path, param = {}) {
  return new WebSocket('ws://'+window.location.host+path+'?'+$.param(param))
}

Bacon.fromWebSocket = function(socket) {
  return Bacon.fromEventTarget(socket, "message").map((event) => JSON.parse(event.data));
}

// params
var params = {};

$("#params").children().each ((index, input) => {
  params[input.name] = input.value
})

// userlist
Bacon.fromWebSocket(createWebSocket('/user/list', {room: params["room"]}))
  .map((data)=> tmpl("userlist-tmpl", data)).assign($("#userlist"), "html");

// join
$("#join").asEventStream("submit").doAction(".preventDefault").subscribe((event) => {
  $.post('/user/add', {room: params["room"], name: $("#join [name=name]").val()})
})

// share

// secret
Bacon.fromWebSocket(createWebSocket('/share/list'))
  .map((data)=> tmpl("secretlist-tmpl", data)).assign($("#secretlist"), "html");
