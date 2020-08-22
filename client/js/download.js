const validateId = (id) => {
  try {
    const length = 40
    const regexp = /(0[xX])?[a-fA-F0-9]+$/
    return id.length === length && regexp.test(id)
  } catch (e) {
    return false
  }
}

const parseIdParam = () => {
  const query = window.location.search.substring(1)
  const params = query.split('&')
  for (let i = 0; i < params.length; i++) {
    const pair = params[i].split('=')
    if (decodeURIComponent(pair[0]) == 'id') {
      return decodeURIComponent(pair[1])
    }
  }
}

const displayAvatar = () => {
  return new Promise((resolve) => {
    const id = parseIdParam()
    if (!validateId(id)) {
      displayError('Invalid id: ' + id)
    }
    const box = $('.download-avatar')
    const url = `${baseUrl}/kerbals/${id}.png`
    $(`<img src="${url}" />`)
      .on('load', function () {
        box.empty()
        box.append($(this))
        resolve(url)
      })
      .on('error', () => {
        $('.loading').text('Kerbal not found :(')
      })
  })
}

const downloadImage = (url) => {
  const link = $('<a />')
  link.attr('download', 'kerbal')
  link.attr('href', url)
  link.attr('id', 'kerbal-link')
  $('body').append(link)
  document.querySelector('#kerbal-link').click()
  link.remove()
}

const enableDownload = (url) => {
  const download = $('#download')
  download.removeClass('disabled')
  download.click(function () {
    downloadImage(url)
    // const button = $(this)
    // button.text('Downloading...')
    // button.addClass('disabled')
  })
}

const titleRedirect = () => {
  $('header h1').click(() => {
    window.location.href = '/'
  })
}

const init = () => {
  titleRedirect()
  displayAvatar().then(enableDownload)
}

$(document).ready(init)
