<?php
// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY

$db = new PDO('sqlite:users.db');
$message = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $username = $_POST['username'];
    $password = $_POST['password'];
    $email    = $_POST['email'];

    // VULN: MD5 password hashing without salt (CWE-916, CWE-328)
    $hashed = md5($password);

    // VULN: SQL injection via string concatenation (CWE-89)
    $sql = "INSERT INTO users (username, password, email) VALUES ('" . $username . "', '" . $hashed . "', '" . $email . "')";

    try {
        $db->exec($sql);
        $message = 'Registered successfully!';
    } catch (Exception $e) {
        // VULN: DB error exposed to user (CWE-209)
        $message = 'Error: ' . $e->getMessage();
    }
}
?>
<!DOCTYPE html>
<html>
<head><title>Register</title></head>
<body>
<h1>Register</h1>
<form method="POST">
    <!-- VULN: No CSRF token (CWE-352) -->
    <input type="text" name="username" placeholder="Username"><br>
    <input type="email" name="email" placeholder="Email"><br>
    <input type="password" name="password" placeholder="Password"><br>
    <input type="submit" value="Register">
</form>
<?php if ($message): echo "<p>$message</p>"; endif; ?>
</body>
</html>
