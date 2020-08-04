$(document).foundation()
const $kerbalContainer = $('#kerbal')

const baseUrl = 'https://s3-us-west-2.amazonaws.com/kerbal.me'

let currentKerbal = []

const items = {
  suit: 'helmet',
  color: 'green',
  mouth: 'smile',
  hair: 'hair',
  eyes: 'eyes',
}

const loadKerbal = () => {
  currentKerbal = []
  let suitName
  Object.keys(items).forEach((folder) => {
    let extra = ''
    if (folder === 'suit') {
      extra = '-back'
      suitName = items[folder]
    }
    currentKerbal.push(`${baseUrl}/${folder}/${items[folder]}${extra}.png`)
  })
  currentKerbal.push(`${baseUrl}/suit/${suitName}-front-glare.png`)
}

const displayKerbal = () => {
  loadKerbal()
  $kerbalContainer.empty()
  currentKerbal.forEach((link) => {
    const imageHTML = `<img src="${link}" />`
    console.log(imageHTML)
    $kerbalContainer.append(imageHTML)
  })
}

const init = () => {
  displayKerbal()
}

init()
