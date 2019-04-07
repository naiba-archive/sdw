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
//申请收录
$('#exampleModal').on('show.bs.modal', function (event) {
  var button = $(event.relatedTarget)
  var recipient = button.data('whatever')
  var modal = $(this)
  modal.find('.modal-title').text('New message to ' + recipient)
  modal.find('.modal-body input').val(recipient)
})