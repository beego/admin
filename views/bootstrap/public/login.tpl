{{template "../public/header.tpl"}}
<script type="text/javascript" src="/static/bootstrap/js/login.js"></script>
<body style="padding: 40px;background-color: #eee;">
	<div class="cotainer">
	<form class="form-horizontal login-form" role="form">
		<div class="form-group">
		<div class="col-lg-4 col-lg-offset-4">
			<h2 class="text-center">Beego Admin</h2>
		</div>
		</div>
		<div class="form-group">
			<div class="col-lg-4 col-lg-offset-4">
				<div id="alert-error" class="hide has-error">
				</div>
			</div>
		</div>
		<div id="usernamegroup" class="form-group">
			<div class="col-lg-4 col-lg-offset-4">
				<input id="username" name="username" type="text" class="form-control input-lg" placeholder="Username" />
				<span id="username-messages" class="help-block hide"></span>
			</div>
		</div>
		<div id="passwordgroup" class="form-group ">
			<div class="col-lg-4 col-lg-offset-4">
				<input id="password" name="password" type="password" class="form-control input-lg" placeholder="Password"  />
				<span id="password-messages" class="help-block hide"></span>
			</div>
		</div>
		<div class="form-group">
			<div class="col-lg-4 col-lg-offset-4">
				<input id="submit" type="button" class="btn btn-lg btn-block btn-primary" value="登录">
			</div>
		</div>
	</form>
	</div>
	<script type="text/javascript">
	$(function(){
		Login.init()
	})
	</script>
</body>
</html>
