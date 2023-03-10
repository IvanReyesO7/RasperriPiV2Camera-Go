import $ from 'jquery';

let placeholder = "https://www.shoshinsha-design.com/wp-content/uploads/2020/05/noimage.png"

$('#preview-btn').click(function(){
    $(this).toggleClass('btn-primary btn-danger');
    if ($(this).text() == "On"){
       $(this).text("Off");
       $('#video-frame').attr('src', '/video');
    } else {
       $(this).text("On");
       $('#video-frame').attr('src', placeholder);
    }
})