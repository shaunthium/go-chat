<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Room {{.Pass}}</title>
  </head>
  <body>
    {{if .Text}}
      <span>{{.Text}}</span>
    {{end}}
    <form method="POST">
      <input type="text" name="input">
      <input type="submit" value="Enter">
    </form>
  </body>
</html>
