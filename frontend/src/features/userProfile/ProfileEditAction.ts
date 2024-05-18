'use server';
import { z } from 'zod';
import { apiInstance } from '../../lib/api/apiInstance';
import { getAllCookies } from '../../utils/getCookies';
import { redirect } from 'next/navigation';

export type State = {
  error?: boolean;
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
      .any()
      .optional()
      .refine(
        (file) => {
          return (
            file.name === 'undefined' ||
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
    profileText: formData.get('profileText'),
    profileImage: formData.get('profileImage'),
  });

  if (validatedFields.success === false) {
    const fieldErrors = validatedFields.error.flatten().fieldErrors;
    return {
      error: true,
      fieldErrors: {
        name: fieldErrors.name ? fieldErrors.name[0] : undefined,
        profileText: fieldErrors.profileText
          ? fieldErrors.profileText[0]
          : undefined,
        profileImage: fieldErrors.profileImage
          ? fieldErrors.profileImage[0]
          : undefined,
      },
    };
  }

  const { name, profileText, profileImage } = validatedFields.data;

  try {
    const cookie = getAllCookies();
    const api = apiInstance;
    let res;
    if (profileImage.name !== 'undefined') {
      res = await api.updateUser(name, profileImage, profileText, {
        headers: { Cookie: cookie },
      });
    } else {
      res = await api.updateUser(name, undefined, profileText, {
        headers: { Cookie: cookie },
      });
    }
    const userID = Number(res.data);
    redirect('/user-profile/' + userID);
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
