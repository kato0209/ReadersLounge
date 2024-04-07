import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';
import { CreateConnectionRequest } from '../../../openapi';

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const connectionID = searchParams.get('connectionID');
  if (!connectionID) {
    return NextResponse.json(
      { error: 'connectionID is required' },
      { status: 500 },
    );
  }
  try {
    const req: CreateConnectionRequest = {
      target_user_id: Number(connectionID),
    };
    const api = apiInstance;
    await api.createConnection(req);
  } catch (error: unknown) {
    return Promise.reject(error);
  }
  return NextResponse.json({ status: 200 });
}
