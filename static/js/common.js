var vac = vac||{};

vac.alert = function(msg){
    $.messager.alert('提示信息',msg,'error');
}

//ajax 请求，不需要确认
vac.ajax = function(url,data,type,callback)
{
	if(url == undefined || url == null)
	{
		vac.alert('链接地址不为空');
	}

	data = (data === undefined || data == null) ? {} : data;

	type = (type === undefined || type == null) ? 'POST' : type;
	$.ajax({
		url: url,
		type: type,
		data: data,
		dataType: 'json',
        async: false,
		timeout: 180000,
		error: function(request, type, ex){
			vac.alert('验证超时');
		},
		success: function(data){

			if(callback === undefined || callback == null)
			{
				if (!data.status)
				{
					vac.alert(data.info);
				}
				else
				{
					location.reload();
				}
			}
			else
			{
				callback(data);
			}
   		}
   	});
}
vac.isEmpty = function(obj){
    var isEmpty;
    for (var n in obj) {
        if( n) {
            isEmpty = false;
            break;
        }
    }
    if(obj == null || obj.length == 0){
        isEmpty = true;
    }
    if ( isEmpty == null ) {
    isEmpty = true;
    }
    return isEmpty;
}

vac.getindex = function (str){
    var index = $("#"+str+"").datagrid("getRowIndex",$("#"+str+"").datagrid("getSelected"));
    return index;
}
//获取鼠标当前坐标
vac.mousPos = function (e){
    var x,y;
    var e = e||window.event;
    return {
        x:e.clientX+document.body.scrollLeft+document.documentElement.scrollLeft,
        y:e.clientY+document.body.scrollTop+document.documentElement.scrollTop
    };
}
/**
 * @author 孙宇
 *
 * @requires jQuery
 *
 * 将form表单元素的值序列化成对象
 *
 * @returns object
 */
vac.serializeObject = function(form) {
    var o = {};
    $.each(form.serializeArray(), function(index) {
        if (o[this['name']]) {
            o[this['name']] = o[this['name']] + "," + this['value'];
        } else {
            o[this['name']] = this['value'];
        }
    });
    return o;
};




