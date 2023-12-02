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
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import { TreeView } from '@mui/x-tree-view/TreeView';
import { TreeItem } from '@mui/x-tree-view/TreeItem';
import { BookGenre } from '../../openapi/models';
import { IconButton } from '@mui/material';


type bookGenreNode = {
    current: BookGenre,
    children: bookGenreNode[],
}

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
    const [bookGenreNodes, setBookGenreNodes] = React.useState<bookGenreNode[]>([]);

    React.useEffect(() => {
        const fetchBookGenres = async (bookGenreID: string) => {
        
            try {
                const api = await apiInstance;
                const res = await api.getBooksGenres(bookGenreID);
                
                if (res.data && Array.isArray(res.data)) {
                    const bookGenreNodes: bookGenreNode[] = res.data.map(item => ({
                        current: item,
                        children: [],
                    }));
                    setBookGenreNodes(bookGenreNodes);
                }
            } catch (error: unknown) {
                errorHandler(error);
            }
                
        };
    
        fetchBookGenres("001");
      }, []);
    
    const appendChildren = async (nodes: bookGenreNode) => {
        console.log(99);
        if (nodes.children.length > 0) {
            console.log(100);
            console.log(bookGenreNodes)
            return;
        }
        console.log(101);
        try {
            const api = await apiInstance;
            const res = await api.getBooksGenres(nodes.current.books_genre_id);
            
            if (res.data && Array.isArray(res.data)) {
                
                const newChildrens: bookGenreNode[] = res.data.map(item => ({
                    current: item,
                    children: [],
                }));
                
                const newTree = addChildrenToTree(bookGenreNodes, nodes.current.id, newChildrens);
                setBookGenreNodes(newTree); // 新しい状態を設定
            }
        } catch (error: unknown) {
            errorHandler(error);
        }
    };

    const addChildrenToTree = (nodes: bookGenreNode[], targetId: number, newChildren: bookGenreNode[]): bookGenreNode[] => {
        return nodes.map(node => {
            if (node.current.id === targetId) {
                return {...node, children: [...node.children, ...newChildren]};
            } else if (node.children) {
                return {...node, children: addChildrenToTree(node.children, targetId, newChildren)};
            } else {
                return node;
            }
        });
    };


    const renderTree = (nodes: bookGenreNode) => (
        <TreeItem key={nodes.current.id} nodeId={String(nodes.current.genre_level-1)} label={
            <Box sx={{display: 'flex', alignItems: 'center', justifyContent: 'start'}}>
                <IconButton onClick={() => appendChildren(nodes)}>
                    <ChevronRightIcon/>
                </IconButton>
                <Box sx={{width: '100%'}}>
                    <Button variant="text" sx={{ width: '100%', justifyContent: 'flex-start', color: 'black' }}>
                        {nodes.current.books_genre_name}
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
        <Box sx={{mt: '2rem'}}>
            <TreeView 
                disableSelection={true}
            >
                {bookGenreNodes.map((data) => renderTree(data))}
            </TreeView>
        </Box>
        <BookList books={books} />
    </Container>
  );
}