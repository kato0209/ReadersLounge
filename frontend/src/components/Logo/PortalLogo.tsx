import React from 'react';
import { Box, Typography } from '@mui/material';
import ReadersLoungeLogo from '../../assets/images/ReadersLounge-logo-book.png';

function PortalLogo() {
    return (
        <Box>
            <Box component="img" src={ReadersLoungeLogo} alt="ReadersLounge Logo"/>
            <Typography component="h1" variant="h3" style={{ color: '#FF7E73' }}>
                ReadersLounge
            </Typography>
        </Box>
    );
}

export default PortalLogo;