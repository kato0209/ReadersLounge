export function isValidUrl(string: string | undefined): boolean {
    if (typeof string !== 'string') {
      return false;
    }
  
    let url: URL;
  
    try {
      url = new URL(string);
    } catch (_) {
      return false;  
    }
  
    return true;
}