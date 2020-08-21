$(document).foundation()

const baseUrl = 'https://s3-us-west-2.amazonaws.com/kerbal.me'
const endpoints = {
  items: '/api/items',
  kerbal: '/kerbal/',
}
const suitFolder = 'suit'
const previewKey = '-preview'
const frontKey = '-front'

let currentKerbal = {}

const initialItems = {
  suit: 'helmet.png',
  color: 'green.png',
  mouth: 'smile.png',
  hair: 'hair.png',
  eyes: 'eyes.png',
  'facial-hair': '',
  glasses: '',
  extras: '',
}

const removeKey = (word, key) => {
  const index = word.indexOf(key)
  if (index >= 0) {
    return word.slice(0, index) + word.slice(key.length + index, word.length)
  }
  return word
}

const updateKerbal = (folder, item) => {
  const box = $(`#kerbal-${folder}`)
  box.empty()
  if (item) {
    item = removeKey(item, previewKey)
    const imageHTML = `<img src="${baseUrl}/${folder}/${item}" />`
    box.append(imageHTML)
    if (folder === suitFolder) {
      const helmetFront = $('#kerbal-suit-front')
      helmetFront.empty()
      item = removeKey(item, '.png')
      helmetFront.append(`<img src="${baseUrl}/${folder}/${item}${frontKey}.png" />`)
    }
  }
}

const removeActive = (folder) => {
  $(`#${folder}`)
    .children()
    .each(function () {
      $(this).removeClass('active')
    })
}

const activateNoneCards = () => {
  $('.none').click(function () {
    const card = $(this)
    const folder = card.parent().attr('id')
    removeActive(folder)
    card.addClass('active')
    currentKerbal[folder] = ''
    updateKerbal(folder)
  })
}

const generateCard = (folder, item) => {
  const img = $('<img />')
  img.addClass('image-card')
  img.attr('src', `${baseUrl}/${folder}/${item}`)
  img.attr('id', removeKey(item, previewKey))
  img.click(function () {
    removeActive(folder)
    img.addClass('active')
    currentKerbal[folder] = removeKey(item, previewKey)
    updateKerbal(folder, item)
  })
  return img
}

const displayImages = () => {
  return fetch(endpoints.items)
    .then((response) => response.json())
    .then((data) => {
      for (const folder in data) {
        const box = $(`#${folder}`)
        if (data[folder].length === 0) {
          continue
        }
        data[folder].forEach((item) => {
          if (item.indexOf('preview') > 0) box.append(generateCard(folder, item))
        })
      }
    })
    .catch(console.error)
}

const loadKerbal = () => {
  Object.keys(initialItems).forEach((folder) => {
    removeActive(folder)
    currentKerbal[folder] = initialItems[folder]
    if (initialItems[folder] === '') {
      $(`#${folder} .none`).addClass('active')
      return
    }
    $(`[id="${initialItems[folder]}"]`).addClass('active')
  })
}

const displayKerbal = () => {
  Object.keys(currentKerbal).forEach((folder) => {
    updateKerbal(folder, currentKerbal[folder])
  })
}

const activateButtons = () => {
  const reset = $('#reset')
  reset.removeClass('disabled')
  reset.click(() => {
    loadKerbal()
    displayKerbal()
  })
  const save = $('#save')
  save.removeClass('disabled')
  save.click(() => {
    const body = {
      ...currentKerbal,
      'suit-front': removeKey(currentKerbal.suit, '.png') + frontKey + '.png',
    }
    fetch(endpoints.kerbal, {
      method: 'post',
      body: JSON.stringify(body),
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((response) => response.json())
      .then(console.log)
      .catch(console.error)
  })
}

const init = () => {
  displayImages().then(() => {
    loadKerbal()
    displayKerbal()
    activateNoneCards()
    activateButtons()
  })
}

$(document).ready(init)
