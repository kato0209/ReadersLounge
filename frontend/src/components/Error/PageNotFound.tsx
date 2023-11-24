import BackToHomeButton from '../Button/BackToHomeButton'
import { Box } from '@mui/material'

function PageNotFound() {
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
        <h2>指定されたページは存在しません</h2>
        <BackToHomeButton />
      </Box>
    </Box>
  )
}
export default PageNotFound