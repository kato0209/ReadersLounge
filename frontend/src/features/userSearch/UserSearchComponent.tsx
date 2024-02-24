import * as React from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { apiInstance } from '../../lib/api/apiInstance';
import { useErrorHandler } from 'react-error-boundary';
import { useForm } from 'react-hook-form';
import { User } from '../../openapi';
import { useNavigate } from 'react-router-dom';
import UserAvatar from '../../components/Avatar/UserAvatar';

const searchUserSchema = z.object({
  keyword: z.string().nonempty(),
});

type FormData = z.infer<typeof searchUserSchema>;

export default function UserSearchComponent() {
  const { register, handleSubmit } = useForm<FormData>({
    resolver: zodResolver(searchUserSchema),
  });
  const errorHandler = useErrorHandler();
  const [users, setUsers] = React.useState<User[]>([]);
  const [userNotFound, setUserNotFound] = React.useState<boolean>(false);
  const navigate = useNavigate();

  const onSubmit = async (data: FormData) => {
    try {
      setUserNotFound(false);
      const api = await apiInstance;
      const res = await api.searchUser(data.keyword);
      if (res.data && Array.isArray(res.data)) {
        const SearchedUsers: User[] = res.data.map((item) => ({
          user_id: item.user_id,
          name: item.name,
          profile_image: item.profile_image,
        }));
        setUsers(SearchedUsers);
        if (res.data.length === 0) {
          setUserNotFound(true);
        }
      }
    } catch (error: unknown) {
      errorHandler(error);
    }
  };

  return (
    <Container component="main">
      <Box
        component="form"
        onSubmit={handleSubmit(onSubmit)}
        noValidate
        sx={{ mt: '1rem' }}
      >
        <Box sx={{ display: 'flex' }}>
          <TextField
            {...register('keyword')}
            fullWidth
            id="keyword"
            label="Search"
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
              },
            }}
          >
            検索
          </Button>
        </Box>
      </Box>
      <Box sx={{ marginTop: '1rem' }}>
        {users.map((user) => (
          <Box
            key={user.user_id}
            onClick={() => navigate(`/user-profile/${user.user_id}`)}
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
      {userNotFound && (
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
