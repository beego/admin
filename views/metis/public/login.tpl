<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Login Page</title>
    <meta name="msapplication-TileColor" content="#5bc0de" />
    <meta name="msapplication-TileImage" content="/static/metis/img/metis-tile.png" />
    <link rel="stylesheet" href="/static/metis/lib/bootstrap/css/bootstrap.css">
    <link rel="stylesheet" href="/static/metis/css/main.css">
    <link rel="stylesheet" href="/static/metis/lib/magic/magic.css">
  </head>
  <body class="login">
    <div class="container">
      <div class="text-center">
        <img src="/static/metis/img/logo.png" alt="Metis Logo">
      </div>
      <div class="tab-content">
        <div id="login" class="tab-pane active">
          <form action="/public/login" class="form-signin" onsubmit="return login()">
            <p class="text-muted text-center">
              Enter your username and password
            </p>
            <input id="username" name="username" type="text" placeholder="Username" required="required" class="form-control">
            <input id="password" name="password" type="password" placeholder="Password" required="required" class="form-control">
            <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
          </form>
        </div>
        <div id="forgot" class="tab-pane">
          <form action="index.html" class="form-signin">
            <p class="text-muted text-center">Enter your valid e-mail</p>
            <input type="email" placeholder="mail@domain.com" required="required" class="form-control">
            <br>
            <button class="btn btn-lg btn-danger btn-block" type="submit">Recover Password</button>
          </form>
        </div>
        <div id="signup" class="tab-pane">
          <form action="index.html" class="form-signin">
            <input type="text" placeholder="username" class="form-control">
            <input type="email" placeholder="mail@domain.com" class="form-control">
            <input type="password" placeholder="password" class="form-control">
            <button class="btn btn-lg btn-success btn-block" type="submit">Register</button>
          </form>
        </div>
      </div>
      <div class="text-center">
        <ul class="list-inline">
          <li> <a class="text-muted" href="#login" data-toggle="tab">Login</a> </li>
          <li> <a class="text-muted" href="#forgot" data-toggle="tab">Forgot Password</a> </li>
          <li> <a class="text-muted" href="#signup" data-toggle="tab">Signup</a> </li>
        </ul>
      </div>
    </div><!-- /container -->
    <script src="/static/metis/lib/jquery.min.js"></script>
    <script src="/static/metis/lib/bootstrap/js/bootstrap.js"></script>
    <script>
      $('.list-inline li > a').click(function() {
        var activeForm = $(this).attr('href') + ' > form';
        //console.log(activeForm);
        $(activeForm).addClass('magictime swap');
        //set timer to 1 seconds, after that, unload the magic animation
        setTimeout(function() {
          $(activeForm).removeClass('magictime swap');
        }, 1000);
      });
      function login () {
        var username = $("#username").val();
        var password = $("#password").val();
        $.ajax({
          url:"/public/login?isajax=1",
          type:"POST",
          data:{username:username,password:password},
          success:function(r){
            if(r.status){
              location.href = "/public/index"
            }else{
              alert(r.info)
            }
          }
        });
        return false;
      }
    </script>
  </body>
</html>