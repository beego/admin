{// 加载头部公共文件 }
<include file="../../Pub/Public/header_easyui" />
<script type="text/javascript">
    var roleid = {$roleid};
$(function(){
    //用户列表
    $("#combobox").combobox({
        url:URL+'/index?from=only',
        valueField:'id',
        textField:'name',
        value:roleid,
        onSelect:function(record){
            $("#datagrid2").datagrid("reload",{id:record.id});
        }
    });
    //组用户列表
    $("#datagrid2").datagrid({
        url:URL+'/RoleToUserList?id='+roleid,
        method:'get',
        fitColumns:false,
        striped:true,
        rownumbers:true,
        idField:'id',
        columns:[[
            {field:'id',title:'ID',width:50,align:'center'},
            {field:'account',title:'用户名',width:140,align:'center'},
            {field:'nickname',title:'昵称',width:140,align:'center'}
        ]],
        onLoadSuccess:function(data){
            $("#datagrid2").datagrid('unselectAll');
            //默认选中已存在的对应关系
            for(var i=0;i<data.rows.length;i++){
                if(data.rows[i].select == 1){
                    $(this).datagrid('selectRecord',data.rows[i].id);
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
            vac.alert("最少要选中一行");
        }
        var ids = [];
        for(var i=0; i<rows.length; i++){
            ids.push(rows[i].id);
        }
        var id = $("#combobox").combobox('getValue');
        vac.ajax(URL+'/AddRoleToUser', {id:id,ids:ids.join(',')}, 'POST', function(r){
            $.messager.alert('提示',r.info,'info');
        })
    }
</script>
<body>
<table id="datagrid2" toolbar="#tb2"></table>
<div id="tb2" style="padding:5px;height:auto">
    <div style="margin-bottom:5px">
        当前组：<input id="combobox" name="name" >
        <a href="#"  class="easyui-linkbutton" iconCls="icon-save" plain="true" onclick="saveselect()">保存</a>
    </div>
    <div style="margin-bottom:5px">
        <a href="#"  class="easyui-linkbutton" iconCls="icon-ok" plain="true" onclick="selectall()">全选</a>
        <a href="#"  class="easyui-linkbutton" iconCls="icon-no" plain="true" onclick="unselectall()">全否</a>
    </div>
</div>
</body>
</html>