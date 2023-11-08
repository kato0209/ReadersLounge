

export const getGoogleAuthUrl = (from: string) => {
    const rootUrl = `https://accounts.google.com/o/oauth2/v2/auth`;

    const apiUrl = import.meta.env.VITE_API_URL as string;
    const googleOauthRedirectPath = import.meta.env.VITE_GOOGLE_OAUTH_REDIRECT_PATH as string;
    const redirectUri = `${apiUrl}${googleOauthRedirectPath}`;
    console.log('redirectUri', redirectUri);
    const options = {
      redirect_uri: redirectUri,
      client_id: import.meta.env.VITE_GOOGLE_OAUTH_CLIENT_ID as string,
      access_type: 'offline',
      response_type: 'code',
      prompt: 'consent',
      scope: [
        'https://www.googleapis.com/auth/userinfo.profile',
        'https://www.googleapis.com/auth/userinfo.email',
      ].join(' '),
      state: from,
    };
  
    const qs = new URLSearchParams(options);
  
    return `${rootUrl}?${qs.toString()}`;
};
  
  