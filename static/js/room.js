$(document).ready(function() {
  function fetch() {
    $.getJSON("/messages", function(data) {
      var result = "";
      for (var i=data.length-1; i>=0; i--) {
        var temp = data[i];
        if (temp == '') {
          break;
        } else {
          result += "<li>" + temp + "</li>";
        }
      }
      $('#chat-history').html(result);
    })
  }
  $('#submit-btn').click(function() {
    var message = $("#input").val();
    console.log('message:', message);
    $.post("/messages", {
      message: message
    });
    console.log('posted');
  });
  setInterval(fetch, 3000);
});
