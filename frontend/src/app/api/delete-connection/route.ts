import { apiInstance } from '../../../lib/api/apiInstance';
import { type NextRequest } from 'next/server';
import { NextResponse } from 'next/server';

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
    const api = apiInstance;
    await api.deleteConnection(Number(connectionID));
  } catch (error: unknown) {
    return Promise.reject(error);
  }
  return NextResponse.json({ status: 200 });
}
