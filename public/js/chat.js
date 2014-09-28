$(document).ready(function() {
  var name = prompt("Enter username");
  if (name == null || name == "") {
    name = "Guest";
  }

  function msg(data) {
    console.log(data);

    $ul = $(".messages > ul");
    $li = $("<li/>");

    $user = $("<span class=\"user\"/>");
    $user.append(data.from);

    $msg = $("<p/>");
    $msg.append(data.msg);

    $div = $("<div class=\"message\"/>");
    $div.append($user);
    $div.append($msg);

    $li.append($div);
    $ul.append($li);
  }

  $(".messages .new form").submit(function(e) {
    e.preventDefault();

    var val = $(this).children("input")[0].value;
    if (val == "") {
      return;
    }

    var data = { from: name, msg: val };
    msg(data);
    conn.emit("client", data);

    $(this).children("input")[0].value = "";
  });

  var conn = new golem.Connection("ws://127.0.0.1:3000/ws", true);
  conn.on("open", function() {
    conn.emit("hello", { name: name });
  });

  conn.on("users", function(data) {
    $ul = $(".users > ul").empty();
    for (var i = 0; i < data.length; ++i) {
      var user = data[i];
      $li = $("<li/>");
      $li.append(user.name);
      $ul.append($li);
    }
  });

  conn.on("server", function(data) {
    msg(data);
  });
});
