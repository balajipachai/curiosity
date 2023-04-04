const express = require('express')
const app = express()
const redis = require('./redis')


app.get('/main', async (req, res) => {
    const ip = req.headers['x-forwarded-for'] || req.socket.remoteAddress;
    console.log(ip)
    const requests = await redis.incr(ip)
    console.log(`Number of requests made so far ${requests}`);

    if (requests === 1) {
        await redis.expire(ip, 60);
    }
    
    if (requests > 5) {
        res.status(503)
            .json({
                response: 'Error',
                callsMade: requests,
                msg: 'Too many calls made'
            });
    } else 
    res.json('You have successfuly hit main route')
})

app.listen(3000, () => {
    console.log('Server started')
})