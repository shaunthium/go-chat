$(document).ready(function() {
  console.log('jquery');
  function fetch() {
    $.getJSON("/messages", function(data) {
      console.log('data is:', data);
      $('#chat-history').append(data);
    })
  }
  setInterval(fetch, 3000);
});
