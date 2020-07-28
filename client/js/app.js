$(document).foundation()
const $kerbalContainer = $('#kerbal')

const baseUrl = 'https://s3-us-west-2.amazonaws.com/kerbal.me'

let currentKerbal = []

const initialItems = {
  face: 'face',
  extras: 'glare',
  hair: 'hair',
}

const loadInitialKerbal = () => {
  Object.keys(initialItems).forEach((folder) => {
    currentKerbal.push(`${baseUrl}/${folder}/${initialItems[folder]}.png`)
  })
}

const displayKerbal = () => {
  console.log(currentKerbal)
  $kerbalContainer.empty()
  currentKerbal.forEach((link) => {
    const imageHTML = `<img src="${link}" />`
    console.log(imageHTML)
    $kerbalContainer.append(imageHTML)
  })
}

const init = () => {
  loadInitialKerbal()
  displayKerbal()
}

init()
