import { apiInstance } from '../../lib/api/apiInstance';
import { Message } from '../../openapi';
import RoomClientComponent from './RoomCC';
import { getAllCookies } from '../../utils/getCookies';

type RoomProps = {
  roomID: number;
};

export default async function Room(props: RoomProps) {
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

  const messages = await fetchMessages(props.roomID);

  return <RoomClientComponent roomID={props.roomID} messages={messages} />;
}
