function docReady(){
	var username, password;
	var withAuth = false;
	var loadFileList;
	//加载目录信息
	make_base_auth = function(){
		return "Basic " + $.base64.encode( username +":" + password );
	}
	var dirLoaded = function(list,root){
		var html ="<ul id='dirList'>"
		html += "<li><a href='../' isDir='true'>【上一级】</a></li>"
		$.each(list,function(e){
			var f = "<li><a href='";
			f += this.Name;
			f += "' isDir='";
			f += this.IsDir;
			f += "'>"
			f += this.Name + "; 大小:" + (this.Size/1024) + "KB; " + (this.IsDir ? "文件夹" : "文件");
			f += "</a></li>"
			html += f
			console.log(f)
		})
		$('#view').html(html+"</ul>");
		$('#dirList a').unbind('click');
		$('#dirList a').click(function(e){
			e.preventDefault();
			var $f = $(this);
			console.log($f)
			if($f.attr('isDir') == "true"){
				loadFileList($f.attr('href'),"false");
			}else{
				window.open("http://ydc-dev-0:20000/file?name=" + $f.attr('href'));
			}
		});
		
	};
	loadFileList = function(root,recursion){
		
		$.ajax("http://ydc-dev-0:20000/dir?name="+root+"&recursion="+recursion,{
			type:'GET',
			dataType:'json',
			success : function(dir){
				dirLoaded(dir,root);
			},
			error : function(e){
				$('#loginDlg').modal();
			},
			beforeSend : function(xhr){
				if(withAuth == true){
					xhr.setRequestHeader('Authorization', make_base_auth());
				}
				
			}
		});
	}
	
	loadFileList("/","false");
	$('#loginDlgOk').click(function(){
		username = $('#inputName').val();
		password = $('#inputPassword').val();
		$('#loginForm')[0].reset();
		$('#loginDlg').modal('hide');
		withAuth = true;
		loadFileList("/","false");
	});
	
	//上传消息
	/*$('#uploadTarget').load(function(){
		console.log($(this).contents())
		var bodyHtml = $(this).contents().find('body').html();
		alert(bodyHtml);
	});*/
	var uploadFileChanged = function(files){
		$('p#dropZone').empty();
		$('p#dropZone').append($('<div>稍等，正在上传文件：</div>'));
			
		$.each(files,function(index,file){
			$('p#dropZone').append($('<div>'+file.name+'</div>'));
		});
	};
    $('#fileupload').fileupload({
    	'option':{
    		'dropZone':$('p#dropZone')
    	},
    	add : function(e,data){
			data.submit();
    	},
    	change: function (e, data) {
    		uploadFileChanged(data.files);
	    },
	    drop : function(e,data){
	    	uploadFileChanged(data.files);
	    },
	    done : function(){
	    	$('p#dropZone').empty();
	    	$('p#dropZone').append($('<div>上传完成！</div>'));
	    	loadFileList("/","false");
	    },
	    error : function(){
	    	$('p#dropZone').empty();
	    	$('p#dropZone').append($('<div>上传出错了！</div>'));
	    }
    });
}