{{template "../public/header.tpl"}}

<script type="text/javascript">
var statuslist = [
    {id:'1',text:'Disable'},
    {id:'2',text:'Enable'}
];
var URL="/rbac/group";
$(function(){
    //用户列表
    $("#datagrid").datagrid({
        title:'Group List',
        url:URL+'/index',
        method:'POST',
        pagination:true,
        fitColumns:true,
        striped:true,
        rownumbers:true,
        singleSelect:true,
        idField:'id',
        pagination:true,
        pageSize:20,
        pageList:[10,20,30,50,100],
        columns:[[
            {field:'Id',title:'ID',width:50},
            {field:'Name',title:'Group Name',width:100,align:'center',editor:'text'},
            {field:'Title',title:'Display Name',width:100,align:'center',editor:'text'},

            {field:'Sort',title:'Sort',width:50,align:'center',editor:'numberbox'},
            {field:'Status',title:'Status',width:50,align:'center',
                formatter:function(value){
                    for(var i=0; i<statuslist.length; i++){
                        if (statuslist[i].id == value) return statuslist[i].text;
                    }
                    return value;
                },
                editor:{
                    type:'combobox',
                    options:{
                        valueField:'id',
                        textField:'text',
                        data:statuslist,
                        required:true
                    }
                }
            }
        ]],
        onAfterEdit:function(index, data, changes){
            if(vac.isEmpty(changes)){
                return;
            }
            changes.Id = data.Id;
            vac.ajax(URL+'/UpdateGroup', changes, 'POST', function(r){
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
    //创建添加分组窗口
    $("#dialog").dialog({
        modal:true,
        resizable:true,
        top:150,
        closed:true,
        buttons:[{
            text:'Save',
            iconCls:'icon-save',
            handler:function(){
                $("#form1").form('submit',{
                    url:URL+'/AddGroup',
                    onSubmit:function(){
                        return $("#form1").form('validate');
                    },
                    success:function(r){
                        var r = $.parseJSON( r );
                        if(r.status){
                            $("#dialog").dialog("close");
                            $("#datagrid").datagrid('reload');
                        }else{
                            vac.alert(r.info);
                        }
                    }
                });
            }
        },{
            text:'Cancel',
            iconCls:'icon-cancel',
            handler:function(){
                $("#dialog").dialog("close");
            }
        }]
    });

})

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

//添加用户弹窗
function addrow(){
    $("#dialog").dialog('open');
    $("#form1").form('clear');
}

//删除
function delrow(){
    $.messager.confirm('Confirm','Are you sure you want to delete?',function(r){
        if (r){
            var row = $("#datagrid").datagrid("getSelected");
            if(!row){
                vac.alert("Please select the rows to be deleted");
                return;
            }
            vac.ajax(URL+'/DelGroup', {Id:row.Id}, 'POST', function(r){
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
<!--Right menu within the table-->
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
<!--Header right menu-->
<div id="mm1" class="easyui-menu" style="width:120px;display: none"  >
    <div icon='icon-add' onclick="addrow()">Add</div>
</div>
<div id="dialog" title="Add Group" style="width:400px;height:300px;">
    <div style="padding:20px 20px 40px 80px;" >
        <form id="form1" method="post">
            <table>
                <tr>
                    <td>Group Name:</td>
                    <td><input name="Name" class="easyui-validatebox" required="true"/></td>
                </tr>
                <tr>
                    <td>Display Name:</td>
                    <td><input name="Title" class="easyui-validatebox" required="true"  /></td>
                </tr>
                <tr>
                    <td>Sort:</td>
                    <td><input name="Sort" class="easyui-numberbox" required="true"  /></td>
                </tr>
                <tr>
                    <td>Status:</td>
                    <td>
                        <select name="Status"  style="width:153px;" class="easyui-combobox " data-options="value:2" editable="false" required="true"  >
                            <option value="2" >Enable</option>
                            <option value="1">Disable</option>
                        </select>
                    </td>
                </tr>
            </table>
        </form>
    </div>
</div>
</body>
</html>
