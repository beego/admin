var Login = function () {
    
    return {
    	Loginsubmit:function(){
    		var username = $(".login-form input[name='username']").val();
    		var password = $(".login-form input[name='password']").val();
    		var remember = $(".login-form input[name='remember']").val();
    		jQuery.ajax({
    			url:"/public/login?isajax=1",
    			type: "POST",
    			data:{username:username,password:password,remember:remember},
    			success: function(r){
				    if(r.status){
		                location.href = "/public/index"
		            }else{
		                $('.alert-error', $('.login-form')).show();
		            }
				}
    		});
    	},
        //main function to initiate the module
        init: function () {
        	
           $('.login-form').validate({
	            errorElement: 'label', //default input error message container
	            errorClass: 'help-inline', // default input error message class
	            focusInvalid: false, // do not focus the last invalid input
	            rules: {
	                username: {
	                    required: true
	                },
	                password: {
	                    required: true
	                },
	                remember: {
	                    required: false
	                }
	            },

	            messages: {
	                username: {
	                    required: "Username is required."
	                },
	                password: {
	                    required: "Password is required."
	                }
	            },

	            invalidHandler: function (event, validator) { //display error alert on form submit   
	                $('.alert-error', $('.login-form')).show();
	            },

	            highlight: function (element) { // hightlight error inputs
	                $(element)
	                    .closest('.control-group').addClass('error'); // set error class to the control group
	            },

	            success: function (label) {
	                label.closest('.control-group').removeClass('error');
	                label.remove();
	            },

	            errorPlacement: function (error, element) {
	                error.addClass('help-small no-left-padding').insertAfter(element.closest('.input-icon'));
	            },

	            submitHandler: function (form) {
	                return Login.Loginsubmit();
	            }
	        });

	        $('.login-form input').keypress(function (e) {
	            if (e.which == 13) {
	                if ($('.login-form').validate().form()) {
	                    return Login.Loginsubmit();
	                }
	                return false;
	            }
	        });
        }

    };

}();