export function isValidUrl(string: string | undefined): boolean {
  if (typeof string !== 'string') {
    return false;
  }

  try {
    new URL(string);
  } catch (_) {
    return false;
  }

  return true;
}
