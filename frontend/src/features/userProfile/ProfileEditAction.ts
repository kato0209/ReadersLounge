'use server';
import { z } from 'zod';
import { apiInstance } from '../../lib/api/apiInstance';

export type State = {
  error?: string;
  fieldErrors?: {
    name?: string;
    profileText?: string;
    profileImage?: string;
  };
};

const IMAGE_TYPES = ['image/jpeg', 'image/png'];

export async function profileEdit(
  state: State,
  formData: FormData,
): Promise<State> {
  const ProfileSchema = z.object({
    name: z.string().max(20, {
      message: '名前は20文字以内で入力してください',
    }),
    profileText: z.string().max(255, {
      message: 'プロフィールは255文字以内で入力してください',
    }),
    profileImage: z
      .instanceof(File)
      .optional()
      .refine(
        (file) => {
          return (
            file === undefined ||
            (IMAGE_TYPES.includes(file.type) &&
              file.name.split('.').pop()?.toLowerCase() !== 'jpg')
          );
        },
        {
          message: '.jpegもしくは.pngのみ可能です',
        },
      ),
  });

  const validatedFields = ProfileSchema.safeParse({
    name: formData.get('name'),
    profileText: Number(formData.get('profileText')),
    profileImage: formData.get('profileImage'),
  });

  if (!validatedFields.success) {
    throw validatedFields.error.flatten().fieldErrors;
  }

  const { name, profileText, profileImage } = validatedFields.data;

  try {
    const api = apiInstance;
    if (profileImage) {
      await api.updateUser(name, profileImage, profileText);
    } else {
      await api.updateUser(name, undefined, profileText);
    }
    return {};
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
