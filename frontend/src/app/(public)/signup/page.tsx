import Signup from '../../../features/signup/Signup';
import GoogleAuth from '../../../components/OAuth/GoogleAuth';

export default function SignupPage() {
  return (
    <>
      <Signup GoogleAuth={<GoogleAuth />} />
    </>
  );
}
