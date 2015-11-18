{{template "../public/header.tpl"}}
<script type="text/javascript">
    var grouplist=$.parseJSON({{.grouplist | stringsToJson}});
    var products = [
        {productid:'1',name:'Disable'},
        {productid:'2',name:'Enable'}
    ];
    var URL="/rbac/node";
    $(function(){
        $("#treegrid").treegrid({
            url:URL+"/index",
            idField:"Id",
            treeField:"Title",
            fitColumns:"true",
            columns:[[
                {field:'Title',title:'Display Name',width:150,editor:'text'},
                {field:'Id',title:'ID',width:50},
                {field:'Name',title:'Application Name',width:100,editor:'text'},
                {field:'Group__Id',title:'Group',width:80,
                    formatter:function(value){
                        for(var i=0; i<grouplist.length; i++){
                            if (grouplist[i].Id == value) return grouplist[i].Title;
                        }
                        return value;
                    }
                },
                {field:'Status',title:'Status',width:50,align:'center',
                    formatter:function(value){
                        for(var i=0; i<products.length; i++){
                            if (products[i].productid == value) return products[i].name;
                        }
                        return value;
                    },
                    editor:{
                        type:'combobox',
                        options:{
                            valueField:'productid',
                            textField:'name',
                            data:products,
                            required:true
                        }
                    }
                },
                {field:'Remark',title:'Description',width:150,editor:'text'}
            ]],
            onAfterEdit:function(c){
                if(vac.isEmpty(c)){
                    return;
                }
                vac.ajax(URL+'/AddAndEdit', c, 'POST', function(r){
                    if(!r.status){
                        vac.alert(r.info);
                    }else{
                        var group_id = $("#group").combobox("getValue");
                        vac.ajax(URL+"/index",{group_id:group_id},"POST",function(data){
                                    $("#treegrid").treegrid("loadData",data)
                                }
                        );
                    }
                })
            },
            onDblClickRow:function(index,row){
                editrow();
            },
            onContextMenu:function(e, row){
                e.preventDefault();
                $(this).treegrid('select', row.Id);
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
    $("#group").combobox({
        "valueField":'Id',
        "textField":'Title',
        data:grouplist,
        value:1,
        onSelect:function(record){
            vac.ajax(URL+"/index",{group_id:record.Id},"POST",function(data){
                        $("#treegrid").treegrid("loadData",data)
                    }
            )
        }
    });
});
    //新增行
    function addrow(){
        var Row = $("#treegrid").treegrid("getSelected");
        var _group = $("#group").combobox("getValue");
        var data = [];
        data[0] = {Id:0,Title:'',Name:'',Remark:'',Status:'1',Pid:0,Group_id:_group};
        if(!vac.isEmpty(Row)){
            data[0].Pid =Row.Id;
            $("#treegrid").treegrid("expand",Row.Id);//展开节点
            if($("#treegrid").treegrid("getLevel",Row.Id) >2){
                vac.alert("Not allowed to add");
                return false;
            }
        }
        //如果没有数据，则从0行开始新增
        $("#treegrid").treegrid("append",{
            parent: (Row?Row.Id:null),
            data:data
        });
        $("#treegrid").treegrid("select",0);//选中
        $("#treegrid").treegrid("beginEdit",0);//编辑输入
    }
    //编辑
    function editrow(){
        var row = $("#treegrid").treegrid("getSelected");
        if(!row){
            vac.alert("Please select the row you want to edit");
            return;
        }
        $("#treegrid").treegrid("beginEdit",row.Id);
    }
    //保存
    function saverow(){
        var row = $("#treegrid").treegrid("getSelected");
        if(!row){
            vac.alert("Please select the row you want to save");
            return;
        }
        $("#treegrid").treegrid("endEdit",row.Id);
    }
    //取消
    function cancelrow(){
        var row = $("#treegrid").treegrid("getSelected");
        if(!row){
            vac.alert("Please select the row you want to cancel");
            return;
        }
        $("#treegrid").treegrid("cancelEdit",row.Id);
    }
    //删除
    function delrow(){
        $.messager.confirm('Confirm','Are you sure you want to delete?',function(r){
            if (r){
                var row = $("#treegrid").treegrid("getSelected");
                if(!row){
                    vac.alert("Please select the rows to be deleted");
                    return;
                }
                vac.ajax(URL+'/DelNode', {Id:row.Id}, 'POST', function(r){
                    if(!r.status){
                        vac.alert(r.info);
                    }else{
                        $("#treegrid").treegrid("reload");
                    }
                })
            }
        });
    }
    //刷新
    function reloadrow(){
        $("#treegrid").treegrid("reload");
    }

</script>
<body>
<table id="treegrid" title="Node Manager" class="easyui-treegrid" toolbar="#tb"></table>
<div id="tb" style="padding:5px;height:auto">
    <input id="group"/>
    <a href="#" icon='icon-add' plain="true" onclick="addrow()" class="easyui-linkbutton" >Add</a>
    <a href="#" icon='icon-edit' plain="true" onclick="editrow()" class="easyui-linkbutton" >Edit</a>
    <a href="#" icon='icon-save' plain="true" onclick="saverow()" class="easyui-linkbutton" >Save</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="cancelrow()" class="easyui-linkbutton" >Cancel</a>
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
