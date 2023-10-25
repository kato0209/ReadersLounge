import { FallbackProps } from 'react-error-boundary'
import BackToHomeButton from '../Button/BackToHomeButton'

function ErrorFallback({ error }: FallbackProps) {
  return (
    <>
      <h2>エラーが発生しました</h2>
      <pre style={{ marginBottom: '1.5rem' }}>{error.message}</pre>
      <BackToHomeButton />
    </>
  )
}
export default ErrorFallback