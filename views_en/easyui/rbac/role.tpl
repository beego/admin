{{template "../public/header.tpl"}}
<script type="text/javascript">
var statuslist = [
    {statusid:'1',name:'Disable'},
    {statusid:'2',name:'Enable'}
];
var URL="/rbac/role";
$(function(){
    //角色列表
    $("#datagrid").datagrid({
        title:'Role Management',
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
            {field:'Name',title:'Role Name',width:150,align:'center',editor:'text'},
            {field:'Remark',title:'Descriptjion',width:250,align:'center',editor:'text'},
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
            {field:'Status',title:'Status',width:100,align:'center',
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
            {field:'action',title:'Operating',width:200,align:'center',
                formatter:function(value,row,index){
                    var c = '<a href="'+URL+'/AccessToNode?Id='+row.Id+'" target="_blank">Authorize</a> ';
                    var d = '<a href="'+URL+'/RoleToUserList?Id='+row.Id+'" target="_blank">User List</a> ';
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
        vac.alert("Please select the row you want to edit");
        return;
    }
    $('#datagrid').datagrid('beginEdit', vac.getindex("datagrid"));
}
function saverow(index){
    if(!$("#datagrid").datagrid("getSelected")){
        vac.alert("Please select the row you want to save");
        return;
    }
    $('#datagrid').datagrid('endEdit', vac.getindex("datagrid"));
}
//取消
function cancelrow(){
    if(! $("#datagrid").datagrid("getSelected")){
        vac.alert("Please select the row you want to cancel");
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
    $.messager.confirm('Confirm','Are you sure you want to delete?',function(r){
        if (r){
            var row = $("#datagrid").datagrid("getSelected");
            if(! row){
                vac.alert("Please select the rows to be deleted");
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
    <a href="#" icon='icon-add' plain="true" onclick="addrow()" class="easyui-linkbutton" >Add</a>
    <a href="#" icon='icon-edit' plain="true" onclick="editrow()" class="easyui-linkbutton" >Edit</a>
    <a href="#" icon='icon-save' plain="true" onclick="saverow()" class="easyui-linkbutton" >Save</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="delrow()" class="easyui-linkbutton" >Delete</a>
    <a href="#" icon='icon-reload' plain="true" onclick="reloadrow()" class="easyui-linkbutton" >Reload</a>
</div>
<!--表格内的右键菜单-->
<div id="mm" class="easyui-menu" style="width:120px;display: none" >
    <div iconCls='icon-add' onclick="addrow()">Add</div>
    <div iconCls="icon-edit" onclick="editrow()">Edit</div>
    <div iconCls='icon-save' onclick="saverow()">Save</div>
    <div iconCls='icon-cancel' onclick="cancelrow()">Cancel</div>
    <div class="menu-sep"></div>
    <div iconCls='icon-cancel' onclick="delrow()">Delete</div>
    <div iconCls='icon-reload' onclick="reloadrow()">Reload</div>
    <div class="menu-sep"></div>
    <div>Exit</div>
</div>
<!--表头的右键菜单-->
<div id="mm1" class="easyui-menu" style="width:120px;display: none"  >
    <div icon='icon-add' onclick="addrow()">Add</div>
</div>
</body>
</html>
