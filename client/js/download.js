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
    $(`<img src="${baseUrl}/kerbals/${id}.png" />`)
      .on('load', function () {
        box.empty()
        box.append($(this))
        resolve('loaded')
      })
      .on('error', () => {
        $('.loading').text('Kerbal not found :(')
      })
  })
}

const enableDownload = () => {
  const download = $('#download')
  download.removeClass('disabled')
  download.click(function () {
    const button = $(this)
    button.text('Downloading...')
    button.addClass('disabled')
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
