import RoomList from '../../../../features/chat/RoomList';
import { authenticate } from '../../../../lib/auth/authenticate';
import { redirect } from 'next/navigation';
import { apiInstance } from '../../../../lib/api/apiInstance';
import { ChatRoom, Message } from '../../../../openapi';
import { getAllCookies } from '../../../../utils/getCookies';

export default async function ChatRoomListPage({
  params,
}: {
  params: { id: string };
}) {
  if (!authenticate()) {
    redirect('/login');
  }

  const roomID = Number(params.id);
  const fetchRooms = async (): Promise<ChatRoom[]> => {
    try {
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getChatRooms({ headers: { Cookie: cookie } });
      if (res.data && Array.isArray(res.data)) {
        const fetchedRooms: ChatRoom[] = res.data.map((item) => ({
          room_id: item.room_id,
          target_user_id: item.target_user_id,
          target_user_name: item.target_user_name,
          target_user_profile_image: item.target_user_profile_image,
          last_message: item.last_message,
          last_message_sent_at: item.last_message_sent_at,
        }));
        return fetchedRooms;
      } else {
        return [];
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const fetchMessages = async (roomID: number): Promise<Message[]> => {
    try {
      const cookie = getAllCookies();
      const api = apiInstance;
      const res = await api.getMessages(roomID, {
        headers: { Cookie: cookie },
      });
      if (res.data && Array.isArray(res.data)) {
        const fetchedMessages: Message[] = res.data.map((item) => ({
          message_id: item.message_id,
          user_id: item.user_id,
          content: item.content,
          sent_at: item.sent_at,
        }));
        return fetchedMessages;
      } else {
        return [];
      }
    } catch (error: unknown) {
      return Promise.reject(error);
    }
  };

  const chatRooms = await fetchRooms();
  if (roomID) {
    const messages = await fetchMessages(roomID);
    return (
      <RoomList roomID={roomID} chatRooms={chatRooms} messages={messages} />
    );
  }

  return <RoomList roomID={roomID} chatRooms={chatRooms} messages={[]} />;
}
