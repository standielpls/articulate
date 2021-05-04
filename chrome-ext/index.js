const PLACEHOLDER_OUTPUT = "Highlight text for it to appear here!";

function processText(tab) {
  chrome.tabs.executeScript(
    tab.id,
    {
      code: "window.getSelection().toString();",
    },
    function (selection) {
      if (!selection[0] || selection[0] === "") {
        SaveOutputText(PLACEHOLDER_OUTPUT);
      } else {
        SaveOutputText(selection[0]);
      }
    }
  );
}

function processVideo(tab, url) {
  const urlObj = new URL(url);
  const ytId = urlObj.searchParams.get("v");
  const thumbnailURL = `https://img.youtube.com/vi/${ytId}/maxresdefault.jpg`;

  chrome.tabs.executeScript(
    tab.id,
    {
      code: "yt = document.querySelector('.video-stream').currentTime; yt;",
    },
    function (data) {
      const sec = Math.floor(data[0]);
      url = `${url}&t=${sec}`;
      SaveOutputVideo({
        src: thumbnailURL,
        url: url,
      });
    }
  );
}

function render(url) {
  const quoteText = GetOutputText().textContent;
  const noteText = GetNote().value;
  const addNotesButton = GetAddNoteButton();
  addNotesButton.addEventListener(
    "click",
    AddNotes({ quote: quoteText, note: noteText, url })
  );
}

function load() {
  let url;
  document.addEventListener("DOMContentLoaded", () => {
    chrome.tabs.query({ active: true, currentWindow: true }, (tab) => {
      url = tab[0].url;
      let tabTitle = tab[0].title;

      document.getElementById("app-name").textContent = tabTitle;
      return url && url.includes("youtube")
        ? processVideo(tab, url)
        : processText(tab);
    });

    render(url);
  });
}
load();
