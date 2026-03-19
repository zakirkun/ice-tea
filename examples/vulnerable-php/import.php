<?php
// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY
// VULN: XXE Injection (CWE-611)

$result = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $xml_data = $_POST['xml'] ?? '';

    // VULN: DOMDocument with external entity loading enabled (CWE-611)
    // Attacker payload:
    // <?xml version="1.0"?>
    // <!DOCTYPE foo [<!ENTITY xxe SYSTEM "file:///etc/passwd">]>
    // <user><name>&xxe;</name></user>
    
    $dom = new DOMDocument();
    // VULN: LIBXML_NOENT enables entity substitution
    $dom->loadXML($xml_data, LIBXML_NOENT | LIBXML_DTDLOAD);
    
    $names = $dom->getElementsByTagName('name');
    foreach ($names as $name) {
        $result .= htmlspecialchars($name->textContent) . '<br>';
    }
}
?>
<!DOCTYPE html>
<html>
<head><title>XML Import</title></head>
<body>
<h1>XML User Import</h1>
<form method="POST">
    <textarea name="xml" rows="10" cols="60"><?xml version="1.0"?>
<users>
  <user><name>John</name><email>john@example.com</email></user>
</users></textarea><br>
    <input type="submit" value="Import">
</form>
<?php if ($result): ?>
    <h2>Imported names:</h2>
    <p><?= $result ?></p>
<?php endif; ?>
</body>
</html>
