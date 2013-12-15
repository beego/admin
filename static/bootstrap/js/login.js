var Login = function(){
	return {
		init:function(){
			function checkform()
			{
				var status = true;
				var username = jQuery("#username").val();
				var password = jQuery("#password").val();
				if(username == ""){
					status = false;
					jQuery("#usernamegroup").addClass("has-error");
					jQuery("#username-messages").text("Username is required.").removeClass("hide");
				}else{
					jQuery("#usernamegroup").removeClass("has-error").addClass("has-success");
					jQuery("#username-messages").addClass("hide");
				}
				if(password == ""){
					status = false;
					jQuery("#passwordgroup").addClass("has-error");
					jQuery("#password-messages").text("Password is required.").removeClass("hide");
				}else{
					jQuery("#passwordgroup").removeClass("has-error").addClass("has-success");
					jQuery("#password-messages").addClass("hide");
				}
				return status;
			}
			jQuery("#submit").click(function(){
				if(!checkform()){
					return false;
				}
				var username = jQuery("#username").val();
				var password = jQuery("#password").val();
				jQuery.ajax({
	    			url:"/public/login?isajax=1",
	    			type: "POST",
	    			data:{username:username,password:password},
	    			success: function(r){
					    if(r.status){
			                location.href = "/public/index"
			            }else{
			                $('#alert-error').removeClass("hide").html("<span>"+r.info+"</span>");
			            }
				}
    		});
			});
		}
	}
}()