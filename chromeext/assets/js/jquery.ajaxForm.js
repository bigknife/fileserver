(function($){
	$.fn.ajaxFormSubmit = function(params){
		params = $.extend( {}, params);
		var beforeSubmit = function($form){
			if(params.disabledButtonIds){
				$.each(params.disabledButtonIds,function(){
					$('#' + this).attr('disabled','disabled');
				});
			}else{
				//disabled all buttons in form
				$form.find('button').attr('disabled','disabled');
			}
			//dataloading
			if(params.loadingButtonIds){
				$.each(params.loadingButtonIds,function(){
					$('#' + this).button('loading');
				});
			}
		};
		var endSubmit = function($form){
			if(params.disabledButtonIds){
				$.each(params.disabledButtonIds,function(){
					$('#' + this).attr('disabled',null);
				});
			}else{
				//disabled all buttons in form
				$form.find('button').attr('disabled',null);
			}
			//dataloading
			if(params.loadingButtonIds){
				$.each(params.loadingButtonIds,function(){
					$('#' + this).button('reset');
				});
			}
		};
		this.each(function(){
			var $form = $(this);
			$form.submit(Function.literal('return false;'));
			var submit = function(){
				var data = $form.serialize();
				var url = params.url || $form.attr('action');
				var type = params.type || $form.attr('method') || 'POST';
				var dataType = params.dataType || $form.attr('data-type') || 'html';
				var success = params.success || Function.literal('return false;');
				var error = params.error || Function.literal('return false;');
				var cache = params.cache || false;
				
				beforeSubmit($form);
				if(params.beforeSubmit){
					params.beforeSubmit();
				}
				$.ajax(url,{
					cache : cache,
					data : data,
					type : type,
					dataType : dataType
				}).success(function(){
					success.apply(arguments.callee,arguments);
					endSubmit($form);
					if(params.afterSuccess){
						params.afterSuccess(arguments.callee,arguments);
					}
				}).error(function(){
					error.apply(arguments.callee,arguments);
					endSubmit($form);
					if(params.afterError){
						params.afterError(arguments.callee,arguments);
					}
				});
			};
			submit();
			
		});
		return this;
	};
})(jQuery);