'use strict';

(function (windows) {
  function uploadText(uploadBody, baseUrl = '') {
    return fetch(baseUrl + '/api/paste', {
      method: 'POST',
      body: JSON.stringify(uploadBody),
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((response) => {
        const contentType = response.headers.get('content-type');
        const isJson =
          contentType && contentType.indexOf('application/json') !== -1;
        // Success case is an HTTP 200 response and a JSON body.
        if (response.status === 200 && isJson) {
          return Promise.resolve(response.json());
        }
        // Treat any other response as an error.
        return response.text().then((text) => {
          if (text) {
            let jsonResponse = JSON.parse(text);
            return Promise.reject(new Error(jsonResponse.response));
          } else {
            return Promise.reject(new Error(response.statusText));
          }
        });
      })
      .then((data) => {
        return {
          id: data.id,
          hash: data.hash,
        };
      });
  }

  if (!window.hasOwnProperty('logpaste')) {
    window.logpaste = {};
  }
  window.logpaste.uploadText = uploadText;
})(window);
