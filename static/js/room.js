$(document).ready(function() {
  var username = $('#username').text();
  var roomName = $('#roomName').text();
  function fetch() {
    $.getJSON("/messages", {roomName: roomName}, function(data) {
      if (data == null) return;
      var result = "";
      for (var i=0; i<data.length; i++) {
        var temp = data[i];
        result +=
          "<li><b>" + temp.sender + ": " + "</b>" + temp.content + "</li>";
      }
      $('#chat-history').html(result);
    })
  }
  $('#submit-btn').click(function(e) {
    e.preventDefault();
    var content = $("#input").val();
    $("#input").val('');
    $.post("/messages", {
      content: content,
      sender: username,
      roomName: roomName
    });
  });
  fetch();
  setInterval(fetch, 1000);
});
