import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';
import { CreateChatRoomRequest } from '../../../openapi';
import { getAllCookies } from '../../../utils/getCookies';

export async function GET(request: NextRequest): Promise<NextResponse> {
  const searchParams = request.nextUrl.searchParams;
  const chatPartnerID = searchParams.get('chatPartnerID');
  if (!chatPartnerID) {
    Promise.reject(new Error('chatPartnerID is required'));
  }
  try {
    const req: CreateChatRoomRequest = {
      chat_partner_id: Number(chatPartnerID),
    };
    const cookie = getAllCookies();
    const api = apiInstance;
    const res = await api.createChatRoom(req, { headers: { Cookie: cookie } });
    if (res.status === 201 && res.data.room_id) {
      return NextResponse.json({ data: res.data.room_id }, { status: 201 });
    } else {
      return Promise.reject(new Error('Failed to create chat room'));
    }
  } catch (error: unknown) {
    return Promise.reject(error);
  }
}
