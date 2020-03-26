fetch("http://localhost:5000/notes") // Call the fetch function passing the url of the API as a parameter
  .then(function(response) {
    console.log(response.status)
    if (response.status !== 200) {
      console.log("Looks like there was a problem. Status Code: " + response.status);
      return;
    }
    // Examine the text in the response
    response.json().then(function(notes) {
        for (note of notes) {
            var parent = document.getElementById("content")
            var aNode = document.createElement("a")
            aNode.href = note['url']
            aNode.innerHTML = note['url']
            parent.appendChild(aNode)

            var div = document.createElement("div");
            var headerEle = document.createElement("blockquote");
            var quoteNode = document.createTextNode(note['quote']);
            headerEle.appendChild(quoteNode)
            div.appendChild(headerEle)

            
            var noteEle = document.createElement("p");
            var noteNode = document.createTextNode(note['note']);
            noteEle.appendChild(noteNode)
            div.appendChild(noteEle)
            
            parent.appendChild(div)
        }
        });
  })
  .catch(function(err) {
    // This is where you run code if the server returns any errors
    console.log('Fetch Error :-S', err);
  });

