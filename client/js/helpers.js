const baseUrl = 'https://s3-us-west-2.amazonaws.com/kerbal.me'
const endpoints = {
  items: '/api/items',
  kerbal: '/kerbal/',
}
const displayError = (error) => {
  const alertBox = $('.callout')
  const fadeOutTime = 5000
  alertBox.fadeIn()
  $('#error-message').text(error)
  setTimeout(() => {
    console.log('i')
    alertBox.fadeOut(1000)
  }, fadeOutTime)
}
