{{template "../public/header.tpl"}}
<script type="text/javascript">
    var roleid = {{.roleid}};
    var grouplist=$.parseJSON({{.grouplist | stringsToJson}});
    var URL="/rbac/role"
$(function(){
    //授权列表
    $("#combobox1").combobox({
        url:URL+'/Getlist',
        valueField:'Id',
        textField:'Name',
        value:roleid,
        onSelect:function(record){
            var group_id = $("#group").combobox("getValue");
            vac.ajax(URL+"/AccessToNode",{Id:record.Id,group_id:group_id},"POST",function(data){
                        $("#treegrid").treegrid("loadData",data)
                    }
            )
        }
    });
    //加载树
    $("#treegrid").treegrid({
        'url':URL+'/AccessToNode?group_id=1&Id='+roleid,
        'idField':'Id',
        'treeField':'Title',
        'fitColumns':true,
        'singleSelect':false,
        columns:[[
            {field:'Id',title:'ID',hidden:true},
            {field:'Title',title:'显示名',width:150},
            {field:'Name',title:'应用名',width:150}
        ]],
        onLoadSuccess:function(node,data){
            //默认选中已存在的对应关系
            for(var i=0;i<data.rows.length;i++){
                if(data.rows[i].checked == 1){
                    $(this).treegrid('select',data.rows[i].Id);
                }
            }
        },
        onSelect:function(row){
            $(this).treegrid('expandAll',row.Id);
            if(row._parentId != 0){
                $(this).treegrid('select',row._parentId);
            }
            
        },
        onUnselect:function(row){
            //选中level=2的节点
            if(row.children != undefined){
                for(var i=0;i<row.children.length;i++){
                    $(this).treegrid('unselect',row.children[i].Id);
                    //选中level=3的节点
                    if(row.children[i].children != undefined){
                        for(var j=0;j<row.children[i].children.length;j++){
                            $(this).treegrid('unselect',row.children[i].children[j].Id);
                        }
                    }
                }
            }
        }
    });
    $("#group").combobox({
        "valueField":'Id',
        "textField":'Title',
        data:grouplist,
        value:1,
        onSelect:function(record){
            var roleid = $("#combobox1").combobox("getValue");
            vac.ajax(URL+"/AccessToNode",{group_id:record.Id,Id:roleid},"POST",function(data){
                        $("#treegrid").treegrid("loadData",data)
                    }
            )
        }
    });
});
    //保存授权
    function saveaccess(){
        $.messager.progress();
        var tdata = $("#treegrid").treegrid('getSelections');
        var data=new Array(tdata.length);
        for(var i=0;i<tdata.length;i++){
            data[i] = tdata[i].Id;
        }
        var roleid = $("#combobox1").combobox("getValue");
        var group_id = $("#group").combobox("getValue");
        vac.ajax(URL+'/AddAccess', {roleid:roleid,group_id:group_id,ids:data.join(",")}, 'POST', function(r){
            $.messager.alert('提示',r.info,'info');
            $.messager.progress('close');
        })
    }
</script>
<body>
<table id="treegrid" toolbar="#tbr"></table>
<div id="tbr" style="padding:5px;height:auto">
    <div style="margin-bottom:5px">
        分组：<input id="group" name="name" >
        当前组：<input id="combobox1" name="name" >
        <a href="#"  class="easyui-linkbutton" iconCls="icon-save" title="保存" plain="true" onclick="saveaccess()">保存</a>
    </div>
</div>
</body>
</html>