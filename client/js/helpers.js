const imagesURL = '/images'
const kerbalURL = '/kerbals'
const downloadURL = '/download.html'
const serverURL = 'https://dq8r27wezj.execute-api.us-west-2.amazonaws.com/prod'

const endpoints = {
  items: `${serverURL}`,
  kerbal: `${serverURL}`,
}
const displayError = (error) => {
  const alertBox = $('.callout')
  const fadeOutTime = 5000
  alertBox.fadeIn()
  $('#error-message').text(error)
  setTimeout(() => {
    alertBox.fadeOut(1000)
  }, fadeOutTime)
}

const handleResponse = (response) => {
  if (response.status !== 200) {
    return Promise.reject(response.json())
  }
  return response.json()
}
