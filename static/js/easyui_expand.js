/**
 * 自定义的easyui扩展
 */

//自己写的扩展easyui表单的验证
$.extend($.fn.validatebox.defaults.rules, {
    //验证汉字
    CHS: {
        validator: function (value) {
            return /^[\u0391-\uFFE5]+$/.test(value);
        },
        message: '只能输入汉字'
    },
    //移动手机号码验证
    mobile: {//value值为文本框中的值
        validator: function (value) {
            var reg = /^1[3|4|5|8|9]\d{9}$/;
            return reg.test(value);
        },
        message: '输入手机号码格式不准确.'
    },
    //国内邮编验证
    zipcode: {
        validator: function (value) {
            var reg = /^[1-9]\d{5}$/;
            return reg.test(value);
        },
        message: '邮编必须是非0开始的6位数字.'
    },
    //用户账号验证(只能包括 _ 数字 字母)
    account: {//param的值为[]中值
        validator: function (value, param) {
            if (value.length < param[0] || value.length > param[1]) {
                $.fn.validatebox.defaults.rules.account.message = '用户名长度必须在' + param[0] + '至' + param[1] + '范围';
                return false;
            } else {
                if (!/^[a-zA-Z][\w]+$/.test(value)) {
                    $.fn.validatebox.defaults.rules.account.message = '用户名只能字母开头，平且由数字、字母、下划线组成.';
                    return false;
                } else {
                    return true;
                }
            }
        }, message: ''
    },
    password:{
        validator:function(value, param){
                if (value.length < param[0] || value.length > param[1]) {
                    $.fn.validatebox.defaults.rules.password.message = '密码长度必须在' + param[0] + '至' + param[1] + '范围';
                    return false;
                } else {
                    if (!/^[\w|\!|\@|\#|\$|\%|\^|\&|\*|\~|\(|\)]+$/.test(value)) {
                        $.fn.validatebox.defaults.rules.password.message = '密码格式不正确.';
                        return false;
                    } else {
                        return true;
                    }
                }
            }, message: ''
        }
})
//时间格式化
Date.prototype.format = function(format){
    /*
     * eg:format="yyyy-MM-dd hh:mm:ss";
     */
    if(!format){
        format = "yyyy-MM-dd hh:mm:ss";
    }

    var o = {
        "M+": this.getMonth() + 1, // month
        "d+": this.getDate(), // day
        "h+": this.getHours(), // hour
        "m+": this.getMinutes(), // minute
        "s+": this.getSeconds(), // second
        "q+": Math.floor((this.getMonth() + 3) / 3), // quarter
        "S": this.getMilliseconds()
        // millisecond
    };

    if (/(y+)/.test(format)) {
        format = format.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    }

    for (var k in o) {
        if (new RegExp("(" + k + ")").test(format)) {
            format = format.replace(RegExp.$1, RegExp.$1.length == 1 ? o[k] : ("00" + o[k]).substr(("" +o[k]).length));
        }
    }
    return format;
};
//扩展datagrid的日期编辑框
$.extend($.fn.datagrid.defaults.editors, {
    datetimebox: {//datetimebox就是你要自定义editor的名称
        init: function(container, options){
            var input = $('<input class="easyuidatetimebox">').appendTo(container);
            return input.datetimebox({
                formatter:function(date){
                    return new Date(date).format("yyyy-MM-dd hh:mm:ss");
                }
            });
        },
        getValue: function(target){
            return $(target).parent().find('input.combo-value').val();
        },
        setValue: function(target, value){
            $(target).datetimebox("setValue",value);
        },
        resize: function(target, width){
            var input = $(target);
            if ($.boxModel == true){
                input.width(width - (input.outerWidth() - input.width()));
            } else {
                input.width(width);
            }
        }
    }
});
// 扩展动态添加删除editor
$.extend($.fn.datagrid.methods, {
    addEditor : function(jq, param) {
        if (param instanceof Array) {
            $.each(param, function(index, item) {
                var e = $(jq).datagrid('getColumnOption', item.field);
                e.editor = item.editor;
            });
        } else {
            var e = $(jq).datagrid('getColumnOption', param.field);
            e.editor = param.editor;
        }
    },
    removeEditor : function(jq, param) {
        if (param instanceof Array) {
            $.each(param, function(index, item) {
                var e = $(jq).datagrid('getColumnOption', item);
                e.editor = {};
            });
        } else {
            var e = $(jq).datagrid('getColumnOption', param);
            e.editor = {};
        }
    }
});
//这是自己添加的代码，为了释放tabs中引入iframe内存未释放的问题。
$.fn.panel.defaults = $.extend({},$.fn.panel.defaults,{onBeforeDestroy:function(){
	var frame=$('iframe', this);
	if(frame.length>0){
		frame[0].contentWindow.document.write('');
		frame[0].contentWindow.close();
		frame.remove();
		if($.support.leadingWhitespace){
			CollectGarbage();
		}
	}
	}
});
//扩展日历选中时间的格式化
$.fn.datebox.defaults.formatter = function(date){
    var y = date.getFullYear();
    var m = date.getMonth()+1;
    var d = date.getDate();
    return y+'-'+m+'-'+d;
}

