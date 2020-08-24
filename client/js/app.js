$(document).foundation()

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
  if (item) {
    item = removeKey(item, previewKey)
    $(`<img src="${imagesURL}/${folder}/${item}" />`).on('load', function () {
      box.empty()
      box.append($(this))
    })
    if (folder === suitFolder) {
      const helmetFront = $('#kerbal-suit-front')
      item = removeKey(item, '.png')
      $(`<img src="${imagesURL}/${folder}/${item}${frontKey}.png" />`).on('load', function () {
        helmetFront.empty()
        helmetFront.append($(this))
      })
    }
  } else {
    box.empty()
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
  img.attr('src', `${imagesURL}/${folder}/${item}`)
  img.attr('id', removeKey(item, previewKey))
  img.click(function () {
    if ($(this).hasClass('active')) return
    removeActive(folder)
    img.addClass('active')
    currentKerbal[folder] = removeKey(item, previewKey)
    updateKerbal(folder, item)
  })
  return img
}

const displayImages = () => {
  return fetch(endpoints.items)
    .then(handleResponse)
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
    .catch((response) => {
      const message = response.error || 'Failed to load images, please refresh'
      displayError(message)
    })
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
    save.text('Launching Kerbal...')
    save.addClass('disabled')
    const body = {
      ...currentKerbal,
      'suit-front': removeKey(currentKerbal.suit, '.png') + frontKey + '.png',
    }
    const errorMessage = 'Failed to create Kerbal, please try again'
    fetch(endpoints.kerbal, {
      method: 'post',
      body: JSON.stringify(body),
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then(handleResponse)
      .then((response) => {
        console.log(response)
        if (!response || !response.id) {
          return Promise.reject(errorMessage)
        }
        window.location.href = `${downloadURL}?id=` + response.id
      })
      .catch((response) => {
        save.removeClass('disabled')
        save.text('Save and Continue...')
        const message = response.error || errorMessage
        displayError(message)
      })
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
