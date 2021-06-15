"use strict";

// Call the GET method of the api to refresh messages every second
var intervalId = window.setInterval(() => {
  var lastIndex = parseInt(localStorage.getItem("lastIndex"));

  getChats(lastIndex).then((ok) => {
    let payload = JSON.parse(ok);
    if (payload['Messages']) {
      lastIndex += payload['Messages'].length;

      payload['Messages'].forEach(item => {
        console.log(item);
        appendNewMessage(item);
      });

    }
    localStorage.setItem('lastIndex', lastIndex);
  }).catch((bad) => { console.log(bad); });
}, 1000);

function postMessage() {
  const msg = document.getElementById('message').value;
  const username = localStorage.getItem('username');

  const t = new Date();
  const dateStr = t.getDate() + '-' + (t.getMonth()+1) + '-' + t.getFullYear() + ' '
    + t.getHours() + ':' + t.getMinutes();

  const payload = {
    From: username,
    Content: msg,
    Date: dateStr,
  };

  console.log(payload);

  const myProm = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open('POST', '/API/messages', true);
    xhr.onload = () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        resolve(xhr.response);
      } else {
        reject({
          status: xhr.status,
          statusText: xhr.statusText
        });
      }
    };
    xhr.onerror = () => {
      reject({
        status: xhr.status,
        statusText: xhr.statusText
      });
    };

    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.send(JSON.stringify(payload));
  });

  myProm.then( ok => {
    console.log(ok);
  }).catch( bad => {
    console.log(bad);
  });

  // Clear input
  document.getElementById('message').value = '';
}

function getChats(lastIndex) {
  const params = 'lastIndex='+lastIndex;

  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', '/API/messages?'+params, true);
    xhr.onload = () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        resolve(xhr.response);
      } else {
        reject({
          status: xhr.status,
          statusText: xhr.statusText
        });
      }
    };
    xhr.onerror = () => {
      reject({
        status: xhr.status,
        statusText: xhr.statusText
      });
    };

    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.send(null);
  });
}

function appendNewMessage(message) {
  const text = '(' + message.Date + ') ' + message.From + ': ' + message.Content;
  const parent = document.getElementById('parentWindow');
  const messageContainer = document.createElement('div');
  const msg = document.createElement('p');
  msg.innerText = text;

  messageContainer.appendChild(msg);
  parent.appendChild(messageContainer);
}
