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
import { set, useForm } from 'react-hook-form';
import { Book } from '../../openapi';
import { BookList } from './BookList';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import { TreeView } from '@mui/x-tree-view/TreeView';
import { TreeItem } from '@mui/x-tree-view/TreeItem';
import { BookGenreNode } from '../../openapi/models';
import { IconButton } from '@mui/material';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import Chip from '@mui/material/Chip';
import { PostSchema } from '../../types/PostSchema';

const searchBookSchema = z.object({
    keyword: z.string().optional(),
    bookGenreID: z.string().optional(),
});

type FormData = z.infer<typeof searchBookSchema>;

type PostFormData = z.infer<typeof PostSchema>;
type SearchBookProps  = {
  formData?: PostFormData;
};



export const SearchBook: React.FC<SearchBookProps> = ({ formData }) => {
    const { register, handleSubmit, setValue, getValues, formState: { errors } } = useForm<FormData>({
        resolver: zodResolver(searchBookSchema),
    });
    const errorHandler = useErrorHandler();
    const [books, setBooks] = React.useState<Book[]>([]);
    const [bookGenreNodes, setBookGenreNodes] = React.useState<BookGenreNode[]>([]);
    const [hasGenre, setHasGenre] = React.useState<boolean>(false);
    const [selectedGenre, setSelectedGenre] = React.useState<string>("");
    const [bookNotFound, setBookNotFound] = React.useState<boolean>(false);

    React.useEffect(() => {
        const fetchBookGenres = async () => {
        
            try {
                const api = await apiInstance;
                const res = await api.getBooksGenres();
                
                if (res.data && Array.isArray(res.data)) {
                    setBookGenreNodes(res.data);
                } 
            } catch (error: unknown) {
                errorHandler(error);
            }
                
        };
    
        fetchBookGenres();
      }, []);

    const renderTree = (nodes: BookGenreNode) => (
        <TreeItem key={nodes.id} nodeId={String(nodes.id)} label={
            <Box sx={{display: 'flex', alignItems: 'center', justifyContent: 'start'}}>
                <Box sx={{width: '100%'}}>
                    <Button 
                        variant="text" 
                        sx={{ 
                            width: '100%', 
                            justifyContent: 'flex-start', 
                            color: 'black' 
                        }}
                        onClick={() => handleGenreSelect(nodes.books_genre_name, nodes.books_genre_id)}
                    >
                        {nodes.books_genre_name}
                    </Button>
                </Box>
            </Box>
            
        }>
            {Array.isArray(nodes.children)
        ? nodes.children.map((node) => renderTree(node))
        : null}
        </TreeItem>
        
    );
    
    
    const onSubmit = async (data: FormData) => {
        try {
            const api = await apiInstance;
            const res = await api.fetchBookData(data.bookGenreID, data.keyword);
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
                setHasGenre(false);
                if (res.data.length === 0) {
                    setBookNotFound(true);
                }
            }
            
        } catch (error: unknown) {
            errorHandler(error);
        }
    }

    const handleGenreDisplay = () => {
        setHasGenre(!hasGenre);
    };

    const handleGenreSelect = (bookGenreName: string, bookGenreID: string) => {
        setSelectedGenre(bookGenreName);
        setValue('bookGenreID', bookGenreID, { shouldValidate: true });
    }

    const handleGenreDelete = () => {
        setSelectedGenre("");
        setValue('bookGenreID', undefined, { shouldValidate: true });
    }

  return (
    <Container component="main">
        <Box component="form" onSubmit={handleSubmit(onSubmit)} noValidate sx={{mt: '1rem'}}>
            <Box sx={{display: 'flex'}}>
                <TextField
                    {...register("keyword")}
                    fullWidth
                    id="keyword"
                    label="本のタイトル"
                    name="keyword"
                />
                <Button 
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
        <Box sx={{mt: '2rem'}}>
            <Box sx={{display: 'flex', textAlign: 'center'}}>
                <Button 
                    sx={{
                        backgroundColor: '#FF7E73',
                        color: '#fff',
                        marginBottom: '1rem',
                        marginRight: '1rem',
                        '&:hover': {
                        backgroundColor: '#E56A67',
                        },
                    }}
                    onClick={handleGenreDisplay}
                >
                    本のジャンルを選択 {hasGenre ? <KeyboardArrowUpIcon/> : <KeyboardArrowDownIcon/>}
                </Button>
                {selectedGenre && <Chip label={selectedGenre} variant="outlined" onDelete={handleGenreDelete} />}
            </Box>
            {hasGenre &&
                <TreeView 
                    defaultCollapseIcon={<ExpandMoreIcon />}
                    defaultExpandIcon={<ChevronRightIcon />}
                >
                    {bookGenreNodes.map((data) => renderTree(data))}
                </TreeView>
            } 
        </Box>
        <BookList books={books} formData={formData} />
        {bookNotFound && 
        <Typography component="div" sx={{fontSize: '1.5rem', marginTop: '1rem'}}>
            該当するものはありません
        </Typography>
        }
    </Container>
  );
}