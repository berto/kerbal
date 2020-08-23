const awsURL = 'https://s3-us-west-2.amazonaws.com/kerbal.me'
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
    console.log('i')
    alertBox.fadeOut(1000)
  }, fadeOutTime)
}

const handleResponse = (response) => {
  console.warn(response)
  if (response.status !== 200) {
    return Promise.reject(response.json())
  }
  return response.json()
}
