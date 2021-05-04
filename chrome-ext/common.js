function GetOutputText() {
  return document.getElementById("output-text");
}
function GetNote() {
  return document.getElementById("comments");
}
function GetAddNoteButton() {
  return document.getElementById("add-notes");
}
function GetOutputVideo() {
  return document.getElementById("output-video");
}
function GetOutputVideoLink() {
  return document.getElementById("output-video-link");
}

function SaveOutputText(output) {
  GetOutputText().textContent = output;
}

function HideOutputPlaceholder() {}

function SaveOutputVideo(output) {
  const outputVideo = GetOutputVideo();
  outputVideo.setAttribute("src", output.src);

  const outputVideoLink = GetOutputVideoLink();
  outputVideoLink.setAttribute("href", output.url);
  outputVideoLink.setAttribute("target", output.url);
  outputVideoLink.style.visibility = "visible";

  GetOutputText().style.visibility = "hidden";
}

function SaveAddNoteButton(text) {
  GetAddNoteButton().setAttribute("textContent", text);
}
