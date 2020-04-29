let title;
let url;

function apiAddNotes() {
  let quote = document.getElementById('output').textContent;
  let note = document.getElementById('comments').value;

  let ipaddr = "555";
  fetch('https://www.cloudflare.com/cdn-cgi/trace').then((response) => {
    response.text().then(function (data) {
      ipaddr = data.match(/ip=(.+)/)[1]

      fetch("https://us-central1-lancelot-274021.cloudfunctions.net/createNote", {
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(
          {
            "user_id": ipaddr,
            "comment": quote,
            "article": note,
            "url": url
          })
      }) // Call the fetch function passing the url of the API as a parameter
        .then(function () {
          document.getElementById("fireAddNotes").textContent = "Added!"
        })
        .catch(function (err) {
          // This is where you run code if the server returns any errors
          console.log("error: ", err)
        });
    })
  });
}


document.addEventListener("DOMContentLoaded", () => {
  document.getElementById("fireAddNotes").addEventListener("click", apiAddNotes);

  chrome.tabs.query({ active: true, currentWindow: true }, function (tab) {
    url = tab[0].url;
    let tabTitle = tab[0].title;

    document.getElementById("app-name").textContent = tabTitle;

    const newUrl = new URL(url);
    const urlParams = new URLSearchParams(newUrl.search);
    const ytId = urlParams.get('v');
    // const thumbnail = `https://img.youtube.com/vi/${ytId}/maxresdefault.jpg`
    const embed = `https://www.youtube.com/embed/${ytId}`

    if (url && url.includes("youtube")) {
      chrome.tabs.executeScript(tab.id,
        {
          code: "yt = document.querySelector('.video-stream').currentTime; yt;"
        },
        function (data) {
          const sec = Math.floor(data[0]);
          url = `${url}&t=${sec}`;
          var yt = document.createElement("iframe");
          var linkText = document.createTextNode(url);
          yt.appendChild(linkText)
          yt.setAttribute("src", embed+`?t=${sec}`)
          yt.classList.add("popup");
          document.getElementById("output").hidden = true;
          document.getElementById("video-link").appendChild(yt);
        }
      );
    } else {
      chrome.tabs.executeScript(tab.id,
        {
          code: "window.getSelection().toString();"
        },
        function (selection) {
          if (!selection[0] || selection[0] === '') {
            document.getElementById("output").textContent = 'Highlight text for it to appear here!';
          } else {
            document.getElementById("output").textContent = selection[0];
          }
        }
      );
    };
  });
});
