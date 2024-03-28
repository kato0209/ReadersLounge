'use client';
import * as React from 'react';
import Container from '@mui/material/Container';
import { z } from 'zod';
import { Book } from '../../openapi';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Link from '@mui/material/Link';
import { CreatePost } from '../home/CreatePost';
import { PostSchema } from '../../types/PostSchema';

type PostFormData = z.infer<typeof PostSchema>;

type BookListProps = {
  books: Book[];
  formData?: PostFormData;
};

export const BookList: React.FC<BookListProps> = ({ books, formData }) => {
  return (
    <Container
      component="main"
      sx={{
        marginTop: '1rem',
        '@media (max-width: 500px)': {
          padding: '0',
        },
      }}
    >
      {books.length > 0 ? (
        <>
          {books.map((book) => (
            <Card
              sx={{
                display: 'flex',
                justifyContent: 'space-between',
                width: '90%',
                minWidth: '600',
                backgroundColor: '#EFEBE5',
                boxShadow: 'none',
                borderTop: '1px solid #BDBDBD',
                borderRight: '1px solid #BDBDBD',
                borderLeft: '1px solid #BDBDBD',
                borderBottom: '1px solid #BDBDBD',
                cursor: 'pointer',
                '&:hover': {
                  color: 'inherit',
                  backgroundColor: '#EAE6E0',
                },
                '@media (max-width: 500px)': {
                  width: '100%',
                },
              }}
              key={book.ISBNcode}
            >
              <Box sx={{ display: 'flex', flexDirection: 'column', flex: 1 }}>
                <CardContent sx={{ flex: '1 0 auto' }}>
                  <Link href={book.item_url} underline="hover">
                    <Typography
                      component="div"
                      sx={{
                        fontSize: '1.5rem',
                        '@media (max-width: 500px)': { fontSize: '1.0rem' },
                      }}
                    >
                      {book.title}
                    </Typography>
                  </Link>
                  <Typography
                    variant="subtitle1"
                    color="text.secondary"
                    component="div"
                  >
                    著者：{book.author}
                  </Typography>
                  <Typography
                    variant="subtitle1"
                    color="text.secondary"
                    component="div"
                  >
                    出版社：{book.publisher}, 出版日：{book.published_at}
                  </Typography>
                  <Box sx={{ display: 'flex', alignItems: 'center' }}>
                    <Typography
                      component="div"
                      sx={{
                        color: 'red',
                        fontSize: '1.5rem',
                        marginRight: '0.3rem',
                      }}
                    >
                      {book.price}円
                    </Typography>
                    <span style={{ color: 'gray' }}>(税込み)</span>
                  </Box>
                </CardContent>
                <Box
                  sx={{
                    display: 'flex',
                    alignItems: 'center',
                    pl: 1,
                    pb: 1,
                    width: '150px',
                    '@media (max-width: 500px)': {
                      width: '120px',
                    },
                  }}
                >
                  <CreatePost
                    displayString="本の感想を投稿"
                    book={book}
                    formData={formData}
                  />
                </Box>
              </Box>
              <Box
                sx={{
                  width: '30%',
                  margin: '1rem',
                  display: 'flex',
                  justifyContent: 'flex-end',
                  '@media (max-width: 500px)': {
                    margin: '0.2rem',
                    width: '35%',
                  },
                }}
              >
                <CardMedia
                  component="img"
                  sx={{
                    width: '60%',
                    '@media (max-width: 500px)': {
                      width: '100%',
                    },
                  }}
                  image={book.image}
                />
              </Box>
            </Card>
          ))}
        </>
      ) : (
        <></>
      )}
    </Container>
  );
};
