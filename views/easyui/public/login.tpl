{{template "../public/header.tpl"}}
<script type="text/javascript">
var URL="/public"
	$(function(){
		$("#dialog").dialog({
			closable:false,
			buttons:[{
            text:'登录',
            iconCls:'icon-save',
            handler:function(){
                $("#form").form('submit',{
                    url:URL+'/login?isajax=1',
                    onSubmit:function(){
                        return $("#form").form('validate');
                    },
                    success:function(r){
                        var r = $.parseJSON( r );
                        if(r.status){
                            location.href = URL+"/index"
                        }else{
                            vac.alert(r.info);
                        }
                    }
                });
            }
        },{
            text:'重置',
            iconCls:'icon-cancel',
            handler:function(){
                $("#form").from("reset");
            }
        }]
		})
	})
</script>
<body>
<div style="text-align:center;margin:0 auto;width:350px;height:250px;" id="dialog" title="登录">
<div style="padding:20px 20px 20px 40px;" >
<form id="form" method="post">
<table >
	<tr>
		<td>用户名：</td><td><input type="text" class="easyui-validatebox" name="username"/></td>
	</tr>
	<tr>
		<td>密码：</td><td><input type="password" class="easyui-validatebox" name="password"/></td>
	</tr>
</table>
</form>
</div>
</div>
</body>
</html>