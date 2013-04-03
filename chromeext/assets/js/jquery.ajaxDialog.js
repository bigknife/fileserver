/**
 * 先定义一个modal，然后如下方法调用
 * $('#tenantDlg').ajaxDialog({title:'hello,world',url : '/abc',method:'POST'})
 * 
 */
(function($){
	$.fn.ajaxDialog = function(params){
		params = $.extend( {
			title : '未标题对话框',
			method : 'GET'
			
		}, params);
		
		this.each(function(){
			var $dlg = $(this);
			var titleId = $dlg.attr('data-title-id');
			$dlg.find('#' + titleId).html(params.title);
			
			var loaderPic = $dlg.attr('data-loading-pic');
			var loadingText = $dlg.attr('data-loading-text');
			var loadingError = $dlg.attr('data-loading-error');
			var loadingHtml = '<div class="ajaxDialogLoading"><img src="' + loaderPic + '" />' + loadingText + '</div>';
			$dlg.find('.modal-body').html(loadingHtml);		
			$dlg.modal();
			
			var url = params.url;
			url = url || $dlg.attr('data-url');
			var method = params.method;
			method = method || $dlg.attr('data-method');
			$.ajax(url,{
				type : method || 'GET',
				dataType : 'html',
				success : function(html){
					var $modalBody = $dlg.find('.modal-body');
					var bodyEl = null;
					try
					{
						bodyEl = $(html);
						$modalBody.empty();
						$modalBody.append(bodyEl);
						$(bodyEl).css({'margin':'0px auto'});
					}catch(error)
					{
						
					}
					if(bodyEl != null){
						$dlg.modal().css({
						       'width': function () { 
						           return ($(bodyEl).width() + 50) + 'px';  
						       },
						       'margin-left': function () { 
						           return -($(this).width() / 2); 
						       }
						});
						if(params.loaded){
							params.loaded($dlg);
						}
					}
				},
				error : function(){
					$dlg.find('.modal-body').html(loadingError || '<div class="ajaxDialogLoadError">系统繁忙，请稍后重试！</div>');
				}
			});
		});
		return this;
	};
})(jQuery);