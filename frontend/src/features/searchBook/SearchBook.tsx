import * as React from 'react';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { useForm } from 'react-hook-form';
import { Book } from '../../openapi';
import { BookList } from './BookList';

const searchBookSchema = z.object({
    keyword: z.string()
});

type FormData = z.infer<typeof searchBookSchema>;


export default function SearchBook() {
    const { register, handleSubmit, setValue, getValues, formState: { errors } } = useForm<FormData>({
        resolver: zodResolver(searchBookSchema),
    });
    const errorHandler = useErrorHandler();
    const [books, setBooks] = React.useState<Book[]>([]);
    
    const onSubmit = async (data: FormData) => {
        try {
            const api = await apiInstance;
            const res = await api.fetchBookData(data.keyword);
            if (res.data && Array.isArray(res.data)) {
                const fetchedBooks: Book[] = res.data.map(item => ({
                    book_id: item.book_id,
                    ISBNcode: item.ISBNcode,
                    title: item.title,
                    author: item.author,
                    price: item.price,
                    publisher: item.publisher,
                    published_at: item.published_at,
                    item_url: item.item_url,
                    image: item.image,
                }));
                setBooks(fetchedBooks);
            }
            
        } catch (error: unknown) {
            errorHandler(error);
        }
    }

  return (
    <Container component="main">
        <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{mt: '1rem'}}>
            <Box sx={{display: 'flex'}}>
                <TextField
                    required
                    fullWidth
                    id="serchBook"
                    label="本のタイトル"
                    name="serchBook"
                />
                <Button 
                    {...register("keyword")}
                    type="submit"
                    sx={{
                    backgroundColor: '#FF7E73',
                    color: '#fff',
                    '&:hover': {
                    backgroundColor: '#E56A67',
                    },
                    '&.Mui-disabled': {
                        backgroundColor: '#FFA49D',
                        color: '#fff',
                    }
                    }}
                >
                送信
                </Button>
            </Box>
        </Box>
        <BookList books={books} />
    </Container>
  );
}