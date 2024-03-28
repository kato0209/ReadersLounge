import Login from '../../../features/login/Login';
import GoogleAuth from '../../../components/OAuth/GoogleAuth';

export default function LoginPage() {
  return (
    <>
      <Login GoogleAuth={<GoogleAuth />} />
    </>
  );
}
