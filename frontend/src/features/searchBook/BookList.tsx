import * as React from 'react';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import { z } from 'zod';
import { Book } from '../../openapi';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';

type BookListProps  = {
    books: Book[];
};

export const BookList: React.FC<BookListProps> = ({ books }) => {

  return (
    <Container component="main">
        <List>
            {books.map((book) => (
            <ListItem key={book.ISBNcode}>
                <ListItemText primary={book.title} secondary={book.author} />
            </ListItem>
            ))}
      </List>
    </Container>
  );
}