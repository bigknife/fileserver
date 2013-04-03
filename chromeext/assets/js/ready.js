/** some extra */
(function() {
    String.prototype.template = function(o) {
        return this.replace(/{([^{}]*)}/g, function(a, b) {
            var r = o[b];
            return typeof r === 'string' || typeof r === 'number' ? r : a;
        });
    };
    var literalFuncCache = {};
    var literal = Function.prototype.literal = function(str) {
            var f = literalFuncCache[str];
            if (!f) {
                var func = '(function(){return function(){' + str + '};})();'
                f = eval(func);
                literalFuncCache[str] = f;
            }
            return f;
        };
    Function.prototype.runLiteral = function(str) {
        return literal(str)();
    };
})(); /*界面共同元素初始化*/
(function() {
    $('[rel="tooltip"]').tooltip({
        placement: 'bottom'
    });
    $("input[type=text]").focus(function() {
        Function.literal('$(this).select()');
    });

    //最小化按钮
    $('.btn-minimize').click(

    function(e) {
        e.preventDefault();
        var $target = $(this).parent().parent().next('.box-content');
        if ($target.is(':visible')) {
            $('i', $(this)).removeClass('icon-chevron-up').addClass('icon-chevron-down');
        } else {
            $('i', $(this)).removeClass('icon-chevron-down').addClass('icon-chevron-up');
        }
        $target.slideToggle();
    });
})(); /** 设置菜单active */
(function() {
    // set init
    $('li[rel="leftmenu"]').each(function() {
        $(this).find('a').each(function() {
            if ($(this).attr('href') == $.cookie('leftmenu_active_url')) {
                $(this).parent().addClass('active');
            }
        });
    });

    $('li[rel="navmenu"]').each(function() {
        $(this).find('a').each(function() {
            if ($(this).attr('href') == $.cookie('navmenu_active_url')) {
                $(this).parent().addClass('active');
            }
        });
    });

    $('li[rel="leftmenu"] > a').each(function() {
        $(this).click(function(e) {
            e.preventDefault();
            $.cookie('leftmenu_active_url', $(this).attr('href'), {
                path: '/'
            });

            $('li[rel="leftmenu"]').removeClass('active');

            $(this).parent().addClass('active');
            window.location.href = $(this).attr('href');
        });
    });
/*
	$('li[rel="navmenu"] > a').each(function() {
		$(this).click(function(e) {
			e.preventDefault();
			$.cookie('navmenu_active_url', $(this).attr('href'), {
				path : '/'
			});

			$('li[rel="leftmenu"]').removeClass('active');

			$(this).parent().addClass('active');
			window.location.href = $(this).attr('href');
		});
	});
	*/
    //$('li[rel="navmenu"] > a').attr('target','_self');
})();

/** called by mainFrame */
$(function() {

    // some global function
    $.ajaxSetup({
        complete: function(xhr, status) {

            if (xhr.responseText == '{"filterFoundForbidden":true}') {
                // 访问禁止
                window.location.href = g.dynamicResRoot + "/forbidden";
            }

            if (xhr.responseText == '{"filterFoundNotLogin":true,"redirect":"relogin"}') {
                // 重新登录
                alert('登录过期，需要重新登录！')
                window.location.href = g.dynamicResRoot;
            }
        }
    });

    if (docReady) {
        docReady();
    }
});

/**主菜单在新窗口打开*/
$(function() {

});