import { FallbackProps } from 'react-error-boundary'
import BackToHomeButton from '../Button/BackToHomeButton'
import { Box } from '@mui/material'

function ErrorFallback({ error }: FallbackProps) {
  return (
    <Box sx={{ 
      display: 'flex', 
      justifyContent: 'center',
      alignItems: 'center',    
      height: '100vh',         
      width: '100vw'           
    }}>
      <Box sx={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center', 
        alignItems: 'center',     
        height: '100vh',          
        width: '50vw',            
      }}>
        <h2>エラーが発生しました</h2>
        <pre style={{ marginBottom: '1.5rem' }}>{error.message}</pre>
        <BackToHomeButton />
      </Box>
    </Box>
  )
}
export default ErrorFallback