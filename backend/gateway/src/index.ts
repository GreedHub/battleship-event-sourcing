import express, { Request, Response }  from  'express'

const PORT = 5000
const app = express()

app.get('/',(req: Request,res:Response)=>{
    res.send('hello')
})

app.listen(PORT,()=>{console.log(`listening on: http://localhost:${PORT}/`)})