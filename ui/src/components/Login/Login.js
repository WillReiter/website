import React, { useState } from 'react';
import { TextField, Button, Box, Typography, Container } from '@mui/material';

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log('Username:', username);
        console.log('Password:', password);
        // Add authentication logic here
    };

    return (
        <Container maxWidth="sm">
            <Box
                component="form"
                onSubmit={handleSubmit}
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    justifyContent: 'center',
                    height: '100vh',
                }}
            >
                <Typography variant="h4" gutterBottom>
                    Login
                </Typography>
                <TextField
                    label="Username"
                    variant="outlined"
                    margin="normal"
                    fullWidth
                    required
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <TextField
                    label="Password"
                    type="password"
                    variant="outlined"
                    margin="normal"
                    fullWidth
                    required
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <Button
                    type="submit"
                    variant="contained"
                    color="primary"
                    sx={{ marginTop: 2 }}
                    fullWidth
                >
                    Login
                </Button>
            </Box>
        </Container>
    );
};

export default Login;
