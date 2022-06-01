'use strict';

function displayResult(resultId, key) {
  clearError();
  clearResult();

  const resultDiv = document.getElementById('result');

  let resultUrl = `${document.location}paste/${resultId}`;
  if (key !== '') {
    resultUrl += `?key=${key}`;
  }

  const header = document.createElement('h3');
  header.innerText = 'Shareable link';
  resultDiv.appendChild(header);

  const input = document.createElement('input');
  input.classList.add('result');
  input.value = resultUrl;

  const copy = document.createElement('button');
  copy.classList.add('btn');
  copy.innerHTML = 'Copy';
  copy.addEventListener('click', function () {
    navigator.clipboard.writeText(input.value);
    copy.innerHTML = 'Copied!';
  });

  resultDiv.appendChild(input);
  resultDiv.appendChild(copy);

  resultDiv.style.display = 'block';
}

function clearResult() {
  const resultDiv = document.getElementById('result');
  while (resultDiv.firstChild) {
    resultDiv.removeChild(resultDiv.lastChild);
  }
  resultDiv.style.display = 'none';
}

function clearError() {
  const uploadError = document.getElementById('form-upload-error');
  uploadError.innerText = ' ';
  uploadError.style.display = 'none';
}

function displayError(error) {
  const uploadError = document.getElementById('form-upload-error');
  uploadError.innerText = error;
  uploadError.style.display = 'block';
}

document.getElementById('upload').addEventListener('click', (evt) => {
  const textToUpload = document.getElementById('upload-textarea').value;
  const encryptPaste = document.getElementById('encrypt-paste').checked;
  const expirePaste = document.getElementById('expiry');

  const pasteVal = {
    content: textToUpload,
    encrypt: encryptPaste,
    expire: parseInt(expirePaste.options[expirePaste.selectedIndex].value),
  };

  pastee
    .uploadText(pasteVal)
    .then(({ id, key }) => {
      displayResult(id, key);
    })
    .catch((error) => {
      clearResult();
      displayError(error);
    });
});
