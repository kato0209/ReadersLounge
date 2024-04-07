import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { CreateChatRoomRequest } from '../../../openapi';

export async function GET(request: NextRequest): Promise<number> {
  const searchParams = request.nextUrl.searchParams;
  const chatPartnerID = searchParams.get('chatPartnerID');
  if (!chatPartnerID) {
    Promise.reject(new Error('chatPartnerID is required'));
  }
  try {
    const req: CreateChatRoomRequest = {
      chat_partner_id: Number(chatPartnerID),
    };
    const api = apiInstance;
    const res = await api.createChatRoom(req);
    if (res.status === 201 && res.data.room_id) {
      return res.data.room_id;
    } else {
      return Promise.reject(new Error('Failed to create chat room'));
    }
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
