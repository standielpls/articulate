async function GetIPAddress() {
  const trace = await fetch("https://www.cloudflare.com/cdn-cgi/trace");
  const traceText = await trace.text();
  return traceText.match(/ip=(.+)/)[1];
}

async function PostNote(quote, note, url, address) {
  const response = await fetch(
    "https://us-central1-lancelot-274021.cloudfunctions.net/createNote",
    {
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      method: "POST",
      body: JSON.stringify({
        user_id: address,
        comment: quote,
        article: note,
        url: url,
      }),
    }
  );
  if (response.ok) {
    SaveAddNoteButton("Added!");
  }
}

async function AddNotes(quote, note, url) {
  const originAddress = await GetIPAddress();
  try {
    await PostNote(quote, note, url, originAddress);
  } catch (error) {
    console.log(error);
  }
}
