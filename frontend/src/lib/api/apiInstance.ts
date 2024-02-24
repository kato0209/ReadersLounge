import { DefaultApi } from '../../openapi/api';
import { Configuration } from '../../openapi';
import { ResCsrfToken } from '../../openapi/models';

const BASE_API_URL = import.meta.env.VITE_API_URL as string;

if (!BASE_API_URL) {
  throw new Error('Environment variable VITE_API_URL is not set.');
}

const createApiInstance = async (): Promise<DefaultApi> => {
  const config = new Configuration({
    basePath: BASE_API_URL,
    baseOptions: {
      withCredentials: true,
    },
  });

  const tmpApiInstance = new DefaultApi(config);
  const resCsrfToken: ResCsrfToken = {};
  try {
    const results = await tmpApiInstance.csrftoken();
    if (results && results.data && results.data.csrf_token) {
      resCsrfToken.csrf_token = results.data.csrf_token;
      config.apiKey = resCsrfToken.csrf_token;
      return new DefaultApi(config);
    } else {
      throw new Error('Failed to retrieve CSRF token from API.');
    }
  } catch (error) {
    console.error('Error initializing API:', error);
    throw error;
  }
};

export const apiInstance = createApiInstance();
