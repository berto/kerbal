const awsURL = 'https://s3-us-west-2.amazonaws.com/kerbal.me'
const serverURL = 'http://localhost:3000'

const endpoints = {
  items: `${serverURL}/api/items`,
  kerbal: `${serverURL}/kerbal/`,
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
