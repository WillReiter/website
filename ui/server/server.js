const express = require('express');
const path = require('path');
const app = express();
const PORT = process.env.PORT || 5000;

// Serve static files from the React build folder
app.use(express.static(path.join(__dirname, '../dist')));  // or '../build' for CRA

// API route (Example)
app.get('/api/hello', (req, res) => {
    res.json({ message: 'Hello from server!' });
});

// Catch-all handler to serve React app for unknown routes
app.get('*', (req, res) => {
    res.sendFile(path.join(__dirname, '../dist', 'index.html'));  // or '../build'
});

// Start server
app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
});