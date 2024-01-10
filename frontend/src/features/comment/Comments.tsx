import * as React from 'react';
import Sidebar from '../../components/Sidebar/Sidebar';
import { PostList } from '../../components/PostList/PostList';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useIsMobileContext } from '../../providers/mobile/isMobile';
import { Box } from '@mui/material';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { Post, Comment } from '../../openapi';
import { CommentComponent } from './CommentComponent';
import { useParams } from 'react-router-dom';


export default function Comments() {
    const isMobile = useIsMobileContext();
    

    return (
    <>
        {!isMobile ? (
            <Box style={{ display: 'flex'}}>
                <Box style={{ flex: '0 0 30%', display: 'flex' }}>
                    <Sidebar />
                </Box>
                <Box style={{ flex: '0 0 40%', overflowX: 'hidden' }}>
                    <CommentComponent />
                </Box>
                <Box style={{ flex: '0 0 30%' }}>

                </Box>
            </Box>
        ): (
        <Box style={{ display: 'flex', justifyContent: "center"}}>
            <Box style={{ flex: '0 0 80%', overflowX: 'hidden' }}>
                <CommentComponent />
            </Box>
        </Box>
    )}
    </>
    );
}