import Login from '../../../features/login/Login';
import GoogleAuth from '../../../components/OAuth/GoogleAuth';
import { apiInstance } from '../../../lib/api/apiInstance';

export default function LoginPage() {
  async function getApiInstance() {
    const api = await apiInstance;
    return api;
  }
  const api = getApiInstance();
  return (
    <>
      <Login GoogleAuth={<GoogleAuth />} ApiInstance={api} />
    </>
  );
}
