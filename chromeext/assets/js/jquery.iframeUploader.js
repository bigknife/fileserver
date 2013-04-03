(function($) {
    $.fn.iframeUpload = function(params) {
        params = $.extend({}, params);

        this.each(function() {
            var $form = $(this);
            var targetId = '_iframeuploadtarget_' + $form.attr('id');
            var $target = $form.find('#' + targetId);
            if ($target.length == 0) {
                var iframe = $('<iframe></iframe>');
                $target = $(iframe);
                $target.attr('name',targetId);
                $target.attr('id',targetId);
                $target.css({
                	'width':'0px',
                	'height':'0px',
                    'display': 'none'
                })
                $form.append(iframe);
            }

            $form.attr('target', targetId);
            $form.attr('enctype', 'multipart/form-data');
            $form.attr('method', 'POST');

            var url = params['url'] || $form.attr('action');
            $form.attr('action', url);

            //事件处理
            var start = params['start'];
            if (start) {
                start();
            }

            if (params.disabledButtonIds) {
                $.each(params.disabledButtonIds, function() {
                    $('#' + this).attr('disabled', 'disabled');
                });
            } else {
                //disabled all buttons in form
                $form.find('button').attr('disabled', 'disabled');
            }
            //dataloading
            if (params.loadingButtonIds) {
                $.each(params.loadingButtonIds, function() {
                    $('#' + this).button('loading');
                });
            }


            var load = params['load'];
            $target.unbind('load');
            $target.load(function() {
                if (params.disabledButtonIds) {
                    $.each(params.disabledButtonIds, function() {
                        $('#' + this).attr('disabled', null);
                    });
                } else {
                    //disabled all buttons in form
                    $form.find('button').attr('disabled', null);
                }
                //dataloading
                if (params.loadingButtonIds) {
                    $.each(params.loadingButtonIds, function() {
                        $('#' + this).button('reset');
                    });
                }
                if (load) {
                    var bodyHtml = $(this).contents().find('body').html();
                    load(bodyHtml, $(this));
                }
            });

            //如果有禁用的按钮，自动禁用
            $form.submit();
        });
        return this;
    };
})(jQuery);