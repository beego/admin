{{template "../public/header.tpl"}}
<!-- BEGIN PAGE LEVEL STYLES -->
<link href="/static/metronic/css/login.css" rel="stylesheet" type="text/css"/>
<!-- BEGIN PAGE LEVEL PLUGINS -->
<script src="/static/metronic/js/jquery.validate.min.js" type="text/javascript"></script>
<!-- END PAGE LEVEL PLUGINS -->
<script src="/static/metronic/js/app.js" type="text/javascript"></script>
<script src="/static/metronic/js/login.js" type="text/javascript"></script> 
<!-- BEGIN BODY -->
<body class="login">
	<!-- BEGIN LOGO -->
	<div class="logo">
<!-- 		<img src="/static/metronic/image/logo-big.png" alt="" />  -->
	</div>
	<!-- END LOGO -->
	<!-- BEGIN LOGIN -->
	<div class="content">
	<!-- BEGIN LOGIN FORM -->
		<form class="form-vertical login-form" action="/public/login?isajax=1">
			<h3 class="form-title">Beego Admin</h3>
			<div class="alert alert-error hide">
				<button class="close" data-dismiss="alert"></button>
				<span>Enter any username and password.</span>
			</div>
			<div class="control-group">
				<!--ie8, ie9 does not support html5 placeholder, so we just show field title for that-->
				<label class="control-label visible-ie8 visible-ie9">Username</label>
				<div class="controls">
					<div class="input-icon left">
						<i class="icon-user"></i>
						<input class="m-wrap placeholder-no-fix" type="text" placeholder="Username" name="username"/>
					</div>
				</div>
			</div>
			<div class="control-group">
				<label class="control-label visible-ie8 visible-ie9">Password</label>
				<div class="controls">
					<div class="input-icon left">
						<i class="icon-lock"></i>
						<input class="m-wrap placeholder-no-fix" type="password" placeholder="Password" name="password"/>
					</div>
				</div>
			</div>
			<div class="form-actions">
				<label class="checkbox">
				<input type="checkbox" name="remember" value="1"/> Remember me
				</label>
				<button type="submit" class="btn green pull-right">
				Login <i class="m-icon-swapright m-icon-white"></i>
				</button>            
			</div>
		</form>
		<!-- END LOGIN FORM --> 
	</div>
	<!-- END LOGIN -->
<!-- BEGIN COPYRIGHT -->
<div class="copyright">
	2013 &copy; Beego. Admin.
</div>
</body>
<script>
	jQuery(document).ready(function() {     
	  App.init();
	  Login.init();
	});
</script>
</html>