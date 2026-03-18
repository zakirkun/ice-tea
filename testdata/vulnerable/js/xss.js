// Vulnerable Javascript Code with typical DOM XSS
function displayMessage() {
    // Read input from URL fragment
    var message = decodeURIComponent(window.location.hash.substring(1));
    
    // Dangerous sink: Assigns un-sanitized input to innerHTML
    document.getElementById("message-box").innerHTML = message;

    // Another dangerous sink using eval
    eval("console.log('User message: ' + " + message + ")");
}

window.onload = displayMessage;
