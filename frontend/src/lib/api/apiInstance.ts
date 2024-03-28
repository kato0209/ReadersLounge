import { DefaultApi } from '../../openapi/api';
import { Configuration } from '../../openapi';

const BASE_API_URL = process.env.API_URL as string;

if (!BASE_API_URL) {
  throw new Error('Environment variable API_URL is not set.');
}

const createApiInstance = async (): Promise<DefaultApi> => {
  const config = new Configuration({
    basePath: BASE_API_URL,
    baseOptions: {
      withCredentials: true,
    },
  });

  return new DefaultApi(config);
};

export const apiInstance = createApiInstance();
