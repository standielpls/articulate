let title;
let url;

function apiAddNotes() {
  let quote = document.getElementById('output').textContent
  let note = document.getElementById('comments').value
  
  fetch("https://us-central1-lancelot-274021.cloudfunctions.net/createNote", {
      headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
      },
      method: "POST",
      body: JSON.stringify({comment: quote, article: note, url: url, time: Date.now()})
  }) // Call the fetch function passing the url of the API as a parameter
    .then(function() {
      console.log("hello")
    })
    .catch(function(err) {
      // This is where you run code if the server returns any errors
      console.log("error: ", err)
    });
}

document.addEventListener("DOMContentLoaded", function() {
  chrome.tabs.query({ active: true, currentWindow: true }, function(tabs) {
    title = tabs[0].title;
    url = tabs[0].url;
    document.getElementById("app-name").textContent = title
  });

  document
    .getElementById("fireAddNotes")
    .addEventListener("click", apiAddNotes);

  chrome.tabs.executeScript(
    {
      code: "window.getSelection().toString();"
    },
    function(selection) {
      document.getElementById("output").textContent = selection[0];
    }
  );
});
