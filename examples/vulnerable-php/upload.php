<?php
// INTENTIONALLY VULNERABLE — FOR SECURITY TESTING ONLY

$message = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST' && isset($_FILES['file'])) {
    $file = $_FILES['file'];

    // VULN: Only checks Content-Type header (user-controlled) (CWE-434)
    $allowed_types = ['image/jpeg', 'image/png', 'image/gif'];
    
    if (in_array($file['type'], $allowed_types)) {
        // VULN: Preserves original filename — allows shell.php to be uploaded (CWE-434)
        $upload_dir = 'uploads/';
        $destination = $upload_dir . $file['name'];

        // VULN: Stores file in web-accessible directory (CWE-434)
        if (move_uploaded_file($file['tmp_name'], $destination)) {
            $message = 'File uploaded: <a href="' . $destination . '">' . $file['name'] . '</a>';
        } else {
            $message = 'Upload failed';
        }
    } else {
        $message = 'Invalid file type: ' . $file['type'];
    }
}
?>
<!DOCTYPE html>
<html>
<head><title>File Upload</title></head>
<body>
<h1>Upload File</h1>
<form method="POST" enctype="multipart/form-data">
    <!-- VULN: No CSRF token -->
    <input type="file" name="file">
    <input type="submit" value="Upload">
</form>
<?php if ($message): ?>
    <p><?= $message ?></p>
<?php endif; ?>
</body>
</html>
