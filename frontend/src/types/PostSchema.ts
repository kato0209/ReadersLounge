import { z } from 'zod';

const IMAGE_TYPES = ['image/jpeg', 'image/png'];

export const PostSchema = z.object({
    content: z.string().nonempty('投稿内容は必須です').max(255, {
      message: "投稿内容は255文字以内で入力してください"
    }),
    rating: z.number().positive(),
    ISBNcode: z.string().nonempty('本が選択されていません'), 
    postImage: z.instanceof(File).optional()
    .refine((file) => {
      return (
        file === undefined ||
        (IMAGE_TYPES.includes(file.type) &&
          file.name.split('.').pop()?.toLowerCase() !== 'jpg')
      );
    }, {
      message: '.jpegもしくは.pngのみ可能です',
    })
  });