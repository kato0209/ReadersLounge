export const getGoogleAuthUrl = (state: string) => {
  const rootUrl = process.env.NEXT_PUBLIC_GOOGLE_OAUTH_URL as string;

  const apiUrl = process.env.NEXT_PUBLIC_API_URL as string;
  const googleOauthRedirectPath = process.env
    .NEXT_PUBLIC_GOOGLE_OAUTH_REDIRECT_PATH as string;
  const redirectUri = `${apiUrl}${googleOauthRedirectPath}`;
  const options = {
    redirect_uri: redirectUri,
    client_id: process.env.GOOGLE_OAUTH_CLIENT_ID as string,
    access_type: 'offline',
    response_type: 'code',
    prompt: 'consent',
    scope: [
      process.env.NEXT_PUBLIC_GOOGLE_OAUTH_USER_INFO_EMAIL_URL as string,
      process.env.NEXT_PUBLIC_GOOGLE_OAUTH_USER_INFO_PROFILE_URL as string,
    ].join(' '),
    state: state,
  };

  const qs = new URLSearchParams(options);

  return `${rootUrl}?${qs.toString()}`;
};
