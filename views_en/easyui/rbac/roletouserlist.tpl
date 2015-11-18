{{template "../public/header.tpl"}}
<script type="text/javascript">
    var roleid = {{.roleid}};
    var URL="/rbac/role"
$(function(){
    //用户列表
    $("#combobox").combobox({
        url:URL+'/Getlist',
        valueField:'Id',
        textField:'Name',
        value:roleid,
        onSelect:function(record){
            $("#datagrid2").datagrid("reload",{Id:record.Id});
        }
    });
    //组用户列表
    $("#datagrid2").datagrid({
        url:URL+'/RoleToUserList?Id='+roleid,
        method:'get',
        fitColumns:false,
        striped:true,
        rownumbers:true,
        idField:'Id',
        columns:[[
            {field:'Id',title:'ID',width:50,align:'center'},
            {field:'Username',title:'Username',width:140,align:'center'},
            {field:'Nickname',title:'Nickname',width:140,align:'center'}
        ]],
        onLoadSuccess:function(data){
            $("#datagrid2").datagrid('unselectAll');
            //默认选中已存在的对应关系
            for(var i=0;i<data.rows.length;i++){
                if(data.rows[i].checked == 1){
                    $(this).datagrid('selectRecord',data.rows[i].Id);
                }
            }
        }
    });
});
    //全选
    function selectall(){
        $("#datagrid2").datagrid('selectAll');
    }
    //全否
    function unselectall(){
        $("#datagrid2").datagrid('unselectAll');
    }
    //保存选择
    function saveselect(){
        var rows = $("#datagrid2").datagrid('getSelections');
        if(rows == null){
            vac.alert("Select a minimum of one row");
        }
        var ids = [];
        for(var i=0; i<rows.length; i++){
            ids.push(rows[i].Id);
        }
        var id = $("#combobox").combobox('getValue');
        vac.ajax(URL+'/AddRoleToUser', {Id:id,ids:ids.join(',')}, 'POST', function(r){
            $.messager.alert('Prompt',r.info,'info');
        })
    }
</script>
<body>
<table id="datagrid2" toolbar="#tb2"></table>
<div id="tb2" style="padding:5px;height:auto">
    <div style="margin-bottom:5px">
        Current Group:<input id="combobox" name="name" >
        <a href="#"  class="easyui-linkbutton" iconCls="icon-save" plain="true" onclick="saveselect()">Save</a>
    </div>
    <div style="margin-bottom:5px">
        <a href="#"  class="easyui-linkbutton" iconCls="icon-ok" plain="true" onclick="selectall()">Select All</a>
        <a href="#"  class="easyui-linkbutton" iconCls="icon-no" plain="true" onclick="unselectall()">UnSelect All</a>
    </div>
</div>
</body>
</html>
