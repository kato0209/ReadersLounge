'use client';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import UserAvatar from '../../components/Avatar/UserAvatar';
import { redirect } from 'next/navigation';
import { State } from './UserSearchAction';
import { useFormState } from 'react-dom';
import { searchUser } from './UserSearchAction';

const initialState: State = {
  error: '',
  fieldErrors: {
    keyword: '',
  },
  users: [],
  userNotFound: false,
};

export default function UserSearchSection() {
  const [state, formAction] = useFormState(searchUser, initialState);

  return (
    <Container component="main">
      <form action={formAction}>
        <Box sx={{ display: 'flex' }}>
          <TextField fullWidth id="keyword" label="Search" name="keyword" />
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
      <Box sx={{ marginTop: '1rem' }}>
        {state.users.map((user) => (
          <Box
            key={user.user_id}
            onClick={() => redirect(`/user-profile/${user.user_id}`)}
            sx={{
              display: 'flex',
              alignItems: 'center',
              padding: '0.5rem',
              cursor: 'pointer',
              '&:hover': {
                color: 'inherit',
                backgroundColor: '#EAE6E0',
              },
            }}
          >
            <Box sx={{ marginRight: '1rem' }}>
              <UserAvatar image={user.profile_image} userID={user.user_id} />
            </Box>
            <Box>
              <Typography variant="body1">{user.name}</Typography>
            </Box>
          </Box>
        ))}
      </Box>
      {state.userNotFound && (
        <Typography
          component="div"
          sx={{ fontSize: '1.5rem', marginTop: '1rem' }}
        >
          該当するユーザーがいません
        </Typography>
      )}
    </Container>
  );
}
