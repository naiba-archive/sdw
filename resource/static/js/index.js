$('a').hover(function () {
  if (!$(this).data('desc')) {
    return
  }
  let offsetTop = $(this).offset().top + 20
  let offsetLeft = $(this).offset().left
  $('.dialog').css({
    'display': 'block', 'top': offsetTop,
    'left': offsetLeft, 'width': 'auto',
  })
  $('.dialog').html(`<p>${$(this).data('desc')}</p>`)
}, function () {
  $('.dialog').css({ 'display': 'none' })
})
//提交数据
$('#submit').click(function () {
  if (!$('#content').val()) {
    alert("请填写站点及站点简介")
    return
  }
  if (!$('#g-recaptcha-response').val()) {
    alert("请通过机器人验证")
    return
  }
  $.post("/submit", {
    content: $('#content').val(),
    captcha: $('#g-recaptcha-response').val()
  }).done(function () {
    alert('提交成功，管理员已收到您的请求')
    $('#myModal').modal('hide')
  }).fail(function (resp) {
    alert(resp.responseText)
  })
})