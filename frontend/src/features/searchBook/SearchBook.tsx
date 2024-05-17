'use client';
import * as React from 'react';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import { z } from 'zod';
import { BookList } from './BookList';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import { TreeView } from '@mui/x-tree-view/TreeView';
import { TreeItem } from '@mui/x-tree-view/TreeItem';
import { BookGenreNode } from '../../openapi/models';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import Chip from '@mui/material/Chip';
import { PostSchema } from '../../types/PostSchema';
import { useFormState } from 'react-dom';
import { searchBook } from './SearchBookAction';
import { State } from './SearchBookAction';
import { useErrorHandler } from 'react-error-boundary';
import axios from 'axios';

const initialState: State = {
  error: '',
  fieldErrors: {
    keyword: '',
    bookGenreID: '',
  },
  fetchedBooks: [],
  bookNotFound: false,
};

type PostFormData = z.infer<typeof PostSchema>;
type SearchBookProps = {
  formData?: PostFormData;
};

export const SearchBook: React.FC<SearchBookProps> = ({ formData }) => {
  const [state, formAction] = useFormState(searchBook, initialState);
  const [hasGenre, setHasGenre] = React.useState<boolean>(false);
  const [selectedGenre, setSelectedGenre] = React.useState<string>('');
  const [bookGenreID, setBookGenreID] = React.useState<string | undefined>(
    undefined,
  );
  const [bookGenreNodes, setBookGenreNodes] = React.useState<BookGenreNode[]>(
    [],
  );
  const errorHandler = useErrorHandler();

  async function fetchBookGenres() {
    try {
      const res = await axios.get(`/api/fetch-book-genres`);
      return res.data;
    } catch (error: unknown) {
      errorHandler(error);
    }
  }
  React.useEffect(() => {
    fetchBookGenres().then((data) => {
      setBookGenreNodes(data);
    });
  }, []);

  const handleGenreDisplay = () => {
    setHasGenre(!hasGenre);
  };

  const handleGenreSelect = (bookGenreName: string, bookGenreID: string) => {
    setSelectedGenre(bookGenreName);
    setBookGenreID(bookGenreID);
  };

  const handleGenreDelete = () => {
    setSelectedGenre('');
    setBookGenreID(bookGenreID);
  };

  const renderTree = (nodes: BookGenreNode) => (
    <TreeItem
      key={nodes.id}
      nodeId={String(nodes.id)}
      label={
        <Box
          sx={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'start',
          }}
        >
          <Box sx={{ width: '100%' }}>
            <Button
              variant="text"
              sx={{
                width: '100%',
                justifyContent: 'flex-start',
                color: 'black',
              }}
              onClick={() =>
                handleGenreSelect(nodes.books_genre_name, nodes.books_genre_id)
              }
            >
              {nodes.books_genre_name}
            </Button>
          </Box>
        </Box>
      }
    >
      {Array.isArray(nodes.children)
        ? nodes.children.map((node) => renderTree(node))
        : null}
    </TreeItem>
  );

  return (
    <Container component="main" sx={{ mt: 2 }}>
      <form action={formAction}>
        <Box sx={{ display: 'flex' }}>
          <TextField
            fullWidth
            id="keyword"
            label="本のタイトル"
            name="keyword"
          />
          <input
            type="hidden"
            id="bookGenreID"
            name="bookGenreID"
            value={bookGenreID}
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
              },
            }}
          >
            検索
          </Button>
        </Box>
      </form>
      <Box sx={{ mt: '2rem' }}>
        <Box sx={{ display: 'flex', textAlign: 'center' }}>
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
            本のジャンルを選択{' '}
            {hasGenre ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
          </Button>
          {selectedGenre && (
            <Chip
              label={selectedGenre}
              variant="outlined"
              onDelete={handleGenreDelete}
            />
          )}
        </Box>
        {hasGenre && (
          <TreeView
            defaultCollapseIcon={<ExpandMoreIcon />}
            defaultExpandIcon={<ChevronRightIcon />}
          >
            {bookGenreNodes.map((data) => renderTree(data))}
          </TreeView>
        )}
      </Box>
      <BookList books={state.fetchedBooks ?? []} formData={formData} />
      {state.bookNotFound && (
        <Typography
          component="div"
          sx={{ fontSize: '1.5rem', marginTop: '1rem' }}
        >
          該当するものはありません
        </Typography>
      )}
    </Container>
  );
};
