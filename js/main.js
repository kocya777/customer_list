$(function(){
	$(".iziModal").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal2").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal_alert").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal3").iziModal({
		headerColor: '#333333',
		width: 935,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal4").iziModal({
		headerColor: '#333333',
		width: 935,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal5").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal6").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal7").iziModal({
		headerColor: '#333333',
		width: 935,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal8").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal9").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal10").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal11").iziModal({
		headerColor: '#333333',
		width: 935,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal12").iziModal({
		headerColor: '#333333',
		width: 935,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal13").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal14").iziModal({
		headerColor: '#333333',
		width: 935,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal15").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal16").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal17").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})

$(function(){
	$(".iziModal18").iziModal({
		headerColor: '#333333',
		width: 400,
		overlayColor: 'rgba(0, 0, 0, 0.5)',
		transitionIn: 'fadeInUp',
		padding:15,
	});
})


// HTMLの要素が全て準備できれば処理を行う
var setBoxId = '#secondary_button';      // スクロールさせる要素
var initOffsetTop = null;   // 要素の初期位置
$(document).ready(function() {
    // 初期位置取得
    initOffsetTop = $(setBoxId).offset().top;

});

//スクロールしたらこの処理が走る
$(window).scroll(function() {
    scrollbox();
});

// スクロール処理
function scrollbox(){
    // 初期位置が取れていなければ処理を抜ける
    if(initOffsetTop == null) return;

    // 現在のスクロール位置を取得
    var scrollTop = $(document).scrollTop();

    // スクロールさせる要素の初期位置と現在のスクロールの位置を比較
    //初期位置より下にスクロールした時
    if(initOffsetTop < scrollTop) {
        // positionを設定
        $(setBoxId).css('position', 'fixed');
        // topの位置を設定
        $(setBoxId).animate({top: '10'}, {duration: 0});
    } else {
        // 設定したスタイルを持とに戻す
        $(setBoxId).css('position', 'absolute');
        $(setBoxId).animate({top: initOffsetTop}, {duration: 0});
    }
}

$(window).load(function(){
	$('.loading').fadeOut();	
});

$('#datepicker').datepicker();