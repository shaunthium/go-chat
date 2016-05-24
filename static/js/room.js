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
  $('#submit-btn').click(function(e) {
    e.preventDefault();
    var message = $("#input").val();
    $("#input").val('');
    $.post("/messages", {
      message: message
    });
  });
  setInterval(fetch, 3000);
});
