/**
 * combobox - jQuery EasyUI
 * 
 * Copyright (c) 2009-2013 www.jeasyui.com. All rights reserved.
 *
 * Licensed under the GPL or commercial licenses
 * To use it on other terms please contact us: info@jeasyui.com
 * http://www.gnu.org/licenses/gpl.txt
 * http://www.jeasyui.com/license_commercial.php
 * 
 * Dependencies:
 *   combo
 * 
 */
(function($){
	function findDataItem(data,key,value){
		for(var i=0; i<data.length; i++){
			var item = data[i];
			if (item[key] == value){return item}
		}
		return null;
	}
	
	/**
	 * scroll panel to display the specified item
	 */
	function scrollTo(target, value){
		var panel = $(target).combo('panel');
		var item = panel.find('div.combobox-item[value="' + value + '"]');
		if (item.length){
			if (item.position().top <= 0){
				var h = panel.scrollTop() + item.position().top;
				panel.scrollTop(h);
			} else if (item.position().top + item.outerHeight() > panel.height()){
				var h = panel.scrollTop() + item.position().top + item.outerHeight() - panel.height();
				panel.scrollTop(h);
			}
		}
	}
	
	function nav(target, dir){
		var opts = $(target).combobox('options');
		var panel = $(target).combobox('panel');
		var item = panel.children('div.combobox-item-hover');
		if (!item.length){
			item = panel.children('div.combobox-item-selected');
		}
		item.removeClass('combobox-item-hover');
		if (!item.length){
			item = panel.children('div.combobox-item:visible:' + (dir=='next'?'first':'last'));
		} else {
			if (dir == 'next'){
				item = item.nextAll('div.combobox-item:visible:first');
				if (!item.length){
					item = panel.children('div.combobox-item:visible:first');
				}
			} else {
				item = item.prevAll('div.combobox-item:visible:first');
				if (!item.length){
					item = panel.children('div.combobox-item:visible:last');
				}
			}
		}
		if (item.length){
			item.addClass('combobox-item-hover');
			scrollTo(target, item.attr('value'));
			if (opts.selectOnNavigation){
				select(target, item.attr('value'));
			}
		}
	}
	
	/**
	 * select the specified value
	 */
	function select(target, value){
		var opts = $.data(target, 'combobox').options;
		var data = $.data(target, 'combobox').data;
		
		if (opts.multiple){
			var values = $(target).combo('getValues');
			for(var i=0; i<values.length; i++){
				if (values[i] == value) return;
			}
			values.push(value);
			setValues(target, values);
		} else {
			setValues(target, [value]);
		}
		
		var item = findDataItem(data, opts.valueField, value);
		if (item){
			opts.onSelect.call(target, item);
		}
	}
	
	/**
	 * unselect the specified value
	 */
	function unselect(target, value){
		var state = $.data(target, 'combobox');
		var opts = state.options;
		var values = $(target).combo('getValues');
//		var index = values.indexOf(value+'');
		var index = $.inArray(value+'', values);
		if (index >= 0){
			values.splice(index, 1);
			setValues(target, values);
		}
		var item = findDataItem(state.data, opts.valueField, value);
		if (item){
			opts.onUnselect.call(target, item);
		}
	}
	
	/**
	 * set values
	 */
	function setValues(target, values, remainText){
		var opts = $.data(target, 'combobox').options;
		var data = $.data(target, 'combobox').data;
		var panel = $(target).combo('panel');
		
		panel.find('div.combobox-item-selected').removeClass('combobox-item-selected');
		var vv = [], ss = [];
		for(var i=0; i<values.length; i++){
			var v = values[i];
			var s = v;
			var item = findDataItem(data, opts.valueField, v);
			if (item){
				s = item[opts.textField];
			}
			vv.push(v);
			ss.push(s);
			panel.find('div.combobox-item[value="' + v + '"]').addClass('combobox-item-selected');
		}
		
		$(target).combo('setValues', vv);
		if (!remainText){
			$(target).combo('setText', ss.join(opts.separator));
		}
	}
	
	/**
	 * load data, the old list items will be removed.
	 */
	function loadData(target, data, remainText){
		var state = $.data(target, 'combobox');
		var opts = state.options;
		state.data = opts.loadFilter.call(target, data);
		data = state.data;
		
		var selected = $(target).combobox('getValues');
		var dd = [];
		var group = undefined;
		for(var i=0; i<data.length; i++){
			var item = data[i];
			var v = item[opts.valueField];
			var s = item[opts.textField];
			var g = item[opts.groupField];
			
			if (g){
				if (group != g){
					group = g;
					dd.push('<div class="combobox-group" value="' + g + '">');
					dd.push(opts.groupFormatter ? opts.groupFormatter.call(target, g) : g);
					dd.push('</div>');
				}
			} else {
				group = undefined;
			}
			
			dd.push('<div class="combobox-item' + (g ? ' combobox-gitem' : '') + '" value="' + v + '"' + (g ? ' group="' + g + '"' : '') + '>');
			dd.push(opts.formatter ? opts.formatter.call(target, item) : s);
			dd.push('</div>');
			
			if (item['selected']){
				(function(){
					for(var i=0; i<selected.length; i++){
						if (v == selected[i]) return;
					}
					selected.push(v);
				})();
			}
		}
		$(target).combo('panel').html(dd.join(''));
		
		if (opts.multiple){
			setValues(target, selected, remainText);
		} else {
			if (selected.length){
				setValues(target, [selected[selected.length-1]], remainText);
			} else {
				setValues(target, [], remainText);
			}
		}
		
		opts.onLoadSuccess.call(target, data);
	}
	
	/**
	 * request remote data if the url property is setted.
	 */
	function request(target, url, param, remainText){
		var opts = $.data(target, 'combobox').options;
		if (url){
			opts.url = url;
		}
//		if (!opts.url) return;
		param = param || {};
		
		if (opts.onBeforeLoad.call(target, param) == false) return;

		opts.loader.call(target, param, function(data){
			loadData(target, data, remainText);
		}, function(){
			opts.onLoadError.apply(this, arguments);
		});
	}
	
	/**
	 * do the query action
	 */
	function doQuery(target, q){
		var state = $.data(target, 'combobox');
		var opts = state.options;
		
		if (opts.multiple && !q){
			setValues(target, [], true);
		} else {
			setValues(target, [q], true);
		}
		
		if (opts.mode == 'remote'){
			request(target, null, {q:q}, true);
		} else {
			var panel = $(target).combo('panel');
			panel.find('div.combobox-item,div.combobox-group').hide();
			var data = state.data;
			var group = undefined;
			for(var i=0; i<data.length; i++){
				var item = data[i];
				if (opts.filter.call(target, q, item)){
					var v = item[opts.valueField];
					var s = item[opts.textField];
					var g = item[opts.groupField];
					var item = panel.find('div.combobox-item[value="' + v + '"]');
					item.show();
					if (s == q){
						setValues(target, [v], true);
						item.addClass('combobox-item-selected');
					}
					if (opts.groupField && group != g){
						panel.find('div.combobox-group[value="' + g + '"]').show();
						group = g;
					}
				}
			}
		}
	}
	
	function doEnter(target){
		var t = $(target);
		var panel = t.combobox('panel');
		var opts = t.combobox('options');
		var data = t.combobox('getData');
		var item = panel.children('div.combobox-item-hover');
		if (!item.length){
			item = panel.children('div.combobox-item-selected');
		}
		if (!item.length){return}
		if (opts.multiple){
			if (item.hasClass('combobox-item-selected')){
				t.combobox('unselect', item.attr('value'));
			} else {
				t.combobox('select', item.attr('value'));
			}
		} else {
			t.combobox('select', item.attr('value'));
			t.combobox('hidePanel');
		}
		var vv = [];
		var values = t.combobox('getValues');
		for(var i=0; i<values.length; i++){
			if (findDataItem(data, opts.valueField, values[i])){
				vv.push(values[i]);
			}
		}
		t.combobox('setValues', vv);
	}
	
	/**
	 * create the component
	 */
	function create(target){
		var opts = $.data(target, 'combobox').options;
		$(target).addClass('combobox-f');
		$(target).combo($.extend({}, opts, {
			onShowPanel: function(){
				$(target).combo('panel').find('div.combobox-item').show();
				scrollTo(target, $(target).combobox('getValue'));
				opts.onShowPanel.call(target);
			}
		}));
		
		$(target).combo('panel').unbind().bind('mouseover', function(e){
			$(this).children('div.combobox-item-hover').removeClass('combobox-item-hover');
			$(e.target).closest('div.combobox-item').addClass('combobox-item-hover');
			e.stopPropagation();
		}).bind('mouseout', function(e){
			$(e.target).closest('div.combobox-item').removeClass('combobox-item-hover');
			e.stopPropagation();
		}).bind('click', function(e){
			var item = $(e.target).closest('div.combobox-item');
			if (!item.length){return}
			var value = item.attr('value');
			if (opts.multiple){
				if (item.hasClass('combobox-item-selected')){
					unselect(target, value);
				} else {
					select(target, value);
				}
			} else {
				select(target, value);
				$(target).combo('hidePanel');
			}
			e.stopPropagation();
		});
	}
	
	$.fn.combobox = function(options, param){
		if (typeof options == 'string'){
			var method = $.fn.combobox.methods[options];
			if (method){
				return method(this, param);
			} else {
				return this.combo(options, param);
			}
		}
		
		options = options || {};
		return this.each(function(){
			var state = $.data(this, 'combobox');
			if (state){
				$.extend(state.options, options);
				create(this);
			} else {
				state = $.data(this, 'combobox', {
					options: $.extend({}, $.fn.combobox.defaults, $.fn.combobox.parseOptions(this), options),
					data: []
				});
				create(this);
				var data = $.fn.combobox.parseData(this);
				if (data.length){
					loadData(this, data);
				}
			}
			if (state.options.data){
				loadData(this, state.options.data);
			}
			request(this);
		});
	};
	
	
	$.fn.combobox.methods = {
		options: function(jq){
			var copts = jq.combo('options');
			return $.extend($.data(jq[0], 'combobox').options, {
				originalValue: copts.originalValue,
				disabled: copts.disabled,
				readonly: copts.readonly
			});
		},
		getData: function(jq){
			return $.data(jq[0], 'combobox').data;
		},
		setValues: function(jq, values){
			return jq.each(function(){
				setValues(this, values);
			});
		},
		setValue: function(jq, value){
			return jq.each(function(){
				setValues(this, [value]);
			});
		},
		clear: function(jq){
			return jq.each(function(){
				$(this).combo('clear');
				var panel = $(this).combo('panel');
				panel.find('div.combobox-item-selected').removeClass('combobox-item-selected');
			});
		},
		reset: function(jq){
			return jq.each(function(){
				var opts = $(this).combobox('options');
				if (opts.multiple){
					$(this).combobox('setValues', opts.originalValue);
				} else {
					$(this).combobox('setValue', opts.originalValue);
				}
			});
		},
		loadData: function(jq, data){
			return jq.each(function(){
				loadData(this, data);
			});
		},
		reload: function(jq, url){
			return jq.each(function(){
				request(this, url);
			});
		},
		select: function(jq, value){
			return jq.each(function(){
				select(this, value);
			});
		},
		unselect: function(jq, value){
			return jq.each(function(){
				unselect(this, value);
			});
		}
	};
	
	$.fn.combobox.parseOptions = function(target){
		var t = $(target);
		return $.extend({}, $.fn.combo.parseOptions(target), $.parser.parseOptions(target,[
			'valueField','textField','groupField','mode','method','url'
		]));
	};
	
	$.fn.combobox.parseData = function(target){
		var data = [];
		var opts = $(target).combobox('options');
		$(target).children().each(function(){
			if (this.tagName.toLowerCase() == 'optgroup'){
				var group = $(this).attr('label');
				$(this).children().each(function(){
					_parseItem(this, group);
				});
			} else {
				_parseItem(this);
			}
		});
		return data;
		
		function _parseItem(el, group){
			var t = $(el);
			var item = {};
			item[opts.valueField] = t.attr('value')!=undefined ? t.attr('value') : t.html();
			item[opts.textField] = t.html();
			item['selected'] = t.is(':selected');
			if (group){
				opts.groupField = opts.groupField || 'group';
				item[opts.groupField] = group;
			}
			data.push(item);
		}
	};
	
	$.fn.combobox.defaults = $.extend({}, $.fn.combo.defaults, {
		valueField: 'value',
		textField: 'text',
		groupField: null,
		groupFormatter: function(group){return group;},
		mode: 'local',	// or 'remote'
		method: 'post',
		url: null,
		data: null,
		
		keyHandler: {
			up: function(){nav(this,'prev')},
			down: function(){nav(this,'next')},
			enter: function(){doEnter(this)},
			query: function(q){doQuery(this, q)}
		},
		filter: function(q, row){
			var opts = $(this).combobox('options');
			return row[opts.textField].indexOf(q) == 0;
		},
		formatter: function(row){
			var opts = $(this).combobox('options');
			return row[opts.textField];
		},
		loader: function(param, success, error){
			var opts = $(this).combobox('options');
			if (!opts.url) return false;
			$.ajax({
				type: opts.method,
				url: opts.url,
				data: param,
				dataType: 'json',
				success: function(data){
					success(data);
				},
				error: function(){
					error.apply(this, arguments);
				}
			});
		},
		loadFilter: function(data){
			return data;
		},
		
		onBeforeLoad: function(param){},
		onLoadSuccess: function(){},
		onLoadError: function(){},
		onSelect: function(record){},
		onUnselect: function(record){}
	});
})(jQuery);
