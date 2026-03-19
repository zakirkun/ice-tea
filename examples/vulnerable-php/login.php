<?php
// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY

session_start();
// VULN: No session_regenerate_id() after login (CWE-384)

$db = new PDO('sqlite:users.db');
$db->exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, email TEXT, role TEXT DEFAULT 'user')");

$error = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $username = $_POST['username'];
    $password = $_POST['password'];

    // VULN: SQL Injection via string concatenation (CWE-89)
    $query = "SELECT * FROM users WHERE username = '" . $username . "' AND password = '" . $password . "'";
    
    try {
        $stmt = $db->query($query);
        $user = $stmt->fetch(PDO::FETCH_ASSOC);

        if ($user) {
            // VULN: No session regeneration (CWE-384)
            $_SESSION['user_id'] = $user['id'];
            $_SESSION['username'] = $user['username'];
            $_SESSION['role'] = $user['role'];

            // VULN: Cookie without Secure, HttpOnly, SameSite (CWE-614)
            setcookie('remember_user', $user['username'], time() + 86400);

            header('Location: dashboard.php');
            exit;
        } else {
            $error = 'Invalid username or password';
        }
    } catch (PDOException $e) {
        // VULN: Full SQL error exposed to user (CWE-209)
        $error = 'Database error: ' . $e->getMessage() . ' (Query: ' . $query . ')';
    }
}
?>
<!DOCTYPE html>
<html>
<head><title>Login</title></head>
<body>
<h1>Login</h1>
<?php if ($error): ?>
    <p style="color:red"><?= $error ?></p>
<?php endif; ?>
<form method="POST" action="login.php">
    <!-- VULN: No CSRF token (CWE-352) -->
    <label>Username: <input type="text" name="username"></label><br>
    <label>Password: <input type="password" name="password"></label><br>
    <input type="submit" value="Login">
</form>
</body>
</html>
