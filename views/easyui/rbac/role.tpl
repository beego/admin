{{template "../public/header.tpl"}}
<script type="text/javascript">
var statuslist = [
    {statusid:'1',name:'禁用'},
    {statusid:'2',name:'启用'}
];
var URL="/rbac/role";
$(function(){
    //角色列表
    $("#datagrid").datagrid({
        title:'角色管理',
        url:URL+"/index",
        method:'POST',
        pagination:true,
        fitColumns:true,
        striped:true,
        rownumbers:true,
        singleSelect:true,
        idField:'Id',
        columns:[[
            {field:'Id',title:'ID',width:50,align:'center'},
            {field:'Name',title:'组名',width:150,align:'center',editor:'text'},
            {field:'Remark',title:'描述',width:250,align:'center',editor:'text'},
            // {field:'create_time',title:'添加时间',width:150,align:'center',
            //     formatter:function(value,row,index){
            //         if(value) return phpjs.date("Y-m-d H:i:s",value);
            //         return value;
            //     }
            // },
            // {field:'update_time',title:'更新时间',width:150,align:'center',
            //     formatter:function(value,row,index){
            //         if(value) return phpjs.date("Y-m-d H:i:s",value);
            //         return value;
            //     }
            // },
            {field:'Status',title:'状态',width:100,align:'center',
                formatter:function(value){
                    for(var i=0; i<statuslist.length; i++){
                        if (statuslist[i].statusid == value) return statuslist[i].name;
                    }
                    return value;
                },
                editor:{
                    type:'combobox',
                    options:{
                        valueField:'statusid',
                        textField:'name',
                        data:statuslist,
                        required:true
                    }
                }
            },
            {field:'action',title:'操作',width:200,align:'center',
                formatter:function(value,row,index){
                    var c = '<a href="'+URL+'/AccessToNode?Id='+row.Id+'" target="_blank">授权</a> ';
                    var d = '<a href="'+URL+'/RoleToUserList?Id='+row.Id+'" target="_blank">用户列表</a> ';
                    return c+d;
                }
            }
        ]],
        onAfterEdit:function(index, data, changes){
            if(vac.isEmpty(changes)){
                return;
            }
            if(data.Id == undefined){
                changes.Id = 0;
            }else{
                changes.Id = data.Id;
            }
            vac.ajax(URL+'/AddAndEdit', changes, 'POST', function(r){
                if(!r.status){
                    vac.alert(r.info);
                }else{
                    $("#datagrid").datagrid("reload");
                }
            })
        },
        onDblClickRow:function(index,row){
            editrow();
        },
        onRowContextMenu:function(e, index, row){
            e.preventDefault();
            $(this).datagrid("selectRow",index);
            $('#mm').menu('show',{
                left: e.clientX,
                top: e.clientY
            });
        },
        onHeaderContextMenu:function(e, field){
            e.preventDefault();
            $('#mm1').menu('show',{
                left: e.clientX,
                top: e.clientY
            });
        }
    });
})
//新增行
function addrow(){
    var getRows = $("#datagrid").datagrid("getRows");

    //如果没有数据，则从0行开始新增
    if(vac.isEmpty(getRows)){
        var lenght = 0;
    }else{
        var lenght = getRows.length;
    }
    $("#datagrid").datagrid("appendRow",{Status:2});//插入
    $("#datagrid").datagrid("selectRow",lenght);//选中
    $("#datagrid").datagrid("beginEdit",lenght);//编辑输入
}
function editrow(){
    if(!$("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要编辑的行");
        return;
    }
    $('#datagrid').datagrid('beginEdit', vac.getindex("datagrid"));
}
function saverow(index){
    if(!$("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要保存的行");
        return;
    }
    $('#datagrid').datagrid('endEdit', vac.getindex("datagrid"));
}
//取消
function cancelrow(){
    if(! $("#datagrid").datagrid("getSelected")){
        vac.alert("请选择要取消的行");
        return;
    }
    $("#datagrid").datagrid("cancelEdit",vac.getindex("datagrid"));
}
//刷新
function reloadrow(){
    $("#datagrid").datagrid("reload");
}

//删除
function delrow(){
    $.messager.confirm('Confirm','你确定要删除?',function(r){
        if (r){
            var row = $("#datagrid").datagrid("getSelected");
            if(! row){
                vac.alert("请选择要删除的行");
                return;
            }
            vac.ajax(URL+'/DelRole', {Id:row.Id}, 'POST', function(r){
                if(r.status){
                    $("#datagrid").datagrid('reload');
                }else{
                    vac.alert(r.info);
                }
            })
        }
    });
}
</script>
<body>
<table id="datagrid" toolbar="#tb"></table>
<div id="tb" style="padding:5px;height:auto">
    <a href="#" icon='icon-add' plain="true" onclick="addrow()" class="easyui-linkbutton" >新增</a>
    <a href="#" icon='icon-edit' plain="true" onclick="editrow()" class="easyui-linkbutton" >编辑</a>
    <a href="#" icon='icon-save' plain="true" onclick="saverow()" class="easyui-linkbutton" >保存</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="delrow()" class="easyui-linkbutton" >删除</a>
    <a href="#" icon='icon-reload' plain="true" onclick="reloadrow()" class="easyui-linkbutton" >刷新</a>
</div>
<!--表格内的右键菜单-->
<div id="mm" class="easyui-menu" style="width:120px;display: none" >
    <div iconCls='icon-add' onclick="addrow()">新增</div>
    <div iconCls="icon-edit" onclick="editrow()">编辑</div>
    <div iconCls='icon-save' onclick="saverow()">保存</div>
    <div iconCls='icon-cancel' onclick="cancelrow()">取消</div>
    <div class="menu-sep"></div>
    <div iconCls='icon-cancel' onclick="delrow()">删除</div>
    <div iconCls='icon-reload' onclick="reloadrow()">刷新</div>
    <div class="menu-sep"></div>
    <div>Exit</div>
</div>
<!--表头的右键菜单-->
<div id="mm1" class="easyui-menu" style="width:120px;display: none"  >
    <div icon='icon-add' onclick="addrow()">新增</div>
</div>
</body>
</html>