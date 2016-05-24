<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.2.4/jquery.min.js" charset="utf-8"></script>
    <script src="/static/js/room.js" charset="utf-8"></script>
    <title>Room {{.Pass}}</title>
  </head>
  <body>
    <div id="chat-history">

    </div>
    <form method="POST" action="/messages">
      <input type="text" name="input">
      <input type="submit" value="Enter">
    </form>
  </body>
</html>
