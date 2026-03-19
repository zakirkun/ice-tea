// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY

// VULN: DOM XSS via innerHTML with URL hash (CWE-79)
// Attacker: http://localhost:3000/#<img src=x onerror=alert(1)>
function loadMessageFromHash() {
    const message = decodeURIComponent(window.location.hash.substring(1));
    // VULN: Unsanitized input assigned to innerHTML
    document.getElementById('message').innerHTML = message;
}

// VULN: eval() with user-controlled string (CWE-94)
function executeUserCode() {
    const code = document.getElementById('code-input').value;
    eval(code);
}

// VULN: document.write with user input
function renderTemplate() {
    const template = new URLSearchParams(window.location.search).get('template');
    document.write(template);
}

// VULN: postMessage to * (any origin) — insecure cross-origin messaging (CWE-942)
function sendMessageToParent(data) {
    window.parent.postMessage(data, '*');
}

// VULN: localStorage used to store JWT token (CWE-312)
function storeToken(token) {
    localStorage.setItem('jwt_token', token);
}

// VULN: Reading token from localStorage without sanitization
function getStoredToken() {
    return localStorage.getItem('jwt_token');
}

// VULN: Using Math.random() for security-sensitive OTP (CWE-338)
function generateOTP() {
    const otp = Math.floor(Math.random() * 1000000);
    return otp.toString().padStart(6, '0');
}

window.addEventListener('load', () => {
    loadMessageFromHash();

    // VULN: Insecure postMessage listener without origin check (CWE-346)
    window.addEventListener('message', (event) => {
        // No origin check!
        document.getElementById('received').innerHTML = event.data;
    });
});
