$(document).ready(function() {
  var username = $('#username').text();
  var roomName = $('#roomName').text();
  $(window).unload(function() {
    // Ensure that ajax request completes before leaving
    $.ajaxSetup({
      async: false
    });
    $.post("/room/" + roomName, {
      roomName: roomName
    });
  });
  function fetch() {
    $.getJSON("/messages", {roomName: roomName}, function(data) {
      if (data == null) return;
      var result = "";
      for (var i=0; i<data.length; i++) {
        var temp = data[i];
        result +=
          "<li><b>" + temp.sender + ": " + "</b>" + temp.content + "</li>";
      }
      var chatHistory = $('#chat-history');
      chatHistory.html(result);
      // chatHistory[0].scrollTop = chatHistory[0].scrollHeight;
    })
  }
  $('#submit-btn').click(function(e) {
    e.preventDefault();
    var input = $("#input");
    var content = input.val();
    input.val('');
    $.post("/messages", {
      content: content,
      sender: username,
      roomName: roomName
    });
  });
  fetch();
  setInterval(fetch, 1000);
});
