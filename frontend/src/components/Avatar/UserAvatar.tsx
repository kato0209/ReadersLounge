import Avatar from '@mui/material/Avatar';
import { isValidUrl } from '../../utils/isValidUrl';
import Link from 'next/link';

function UserAvatar(props: { image: string; userID: number }) {
  return (
    <Link href={`/user-profile/${props.userID}`}>
      <Avatar
        src={
          isValidUrl(props.image)
            ? props.image
            : `data:image/png;base64,${props.image}`
        }
        onClick={(event) => event.stopPropagation()}
      ></Avatar>
    </Link>
  );
}

export default UserAvatar;
