import React from 'react';
import { Box, Typography } from '@mui/material';
import ReadersLoungeLogo from '../../assets/images/ReadersLounge-logo-book.png';

function PortalLogo() {
    return (
        <Box
            sx={{
                display: 'flex',         
                flexDirection: 'column',  
                justifyContent: 'center',
                alignItems: 'center',     
                textAlign: 'center',     
                height: '100%',          
            }}
        >
            <Box component="img" src={ReadersLoungeLogo} alt="ReadersLounge Logo"/>
            <Typography component="h1" variant="h3" style={{ color: '#FF7E73' }}>
                ReadersLounge
            </Typography>
        </Box>
    );
}

export default PortalLogo;