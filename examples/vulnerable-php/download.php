<?php
// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY

// VULN: Path traversal — no sanitization of filename (CWE-22)
// Attacker: ?file=../../../../etc/passwd
$filename = $_GET['file'] ?? '';
$base_dir = '/var/www/html/downloads/';
$full_path = $base_dir . $filename;

if (file_exists($full_path)) {
    header('Content-Type: application/octet-stream');
    header('Content-Disposition: attachment; filename="' . basename($filename) . '"');
    readfile($full_path);
} else {
    // VULN: Exposes full path in error message (CWE-209)
    die('File not found: ' . $full_path);
}
