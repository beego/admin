{{template "../public/header.tpl"}}
<script type="text/javascript">
    var grouplist=$.parseJSON({{.grouplist | stringsToJson}});
    var products = [
        {productid:'1',name:'禁用'},
        {productid:'2',name:'启用'}
    ];
    var URL="/rbac/node";
    $(function(){
        $("#treegrid").treegrid({
            url:URL+"/index",
            idField:"Id",
            treeField:"Title",
            fitColumns:"true",
            columns:[[
                {field:'Title',title:'显示名',width:150,editor:'text'},
                {field:'Id',title:'ID',width:50},
                {field:'Name',title:'应用名',width:100,editor:'text'},
                {field:'Group__Id',title:'分组',width:80,
                    formatter:function(value){
                        for(var i=0; i<grouplist.length; i++){
                            if (grouplist[i].Id == value) return grouplist[i].Title;
                        }
                        return value;
                    }
                },
                {field:'Status',title:'状态',width:50,align:'center',
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
                {field:'Remark',title:'描述',width:150,editor:'text'}
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
                vac.alert("不允许添加");
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
            vac.alert("请选择要编辑的行");
            return;
        }
        $("#treegrid").treegrid("beginEdit",row.Id);
    }
    //保存
    function saverow(){
        var row = $("#treegrid").treegrid("getSelected");
        if(!row){
            vac.alert("请选择要保存的行");
            return;
        }
        $("#treegrid").treegrid("endEdit",row.Id);
    }
    //取消
    function cancelrow(){
        var row = $("#treegrid").treegrid("getSelected");
        if(!row){
            vac.alert("请选择要取消的行");
            return;
        }
        $("#treegrid").treegrid("cancelEdit",row.Id);
    }
    //删除
    function delrow(){
        $.messager.confirm('Confirm','你确定要删除?',function(r){
            if (r){
                var row = $("#treegrid").treegrid("getSelected");
                if(!row){
                    vac.alert("请选择要删除的行");
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
<table id="treegrid" title="节点管理" class="easyui-treegrid" toolbar="#tb"></table>
<div id="tb" style="padding:5px;height:auto">
    <input id="group"/>
    <a href="#" icon='icon-add' plain="true" onclick="addrow()" class="easyui-linkbutton" >新增</a>
    <a href="#" icon='icon-edit' plain="true" onclick="editrow()" class="easyui-linkbutton" >编辑</a>
    <a href="#" icon='icon-save' plain="true" onclick="saverow()" class="easyui-linkbutton" >保存</a>
    <a href="#" icon='icon-cancel' plain="true" onclick="cancelrow()" class="easyui-linkbutton" >取消</a>
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
