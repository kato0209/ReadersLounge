const OnError = (error: Error, info: { componentStack: string }) => {
  console.log('error.message', error.message);
  console.log('info.componentStack:', info.componentStack);
};

export default OnError;
