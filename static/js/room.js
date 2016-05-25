$(document).ready(function() {
  function fetch() {
    $.getJSON("/messages", function(data) {
      var result = "";
      console.log('data is', data);
      for (var i=data.length-1; i>=0; i--) {
        var temp = data[i];
        if (temp.content == '') {
          break;
        } else {
          result +=
            "<li><b>" + temp.sender + ": " + "</b>" + temp.content + "</li>";
        }
      }
      $('#chat-history').html(result);
    })
  }
  $('#submit-btn').click(function(e) {
    e.preventDefault();
    var content = $("#input").val();
    $("#input").val('');
    var username = $('#username').text();
    $.post("/messages", {
      content: content,
      sender: username
    });
  });
  setInterval(fetch, 3000);
});
