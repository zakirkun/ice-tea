<?php
// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY

$db = new PDO('sqlite:users.db');
$query_term = $_GET['q'] ?? '';

// VULN: SQL Injection (CWE-89)
$sql = "SELECT id, username, email FROM users WHERE username LIKE '%" . $query_term . "%'";

$results = [];
try {
    $stmt = $db->query($sql);
    $results = $stmt->fetchAll(PDO::FETCH_ASSOC);
} catch (Exception $e) {
    die('Error: ' . $e->getMessage());
}
?>
<!DOCTYPE html>
<html>
<head><title>Search</title></head>
<body>
<form method="GET">
    <input type="text" name="q" value="<?php echo $_GET['q'] ?? ''; ?>">
    <!-- VULN: XSS — unescaped output (CWE-79) -->
    <input type="submit" value="Search">
</form>

<h2>Results for: <?php echo $query_term; /* VULN: Reflected XSS — no htmlspecialchars */ ?></h2>

<?php foreach ($results as $row): ?>
    <div>
        <!-- VULN: XSS in output -->
        <strong><?php echo $row['username']; ?></strong> - <?php echo $row['email']; ?>
    </div>
<?php endforeach; ?>
</body>
</html>
