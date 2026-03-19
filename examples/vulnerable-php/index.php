<?php
// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY

// VULN: debug mode — display_errors on (CWE-209)
ini_set('display_errors', '1');
error_reporting(E_ALL);

// VULN: Local File Inclusion (CWE-98)
// Attacker: ?page=../../../../etc/passwd%00
$page = $_GET['page'] ?? 'home';

// No sanitization — direct file inclusion
include($page . '.php');
?>
