<?php
// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY

$output = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $host = $_POST['host'] ?? '';
    $tool = $_POST['tool'] ?? 'ping';

    // VULN: Command injection — user input directly in shell command (CWE-78)
    if ($tool === 'ping') {
        $output = shell_exec('ping -c 4 ' . $host);
    } elseif ($tool === 'nslookup') {
        $output = shell_exec('nslookup ' . $host);
    } elseif ($tool === 'curl') {
        // VULN: SSRF via curl with no URL validation
        $output = shell_exec('curl -s ' . $host);
    }
}
?>
<!DOCTYPE html>
<html>
<head><title>Network Tools</title></head>
<body>
<h1>Network Tools</h1>
<form method="POST">
    <select name="tool">
        <option value="ping">Ping</option>
        <option value="nslookup">NSLookup</option>
        <option value="curl">Fetch URL</option>
    </select>
    <input type="text" name="host" placeholder="hostname or IP">
    <input type="submit" value="Run">
</form>
<?php if ($output): ?>
    <pre><?= htmlspecialchars($output) ?></pre>
<?php endif; ?>
</body>
</html>
